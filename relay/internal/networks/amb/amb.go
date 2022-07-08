package amb

import (
	"fmt"

	"github.com/ambrosus/ambrosus-bridge/relay/internal/bindings"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/config"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/networks"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/networks/amb/aura_proof"
	nc "github.com/ambrosus/ambrosus-bridge/relay/internal/networks/common"
	"github.com/ambrosus/ambrosus-bridge/relay/pkg/ethclients/parity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
)

const BridgeName = "ambrosus"

type Bridge struct {
	nc.CommonBridge
	ParityClient *parity.Client
	vSContract   *bindings.Vs
	sideBridge   networks.BridgeReceiveAura

	auraEncoder *aura_proof.AuraEncoder
}

// New creates a new ambrosus bridge.
func New(cfg *config.AMBConfig, baseLogger zerolog.Logger) (*Bridge, error) {
	commonBridge, err := nc.New(cfg.Network, BridgeName)
	if err != nil {
		return nil, fmt.Errorf("create commonBridge: %w", err)
	}

	commonBridge.Logger = baseLogger.With().Str("bridge", BridgeName).Logger()

	// ///////////////////

	parityClient, err := parity.Dial(cfg.HttpURL)
	if err != nil {
		return nil, fmt.Errorf("dial http: %w", err)
	}
	commonBridge.Client = parityClient

	// Creating a new bridge contract instance.
	commonBridge.Contract, err = bindings.NewBridge(common.HexToAddress(cfg.ContractAddr), commonBridge.Client)
	if err != nil {
		return nil, fmt.Errorf("create contract http: %w", err)
	}

	// Create websocket instances if wsUrl provided
	if cfg.WsURL != "" {
		commonBridge.WsClient, err = parity.Dial(cfg.WsURL)
		if err != nil {
			return nil, fmt.Errorf("dial ws: %w", err)
		}

		commonBridge.WsContract, err = bindings.NewBridge(common.HexToAddress(cfg.ContractAddr), commonBridge.WsClient)
		if err != nil {
			return nil, fmt.Errorf("create contract ws: %w", err)
		}
	}

	// Creating a new ambrosus VS contract instance.
	vsContract, err := bindings.NewVs(common.HexToAddress(cfg.VSContractAddr), commonBridge.Client)
	if err != nil {
		return nil, fmt.Errorf("create vs contract: %w", err)
	}

	b := &Bridge{
		CommonBridge: commonBridge,
		ParityClient: parityClient,
		vSContract:   vsContract,
	}
	b.CommonBridge.Bridge = b
	return b, nil
}

func (b *Bridge) SetSideBridge(sideBridge networks.BridgeReceiveAura) {
	b.sideBridge = sideBridge
	b.CommonBridge.SideBridge = sideBridge

	// todo refactor: move to separate service along with vsContract creation
	b.auraEncoder = aura_proof.NewAuraEncoder(b, sideBridge, b.vSContract, b.ParityClient)
}

func (b *Bridge) Run() {
	b.Logger.Debug().Msg("Running ambrosus bridge...")

	go b.UnlockTransfersLoop()
	go b.TriggerTransfersLoop()
	b.SubmitTransfersLoop()
}

func (b *Bridge) SendEvent(event *bindings.BridgeTransfer, safetyBlocks uint64) error {
	auraProof, err := b.auraEncoder.EncodeAuraProof(event, safetyBlocks)
	if err != nil {
		return fmt.Errorf("encodeAuraProof: %w", err)
	}

	b.Logger.Info().Str("event_id", event.EventId.String()).Msg("Submit transfer Aura...")
	err = b.sideBridge.SubmitTransferAura(auraProof)
	if err != nil {
		return fmt.Errorf("SubmitTransferAura: %w", err)
	}
	return nil
}
