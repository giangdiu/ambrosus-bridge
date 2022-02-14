package eth

import (
	"github.com/ambrosus/ambrosus-bridge/relay/contracts"
	"github.com/ambrosus/ambrosus-bridge/relay/helpers"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/rs/zerolog/log"
)

func EncodeBlock(header *types.Header, isEventBlock bool) *contracts.CheckPoWBlockPoW {
	// split rlp encoded header (bytes) by
	// - receiptHash (for event block) / parentHash (for safety block)
	// - Difficulty (for PoW)

	rlpHeader, err := rlp.EncodeToBytes(header)
	if err != nil {
		log.Error().Err(err).Msg("header not encode to bytes")
	}

	splitEls := make([][]byte, 2)
	if isEventBlock {
		splitEls[0] = header.ReceiptHash.Bytes()
	} else {
		splitEls[0] = header.ParentHash.Bytes()
	}

	splitEls[1] = header.Difficulty.Bytes()

	splitted, err := helpers.BytesSplit(rlpHeader, splitEls)
	if err != nil {
		log.Error().Err(err).Msg("rlp header not splited")
	}

	return &contracts.CheckPoWBlockPoW{
		P1:                    splitted[0],
		PrevHashOrReceiptRoot: helpers.BytesToBytes32(splitEls[0]),
		P2:                    splitted[1],
		Difficulty:            splitEls[1],
		P3:                    splitted[2],
	}

}
