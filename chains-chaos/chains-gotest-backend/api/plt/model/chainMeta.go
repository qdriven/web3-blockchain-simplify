package model

import "chains-gotest-backend/global"

type ChainMetaData struct {
	global.BaseModel
	ChainName string
	ChainId   string
	ChainType string //mainnet/testnet
	RPCUrl    string
}

type EvmTransactionRecord struct {
	global.BaseModel
	FromChainId string
	ToChainId   string
	AssetHash   string //mainnet/testnet
	Amount      string
	ToAddress   string
}

type CoinMetaData struct {
	global.BaseModel
	ChainName string
	ChainId   string
	ChainType string //mainnet/testnet
	CoinName  string
	Symbol    string
	Decimal   string
	Hash      string
	Address   string
}

type CrossChainLockProxy struct {
	global.BaseModel
	ChainId          string
	ChainName        string
	ChainType        string
	Active           bool
	LockProxyAddress string
}

type CrossChainAssetMapping struct {
	global.BaseModel
	LockProxyAddress string
	FromAssetHash    string
	ToAssetHash      string
	ToChainId        uint
	Active           bool
}

type CrossChainProxyHashMap struct {
	global.BaseModel
	LockProxyAddress string
	ToProxyHash      string
	ToChainId        uint
}
