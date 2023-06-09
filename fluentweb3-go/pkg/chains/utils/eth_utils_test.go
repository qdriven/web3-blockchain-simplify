package utils

import (
	"log"
	"testing"
)

func TestNewEthUtil(t *testing.T) {
	evm, err := NewEthUtil("http://127.0.0.1:8545/")
	if err != nil {
		log.Fatal(err)
	}
	println(evm)
}
