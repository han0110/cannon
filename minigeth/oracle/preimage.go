//go:build !mips
// +build !mips

package oracle

import (
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
)

var preimages = make(map[common.Hash][]byte)

func Preimage(hash common.Hash) []byte {
	val, ok := preimages[hash]
	key := fmt.Sprintf("/tmp/eth/%s", hash)
	ioutil.WriteFile(key, val, 0644)
	if !ok {
		fmt.Println("can't find preimage", hash)
		panic("preimage missing")
	}
	return val
}
