package model

import "chains-gotest-backend/global"

type PaletteNode struct {
	global.BaseModel
	NodeAddress    string
	StakeAmount  string
	CurrentBlockNum string
	NodeType int
	SeedNode int
	NodeHost string
	NodeRPCPort string
	NodeP2PPort string
	NodePrivateKey string
	NodeDir string
	Active int
}
