package api

import (
	"chains-gotest-backend/api/evm/api/validator"
	"chains-gotest-backend/api/evm/models"
	"chains-gotest-backend/api/evm/service"
	"chains-gotest-backend/global"
	"chains-gotest-backend/model/response"
	"chains-gotest-backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ChainMetaDataApi
// @Summary 创建Chain的基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body models.ChainMetaData true "	ChainName,ChainI,ChainType,RPCUrl string"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/chainMetaData [post]
func CreateChainMetaData(c *gin.Context) {
	var request = &models.ChainMetaData{}
	_ = c.ShouldBindJSON(&request)
	if err := utils.Verify(request, validator.ChainMetaValidator); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateChainMetaData(request); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ChainMetaDataApi
// @Summary 获取Chain MetaData信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/chainMetaData [get]
func QueryChainMetaData(c *gin.Context) {
	result := service.QueryAllChainMetaData()
	response.OkWithData(result, c)
}


// @Tags ChainMetaDataApi
// @Summary 创建Chain的基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body models.ChainMetaData true "	ChainName,ChainI,ChainType,RPCUrl string"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/coinMetaData [post]
func CreateCoinMetaData(c *gin.Context) {
	var request = &models.CoinMetaData{}
	_ = c.ShouldBindJSON(&request)
	if err := utils.Verify(request, validator.CoinMetaDataValidator); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateCoinMetaData(request); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ChainMetaDataApi
// @Summary 获取Chain MetaData信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/coinMetaData [get]
func QueryCoinMetaData(c *gin.Context) {
	result := service.QueryAllChainMetaData()
	response.OkWithData(result, c)
}


// @Tags ChainMetaDataApi
// @Summary 创建Chain的基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body models.ChainMetaData true "	ChainName,ChainI,ChainType,RPCUrl string"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/CrossChainAssetMapping [post]
func CreateCrossChainAssetMapping(c *gin.Context) {
	var request = &models.CrossChainAssetMapping{}
	_ = c.ShouldBindJSON(&request)
	if err := utils.Verify(request, validator.CrossChainAssetMappingValidator); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateCrossChainAssetMapping(request); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ChainMetaDataApi
// @Summary 获取Chain MetaData信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/queryCrossChainAssetMapping [get]
func QueryCrossChainAssetMapping(c *gin.Context) {
	result := service.QueryAllCrossChainAssetMapping()
	response.OkWithData(result, c)
}

// @Tags ChainMetaDataApi
// @Summary 创建Chain的基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body models.ChainMetaData true "	ChainName,ChainI,ChainType,RPCUrl string"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/CrossChainLockProxy [post]
func CreateCrossChainLockProxy(c *gin.Context) {
	var request = &models.CrossChainLockProxy{}
	_ = c.ShouldBindJSON(&request)
	if err := utils.Verify(request, validator.CrossChainLockProxyValidator); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateCrossChainLockProxy(request); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ChainMetaDataApi
// @Summary 获取Chain MetaData信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/queryCrossChainLockProxy [get]
func QueryCrossChainLockProxy(c *gin.Context) {
	result := service.QueryAllCrossChainLockProxy()
	response.OkWithData(result, c)
}


// @Tags ChainMetaDataApi
// @Summary 创建Chain的基础信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body models.ChainMetaData true "	ChainName,ChainI,ChainType,RPCUrl string"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/CrossChainProxyHashMap [post]
func CreateCrossChainProxyHashMap(c *gin.Context) {
	var request = &models.CrossChainProxyHashMap{}
	_ = c.ShouldBindJSON(&request)
	if err := utils.Verify(request, validator.CrossChainProxyHashMapValidator); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateCrossChainProxyHashMap(request); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ChainMetaDataApi
// @Summary 获取Chain MetaData信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/queryCrossChainProxyHashMap [get]
func QueryCrossChainProxyHashMap(c *gin.Context) {
	result := service.QueryAllCrossChainProxyHashMap()
	response.OkWithData(result, c)
}