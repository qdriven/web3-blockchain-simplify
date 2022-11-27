package tracker

import (
	"chains-gotest-backend/api/evm/log"
	"chains-gotest-backend/api/plt/core"
	"chains-gotest-backend/global"
)

//TODO: Pass functions and table name to do the same thing
//TODO: Refactor Event Tracking
func TrackEvents() {
	var maxBlockNum int
	global.GVA_DB.Raw("select max(current_block_num) from reward_events").Scan(&maxBlockNum)
	log.Info("current scanned block number is ", maxBlockNum)
	startIndex := uint64(maxBlockNum)
	var indexInc uint64 = 2 //TODO: Move indexInc to configuration

	currentBlockNum := core.Client.GetBlockNumber()
	log.Info("current block is ", currentBlockNum)
	//FilterStakeEvents(0, 300000)

	for currentBlockNum > startIndex {
		//ADD All Tracker needed
		_, _ = FilterRewardEvent(int64(startIndex), int64(startIndex+indexInc))
		//TODO: Filter Different Logs
		_, _ = FilterPLTTransferLogs(int64(startIndex), int64(startIndex+indexInc))
		_, _ = FilterPLTApprovalLogs(int64(startIndex), int64(startIndex+indexInc))
		_, _ = FilterStakeEvents(int64(startIndex), int64(startIndex+indexInc))

		//TRACKING NFT Events
		TrackNFTEvents(int64(startIndex), int64(startIndex+indexInc))
		startIndex = startIndex + indexInc+1
		currentBlockNum = core.Client.GetBlockNumber()
		for currentBlockNum <= startIndex {
			core.WaitBlocks(2)
			currentBlockNum = core.Client.GetBlockNumber()
			log.Info("current block height is ", currentBlockNum)
		}
	}
}
