package rpctest

type MarshalerForRpcTest interface {
	MarshalJSONForRPCTest(indent ...bool) ([]byte, error)
}
