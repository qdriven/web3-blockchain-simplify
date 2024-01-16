package rpctest

type ConvertParamError error

type CallExampleResultHandler func(example RpcExample, rpcReulst interface{}, rpcError interface{}) error

type RpcExample struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Params      []interface{} `json:"params"`
	Result      interface{}   `json:"result"`
	Error       interface{}   `json:"error"`
}
type MockRPC struct {
	Version  string                  `json:"version"`
	Examples map[string][]RpcExample `json:"examples"`
}

type RpcTestConfig struct {
	ExamplesUrl string
	Client      interface{}

	Rpc2Func         map[string]string
	Rpc2FuncSelector map[string]func(params []interface{}) (realFuncName string, realParams []interface{})
	// convert sdk rpc result back to pre-unmarshal for comparing with example result, becasue sdk may change result type for user convinent, such as web3go
	Rpc2FuncResultHandler map[string]func(result interface{}) (handlerdResult interface{})
	// ignored testing rpcs, ignoreRpc priority is higher than onlyTestRpc
	IgnoreRpcs map[string]bool
	// onlyTestRpc priority is lower than ignoreRpc
	OnlyTestRpcs map[string]bool
	// ignored testing examples, ignoreExamples priority is higher than onlyExamples
	IgnoreExamples map[string]bool
	// onlyExamples priority is lower than ignoreExamples
	OnlyExamples map[string]bool
}
