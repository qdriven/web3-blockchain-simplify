package model

import "chains-gotest-backend/global"

type CrossChainTx struct {
	global.BaseModel
	FromChain   string
	FromAddress string
	ToAddress   string
	Amount      string
	TxHash      string
	ToChain     string
	Status	string
}
