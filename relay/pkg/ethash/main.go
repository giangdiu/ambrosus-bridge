package ethash

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ambrosus/ambrosus-bridge/relay/pkg/ethash/merkle"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/rs/zerolog"
)

type Ethash struct {
	dir    string
	logger zerolog.Logger

	keepPrevEpochs uint64
	genNextEpochs  uint64

	// cache (in default meaning)
	caches   map[uint64][]byte
	dags     map[uint64][]byte
	dagKLock *Kmutex // dagKLock need to make sure that for each epoch only one dag is generating at a time
	// no lock for cache because is doesn't take that long to generate it
}

func New(dir string, keepPrevEpochs, genNextEpochs uint64, baseLogger zerolog.Logger) *Ethash {
	logger := baseLogger.With().Str("pkg", "ethash").Logger()

	if dir == "" {
		logger.Info().Msg("No ethash dir provided, working in memory only")
	}

	return &Ethash{
		dir:    dir,
		logger: logger,

		keepPrevEpochs: keepPrevEpochs,
		genNextEpochs:  genNextEpochs,

		caches:   map[uint64][]byte{},
		dags:     map[uint64][]byte{},
		dagKLock: NewKmutex(),
	}
}

type EpochData struct {
	Epoch                   *big.Int   `json:"epoch"`
	FullSizeIn128Resolution *big.Int   `json:"full_size_in_128_resolution"`
	BranchDepth             *big.Int   `json:"branch_depth"`
	MerkleNodes             []*big.Int `json:"merkle_nodes"`
}

func (e *Ethash) GenDagForEpoch(epoch uint64) {
	e.getDag(epoch)
}

func (e *Ethash) GetEpochData(epoch uint64) (*EpochData, error) {
	defer e.UpdateCache(epoch)

	e.logger.Debug().Uint64("epoch", epoch).Msg("Getting epoch data")

	mt := merkle.NewDatasetTree()
	fullSize, branchDepth, err := e.populateMerkle(epoch, mt)
	if err != nil {
		return nil, err
	}

	return &EpochData{
		Epoch:                   big.NewInt(int64(epoch)),
		FullSizeIn128Resolution: big.NewInt(int64(fullSize / 128)),
		BranchDepth:             big.NewInt(int64(branchDepth - 10)),
		MerkleNodes:             mt.MerkleNodes(),
	}, nil
}

func (e *Ethash) GetBlockLookups(blockNumber, nonce uint64, hashNoNonce [32]byte) (dataSetLookup, witnessForLookup []*big.Int, err error) {
	ep := epoch(blockNumber)
	defer e.UpdateCache(ep)

	e.logger.Debug().Uint64("blockNumber", blockNumber).Msg("Getting block lookups")

	indices, err := e.getVerificationIndices(ep, hashNoNonce, nonce)
	if err != nil {
		return
	}

	mt := merkle.NewDatasetTree()
	mt.RegisterIndex(indices...)
	_, _, err = e.populateMerkle(ep, mt)
	if err != nil {
		return nil, nil, err
	}

	dataSetLookup, witnessForLookup = mt.Lookups()
	return
}

func (e *Ethash) UpdateCache(currentEpoch uint64) {
	if e.genNextEpochs != 0 {
		e.logger.Debug().Msg("Generating data for next epochs")

		go func() {
			for i := uint64(0); i < e.genNextEpochs; i++ {
				if _, err := os.Stat(e.pathToDag(currentEpoch + i + 1)); errors.Is(err, os.ErrNotExist) {
					_, _ = e.getDag(currentEpoch + i + 1)
					delete(e.dags, currentEpoch+i+1) // we need to keep in memory only the current epoch's DAG
				}
			}
		}()
	}

	ep, of := math.SafeSub(currentEpoch, e.keepPrevEpochs+1)
	if of { // fucking golang, 21 century already -_-
		ep = 0
	}

	e.logger.Debug().Msgf("Deleting data for older epochs older than %d", ep)

	for ; ep > 0; ep-- {
		if e.useFs() {
			_ = os.Remove(e.pathToDag(ep))
			_ = os.Remove(e.pathToCache(ep))
		}
		delete(e.dags, ep)
		delete(e.caches, ep)
	}
}

func (e *Ethash) useFs() bool {
	return e.dir != ""
}
func (e *Ethash) pathToCache(epoch uint64) string {
	return filepath.Join(e.dir, fmt.Sprintf("cache-%d", epoch))
}
func (e *Ethash) pathToDag(epoch uint64) string {
	return filepath.Join(e.dir, fmt.Sprintf("dag-%d", epoch))
}
