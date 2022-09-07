package sdk

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/contracts/native"
	"github.com/ethereum/go-ethereum/contracts/native/nft"
	"github.com/ethereum/go-ethereum/contracts/native/nftmanager"
	"github.com/ethereum/go-ethereum/contracts/native/utils"
	"github.com/ethereum/go-ethereum/core/types"
	polycm "github.com/polynetwork/poly/common"
	"math/big"
	"time"
	"chains-gotest-backend/api/plt/pkg/log"
)

var (
	NFTMangerAddress          = common.HexToAddress(native.NFTContractCreateAddress)
	NFTABI                    abi.ABI
	NFTManagerABI             abi.ABI
	EMPTY_ADDRESS             = common.BytesToAddress([]byte{})
	NFTEventID_Transfer       string
	NFTEventID_Approval       string
	NFTEventID_ApprovalForAll string
	NFTEventID_Deploy         string
)

type NFTTransferEvent struct {
	Event    string
	Contract common.Address
	Height   uint64
	TxHash   common.Hash
	From     common.Address
	To       common.Address
	Token    *big.Int
}

type NFTApprovalEvent struct {
	Event    string
	Contract common.Address
	Height   uint64
	TxHash   common.Hash
	Owner    common.Address
	Operator common.Address
	Token    *big.Int
}

type NFTApprovalForAllEvent struct {
	Event    string
	Contract common.Address
	Height   uint64
	TxHash   common.Hash
	Owner    common.Address
	Operator common.Address
	Approved bool
}

type NFTDeployEvent struct {
	Event    string
	Contract common.Address
	Height   uint64
	TxHash   common.Hash
	Owner    common.Address
	Nft      common.Address
}

func init() {
	NFTABI = nft.GetABI()
	NFTManagerABI = nftmanager.GetABI()
	NFTEventID_Transfer = NFTABI.Events[nft.EventTransfer].ID().String()
	NFTEventID_Approval = NFTABI.Events[nft.EventApproval].ID().String()
	NFTEventID_ApprovalForAll = NFTABI.Events[nft.EventApprovalForAll].ID().String()
	NFTEventID_Deploy = NFTABI.Events[nft.EventDeploy].ID().String()
}

func (c *Client) NFTDeploy(name string, symbol string) (common.Hash, common.Address, error) {
	payload, err := utils.PackMethod(NFTManagerABI, nftmanager.MethodDeploy, name, symbol, c.Address())
	if err != nil {
		return common.BytesToHash([]byte{}), common.BytesToAddress([]byte{}), err
	}

	txHash, err := c.SendNFTManagerTransaction(payload)
	if err != nil {
		return common.BytesToHash([]byte{}), common.BytesToAddress([]byte{}), err
	}
	tx := c.TransactionByHash(txHash)
	fmt.Printf("tx : %s\n", tx.String())
	counter := 0
	for counter < 60 {
		time.Sleep(time.Second * 1)
		event, err := c.GetDeployEvent(txHash)
		counter++
		if err != nil {
			continue
		}
		return txHash, event.Nft, nil
	}
	return common.BytesToHash([]byte{}), common.BytesToAddress([]byte{}), nil
}

func (c *Client) GetDeployEvent(txHash common.Hash) (*NFTDeployEvent, error) {
	events, err := c.GetEventByHash(txHash)
	if err != nil {
		return nil, err
	}
	if len(events) == 0 {
		return nil, fmt.Errorf("There is no event yet!")
	}
	for _, log := range events {
		if log.Topics[0] == NFTABI.Events[nft.EventDeploy].ID() {
			event := &NFTDeployEvent{
				Event:    nft.EventDeploy,
				Contract: log.Address,
				Height:   log.BlockNumber,
				TxHash:   log.TxHash,
				Owner:    common.BytesToAddress(log.Topics[1].Bytes()),
				Nft:      common.BytesToAddress(log.Topics[2].Bytes()),
			}
			return event, nil
		}
	}
	return nil, nil
}

func (c *Client) NFTName(nftAddr common.Address) (string, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodName)
	if err != nil {
		return "", err
	}
	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return "", err
	}
	fmt.Printf("data: %s\n", hex.EncodeToString(data))
	result := &nft.NameResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodName, result, data)
	if err != nil {
		return "", err
	}

	return result.Name, nil
}

func (c *Client) NFTSymbol(nftAddr common.Address) (string, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodSymbol)
	if err != nil {
		return "", err
	}
	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return "", err
	}
	fmt.Printf("data: %s\n", hex.EncodeToString(data))
	result := &nft.SymbolResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodSymbol, result, data)
	if err != nil {
		return "", err
	}

	return result.Symbol, nil
}

func (c *Client) NFTOwner(nftAddr common.Address) (common.Address, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodOwner)
	if err != nil {
		return EMPTY_ADDRESS, err
	}
	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return EMPTY_ADDRESS, err
	}
	fmt.Printf("data: %s\n", hex.EncodeToString(data))
	result := &nft.OwnerResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodOwner, result, data)
	if err != nil {
		return EMPTY_ADDRESS, err
	}

	return result.Owner, nil
}

func (c *Client) NFTTotalSupply(nftAddr common.Address) (*big.Int, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodTotalSupply)
	if err != nil {
		return big.NewInt(0), err
	}
	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return big.NewInt(0), err
	}
	fmt.Printf("data: %s\n", hex.EncodeToString(data))
	result := &nft.TotalSupplyResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodTotalSupply, result, data)
	if err != nil {
		return big.NewInt(0), err
	}

	return result.Supply, nil
}

func (c *Client) NFTMint(nftAddr common.Address, addr common.Address, token *big.Int, uri string) (common.Hash, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodMint, addr, token, uri)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}

	hash, err := c.SendNFTTransaction(nftAddr, payload)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}
	return hash, nil
}

func (c *Client) NFTBurn(nftAddr common.Address, token *big.Int) (common.Hash, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodBurn, token)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}

	hash, err := c.SendNFTTransaction(nftAddr, payload)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}
	return hash, nil
}

func (c *Client) NFTBalanceOf(nftAddr common.Address, addr common.Address) (*big.Int, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodBalanceOf, addr)
	if err != nil {
		return big.NewInt(0), err
	}

	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return big.NewInt(0), err
	}

	fmt.Printf("data: %s\n", hex.EncodeToString(data))

	result := &nft.BalanceOfResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodBalanceOf, result, data)
	if err != nil {
		return big.NewInt(0), err
	}

	return result.Balance, nil
}

func (c *Client) NFTOwnerOf(nftAddr common.Address, token *big.Int) (common.Address, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodOwnerOf, token)
	if err != nil {
		return EMPTY_ADDRESS, err
	}

	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return EMPTY_ADDRESS, err
	}

	fmt.Printf("data: %s\n", hex.EncodeToString(data))

	result := &nft.OwnerOfResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodOwnerOf, result, data)
	if err != nil {
		return EMPTY_ADDRESS, err
	}

	return result.Owner, nil
}

/*
func (c *Client) NFTSetTokenUri(nftAddr common.Address, token *big.Int, uri string) (common.Hash, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodSetTokenUri, token, uri)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}

	hash, err := c.SendNFTTransaction(nftAddr, payload)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}
	return hash, nil
}
*/

func (c *Client) NFTTokenUri(nftAddr common.Address, token *big.Int) (string, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodTokenUri, token)
	if err != nil {
		return "", err
	}

	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return "", err
	}

	fmt.Printf("data: %s\n", hex.EncodeToString(data))

	result := &nft.TokenUriResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodTokenUri, result, data)
	if err != nil {
		return "", err
	}

	return result.Uri, nil
}

func (c *Client) NFTTransferFrom(nftAddr common.Address, from common.Address, to common.Address, token *big.Int) (common.Hash, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodTransferFrom, from, to, token)
	if err != nil {
		return common.Hash{}, err
	}

	return c.SendNFTTransaction(nftAddr, payload)
}

func (c *Client) NFTRawSafeTransferFrom(nftAddr common.Address, from common.Address,
	to common.Address, token *big.Int, data []byte) (common.Hash, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodSafeTransferFrom, from, to, token, data)
	if err != nil {
		return common.Hash{}, err
	}

	return c.SendNFTTransaction(nftAddr, payload)
}

func (c *Client) NFTApprove(nftAddr common.Address, to common.Address, token *big.Int) (common.Hash, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodApprove, to, token)
	if err != nil {
		return common.Hash{}, err
	}

	return c.SendNFTTransaction(nftAddr, payload)
}

func (c *Client) NFTGetApproved(nftAddr common.Address, token *big.Int) (common.Address, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodGetApproved, token)
	if err != nil {
		return EMPTY_ADDRESS, err
	}

	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return EMPTY_ADDRESS, err
	}

	fmt.Printf("data: %s\n", hex.EncodeToString(data))

	result := &nft.GetApprovedResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodGetApproved, result, data)
	if err != nil {
		return EMPTY_ADDRESS, err
	}

	return result.Spender, nil
}

func (c *Client) NFTSetApprovalForAll(nftAddr common.Address, to common.Address, approved bool) (common.Hash, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodSetApprovalForAll, to, approved)
	if err != nil {
		return common.Hash{}, err
	}

	return c.SendNFTTransaction(nftAddr, payload)
}

func (c *Client) NFTIsApprovedForAll(nftAddr common.Address, owner common.Address, to common.Address) (bool, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodIsApprovedForAll, owner, to)
	if err != nil {
		return false, err
	}

	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return false, err
	}

	fmt.Printf("data: %s\n", hex.EncodeToString(data))

	result := &nft.IsApprovedForAllResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodIsApprovedForAll, result, data)
	if err != nil {
		return false, err
	}

	return result.Result, nil
}

func (c *Client) NFTTokenByIndex(nftAddr common.Address, index *big.Int) (*big.Int, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodTokenByIndex, index)
	if err != nil {
		return big.NewInt(0), err
	}

	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return big.NewInt(0), err
	}

	fmt.Printf("data: %s\n", hex.EncodeToString(data))

	result := &nft.TokenByIndexResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodTokenByIndex, result, data)
	if err != nil {
		return big.NewInt(0), err
	}

	return result.Token, nil
}

func (c *Client) NFTTokenOfOwnerByIndex(nftAddr common.Address, owner common.Address, index *big.Int) (*big.Int, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodTokenOfOwnerByIndex, owner, index)
	if err != nil {
		return big.NewInt(0), err
	}

	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return big.NewInt(0), err
	}

	fmt.Printf("data: %s\n", hex.EncodeToString(data))

	result := &nft.TokenOfOwnerByIndexResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodTokenOfOwnerByIndex, result, data)
	if err != nil {
		return big.NewInt(0), err
	}

	return result.Token, nil
}

func (c *Client) FilterNFTTransferLogs(nftAddr common.Address, startHeight int64, endHeight int64) ([]*NFTTransferEvent, error) {
	query := [][]interface{}{{NFTABI.Events[nft.EventTransfer].ID()}}
	topics, err := MakeTopics(query...)
	if err != nil {
		return nil, err
	}
	config := ethereum.FilterQuery{
		Addresses: []common.Address{nftAddr},
		Topics:    topics,
		FromBlock: new(big.Int).SetInt64(startHeight),
		ToBlock:   new(big.Int).SetInt64(endHeight),
	}
	logs, err := c.FilterLogs(context.Background(), config)
	if err != nil {
		return nil, err
	}
	events := make([]*NFTTransferEvent, 0)
	for _, log := range logs {
		event := &NFTTransferEvent{
			Event:    nft.EventTransfer,
			Contract: log.Address,
			Height:   log.BlockNumber,
			TxHash:   log.TxHash,
			From:     common.BytesToAddress(log.Topics[1].Bytes()),
			To:       common.BytesToAddress(log.Topics[2].Bytes()),
			Token:    new(big.Int).SetBytes(log.Data),
		}
		events = append(events, event)
	}
	return events, nil

}

func (c *Client) FilterNFTApprovalLogs(nftAddr common.Address, startHeight int64, endHeight int64) ([]*NFTApprovalEvent, error) {
	query := [][]interface{}{{NFTABI.Events[nft.EventApproval].ID()}}
	topics, err := MakeTopics(query...)
	if err != nil {
		return nil, err
	}
	config := ethereum.FilterQuery{
		Addresses: []common.Address{nftAddr},
		Topics:    topics,
		FromBlock: new(big.Int).SetInt64(startHeight),
		ToBlock:   new(big.Int).SetInt64(endHeight),
	}
	logs, err := c.FilterLogs(context.Background(), config)
	if err != nil {
		return nil, err
	}
	events := make([]*NFTApprovalEvent, 0)
	for _, log := range logs {
		event := &NFTApprovalEvent{
			Event:    nft.EventApproval,
			Contract: log.Address,
			Height:   log.BlockNumber,
			TxHash:   log.TxHash,
			Owner:    common.BytesToAddress(log.Topics[1].Bytes()),
			Operator: common.BytesToAddress(log.Topics[2].Bytes()),
			Token:    new(big.Int).SetBytes(log.Data),
		}
		events = append(events, event)
	}
	return events, nil

}

func (c *Client) FilterNFTApprovalForAllLogs(nftAddr common.Address, startHeight int64, endHeight int64) ([]*NFTApprovalForAllEvent, error) {
	query := [][]interface{}{{NFTABI.Events[nft.EventApprovalForAll].ID()}}
	topics, err := MakeTopics(query...)
	if err != nil {
		return nil, err
	}
	config := ethereum.FilterQuery{
		Addresses: []common.Address{nftAddr},
		Topics:    topics,
		FromBlock: new(big.Int).SetInt64(startHeight),
		ToBlock:   new(big.Int).SetInt64(endHeight),
	}
	logs, err := c.FilterLogs(context.Background(), config)
	if err != nil {
		return nil, err
	}
	events := make([]*NFTApprovalForAllEvent, 0)
	for _, log := range logs {
		event := &NFTApprovalForAllEvent{
			Event:    nft.EventApprovalForAll,
			Contract: log.Address,
			Height:   log.BlockNumber,
			TxHash:   log.TxHash,
			Owner:    common.BytesToAddress(log.Topics[1].Bytes()),
			Operator: common.BytesToAddress(log.Topics[2].Bytes()),
		}
		if bytes.Equal(log.Data, []byte{1}) {
			event.Approved = true
		}
		events = append(events, event)
	}
	return events, nil

}

func (c *Client) SendNFTTransaction(nft common.Address, payload []byte) (common.Hash, error) {
	addr := c.Address()
	nonce := c.GetNonce(addr.String())
	tx := types.NewTransaction(
		nonce,
		nft,
		big.NewInt(0),
		GasNormal,
		big.NewInt(GasPrice),
		payload,
	)

	signedTx, err := c.SignTransaction(tx)
	if err != nil {
		return common.Hash{}, err
	}
	return c.SendRawTransaction(tx.Hash(), signedTx)
}

func (c *Client) CallNFTContract(nft common.Address, payload []byte) ([]byte, error) {
	data, err := c.CallContract(c.Address(), nft, payload, "latest")
	if err != nil {
		return nil, err
	}
	return data, err
}

func (c *Client) SendNFTManagerTransaction(payload []byte) (common.Hash, error) {
	addr := c.Address()
	nonce := c.GetNonce(addr.String())
	tx := types.NewTransaction(
		nonce,
		NFTMangerAddress,
		big.NewInt(0),
		GasNormal,
		big.NewInt(GasPrice),
		payload,
	)

	signedTx, err := c.SignTransaction(tx)
	if err != nil {
		return common.Hash{}, err
	}
	return c.SendRawTransaction(tx.Hash(), signedTx)
}

func (c *Client) NFTSetBaseUri(nftAddr common.Address, uri string) (common.Hash, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodSetBaseUri, uri)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}

	hash, err := c.SendNFTTransaction(nftAddr, payload)
	if err != nil {
		return common.BytesToHash([]byte{}), err
	}
	return hash, nil
}

func (c *Client) NFTBaseUri(nftAddr common.Address) (string, error) {
	payload, err := utils.PackMethod(NFTABI, nft.MethodBaseUri)
	if err != nil {
		return "", err
	}
	data, err := c.CallNFTContract(nftAddr, payload)
	if err != nil {
		return "", err
	}

	result := &nft.BaseUriResult{}
	err = utils.UnpackOutputs(NFTABI, nft.MethodBaseUri, result, data)
	if err != nil {
		return "", err
	}

	return result.Uri, nil
}

func assembleSafeTransferCallData(toAddress common.Address, chainID uint64) []byte {
	sink := polycm.NewZeroCopySink(nil)
	sink.WriteVarBytes(toAddress.Bytes())
	sink.WriteUint64(chainID)
	return sink.Bytes()
}

// safe transfer to contract
func (c *Client) NFTSafeTransferFrom(
	asset common.Address,
	from common.Address,
	proxy common.Address,
	tokenID *big.Int,
	to common.Address,
	toChainID uint64,
) (common.Hash, error) {

	data := assembleSafeTransferCallData(to, toChainID)
	log.Infof("asset %s, from %s, proxy %s, tokenID %d, data %s",
		asset.Hex(), from.Hex(), proxy.Hex(), tokenID.Uint64(), hexutil.Encode(data))

	payload, err := c.packNFT(nft.MethodSafeTransferFrom, from, proxy, tokenID, data)
	if err != nil {
		return utils.EmptyHash, err
	}

	return c.sendNFT(asset, payload)
}

// NFT
func (c *Client) packNFT(method string, args ...interface{}) ([]byte, error) {
	return utils.PackMethod(NFTABI, method, args...)
}
func (c *Client) unpackNFT(method string, output interface{}, enc []byte) error {
	return utils.UnpackOutputs(NFTABI, method, output, enc)
}
func (c *Client) sendNFT(nftAddr common.Address, payload []byte) (common.Hash, error) {
	return c.SendTransaction(nftAddr, payload)
}
func (c *Client) callNFT(nftAddr common.Address, payload []byte, blockNum string) ([]byte, error) {
	return c.CallContract(c.Address(), nftAddr, payload, blockNum)
}

// nft manager
func (c *Client) packNFTManager(method string, args ...interface{}) ([]byte, error) {
	return utils.PackMethod(NFTManagerABI, method, args...)
}
func (c *Client) unpackNFTManager(method string, output interface{}, enc []byte) error {
	return utils.UnpackOutputs(NFTManagerABI, method, output, enc)
}
func (c *Client) sendNFTManager(payload []byte) (common.Hash, error) {
	return c.SendTransaction(NFTMangerAddress, payload)
}
func (c *Client) callNFTManager(payload []byte, blockNum string) ([]byte, error) {
	return c.CallContract(c.Address(), NFTMangerAddress, payload, blockNum)
}
