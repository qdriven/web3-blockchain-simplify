package scd

import (
	"context"
	"encoding/hex"
	"fluentweb3-go/pkg/address"
	log "fluentweb3-go/pkg/logger"
	"fluentweb3-go/pkg/utils"
	"fluentweb3-go/scd/abi"
	"fluentweb3-go/scd/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

const RPC_URL = "http://10.100.90.82:8545"

var chainId = big.NewInt(1337)

type WrappedCdpContractV1 struct {
	CdpV1       *contract.Cdp
	ChainClient *ethclient.Client
	AbiDecoder  *utils.ABIDecoder
}

type WrappedTxData struct {
	OnChainTime     uint64
	BlockHash       string
	BlockNumber     *big.Int
	TransactionHash string
	Logs            *[]utils.DecodedLog
}

func NewCdpContractV1(contractAddressHex string) *WrappedCdpContractV1 {
	client := utils.DialNode(RPC_URL)
	contractAddress := common.HexToAddress(contractAddressHex)
	cdpContractV1, err := contract.NewCdp(contractAddress, client)
	decoder := utils.NewABIDecoderByAbi(abi.CdpV1Abi)
	if err != nil {
		panic(fmt.Sprintf("failed to load cdp contract v1 [%s]", contractAddress))
	}

	return &WrappedCdpContractV1{
		CdpV1: cdpContractV1, ChainClient: client, AbiDecoder: decoder,
	}
}

func (w *WrappedCdpContractV1) AddToWhiteList(fromAddress *address.WrappedAddress) (*types.Transaction, error) {
	auth := w.CreateTransactOpts(fromAddress)
	tx, err := w.CdpV1.AddToWhiteList(auth, fromAddress.Address)
	return tx, err
}

func (w *WrappedCdpContractV1) RemoveFromWhiteList(fromAddress *address.WrappedAddress) (*types.Transaction, error) {
	auth := w.CreateTransactOpts(fromAddress)
	tx, err := w.CdpV1.RemoveFromWhiteList(auth, fromAddress.Address)
	return tx, err
}

func (w *WrappedCdpContractV1) AddEmissionData(fromAddress *address.WrappedAddress, emissionData string) (*types.Transaction, error) {
	auth := w.CreateTransactOpts(fromAddress)
	tx, err := w.CdpV1.AddEmissionData(auth, emissionData)
	return tx, err
}

func (w *WrappedCdpContractV1) GetEmissionData() ([]contract.EmissionAppContractEmissionData, error) {
	result, err := w.CdpV1.GetEmisionData(nil)
	return result, err
}

func (w *WrappedCdpContractV1) CreateTransactOpts(fromAddress *address.WrappedAddress) *bind.TransactOpts {
	nonce, _ := w.ChainClient.PendingNonceAt(context.Background(), fromAddress.Address)
	auth, _ := bind.NewKeyedTransactorWithChainID(fromAddress.PriKey, chainId)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = big.NewInt(0)
	return auth
}

func (w *WrappedCdpContractV1) GetTransactionData(txHashStr string) *WrappedTxData {
	txHash := common.HexToHash(txHashStr)
	txInfo := &WrappedTxData{}
	for {
		txDetail, pending, err := w.ChainClient.TransactionByHash(context.Background(), txHash)
		if err != nil {
			log.Errorf("failed to call TransactionByHash: %v", err)
			continue
		}
		if pending == true {
			time.Sleep(time.Millisecond * 250)
			continue
		} else {
			method, err := w.AbiDecoder.DecodeMethod(hex.EncodeToString(txDetail.Data()))
			log.Info("method nam is ", method)
			utils.ParseTransactionBaseInfo(txDetail)
			if err != nil {
				log.Fatal(err)
			}
			txReceipt, _ := w.ChainClient.TransactionReceipt(context.Background(), txHash)
			txInfo.BlockNumber = txReceipt.BlockNumber
			txInfo.BlockHash = txReceipt.BlockHash.Hex()
			txInfo.TransactionHash = txReceipt.TxHash.Hex()
			block, _ := w.ChainClient.BlockByHash(context.Background(), txReceipt.BlockHash)
			txInfo.OnChainTime = block.Time()
			decodedLogs, _ := w.AbiDecoder.DecodeLogs(txReceipt.Logs)
			txInfo.Logs = &decodedLogs
			break
		}
	}
	return txInfo
}
