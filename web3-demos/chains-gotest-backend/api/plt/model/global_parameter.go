package model

import "chains-gotest-backend/global"

type GlobalParameter struct {
	global.BaseModel
	CurrentBlockNo    string
	NftGasPriceInPlt  string
	NftMintPriceInPlt string
	RewardPeriod string
}