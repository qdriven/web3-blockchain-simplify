package core

import (
	"chains-gotest-backend/global"
	"chains-gotest-backend/initialize"
	"fmt"
	"go.uber.org/zap"
	"time"
)
/**
Init Linux Server
 */
type server interface {
	ListenAndServe() error
}
/**
Init Windows Server
 */
func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	Welcome to User Go-Chains Server for easy testing
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
