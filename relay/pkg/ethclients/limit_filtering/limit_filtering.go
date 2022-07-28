package limit_filtering

import (
	"context"
	"math/big"
	"time"

	common_ethclient "github.com/ambrosus/ambrosus-bridge/relay/pkg/ethclients/common"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	common_ethclient.Client
	c *rpc.Client

	defaultFilterLogsFromBlock int64
}

// Dial connects a client to the given URL.
func Dial(rawurl string, defaultFilterLogsFromBlock int64) (*Client, error) {
	return DialContext(context.Background(), rawurl, defaultFilterLogsFromBlock)
}

func DialContext(ctx context.Context, rawurl string, defaultFilterLogsFromBlock int64) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c, defaultFilterLogsFromBlock), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client, defaultFilterLogsFromBlock int64) (client *Client) {
	return &Client{
		Client:                     *common_ethclient.NewClient(c),
		c:                          c,
		defaultFilterLogsFromBlock: defaultFilterLogsFromBlock,
	}
}

// FilterLogs executes a filter query.
func (ec *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	var limit = int64(4999)
	var result []types.Log

	if q.FromBlock.Cmp(big.NewInt(0)) == 0 {
		q.FromBlock = big.NewInt(ec.defaultFilterLogsFromBlock)
	}

	if q.ToBlock == nil {
		currBlockNum, err := ec.BlockNumber(ctx)
		if err != nil {
			return nil, err
		}

		q.ToBlock = big.NewInt(int64(currBlockNum))
	}

	for offset := int64(0); ; offset += limit {
		fromBlock := new(big.Int).Add(q.FromBlock, big.NewInt(offset))
		toBlock := new(big.Int).Add(fromBlock, big.NewInt(limit))
		offset += 1

		if toBlock.Cmp(q.ToBlock) > 0 {
			toBlock = q.ToBlock
		}

		if fromBlock.Cmp(toBlock) >= 0 {
			break
		}

		editedQuery := q
		editedQuery.FromBlock = fromBlock
		editedQuery.ToBlock = toBlock

		logs, err := ec.Client.FilterLogs(ctx, editedQuery)
		if err != nil {
			return nil, err
		}

		result = append(result, logs...)
		time.Sleep(1) // don't spam the node
	}

	return result, nil
}