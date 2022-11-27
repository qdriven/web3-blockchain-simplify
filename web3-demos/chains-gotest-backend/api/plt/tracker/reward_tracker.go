package tracker

import (
	"context"
	"gin-vue-admin/chain/core"
	chainData "gin-vue-admin/chain/model"
	"gin-vue-admin/chain/pkg/sdk"
	"gin-vue-admin/global"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native/governance"
	"math/big"
)

func FilterRewardEvent(startHeight int64, endHeight int64) ([]*chainData.RewardEvent, error) {

	query := [][]interface{}{{sdk.GovernanceABI.Events[governance.EventReward].ID()}}
	topics, err := sdk.MakeTopics(query...)
	if err != nil {
		return nil, err
	}
	config := ethereum.FilterQuery{
		Addresses: []common.Address{sdk.GovernanceAddress},
		Topics:    topics,
		FromBlock: new(big.Int).SetInt64(startHeight),
		ToBlock:   new(big.Int).SetInt64(endHeight),
	}
	logs, err := core.Client.FilterLogs(context.Background(), config)
	events := make([]*chainData.RewardEvent, 0)
	for _, log := range logs {
		event := &chainData.RewardEvent{
			EventName:       governance.EventReward,
			ContractAddress: log.Address.Hex(),
			CurrentBlockNum: log.BlockNumber,
			Amount:          new(big.Int).SetBytes(log.Data).String(),
		}
		global.GVA_DB.Save(event)
		events = append(events, event)
	}
	return events, nil
}
