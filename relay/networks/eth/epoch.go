package eth

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/ambrosus/ambrosus-bridge/relay/pkg/ethash"
)

func (b *Bridge) SetEpochData(epochData ethash.EpochData) error {
	var nodes []*big.Int
	start := big.NewInt(0)

	for i, node := range epochData.MerkleNodes {
		nodes = append(nodes, node)

		if len(nodes) == 40 || i == len(epochData.MerkleNodes) {
			merkelNodesNumber := big.NewInt(int64(len(nodes)))

			if i < 440 && epochData.Epoch.Uint64() == 128 {
				start.Add(start, merkelNodesNumber)
				nodes = []*big.Int{}

				continue
			}

			err := b.sideBridge.SubmitEpochData(
				epochData.Epoch, epochData.FullSizeIn128Resolution, epochData.BranchDepth,
				nodes, start, merkelNodesNumber,
			)
			if err != nil {
				return err
			}

			start.Add(start, merkelNodesNumber)
			nodes = []*big.Int{}
		}
	}

	return nil
}

func (b *Bridge) loadEpochDataFile(epoch uint64) (*ethash.EpochData, error) {
	data, err := os.ReadFile(fmt.Sprintf("./assets/epoch/%d.json", epoch))
	if err != nil {
		return nil, err
	}

	var epochData *ethash.EpochData

	if err := json.Unmarshal(data, &epochData); err != nil {
		return nil, err
	}

	return epochData, nil
}

func (b *Bridge) createEpochDataFile(epoch uint64) (*ethash.EpochData, error) {
	data, err := ethash.GenerateEpochData(epoch)
	if err != nil {
		return nil, err
	}

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(fmt.Sprintf("./assets/epoch/%d.json", epoch), file, 0644); err != nil {
		return nil, err
	}

	return data, nil
}

func (b *Bridge) deleteEpochDataFile(epoch uint64) error {
	return os.Remove(fmt.Sprintf("./assets/epoch/%d.json", epoch))
}
