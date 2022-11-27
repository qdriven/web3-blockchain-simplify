package sdk

import (
	"chains-gotest-backend/api/plt/pkg/log"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/contracts/native"
	"github.com/ethereum/go-ethereum/contracts/native/governance"
	"github.com/ethereum/go-ethereum/contracts/native/plt"
	"github.com/ethereum/go-ethereum/contracts/native/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	PLTABI, GovernanceABI abi.ABI
	PLTAddress            = common.HexToAddress(native.PLTContractAddress)
	GovernanceAddress     = common.HexToAddress(native.GovernanceContractAddress)

	gasLimit, deployGasLimit uint64
	blockPeriod              time.Duration
)

const (
	gasPrice = 0
)

func Init(_gasLimit, _deployGasLimit uint64, _blockPeriod time.Duration) {
	PLTABI = plt.GetABI()
	GovernanceABI = governance.GetABI()

	gasLimit = _gasLimit
	deployGasLimit = _deployGasLimit
	blockPeriod = _blockPeriod
}

func (c *Client) GetBlockNumber() uint64 {
	var raw string

	if err := c.Call(
		&raw,
		"eth_blockNumber",
	); err != nil {
		fmt.Errorf("failed to get nonce: [%v]", err)
		return uint64(0)
	}

	without0xStr := strings.Replace(raw, "0x", "", -1)
	bigNonce, _ := new(big.Int).SetString(without0xStr, 16)
	return bigNonce.Uint64()
}

func (c *Client) GetNonce(address string) uint64 {
	var raw string

	if err := c.Call(
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

func (c *Client) SendTransaction(contractAddr common.Address, payload []byte) (common.Hash, error) {
	addr := c.Address()

	nonce := c.GetNonce(addr.Hex())
	if c.currentNonce < nonce {
		c.currentNonce = nonce
	}
	log.Debugf("%s current nonce %d, valid nonce %d", addr.Hex(), c.currentNonce, nonce)
	tx := types.NewTransaction(
		c.currentNonce,
		contractAddr,
		big.NewInt(0),
		gasLimit,
		big.NewInt(gasPrice),
		payload,
	)
	hash := tx.Hash()

	signedTx, err := c.SignTransaction(tx)
	if err != nil {
		return hash, err
	}
	c.currentNonce += 1
	return c.SendRawTransaction(hash, signedTx)
}

func (c *Client) SendTransactionAndDumpEvent(contract common.Address, payload []byte) error {
	hash, err := c.SendTransaction(contract, payload)
	if err != nil {
		return err
	}
	time.Sleep(blockPeriod)
	return c.DumpEventLog(hash)
}

func (c *Client) RepeatSendTransactionAndDumpEvent(contract common.Address, payload []byte, repeat int) error {
	hashList := make([]common.Hash, repeat)

	for i := 0; i < repeat; i++ {
		hash, err := c.SendTransaction(contract, payload)
		if err != nil {
			return err
		}
		hashList[i] = hash
	}

	time.Sleep(blockPeriod)

	for i := 0; i < repeat; i++ {
		if err := c.DumpEventLog(hashList[i]); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) SignTransaction(tx *types.Transaction) (string, error) {

	signer := types.HomesteadSigner{}
	signedTx, err := types.SignTx(
		tx,
		signer,
		c.Key,
	)
	if err != nil {
		return "", fmt.Errorf("failed to sign tx: [%v]", err)
	}

	bz, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to rlp encode bytes: [%v]", err)
	}
	return "0x" + hex.EncodeToString(bz), nil
}

func (c *Client) SendRawTransaction(hash common.Hash, signedTx string) (common.Hash, error) {
	var result common.Hash
	if err := c.Client.Call(&result, "eth_sendRawTransaction", signedTx); err != nil {
		return hash, fmt.Errorf("failed to send raw transaction: [%v]", err)
	}

	return result, nil
}

func (c *Client) DeployContract(abiStr, binStr string, params ...interface{}) (common.Address, *bind.BoundContract, error) {
	auth := bind.NewKeyedTransactor(c.Key)
	auth.GasLimit = 1e9
	auth.Nonce = new(big.Int).SetUint64(c.GetNonce(c.Address().Hex()))

	parsedABI, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Errorf("failed to read abi json, err: %v", err)
		return utils.EmptyAddress, nil, err
	}
	parsedBin := common.FromHex(binStr)
	backend := ethclient.NewClient(c.Client)

	address, tx, contract, err := bind.DeployContract(auth, parsedABI, parsedBin, backend, params...)
	if err != nil {
		return utils.EmptyAddress, nil, err
	}

	log.Infof("deploy contract tx %v\r\n, contract %v\r\n, address %s\r\n", tx, contract, address.Hex())
	return address, contract, nil
}

func (c *Client) getDeployAuth() *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(c.Key)
	auth.GasLimit = 1e7
	auth.Nonce = new(big.Int).SetUint64(c.GetNonce(c.Address().Hex()))
	return auth
}

func (c *Client) getTransactAuth() *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(c.Key)
	auth.GasLimit = 21000
	return auth
}

func (c *Client) getCallOpts() *bind.CallOpts {
	auth := new(bind.CallOpts)
	auth.From = c.Address()
	auth.Context = context.Background()
	return auth
}

func (c *Client) DumpEventLog(hash common.Hash) error {
	raw, err := c.GetReceipt(hash)
	if err != nil {
		return fmt.Errorf("faild to get receipt %s", hash.Hex())
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

func (c *Client) GetReceipt(hash common.Hash) (*types.Receipt, error) {
	raw := &types.Receipt{}
	if err := c.Call(raw, "eth_getTransactionReceipt", hash.Hex()); err != nil {
		return nil, err
	}
	return raw, nil
}

type ProofRsp struct {
	JsonRPC string       `json:"jsonrpc"`
	Result  PaletteProof `json:"result,omitempty"`
	Error   *JsonError   `json:"error,omitempty"`
	Id      uint         `json:"id"`
}

type PaletteProof struct {
	Address       string         `json:"address"`
	Balance       string         `json:"balance"`
	CodeHash      string         `json:"codeHash"`
	Nonce         string         `json:"nonce"`
	StorageHash   string         `json:"storageHash"`
	AccountProof  []string       `json:"accountProof"`
	StorageProofs []StorageProof `json:"StorageProof"`
}

type StorageProof struct {
	Key   string   `json:"key"`
	Value string   `json:"value"`
	Proof []string `json:"proof"`
}

type JsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (c *Client) GetProof(contractAddr common.Address, key string, blockNum string) (*PaletteProof, error) {
	res := new(ProofRsp)
	if err := c.Call(res, "eth_getProof", contractAddr, []string{key}, blockNum); err != nil {
		return nil, err
	}

	if res.Error != nil {
		return nil, fmt.Errorf(res.Error.Message)
	}

	return &res.Result, nil
}

func (c *Client) CallContract(caller, contractAddr common.Address, payload []byte, blockNum string) ([]byte, error) {
	var res hexutil.Bytes
	arg := ethereum.CallMsg{
		From: caller,
		To:   &contractAddr,
		Data: payload,
	}
	err := c.CallContext(context.Background(), &res, "eth_call", toCallArg(arg), blockNum)
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

func (c *Client) makeDeployAuth() *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(c.Key)
	auth.GasLimit = 1e7
	auth.Nonce = new(big.Int).SetUint64(c.GetNonce(c.Address().Hex()))
	return auth
}

func (c *Client) makeAuth() *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(c.Key)
	auth.GasLimit = 2100000
	return auth
}
