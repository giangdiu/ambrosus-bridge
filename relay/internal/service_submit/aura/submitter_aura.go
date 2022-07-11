package aura

import (
	"fmt"

	"github.com/ambrosus/ambrosus-bridge/relay/internal/bindings"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/config"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/networks"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_submit"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_submit/aura/aura_proof"
	"github.com/ambrosus/ambrosus-bridge/relay/pkg/ethclients/parity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
)

type SubmitterAura struct {
	networks.Bridge
	auraReceiver service_submit.ReceiverAura
	auraEncoder  *aura_proof.AuraEncoder
	logger       *zerolog.Logger
}

func NewSubmitterAura(bridge networks.Bridge, auraReceiver service_submit.ReceiverAura, cfg *config.SubmitterAura) (*SubmitterAura, error) {
	parityClient := bridge.GetClient().(*parity.Client)

	// Creating a new ambrosus VS contract instance.
	vsContract, err := bindings.NewVs(common.HexToAddress(cfg.VSContractAddr), parityClient)
	if err != nil {
		return nil, fmt.Errorf("create vs contract: %w", err)
	}

	return &SubmitterAura{
		Bridge:       bridge,
		auraReceiver: auraReceiver,
		auraEncoder:  aura_proof.NewAuraEncoder(bridge, auraReceiver, vsContract, parityClient),
		logger:       bridge.GetLogger(), // todo maybe sublogger?
	}, nil
}

func (b *SubmitterAura) SendEvent(event *bindings.BridgeTransfer, safetyBlocks uint64) error {
	auraProof, err := b.auraEncoder.EncodeAuraProof(event, safetyBlocks)
	if err != nil {
		return fmt.Errorf("encodeAuraProof: %w", err)
	}

	b.logger.Info().Str("event_id", event.EventId.String()).Msg("Submit transfer Aura...")
	err = b.auraReceiver.SubmitTransferAura(auraProof)
	if err != nil {
		return fmt.Errorf("SubmitTransferAura: %w", err)
	}
	return nil
}