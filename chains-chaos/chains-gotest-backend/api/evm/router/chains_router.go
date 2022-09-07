package router

import (
	"github.com/gin-gonic/gin"
	"go-chains/chains/api"
	"go-chains/middleware"
)

func InitChainsApis(Router *gin.RouterGroup) {

	ApiRouter := Router.Group("chains").Use(middleware.OperationRecord())
	{
		ApiRouter.POST("chainMetaData", api.CreateChainMetaData) // Get All Global Parameters
		ApiRouter.GET("chainMetaData", api.QueryChainMetaData)
		ApiRouter.POST("approveAndSentPltToPlt", api.ApproveAndSendPltToPaletteChainTx)
		ApiRouter.POST("sendPltToPalletChain", api.SentPltToPlt)
	}
}
