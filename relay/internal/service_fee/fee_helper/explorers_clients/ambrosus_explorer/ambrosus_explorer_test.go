//go:build !ci

package ambrosus_explorer

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ambrosus/ambrosus-bridge/relay/internal/service_fee/fee_helper/explorers_clients"
	"github.com/stretchr/testify/assert"
)

func TestGetTxs(t *testing.T) {
	explorer, err := NewAmbrosusExplorer("https://explorer-api.ambrosus-dev.io", nil)
	if err != nil {
		t.Fatal(err)
	}
	r, err := explorer.TxListByFromToAddresses("0x295C2707319ad4BecA6b5bb4086617fD6F240CfE", "0xf7E15b720867747a536137f4EFdAB4309225f8D6", explorers_clients.TxFilters{0, nil})
	if err != nil {
		t.Fatal(err)
	}

	jr, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(jr))
}

func TestGetTxsByFromList(t *testing.T) {
	explorer, err := NewAmbrosusExplorer("https://explorer-api.ambrosus-dev.io", nil)
	if err != nil {
		t.Fatal(err)
	}

	fromList := []string{"0xD693a3cc5686e74Ca2e72e8120A2F2013B8eE66E", "0x295C2707319ad4BecA6b5bb4086617fD6F240CfE"}
	r, err := explorer.TxListByFromListToAddresses(fromList, "0xf7E15b720867747a536137f4EFdAB4309225f8D6", explorers_clients.TxFilters{0, nil})
	if err != nil {
		t.Fatal(err)
	}

	for _, tx := range r {
		assert.Contains(t, fromList, tx.From)
	}
}

func TestGetTxsByFromListAndFromBlock(t *testing.T) {
	explorer, err := NewAmbrosusExplorer("https://explorer-api.ambrosus-dev.io", nil)
	if err != nil {
		t.Fatal(err)
	}

	fromBlock := uint64(800_000)
	fromList := []string{"0xD693a3cc5686e74Ca2e72e8120A2F2013B8eE66E", "0x295C2707319ad4BecA6b5bb4086617fD6F240CfE"}
	r, err := explorer.TxListByFromListToAddresses(fromList, "0xf7E15b720867747a536137f4EFdAB4309225f8D6", explorers_clients.TxFilters{fromBlock, nil})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(r))

	for _, tx := range r {
		assert.Contains(t, fromList, tx.From)
		assert.GreaterOrEqual(t, tx.BlockNumber, fromBlock)
	}
}
