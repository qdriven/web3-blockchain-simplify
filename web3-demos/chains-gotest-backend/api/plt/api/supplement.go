package api

import (
	"chains-gotest-backend/api/plt/config"
	"chains-gotest-backend/api/plt/core"
	"chains-gotest-backend/api/plt/model"
	"chains-gotest-backend/api/plt/pkg/hdwallet"
	"chains-gotest-backend/api/plt/pkg/sdk"
	"chains-gotest-backend/global"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

//TODO: 有全局变量
func SetTxAccount(From string) common.Address {
	accountAddress := config.AdminAddr
	if From != "" { // is empty, then use admin account
		var testAccount model.TestAccount
		global.GVA_DB.Where("address= ?", From).First(&testAccount)
		w := hdwallet.DefaultWallet()
		account, _ := hdwallet.Drive(testAccount.AccIndex)
		privateKey, _ := w.PrivateKey(account)
		accountAddress = account.Address
		core.Client.Reset(privateKey)
	}
	return accountAddress
}

func CreateNewClient(From string) *sdk.Client {
	if From != "" { // is empty, then use admin account
		if From == config.AdminAddr.String() {
			return core.Client
		}
		var testAccount model.TestAccount
		global.GVA_DB.Where("address= ?", From).First(&testAccount)
		if testAccount.DerivationPath=="" && testAccount.PrivateKey!=""{
			for _, node := range config.Conf.Nodes {
				if strings.ToLower(node.Address)==strings.ToLower(From){
					return core.NewClient(node.NodeKey)
				}
			}
			//use private key directly
			return core.NewClient(testAccount.PrivateKey)
		}
		w := hdwallet.DefaultWallet()
		account, _ := hdwallet.Drive(testAccount.AccIndex)
		privateKey, _ := w.PrivateKey(account)
		return sdk.NewSender(config.BakConf.BaseRPCUrl, privateKey)
	}
	return core.Client
}

//2000000000000100001
//1000000000000100001