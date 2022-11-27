package service

import "net/http"

func GetEthTestCoin(address string) (*http.Response, error) {
	url := "https://faucet.ropsten.be/donate/"
	resp, err := http.Get(url + address)
	return resp, err
}


//1,999,999,994.998199999999899999

//TODO: Transaction Balances