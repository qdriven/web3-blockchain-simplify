package tracker

import (
	"gin-vue-admin/chain/config"
	"gin-vue-admin/chain/core"
	gCore "gin-vue-admin/core"
	chainInitialize "gin-vue-admin/chain/initialize"
	"gin-vue-admin/chain/pkg/log"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	defer time.Sleep(time.Second)
	log.InitLog(1, log.Stdout)
	config.Init("/Users/Patrick/go/src/github.com/palettechain/onRobot/service/config.json") //TODO: reading from setting
	core.Initialize()
	global.GVA_VP = gCore.Viper()          // 初始化Viper
	global.GVA_LOG = gCore.Zap()           // 初始化zap日志库
	global.GVA_DB = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.GVA_DB) // 初始化表
	chainInitialize.ChainDataBase(global.GVA_DB)  //initialize blockchain database
	// 程序结束前关闭数据库链接
	db, _ := global.GVA_DB.DB()

	defer db.Close()
}

func TestFilterTransferEvent(t *testing.T) {
	FilterStakeEvents(0,40000)
}