package sdk

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type EthClient struct {
	*ethclient.Client
}
type Client struct {
	*rpc.Client
	*EthClient
	backend      *ethclient.Client
	url          string
	Key          *ecdsa.PrivateKey
	currentNonce uint64
}

func NewSender(url string, key *ecdsa.PrivateKey) *Client {
	cli := ethClientDialNode(url)
	return &Client{
		url:       url,
		Client:    dialNode(url),
		EthClient: &EthClient{Client: cli},
		Key:       key,
		backend:   cli,
	}
}

func (c *Client) Url() string {
	return c.url
}

func (c *Client) Address() common.Address {
	return crypto.PubkeyToAddress(c.Key.PublicKey)
}

func (c *Client) Reset(key *ecdsa.PrivateKey) *Client {
	c.Key = key
	return c
}

func ethClientDialNode(url string) *ethclient.Client {
	ethClient, err := ethclient.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("failed to dial geth rpc [%v]", err))
	}

	return ethClient
}

func dialNode(url string) *rpc.Client {
	client, err := rpc.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("failed to dial geth rpc [%v]", err))
	}

	return client
}

func (c *Client) TransactionByHash(hash common.Hash) *types.Transaction {
	tx, _, err := c.EthClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		fmt.Printf("TransactionByHash err: %s\n", err.Error())
	}
	return tx
}

func (c *Client) GetEventByHash(hash common.Hash) ([]*types.Log, error) {
	raw := &types.Receipt{}
	if err := c.Call(raw, "eth_getTransactionReceipt", hash.Hex()); err != nil {
		return nil, fmt.Errorf("failed to get transaction receipt: [%v]", err)
	}
	return raw.Logs, nil
}
