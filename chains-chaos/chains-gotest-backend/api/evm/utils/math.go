package utils

import (
	"math/big"
	"strings"
)

func ToStringByPrecise(bigNum *big.Int, precise uint64) string {
	result := ""
	destStr := bigNum.String()
	destLen := uint64(len(destStr))
	if precise >= destLen { // add "0.000..." at former of destStr
		var i uint64 = 0
		prefix := "0."
		for ; i < precise-destLen; i++ {
			prefix += "0"
		}
		result = prefix + destStr
	} else { // add "."
		pointIndex := destLen - precise
		result = destStr[0:pointIndex] + "." + destStr[pointIndex:]
	}
	result = removeZeroAtTail(result)
	return result
}

// delete no need "0" at last of result
func removeZeroAtTail(str string) string {
	i := len(str) - 1
	for ; i >= 0; i-- {
		if str[i] != '0' {
			break
		}
	}
	str = str[:i+1]
	// delete "." at last of result
	if str[len(str)-1] == '.' {
		str = str[:len(str)-1]
	}
	return str
}

func ToIntByPrecise(str string, precise uint64) *big.Int {
	result := new(big.Int)
	splits := strings.Split(str, ".")
	if len(splits) == 1 { // doesn't contain "."
		var i uint64 = 0
		for ; i < precise; i++ {
			str += "0"
		}
		intValue, ok := new(big.Int).SetString(str, 10)
		if ok {
			result.Set(intValue)
		}
	} else if len(splits) == 2 {
		value := new(big.Int)
		ok := false
		floatLen := uint64(len(splits[1]))
		if floatLen <= precise { // add "0" at last of str
			parseString := strings.Replace(str, ".", "", 1)
			var i uint64 = 0
			for ; i < precise-floatLen; i++ {
				parseString += "0"
			}
			value, ok = value.SetString(parseString, 10)
		} else { // remove redundant digits after "."
			splits[1] = splits[1][:precise]
			parseString := splits[0] + splits[1]
			value, ok = value.SetString(parseString, 10)
		}
		if ok {
			result.Set(value)
		}
	}

	return result
}
