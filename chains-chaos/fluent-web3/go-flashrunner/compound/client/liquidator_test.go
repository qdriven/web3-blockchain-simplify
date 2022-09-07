package client

import (
	"fmt"
	"testing"
)

var liquidator = NewLiquidator()


func TestLoadMarketData(t *testing.T) {
	fmt.Println(liquidator.cTokenMap)
}

func TestGetLiquidationOps(t *testing.T) {
	cmpApiClient := NewClient(Cmp_Api_Url)

	accounts,err:=cmpApiClient.GetUnhealthyAccount("1.00","1")

	if err!=nil{
		fmt.Println(err)
	}

	for _, account := range accounts {
		liqInfo :=liquidator.GetLiquidationInfo(&account)
		fmt.Println(liqInfo)
	}
}