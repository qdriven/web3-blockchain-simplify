# ETH Intro

- ETH Client Setup 
- ETH 

## ETH Client Setup

- Local: ganache
```sh 
npm install -g ganache-cli
```

- setup go-ethereum client

```golang 
func NeEthClient(ethConfig *EthConfig) *ethclient.Client {
	client,err := ethclient.Dial(ethConfig.EthNodeRpcAddress)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("eth client is connected.......")
	return client
}
```