package chains

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

var (
	QGasPrice *big.Int = big.NewInt(0)
	QGasLimit uint64   = 0
)

type ChainTransact struct {
	Address  *common.Address
	Contract *bind.BoundContract
	TxHash   *common.Hash
	Tx       *types.Transaction
}

type ChainClient struct {
	conn            bind.ContractBackend //*ethclient.Client
	c               *ethclient.Client
	auth            *bind.TransactOpts
	LastTransaction *ChainTransact
	RpcUrl          string
	rpcClient       *rpc.Client
	Key             *ecdsa.PrivateKey
}

func NewChainClient(rpcUrl string) *ChainClient {
	conn := DialNode(rpcUrl)
	rpcClient := DialRpcNode(rpcUrl)
	return &ChainClient{
		conn: conn,
		c:    conn, LastTransaction: &ChainTransact{},
		rpcClient: rpcClient,
	}
}

func DialNode(url string) *ethclient.Client {
	ethClient, err := ethclient.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("failed to dial geth rpc [%v]", err))
	}

	return ethClient
}

func DialRpcNode(url string) *rpc.Client {
	client, err := rpc.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("failed to dial geth rpc [%v]", err))
	}

	return client
}
