package fee

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ambrosus/ambrosus-bridge/relay/internal/config"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/networks"
	"github.com/ambrosus/ambrosus-bridge/relay/pkg/helpers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
)

type BridgeFee struct {
	networks.Bridge

	minBridgeFee       decimal.Decimal
	defaultTransferFee decimal.Decimal

	wrapperAddress common.Address
	privateKey     *ecdsa.PrivateKey

	transferFeeTracker *transferFeeTracker
}

func NewBridgeFee(bridge, sideBridge networks.Bridge, cfg config.FeeApiNetwork) (*BridgeFee, error) {
	wrapperAddress, err := bridge.GetContract().WrapperAddress(nil)
	if err != nil {
		return nil, err
	}

	transferFee, err := newTransferFeeTracker(bridge, sideBridge)
	if err != nil {
		return nil, err
	}

	privateKey, err := helpers.ParsePK(cfg.PrivateKey)
	if err != nil {
		return nil, err
	}

	defaultTransferFee, err := decimal.NewFromString(cfg.DefaultTransferFee)
	if err != nil {
		return nil, fmt.Errorf("failed to parse defaultTransferFee (%s)", cfg.DefaultTransferFee)
	}

	return &BridgeFee{
		Bridge:             bridge,
		minBridgeFee:       decimal.NewFromFloat(cfg.MinBridgeFee),
		defaultTransferFee: defaultTransferFee,
		privateKey:         privateKey,
		wrapperAddress:     wrapperAddress,
		transferFeeTracker: transferFee,
	}, nil
}

func (b *BridgeFee) Sign(digestHash []byte) ([]byte, error) {
	return crypto.Sign(digestHash, b.privateKey)
}

func (b *BridgeFee) GetTransferFee() *big.Int {
	return b.transferFeeTracker.GasPerWithdraw()
}

func (b *BridgeFee) GetWrapperAddress() common.Address {
	return b.wrapperAddress
}

func (b *BridgeFee) GetMinBridgeFee() decimal.Decimal {
	return b.minBridgeFee
}

func (b *BridgeFee) GetDefaultTransferFee() decimal.Decimal {
	return b.defaultTransferFee
}
