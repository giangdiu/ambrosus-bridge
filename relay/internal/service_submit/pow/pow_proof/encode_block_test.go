package pow_proof

import (
	"context"
	"math/big"
	"testing"

	"github.com/ambrosus/ambrosus-bridge/relay/pkg/helpers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ab050ca98686478e9e9b06dfc3b2f069")
	if err != nil {
		t.Fatal(err)
	}

	blockOld, err := client.BlockByNumber(context.Background(), big.NewInt(10000))
	if err != nil {
		t.Fatal(err)
	}
	blockNew, err := client.BlockByNumber(context.Background(), big.NewInt(14264072))
	if err != nil {
		t.Fatal(err)
	}

	testEncodeBlock(t, blockOld)
	testEncodeBlock(t, blockNew)

}

func testEncodeBlock(t *testing.T, block *types.Block) {
	b, err := splitBlock(block.Header(), true)
	if err != nil {
		t.Fatal(err)
	}

	rlpWithoutNonce := helpers.BytesConcat(
		b.P0WithoutNonce[:],
		b.P1, b.ParentOrReceiptHash[:],
		b.P2, b.Difficulty,
		b.P3, b.Number,
		b.P4,
		// here was nonce
		b.P6,
	)
	rlpWithNonce := helpers.BytesConcat(
		b.P0WithNonce[:],
		b.P1, b.ParentOrReceiptHash[:],
		b.P2, b.Difficulty,
		b.P3, b.Number,
		b.P4, b.P5,
		b.Nonce, b.P6,
	)

	expectedRlpWithoutNonce, err := headerRlp(block.Header(), false)
	if err != nil {
		t.Fatal(err)
	}

	expectedRlpWithNonce, err := rlp.EncodeToBytes(block.Header())
	if err != nil {
		t.Fatal(err)
	}

	hash := common.BytesToHash(crypto.Keccak256(rlpWithNonce))
	hashTest := common.BytesToHash(crypto.Keccak256(expectedRlpWithNonce))

	assert.Equal(t, block.Hash(), hashTest) // =>  expectedRlpWithNonce ok
	assert.Equal(t, block.Hash(), hash)     // =>  rlpWithNonce ok

	assert.Equal(t, expectedRlpWithNonce, rlpWithNonce)
	assert.Equal(t, expectedRlpWithoutNonce, rlpWithoutNonce)
}
