package api

import (
	"chains-gotest-backend/api/evm/ethereum/clients"
	"chains-gotest-backend/api/evm/log"
	"chains-gotest-backend/model/response"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)


//api generation by abi json
//template for invocation
//
var (
	PLT_ETH_LOCKPROXY = "0x33082b802c4099040420533a9234bD80D71726ba"
	PLT_ETH_ADDRESS   = "0x1AC77d322e3347B8de181304B7c118a7e89c6Bc1" //PLT Token Address
)

type PltCrossChainRequest struct {
	FromAddress   string
	ToAddress     string
	Amount        string
	ApproveAmount string
	Decimal       uint64
}
// @Tags cross-chain
// @Summary send plt to ETH
// @accept application/json
// @Produce application/json
// @Param data body PltCrossChainRequest true "FromAddress,ToAddress,Amount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/approveAndSentPltToPlt [post]
func ApproveAndSendPltToPaletteChainTx(c *gin.Context) {
	request := &PltCrossChainRequest{}
	c.ShouldBind(request)
	tx, err := ApproveAndSendPltToPaletteChain(request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if tx != nil {
		response.OkWithData(tx.Hash(), c)
	} else {
		response.FailWithMessage("transaction is not generated", c)
	}
}


// @Tags cross-chain
// @Summary send plt to ETH
// @accept application/json
// @Produce application/json
// @Param data body PltCrossChainRequest true "FromAddress,ToAddress,Amount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chains/sendPltToPalletChain [post]
func SentPltToPlt(c *gin.Context) {
	request := &PltCrossChainRequest{}
	_ = c.ShouldBind(request)
	tx, err := SendPltToPalletChain(request)
	log.Info("transaction hash is ",tx.Hash().String())
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if tx != nil {
		response.OkWithData(tx.Hash(), c)
	} else {
		response.FailWithMessage("transaction is not generated", c)
	}
}

func SendPltToPalletChain(request *PltCrossChainRequest) (*types.Transaction, error) {
	url := "https://ropsten.infura.io/v3/40961bdd0891462db115cde9ce6a3ad2"
	//privateKey := strings.ReplaceAll("0x4f17afd61d017aa256a94d4edcd3ec3241e75ef6ca5309787d2ad8ca96f1d613", "0x", "")
	privateKey := "aec101ecdb5c86931e0ca5e635824f0d0f05240760c01ddee64be90b2a2608a9" //TODO: Make it different Address
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Error(err)
	}
	evmClient := clients.NewEvmClient(url, key)
	ccRequest := &clients.CrossChainRequestContext{
		LockProxyAddress: PLT_ETH_LOCKPROXY,
		FromAddressHex:   PLT_ETH_ADDRESS,
		Amount:           request.Amount,
		ToAddressHex:     request.ToAddress,
		Decimal:          18,
		ToChainId:        uint64(106),
	}

	return evmClient.SendLockTx(ccRequest)
}

func ApproveAndSendPltToPaletteChain(request *PltCrossChainRequest) (*types.Transaction, error) {

	//url := "http://127.0.0.1:8545"
	url := "https://ropsten.infura.io/v3/40961bdd0891462db115cde9ce6a3ad2" //TODO: Move to Redis or Database
	//privateKey := strings.ReplaceAll("0x4f17afd61d017aa256a94d4edcd3ec3241e75ef6ca5309787d2ad8ca96f1d613", "0x", "")
	privateKey := "aec101ecdb5c86931e0ca5e635824f0d0f05240760c01ddee64be90b2a2608a9" //TODO: Make it different Address
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Error(err)
	}
	evmClient := clients.NewEvmClient(url, key)
	currentBalance, _ := evmClient.BalanceAt(context.Background(), evmClient.Address(), nil)
	fmt.Println(currentBalance)
	ccRequest := &clients.CrossChainRequestContext{
		LockProxyAddress: PLT_ETH_LOCKPROXY,
		FromAddressHex:   PLT_ETH_ADDRESS,
		Amount:           request.Amount,
		ToAddressHex:     request.ToAddress,
		Decimal:          18,
		ToChainId:        uint64(106),
	}

	return evmClient.SendApproveAndLockTx(ccRequest)

}
