package tracker

import (
	"context"
	"gin-vue-admin/chain/core"
	"gin-vue-admin/chain/model"
	"gin-vue-admin/chain/pkg/sdk"
	"gin-vue-admin/global"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native/governance"
	"github.com/ethereum/go-ethereum/contracts/native/utils"
	"math/big"
)

//
//import (
//	"context"
//	"gin-vue-admin/chain/core"
//	"gin-vue-admin/chain/model"
//	"gin-vue-admin/chain/pkg/sdk"
//	"gin-vue-admin/global"
//	"github.com/ethereum/go-ethereum"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/contracts/native/plt"
//	"math/big"
//)
//
////get gov event
////add validator
////vote processing
////nft transaction
////plt transaction
//

func FilterStakeEvents(startHeight int64, endHeight int64) ([]*model.StakingEvent, error) {
	query := [][]interface{}{{sdk.GovernanceABI.Events[governance.EventStake].ID()}}
	topics, err := sdk.MakeTopics(query...)
	if err != nil {
		return nil, err
	}
	config := ethereum.FilterQuery{
		Addresses: []common.Address{GovAddress},
		Topics:    topics,
		FromBlock: new(big.Int).SetInt64(startHeight),
		ToBlock:   new(big.Int).SetInt64(endHeight),
	}
	logs, err := core.Client.FilterLogs(context.Background(), config)
	events := make([]*model.StakingEvent, 0)
	for _, eventLog := range logs {
		event := &model.StakingEvent{
			EventName:          governance.EventStake,
			Contract:           eventLog.Address.String(),
			CurrentBlockHeight: eventLog.BlockNumber,
			TxHash:             eventLog.TxHash.String(),
			Owner:              utils.Hash2Address(eventLog.Topics[1]).String(),
			Validator:          utils.Hash2Address(eventLog.Topics[2]).String(),
			StakeAccount:       utils.Hash2Address(eventLog.Topics[3]).String(),
			Revoke:             utils.Hash2Bool(eventLog.Topics[4]),
			Data:               new(big.Int).SetBytes(eventLog.Data).String(),
		}
		global.GVA_DB.Save(event)
		events = append(events, event)
	}
	return events, nil
}
