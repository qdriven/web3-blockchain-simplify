package initializer

import (
	"chains-gotest-backend/api/evm/models"
	"chains-gotest-backend/core"
	"chains-gotest-backend/global"
	"chains-gotest-backend/initialize"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"strings"
)

func GetProjectPath() string {
	path, _ := os.Getwd()
	index := strings.Index(path, "go-chains")
	return path[0:index] + "go-chains/"
}

func GetProjectConfigPath() string {
	basePath := GetProjectPath()
	return basePath + "config.yaml"
}

func Setup4Test() {
	configPath := GetProjectConfigPath()
	global.GVA_VP = core.Viper(configPath) // 初始化Viper
	global.GVA_LOG = core.Zap()            // 初始化zap日志库
	global.GVA_DB = initialize.Gorm()      // gorm连接数据库
	initialize.MysqlTables(global.GVA_DB)  // 初始化表
	// 程序结束前关闭数据库链接
	MigrateChainsTables(global.GVA_DB)
}

func MigrateChainsTables(db *gorm.DB) {

	err :=db.AutoMigrate(
		models.ChainMetaData{},
		models.CrossChainTx{},
		models.CrossChainAssetMapping{},
		models.CrossChainLockProxy{},
		models.CrossChainProxyHashMap{},
		models.CoinMetaData{},
		models.EvmTransactionRecord{},
	)
	if err != nil {
		global.GVA_LOG.Error("register blockchain table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register blockchain tables success")

}
