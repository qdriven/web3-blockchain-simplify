package service

import (
	"fmt"
	"go-chains/chains/models"
	"testing"
)

//func init() {
//	initializer.Setup4Test()
//}

func TestQueryData(t *testing.T) {
	result := QueryAllChainMetaData()
	fmt.Println(result)
}

func TestCreateMetaData(t *testing.T) {
	metaData := &models.ChainMetaData{
		ChainName: "PalletChain",
		ChainId:   "106",
		ChainType: "testnet",
		RPCUrl:    "http://106.75.250.160:22001",
	}
	result := UpdateChainMetaData(metaData)
	fmt.Println(result)
}
