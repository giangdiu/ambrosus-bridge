package eth

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ambrosus/ambrosus-bridge/relay/config"
	"github.com/ambrosus/ambrosus-bridge/relay/contracts"
	"github.com/ambrosus/ambrosus-bridge/relay/networks"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
)

type EventNotFoundErr struct {
	EventId *big.Int
}

func (e EventNotFoundErr) Error() string {
	return fmt.Sprintf("event with id %s not found", e.EventId)
}

func NewEventNotFoundErr(eventId *big.Int) *EventNotFoundErr {
	return &EventNotFoundErr{EventId: eventId}
}

// maybe move to helpers pkg
type JsonError interface {
	Error() string
	ErrorCode() int
	ErrorData() interface{}
}

type Bridge struct {
	Client     *ethclient.Client
	Contract   *contracts.Eth
	sideBridge networks.Bridge
	config     *config.Bridge
}

// Creating a new ethereum bridge.
func New(cfg *config.Bridge) (*Bridge, error) {
	// Creating a new ethereum client.
	client, err := ethclient.Dial(cfg.Url)
	if err != nil {
		return nil, err
	}

	// Creating a new ethereum bridge contract instance.
	contract, err := contracts.NewEth(cfg.ContractAddress, client)
	if err != nil {
		return nil, err
	}

	return &Bridge{Client: client, Contract: contract, config: cfg}, nil
}

func (b *Bridge) SubmitTransfer(proof contracts.TransferProof) error {
	var castProof *contracts.CheckAuraAuraProof
	switch proof.(type) {
	case *contracts.CheckAuraAuraProof:
		castProof = proof.(*contracts.CheckAuraAuraProof)
	default:
		return fmt.Errorf("unknown proof type %T, expected %T", proof, castProof)

	}

	// todo

	return nil
}

func (b *Bridge) GetLastEventId() (*big.Int, error) {
	return b.Contract.InputEventId(nil)
}

// Getting contract event by id.
func (b *Bridge) GetEventById(eventId *big.Int) (*contracts.TransferEvent, error) {
	opts := &bind.FilterOpts{Context: context.Background()}

	logs, err := b.Contract.FilterTransfer(opts, []*big.Int{eventId})
	if err != nil {
		return nil, err
	}

	if logs.Next() {
		return &logs.Event.TransferEvent, nil
	}

	return nil, NewEventNotFoundErr(eventId)
}

func (b *Bridge) GetValidatorSet() ([]common.Address, error) {
	return b.Contract.GetValidatorSet(nil)
}

// todo code below may be common for all networks?

func (b *Bridge) Run(sideBridge networks.Bridge) {
	b.sideBridge = sideBridge

	for {
		if err := b.listen(); err != nil {
			log.Error().Err(err).Msg("listen ambrosus error")
		}
	}
}

func (b *Bridge) checkOldEvents() error {
	lastEventId, err := b.sideBridge.GetLastEventId()
	if err != nil {
		return err
	}

	i := big.NewInt(1)
	for {
		nextEventId := big.NewInt(0).Add(lastEventId, i)
		nextEvent, err := b.GetEventById(nextEventId)
		if err != nil {
			var eventNotFoundErr *EventNotFoundErr
			if errors.As(err, &eventNotFoundErr) {
				return nil
			}
			return err
		}

		err = b.sendEvent(nextEvent)
		if err != nil {
			return err
		}

		i = big.NewInt(0).Add(i, big.NewInt(1))
	}
}

func (b *Bridge) listen() error {
	err := b.checkOldEvents()
	if err != nil {
		return err
	}

	// Subscribe to events
	watchOpts := &bind.WatchOpts{Context: context.Background()}
	eventChannel := make(chan *contracts.EthTransfer) // <-- тут я хз как сделать общий(common) тип для канала
	eventSub, err := b.Contract.WatchTransfer(watchOpts, eventChannel, nil)
	if err != nil {
		return err
	}

	defer eventSub.Unsubscribe()

	// main loop
	for {
		select {
		case err := <-eventSub.Err():
			return err
		case event := <-eventChannel:
			if err := b.sendEvent(&event.TransferEvent); err != nil {
				return err
			}
		}
	}
}

func (b *Bridge) sendEvent(event *contracts.TransferEvent) error {
	// Wait for safety blocks.
	safetyBlocks, err := b.getSafetyBlocksNum()
	if err != nil {
		return err
	}

	if err := b.waitForBlock(event.Raw.BlockNumber + safetyBlocks); err != nil {
		return err
	}

	// Check if the event has been removed.
	if err := b.isEventRemoved(event); err != nil {
		return err
	}

	ambTransfer, err := b.getBlocksAndEvents(event)
	if err != nil {
		return err
	}

	return b.sideBridge.SubmitTransfer(ambTransfer)
}

func (b *Bridge) GetReceipts(blockHash common.Hash) ([]*types.Receipt, error) {
	txsCount, err := b.Client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		return nil, err
	}

	receipts := make([]*types.Receipt, txsCount)

	errGroup := new(errgroup.Group)
	for i := uint(0); i < txsCount; i++ {
		i := i // https://golang.org/doc/faq#closures_and_goroutines ¯\_(ツ)_/¯
		errGroup.Go(func() error {
			tx, err := b.Client.TransactionInBlock(context.Background(), blockHash, i)
			if err != nil {
				return err
			}
			receipt, err := b.Client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				return err
			}

			receipts[i] = receipt
			return nil
		})
	}

	err = errGroup.Wait()
	if err != nil {
		return nil, err
	}
	return receipts, nil
}

func (b *Bridge) getAuth() (*bind.TransactOpts, error) {
	chainId, err := b.Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(b.config.PrivateKey, chainId)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (b *Bridge) isEventRemoved(event *contracts.TransferEvent) error {
	block, err := b.Client.BlockByNumber(context.Background(), big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return err
	}

	if block.Hash() != event.Raw.BlockHash {
		return fmt.Errorf("block hash != event's block hash")
	}
	return nil
}

func (b *Bridge) waitForBlock(targetBlockNum uint64) error {
	// todo maybe timeout (context)
	blockChannel := make(chan *types.Header)
	blockSub, err := b.Client.SubscribeNewHead(context.Background(), blockChannel)
	if err != nil {
		return err
	}

	currentBlockNum, err := b.Client.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	for currentBlockNum < targetBlockNum {
		select {
		case err := <-blockSub.Err():
			return err

		case block := <-blockChannel:
			currentBlockNum = block.Number.Uint64()
		}
	}

	return nil
}

func (b *Bridge) getSafetyBlocksNum() (uint64, error) {
	safetyBlocks, err := b.Contract.MinSafetyBlocks(nil)
	if err != nil {
		return 0, err
	}
	return safetyBlocks.Uint64(), nil
}

// maybe move the functions below to helpers pkg
func getFailureReason(client *ethclient.Client, from common.Address, tx *types.Transaction) error {
	_, err := client.CallContract(context.Background(), ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}, nil)

	return err
}

func getJsonErrData(err error) string {
	var jsonErr = err.(JsonError)
	return jsonErr.ErrorData().(string)
}
