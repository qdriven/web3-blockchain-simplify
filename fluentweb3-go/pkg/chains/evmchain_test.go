package chains

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"
)

var RpcUrl = "http://localhost:8545"

func TestNewChainClient(t *testing.T) {
	q := NewChainClient(RpcUrl)
	fmt.Println(q)
}

func TestGetTxDetailForContract(t *testing.T) {
	q := NewChainClient(RpcUrl)
	result := q.GetBlockNumber()
	fmt.Println(result)
	txHash := common.HexToHash("0xedeccdd04368cb567e695b8a7eef9d1294048e322beaf03bab3d6c9a0822ee8b")
	tx := q.TransactionByHash(txHash)
	fmt.Println(tx.Data())
	fmt.Println(hexutil.Encode(tx.Data()))
}

func TestGetEventLogs(t *testing.T) {
	q := NewChainClient(RpcUrl)
	txHash := common.HexToHash("0xedeccdd04368cb567e695b8a7eef9d1294048e322beaf03bab3d6c9a0822ee8b")
	logs, _ := q.GetEventByHash(txHash)
	fmt.Println(logs)
	for _, vlog := range logs {
		fmt.Printf("blocknumber is %d", vlog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vlog.Index)
		fmt.Println(vlog.Topics[0].Hex())
	}
}
