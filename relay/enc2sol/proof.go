package enc2sol

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"test/mytrie"
)

type proofGen struct {
	whatSearch []byte
	Result     [][]byte
}

func CheckProof(whatSearch []byte, proof [][]byte, realRoot common.Hash) bool {

	el := whatSearch
	for i := 0; i < len(proof)-1; i += 2 {
		el = append(proof[i][:], el[:]...)
		el = append(el, proof[i+1][:]...)
		//fmt.Printf("%x\n", el)
		if len(el) > 32 {
			el = mytrie.Hash(el)
		}
		//fmt.Printf("%x\n", el)
	}

	return common.BytesToHash(el) == realRoot
}

func CalcProof(root *mytrie.ModifiedStackTrie, whatSearch []byte) [][]byte {
	p := proofGen{whatSearch: whatSearch, Result: [][]byte{}}
	p.calcProof(root)
	return p.Result
}

// todo optimize, using path
func (p *proofGen) calcProof(st *mytrie.ModifiedStackTrie) bool {
	if bytes.Contains(st.UnhashedVal, p.whatSearch) {
		//fmt.Printf("%x\n", p.whatSearch)
		//fmt.Printf("%x\n", st.UnhashedVal)

		r := bytes.SplitN(st.UnhashedVal, p.whatSearch, 2)
		if len(r) != 2 {
			panic("split not 2")
		}
		p.Result = append(p.Result, r[0], r[1])
		return true
	}

	for _, c := range st.Children {
		if c != nil && p.calcProof(c) {
			//fmt.Printf("%x\n", c.Val)
			//fmt.Printf("%x\n", st.UnhashedVal)

			r := bytes.Split(st.UnhashedVal, c.Val)
			if len(r) != 2 {
				panic("split not 2")
			}
			p.Result = append(p.Result, r[0], r[1])

			p.whatSearch = st.Val
			return true
		}

	}
	return false
}
