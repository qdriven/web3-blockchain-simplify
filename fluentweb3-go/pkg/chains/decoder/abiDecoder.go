package decoder

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/mitchellh/mapstructure"
	"strings"
)

type (
	Abi abi.ABI
)

func NewAbi(abiStr string) (*Abi, error) {
	a, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, err
	}
	var abiObj = Abi(a)
	return &abiObj, nil
}

func (a *Abi) AsABI() abi.ABI {
	return abi.ABI(*a)
}

func (a *Abi) Decode(txInput string, inputData ...interface{}) (method *abi.Method, decodedInput map[string]interface{}, err error) {
	var ABI = a.AsABI()
	// decode txInput method signature
	decodedSig, err := hex.DecodeString(txInput[2:10])
	if err != nil {
		return nil, nil, err
	}

	method, err = ABI.MethodById(decodedSig)
	if err != nil {
		return nil, nil, err
	}

	// decode txInput Payload
	if len(txInput) < 10 {
		return nil, nil, nil
	}
	decodedData, err := hex.DecodeString(txInput[10:])
	if err != nil {
		return nil, nil, err
	}

	// unpack method inputs
	inputMap := make(map[string]interface{}, 0)
	err = method.Inputs.UnpackIntoMap(inputMap, decodedData)
	if err != nil {
		return nil, nil, err
	}

	if len(inputData) > 0 {
		err = mapstructure.Decode(inputMap, inputData[0])
		if err != nil {
			return nil, nil, err
		}
	}
	return method, inputMap, nil
}
