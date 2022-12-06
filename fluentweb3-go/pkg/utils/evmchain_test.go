package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"
)

func TestNewChainClient(t *testing.T) {
	q := NewChainClient("http://10.100.90.82:8545")
	fmt.Println(q)
}

func TestGetTxDetailForContract(t *testing.T) {
	q := NewChainClient("http://10.100.90.82:8545")
	result := q.GetBlockNumber()
	fmt.Println(result)
	//tx: 0x74e9ac43b811031a971cdbe94d94d49db72bf0d3dcd730b797a91d08b76044ff
	//contract address: 0xb2eE5989a4FF255f6bD1f015468a2069b92b2EF8
	//txHash := common.HexToHash("0x2a15db8d065531b2b82d9d41ce7dab5db08c329c326ab263ae02239734a11cec")
	txHash := common.HexToHash("0xedeccdd04368cb567e695b8a7eef9d1294048e322beaf03bab3d6c9a0822ee8b")
	tx := q.TransactionByHash(txHash)
	fmt.Println(tx.Data())
	fmt.Println(hexutil.Encode(tx.Data()))
	//addWhiteList: 0x47ee03940000000000000000000000001288aa3abe94938acb8682947a00c54cfafbd456
	//
	//result, err := hexutil.Decode("0x" + string(tx.Data()))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(result))
}

func TestGetEventLogs(t *testing.T) {
	q := NewChainClient("http://10.100.90.82:8545")
	//tx: 0x74e9ac43b811031a971cdbe94d94d49db72bf0d3dcd730b797a91d08b76044ff
	//contract address: 0xb2eE5989a4FF255f6bD1f015468a2069b92b2EF8
	txHash := common.HexToHash("0xedeccdd04368cb567e695b8a7eef9d1294048e322beaf03bab3d6c9a0822ee8b")
	logs, _ := q.GetEventByHash(txHash)
	fmt.Println(logs)
	//result, err := hexutil.Decode("0x" + string(tx.Data()))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(result))
	for _, vlog := range logs {
		fmt.Printf("blocknumber is %d", vlog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vlog.Index)
		fmt.Println(vlog.Topics[0].Hex())
	}
}
