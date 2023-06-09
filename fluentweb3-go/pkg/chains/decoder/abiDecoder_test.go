package decoder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"testing"
)

func TestDecodeContractFunc(t *testing.T) {
	txInput := ""
	fileBytes, err := ioutil.ReadFile("uniswapv2.json")
	if err != nil {
		logrus.Error(err)
		return
	}
	a, err := NewAbi(string(fileBytes))
	if err != nil {
		logrus.Error(err)
		return
	}
	method, decodeInput, err := a.Decode(txInput)
	if err != nil {
		logrus.Error(err)
		return
	}
	res := struct {
		Function string      `json:"function"`
		Id       string      `json:"id"`
		Data     interface{} `json:"data"`
	}{
		Function: method.String(),
		Id:       hexutil.Encode(method.ID),
		Data:     decodeInput,
	}
	bs, err := json.Marshal(res)
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(jsonPrettyPrint(bs))
}

func jsonPrettyPrint(bs []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, bs, "", "    ")
	if err != nil {
		return string(bs)
	}
	return out.String()
}
