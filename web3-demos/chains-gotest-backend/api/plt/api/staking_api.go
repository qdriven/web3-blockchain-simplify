package api

import (
	"chains-gotest-backend/api/plt/pkg/log"
	"chains-gotest-backend/api/plt/service"
	"chains-gotest-backend/global"
	"chains-gotest-backend/model/response"
	"github.com/ethereum/go-ethereum/contracts/native/plt"
	"github.com/ethereum/go-ethereum/contracts/native/utils"
	"github.com/ethereum/go-ethereum/contracts/native/utils/decimal"
	"github.com/gin-gonic/gin"
	"math/big"
)

// @Tags PltTokenApi
// @Summary Get All Staking Information
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/GetAllStakingInformation [get]
func GetAllStakingInformation(c *gin.Context) {

	result :=service.GetAllValidatorStaking()
	response.OkWithData(result,c)
}

type StakingAccountRecord struct {
	StakingAmount string
	Validator string
	StakeAccount string
	UserAccount string
}

func GetDistributionResultByHeight(c *gin.Context){
	//Get All Staking Event Before A Given Height
	//Getting Report
	result:=make([]StakingAccountRecord,0)
	sql :="SELECT\n\tsum(\n\tCASE\n\t\t\t\n\t\t\tWHEN `REVOKE` = 0 THEN\n\t\t\tDATA ELSE - DATA \n\t\tEND \n\t\t),\n\t\t`owner`,\n\t\tstake_account,\n\t\tvalidator \n\tFROM\n\t\tstaking_events \n  where current_block_height < ?\n\tGROUP BY\n\t\t`owner`,\n\t\t`stake_account`,\n\t\tvalidator"
	//scan
	global.GVA_DB.Raw(sql,0).Scan(&result)
	BaseAmount := plt.MulOnePLT(decimal.NewFromFloat(60))
	a:=0.2
	//b:=0.8
	BaseAward:=BaseAmount.Mul(decimal.NewFromFloat(a))
	log.Info("base reward is ",BaseAward)
	ValidatorStakingMap := make(map[string]*big.Int)
	for i, record := range result {
		log.Info("process record ",i)
		value, _ :=new(big.Int).SetString(record.StakingAmount,18)
		ValidatorStakingMap[record.Validator]= utils.SafeAdd(ValidatorStakingMap[record.Validator],value)
	}
	//Validator Address

}