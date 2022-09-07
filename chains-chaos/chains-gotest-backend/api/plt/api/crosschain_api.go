package api

import (
	"chains-gotest-backend/api/plt/model"
	"chains-gotest-backend/api/plt/pkg/log"
	"chains-gotest-backend/global"
	"chains-gotest-backend/model/response"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/native/utils/decimal"
	"github.com/gin-gonic/gin"
	"math/big"
)

type PltCrossChainRequest struct {
	FromAddress string
	ToAddress   string
	Amount      string
}

// @Tags cross-chain
// @Summary send plt to ETH
// @accept application/json
// @Produce application/json
// @Param data body PltCrossChainRequest true "FromAddress,ToAddress,Amount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/sendPltToEth [post]
func SendPltToEth(c *gin.Context) {
	request := &PltCrossChainRequest{}
	_ = c.ShouldBind(&request)

	client := CreateNewClient(request.FromAddress)
	currentBalance, _ := client.BalanceOf(client.Address(), "latest")
	log.Info(currentBalance)
	toAddress := common.HexToAddress(request.ToAddress) //to address
	amount, _ := decimal.NewFromString(request.Amount)
	pltDecimal, _ := decimal.NewFromString("1000000000000000000")
	tx, err := client.Lock(2, toAddress, amount.Mul(pltDecimal).BigInt())
	crossChainMode := &model.CrossChainTx{
		TxHash:      tx.String(),
		FromAddress: client.Address().String(),
		ToAddress:   toAddress.String(),
		Amount:      request.Amount,
		FromChain:   "PLT",
		ToChain:     "ETH",
	}
	global.GVA_DB.Save(crossChainMode)
	if err != nil {
		log.Error("fail to send cross chains transaction,", err)
	}
	log.Info(tx.String())
	response.OkWithData(tx.String(), c)
}

type NftCrossChainRequest struct {
	AssetAddress string
	FromAddress  string
	ToAddress    string
	TokenId      int
	ProxyAddress string
}

// @Tags cross-chain
// @Summary send nft to ETH
// @accept application/json
// @Produce application/json
// @Param data body NftCrossChainRequest true "AssetAddress,FromAddress,ToAddress,TokenId,ProxyAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/sendNFTToEth [post]
func SendNFTToEth(c *gin.Context) {
	request := &NftCrossChainRequest{}
	_ = c.ShouldBind(&request)
	client := CreateNewClient(request.FromAddress)
	//asset, from, proxy, token, to, sideChainID
	assetAddr := common.HexToAddress(request.AssetAddress)
	fromAddr := common.HexToAddress(request.FromAddress)
	proxyAddr := common.HexToAddress(request.ProxyAddress)
	toAddr := common.HexToAddress(request.ToAddress)
	tx, _ := client.NFTSafeTransferFrom(assetAddr, fromAddr, proxyAddr, big.NewInt(int64(request.TokenId)), toAddr, 2)
	response.OkWithData(tx, c)
}



//How To Use NFT:
//1. Deploy NFT Contract
//2. Mint NFT to Address
//3. Send Crosschain Request
//4. Get all NFT Mapping information in LockProxy Contract
//5. Event Track with Event Bus mode
//send nft cross chain
func DemoPLTCrossChainFunction() {
	client := CreateNewClient("0xDe6baAE1014755a0C3F3c8Fb929ca6937f695669")
	currentBalance, _ := client.BalanceOf(client.Address(), "latest")
	log.Info(currentBalance)
	toAddress := common.HexToAddress("0x344cfc3b8635f72f14200aaf2168d9f75df86fd3") //to address
	amount, _ := decimal.NewFromString("0.0000000000000001")
	pltDecimal, _ := decimal.NewFromString("1000000000000000000")
	tx, err := client.Lock(2, toAddress, amount.Mul(pltDecimal).BigInt())
	if err != nil {
		log.Error("fail to send cross chains transaction,", err)
	}
	log.Info(tx.String())
}


//func GetNFTAssetMap(){
//	client:=CreateNewClient("")
//}