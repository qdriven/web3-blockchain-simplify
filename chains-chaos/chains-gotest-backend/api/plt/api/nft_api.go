package api

import (
	"chains-gotest-backend/api/plt/core"
	"chains-gotest-backend/api/plt/model"
	"chains-gotest-backend/api/plt/pkg/log"
	"chains-gotest-backend/global"
	"chains-gotest-backend/model/response"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"strconv"
)

type DeployRequest struct {
	Name   string
	Symbol string
	Owner  string
}

/**
部署Nft合约
*/
// @Tags NFTAPI
// @Summary Deploy NFT Contract
// @accept application/json
// @Produce application/json
// @Param data body DeployRequest true "Name,Symbol,Owner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftDeploy [post]
func Deploy(ctx *gin.Context) {
	request := &DeployRequest{}
	_ = ctx.ShouldBind(&request)
	//Get Owner Address, if owner is empty use admin
	//TODO: validatornftTokenuri
	SetTxAccount(request.Owner)
	request.Owner = core.Client.Address().String()
	tx, nftAddress, _ := core.Client.NFTDeploy(request.Name, request.Symbol)
	txReceipt, _ := core.Client.TransactionReceipt(context.Background(), tx)
	//todo: write to database
	NftContract := &model.NftContract{
		NFTContractAddress: nftAddress.Hex(),
		TransactionHash:    tx.String(),
		Name:               request.Name,
		Symbol:             request.Symbol,
		Creator:            request.Owner,
		Status:             txReceipt.Status,
	}
	global.GVA_DB.Save(NftContract)
	response.OkWithData(NftContract, ctx)
}

type BatchDeployRequest struct {
	Num   string `json:"num" form:"num"`
	Owner string `json:"owner" form:"owner"`
	Name string `json:"name" form:"name"`
	Symbol string `json:"symbol" form:"symbol"`
}

func setupDefaultLink() (string, [12]string) {
	baseUrl := "http://r6d.cn/"
	links := [12]string{
		"bw3hx",
		"bw3h2",
		"bw3h4",
		"bw3hn",
		"bw3h3",
		"bw3h5",
		"bw3h7",
		"bw3h1",
		"bw3ya",
		"bw3yk",
		"bw3yr",
		"bw3yy",
	}
	return baseUrl, links
}

// @Tags NFTAPI
// @Summary Deploy NFT Contract
// @accept application/json
// @Produce application/json
// @Param data body BatchDeployRequest true "Num,Owner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftBatchDeploy [post]
func BatchDeploy(ctx *gin.Context) {
	request := &BatchDeployRequest{}
	_ = ctx.ShouldBind(&request)
	num, _ := strconv.Atoi(request.Num)
	SetTxAccount(request.Owner)
	baseUrl, links := setupDefaultLink()
	name := request.Name
	symbol := request.Symbol
	resultMap := make(map[string]string)
	commonAddress := common.StringToAddress("0x3FbDaFDB7ff9f56F25567d3f596c8cfCB0B2E2F8")
	for i := 0; i < num; i++ {
		tx, nftAddress, _ := core.Client.NFTDeploy(name+strconv.Itoa(i), symbol+strconv.Itoa(i))
		core.Client.WaitTransaction(tx)
		log.Info("setup NFT BaseURI")
		tx1, err := core.Client.NFTSetBaseUri(nftAddress, baseUrl)
		if err != nil {
			log.Error(err)
		}
		core.Client.WaitTransaction(tx1)
		log.Info("mint NFT BaseURI")
		tx2, err := core.Client.NFTMint(nftAddress, commonAddress, big.NewInt(1), links[i])
		if err != nil {
			log.Error(err)
		}
		core.Client.WaitTransaction(tx2)
		resultMap[tx.String()] = nftAddress.String()
	}
	response.OkWithData(resultMap, ctx)
}



type MintRequest struct {
	NFTAddress   string
	To           string
	Token        int64
	Uri          string
	OwnerAddress string
}

// @Tags NFTAPI
// @Summary Mint NFT Contract
// @accept application/json
// @Produce application/json
// @Param data body MintRequest true "NFTAddress,To,Token,Uri"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftMint [post]
func Mint(ctx *gin.Context) {
	request := &MintRequest{}
	_ = ctx.ShouldBind(&request)
	SetTxAccount(request.OwnerAddress)
	result, _ := core.Client.NFTMint(common.HexToAddress(request.NFTAddress),
		common.HexToAddress(request.To),
		big.NewInt(request.Token), request.Uri)
	//填写逻辑
	response.OkWithData(result.String(), ctx)
}

type BalanceofRequest struct {
	Owner      string
	NFTAddress string
}

// @Tags NFTAPI
// @Summary NFT Contract Balance Of a given NTFAddress
// @accept application/json
// @Produce application/json
// @Param data body BalanceofRequest true "Owner,NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftBalanceof [post]
func Balanceof(ctx *gin.Context) {
	request := &BalanceofRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.NFTBalanceOf(common.HexToAddress(request.NFTAddress), common.HexToAddress(request.Owner))
	//填写逻辑
	response.OkWithData(result.String(), ctx)
}

type SafetransferfromRequest struct {
	From       string
	To         string
	Tokenid    int64
	NFTAddress string
	Data       string
}

//func Safetransferfrom(ctx *gin.Context) {
//	request := &SafetransferfromRequest{}
//	_ = ctx.ShouldBind(&request)
//	SetTxAccount(request.From)
//	result, err := core.Client.NFTSafeTransferFrom(common.HexToAddress(request.NFTAddress),
//		common.HexToAddress(request.From), common.HexToAddress(request.To), big.NewInt(request.Tokenid))
//	//填写逻辑
//	if err != nil {
//		log.Info("error is ", err)
//	}
//	response.OkWithData(result.String(), ctx)
//}

type TransferfromRequest struct {
	From       string
	To         string
	Tokenid    int64
	NFTAddress string
}

// @Tags NFTAPI
// @Summary Transfer From NFT Contract to a given NTFAddress
// @accept application/json
// @Produce application/json
// @Param data body TransferfromRequest true "From,To,Tokenid,NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftTransferfrom [post]
func Transferfrom(ctx *gin.Context) {
	request := &TransferfromRequest{}
	_ = ctx.ShouldBind(&request)
	SetTxAccount(request.From)
	result, _ := core.Client.NFTTransferFrom(common.HexToAddress(request.NFTAddress),
		common.HexToAddress(request.From), common.HexToAddress(request.To), big.NewInt(request.Tokenid))
	//填写逻辑
	response.OkWithData(result.String(), ctx)
}

type ApproveRequest struct {
	ToApproveAddress string
	Tokenid          int64
	NFTAddress       string
	OwnerAddress     string
}

// @Tags NFTAPI
// @Summary Approve From NFT Contract to a given NTFAddress
// @accept application/json
// @Produce application/json
// @Param data body ApproveRequest true "ToApproveAddress,Tokenid,NFTAddress,OwnerAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftApprove [post]
func Approve(ctx *gin.Context) {
	request := &ApproveRequest{}
	_ = ctx.ShouldBind(&request)
	SetTxAccount(request.OwnerAddress)
	result, _ := core.Client.NFTApprove(common.HexToAddress(request.NFTAddress),
		common.HexToAddress(request.ToApproveAddress), big.NewInt(request.Tokenid))
	response.OkWithData(result.String(), ctx)
}

type SetapprovalforallRequest struct {
	Operator     string
	Approved     bool
	NFTAddress   string
	OwnerAddress string
}

// @Tags NFTAPI
// @Summary set approval for all From NFT Contract to a given NTFAddress
// @accept application/json
// @Produce application/json
// @Param data body SetapprovalforallRequest true "Operator,Approved,NFTAddress,OwnerAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftSetapprovalforall [post]
func Setapprovalforall(ctx *gin.Context) {
	request := &SetapprovalforallRequest{}
	_ = ctx.ShouldBind(&request)
	SetTxAccount(request.OwnerAddress)
	result, _ := core.Client.NFTSetApprovalForAll(common.HexToAddress(request.NFTAddress), common.HexToAddress(request.Operator), request.Approved)
	//填写逻辑
	response.OkWithData(result.String(), ctx)
}

type OwnerofRequest struct {
	Tokenid    int64
	NFTAddress string
}

// @Tags NFTAPI
// @Summary owner of given nft address
// @accept application/json
// @Produce application/json
// @Param data body OwnerofRequest true "Tokenid,NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftOwnerof [post]
func Ownerof(ctx *gin.Context) {
	request := &OwnerofRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.NFTOwnerOf(common.HexToAddress(request.NFTAddress), big.NewInt(request.Tokenid))
	//填写逻辑
	response.OkWithData(result.String(), ctx)
}

type GetapprovedRequest struct {
	Tokenid    int64
	NFTAddress string
}

// @Tags NFTAPI
// @Summary approved address of given nft address
// @accept application/json
// @Produce application/json
// @Param data body GetapprovedRequest true "Tokenid,NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftGetApproved [post]
func Getapproved(ctx *gin.Context) {
	request := &GetapprovedRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.NFTGetApproved(common.HexToAddress(request.NFTAddress), big.NewInt(request.Tokenid))
	//填写逻辑
	response.OkWithData(result.String(), ctx)
}

type IsapprovedforallRequest struct {
	Owner      string
	Operator   string
	NFTAddress string
}

// @Tags NFTAPI
// @Summary  is approved for all
// @accept application/json
// @Produce application/json
// @Param data body IsapprovedforallRequest true "Owner,Operator,NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftIsapprovedforall [post]
func Isapprovedforall(ctx *gin.Context) {
	request := &IsapprovedforallRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.NFTIsApprovedForAll(common.HexToAddress(request.NFTAddress),
		common.HexToAddress(request.Owner), common.HexToAddress(request.Operator))
	//填写逻辑
	response.OkWithData(result, ctx)
}

type TotalsupplyRequest struct {
	NFTAddress string
}

// @Tags NFTAPI
// @Summary total supply
// @accept application/json
// @Produce application/json
// @Param data body TotalsupplyRequest true "NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftTotalSupply [post]
func TotalSupply(ctx *gin.Context) {
	request := &TotalsupplyRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.NFTTotalSupply(common.HexToAddress(request.NFTAddress))
	//填写逻辑
	response.OkWithData(result, ctx)
}

type TokenbyindexRequest struct {
	Index      int    `json:"index"`
	NFTAddress string `json:"nftAddress"`
}

// @Tags NFTAPI
// @Summary Tokenbyindex
// @accept application/json
// @Produce application/json
// @Param data body TokenbyindexRequest true "index,NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftTokenbyindex [post]
func Tokenbyindex(ctx *gin.Context) {
	request := &TokenbyindexRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.NFTTokenByIndex(common.HexToAddress(request.NFTAddress), big.NewInt(int64(request.Index)))
	response.OkWithData(result, ctx)
}

type TokenofownerbyindexRequest struct {
	Owner      string `json:"owner"`
	Index      int    `json:"index"`
	NFTAddress string `json:"nftAddress"`
}

// @Tags NFTAPI
// @Summary Tokenofownerbyindex
// @accept application/json
// @Produce application/json
// @Param data body TokenofownerbyindexRequest true "Owner,index,NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftTokenofownerbyindex [post]
func Tokenofownerbyindex(ctx *gin.Context) {
	request := &TokenofownerbyindexRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.NFTTokenOfOwnerByIndex(common.HexToAddress(request.NFTAddress), common.HexToAddress(request.Owner), big.NewInt(int64(request.Index)))
	//填写逻辑
	response.OkWithData(result, ctx)
}

type BurnRequest struct {
	Token        int64
	NFTAddress   string
	OwnerAddress string
}

// @Tags NFTAPI
// @Summary Burn NFT token
// @accept application/json
// @Produce application/json
// @Param data body BurnRequest true "Token,NFTAddress,OwnerAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftBurn [post]
func Burn(ctx *gin.Context) {
	request := &BurnRequest{}
	_ = ctx.ShouldBind(&request)
	SetTxAccount(request.OwnerAddress)
	result, _ := core.Client.NFTBurn(common.HexToAddress(request.NFTAddress), big.NewInt(request.Token))
	//填写逻辑
	response.OkWithData(result.String(), ctx)
}

type NameRequest struct {
	NFTAddress string
}

// @Tags NFTAPI
// @Summary Name NFT token
// @accept application/json
// @Produce application/json
// @Param data body NameRequest true "NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftName [post]
func Name(ctx *gin.Context) {
	request := &NameRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.NFTName(common.HexToAddress(request.NFTAddress))
	response.OkWithData(result, ctx)
}

type SymbolRequest struct {
	NFTAddress string
}

// @Tags NFTAPI
// @Summary Symbol NFT token
// @accept application/json
// @Produce application/json
// @Param data body SymbolRequest true "NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftSymbol [post]
func Symbol(ctx *gin.Context) {
	request := &SymbolRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.NFTSymbol(common.HexToAddress(request.NFTAddress))
	response.OkWithData(result, ctx)
}

type OwnerRequest struct {
	NFTAddress string
}

// @Tags NFTAPI
// @Summary owner of  NFT token
// @accept application/json
// @Produce application/json
// @Param data body OwnerRequest true "NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftOwner [post]
func Owner(ctx *gin.Context) {
	request := &OwnerRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.NFTOwner(common.HexToAddress(request.NFTAddress))
	response.OkWithData(result, ctx)
}

type TokenuriRequest struct {
	Token      int64
	NFTAddress string
}

// @Tags NFTAPI
// @Summary token uri of  NFT token
// @accept application/json
// @Produce application/json
// @Param data body TokenuriRequest true "Token,NFTAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftTokenuri [post]
func Tokenuri(ctx *gin.Context) {
	request := &TokenuriRequest{}
	_ = ctx.ShouldBind(&request)
	result, _ := core.Client.NFTTokenUri(common.HexToAddress(request.NFTAddress),
		big.NewInt(request.Token))
	//填写逻辑
	log.Info("current result is {}", result)
	response.OkWithData(result, ctx)
}

type SetTokenBaseUriRequest struct {
	NFTAddress              string
	BaseUri                 string
	NFTContractOwnerAddress string
}

// @Tags NFTAPI
// @Summary set token base uri of  NFT token
// @accept application/json
// @Produce application/json
// @Param data body SetTokenBaseUriRequest true "BaseUri,NFTAddress,NFTContractOwnerAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftSetBaseUri [post]
func SetTokenBaseUri(ctx *gin.Context) {
	request := &SetTokenBaseUriRequest{}
	_ = ctx.ShouldBind(&request)
	client := CreateNewClient(request.NFTContractOwnerAddress)
	result, _ := client.NFTSetBaseUri(common.HexToAddress(request.NFTAddress), request.BaseUri)
	//填写逻辑
	response.OkWithData(result.String(), ctx)
}

type TokenBaseUriRequest struct {
	NFTAddress              string
	NFTContractOwnerAddress string
}

// @Tags NFTAPI
// @Summary set token base uri of  NFT token
// @accept application/json
// @Produce application/json
// @Param data body TokenBaseUriRequest true "NFTAddress,NFTContractOwnerAddress"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /pallet/nftGetTokenBaseUri [post]
func GetTokenBaseUri(ctx *gin.Context) {
	request := &SetTokenBaseUriRequest{}
	_ = ctx.ShouldBind(&request)
	client := CreateNewClient(request.NFTContractOwnerAddress)
	result, _ := client.NFTBaseUri(common.HexToAddress(request.NFTAddress))
	//填写逻辑
	response.OkWithData(result, ctx)
}
