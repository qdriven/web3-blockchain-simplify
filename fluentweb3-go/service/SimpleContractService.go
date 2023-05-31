package service

import (
	"fluentweb3-go/chains/evm"
	"log"
)

func Publish() {
	evm, err := evm.NewEthUtil("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	print(evm)
}
