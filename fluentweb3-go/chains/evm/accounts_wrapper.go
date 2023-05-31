package evm

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountBalance(client ethclient.Client,
	accountAddressHex string, blockNum *big.Int) *big.Int {
	accountAddress := common.HexToAddress(accountAddressHex)
	balance, err := client.BalanceAt(context.Background(),
		accountAddress, blockNum)
	if err != nil {
		log.Fatal(err)
		//TODO: raise error
		return nil
	}

	return balance
}

func GetAccountBalanceInEth(client ethclient.Client, accountAddressHex string,
	blockNum *big.Int) *big.Float {
	balance := GetAccountBalance(client, accountAddressHex, blockNum)
	return toEthValue(balance)
}

func GetAccountPendingBalance(client ethclient.Client,
	accountAddressHex string, blockNum *big.Int) *big.Int {
	accountAddress := common.HexToAddress(accountAddressHex)
	balance, err := client.PendingBalanceAt(context.Background(),
		accountAddress)
	if err != nil {
		log.Fatal(err)
		//TODO: raise error
		return nil
	}
	return balance
}

func GetAccountPendingBalanceInEth(client ethclient.Client,
	accountAddressHex string, blockNum *big.Int) *big.Float {
	balance := GetAccountPendingBalance(client, accountAddressHex, blockNum)
	ethValue := toEthValue(balance)
	return ethValue
}

func toEthValue(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}

// TODO: call any contract have balanceOf function
func GetBalanceAtContract(client ethclient.Client, contractAdddress string,
	userAccount string, blockNum *big.Int) *big.Int {
	contractAddress := common.HexToAddress(contractAdddress)
	balance, err := client.BalanceAt(context.Background(),
		contractAddress, blockNum)
	if err != nil {
		log.Fatal(err)
		//TODO: raise error
		return nil
	}
	return balance

}

type AddressType string

const (
	AddressTypeInit AddressType = "" // Init value

	// After detection
	AddressTypeErc20         AddressType = "Erc20"
	AddressTypeErc721        AddressType = "Erc721"
	AddressTypeOtherContract AddressType = "OtherContract"
	AddressTypeEOA           AddressType = "EOA" // couldn't detect a smart contract, classify as Externally Owned Address (EOA)
)

type AddressDetail struct {
	Address  string      `json:"address"`
	Type     AddressType `json:"type"`
	Name     string      `json:"name"`
	Symbol   string      `json:"symbol"`
	Decimals uint8       `json:"decimals"`
}

// Returns a new unknown address detail
func NewAddressDetail(address string) AddressDetail {
	return AddressDetail{Address: address, Type: AddressTypeInit}
}

func (a AddressDetail) String() string {
	return fmt.Sprintf("%s [%s] name=%s, symbol=%s, decimals=%d", a.Address, a.Type, a.Name, a.Symbol, a.Decimals)
}

// func (a *AddressDetail) IsLoaded() bool {
func (a *AddressDetail) IsInitial() bool {
	return a.Type == AddressTypeInit
}

func (a *AddressDetail) IsErc20() bool {
	return a.Type == AddressTypeErc20
}

func (a *AddressDetail) IsErc721() bool {
	return a.Type == AddressTypeErc721
}
