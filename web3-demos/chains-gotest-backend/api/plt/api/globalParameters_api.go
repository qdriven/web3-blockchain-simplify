package api

import (
	"chains-gotest-backend/api/plt/core"
	chainData "chains-gotest-backend/api/plt/model"
	"chains-gotest-backend/global"
	"chains-gotest-backend/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ProposalType uint8

const (
	ProposalTypeUnknown      ProposalType = iota //0
	ProposalTypeMintPrice                        //1
	ProposalTypeGasFee                           //2
	ProposalTypeRewardPeriod                     //3
)

// @Tags PltParameters
// @Summary Get All global Parameters
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/getAllGlobalParameters [get]
func GetAllGlobalParameters(c *gin.Context) {
	params := make(map[string]string)
	getGlobalParameters(params, ProposalTypeMintPrice, "nft_mint_price_in_plt")
	getGlobalParameters(params, ProposalTypeGasFee, "nft_gas_price_in_plt")
	getGlobalParameters(params, ProposalTypeRewardPeriod, "reward_period")
	currentBlockNum := core.Client.GetBlockNumber()
	params["current_block_no"] = strconv.FormatUint(currentBlockNum, 10)
	globalParameter := &chainData.GlobalParameter{
		CurrentBlockNo:    strconv.FormatUint(currentBlockNum, 10),
		NftGasPriceInPlt:  params["nft_gas_price_in_plt"],
		NftMintPriceInPlt: params["nft_mint_price_in_plt"],
		RewardPeriod:      params["reward_period"],
	}
	global.GVA_DB.Create(globalParameter)
	response.OkWithData(params, c)
}

func getGlobalParameters(params map[string]string, proposalType ProposalType, parameterName string) {
	result, _ := core.Client.GetGlobalParams(uint8(proposalType), "latest")
	params[parameterName] = result.String()
}
