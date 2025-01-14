package aura_proof

import (
	"context"
	"fmt"

	c "github.com/ambrosus/ambrosus-bridge/relay/internal/bindings"
	"github.com/ambrosus/ambrosus-bridge/relay/pkg/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type vsChangeInBlock struct {
	c.CheckAuraValidatorSetProof                     // to be sent to the bridge
	finalizedBlock               uint64              // block number when this event was finalized
	lastEvent                    *c.VsInitiateChange // last event in this block, used to generate proof
}

func (e *AuraEncoder) getVsChanges(toBlock uint64) ([]*vsChangeInBlock, error) {
	// note: lastProcessedBlock *should* be a block when latest *finalized* `VSChangeEvent` is *emitted*

	lastProcessedBlockNum, err := e.getLastProcessedBlockNum()
	if err != nil {
		return nil, fmt.Errorf("getLastProcessedBlockNum: %w", err)
	}

	initialValidatorSet, err := e.auraReceiver.GetValidatorSet()
	if err != nil {
		return nil, fmt.Errorf("GetValidatorSet: %w", err)
	}

	vsChangeEvents, err := e.fetchVSChangeEvents(lastProcessedBlockNum+1, toBlock)
	if err != nil {
		return nil, fmt.Errorf("fetchVSChangeEvents: %w", err)
	}

	blockToEvents, err := e.preprocessVSChangeEvents(initialValidatorSet, vsChangeEvents)
	if err != nil {
		return nil, fmt.Errorf("preprocessVSChangeEvents: %w", err)
	}

	err = e.findWhenFinalize(lastProcessedBlockNum, blockToEvents)
	if err != nil {
		return nil, fmt.Errorf("findWhenFinalize: %w", err)
	}

	// delete events that finalized after `toBlock`
	// otherwise, next time we can miss some events emitted after `toBlock` and before current highest finalized
	filterBlocks(blockToEvents, toBlock)

	vsChanges := helpers.SortedValues(blockToEvents)

	return vsChanges, nil
}

func filterBlocks(blockToEvents map[uint64]*vsChangeInBlock, toBlock uint64) {
	for blockNum, event := range blockToEvents {
		if event.finalizedBlock >= toBlock {
			delete(blockToEvents, blockNum)
		}
	}
}

func (e *AuraEncoder) findWhenFinalize(lastProcessedBlockNum uint64, blockToEvents map[uint64]*vsChangeInBlock) error {
	currentEpoch, err := e.finalizeService.GetBlockWhenFinalize(lastProcessedBlockNum)
	if err != nil {
		return fmt.Errorf("getBlockWhenFinalize: %w", err)
	}

	eventBlockNums := helpers.SortedKeys(blockToEvents)

	for i, eventBlockNum := range eventBlockNums {
		event := blockToEvents[eventBlockNum]

		// current event implicitly finalized, but we don't know in which block.
		// so, it's easier to pretend that this event doesn't exist.
		// first explicit finalized event will use `changes` from this event
		if event.EventBlock <= currentEpoch {
			if i+1 < len(eventBlockNums) {
				// save this event `changes` before the next event `changes`
				nextEvent := blockToEvents[eventBlockNums[i+1]]
				nextEvent.Changes = append(event.Changes, nextEvent.Changes...)
			}
			delete(blockToEvents, eventBlockNum)
			e.logger.Trace().Uint64("block", eventBlockNum).Msg("aura implicitly finalized event block")

		} else {
			currentEpoch, err = e.finalizeService.GetBlockWhenFinalize(eventBlockNum)
			if err != nil {
				return fmt.Errorf("getBlockWhenFinalize: %w", err)
			}
			event.finalizedBlock = currentEpoch

			e.logger.Trace().Uint64("block", eventBlockNum).Uint64("finalized", event.finalizedBlock).Msg("aura finalized event block")
		}
	}
	return nil
}

func (e *AuraEncoder) preprocessVSChangeEvents(initialValidatorSet []common.Address, events []*c.VsInitiateChange) (map[uint64]*vsChangeInBlock, error) {
	blocksToEvents := map[uint64]*vsChangeInBlock{}

	prevSet := initialValidatorSet

	for _, event := range events {
		address, index, err := deltaVS(prevSet, event.NewSet)
		if err != nil {
			return nil, fmt.Errorf("deltaVS: %w", err)
		}
		vsChange := c.CheckAuraValidatorSetChange{DeltaAddress: address, DeltaIndex: index}
		prevSet = event.NewSet

		if _, ok := blocksToEvents[event.Raw.BlockNumber]; !ok {
			blocksToEvents[event.Raw.BlockNumber] = &vsChangeInBlock{CheckAuraValidatorSetProof: c.CheckAuraValidatorSetProof{EventBlock: event.Raw.BlockNumber}}
		}
		blocksToEvents[event.Raw.BlockNumber].Changes = append(blocksToEvents[event.Raw.BlockNumber].Changes, vsChange)
		blocksToEvents[event.Raw.BlockNumber].lastEvent = event
	}

	return blocksToEvents, nil
}

func (e *AuraEncoder) fetchVSChangeEvents(start, end uint64) ([]*c.VsInitiateChange, error) {
	opts := &bind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: context.Background(),
	}

	logs, err := e.vsContract.FilterInitiateChange(opts, nil)
	if err != nil {
		return nil, fmt.Errorf("filter initiate changes: %w", err)
	}

	var res []*c.VsInitiateChange
	for logs.Next() {
		res = append(res, logs.Event)
	}

	return res, nil
}

func (e *AuraEncoder) getLastProcessedBlockNum() (uint64, error) {
	blockHash, err := e.auraReceiver.GetLastProcessedBlockHash()
	if err != nil {
		return 0, fmt.Errorf("GetLastProcessedBlockHash: %w", err)
	}

	block, err := e.bridge.GetClient().BlockByHash(context.Background(), *blockHash)
	if err != nil {
		return 0, fmt.Errorf("get rfBlock by hash: %w", err)
	}

	// last processed block is *parentHash* of block in which last processed event was *emitted*
	// so, processed event emitted at `start+1` block num
	return block.Number().Uint64() + 1, nil
}

func deltaVS(prev, curr []common.Address) (common.Address, uint16, error) {
	d := len(curr) - len(prev)
	if d != 1 && d != -1 {
		return common.Address{}, 0, fmt.Errorf("delta has more (or less) than 1 change")
	}

	for i, prevEl := range prev {
		if i >= len(curr) { // deleted at the end
			return prev[i], uint16(i + 1), nil
		}

		if curr[i] != prevEl {
			if d == 1 { // added
				return curr[i], 0, nil
			} else { // deleted
				return prev[i], uint16(i + 1), nil
			}
		}
	}

	// add at the end
	i := len(curr) - 1
	return curr[i], 0, nil

	// return common.Address{}, 0, fmt.Errorf("this error shouln't exist")
}

// todo if set after applying some changes equal to initialSet => this changes can be skipped
func optimizeVsChanges(initialSet []common.Address, changes []c.CheckAuraValidatorSetChange) {
}
