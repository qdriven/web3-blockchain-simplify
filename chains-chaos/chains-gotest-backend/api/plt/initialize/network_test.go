package initialize

import (
	"chains-gotest-backend/api/plt/api"
	"chains-gotest-backend/api/plt/tracker"
	"chains-gotest-backend/core"
	"chains-gotest-backend/global"
	"chains-gotest-backend/initialize"
	"testing"
)

func init()  {
	SetupTestCases("/Users/Patrick/go/src/github.com/palettechain/onRobot/service/config.json")
	global.GVA_VP = core.Viper("/Users/Patrick/go/src/github.com/palettechain/gin-vue-admin/server/plt-server/config.yaml")          // 初始化Viper
	global.GVA_LOG = core.Zap()           // 初始化zap日志库
	global.GVA_DB = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.GVA_DB) // 初始化表
}

//func TestStartAllNetwork(t *testing.T) {
//	core.StartAllNetwork()
//}
//
//func TestStopAllNetwork(t *testing.T){
//	core.StopAllNetwork()
//}
//
//func TestInitValidatorWorks(t *testing.T){
//	//core.InitValidatorNetwork()
//	core.GetValidators()
//	//core.Transfer()
//	//core.AdminBalance()
//	//core.Proposal()
//	//result, err :=utils.PackMethod(sdk.PLTABI,"balanceOf","test")
//	//if(err!=nil){
//	//	log.Info("this is error,{}",err)
//	//}
//	//fmt.Println(result)
//}

//func TestHex(t *testing.T) {
//	//result:= utils.Big2Address(big.NewInt(1))
//	bigNum :=utils.Address2Big(common.HexToAddress("0x0000000000000000000000010000000000000005"))
//	fmt.Println(bigNum)
//	//log.Info(result.String())
//	//log.Info(addr.String())
//
//	result, err :=core.Client.GetProposal(utils.Big2Address(bigNum),"latest")
//	if(err!=nil){
//		fmt.Println(err)
//	}
//	fmt.Println(result)
//}

func TestSendEth(t *testing.T) {
	api.DemoPLTCrossChainFunction()
}

func TestQueryLog(t *testing.T) {
	//tracker.FilterRewardEvent(174585,174585)
	//tracker.FilterPLTTransferLogs(174585,174585)
	tracker.FilterStakeEvents(174500,204585)
}

//0x4fb7e37e4abcf2eb97098579493ad380e8f7178ecbd6d8ef472f0782373332ba
//0xa765d90ec5bdcbdd8f28c6ba6ddfd1dcc4f8d1da85536bfadeba0a3ad2bb9b8d
//1. Getting Mapping setting
//2. Getting Asset Mapping
//3. Send Transaction
