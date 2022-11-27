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
	"github.com/ethereum/go-ethereum/contracts/native/nft"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

var (
	NFTManagerAddress = common.HexToAddress(native.NFTContractCreateAddress)
	NFTABI            = nft.GetABI()
)

func FilterNFTLogs(startHeight, endHeight int64, EventName string) ([]types.Log, error) {
	query := [][]interface{}{{NFTABI.Events[EventName].ID()}}
	topics, err := sdk.MakeTopics(query...)
	if err != nil {
		return nil, err
	}
	//filterAddress := make([]common.Address, 0)
	//if EventName != nft.EventDeploy {
	//	addressStr := make([]string, 0)
	//	global.GVA_DB.Raw("select distinct deployed_nft_address from nft_events where event_name='deploy'").Scan(&addressStr)
	//	for _, value := range addressStr {
	//		filterAddress = append(filterAddress, common.HexToAddress(value))
	//	}
	//} else {
	//	filterAddress = append(filterAddress)
	//}
	config := ethereum.FilterQuery{
		//Addresses: filterAddress,
		Topics:    topics,
		FromBlock: new(big.Int).SetInt64(startHeight),
		ToBlock:   new(big.Int).SetInt64(endHeight),
	}
	logs, err := core.Client.FilterLogs(context.Background(), config)
	return logs, err
}

func TrackNFTApprovalLogs(startHeight int64, endHeight int64) []*model.NFTEvent {
	logs, err := FilterNFTLogs(startHeight, endHeight, nft.EventApproval)
	if err != nil {
		log.Errorf("error is {}", err)
		return nil
	}
	events := make([]*model.NFTEvent, 0)
	for _, log := range logs {
		if log.Address.String() != native.PLTContractAddress {
			event := &model.NFTEvent{
				EventName:       nft.EventApproval,
				ContractAddress: log.Address.String(),
				BlockHeight:     log.BlockNumber,
				TxHash:          log.TxHash.String(),
				Owner:           common.BytesToAddress(log.Topics[1].Bytes()).String(),
				Spender:         common.BytesToAddress(log.Topics[2].Bytes()).String(),
				TokenId:         new(big.Int).SetBytes(log.Data).String(),
			}
			global.GVA_DB.Save(event)
			events = append(events, event)
		}
	}
	return events
}

func TrackNFTApprovalAllLogs(startHeight int64, endHeight int64) []*model.NFTEvent {
	logs, err := FilterNFTLogs(startHeight, endHeight, nft.EventApprovalForAll)
	if err != nil {
		log.Errorf("error is {}", err)
		return nil
	}
	events := make([]*model.NFTEvent, 0)

	for _, log := range logs {
		if log.Address.String() != native.PLTContractAddress {

			event := &model.NFTEvent{
				EventName:       nft.EventApprovalForAll,
				ContractAddress: log.Address.String(),
				BlockHeight:     log.BlockNumber,
				TxHash:          log.TxHash.String(),
				Owner:           common.BytesToAddress(log.Topics[1].Bytes()).String(),
				Spender:         common.BytesToAddress(log.Topics[2].Bytes()).String(),
				Approved:        new(big.Int).SetBytes(log.Data).String(),
			}
			global.GVA_DB.Save(event)
			events = append(events, event)
		}
	}
	return events
}

func TrackNFTTransferLogs(startHeight int64, endHeight int64) []*model.NFTEvent {
	logs, err := FilterNFTLogs(startHeight, endHeight, nft.EventTransfer)
	if err != nil {
		log.Errorf("error is {}", err)
		return nil
	}
	events := make([]*model.NFTEvent, 0)
	for _, log := range logs {
		if log.Address.String() != native.PLTContractAddress {
			event := &model.NFTEvent{
				EventName:       nft.EventTransfer,
				ContractAddress: log.Address.String(),
				BlockHeight:     log.BlockNumber,
				TxHash:          log.TxHash.String(),
				From:            common.BytesToAddress(log.Topics[1].Bytes()).String(),
				To:              common.BytesToAddress(log.Topics[2].Bytes()).String(),
				TokenId:         new(big.Int).SetBytes(log.Data).String(),
			}
			global.GVA_DB.Save(event)
			events = append(events, event)
		}
	}
	return events
}

func TrackNFTDeployLogs(startHeight int64, endHeight int64) []*model.NFTEvent {
	logs, err := FilterNFTLogs(startHeight, endHeight, nft.EventDeploy)
	if err != nil {
		log.Errorf("error is {}", err)
		return nil
	}
	events := make([]*model.NFTEvent, 0)
	for _, log := range logs {
		if log.Address.String() != native.PLTContractAddress {

			event := &model.NFTEvent{
				EventName:          nft.EventDeploy,
				ContractAddress:    log.Address.String(),
				BlockHeight:        log.BlockNumber,
				TxHash:             log.TxHash.String(),
				Owner:              common.BytesToAddress(log.Topics[1].Bytes()).String(),
				DeployedNFTAddress: common.BytesToAddress(log.Topics[2].Bytes()).String(),
				DeployNFTData:      new(big.Int).SetBytes(log.Data).String(),
			}
			global.GVA_DB.Save(event)
			events = append(events, event)
		}
	}
	return events
}

func TrackNFTEvents(startHeight int64, endHeight int64) {
	TrackNFTDeployLogs(startHeight, endHeight)
	TrackNFTApprovalAllLogs(startHeight, endHeight)
	TrackNFTApprovalLogs(startHeight, endHeight)
	TrackNFTTransferLogs(startHeight, endHeight)
}
