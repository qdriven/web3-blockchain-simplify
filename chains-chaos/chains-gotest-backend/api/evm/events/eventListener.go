package events

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

type EthEventListenerClient interface {
	// SubscribeNewHead subscribes to notifications about the current blockchain head
	// on the given channel.
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	// FilterLogs executes a filter query.
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	// SubscribeFilterLogs subscribes to the results of a streaming filter query.
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	// TransactionReceipt returns the receipt of a transaction by transaction hash.
	// Note that the receipt is not available for pending transactions.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	// BlockByHash returns the given full block.
	//
	// Note that loading full blocks requires two requests. Use HeaderByHash
	// if you don't need all transactions or uncle headers.
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
}

const (
	bufferedLogSize = 1000
)

var logger = log.New()

type EventListener struct {
	client EthEventListenerClient
	logCh  chan types.Log

	// Contract address <-> Contract mapping
	addressMap map[common.Address]*Contract
}

func NewEventListener(client EthEventListenerClient,
	contracts []*Contract) *EventListener {

	l := &EventListener{
		client:     client,
		addressMap: make(map[common.Address]*Contract),
		logCh:      make(chan types.Log, bufferedLogSize),
	}

	for _, c := range contracts {
		l.addressMap[c.Address] = c
	}

	return l
}

func (el *EventListener) Listen(fromBlock *big.Int, eventCh chan<- *ContractEvent, stop <-chan struct{}) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	contracts := make([]common.Address, 0)
	for addr, _ := range el.addressMap {
		contracts = append(contracts, addr)
	}
	q := ethereum.FilterQuery{
		FromBlock: fromBlock,
		Addresses: contracts,
	}
	sub, err := el.client.SubscribeFilterLogs(ctx, q, el.logCh)
	if err != nil {
		logger.Error("Failed to subscribe filter logs", "err", err)
		return err
	}
	defer sub.Unsubscribe()

	// fetch the past logs
	logs, err := el.client.FilterLogs(context.Background(), q)
	if err != nil {
		logger.Error("Failed to execute a filter query command", "err", err)
		return err
	}

	go func() {
		for _, l := range logs {
			el.logCh <- l
		}
	}()

	for {
		select {
		case err := <-sub.Err():
			logger.Error("Unexpected subscription error", "err", err)
			return err
		case log := <-el.logCh:
			if cEvent := el.Parse(log); cEvent != nil {
				eventCh <- cEvent
			}
		case <-stop:
			logger.Warn("Received a stop signal")
			return nil
		}
	}
}

func (el *EventListener) Parse(l types.Log) *ContractEvent {
	c, ok := el.addressMap[l.Address]
	if !ok {
		return nil
	}
	if len(l.Topics) == 0 {
		return nil
	}
	name, ok := c.events[l.Topics[0]]
	if !ok {
		return nil
	}
	return &ContractEvent{
		BlockNumber: l.BlockNumber,
		BlockHash:   l.BlockHash,
		TxHash:      l.TxHash,
		Contract:    c,
		Name:        name,
		Data:        l.Data,
		Removed:     l.Removed,
	}
}
