package scd

import (
	"fluentweb3-go/pkg/address"
	"fmt"
	"testing"
)

// Deployment Hash: 0x381327b06733a5529081677a9a108180ed9e2c2a238c5bef54ddfdef29ebf712
// Contract Address: 0x8f5090F08f739E8d65b725cABAF107Ac1CA7d863
// ChainId = 1337
func TestCdpContractOperations(t *testing.T) {
	cdpV1 := "0x8f5090F08f739E8d65b725cABAF107Ac1CA7d863"
	cdpContract := NewCdpContractV1(cdpV1)
	//通过私钥生成地址
	ownerWrapperAddr := address.LoadWrapperAddress("0x3930cdc34dee4959019b5d1e50e06dca524b9556fa426afa81a17172fd18a76c")
	//添加白名单给某个地址
	tx, _ := cdpContract.AddToWhiteList(ownerWrapperAddr)
	fmt.Println(tx.Hash().Hex())
	// 添加排放源数据
	emissionData := "{\n    \"elementName\": \"光气的生产排放因子\",\n    \"elementInfo\": \"1.902938 kgCO2e/kg\",\n    \"publisher\": \"EcoInvent V3.8\"\n}"
	emissionTx, err := cdpContract.AddEmissionData(ownerWrapperAddr, emissionData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("交易Hash TransactionHash:", emissionTx.Hash().Hex())
	//Transaction Data
	txInfo := cdpContract.GetTransactionData(emissionTx.Hash().Hex())
	fmt.Println(txInfo)
}
