package enc2sol

import (
	"bytes"
	"encoding/json"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
	"net/http"
	"relay/enc2sol/mytrie"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}
type Response struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  Header `json:"result"`
	Id      int    `json:"id"`
}

type Header struct {
	ParentHash  *common.Hash    `json:"parentHash"`
	UncleHash   *common.Hash    `json:"sha3Uncles"`
	Coinbase    *common.Address `json:"author"`
	Root        *common.Hash    `json:"stateRoot"`
	TxHash      *common.Hash    `json:"transactionsRoot"`
	ReceiptHash *common.Hash    `json:"receiptsRoot"`
	Bloom       *types.Bloom    `json:"logsBloom"`
	Difficulty  *hexutil.Big    `json:"difficulty"`
	Number      *hexutil.Big    `json:"number"`
	GasLimit    *hexutil.Uint64 `json:"gasLimit"`
	GasUsed     *hexutil.Uint64 `json:"gasUsed"`
	Time        *hexutil.Uint64 `json:"timestamp"`
	Extra       *hexutil.Bytes  `json:"extraData"`

	SealFields []string `json:"sealFields"`
	Signature  *string  `json:"signature"`
}

func newRequest(id int, method string, params []interface{}) *Request {
	return &Request{
		Jsonrpc: "2.0",
		Id:      id,
		Method:  method,
		Params:  params,
	}
}

func HeaderByNumber(number *big.Int) (*Header, error) {
	body := newRequest(1, "eth_getBlockByNumber", []interface{}{hexutil.EncodeBig(number), true})
	payloadBuf := new(bytes.Buffer)
	if err := json.NewEncoder(payloadBuf).Encode(body); err != nil {
		return nil, err
	}
	resp, err := http.Post("https://network.ambrosus.io", "application/json", payloadBuf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData := new(Response)
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}

	return &respData.Result, nil
}

func (h *Header) Rlp(withSeal bool) []byte {
	headerAsSlice := []interface{}{
		h.ParentHash,
		h.UncleHash,
		h.Coinbase,
		h.Root,
		h.TxHash,
		h.ReceiptHash,
		h.Bloom,
		h.Difficulty.ToInt(),
		h.Number.ToInt(),
		h.GasLimit,
		h.GasUsed,
		h.Time,
		h.Extra,
	}

	if withSeal {
		for _, seal := range h.SealFields {
			sealBytes, _ := hexutil.Decode(seal)
			sealRlpDecoded := new([]byte)
			err := rlp.DecodeBytes(sealBytes, sealRlpDecoded)
			if err != nil {
				panic(err)
			}

			headerAsSlice = append(headerAsSlice, *sealRlpDecoded)
			//res = append(res, sealBytes...)
		}
	}
	res, err := rlp.EncodeToBytes(headerAsSlice)
	if err != nil {
		panic(err)
	}
	return res
}

func (h *Header) Hash(seal bool) common.Hash {
	return common.BytesToHash(mytrie.Hash(h.Rlp(seal)))
}
