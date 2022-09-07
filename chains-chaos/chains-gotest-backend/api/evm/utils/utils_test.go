/*
 * Copyright (C) 2020 The onchain Authors
 * This file is part of The onchain library.
 *
 *All rights reserved.
 */
package utils

import (
	"chains-gotest-backend/api/evm/log"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"math/big"
	"strconv"
	"testing"
)

func TestStringAndInt(t *testing.T) {
	testStr := []string{
		"1111",
		"111.2333",
		"11.0000000000000000000000000000000000000",
		"0238462326847234.61637644",
		"39119032.1111111111111111111111111111199999999999999999999999",
		"132716.0000002734824800000000",
		"-1111",
		"-111.2333",
		"-11.0000000000000000000000000000000000000",
		"-0238462326847234.61637644",
		"-39119032.1111111111111111111111111111199999999999999999999999",
		"-132716.0000002734824800000000",
	}
	testDecimals := []uint64{0, 10, 18, 50}
	expectStr := [12][4]string{
		{"1111", "1111", "1111", "1111"},
		{"111", "111.2333", "111.2333", "111.2333"},
		{"11", "11", "11", "11"},
		{"238462326847234", "238462326847234.61637644", "238462326847234.61637644", "238462326847234.61637644"},
		{"39119032", "39119032.1111111111", "39119032.111111111111111111",
			"39119032.11111111111111111111111111111999999999999999999999"},
		{"132716", "132716.0000002734", "132716.00000027348248", "132716.00000027348248"},
		{"-1111", "-1111", "-1111", "-1111"},
		{"-111", "-111.2333", "-111.2333", "-111.2333"},
		{"-11", "-11", "-11", "-11"},
		{"-238462326847234", "-238462326847234.61637644", "-238462326847234.61637644", "-238462326847234.61637644"},
		{"-39119032", "-39119032.1111111111", "-39119032.111111111111111111",
			"-39119032.11111111111111111111111111111999999999999999999999"},
		{"-132716", "-132716.0000002734", "-132716.00000027348248", "-132716.00000027348248"},
	}
	for i, str := range testStr {
		for j, decimals := range testDecimals {
			parsedInt := ToIntByPrecise(str, decimals)
			t.Logf("parsed int is %d", parsedInt)
			parsedStr := ToStringByPrecise(parsedInt, decimals)
			t.Logf("parsed str is \"%s\",", parsedStr)
			assert.Equal(t, expectStr[i][j], parsedStr)
		}
	}
}

func TestReadExcelPlan(t *testing.T) {
	result := ReadExcelPlan("Actions-2.xlsx", "Sheet1")
	fmt.Println(result)
}

func TestStringToInt(t *testing.T) {
	fmt.Println(ToIntByPrecise("10",10))
}

func TestDecimal(t *testing.T) {
	transferDecimal := decimal.NewFromInt(1000000000000000000)

	baseAmount,err := decimal.NewFromString("52000000000000.6186277")
	if(err!=nil){
		fmt.Println(err)
	}
	transferValue := baseAmount.Mul(transferDecimal)
	log.Info(strconv.FormatInt(transferValue.IntPart(), 10))
	transferBigInt,two:= new(big.Int).SetString(strconv.FormatInt(transferValue.IntPart(), 18), 10)
	log.Info(two)
	//transferBigInt := ToIntByPrecise(baseAmount.String(), 6)

	log.Info("transfer amount is ", transferBigInt)
	log.Info(transferValue)
}
