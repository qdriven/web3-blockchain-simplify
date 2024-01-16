package rpctest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"gotest.tools/assert"
)

// request rpc
// compare result
//   order both config result and response result by their fields
//   json marshal then amd compare
func DoClientTest(t *testing.T, config RpcTestConfig) {
	handler := CallExampleResultHandler(func(example RpcExample, rpcResult interface{}, rpcError interface{}) error {
		if example.Error != nil || rpcError != nil {
			assert.Equal(t, mustJsonMarshalForTest(example.Error), mustJsonMarshalForTest(rpcError))
		} else {
			assert.Equal(t, mustJsonMarshalForTest(example.Result), mustJsonMarshalForTest(rpcResult))
		}
		return nil
	})
	if err := ExecuteExamples(config, handler); err != nil {
		t.Fatal(err)
	}
}

func ExecuteExamples(config RpcTestConfig, handler CallExampleResultHandler) error {
	rpc2Func, rpc2FuncSelector, rpc2FuncResultHandler := config.Rpc2Func, config.Rpc2FuncSelector, config.Rpc2FuncResultHandler
	ignoreRpc, ignoreExamples, onlyTestRpc, onlyExamples := config.IgnoreRpcs, config.IgnoreExamples, config.OnlyTestRpcs, config.OnlyExamples

	m, err := getMockExamples(config.ExamplesUrl)
	if err != nil {
		return errors.Wrap(err, "failed to get test examples")
	}

	for rpcName, subExamps := range m.Examples {
		if ignoreRpc[rpcName] {
			continue
		}

		if len(onlyTestRpc) > 0 && !onlyTestRpc[rpcName] {
			continue
		}

		for _, subExamp := range subExamps {

			if ignoreExamples[subExamp.Name] {
				continue
			}

			if len(onlyExamples) > 0 && !onlyExamples[subExamp.Name] {
				continue
			}

			var sdkFunc string
			var params []interface{}

			if _sdkFunc, ok := rpc2Func[rpcName]; ok {
				sdkFunc, params = _sdkFunc, subExamp.Params
			}

			if sdkFuncSelector, ok := rpc2FuncSelector[rpcName]; ok {
				sdkFunc, params = sdkFuncSelector(subExamp.Params)
			}

			if sdkFunc == "" {
				return errors.New("no sdk func for rpc:" + rpcName)
				// t.Fatalf("no sdk func for rpc:%s", rpcName)
			}

			fmt.Printf("\n========== example: %v === rpc: %s === params: %s ==========\n", subExamp.Name, rpcName, mustJsonMarshalForTest(params))
			// reflect call sdkFunc
			rpcReuslt, rpcError, err := reflectCall(config.Client, sdkFunc, params)
			if err != nil {
				var tmp interface{} = err
				switch tmp.(type) {
				case ConvertParamError:
					if subExamp.Error != nil {
						fmt.Printf("convert param err:%v, expect error %v\n", err, subExamp.Error)
						continue
					}
				}
				return errors.Wrapf(err, "failed to call example %v", rpcName)
			}

			if sdkFuncResultHandler, ok := rpc2FuncResultHandler[rpcName]; ok {
				rpcReuslt = sdkFuncResultHandler(rpcReuslt)
			}

			if handler == nil {
				continue
			}

			if err := handler(subExamp, rpcReuslt, rpcError); err != nil {
				return err
			}
		}
	}
	return nil
}

func getMockExamples(url string) (*MockRPC, error) {
	// read json config
	httpClient := &http.Client{}
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	source := resp.Body
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(source)
	if err != nil {
		return nil, err
	}

	m := &MockRPC{}
	err = json.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func reflectCall(c interface{}, sdkFunc string, params []interface{}) (resp interface{}, respError interface{}, err error) {
	typeOfClient := reflect.TypeOf(c)
	if method, ok := typeOfClient.MethodByName(sdkFunc); ok {
		in := make([]reflect.Value, len(params)+1)
		in[0] = reflect.ValueOf(c)
		// params marshal/unmarshal -> func params type
		for i, param := range params {
			// unmarshal params
			pType := method.Type.In(i + 1)

			// get element type if is variadic function for last param
			if method.Type.IsVariadic() && i == method.Type.NumIn()-2 {
				pType = pType.Elem()
			}

			vPtr := reflect.New(pType).Interface()
			vPtr, err = convertType(param, vPtr)
			if err != nil {
				return nil, nil, ConvertParamError(err)
			}
			v := reflect.ValueOf(vPtr).Elem().Interface()
			in[i+1] = reflect.ValueOf(v)
		}
		out := method.Func.Call(in)
		fmt.Printf("\n-- invoke func --\n\tname: %v, \n\tparams: %v, \n\tresp type: %T, \n\trespError type: %T, \n\tresp value: %v, \n\trespError value: %v\n",
			sdkFunc,
			mustJsonMarshalForTest(getReflectValuesInterfaces(in[1:])),
			out[0].Interface(),
			out[1].Interface(),
			mustJsonMarshalForTest(out[0].Interface(), true),
			mustJsonMarshalForTest(out[1].Interface(), true),
		)
		return out[0].Interface(), out[1].Interface(), nil
	}
	return nil, nil, errors.Errorf("not found method %v", sdkFunc)
}

func getReflectValuesInterfaces(values []reflect.Value) []interface{} {
	var result []interface{}
	for _, v := range values {
		result = append(result, v.Interface())
	}
	return result
}

// cfx_getBlockByEpochNumber  GetBlockSummaryByEpoch 0x0, false
// rpc_name => func(params) sdkFuncName sdkFuncParams
func convertType(from interface{}, to interface{}) (interface{}, error) {
	jp, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jp, &to)
	if err != nil {
		return nil, err
	}
	return to, nil
}

func mustConvertType(from interface{}, to interface{}) interface{} {
	v, err := convertType(from, to)
	if err != nil {
		panic(err)
	}
	return v
}

func mustJsonMarshalForTest(v interface{}, indent ...bool) string {
	j, err := JsonMarshalForRpcTest(v, indent...)
	if err != nil {
		panic(err)
	}
	return string(j)
}

// Block
// 	BlockHash
// 	[]Transactions
// 		Creates 'testomit:false'

// JsonMarshalForRpcTest recursively handle fields of struct or sub-struct of v by 'testomit' tag and order json
func JsonMarshalForRpcTest(v interface{}, indent ...bool) ([]byte, error) {
	// fmt.Printf("reflect.ValueOf(v).Kind(): %v\n", reflect.ValueOf(v).Kind())

	converted := mustConvertType(v, interface{}(nil))

	// if converted result type same with type of v, then return
	if reflect.TypeOf(converted) == reflect.TypeOf(v) {
		return json.Marshal(v)
	}

	converted = setFieldsOmitByTag(converted, reflect.TypeOf(v))

	if isIndent(indent...) {
		return json.MarshalIndent(converted, "", "  ")
	} else {
		return json.Marshal(converted)
	}
}

// setFieldsOmitByTag set fields omit by 'testomit' tag
// v must be marshal result of object, t is real object type
func setFieldsOmitByTag(v interface{}, t reflect.Type) interface{} {
	// fmt.Printf("setTestOmit: v: %v, t: %v\n", v, t)
	if t == nil {
		panic(fmt.Sprintf("t of %v is nil", v))
	}

	t = getCoreType(t)

	// only struct/array/slice need be handle
	if t.Kind() != reflect.Struct && t.Kind() != reflect.Array && t.Kind() != reflect.Slice {
		return v
	}

	if method, ok := t.MethodByName("MarshalJSONForRPCTest"); ok {
		// unmarshal to t
		strongValue := reflect.New(t).Interface()
		strongValue, err := convertType(v, strongValue)
		if err != nil {
			panic(err)
		}
		// call marshalJsonForRpcTest
		reflectJ := method.Func.Call([]reflect.Value{reflect.ValueOf(strongValue).Elem()})
		if err != nil {
			panic(err)
		}

		j, ok := reflectJ[0].Interface().([]byte)
		if !ok {
			panic(errors.Errorf("%v.MarshalJSONForRPCTest() return type is not []byte", t))
		}

		// fmt.Printf("MarshalJSONForRPCTest result %v\n", string(j))

		var i interface{}
		err = json.Unmarshal(j, &i)
		if err != nil {
			panic(err)
		}
		return i
	}

	switch v.(type) {
	case map[string]interface{}:
		break
	case []interface{}:
		raw := v.([]interface{})
		for i, vv := range raw {
			raw[i] = setFieldsOmitByTag(vv, t.Elem())
		}
		return raw
	default:
		return v
	}

	m := v.(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		fName := tf.Name
		if jsonTag, ok := tf.Tag.Lookup("json"); ok {
			fName, _ = parseJsonTag(jsonTag)
		}

		// recursive fields
		if m[fName] != nil {
			m[fName] = setFieldsOmitByTag(m[fName], tf.Type)
			continue
		}

		isOmit, ok := tf.Tag.Lookup("testomit")
		if !ok {
			continue
		}

		if isOmit == "true" {
			delete(m, fName)
			continue
		}

		m[fName] = nil
	}
	return m
}

func isSelfOrElemBeStruct(v interface{}) bool {
	reflectV := reflect.ValueOf(v)

	if reflectV.Kind() != reflect.Ptr && reflectV.Kind() != reflect.Struct {
		return false
	}

	if reflectV.Kind() == reflect.Ptr && reflectV.Elem().Kind() != reflect.Struct {
		fmt.Printf("reflect.ValueOf(v).Elem().Kind(): %v\n", reflectV.Elem().Kind())
		return false
	}
	return true
}

// Get type of self or elem type if pointer
func getCoreType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

func parseJsonTag(jsonTag string) (jsonName string, isOmitEmpty bool) {
	splits := strings.Split(jsonTag, ",")
	if len(splits) == 1 {
		return splits[0], false
	}
	return splits[0], strings.Contains(splits[1], "omitempty")
}

func isIndent(indent ...bool) bool {
	_isIndent := false
	if len(indent) > 0 {
		_isIndent = indent[0]
	}
	return _isIndent
}
