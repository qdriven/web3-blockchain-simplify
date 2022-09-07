package clients

import (
	"chains-gotest-backend/api/evm/evm/binding/poly/lock_proxy_abi"
	"chains-gotest-backend/api/evm/log"
	"chains-gotest-backend/api/evm/models"
	"chains-gotest-backend/api/evm/utils"
	"chains-gotest-backend/global"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

)

type CrossChainRequestContext struct {
	LockProxyAddress string
	FromAddressHex   string
	ToAddressHex     string
	Amount           string
	Decimal          uint64
	FromChainId      uint64
	ToChainId        uint64
	ApproveAmount    string
}

func (c *EvmClient) SendApproveAndLockTx(request *CrossChainRequestContext) (*types.Transaction, error) {
	return c.ApproveAndLock(request.LockProxyAddress,
		request.FromAddressHex,
		request.Amount,
		request.ToAddressHex,
		request.ApproveAmount,
		request.Decimal,
		request.ToChainId)
}

func (c *EvmClient) SendLockTx(request *CrossChainRequestContext) (*types.Transaction, error) {
	return c.LockNoWait(request.LockProxyAddress,
		request.FromAddressHex,
		request.Amount,
		request.ToAddressHex,
		request.Decimal,
		request.ToChainId)
}

func (c *EvmClient) ApproveAndLock(lockProxyAddress, fromAssetHex, amount, toAddressHex, approveAmount string, decimal, chainId uint64) (*types.Transaction, error) {

	//PLT Token Address
	fromAssetHash := common.HexToAddress(fromAssetHex)
	contractAddress := common.HexToAddress(lockProxyAddress)
	if approveAmount == "" {
		approveAmount = "10000000000"
	}
	erc20Tx, _ := c.Approve(fromAssetHash, contractAddress, utils.ToIntByPrecise(approveAmount, decimal))
	log.Info("approve transaction hash is:", erc20Tx.Hash().String())
	return c.Lock(lockProxyAddress, fromAssetHex, amount, toAddressHex, decimal, chainId)
}

func (c *EvmClient) Lock(lockProxyAddress, fromAssetHex, amount, toAddressHex string, decimal, chainId uint64) (*types.Transaction, error) {
	contractAddress := common.HexToAddress(lockProxyAddress)
	lockProxy, err := lock_proxy_abi.NewLockProxy(contractAddress, c.EClient)
	if err != nil {
		log.Info(err)
	}
	//跨链
	fromAssetHash := common.HexToAddress(fromAssetHex)
	auth := c.MakeAuth()
	toAddress := common.HexToAddress(toAddressHex) //Test Accounts
	lockTx, err := lockProxy.Lock(auth, fromAssetHash, chainId, toAddress.Bytes(), utils.ToIntByPrecise(amount, decimal))
	if err != nil {
		log.Error("send lock transaction failed,", err)
	}
	c.WaitTransactionConfirm(lockTx.Hash())
	log.Info("transaction hash:", lockTx.Hash().String())
	return lockTx, nil
}

func (c *EvmClient) LockNoWait(lockProxyAddress, fromAssetHex, amount, toAddressHex string, decimal, chainId uint64) (*types.Transaction, error) {
	contractAddress := common.HexToAddress(lockProxyAddress)
	lockProxy, err := lock_proxy_abi.NewLockProxy(contractAddress, c.EClient)
	if err != nil {
		log.Info(err)
	}
	//跨链
	fromAssetHash := common.HexToAddress(fromAssetHex)
	auth := c.MakeAuth()
	toAddress := common.HexToAddress(toAddressHex) //Test Accounts
	lockTx, err := lockProxy.Lock(auth, fromAssetHash, chainId, toAddress.Bytes(), utils.ToIntByPrecise(amount, decimal))
	if err != nil {
		log.Error("send lock transaction failed,", err)
	}
	log.Info("transaction hash:", lockTx.Hash().String())

	crossChainModel := &models.CrossChainTx{
		TxHash:      lockTx.Hash().String(),
		FromAddress: c.Address().String(),
		ToAddress:   toAddress.String(),
		Amount:      amount,
		FromChain:   "ETH",
		ToChain:     "PLT",
	}
	global.GVA_DB.Save(crossChainModel)
	return lockTx, nil
}

//TODO: search for tx status and update status

//1. Getting Asset Mapping
//2. Getting Get Proxy binding
//3. Asset Database