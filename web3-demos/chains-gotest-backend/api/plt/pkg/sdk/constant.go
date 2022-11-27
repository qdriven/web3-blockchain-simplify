package sdk

import "math/big"

const (
	NetworkID = 10
	GasPrice  = 0
	GasMin    = 21000
	GasNormal = 100000
	Gas1GW    = 1000000000
	Gas10GW   = 10 * Gas1GW
	Gas100GW  = 10 * Gas10GW
	Gas1000GW = 10 * Gas100GW
)

var (
	OneEth, _ = new(big.Int).SetString("1000000000000000000", 10)
	OnePLT, _ = new(big.Int).SetString("1000000000000000000", 10)
)
