package hdwallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDrive(t *testing.T) {
	data := make(map[common.Address]struct{})

	N := 10000
	for i := 0; i < N; i++ {
		acc, err := Drive(i)
		assert.NoError(t, err)
		data[acc.Address] = struct{}{}
		key, _ :=wallet.PrivateKeyHex(acc)
		fmt.Println(acc.URL.Path)
		fmt.Println(key)
	}

	assert.Equal(t, N, len(data))
}
