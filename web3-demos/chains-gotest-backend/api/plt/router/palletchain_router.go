package router

import (
	"chains-gotest-backend/api/plt/api"
	"chains-gotest-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitPalletChainApis(Router *gin.RouterGroup) {

	ApiRouter := Router.Group("pallet").Use(middleware.OperationRecord())
	{
		ApiRouter.GET("getAllGlobalParameters", api.GetAllGlobalParameters) // Get All Global Parameters
		//node:
		ApiRouter.POST("createTestAccounts", api.CreateTestAccounts)
		ApiRouter.GET("getInitNodeInfo", api.GetInitNodeInfo)

		//plt
		ApiRouter.POST("pltTransfer", api.PltTransfer)
		ApiRouter.POST("pltApprove", api.PltApprove)
		ApiRouter.POST("pltAllowance", api.PltAllowance)
		ApiRouter.POST("pltTransferFrom", api.PltSafeTransferFrom)
		ApiRouter.POST("pltBatchTransfer", api.BatchTransfer)
		ApiRouter.POST("pltBalanceOf", api.PltBalanceOf)
		ApiRouter.POST("pltPerformance", api.PltPerformance)
		ApiRouter.GET("pltMeta", api.GetPltMeta)
		//nft
		ApiRouter.POST("nftDeploy", api.Deploy)
		ApiRouter.POST("nftBatchDeploy", api.BatchDeploy)
		//ApiRouter.POST("nftSafetransferfrom", api.Safetransferfrom)
		ApiRouter.POST("nftTransferfrom", api.Transferfrom)
		ApiRouter.POST("nftApprove", api.Approve)
		ApiRouter.POST("nftSetapprovalforall", api.Setapprovalforall)
		ApiRouter.POST("nftBalanceof", api.Balanceof)
		ApiRouter.POST("nftOwnerof", api.Ownerof)
		ApiRouter.POST("nftGetApproved", api.Getapproved)
		ApiRouter.POST("nftIsapprovedforall", api.Isapprovedforall)
		ApiRouter.POST("nftTotalSupply", api.TotalSupply)
		ApiRouter.POST("nftTokenbyindex", api.Tokenbyindex)
		ApiRouter.POST("nftTokenofownerbyindex", api.Tokenofownerbyindex)
		ApiRouter.POST("nftMint", api.Mint)
		ApiRouter.POST("nftBurn", api.Burn)
		ApiRouter.POST("nftName", api.Name)
		ApiRouter.POST("nftSymbol", api.Symbol)
		ApiRouter.POST("nftOwner", api.Owner)
		ApiRouter.POST("nftTokenuri", api.Tokenuri)
		ApiRouter.POST("nftSetBaseUri", api.SetTokenBaseUri)
		ApiRouter.POST("nftGetTokenBaseUri", api.GetTokenBaseUri)

		//gov api
		ApiRouter.POST("govAddValidator", api.AddValidator)
		ApiRouter.GET("govGetEffectiveValidators", api.GetEffectiveValidators)
		ApiRouter.GET("govGetAllValidators", api.GetAllValidators)
		ApiRouter.POST("govCheckValidator", api.CheckValidator)
		ApiRouter.POST("govStake", api.Stake)
		ApiRouter.POST("govGetStakeAmount", api.GetStakeAmount)
		ApiRouter.POST("govGetValidatorTotalStakeAmount", api.GetValidatorTotalStakeAmount)
		//ApiRouter.POST("govGetAllStakeAmount",api.Get)
		ApiRouter.POST("govPropose", api.Propose)
		ApiRouter.POST("govGetProposal", api.GetProposal)
		ApiRouter.POST("govVote", api.Vote)
		//ApiRouter.POST("govGetvotetotalamount",api.Get)
		ApiRouter.POST("govGetGlobalParams", api.GetGlobalParams)
		//ApiRouter.POST("govSetDelegateRewardFactor",api.Setapprovalforall)
		//ApiRouter.POST("govGetdelegaterewardfactor",api.Get)
		ApiRouter.GET("govGetLastRewardBlockHeight", api.Getlastrewardblockheight)
		ApiRouter.POST("govGetStakeAccount", api.GetStakeAccount)
		ApiRouter.POST("govGetProposalByHash", api.GetProposalByHash)

		//node api
		ApiRouter.GET("getAllInitializedNodes", api.GetAllInitializedNodes)
		ApiRouter.POST("startNode", api.StartNode)
		ApiRouter.POST("stopNode", api.StopNode)
		ApiRouter.GET("ResetPltGenesisNetwork", api.ResetPltGenesisNetwork)
		ApiRouter.GET("ReStartPltGenesisNetwork", api.ReStartPltGenesisNetwork)
		ApiRouter.GET("InitPltGenesisNetwork", api.InitPltGenesisNetwork)
		ApiRouter.GET("StartPltGenesisNetwork", api.StartPltGenesisNetwork)
		ApiRouter.GET("StopPltGenesisNetwork", api.StopPltGenesisNetwork)
		ApiRouter.GET("ClearPltGenesisNetwork", api.ClearPltGenesisNetwork)
		ApiRouter.GET("ResetPltValidatorNetwork", api.ResetPltValidatorNetwork)
		ApiRouter.GET("ReStartPltValidatorNetwork", api.ReStartPltValidatorNetwork)
		ApiRouter.GET("StartPltValidatorNetwork", api.StartPltValidatorNetwork)
		ApiRouter.GET("StopPltValidatorNetwork", api.StopPltValidatorNetwork)
		ApiRouter.GET("ClearPltValidatorNetwork", api.ClearPltValidatorNetwork)
		ApiRouter.GET("ReStartPltAllNetwork", api.ReStartPltAllNetwork)
		ApiRouter.GET("InitPltAllNetwork", api.InitPltAllNetwork)
		ApiRouter.GET("StartPltAllNetwork", api.StartPltAllNetwork)
		ApiRouter.GET("StopPltAllNetwork", api.StopPltAllNetwork)
		ApiRouter.GET("ResetPltAllNetwork", api.ResetPltAllNetwork)
		ApiRouter.GET("ClearPltAllNetwork", api.ClearPltAllNetwork)
		ApiRouter.GET("GetAllNodes", api.GetAllNodes)

		//staking
		ApiRouter.GET("GetAllStakingInformation", api.GetAllStakingInformation)

		//crosschain
		ApiRouter.POST("sendPltToEth", api.SendPltToEth)
		ApiRouter.POST("sendNFTToEth", api.SendNFTToEth)
	}
}
