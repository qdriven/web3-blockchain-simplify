package model

import (
	"chains-gotest-backend/global"
)

type RewardEvent struct {
	global.BaseModel
	EventName       string
	ContractAddress string
	Amount          string
	CurrentBlockNum uint64
}
//staking history to track staking events
type StakingEvent struct {
	global.BaseModel
	Contract string
	CurrentBlockHeight uint64
	EventName string
	Owner string
	Validator string
	StakeAccount string
	Revoke bool
	Data string
	TxHash string

}

type NFTEvent struct {
	global.BaseModel
	EventName          string
	ContractAddress    string
	BlockHeight        uint64
	TxHash             string
	From               string
	To                 string
	TokenId            string
	Owner              string
	Spender            string
	Approved           string
	DeployedNFTAddress string
	DeployNFTData string
}

type PLTTransferEvent struct {
	global.BaseModel
	EventName   string
	Contract    string
	BlockHeight uint64
	TxHash      string
	From        string
	To          string
	Value       string
}

type PLTApprovalEvent struct {
	global.BaseModel
	EventName   string
	Contract    string
	BlockHeight uint64
	TxHash      string
	Owner       string
	Spender     string
	Value       string
}
