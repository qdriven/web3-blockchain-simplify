package evm

import (
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKeyStore(walletPath string, pwd string) accounts.Account {
	ks := keystore.NewKeyStore(walletPath, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(pwd)
	if err != nil {
		log.Fatal(err)
	}
	return account
}

func LoadKeystore(keystorePath string, pwd string) accounts.Account {
	jsonBytes, err := ioutil.ReadFile(keystorePath)
	if err != nil {
		log.Fatal(err)
	}
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.Import(jsonBytes, pwd, pwd)
	if err != nil {
		log.Fatal(err)
	}
	return account
}
