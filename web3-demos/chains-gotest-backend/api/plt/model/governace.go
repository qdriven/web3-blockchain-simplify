package model

import "chains-gotest-backend/global"

type ValidatorHistory struct {
	global.BaseModel
	Validator          string
	StackAccount       string
	CurrentBlockHeight int64
	Revoke string
	TxHash string
}

