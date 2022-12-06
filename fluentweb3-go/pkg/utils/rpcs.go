package utils

import (
	"context"
	log "fluentweb3-go/pkg/logger"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
	"time"
)

func (q *ChainClient) TransactionByHash(hash common.Hash) *types.Transaction {
	tx, _, err := q.c.TransactionByHash(context.Background(), hash)
	if err != nil {
		fmt.Printf("TransactionByHash err: %s\n", err.Error())
	}
	return tx
}

func (q *ChainClient) GetEventByHash(hash common.Hash) ([]*types.Log, error) {
	raw := &types.Receipt{}
	if err := q.rpcClient.Call(raw, "eth_getTransactionReceipt", hash.Hex()); err != nil {
		return nil, fmt.Errorf("failed to get transaction receipt: [%v]", err)
	}
	return raw.Logs, nil
}
func (q *ChainClient) GetBlockNumber() uint64 {
	var raw string

	if err := q.rpcClient.Call(
		&raw,
		"eth_blockNumber",
	); err != nil {
		_ = fmt.Errorf("failed to get nonce: [%v]", err)
		return uint64(0)
	}

	without0xStr := strings.Replace(raw, "0x", "", -1)
	bigNonce, _ := new(big.Int).SetString(without0xStr, 16)
	return bigNonce.Uint64()
}

func (q *ChainClient) GetNonce(address string) uint64 {
	var raw string

	if err := q.rpcClient.Call(
		&raw,
		"eth_getTransactionCount",
		address,
		"latest",
	); err != nil {
		panic(fmt.Errorf("failed to get nonce: [%v]", err))
	}

	without0xStr := strings.Replace(raw, "0x", "", -1)
	bigNonce, _ := new(big.Int).SetString(without0xStr, 16)
	return bigNonce.Uint64()
}

func (q *ChainClient) GetProof(contractAddr common.Address, key string, blockNum string) (*Proof, error) {
	res := new(ProofRsp)
	if err := q.rpcClient.Call(res, "eth_getProof", contractAddr, []string{key}, blockNum); err != nil {
		return nil, err
	}

	if res.Error != nil {
		return nil, fmt.Errorf(res.Error.Message)
	}

	return &res.Result, nil
}

func (q *ChainClient) CallContract(caller, contractAddr common.Address, payload []byte, blockNum string) ([]byte, error) {
	var res hexutil.Bytes
	arg := ethereum.CallMsg{
		From: caller,
		To:   &contractAddr,
		Data: payload,
	}
	err := q.rpcClient.CallContext(context.Background(), &res, "eth_call", toCallArg(arg), blockNum)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}
func (q *ChainClient) WaitTransactionsConfirm(hashes []common.Hash) {
	hasPending := true
	for hasPending {
		time.Sleep(time.Second * 1)
		hasPending = false
		for _, hash := range hashes {
			_, isPending, err := q.c.TransactionByHash(context.Background(), hash)
			log.Infof("transaction %s is pending: %d", hash.String(), isPending)
			if err != nil {
				hasPending = true
				continue
			}
			if isPending == true {
				hasPending = true
			} else {
			}
		}
	}
}

func (q *ChainClient) WaitTransactionConfirm(hash common.Hash) {
	for {
		time.Sleep(time.Second * 1)
		_, ispending, err := q.c.TransactionByHash(context.Background(), hash)
		if err != nil {
			log.Errorf("failed to call TransactionByHash: %v", err)
			continue
		}
		if ispending == true {
			continue
		} else {
			break
		}
	}

}

func (q *ChainClient) GetReceipt(hash common.Hash) (*types.Receipt, error) {
	raw := &types.Receipt{}
	if err := q.rpcClient.Call(raw, "eth_getTransactionReceipt", hash.Hex()); err != nil {
		return nil, err
	}
	return raw, nil
}

func (q *ChainClient) DumpEventLog(hash common.Hash) error {
	raw, err := q.GetReceipt(hash)
	if err != nil {
		return fmt.Errorf("failed to get receipt %s", hash.Hex())
	}

	if raw.Status == 0 {
		return fmt.Errorf("receipt failed %s", hash.Hex())
	}

	log.Infof("txhash %s", hash.Hex())
	for _, event := range raw.Logs {
		log.Infof("eventlog address %s", event.Address.Hex())
		log.Infof("eventlog data %s", new(big.Int).SetBytes(event.Data).String())
		for i, topic := range event.Topics {
			log.Infof("eventlog topic[%d] %s", i, topic.String())
		}
	}
	return nil
}

//
//func (q *ChainClient) SendTransaction(contractAddr common.Address, payload []byte) (common.Hash, error) {
//	addr := q.c.Address()
//	nonce := c.GetNonce(addr.Hex())
//	if c.CurrentNonce < nonce {
//		c.CurrentNonce = nonce
//	}
//	log.Debugf("%s current nonce %d, valid nonce %d", addr.Hex(), c.CurrentNonce, nonce)
//	c.GasPrice, _ = c.EClient.SuggestGasPrice(nil)
//	tx := types.NewTransaction(
//		c.CurrentNonce,
//		contractAddr,
//		big.NewInt(0),
//		c.GasLimit,
//		c.GasPrice,
//		payload,
//	)
//	hash := tx.Hash()
//
//	signedTx, err := c.SignTransaction(tx)
//	if err != nil {
//		return hash, err
//	}
//	c.CurrentNonce += 1
//	return c.SendRawTransaction(hash, signedTx)
//}
//
//func (q *ChainClient) SendTransactionAndDumpEvent(contract common.Address, payload []byte) error {
//	hash, err := c.SendTransaction(contract, payload)
//	if err != nil {
//		return err
//	}
//	time.Sleep(c.BlockPeriod)
//	return c.DumpEventLog(hash)
//}
//
//func (c *EvmClient) RepeatSendTransactionAndDumpEvent(contract common.Address, payload []byte, repeat int) error {
//	hashList := make([]common.Hash, repeat)
//
//	for i := 0; i < repeat; i++ {
//		hash, err := c.SendTransaction(contract, payload)
//		if err != nil {
//			return err
//		}
//		hashList[i] = hash
//	}
//
//	time.Sleep(c.BlockPeriod)
//
//	for i := 0; i < repeat; i++ {
//		if err := c.DumpEventLog(hashList[i]); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func (c *EvmClient) SignTransaction(tx *types.Transaction) (string, error) {
//
//	signer := types.HomesteadSigner{}
//	signedTx, err := types.SignTx(
//		tx,
//		signer,
//		c.Key,
//	)
//	if err != nil {
//		return "", fmt.Errorf("failed to sign tx: [%v]", err)
//	}
//
//	bz, err := rlp.EncodeToBytes(signedTx)
//	if err != nil {
//		return "", fmt.Errorf("failed to rlp encode bytes: [%v]", err)
//	}
//	return "0x" + hex.EncodeToString(bz), nil
//}
//
//func (c *EvmClient) SendRawTransaction(hash common.Hash, signedTx string) (common.Hash, error) {
//	var result common.Hash
//	if err := c.Client.Call(&result, "eth_sendRawTransaction", signedTx); err != nil {
//		return hash, fmt.Errorf("failed to send raw transaction: [%v]", err)
//	}
//
//	return result, nil
//}

//func (q *ChainClient) DeployContract(abiStr, binStr string, params ...interface{}) (common.Address, *bind.BoundContract, error) {
//	auth := bind.NewKeyedTransactor(q.Key)
//	auth.GasLimit = 1e9
//	auth.Nonce = new(big.Int).SetUint64(q.GetNonce(q.Address().Hex()))
//	parsedABI, err := abi.JSON(strings.NewReader(abiStr))
//	if err != nil {
//		log.Errorf("failed to read abi json, err: %v", err)
//		return common.HexToAddress("0"), nil, err
//	}
//	parsedBin := common.FromHex(binStr)
//	backend := ethclient.NewClient(q.rpcClient)
//
//	address, tx, contract, err := bind.DeployContract(auth, parsedABI, parsedBin, backend, params...)
//	if err != nil {
//		return common.HexToAddress("0"), nil, err
//	}
//
//	log.Infof("deploy contract tx %v\r\n, contract %v\r\n, address %s\r\n", tx, contract, address.Hex())
//	return address, contract, nil
//}
