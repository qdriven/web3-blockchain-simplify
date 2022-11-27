package clients

import (
	"chains-gotest-backend/api/evm/log"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/net/context"
	"strings"
	"testing"
)

func TestEvmClient(t *testing.T) {
	//address := "0xBA31E9564BbBdD3394b35E3AfDc8985554E50F4F"
	url := "http://127.0.0.1:8545"
	privateKey := strings.ReplaceAll("0x4f17afd61d017aa256a94d4edcd3ec3241e75ef6ca5309787d2ad8ca96f1d613", "0x", "")
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Error(err)
	}
	evmClient := NewEvmClient(url, key)
	nonce := evmClient.CurrentNonce
	address := evmClient.Address().String()
	currentBalance, _ := evmClient.BalanceAt(context.Background(), evmClient.Address(), nil)
	fmt.Println(nonce, address)
	fmt.Println(currentBalance)
	tx, _, err := evmClient.DeployERC20("")
	fmt.Println(tx.Hash().String())
	//hash := common.HexToHash(tx.Hash())
	txReceipt, _ :=evmClient.GetReceipt(tx.Hash())
	fmt.Println(txReceipt.BlockHash)
	fmt.Println(txReceipt.Status)
	//tx1,contract,_:=evmClient.DeployERC20Template()
	//fmt.Println(tx1.Hash().String())
	//totalSupply, _ := contract.TotalSupply(nil)
	//fmt.Println(totalSupply)
}
