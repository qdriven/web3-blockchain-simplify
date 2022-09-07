package model

import (
	"chains-gotest-backend/global"
)

type NftContract struct {
	global.BaseModel
	TransactionHash         string
	NFTContractAddress string
	Status				uint64
	Name               string
	Symbol             string
	Creator            string
}
