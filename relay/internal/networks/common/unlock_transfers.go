package common

import (
	"context"
	"fmt"
	"time"

	"github.com/ambrosus/ambrosus-bridge/relay/internal/networks"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func (b *CommonBridge) UnlockTransfersLoop() {
	b.shouldHavePk()
	for {
		b.EnsureContractUnpaused()

		if err := b.unlockOldTransfers(); err != nil {
			b.Logger.Error().Msgf("UnlockTransfersLoop: %s", err)
		}
		time.Sleep(failSleepTIme)
	}
}

func (b *CommonBridge) unlockOldTransfers() error {
	// Get oldest transfer timestamp.
	oldestLockedEventId, err := b.Contract.OldestLockedEventId(nil)
	if err != nil {
		return fmt.Errorf("get oldest locked event id: %w", err)
	}
	lockedTransferTime, err := b.Contract.LockedTransfers(nil, oldestLockedEventId)
	if err != nil {
		return fmt.Errorf("get locked transfer time %v: %w", oldestLockedEventId, err)
	}
	if lockedTransferTime.Uint64() == 0 {
		lockTime, err := b.Contract.LockTime(nil)
		if err != nil {
			return fmt.Errorf("get lock time: %w", err)
		}

		b.Logger.Debug().Str("event_id", oldestLockedEventId.String()).Msgf(
			"unlockOldTransfers: there are no locked transfers with that id. Sleep %v seconds...",
			lockTime.Uint64(),
		)
		time.Sleep(time.Duration(lockTime.Uint64()) * time.Second)
		return nil
	}

	// Get the latest block.
	latestBlock, err := b.Client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("get latest block: %w", err)
	}

	// Check if the unlocking is allowed and get the sleep time.
	sleepTime := lockedTransferTime.Int64() - int64(latestBlock.Time())
	if sleepTime > 0 {
		b.Logger.Debug().Str("event_id", oldestLockedEventId.String()).Msgf(
			"unlockOldTransfers: sleep %v seconds...",
			sleepTime,
		)
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}

	// Unlock the oldest transfer.
	b.Logger.Info().Str("event_id", oldestLockedEventId.String()).Msg("unlockOldTransfers: unlocking...")
	err = b.unlockTransfers()
	if err != nil {
		return fmt.Errorf("unlock locked transfer %v: %w", oldestLockedEventId, err)
	}
	b.Logger.Info().Str("event_id", oldestLockedEventId.String()).Msg("unlockOldTransfers: unlocked")
	return nil
}

func (b *CommonBridge) unlockTransfers() error {
	// Make tx without sending it for getting the gas limit.
	authNoSend := *b.Auth
	authNoSend.NoSend = true
	tx, err := b.Contract.UnlockTransfersBatch(&authNoSend)
	if err = b.GetTxErr(networks.GetTxErrParams{Tx: tx, TxErr: err, MethodName: "unlockTransfersBatch"}); err != nil {
		return fmt.Errorf("NoSend: %w", err)
	}

	// Send the tx with the gas limit 20% more than the estimated gas limit.
	customGas := uint64(float64(tx.Gas()) * 1.20) // todo: make the multipler configurable
	authCustomGas := *b.Auth
	authCustomGas.GasLimit = customGas
	return b.ProcessTx(func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return b.Contract.UnlockTransfersBatch(&authCustomGas)
	}, networks.GetTxErrParams{MethodName: "unlockTransfersBatch"})
}
