package main

import (
	"github.com/ambrosus/ambrosus-bridge/relay/internal/config"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/metric"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/networks"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/networks/amb"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/networks/bsc"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/networks/eth"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_fee/fee_api"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_fee/fee_helper"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_fee/fee_helper/explorers_clients/ambrosus_explorer"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_fee/fee_helper/explorers_clients/etherscan"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_submit"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_submit/aura"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_submit/posa"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_submit/pow"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_trigger"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_unlock"
	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_watchdog"
	"github.com/rs/zerolog"
)

func runSubmitters(cfg *config.Submitters, ambBridge *amb.Bridge, sideBridge service_submit.Receiver, logger zerolog.Logger) {
	logger.Info().Str("service", "submitter").Bool("enabled", cfg.Enable).Send()
	if !cfg.Enable {
		return
	}

	auraSubmitter, err := aura.NewSubmitterAura(ambBridge, &aura.ReceiverAura{Receiver: sideBridge}, cfg.Aura)
	if err != nil {
		logger.Fatal().Err(err).Msg("auraBridgeSubmitter don't created")
	}

	var sideBridgeSubmitter service_submit.Submitter
	switch sideBridge.(type) {
	case *eth.Bridge:
		sideBridgeSubmitter, err = pow.NewSubmitterPoW(sideBridge, &pow.ReceiverPoW{Receiver: ambBridge}, cfg.Pow)
	case *bsc.Bridge:
		sideBridgeSubmitter, err = posa.NewSubmitterPoSA(sideBridge, &posa.ReceiverPoSA{Receiver: ambBridge})
	}
	if err != nil {
		logger.Fatal().Err(err).Msg("sideBridgeSubmitter don't created")
	}

	go service_submit.NewSubmitTransfers(auraSubmitter, sideBridge).Run()
	go service_submit.NewSubmitTransfers(sideBridgeSubmitter, ambBridge).Run()
}

func runWatchdogs(cfg *config.Watchdogs, ambBridge *amb.Bridge, sideBridge networks.Bridge, logger zerolog.Logger) {
	logger.Info().Str("service", "watchdog").Bool("enabled", cfg.Enable).Send()
	if !cfg.Enable {
		return
	}

	go service_watchdog.NewWatchTransfersValidity(ambBridge, sideBridge.GetContract()).Run()
	go service_watchdog.NewWatchTransfersValidity(sideBridge, ambBridge.GetContract()).Run()
}

func runUnlockers(cfg *config.Unlockers, ambBridge *amb.Bridge, sideBridge networks.Bridge, logger zerolog.Logger) {
	logger.Info().Str("service", "watchdog").Bool("enabled", cfg.Enable).Send()
	if !cfg.Enable {
		return
	}

	ambWatchdog := service_watchdog.NewWatchTransfersValidity(ambBridge, sideBridge.GetContract())
	sideWatchdog := service_watchdog.NewWatchTransfersValidity(sideBridge, ambBridge.GetContract())
	go service_unlock.NewUnlockTransfers(ambBridge, ambWatchdog).Run()
	go service_unlock.NewUnlockTransfers(sideBridge, sideWatchdog).Run()
}

func runTriggers(cfg *config.Triggers, ambBridge *amb.Bridge, sideBridge networks.Bridge, logger zerolog.Logger) {
	logger.Info().Str("service", "triggers").Bool("enabled", cfg.Enable).Send()
	if !cfg.Enable {
		return
	}

	go service_trigger.NewTriggerTransfers(ambBridge).Run()
	go service_trigger.NewTriggerTransfers(sideBridge).Run()
}

func runFeeApi(cfg *config.FeeApi, ambBridge, sideBridge networks.Bridge, logger zerolog.Logger) {
	logger.Info().Str("service", "fee api").Bool("enabled", cfg.Enable).Send()
	if !cfg.Enable {
		return
	}

	explorerAmb, err := ambrosus_explorer.NewAmbrosusExplorer(cfg.Amb.ExplorerURL, nil)
	if err != nil {
		logger.Fatal().Err(err).Msg("explorerAmb not created")
	}
	explorerSide, err := etherscan.NewEtherscan(cfg.Side.ExplorerURL, nil)
	if err != nil {
		logger.Fatal().Err(err).Msg("explorerSide not created")
	}

	feeAmb, err := fee_helper.NewFeeHelper(ambBridge, sideBridge, explorerAmb, explorerSide, cfg.Amb, cfg.Side)
	if err != nil {
		logger.Fatal().Err(err).Msg("feeAmb not created")
	}
	feeSide, err := fee_helper.NewFeeHelper(sideBridge, ambBridge, explorerSide, explorerAmb, cfg.Side, cfg.Amb)
	if err != nil {
		logger.Fatal().Err(err).Msg("feeSide not created")
	}

	feeApi := fee_api.NewFeeAPI(feeAmb, feeSide, logger)
	feeApi.Run(cfg.Endpoint, cfg.Ip, cfg.Port)
}

func runPrometheus(cfg *config.Prometheus, logger zerolog.Logger) {
	logger.Info().Str("service", "prometheus").Bool("enabled", cfg.Enable).Send()
	if !cfg.Enable {
		return
	}

	if err := metric.ServeEndpoint(cfg.Ip, cfg.Port); err != nil {
		logger.Fatal().Err(err).Msg("failed to serve HTTP server (Prometheus endpoint)")
	}
}
