package api

import (
	"chains-gotest-backend/api/evm/log"
	"chains-gotest-backend/api/plt/config"
	"chains-gotest-backend/api/plt/core"
	"chains-gotest-backend/api/plt/model"
	"chains-gotest-backend/api/plt/pkg/hdwallet"
	"chains-gotest-backend/global"
	"chains-gotest-backend/model/response"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

/**
1. Init many accounts
2. Leverage These accounts for different testing purpose,like performance testing/staking/
*/

type TestAccountRequest struct {
	Num int `form:"num" json:"num"`
}

// @Tags Account API
// @Summary Create TestAccounts
// @accept application/json
// @Produce application/json
// @Param data body TestAccountRequest true "Num"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/createTestAccounts [post]
func CreateTestAccounts(ctx *gin.Context) {
	/**
	创建账户: Testing Account for both EVM based wallet
	*/
	request := &TestAccountRequest{}
	_ = ctx.ShouldBind(&request)
	num := request.Num
	var accountIndex int
	//Tracking Create Account Events to save accounts information, all accounts with same password
	global.GVA_DB.Raw("select max(acc_index) from test_accounts").Scan(&accountIndex)
	testAccounts := make([]*model.TestAccount, 0)
	w := hdwallet.DefaultWallet()
	for i := 0; i < num; i++ {
		acc, err := hdwallet.Drive(accountIndex + i + 1)
		if err != nil {
			response.FailWithDetailed(err, "create account failed", ctx)
		}

		privateKey, _ := w.PrivateKeyHex(acc)
		testAccount := &model.TestAccount{
			Address:        acc.Address.Hex(),
			AccIndex:       accountIndex + i + 1,
			PrivateKey:     privateKey,
			DerivationPath: acc.URL.Path,
		}
		global.GVA_DB.Save(testAccount)
		testAccounts = append(testAccounts, testAccount)
	}

	response.OkWithData(testAccounts, ctx)
}

// @Tags Account API
// @Summary Create GetInitNodeInfo
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/getInitNodeInfo [get]
func GetInitNodeInfo(ctx *gin.Context) {
	nodes := config.Conf.AllNodes()
	for _, node := range nodes {
		testAccount := &model.TestAccount{
			Address:    node.Address,
			PrivateKey: node.NodeKey,
		}
		log.Info("test account info is ", testAccount)
		global.GVA_DB.Save(testAccount)
		stakeAccount := &model.TestAccount{
			Address: node.StakeAccount,
			PrivateKey: node.StakePrivateKeyString(),
		}
		log.Info("staking account info is ", stakeAccount)
		global.GVA_DB.Save(stakeAccount)
		//currentBlockNum := core.Client.GetBlockNumber()
		//balance, _ := core.Client.BalanceOf(common.HexToAddress(node.Address), "latest")
		//accountBalance := &model.TestAccountBalance{
		//	CurrentBlockNum: int(currentBlockNum),
		//	Balance:         balance.Uint64(),
		//	Address:         node.Address,
		//}
		//global.GVA_DB.Save(accountBalance)
	}
	adminAccount := &model.TestAccount{
		Address:    config.Conf.AdminAccount,
		PrivateKey: string(crypto.FromECDSA(core.Client.Key)),
		IsAdmin:    true,
	}
	global.GVA_DB.Save(adminAccount)
	currentBlockNum := core.Client.GetBlockNumber()
	balance, _ := core.Client.BalanceOf(common.HexToAddress(config.Conf.AdminAccount), "latest")
	//TODO: Different ChainId
	adminBalance := &model.TestAccountBalance{
		CurrentBlockNum: int(currentBlockNum),
		Balance:         balance.Uint64(),
		Address:         config.AdminAddr.String(),
		Symbol:          "PLT",
	}
	global.GVA_DB.Save(adminBalance)
	response.OkWithData(nodes, ctx)
}
