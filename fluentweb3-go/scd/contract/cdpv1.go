// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EmissionAppContractEmissionData is an auto generated low-level Go binding around an user-defined struct.
type EmissionAppContractEmissionData struct {
	Owner common.Address
	Data  string
}

// CdpMetaData contains all meta data concerning the Cdp contract.
var CdpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"AddEmissionDataEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"targetVal\",\"type\":\"bool\"}],\"name\":\"AuthChangeEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"emissonData\",\"type\":\"string\"}],\"name\":\"addEmissionData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addToWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmisionData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"internalType\":\"structEmissionAppContract.EmissionData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeFromWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CdpABI is the input ABI used to generate the binding from.
// Deprecated: Use CdpMetaData.ABI instead.
var CdpABI = CdpMetaData.ABI

// Cdp is an auto generated Go binding around an Ethereum contract.
type Cdp struct {
	CdpCaller     // Read-only binding to the contract
	CdpTransactor // Write-only binding to the contract
	CdpFilterer   // Log filterer for contract events
}

// CdpCaller is an auto generated read-only Go binding around an Ethereum contract.
type CdpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CdpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CdpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CdpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CdpSession struct {
	Contract     *Cdp              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CdpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CdpCallerSession struct {
	Contract *CdpCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CdpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CdpTransactorSession struct {
	Contract     *CdpTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CdpRaw is an auto generated low-level Go binding around an Ethereum contract.
type CdpRaw struct {
	Contract *Cdp // Generic contract binding to access the raw methods on
}

// CdpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CdpCallerRaw struct {
	Contract *CdpCaller // Generic read-only contract binding to access the raw methods on
}

// CdpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CdpTransactorRaw struct {
	Contract *CdpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCdp creates a new instance of Cdp, bound to a specific deployed contract.
func NewCdp(address common.Address, backend bind.ContractBackend) (*Cdp, error) {
	contract, err := bindCdp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cdp{CdpCaller: CdpCaller{contract: contract}, CdpTransactor: CdpTransactor{contract: contract}, CdpFilterer: CdpFilterer{contract: contract}}, nil
}

// NewCdpCaller creates a new read-only instance of Cdp, bound to a specific deployed contract.
func NewCdpCaller(address common.Address, caller bind.ContractCaller) (*CdpCaller, error) {
	contract, err := bindCdp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CdpCaller{contract: contract}, nil
}

// NewCdpTransactor creates a new write-only instance of Cdp, bound to a specific deployed contract.
func NewCdpTransactor(address common.Address, transactor bind.ContractTransactor) (*CdpTransactor, error) {
	contract, err := bindCdp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CdpTransactor{contract: contract}, nil
}

// NewCdpFilterer creates a new log filterer instance of Cdp, bound to a specific deployed contract.
func NewCdpFilterer(address common.Address, filterer bind.ContractFilterer) (*CdpFilterer, error) {
	contract, err := bindCdp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CdpFilterer{contract: contract}, nil
}

// bindCdp binds a generic wrapper to an already deployed contract.
func bindCdp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CdpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cdp *CdpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cdp.Contract.CdpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cdp *CdpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdp.Contract.CdpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cdp *CdpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cdp.Contract.CdpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cdp *CdpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cdp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cdp *CdpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cdp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cdp *CdpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cdp.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Cdp *CdpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cdp.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Cdp *CdpSession) Owner() (common.Address, error) {
	return _Cdp.Contract.Owner(&_Cdp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Cdp *CdpCallerSession) Owner() (common.Address, error) {
	return _Cdp.Contract.Owner(&_Cdp.CallOpts)
}

// GetEmisionData is a free data retrieval call binding the contract method 0x34fcddc4.
//
// Solidity: function getEmisionData() view returns((address,string)[])
func (_Cdp *CdpCaller) GetEmisionData(opts *bind.CallOpts) ([]EmissionAppContractEmissionData, error) {
	var out []interface{}
	err := _Cdp.contract.Call(opts, &out, "getEmisionData")

	if err != nil {
		return *new([]EmissionAppContractEmissionData), err
	}

	out0 := *abi.ConvertType(out[0], new([]EmissionAppContractEmissionData)).(*[]EmissionAppContractEmissionData)

	return out0, err

}

// GetEmisionData is a free data retrieval call binding the contract method 0x34fcddc4.
//
// Solidity: function getEmisionData() view returns((address,string)[])
func (_Cdp *CdpSession) GetEmisionData() ([]EmissionAppContractEmissionData, error) {
	return _Cdp.Contract.GetEmisionData(&_Cdp.CallOpts)
}

// GetEmisionData is a free data retrieval call binding the contract method 0x34fcddc4.
//
// Solidity: function getEmisionData() view returns((address,string)[])
func (_Cdp *CdpCallerSession) GetEmisionData() ([]EmissionAppContractEmissionData, error) {
	return _Cdp.Contract.GetEmisionData(&_Cdp.CallOpts)
}

// AddEmissionData is a paid mutator transaction binding the contract method 0x8adbe86f.
//
// Solidity: function addEmissionData(string emissonData) returns()
func (_Cdp *CdpTransactor) AddEmissionData(opts *bind.TransactOpts, emissonData string) (*types.Transaction, error) {
	return _Cdp.contract.Transact(opts, "addEmissionData", emissonData)
}

// AddEmissionData is a paid mutator transaction binding the contract method 0x8adbe86f.
//
// Solidity: function addEmissionData(string emissonData) returns()
func (_Cdp *CdpSession) AddEmissionData(emissonData string) (*types.Transaction, error) {
	return _Cdp.Contract.AddEmissionData(&_Cdp.TransactOpts, emissonData)
}

// AddEmissionData is a paid mutator transaction binding the contract method 0x8adbe86f.
//
// Solidity: function addEmissionData(string emissonData) returns()
func (_Cdp *CdpTransactorSession) AddEmissionData(emissonData string) (*types.Transaction, error) {
	return _Cdp.Contract.AddEmissionData(&_Cdp.TransactOpts, emissonData)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(address user) returns()
func (_Cdp *CdpTransactor) AddToWhiteList(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Cdp.contract.Transact(opts, "addToWhiteList", user)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(address user) returns()
func (_Cdp *CdpSession) AddToWhiteList(user common.Address) (*types.Transaction, error) {
	return _Cdp.Contract.AddToWhiteList(&_Cdp.TransactOpts, user)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(address user) returns()
func (_Cdp *CdpTransactorSession) AddToWhiteList(user common.Address) (*types.Transaction, error) {
	return _Cdp.Contract.AddToWhiteList(&_Cdp.TransactOpts, user)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x01bf6648.
//
// Solidity: function removeFromWhiteList(address user) returns()
func (_Cdp *CdpTransactor) RemoveFromWhiteList(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Cdp.contract.Transact(opts, "removeFromWhiteList", user)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x01bf6648.
//
// Solidity: function removeFromWhiteList(address user) returns()
func (_Cdp *CdpSession) RemoveFromWhiteList(user common.Address) (*types.Transaction, error) {
	return _Cdp.Contract.RemoveFromWhiteList(&_Cdp.TransactOpts, user)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x01bf6648.
//
// Solidity: function removeFromWhiteList(address user) returns()
func (_Cdp *CdpTransactorSession) RemoveFromWhiteList(user common.Address) (*types.Transaction, error) {
	return _Cdp.Contract.RemoveFromWhiteList(&_Cdp.TransactOpts, user)
}

// CdpAddEmissionDataEventIterator is returned from FilterAddEmissionDataEvent and is used to iterate over the raw logs and unpacked data for AddEmissionDataEvent events raised by the Cdp contract.
type CdpAddEmissionDataEventIterator struct {
	Event *CdpAddEmissionDataEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CdpAddEmissionDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdpAddEmissionDataEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CdpAddEmissionDataEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CdpAddEmissionDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdpAddEmissionDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdpAddEmissionDataEvent represents a AddEmissionDataEvent event raised by the Cdp contract.
type CdpAddEmissionDataEvent struct {
	Owner common.Address
	Data  string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddEmissionDataEvent is a free log retrieval operation binding the contract event 0xa83e636d9355bfdc3b54880fdf3d15bb5b0434f7573932407eafe4ffa53024b6.
//
// Solidity: event AddEmissionDataEvent(address owner, string data)
func (_Cdp *CdpFilterer) FilterAddEmissionDataEvent(opts *bind.FilterOpts) (*CdpAddEmissionDataEventIterator, error) {

	logs, sub, err := _Cdp.contract.FilterLogs(opts, "AddEmissionDataEvent")
	if err != nil {
		return nil, err
	}
	return &CdpAddEmissionDataEventIterator{contract: _Cdp.contract, event: "AddEmissionDataEvent", logs: logs, sub: sub}, nil
}

// WatchAddEmissionDataEvent is a free log subscription operation binding the contract event 0xa83e636d9355bfdc3b54880fdf3d15bb5b0434f7573932407eafe4ffa53024b6.
//
// Solidity: event AddEmissionDataEvent(address owner, string data)
func (_Cdp *CdpFilterer) WatchAddEmissionDataEvent(opts *bind.WatchOpts, sink chan<- *CdpAddEmissionDataEvent) (event.Subscription, error) {

	logs, sub, err := _Cdp.contract.WatchLogs(opts, "AddEmissionDataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdpAddEmissionDataEvent)
				if err := _Cdp.contract.UnpackLog(event, "AddEmissionDataEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddEmissionDataEvent is a log parse operation binding the contract event 0xa83e636d9355bfdc3b54880fdf3d15bb5b0434f7573932407eafe4ffa53024b6.
//
// Solidity: event AddEmissionDataEvent(address owner, string data)
func (_Cdp *CdpFilterer) ParseAddEmissionDataEvent(log types.Log) (*CdpAddEmissionDataEvent, error) {
	event := new(CdpAddEmissionDataEvent)
	if err := _Cdp.contract.UnpackLog(event, "AddEmissionDataEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CdpAuthChangeEventIterator is returned from FilterAuthChangeEvent and is used to iterate over the raw logs and unpacked data for AuthChangeEvent events raised by the Cdp contract.
type CdpAuthChangeEventIterator struct {
	Event *CdpAuthChangeEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CdpAuthChangeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CdpAuthChangeEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CdpAuthChangeEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CdpAuthChangeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CdpAuthChangeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CdpAuthChangeEvent represents a AuthChangeEvent event raised by the Cdp contract.
type CdpAuthChangeEvent struct {
	User      common.Address
	TargetVal bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthChangeEvent is a free log retrieval operation binding the contract event 0x6033560938c4cc37896e8b06281b1a56bfba5d8fa6349ce3c0b1267860a665d6.
//
// Solidity: event AuthChangeEvent(address user, bool targetVal)
func (_Cdp *CdpFilterer) FilterAuthChangeEvent(opts *bind.FilterOpts) (*CdpAuthChangeEventIterator, error) {

	logs, sub, err := _Cdp.contract.FilterLogs(opts, "AuthChangeEvent")
	if err != nil {
		return nil, err
	}
	return &CdpAuthChangeEventIterator{contract: _Cdp.contract, event: "AuthChangeEvent", logs: logs, sub: sub}, nil
}

// WatchAuthChangeEvent is a free log subscription operation binding the contract event 0x6033560938c4cc37896e8b06281b1a56bfba5d8fa6349ce3c0b1267860a665d6.
//
// Solidity: event AuthChangeEvent(address user, bool targetVal)
func (_Cdp *CdpFilterer) WatchAuthChangeEvent(opts *bind.WatchOpts, sink chan<- *CdpAuthChangeEvent) (event.Subscription, error) {

	logs, sub, err := _Cdp.contract.WatchLogs(opts, "AuthChangeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CdpAuthChangeEvent)
				if err := _Cdp.contract.UnpackLog(event, "AuthChangeEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthChangeEvent is a log parse operation binding the contract event 0x6033560938c4cc37896e8b06281b1a56bfba5d8fa6349ce3c0b1267860a665d6.
//
// Solidity: event AuthChangeEvent(address user, bool targetVal)
func (_Cdp *CdpFilterer) ParseAuthChangeEvent(log types.Log) (*CdpAuthChangeEvent, error) {
	event := new(CdpAuthChangeEvent)
	if err := _Cdp.contract.UnpackLog(event, "AuthChangeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
