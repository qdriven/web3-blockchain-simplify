package clients

import (
	"chains-gotest-backend/api/evm/evm/binding/plttoken/nftlp"
	"chains-gotest-backend/api/evm/evm/binding/plttoken/nftmapping_abi"
	"chains-gotest-backend/api/evm/log"
	"chains-gotest-backend/api/evm/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	polycm "github.com/polynetwork/poly/common"
)

type Erc721CrossChainRequestContext struct {
	LockProxyAddress string
	FromAddressHex   string
	ToAddressHex     string
	TokenId          uint64
	TokenURI         string
	FromChainId      uint64
	ToChainId        uint64
	FromAssetHash    string
	ToAssetHash      string
}

func (c *EvmClient) DeployCrossChainERC721(lockProxyHex, name, symbol string) (*types.Transaction, *nftmapping_abi.CrossChainNFTMapping, error) {
	auth := c.MakeAuth()
	lockProxyAddress := common.HexToAddress(lockProxyHex)
	contractAddress, tx, contract, err := nftmapping_abi.DeployCrossChainNFTMapping(auth, c.EClient, lockProxyAddress, name, symbol)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	c.WaitTransactionConfirm(tx.Hash())

	log.Info("deploy cross chain 721 contract address is ", contractAddress.String())
	totalSupply, err := contract.TotalSupply(nil)
	log.Info("deploy erc20 contract total Supply is ", totalSupply)
	return tx, contract, nil
}

func (c *EvmClient) BindNFTAsset(
	lockProxyAddr,
	fromAssetHash,
	toAssetHash string,
	targetSideChainId uint64) (common.Hash, error) {
	proxy, err := c.MakeNFTLockProxy(lockProxyAddr)
	if err != nil {
		return utils.EmptyHash, err
	}

	auth := c.MakeAuth()
	fromAsset := common.HexToAddress(fromAssetHash)
	toAsset := common.HexToAddress(toAssetHash)
	tx, err := proxy.BindAssetHash(auth, fromAsset, targetSideChainId, toAsset[:])
	if err != nil {
		return utils.EmptyHash, err
	}
	c.WaitTransactionConfirm(tx.Hash())
	return tx.Hash(), nil
}

func (c *EvmClient) MakeNFTLockProxy(lockProxyAddr string) (*nftlp.NFTLockProxy, error) {
	lockProxy := common.HexToAddress(lockProxyAddr)
	proxy, err := nftlp.NewNFTLockProxy(lockProxy, c.backend)

	if err != nil {
		return nil, err
	}
	return proxy, nil
}

func (c *EvmClient) GetAssetMap(lockProxyAddr string, FromAssetHash string, ToChainId uint64) (string, error) {
	proxy, err := c.MakeNFTLockProxy(lockProxyAddr)
	if err != nil {
		return "", err
	}
	fromAsset := common.HexToAddress(FromAssetHash)
	toAssetHash, err := proxy.AssetHashMap(nil, fromAsset, ToChainId)
	if err != nil {
		return "", err
	}
	return string(toAssetHash), nil
}

func (c *EvmClient) GetBindProxy(lockProxyAddr string, ToChainId uint64) (string, error) {
	proxy, err := c.MakeNFTLockProxy(lockProxyAddr)
	if err != nil {
		return "", err
	}
	toProxyHash, err := proxy.ProxyHashMap(nil, ToChainId)
	if err != nil {
		return "", err
	}
	return string(toProxyHash), nil
}

func (c *EvmClient) SafeTransferFrom(contractAddrHex, fromAddressHex, lockProxyAddressHex, tokenId, toAddressHex string, toChainId uint64) (*types.Transaction, error) {
	contractAddr := common.HexToAddress(contractAddrHex)
	erc721Contract, err := nftmapping_abi.NewCrossChainNFTMapping(contractAddr, c.backend)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	auth := c.MakeAuth()
	fromAddress := common.HexToAddress(fromAddressHex)
	lockProxyAddress := common.HexToAddress(lockProxyAddressHex) //should be lockproxy address
	tokenIdBigInt := utils.ToIntByPrecise(tokenId, 10)
	toAddress := common.HexToAddress(toAddressHex)
	callData := assembleSafeTransferCallData(toAddress, toChainId)
	tx, err := erc721Contract.SafeTransferFrom0(auth, fromAddress, lockProxyAddress, tokenIdBigInt, callData)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	c.WaitTransactionConfirm(tx.Hash())
	return tx, nil
}

func assembleSafeTransferCallData(toAddress common.Address, chainID uint64) []byte {
	sink := polycm.NewZeroCopySink(nil)
	sink.WriteVarBytes(toAddress.Bytes())
	sink.WriteUint64(chainID)
	return sink.Bytes()
}

//1. Getting Asset Mapping
//2. Getting Get Proxy binding
//3. Asset Database
