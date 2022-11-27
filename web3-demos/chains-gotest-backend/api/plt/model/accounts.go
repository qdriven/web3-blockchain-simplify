package model

import (
	"chains-gotest-backend/global"
)

type TestAccount struct {
	global.BaseModel
	Address        string
	AccIndex       int
	DerivationPath string
	PrivateKey     string
	IsAdmin        bool
}

type TestAccountBalance struct {
	global.BaseModel
	Address         string
	Balance         uint64
	CurrentBlockNum int
	Symbol string
	ChainId int
}
