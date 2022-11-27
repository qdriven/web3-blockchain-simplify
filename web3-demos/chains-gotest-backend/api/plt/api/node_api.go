package api

import (
	"chains-gotest-backend/api/plt/config"
	"chains-gotest-backend/api/plt/core"
	chainData "chains-gotest-backend/api/plt/model"
	"chains-gotest-backend/api/plt/pkg/log"
	"chains-gotest-backend/global"
	"chains-gotest-backend/model/response"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Tags NodeAPI
// @Summary Get Init Nodes
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/getAllInitializedNodes [get]
func GetAllInitializedNodes(c *gin.Context) {
	/**
	1. NODE Type: 0: GENESIS 1: Validator 2: Other Nodes
	   "GenesisNodeNumber": 5, first five is genesis node
	    "ValidatorsNumber": 3
	*/
	nodes := config.Conf.AllNodes()
	pnodes := make([]*chainData.PaletteNode, 0)
	for i, node := range nodes {
		log.Info("current node number ", i)
		address := common.HexToAddress(node.Address)
		validatorResult := core.Client.CheckValidator(address, "latest")
		nodeType := 0
		if validatorResult {
			nodeType = 1
		}
		totalStakeAmount := core.Client.GetValidatorTotalStakeAmount(address, "latest")
		var pnode = &chainData.PaletteNode{
			CurrentBlockNum: strconv.FormatUint(0, 10),
			NodeAddress:     node.Address,
			NodeHost:        node.Host,
			NodeRPCPort:     node.RPCPort,
			NodeP2PPort:     node.P2PPort,
			NodePrivateKey:  node.NodeKey,
			NodeDir:         node.NodeDirPath(),
			NodeType:        nodeType,
			StakeAmount:     totalStakeAmount.String(),
			SeedNode:        1,
			Active: 1,
		}
		global.GVA_DB.Save(pnode) //TODO: 不能重复
		pnodes = append(pnodes, pnode)
	}
	spareNodes := config.Conf.SpareNodes()
	for i, spareNode := range spareNodes {
		log.Info("current spare node number ", i)
		var pnode = &chainData.PaletteNode{
			CurrentBlockNum: strconv.FormatUint(0, 10),
			NodeAddress:     spareNode.Address,
			NodeHost:        spareNode.Host,
			NodeRPCPort:     spareNode.RPCPort,
			NodeP2PPort:     spareNode.P2PPort,
			NodePrivateKey:  spareNode.NodeKey,
			NodeDir:         spareNode.NodeDirPath(),
			NodeType:        2,
			StakeAmount:     "0",
			SeedNode:        0,
			Active: 1,
		}
		global.GVA_DB.Save(pnode) //TODO: 不能重复
		pnodes = append(pnodes, pnode)
	}
	response.OkWithData(pnodes, c)
}

type NodeRequest struct {
	NodeAddresses []string
}
type NodeFunction func(node *config.Node)

func nodeOperation(c *gin.Context,operation NodeFunction){
	request := &NodeRequest{}
	_ = c.ShouldBind(&request)
	for i, address := range request.NodeAddresses {
		log.Info("start current node,",i)
		node := config.Conf.GetNodeByAddress(address)
		operation(node)
	}
	response.OkWithData(request,c)
}

// @Tags NodeAPI
// @Summary start given nodes
// @accept application/json
// @Produce application/json
// @Param data body NodeRequest true "NodeAddresses"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/startNode [post]
func StartNode(c *gin.Context) {
	//TODO: remove to database
	nodeOperation(c,core.ExecStopNode)
}
// @Tags NodeAPI
// @Summary start given nodes
// @accept application/json
// @Produce application/json
// @Param data body NodeRequest true "NodeAddresses"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/stopNode [post]
func StopNode(c *gin.Context) {
	nodeOperation(c,core.ExecStopNode)
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary start given nodes
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/ResetPltGenesisNetwork [get]
func ResetPltGenesisNetwork(c *gin.Context){
	core.ResetGenesisNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary reset genesis nodes
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/ReStartPltGenesisNetwork [get]
func ReStartPltGenesisNetwork(c *gin.Context){
	core.ReStartGenesisNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary int genesis network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/InitPltGenesisNetwork [get]
func InitPltGenesisNetwork(c *gin.Context){
	core.InitGenesisNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary start genesis network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/StartPltGenesisNetwork [get]
func StartPltGenesisNetwork(c *gin.Context){
	core.StartGenesisNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary stop genesis network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/StopPltGenesisNetwork [get]
func StopPltGenesisNetwork(c *gin.Context){
	core.StopGenesisNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary clear genesis network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/ClearPltGenesisNetwork [get]
func ClearPltGenesisNetwork(c *gin.Context){
	core.ClearGenesisNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary reset validator network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/ResetPltValidatorNetwork [get]
func ResetPltValidatorNetwork(c *gin.Context){
	core.ResetValidatorNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary restart validator network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/ReStartPltValidatorNetwork [get]
func ReStartPltValidatorNetwork(c *gin.Context){
	core.ReStartValidatorNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary start validator network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/StartPltValidatorNetwork [get]
func StartPltValidatorNetwork(c *gin.Context){
	core.StartValidatorNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary stop validator network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/StopPltValidatorNetwork [get]
func StopPltValidatorNetwork(c *gin.Context){
	core.StopValidatorNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary clear validator network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/ClearPltValidatorNetwork [get]
func ClearPltValidatorNetwork(c *gin.Context){
	core.ClearValidatorNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary init all network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/ReStartPltAllNetwork [get]
func ReStartPltAllNetwork(c *gin.Context){
	core.ReStartAllNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary init all network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/InitPltAllNetwork [get]
func InitPltAllNetwork(c *gin.Context){
	core.InitAllNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary start plt all network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/StartPltAllNetwork [get]
func StartPltAllNetwork(c *gin.Context){
	core.StartAllNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary stop all plt network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/StopPltAllNetwork [get]
func StopPltAllNetwork(c *gin.Context){
	core.StopAllNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary reset all  network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/ResetPltAllNetwork [get]
func ResetPltAllNetwork(c *gin.Context){
	core.ResetAllNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary clear all  network
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/ClearPltAllNetwork [get]
func ClearPltAllNetwork(c *gin.Context){
	core.ClearAllNetwork()
	response.Ok(c)
}

// @Tags NodeAPI
// @Summary get working nodes
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/GetAllNodes [get]
func GetAllNodes(c *gin.Context){
	result :=make([]chainData.PaletteNode,0)
	global.GVA_DB.Raw("select * from palette_node").Scan(result)
	response.OkWithData(result,c)
}