package api

import (
	"chains-gotest-backend/api/evm/log"
	response2 "chains-gotest-backend/api/plt/api/response"
	"chains-gotest-backend/api/plt/config"
	"chains-gotest-backend/api/plt/core"
	"chains-gotest-backend/api/plt/model"
	"chains-gotest-backend/global"
	"chains-gotest-backend/model/response"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native/plt"
	"github.com/ethereum/go-ethereum/contracts/native/utils/decimal"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync/atomic"
	"time"
)

type AccountTransferRequest struct {
	From         string `form:"from" json:"from"`
	To           string `form:"to" json:"to"`
	Amount       string `form:"amount" json:"amount"`
	OwnerAddress string `form:"ownerAddress" json:"ownerAddress"`
}

// @Tags PltTokenApi
// @Summary Get PLT Token summary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/pltMeta [get]
func GetPltMeta(ctx *gin.Context) {

	client := CreateNewClient("")
	p := &response2.PltCoinMetaResponse{}
	p.Name, _ = client.PLTName()
	totalSupply, _ := client.PLTTotalSupply("latest")
	p.TotalSupply = totalSupply.String()
	p.Symbol, _ = client.PLTSymbol()
	decim, _ := client.PLTDecimals()
	p.Decimal = strconv.FormatUint(decim, 10)
	response.OkWithData(p, ctx)
}

// @Tags PltTokenApi
// @Summary Transfer PltToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body AccountTransferRequest true "from, to, Amount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/pltTransfer [post]
func PltTransfer(ctx *gin.Context) {
	//account configuration to performance testing
	//Sending Metrics to a centered database
	request := &AccountTransferRequest{}
	_ = ctx.ShouldBind(&request)
	accountAddress := config.AdminAddr
	client := CreateNewClient(request.From)
	balance, _ := client.BalanceOf(accountAddress, "latest")
	log.Info("current balance is ", balance)
	amount, _ := strconv.ParseFloat(request.Amount, 18)
	log.Info("transfer amount is ", amount)
	hash, _ := client.PLTTransfer(common.HexToAddress(request.To), plt.MultiFloatPLT(amount).BigInt())
	result := make(map[string]interface{})
	result["hash"] = hash
	response.OkWithData(result, ctx)
}

// @Tags PltTokenApi
// @Summary Transfer From PltToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body AccountTransferRequest true "from, to, Amount,OwnerAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/pltTransferFrom [post]
func PltSafeTransferFrom(ctx *gin.Context) {
	//account configuration to performance testing
	//Sending Metrics to a centered database
	request := &AccountTransferRequest{}
	_ = ctx.ShouldBind(&request)
	client := CreateNewClient(request.OwnerAddress)
	balance, _ := client.PLTAllowance(client.Address(), common.HexToAddress(request.From), "latest")
	log.Info("current balance is ", balance)
	amount, _ := strconv.ParseFloat(request.Amount, 18)
	log.Info("transfer amount is ", amount)
	hash, err := client.PLTTransferFrom(common.HexToAddress(request.From), common.HexToAddress(request.To),
		plt.MultiFloatPLT(amount).BigInt())
	result := make(map[string]interface{})
	if err != nil {
		result["error"] = err.Error()
	}
	result["hash"] = hash
	response.OkWithData(result, ctx)
}

type PltApproveRequest struct {
	OwnerAddress   string
	SpenderAddress string
	Amount         string
}

// @Tags PltTokenApi
// @Summary Approve PltToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PltApproveRequest true "OwnerAddress, SpenderAddress, Amount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/pltApprove [post]
func PltApprove(ctx *gin.Context) {
	//account configuration to performance testing
	//Sending Metrics to a centered database
	request := &PltApproveRequest{}
	_ = ctx.ShouldBind(&request)
	client := CreateNewClient(request.OwnerAddress)
	balance, _ := client.BalanceOf(client.Address(), "latest")
	log.Info("current balance is ", balance)
	amount, _ := strconv.ParseFloat(request.Amount, 18)
	log.Info("transfer amount is ", amount)
	hash, _ := client.PLTApprove(common.HexToAddress(request.SpenderAddress), plt.MultiFloatPLT(amount).BigInt())
	result := make(map[string]interface{})
	result["hash"] = hash
	response.OkWithData(result, ctx)
}

// @Tags PltTokenApi
// @Summary Approve PltToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PltApproveRequest true "OwnerAddress,FromAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/pltAllowance [post]
func PltAllowance(ctx *gin.Context) {
	//account configuration to performance testing
	//Sending Metrics to a centered database
	request := &PltApproveRequest{}
	_ = ctx.ShouldBind(&request)
	client := CreateNewClient(request.OwnerAddress)
	allowance, _ := client.PLTAllowance(client.Address(), common.HexToAddress(request.SpenderAddress), "latest")
	result := make(map[string]interface{})
	result["allowance"] = allowance
	response.OkWithData(result, ctx)
}

// @Tags PltTokenApi
// @Summary BalanceOf PltToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body AccountTransferRequest true "from"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/pltBalanceOf [post]
func PltBalanceOf(ctx *gin.Context) {
	request := &AccountTransferRequest{}
	_ = ctx.ShouldBind(&request)
	balance, _ := core.Client.BalanceOf(common.HexToAddress(request.From), "latest")
	var testAccountBalance = &model.TestAccountBalance{}
	global.GVA_DB.Where("address=?", request.From).Scan(&testAccountBalance)
	if testAccountBalance != nil {
		testAccountBalance.Balance = balance.Uint64()
		global.GVA_DB.Updates(testAccountBalance)
	} else {
		testAccountBalance.Address = request.From
		testAccountBalance.Balance = balance.Uint64()
		global.GVA_DB.Save(testAccountBalance)
	}
	response.OkWithData(balance, ctx)
}

//TODO：
//1. name/symbol/totalSupply/balanceOf/approve/
//2. batch transfer
type BatchTransferRequest struct {
	AccountNum int `form:"accountNum" json:"accountNum"`
}

// @Tags PltTokenApi
// @Summary batch transfer PltToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body BatchTransferRequest true "accountNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/pltBatchTransfer [post]
func BatchTransfer(ctx *gin.Context) {
	request := &BatchTransferRequest{}
	ctx.ShouldBind(request)
	accountNum := request.AccountNum
	accounts := make([]string, 0)
	hashes := make([]string, 0)
	global.GVA_DB.Raw("select address from test_accounts where derivation_path is not null limit ? ", accountNum).Scan(&accounts)
	balanceBeforeSend, _ := core.Client.BalanceOf(common.HexToAddress(config.Conf.AdminAccount), "latest")
	for i := 0; i < accountNum; i++ {
		txHash, _ := core.Client.PLTTransfer(common.HexToAddress(accounts[i]), plt.MultiPLT(1))
		core.WaitBlocks(1)
		hashes = append(hashes, txHash.Hex())
	}
	afterBeforeSend, _ := core.Client.BalanceOf(common.HexToAddress(config.Conf.AdminAccount), "latest")
	log.Info("beforeBalance,{},afterBalance {}", balanceBeforeSend, afterBeforeSend)
	response.OkWithData(hashes, ctx)
}

type PerformanceRequest struct {
	UserNum       int `form:"userNum" json:"userNum"`
	ExecutionTime int `form:"executionTime" json:"executionTime"`
	RampUpTime    int `form:"rampUpTime" json:"rampUpTime"`
}

// @Tags PltTokenApi
// @Summary BalanceOf PltToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PerformanceRequest true "UserNum,ExecutionTime,RampUpTime"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/pltPerformance [post]
func PltPerformance(ctx *gin.Context) {
	request := &PerformanceRequest{}
	ctx.ShouldBind(request)
	et := request.ExecutionTime
	rut := request.RampUpTime
	tc := request.UserNum //thread nums
	accounts := make([]model.TestAccount, 0)
	userNum := request.UserNum
	global.GVA_DB.Raw("select * from test_accounts where derivation_path is not null limit ? ", userNum).Scan(&accounts)

	//Check if execution time is more than ramp up time
	if et*60 < rut {
		log.Error("Total execution time needs to be more than ramp up time")
	}

	waitTime := rut / tc

	log.Info("Execution will happen with %d users with a ramp up time of %d seconds for %d minutes\n", tc, rut, et)

	tchan := make(chan int)
	go func(c chan<- int) { //Sender
		for ti := 1; ti <= tc; ti++ {
			log.Infof("Thread Count %d", ti)
			c <- ti
			time.Sleep(time.Duration(waitTime) * time.Second)
		}
	}(tchan)

	timeout := time.After(time.Duration(et*60) * time.Second)
	var txCount uint64
loop:
	for {
		select { //Select blocks the flow until one of the channels receives a message
		case <-timeout: //receives a msg when execution duration is over
			log.Infof("Execution completed")
			close(tchan)
			break loop

		case ts := <-tchan: //receives a message when a new user thread has to be initiated
			log.Infof("Thread No %d started", ts)
			go func(t int) {
				//This is the place where you add all your tests
				//In my case they were making rpc calls over rabbitmq with random inputs
				//They keep running till the end of execution
				for {
					client := core.NewClient(accounts[t-1].PrivateKey)
					txHash, _ := client.PLTTransfer(common.HexToAddress(accounts[t-1].Address), plt.MultiFloatPLT(0.00001).BigInt())
					core.WaitBlocks(1)
					atomic.AddUint64(&txCount, 1)
					log.Info("transaction hash is ", txHash.Hex())
				}
			}(ts)
		}
	}
	tps := decimal.NewFromInt(int64(txCount)).Div(decimal.NewFromInt(int64(et)).Mul(decimal.NewFromInt(60)))
	//counter/time
	pMetrics := &model.PMetrics{
		TotalCount:         txCount,
		TotalExecutionTime: et,
		AvgTps:             tps.String(),
	}
	global.GVA_DB.Save(pMetrics)
	response.OkWithData(pMetrics, ctx)
}
