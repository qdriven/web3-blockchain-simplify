package tracker

import (
	"context"
	"gin-vue-admin/chain/core"
	"gin-vue-admin/chain/model"
	"gin-vue-admin/chain/pkg/log"
	"gin-vue-admin/chain/pkg/sdk"
	"gin-vue-admin/global"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native"
	"github.com/ethereum/go-ethereum/contracts/native/plt"
	"math/big"
)

var (
	PLTAddress          = common.HexToAddress(native.PLTContractAddress)
	GovAddress	=common.HexToAddress(native.GovernanceContractAddress)
	PLTABI              = plt.GetABI()
)


func TrackPltTransfer() {
	var maxBlockNum int
	global.GVA_DB.Raw("select max(current_block_num) from plt_transfer_events]").Scan(&maxBlockNum)
	log.Info("current scanned block number is ", maxBlockNum)
	startIndex := uint64(maxBlockNum)
	var indexInc uint64 = 100

	currentBlockNum := core.Client.GetBlockNumber()
	log.Info("current block is ", currentBlockNum)
	for currentBlockNum > startIndex {
		FilterPLTTransferLogs(int64(startIndex), int64(startIndex+indexInc))
		FilterPLTApprovalLogs(int64(startIndex), int64(startIndex+indexInc))
		startIndex = startIndex + indexInc
		currentBlockNum = core.Client.GetBlockNumber()
		for currentBlockNum <= startIndex {
			core.WaitBlocks(int(indexInc))
			currentBlockNum = core.Client.GetBlockNumber()
			log.Info("current block height is ", currentBlockNum)
		}
	}
}

/**
Filter and save PLT Transfer Events
*/
func FilterPLTTransferLogs(startHeight int64, endHeight int64) ([]*model.PLTTransferEvent, error) {
	query := [][]interface{}{{PLTABI.Events[plt.EventTransfer].ID()}}
	topics, err := sdk.MakeTopics(query...)
	if err != nil {
		return nil, err
	}
	config := ethereum.FilterQuery{
		Addresses: []common.Address{PLTAddress},
		Topics:    topics,
		FromBlock: new(big.Int).SetInt64(startHeight),
		ToBlock:   new(big.Int).SetInt64(endHeight),
	}
	logs, err := core.Client.FilterLogs(context.Background(), config)
	events := make([]*model.PLTTransferEvent, 0)
	for _, eventLog := range logs {
		event := &model.PLTTransferEvent{
			EventName:   plt.EventTransfer,
			Contract:    eventLog.Address.String(),
			BlockHeight: eventLog.BlockNumber,
			TxHash:      eventLog.TxHash.String(),
			From:        common.BytesToAddress(eventLog.Topics[1].Bytes()).String(),
			To:          common.BytesToAddress(eventLog.Topics[2].Bytes()).String(),
			Value:       new(big.Int).SetBytes(eventLog.Data).String(),
		}
		global.GVA_DB.Save(event)
		events = append(events, event)
	}
	return events, nil
}

func FilterPLTApprovalLogs(startHeight int64, endHeight int64) ([]*model.PLTApprovalEvent, error) {
	query := [][]interface{}{{PLTABI.Events[plt.EventApproval].ID()}}
	topics, err := sdk.MakeTopics(query...) //ADD MORE TOPICS or subscribe more topics
	if err != nil {
		return nil, err
	}
	config := ethereum.FilterQuery{
		Addresses: []common.Address{PLTAddress},
		Topics:    topics,
		FromBlock: new(big.Int).SetInt64(startHeight),
		ToBlock:   new(big.Int).SetInt64(endHeight),
	}
	logs, err := core.Client.FilterLogs(context.Background(), config)
	events := make([]*model.PLTApprovalEvent, 0)
	for _, log := range logs {
		event := &model.PLTApprovalEvent{
			EventName:   plt.EventApproval,
			Contract:    log.Address.String(),
			BlockHeight: log.BlockNumber,
			TxHash:      log.TxHash.String(),
			Owner:       common.BytesToAddress(log.Topics[1].Bytes()).String(),
			Spender:     common.BytesToAddress(log.Topics[2].Bytes()).String(),
			Value:       new(big.Int).SetBytes(log.Data).String(),
		}
		global.GVA_DB.Save(event)
		events = append(events, event)
	}
	return events, nil
}
