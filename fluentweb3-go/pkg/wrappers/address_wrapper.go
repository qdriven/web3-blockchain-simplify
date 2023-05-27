package wrappers

import (
	"context"
	_crypto "crypto"
	_ecdsa "crypto/ecdsa"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type WrappedAddress struct {
	PriKey       *_ecdsa.PrivateKey
	PrivKeyStr   string
	PrivKeyBytes []byte
	PubKey       _crypto.PublicKey
	PubKeyStr    string
	AddressStr   string
	Address      common.Address
}

const CONFIG_FILE = "./config.json"

type ContractOperators struct {
	Admin     ContractOperator
	WhiteList []ContractOperator
}

type ContractOperator struct {
	Address string
	PriKey  string
}

func NewWrapperAddress() *WrappedAddress {
	addr := &WrappedAddress{}
	addr.generatePrivKeys()
	addr.generatePubKeys(addr.PriKey)
	return addr
}

func LoadWrapperAddress(priKey string) *WrappedAddress {
	addr := &WrappedAddress{}
	privateKey, err := crypto.HexToECDSA(strings.ReplaceAll(priKey, "0x", ""))
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)
	addr.PriKey = privateKey
	addr.PrivKeyStr = privateKeyStr
	addr.PrivKeyBytes = privateKeyBytes
	addr.generatePubKeys(privateKey)
	return addr
}

func (wa *WrappedAddress) generatePrivKeys() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)
	wa.PriKey = privateKey
	wa.PrivKeyStr = privateKeyStr
	wa.PrivKeyBytes = privateKeyBytes
}

func (wa *WrappedAddress) generatePubKeys(privateKey *_ecdsa.PrivateKey) {
	publicKey := privateKey.Public()
	publicKeyECDSA, err := publicKey.(*_ecdsa.PublicKey)
	if !err {
		log.Fatal(err)
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKeyStr := hexutil.Encode(publicKeyBytes)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	addressStr := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	wa.PubKey = publicKey
	wa.PubKeyStr = pubKeyStr
	wa.Address = address
	wa.AddressStr = addressStr
}

func IsContractAddress(client ethclient.Client, addrHex string) bool {
	address := common.HexToAddress(addrHex)
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0
	return isContract
}

//TODO: contract address and use address convertion
