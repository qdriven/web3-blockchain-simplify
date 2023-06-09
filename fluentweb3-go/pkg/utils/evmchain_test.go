package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"
)

func TestNewChainClient(t *testing.T) {
	q := NewChainClient("http://127.0.0.1:8545")
	fmt.Println(q)
}

func TestGetTxDetailForContract(t *testing.T) {
	q := NewChainClient("http://127.0.0.1:8545")
	result := q.GetBlockNumber()
	fmt.Println(result)
	txHash := common.HexToHash("0xedeccdd04368cb567e695b8a7eef9d1294048e322beaf03bab3d6c9a0822ee8b")
	tx := q.TransactionByHash(txHash)
	fmt.Println(tx.Data())
	fmt.Println(hexutil.Encode(tx.Data()))
}

func TestGetEventLogs(t *testing.T) {
	q := NewChainClient("http://10.100.90.82:8545")
	txHash := common.HexToHash("0xedeccdd04368cb567e695b8a7eef9d1294048e322beaf03bab3d6c9a0822ee8b")
	logs, _ := q.GetEventByHash(txHash)
	fmt.Println(logs)
	for _, vlog := range logs {
		fmt.Printf("blocknumber is %d", vlog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vlog.Index)
		fmt.Println(vlog.Topics[0].Hex())
	}
}
