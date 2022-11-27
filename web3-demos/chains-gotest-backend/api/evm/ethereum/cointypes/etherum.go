package cointypes

import "github.com/ethereum/go-ethereum/common"

type rateModel int64

const (
	// StableRate https://medium.com/aave/aave-borrowing-rates-upgraded-f6c8b27973a7
	StableRate rateModel = 1
	// VariableRate https://medium.com/aave/aave-borrowing-rates-upgraded-f6c8b27973a7
	VariableRate rateModel = 2
)

type coinType int

const (
	// ETH weiwu
	ETH coinType = iota
	// BAT weiwu
	BAT coinType = iota
	// COMP weiwu
	COMP coinType = iota
	// DAI weiwu
	DAI coinType = iota
	// REP weiwu
	REP coinType = iota
	// SAI weiwu
	SAI coinType = iota
	// UNI weiwu
	UNI coinType = iota
	// USDC weiwu
	USDC coinType = iota
	// USDT weiwu
	USDT coinType = iota
	// WBTC weiwu
	WBTC coinType = iota
	// ZRX weiwu
	ZRX coinType = iota
	// BUSD weiwu
	BUSD coinType = iota

	// cToken
	cETH = iota

	cDAI = iota

	cUSDC = iota

	yWETH = iota
)

const (
	// uniswapAddr is UniswapV2Router, see here: https://uniswap.org/docs/v2/smart-contracts/router02/#address
	uniswapAddr             string = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	yRegistryAddr           string = "0x3eE41C098f9666ed2eA246f4D2558010e59d63A0"
	yETHVaultAddr           string = "0xe1237aA7f535b0CC33Fd973D66cBf830354D16c7"
	aaveLendingPoolAddr     string = "0x398eC7346DcD622eDc5ae82352F02bE94C62d119"
	aaveLendingPoolCoreAddr string = "0x3dfd23A6c5E8BbcFc9581d2E864a68feb6a076d3"
	// Furucombo related addresses

	// FurucomboAddr is the address of the Furucombo proxy
	FurucomboAddr string = "0x57805e5a227937bac2b0fdacaa30413ddac6b8e1"
	hCEtherAddr   string = "0x9A1049f7f87Dbb0468C745d9B3952e23d5d6CE5e"
	hErcInAddr    string = "0x914490a362f4507058403a99e28bdf685c5c767f"
	hCTokenAddr   string = "0x8973D623d883c5641Dd3906625Aac31cdC8790c5"
	hMakerDaoAddr string = "0x294fbca49c8a855e04d7d82b28256b086d39afea"
	hUniswapAddr  string = "0x58a21cfcee675d65d577b251668f7dc46ea9c3a0"
	hCurveAddr    string = "0xa36dfb057010c419c5917f3d68b4520db3671cdb"
	hYearnAddr    string = "0xC50C8F34c9955217a6b3e385a069184DCE17fD2A"
	hAaveAddr     string = "0xf579b009748a62b1978639d6b54259f8dc915229"
	hOneInch      string = "0x783f5c56e3c8b23d90e4a271d7acbe914bfcd319"
	hFunds        string = "0xf9b03e9ea64b2311b0221b2854edd6df97669c09"
	hKyberAddr    string = "0xe2a3431508cd8e72d53a0e4b57c24af2899322a0"
	// TODO: The following is not on mainnet yet
	hSushiswapAddr string = "0xB6F469a8930dd5111c0EA76571c7E86298A171f7"

	// Curve pool addresses
	cCompound string = "0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56"
	cUsdt     string = "0x52EA46506B9CC5Ef470C5bf89f17Dc28bB35D85C"
	cY        string = "0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51"
	cBusd     string = "0x79a8C46DeA5aDa233ABaFFD40F3A0A2B1e5A4F27"
	cSusd     string = "0xA5407eAE9Ba41422680e2e00537571bcC53efBfD"
	cRen      string = "0x93054188d876f558f4a66B2EF1d97d16eDf0895B"
	cSbtc     string = "0x7fC77b5c7614E1533320Ea6DDc2Eb61fa00A9714"
	cHbtc     string = "0x4ca9b3063ec5866a4b82e437059d2c43d1be596f"
	c3Pool    string = "0xbebc44782c7db0a1a60cb6fe97d0b483032ff1c7"
	cGusd     string = "0x4f062658eaaf2c1ccf8c8e36d6824cdf41167956"
	cHusd     string = "0x3eF6A01A0f81D6046290f3e2A8c5b843e738E604"
	cUsdk     string = "0x3e01dd8a5e1fb3481f0f589056b428fc308af0fb"
	cUsdn     string = "0x0f9cb53Ebe405d49A0bbdBD291A65Ff571bC83e1"

	// Curve token addresses
	compCrv      string = "0x845838DF265Dcd2c412A1Dc9e959c7d08537f8a2"
	usdtCrv      string = "0x9fC689CCaDa600B6DF723D9E47D84d76664a1F23"
	yCrv         string = "0xdF5e0e81Dff6FAF3A7e52BA697820c5e32D806A8"
	busdCrv      string = "0x3B3Ac5386837Dc563660FB6a0937DFAa5924333B"
	susdCrv      string = "0xC25a3A3b969415c80451098fa907EC722572917F"
	renCrv       string = "0x49849C98ae39Fff122806C06791Fa73784FB3675"
	sbtcCrv      string = "0x075b1bb99792c9E1041bA13afEf80C91a1e70fB3"
	hbtcCrv      string = "0xb19059ebb43466C323583928285a49f558E572Fd"
	threePoolCrv string = "0x6c3F90f043a72FA612cbac8115EE7e52BDe6E490"
	gusdCrv      string = "0xD2967f45c4f384DEEa880F807Be904762a3DeA07"
	husdCrv      string = "0x5B5CFE992AdAC0C9D48E05854B2d91C73a003858"
	usdkCrv      string = "0x97E2768e8E73511cA874545DC5Ff8067eB19B787"
	usdnCrv      string = "0x4f3E8F405CF5aFC05D68142F3783bDfE13811522"
)

// CoinToAddressMap returns a mapping from coin to address
var CoinToAddressMap = map[coinType]common.Address{
	ETH:   common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
	BAT:   common.HexToAddress("0x0d8775f648430679a709e98d2b0cb6250d2887ef"),
	COMP:  common.HexToAddress("0xc00e94cb662c3520282e6f5717214004a7f26888"),
	DAI:   common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
	USDC:  common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
	USDT:  common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
	cETH:  common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"),
	cDAI:  common.HexToAddress("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"),
	cUSDC: common.HexToAddress("0x39aa39c021dfbae8fac545936693ac917d5e7563"),
	BUSD:  common.HexToAddress("0x4Fabb145d64652a948d72533023f6E7A623C7C53"),
	yWETH: common.HexToAddress("0xe1237aA7f535b0CC33Fd973D66cBf830354D16c7"),
}

// CoinToCompoundMap returns a mapping from coin to compound address
var CoinToCompoundMap = map[coinType]common.Address{
	ETH:  common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"),
	DAI:  common.HexToAddress("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"),
	USDC: common.HexToAddress("0x39aa39c021dfbae8fac545936693ac917d5e7563"),
}
