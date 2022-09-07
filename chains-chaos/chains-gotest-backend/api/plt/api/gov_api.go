package api

import (
	"chains-gotest-backend/api/plt/core"
	"chains-gotest-backend/api/plt/model"
	"chains-gotest-backend/api/plt/pkg/log"
	"chains-gotest-backend/global"
	"chains-gotest-backend/model/response"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native/plt"
	"github.com/gin-gonic/gin"
	"math/big"
	"strconv"
)

type AddValidatorRequest struct {
	Validator    string
	StakeAccount string
	Revoke       bool
	OwnerAddress string
}

// @Tags GovernanceApi
// @Summary Add Validator
// @accept application/json
// @Produce application/json
// @Param data body AddValidatorRequest true "Validator,StakeAccount,Revoke"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govAddValidator [post]
func AddValidator(ctx *gin.Context) {
	request := &AddValidatorRequest{}
	_ = ctx.ShouldBind(&request)
	client := CreateNewClient(request.OwnerAddress)
	result, err := client.AddValidator(common.HexToAddress(request.Validator),
		common.HexToAddress(request.StakeAccount), request.Revoke)
	//填写逻辑
	if err != nil {
		log.Error(err)
		response.FailWithDetailed(err, "add validator failed ", ctx)
	} else {
		blockNum := core.Client.GetBlockNumber()
		validatorHistory := &model.ValidatorHistory{
			Validator:          request.Validator,
			StackAccount:       request.StakeAccount,
			Revoke:             strconv.FormatBool(request.Revoke),
			TxHash:             result.String(),
			CurrentBlockHeight: int64(blockNum),
		}
		global.GVA_DB.Save(validatorHistory)
		response.OkWithData(validatorHistory, ctx)
	}
}

// @Tags GovernanceApi
// @Summary Get Validator
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govGetEffectiveValidators [get]
func GetEffectiveValidators(ctx *gin.Context) {
	result := core.Client.GetEffectiveValidators("latest")
	//填写逻辑
	pnodes := getValidators(result, false)
	response.OkWithData(pnodes, ctx)
}

// @Tags GovernanceApi
// @Summary Get All Validator include not effective node
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govGetAllValidators [get]
func GetAllValidators(ctx *gin.Context) {

	result := core.Client.GetAllValidators("latest")
	//填写逻辑
	pnodes := getValidators(result, true)
	response.OkWithData(pnodes, ctx)
	//填写逻辑
}

func getValidators(result []common.Address, checkEffective bool) []*model.PaletteNode {
	pnodes := make([]*model.PaletteNode, 0)
	for _, node := range result {
		totalStakeAmount := core.Client.GetValidatorTotalStakeAmount(node, "latest")
		nodeType := 1
		if checkEffective {
			validatorResult := core.Client.CheckValidator(node, "latest")
			if !validatorResult {
				nodeType = 0
			}
		}
		currentBlockNum := core.Client.GetBlockNumber()
		var pnode = &model.PaletteNode{
			CurrentBlockNum: strconv.FormatUint(currentBlockNum, 10),
			NodeAddress:     node.String(),
			NodeType:        nodeType,
			StakeAmount:     totalStakeAmount.String(),
			Active:          1,
		}
		//global.GVA_DB.Save(pnode) //TODO: 不能重复, 不保存数据库，因为已经存在
		pnodes = append(pnodes, pnode)
	}
	return pnodes
}

type CheckValidatorRequest struct {
	Validator string
	BlockNum  string
}

// @Tags GovernanceApi
// @Summary Check Validator Address is effective
// @accept application/json
// @Produce application/json
// @Param data body CheckValidatorRequest true "Validator,BlockNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govCheckValidator [post]
func CheckValidator(ctx *gin.Context) {
	request := &CheckValidatorRequest{}
	_ = ctx.ShouldBind(&request)
	result := true
	if request.BlockNum == "" {
		result = core.Client.CheckValidator(common.HexToAddress(request.Validator), "latest")
	} else {
		result = core.Client.CheckValidator(common.HexToAddress(request.Validator), request.BlockNum)
	}
	response.OkWithData(result, ctx)
}

type StakeRequest struct {
	Validator    string
	StakeAccount string
	Amount       string
	IsAdd        bool
	OwnerAddress string
}

// @Tags GovernanceApi
// @Summary Stack or unstack
// @accept application/json
// @Produce application/json
// @Param data body StakeRequest true "Validator,StakeAccount,Amount,IsAdd,OwnerAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govStake [post]
func Stake(ctx *gin.Context) {
	request := &StakeRequest{}
	_ = ctx.ShouldBind(&request)
	client := CreateNewClient(request.OwnerAddress)
	stakeAmount, _ := strconv.ParseFloat(request.Amount, 18)
	result, _ := client.Stake(common.HexToAddress(request.Validator),
		common.HexToAddress(request.StakeAccount), plt.MultiFloatPLT(stakeAmount).BigInt(), !request.IsAdd)
	//填写逻辑
	response.OkWithData(result.String(), ctx)
}

type GetStakeAmountRequest struct {
	Owner       string
	Validator   string
	BlockNumber string
}

// @Tags GovernanceApi
// @Summary Get Stack amount in a given block number
// @accept application/json
// @Produce application/json
// @Param data body GetStakeAmountRequest true "Validator,Owner,BlockNumber"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govGetStakeAmount [post]
func GetStakeAmount(ctx *gin.Context) {
	request := &GetStakeAmountRequest{}
	_ = ctx.ShouldBind(&request)
	queryBlock := "latest"
	if request.BlockNumber != "" {
		queryBlock = request.BlockNumber
	}
	result := core.Client.GetStakeAmount(common.HexToAddress(request.Validator), common.HexToAddress(request.Owner), queryBlock)
	//填写逻辑
	response.OkWithData(result, ctx)
}

type GetStakeAccountRequest struct {
	Validator string
	BlockNum  string
}

// @Tags GovernanceApi
// @Summary Get Stack amount in a given block number
// @accept application/json
// @Produce application/json
// @Param data body GetStakeAccountRequest true "Validator,BlockNumber"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govGetStakeAccount [post]
func GetStakeAccount(ctx *gin.Context) {
	request := &GetStakeAccountRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.GetStakeAccount(common.HexToAddress(request.Validator), "latest")
	//填写逻辑
	response.OkWithData(result, ctx)
}

type GetValidatorTotalStakeAmountRequest struct {
	Validator   string
	BlockNumber string
}

func getQueryBlockNumber(givenBlockNumber string) string {
	queryBlock := "latest"
	if givenBlockNumber != "" {
		queryBlock = givenBlockNumber
	}
	return queryBlock
}

// @Tags GovernanceApi
// @Summary Get Stack amount in a given block number for a given validator
// @accept application/json
// @Produce application/json
// @Param data body GetValidatorTotalStakeAmountRequest true "Validator,BlockNumber"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govGetValidatorTotalStakeAmount [post]
func GetValidatorTotalStakeAmount(ctx *gin.Context) {
	request := &GetValidatorTotalStakeAmountRequest{}
	_ = ctx.ShouldBind(&request)
	result := core.Client.GetValidatorTotalStakeAmount(common.HexToAddress(request.Validator), getQueryBlockNumber(request.BlockNumber))
	//填写逻辑
	response.OkWithData(result.String(), ctx)
}

//TODO: add Get All Stake Amount
//func GetAllStakeAmount(ctx *gin.Context) {
//
//	result, _ := core.Client.GetAllStakeAmount()
//	//填写逻辑
//	response.OkWithData(balance, ctx)
//}

type ProposeRequest struct {
	ProposalType    uint8 `form:"proposalType" json:"proposalType"`
	Value           int  `form:"value" json:"value"`
	ProposalAddress string `form:"proposalAddress" json:"proposalAddress"`
}

// @Tags GovernanceApi
// @Summary Propose a Issue
// @accept application/json
// @Produce application/json
// @Param data body ProposeRequest true "ProposalType,Value,ProposalAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govPropose [post]
func Propose(ctx *gin.Context) {
	request := &ProposeRequest{}
	_ = ctx.ShouldBind(&request)

	client := CreateNewClient(request.ProposalAddress)
	proposValue := big.NewInt(0)
	if request.ProposalType ==3 {
		proposValue= big.NewInt(int64(request.Value))
	}else{
		proposValue= plt.MultiPLT(request.Value)
	}
	result, err := client.Propose(request.ProposalType, proposValue)
	//填写逻辑
	if err != nil {
		response.FailWithDetailed(err, "Propose failed", ctx)
	}
	response.OkWithData(result.String(), ctx)
}


type GetProposalByHashRequest struct {
	Hash string `form:"hash" json:"hash"`
}
// @Tags GovernanceApi
// @Summary Get Proposal By Hash
// @accept application/json
// @Produce application/json
// @Param data body GetProposalByHashRequest true "hash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govGetProposalByHash [post]
func GetProposalByHash(ctx *gin.Context) {
	request := &GetProposalByHashRequest{}
	_ = ctx.ShouldBindJSON(&request)
	proposalId,output, _ := core.Client.GetProposalFromReceipt(common.HexToHash(request.Hash))
	//填写逻辑
	result :=make(map[string]interface{})
	result["id"]=proposalId.String()
	result["proposal"]=output
	response.OkWithData(result, ctx)
}

type GetProposalRequest struct {
	ProposalId string
	BlockNum   string
}

// @Tags GovernanceApi
// @Summary Get Proposal
// @accept application/json
// @Produce application/json
// @Param data body GetProposalRequest true "ProposalId,BlockNumber"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govGetProposal [post]
func GetProposal(ctx *gin.Context) {
	request := &GetProposalRequest{}
	_ = ctx.ShouldBind(&request)
	proposalAddress := common.HexToAddress(request.ProposalId)
	result, _ := core.Client.GetProposal(proposalAddress, getQueryBlockNumber(request.BlockNum))
	//填写逻辑
	response.OkWithData(result, ctx)
}

type VoteRequest struct {
	ProposalId   string
	VoterAddress string
}

// @Tags GovernanceApi
// @Summary Get Vote information
// @accept application/json
// @Produce application/json
// @Param data body VoteRequest true "ProposalId,VoterAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govVote [post]
func Vote(ctx *gin.Context) {
	request := &VoteRequest{}
	_ = ctx.ShouldBind(&request)
	client:=CreateNewClient(request.VoterAddress)
	proposalIdAddress := common.HexToAddress(request.ProposalId)
	log.Info("current proposal id {}",proposalIdAddress.String())
	result, _ := client.Vote(proposalIdAddress)
	//填写逻辑
	response.OkWithData(result.String(), ctx)
}

//TODO: getVoteTotalAmount
//func GetVoteTotalAmount(ctx *gin.Context) {
//	request := &VoteRequest{}
//	_ = ctx.ShouldBind(&request)
//	result, _ := core.Client.GetVoteTotalAmount()
//	//填写逻辑
//	response.OkWithData(balance, ctx)
//}

type GetGlobalParamsRequest struct {
	ProposalType uint8
	BlockNum     string
}

// @Tags GovernanceApi
// @Summary Get global params information
// @accept application/json
// @Produce application/json
// @Param data body GetGlobalParamsRequest true "ProposalId,BlockNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govGetGlobalParams [post]
func GetGlobalParams(ctx *gin.Context) {
	request := &GetGlobalParamsRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.GetGlobalParams(request.ProposalType, getQueryBlockNumber(request.BlockNum))
	//填写逻辑
	response.OkWithData(result, ctx)
}

//TODO: set delegate reward factor
type SetDelegateRewardFactorRequest struct {
	Factor int64
}

//
//func SetDelegateRewardFactor(ctx *gin.Context) {
//	request := &SetDelegateRewardFactorRequest{}
//	_ = ctx.ShouldBind(&request)
//	result, _ := core.Client.SetD()
//	//填写逻辑
//	response.OkWithData(balance, ctx)
//}

type DelegateRewardFactorRequest struct {
	Validator string
}

//func Getdelegaterewardfactor(ctx *gin.Context) {
//	request := &DelegateRewardFactorRequest{}
//	_ = ctx.ShouldBind(&request)
//	result, _ := core.Client.()
//	//填写逻辑
//	response.OkWithData(balance, ctx)
//}

// @Tags GovernanceApi
// @Summary Get global params information
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/govGetLastRewardBlockHeight [get]
func Getlastrewardblockheight(ctx *gin.Context) {

	result, _ := core.Client.GetLastRewardBlock("latest")
	//填写逻辑
	response.OkWithData(result, ctx)
}
