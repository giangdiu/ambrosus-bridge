package common

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ambrosus/ambrosus-bridge/relay/internal/config"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/contracts"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/networks"
	"github.com/ambrosus/ambrosus-bridge/relay/pkg/helpers"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
)

type CommonBridge struct {
	networks.Bridge
	Client     *ethclient.Client
	WsClient   *ethclient.Client
	Contract   *contracts.Bridge
	WsContract *contracts.Bridge
	Auth       *bind.TransactOpts
	SideBridge networks.Bridge
	Logger     zerolog.Logger
	Name       string
}

func New(cfg config.Network, name string) (b CommonBridge, err error) {
	b.Name = name

	b.Client, err = ethclient.Dial(cfg.HttpURL)
	if err != nil {
		return b, fmt.Errorf("dial http: %w", err)
	}

	// Creating a new bridge contract instance.
	b.Contract, err = contracts.NewBridge(common.HexToAddress(cfg.ContractAddr), b.Client)
	if err != nil {
		return b, fmt.Errorf("create contract http: %w", err)
	}

	// Create websocket instances if wsUrl provided
	if cfg.WsURL != "" {
		b.WsClient, err = ethclient.Dial(cfg.WsURL)
		if err != nil {
			return b, fmt.Errorf("dial ws: %w", err)
		}

		b.WsContract, err = contracts.NewBridge(common.HexToAddress(cfg.ContractAddr), b.WsClient)
		if err != nil {
			return b, fmt.Errorf("create contract ws: %w", err)
		}
	}

	// create auth if privateKey provided
	if cfg.PrivateKey != "" {
		pk, err := parsePK(cfg.PrivateKey)
		if err != nil {
			return b, fmt.Errorf("parse private key: %w", err)
		}
		chainId, err := b.Client.ChainID(context.Background())
		if err != nil {
			return b, fmt.Errorf("chain id: %w", err)
		}
		b.Auth, err = bind.NewKeyedTransactorWithChainID(pk, chainId)
		if err != nil {
			return b, fmt.Errorf("new keyed transactor: %w", err)
		}

		// update metrics
		b.SetRelayBalanceMetric()
	} else {
		b.Logger.Info().Msg("No private key provided")
	}

	return b, nil

}

// GetLastReceivedEventId get last event id submitted in this contract.
func (b *CommonBridge) GetLastReceivedEventId() (*big.Int, error) {
	return b.Contract.InputEventId(nil)
}

// GetEventById get `Transfer` event (emitted by this contract) by id.
func (b *CommonBridge) GetEventById(eventId *big.Int) (*contracts.BridgeTransfer, error) {
	logs, err := b.Contract.FilterTransfer(nil, []*big.Int{eventId})
	if err != nil {
		return nil, fmt.Errorf("filter transfer: %w", err)
	}
	for logs.Next() {
		if !logs.Event.Raw.Removed {
			return logs.Event, nil
		}
	}
	return nil, networks.ErrEventNotFound
}

func (b *CommonBridge) GetMinSafetyBlocksNum(opts *bind.CallOpts) (uint64, error) {
	safetyBlocks, err := b.Contract.MinSafetyBlocks(opts)
	if err != nil {
		return 0, err
	}
	return safetyBlocks.Uint64(), nil
}

func (b *CommonBridge) ProcessTx(params networks.GetTxErrParams) error {
	if err := b.Bridge.GetTxErr(params); err != nil {
		return err
	}

	b.IncTxCountMetric(params.MethodName)

	b.Logger.Info().
		Str("method", params.MethodName).
		Str("tx_hash", params.Tx.Hash().Hex()).
		Interface("full_tx", params.Tx).
		Interface("tx_params", params.TxParams).
		Msgf("Wait the tx to be mined...")

	// TODO: extract to a separate method
	var receipt *types.Receipt
	err := retry.Do(
		func() error {
			var err error

			ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
			defer cancel()

			receipt, err = bind.WaitMined(ctx, b.Client, params.Tx)
			if err != nil {
				return err
			}
			return nil
		},

		retry.RetryIf(func(err error) bool {
			if errors.Is(err, context.DeadlineExceeded) {
				return true
			}
			return false
		}),
		retry.OnRetry(func(n uint, err error) {
			b.Logger.Warn().
				Str("method", params.MethodName).
				Str("tx_hash", params.Tx.Hash().Hex()).
				// Interface("full_tx", params.Tx).
				// Interface("tx_params", params.TxParams).
				Msgf("Timeout waiting for tx to be mined, trying again... (%d/%d)", n+1, 2)
		}),
		retry.Attempts(2),
		retry.LastErrorOnly(true),
	)
	if err != nil {
		return fmt.Errorf("wait mined: %w", err)
	}

	b.SetUsedGasMetric(params.MethodName, receipt.GasUsed, params.Tx.GasPrice())

	if receipt.Status != types.ReceiptStatusSuccessful {
		b.IncFailedTxCountMetric(params.MethodName)
		err = b.GetFailureReason(params.Tx)
		if err != nil {
			return fmt.Errorf("tx %s failed: %w", params.Tx.Hash().Hex(), helpers.ParseError(err))
		}
		b.Logger.Debug().Err(err).Str("tx_hash", params.Tx.Hash().Hex()).Msg("Tx has been mined but failed :(")
	}
	b.Logger.Debug().Str("tx_hash", params.Tx.Hash().Hex()).Msg("Tx has been mined successfully!")

	return nil
}
