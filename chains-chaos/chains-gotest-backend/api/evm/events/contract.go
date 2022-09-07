package events

import (
	"chains-gotest-backend/api/evm/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

type Contract struct {
	Name    string
	ABI     abi.ABI
	Address common.Address

	// event ID <-> event Name mapping
	events map[common.Hash]string
}
type ContractEvent struct {
	// block in which the transaction was included
	BlockNumber uint64
	// hash of the block in which the transaction was included
	BlockHash common.Hash
	// hash of the transaction
	TxHash common.Hash

	// contract which the event belongs to
	Contract *Contract
	// name of the contract event
	Name string
	// supplied by the contract, usually ABI-encoded
	Data []byte

	// The Removed field is true if this log was reverted due to a chain reorganisation.
	// You must pay attention to this field if you receive logs through a filter query.
	Removed bool
}


func NewContract(name string, abiJSON string, address common.Address) (*Contract, error) {
	c := &Contract{
		Name:    name,
		Address: address,
		events:  make(map[common.Hash]string),
	}

	abi, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		log.Error("Failed to parse ABI interface", "name", name, "abi", abiJSON, "address", address.Hex(), "err", err)
		return nil, err
	}
	c.ABI = abi
	for _, event := range abi.Events {
		c.events[event.ID()] = event.Name
	}
	return c, nil
}
