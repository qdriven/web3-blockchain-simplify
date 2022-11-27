package client

import (
	"go-flashrunner/compound/models"
	"math/big"
	"sort"

	"github.com/shopspring/decimal"
)

var (
	LIQ_CLOSE_FACTOR, _  = decimal.NewFromString("0.48")
	Collateral_Factor, _ = decimal.NewFromString("0.6")
	Cmp_Api_Url          = "https://api.compound.finance/api/v2"
	LIQ_ASSET_LIST       = []string{"CETH", "CDAI", "CUSDC", "cETH", "cDAI", "cUSDC"}
	unliqAccounts        = make(map[string]string)
	BASE_EXP             = big.NewInt(10)
)

/**
1. Getting Market Data To Cache
2. Getting Unhealthy Data
3. Calculate Liquidation Information
2. 获取Borrow Tokens信息
3. 获取Supply Tokens 信息
*/

type CompoundMarketData struct {
}

type LiquidateInfo struct {
	BorrowAddress          string
	RepayAmount            string
	RepayAsset             string
	CollateralAssetAddress string
	CollateralAsset        string
	ReturnAmount           string
	RepayAssetAddress      string
}
type Liquidator struct {
	client           *Client
	cTokenMap        map[string]models.CTokenInfo
	cTokenDecimalMap map[string]int64
}

func NewLiquidator() *Liquidator {
	client := &Liquidator{
		client:    NewClient(Cmp_Api_Url),
		cTokenMap: make(map[string]models.CTokenInfo),
	}
	cTokenDecimalMap := make(map[string]int64)
	cTokenDecimalMap["cUSDC"] = 6
	cTokenDecimalMap["cUSDT"] = 6
	cTokenDecimalMap["cDAI"] = 18
	cTokenDecimalMap["cETH"] = 18
	client.LoadMarketData()
	return client
}

type LiquidationCalculator interface {
	LoadMarketData()
	CheckUnLiquidateAccount(account *models.Account) bool
	convertAccountInfo(account *models.Account) *AccountForLiq
	CalculateLiquidationOpts()
	GetLiquidationInfo(account *models.Account) *LiquidateInfo
	AddUnLiquidateAddress()
}

func (l *Liquidator) LoadMarketData() {
	cTokens, _ := l.client.GetCTokens()
	for _, token := range cTokens.CToken {
		l.cTokenMap[token.Symbol] = token
	}
	//TODO: save to database?
}

func (l *Liquidator) getTokenDecimal(tokenName string) int64 {
	if l.cTokenDecimalMap[tokenName] != 0 {
		return l.cTokenDecimalMap[tokenName]
	}
	return 18
}

type AccountForLiq struct {
	SupplyTokens    []models.Token
	BorrowTokens    []models.Token
	BorrowerAddress string
}

func (l *Liquidator) GetLiquidationInfo(account *models.Account) *LiquidateInfo {
	if l.checkUnLiquidateAccount(account) {
		return nil
	}
	accountForLiq := l.convertAccountInfo(account)
	return l.calculateLiquidationOpts(accountForLiq)
}

func (l *Liquidator) convertAccountInfo(account *models.Account) *AccountForLiq {
	accountForLiq := &AccountForLiq{}
	accountForLiq.BorrowerAddress = account.Address
	for _, token := range account.Tokens {
		if token.SupplyBalanceUnderlying.Value != "0.0" {
			accountForLiq.SupplyTokens = append(accountForLiq.SupplyTokens, token)
		}
		if token.BorrowBalanceUnderlying.Value != "0.0" {
			accountForLiq.BorrowTokens = append(accountForLiq.BorrowTokens, token)
		}
	}
	return accountForLiq
}

//liq:
//1. borrow asset:=>

func (l *Liquidator) calculateLiquidationOpts(account *AccountForLiq) *LiquidateInfo {
	//sort the largest eth value
	l.sortToken(account.BorrowTokens, false)
	l.sortToken(account.SupplyTokens, true)

	for i := 0; i < len(account.BorrowTokens); i++ {
		if contains(LIQ_ASSET_LIST, account.BorrowTokens[i].Symbol) {
			borrowToken := account.BorrowTokens[i]
			//repay amount
			for i := 0; i < len(account.SupplyTokens); i++ {
				if contains(LIQ_ASSET_LIST, account.SupplyTokens[i].Symbol) {
					supplyToken := account.SupplyTokens[i]
					if borrowToken.Symbol == supplyToken.Symbol && borrowToken.Symbol == "cUSDC" {
						continue
					}
					liqInfo := &LiquidateInfo{}
					liqInfo.BorrowAddress = account.BorrowerAddress
					liqInfo.CollateralAsset = supplyToken.Symbol
					liqInfo.CollateralAssetAddress = supplyToken.Address
					liqInfo.RepayAmount = l.calculateRepayAmount(borrowToken, supplyToken).String()
					liqInfo.RepayAsset = borrowToken.Symbol
					liqInfo.RepayAssetAddress = borrowToken.Address
					return liqInfo
				}
			}
		}

	}

	return nil
}

func (l *Liquidator) calculateRepayAmount(borrowToken, supplyToken models.Token) decimal.Decimal {
	borrowTokenUnderlyingPrice := l.cTokenMap[borrowToken.Symbol].UnderlyingPrice.Value
	availableForLiqInEth := mulStringValue(borrowTokenUnderlyingPrice, borrowToken.BorrowBalanceUnderlying.Value).Mul(LIQ_CLOSE_FACTOR)
	supplyTokenUnderlyingPrice := l.cTokenMap[supplyToken.Symbol].UnderlyingPrice.Value
	collateralAmountInEth := mulStringValue(supplyTokenUnderlyingPrice, supplyToken.SupplyBalanceUnderlying.Value).Mul(Collateral_Factor)
	price, _ := decimal.NewFromString(borrowTokenUnderlyingPrice)

	tokenDecimal := l.getTokenDecimal(borrowToken.Symbol)
	expNum := decimal.New(1, int32(tokenDecimal))

	if availableForLiqInEth.GreaterThan(collateralAmountInEth) {
		repayAmount := collateralAmountInEth.Div(price).Mul(expNum)
		return repayAmount
	} else {
		repayAmount := availableForLiqInEth.Div(price).Mul(expNum)
		return repayAmount
	}
}

func (l *Liquidator) sortToken(tokens []models.Token, isSupply bool) {
	sort.Slice(tokens, func(i, j int) bool {
		//sort by eth value because eth value is 1, in market underlying price
		iUnderlyPrice := l.cTokenMap[tokens[i].Symbol].UnderlyingPrice.Value
		jUnderlyPrice := l.cTokenMap[tokens[j].Symbol].UnderlyingPrice.Value
		if isSupply {
			iAssetInEth := mulStringValue(iUnderlyPrice, tokens[i].SupplyBalanceUnderlying.Value)
			jAssetInEth := mulStringValue(jUnderlyPrice, tokens[j].SupplyBalanceUnderlying.Value)
			return iAssetInEth.GreaterThan(jAssetInEth)
		} else {
			iAssetInEth := mulStringValue(iUnderlyPrice, tokens[i].BorrowBalanceUnderlying.Value)
			jAssetInEth := mulStringValue(jUnderlyPrice, tokens[j].BorrowBalanceUnderlying.Value)
			return iAssetInEth.GreaterThan(jAssetInEth)
		}
	})
}
func mulStringValue(base, multiper string) decimal.Decimal {
	value, _ := decimal.NewFromString(base)
	mul, _ := decimal.NewFromString(multiper)
	return value.Mul(mul)
}

func contains(s []string, searchterm string) bool {
	for _, item := range s {
		if item == searchterm {
			return true
		}
	}
	return false
}

func (l *Liquidator) checkUnLiquidateAccount(account *models.Account) bool {
	if unliqAccounts[account.Address] != "" {
		return true
	}
	return false
}
