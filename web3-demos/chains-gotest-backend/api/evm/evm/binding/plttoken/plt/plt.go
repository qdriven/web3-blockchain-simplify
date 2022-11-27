// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plt

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bf702f747354a3982e7cd9857da855ad97ca1a5c8222349b267e789e6e4964ac64736f6c634300060c0033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ERC20BurnableABI is the input ABI used to generate the binding from.
const ERC20BurnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initializeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20BurnableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20BurnableFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"79cc6790": "burnFrom(address,uint256)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"1694fc56": "initializeERC20(string,string)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Burnable is an auto generated Go binding around an Ethereum contract.
type ERC20Burnable struct {
	ERC20BurnableCaller     // Read-only binding to the contract
	ERC20BurnableTransactor // Write-only binding to the contract
	ERC20BurnableFilterer   // Log filterer for contract events
}

// ERC20BurnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BurnableSession struct {
	Contract     *ERC20Burnable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BurnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BurnableCallerSession struct {
	Contract *ERC20BurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20BurnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BurnableTransactorSession struct {
	Contract     *ERC20BurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20BurnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BurnableRaw struct {
	Contract *ERC20Burnable // Generic contract binding to access the raw methods on
}

// ERC20BurnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BurnableCallerRaw struct {
	Contract *ERC20BurnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BurnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactorRaw struct {
	Contract *ERC20BurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Burnable creates a new instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20Burnable(address common.Address, backend bind.ContractBackend) (*ERC20Burnable, error) {
	contract, err := bindERC20Burnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Burnable{ERC20BurnableCaller: ERC20BurnableCaller{contract: contract}, ERC20BurnableTransactor: ERC20BurnableTransactor{contract: contract}, ERC20BurnableFilterer: ERC20BurnableFilterer{contract: contract}}, nil
}

// NewERC20BurnableCaller creates a new read-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20BurnableCaller, error) {
	contract, err := bindERC20Burnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableCaller{contract: contract}, nil
}

// NewERC20BurnableTransactor creates a new write-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BurnableTransactor, error) {
	contract, err := bindERC20Burnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransactor{contract: contract}, nil
}

// NewERC20BurnableFilterer creates a new log filterer instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BurnableFilterer, error) {
	contract, err := bindERC20Burnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableFilterer{contract: contract}, nil
}

// bindERC20Burnable binds a generic wrapper to an already deployed contract.
func bindERC20Burnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BurnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.ERC20BurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Burnable.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Burnable.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20Burnable.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCallerSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Burnable.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Burnable.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Burnable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// InitializeERC20 is a paid mutator transaction binding the contract method 0x1694fc56.
//
// Solidity: function initializeERC20(string name, string symbol) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) InitializeERC20(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "initializeERC20", name, symbol)
}

// InitializeERC20 is a paid mutator transaction binding the contract method 0x1694fc56.
//
// Solidity: function initializeERC20(string name, string symbol) returns()
func (_ERC20Burnable *ERC20BurnableSession) InitializeERC20(name string, symbol string) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.InitializeERC20(&_ERC20Burnable.TransactOpts, name, symbol)
}

// InitializeERC20 is a paid mutator transaction binding the contract method 0x1694fc56.
//
// Solidity: function initializeERC20(string name, string symbol) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) InitializeERC20(name string, symbol string) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.InitializeERC20(&_ERC20Burnable.TransactOpts, name, symbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// ERC20BurnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Burnable contract.
type ERC20BurnableApprovalIterator struct {
	Event *ERC20BurnableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableApproval)
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
		it.Event = new(ERC20BurnableApproval)
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
func (it *ERC20BurnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableApproval represents a Approval event raised by the ERC20Burnable contract.
type ERC20BurnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20BurnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableApprovalIterator{contract: _ERC20Burnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20BurnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableApproval)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseApproval(log types.Log) (*ERC20BurnableApproval, error) {
	event := new(ERC20BurnableApproval)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20BurnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Burnable contract.
type ERC20BurnableTransferIterator struct {
	Event *ERC20BurnableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableTransfer)
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
		it.Event = new(ERC20BurnableTransfer)
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
func (it *ERC20BurnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableTransfer represents a Transfer event raised by the ERC20Burnable contract.
type ERC20BurnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BurnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransferIterator{contract: _ERC20Burnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BurnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableTransfer)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) ParseTransfer(log types.Log) (*ERC20BurnableTransfer, error) {
	event := new(ERC20BurnableTransfer)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20InitializableABI is the input ABI used to generate the binding from.
const ERC20InitializableABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initializeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20InitializableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20InitializableFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"1694fc56": "initializeERC20(string,string)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20InitializableBin is the compiled bytecode used for deploying new contracts.
var ERC20InitializableBin = "0x60806040523480156200001157600080fd5b506040516200104738038062001047833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b50604052505050620001b18282620001b960201b60201c565b505062000340565b600054610100900460ff1680620001d55750620001d56200029e565b80620001e4575060005460ff16155b620002215760405162461bcd60e51b815260040180806020018281038252602e81526020018062001019602e913960400191505060405180910390fd5b600054610100900460ff161580156200024d576000805460ff1961ff0019909116610100171660011790555b825162000262906004906020860190620002a4565b50815162000278906005906020850190620002a4565b506006805460ff19166012179055801562000299576000805461ff00191690555b505050565b303b1590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002e757805160ff191683800117855562000317565b8280016001018555821562000317579182015b8281111562000317578251825591602001919060010190620002fa565b506200032592915062000329565b5090565b5b808211156200032557600081556001016200032a565b610cc980620003506000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80633950935111610071578063395093511461031357806370a082311461033f57806395d89b4114610365578063a457c2d71461036d578063a9059cbb14610399578063dd62ed3e146103c5576100b4565b806306fdde03146100b9578063095ea7b3146101365780631694fc561461017657806318160ddd146102a557806323b872dd146102bf578063313ce567146102f5575b600080fd5b6100c16103f3565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100fb5781810151838201526020016100e3565b50505050905090810190601f1680156101285780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101626004803603604081101561014c57600080fd5b506001600160a01b038135169060200135610489565b604080519115158252519081900360200190f35b6102a36004803603604081101561018c57600080fd5b8101906020810181356401000000008111156101a757600080fd5b8201836020820111156101b957600080fd5b803590602001918460018302840111640100000000831117156101db57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561022e57600080fd5b82018360208201111561024057600080fd5b8035906020019184600183028401116401000000008311171561026257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104a6945050505050565b005b6102ad61057f565b60408051918252519081900360200190f35b610162600480360360608110156102d557600080fd5b506001600160a01b03813581169160208101359091169060400135610585565b6102fd61060c565b6040805160ff9092168252519081900360200190f35b6101626004803603604081101561032957600080fd5b506001600160a01b038135169060200135610615565b6102ad6004803603602081101561035557600080fd5b50356001600160a01b0316610663565b6100c161067e565b6101626004803603604081101561038357600080fd5b506001600160a01b0381351690602001356106df565b610162600480360360408110156103af57600080fd5b506001600160a01b038135169060200135610747565b6102ad600480360360408110156103db57600080fd5b506001600160a01b038135811691602001351661075b565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561047f5780601f106104545761010080835404028352916020019161047f565b820191906000526020600020905b81548152906001019060200180831161046257829003601f168201915b5050505050905090565b600061049d610496610786565b848461078a565b50600192915050565b600054610100900460ff16806104bf57506104bf610876565b806104cd575060005460ff16155b6105085760405162461bcd60e51b815260040180806020018281038252602e815260200180610bd0602e913960400191505060405180910390fd5b600054610100900460ff16158015610533576000805460ff1961ff0019909116610100171660011790555b8251610546906004906020860190610ad1565b50815161055a906005906020850190610ad1565b506006805460ff19166012179055801561057a576000805461ff00191690555b505050565b60035490565b600061059284848461087c565b6106028461059e610786565b6105fd85604051806060016040528060288152602001610bfe602891396001600160a01b038a166000908152600260205260408120906105dc610786565b6001600160a01b0316815260208101919091526040016000205491906109d9565b61078a565b5060019392505050565b60065460ff1690565b600061049d610622610786565b846105fd8560026000610633610786565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490610a70565b6001600160a01b031660009081526001602052604090205490565b60058054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561047f5780601f106104545761010080835404028352916020019161047f565b600061049d6106ec610786565b846105fd85604051806060016040528060258152602001610c6f6025913960026000610716610786565b6001600160a01b03908116825260208083019390935260409182016000908120918d168152925290205491906109d9565b600061049d610754610786565b848461087c565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b3390565b6001600160a01b0383166107cf5760405162461bcd60e51b8152600401808060200182810382526024815260200180610c4b6024913960400191505060405180910390fd5b6001600160a01b0382166108145760405162461bcd60e51b8152600401808060200182810382526022815260200180610b886022913960400191505060405180910390fd5b6001600160a01b03808416600081815260026020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b303b1590565b6001600160a01b0383166108c15760405162461bcd60e51b8152600401808060200182810382526025815260200180610c266025913960400191505060405180910390fd5b6001600160a01b0382166109065760405162461bcd60e51b8152600401808060200182810382526023815260200180610b656023913960400191505060405180910390fd5b61091183838361057a565b61094e81604051806060016040528060268152602001610baa602691396001600160a01b03861660009081526001602052604090205491906109d9565b6001600160a01b03808516600090815260016020526040808220939093559084168152205461097d9082610a70565b6001600160a01b0380841660008181526001602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610a685760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a2d578181015183820152602001610a15565b50505050905090810190601f168015610a5a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610aca576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b1257805160ff1916838001178555610b3f565b82800160010185558215610b3f579182015b82811115610b3f578251825591602001919060010190610b24565b50610b4b929150610b4f565b5090565b5b80821115610b4b5760008155600101610b5056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a656445524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220b8d1a31d48ee71377c7d2bc00afe975c50894e4c6df34363445b2dba8bbb355764736f6c634300060c0033496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564"

// DeployERC20Initializable deploys a new Ethereum contract, binding an instance of ERC20Initializable to it.
func DeployERC20Initializable(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *ERC20Initializable, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20InitializableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20InitializableBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Initializable{ERC20InitializableCaller: ERC20InitializableCaller{contract: contract}, ERC20InitializableTransactor: ERC20InitializableTransactor{contract: contract}, ERC20InitializableFilterer: ERC20InitializableFilterer{contract: contract}}, nil
}

// ERC20Initializable is an auto generated Go binding around an Ethereum contract.
type ERC20Initializable struct {
	ERC20InitializableCaller     // Read-only binding to the contract
	ERC20InitializableTransactor // Write-only binding to the contract
	ERC20InitializableFilterer   // Log filterer for contract events
}

// ERC20InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20InitializableSession struct {
	Contract     *ERC20Initializable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ERC20InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20InitializableCallerSession struct {
	Contract *ERC20InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ERC20InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20InitializableTransactorSession struct {
	Contract     *ERC20InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ERC20InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20InitializableRaw struct {
	Contract *ERC20Initializable // Generic contract binding to access the raw methods on
}

// ERC20InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20InitializableCallerRaw struct {
	Contract *ERC20InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20InitializableTransactorRaw struct {
	Contract *ERC20InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Initializable creates a new instance of ERC20Initializable, bound to a specific deployed contract.
func NewERC20Initializable(address common.Address, backend bind.ContractBackend) (*ERC20Initializable, error) {
	contract, err := bindERC20Initializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Initializable{ERC20InitializableCaller: ERC20InitializableCaller{contract: contract}, ERC20InitializableTransactor: ERC20InitializableTransactor{contract: contract}, ERC20InitializableFilterer: ERC20InitializableFilterer{contract: contract}}, nil
}

// NewERC20InitializableCaller creates a new read-only instance of ERC20Initializable, bound to a specific deployed contract.
func NewERC20InitializableCaller(address common.Address, caller bind.ContractCaller) (*ERC20InitializableCaller, error) {
	contract, err := bindERC20Initializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InitializableCaller{contract: contract}, nil
}

// NewERC20InitializableTransactor creates a new write-only instance of ERC20Initializable, bound to a specific deployed contract.
func NewERC20InitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20InitializableTransactor, error) {
	contract, err := bindERC20Initializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InitializableTransactor{contract: contract}, nil
}

// NewERC20InitializableFilterer creates a new log filterer instance of ERC20Initializable, bound to a specific deployed contract.
func NewERC20InitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20InitializableFilterer, error) {
	contract, err := bindERC20Initializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20InitializableFilterer{contract: contract}, nil
}

// bindERC20Initializable binds a generic wrapper to an already deployed contract.
func bindERC20Initializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Initializable *ERC20InitializableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Initializable.Contract.ERC20InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Initializable *ERC20InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.ERC20InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Initializable *ERC20InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.ERC20InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Initializable *ERC20InitializableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Initializable *ERC20InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Initializable *ERC20InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Initializable *ERC20InitializableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Initializable.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Initializable *ERC20InitializableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Initializable.Contract.Allowance(&_ERC20Initializable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Initializable *ERC20InitializableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Initializable.Contract.Allowance(&_ERC20Initializable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Initializable *ERC20InitializableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Initializable.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Initializable *ERC20InitializableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Initializable.Contract.BalanceOf(&_ERC20Initializable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Initializable *ERC20InitializableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Initializable.Contract.BalanceOf(&_ERC20Initializable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Initializable *ERC20InitializableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20Initializable.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Initializable *ERC20InitializableSession) Decimals() (uint8, error) {
	return _ERC20Initializable.Contract.Decimals(&_ERC20Initializable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Initializable *ERC20InitializableCallerSession) Decimals() (uint8, error) {
	return _ERC20Initializable.Contract.Decimals(&_ERC20Initializable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Initializable *ERC20InitializableCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Initializable.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Initializable *ERC20InitializableSession) Name() (string, error) {
	return _ERC20Initializable.Contract.Name(&_ERC20Initializable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Initializable *ERC20InitializableCallerSession) Name() (string, error) {
	return _ERC20Initializable.Contract.Name(&_ERC20Initializable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Initializable *ERC20InitializableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Initializable.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Initializable *ERC20InitializableSession) Symbol() (string, error) {
	return _ERC20Initializable.Contract.Symbol(&_ERC20Initializable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Initializable *ERC20InitializableCallerSession) Symbol() (string, error) {
	return _ERC20Initializable.Contract.Symbol(&_ERC20Initializable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Initializable *ERC20InitializableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Initializable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Initializable *ERC20InitializableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Initializable.Contract.TotalSupply(&_ERC20Initializable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Initializable *ERC20InitializableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Initializable.Contract.TotalSupply(&_ERC20Initializable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Initializable *ERC20InitializableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Initializable *ERC20InitializableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.Approve(&_ERC20Initializable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Initializable *ERC20InitializableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.Approve(&_ERC20Initializable.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Initializable *ERC20InitializableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Initializable *ERC20InitializableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.DecreaseAllowance(&_ERC20Initializable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Initializable *ERC20InitializableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.DecreaseAllowance(&_ERC20Initializable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Initializable *ERC20InitializableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Initializable *ERC20InitializableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.IncreaseAllowance(&_ERC20Initializable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Initializable *ERC20InitializableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.IncreaseAllowance(&_ERC20Initializable.TransactOpts, spender, addedValue)
}

// InitializeERC20 is a paid mutator transaction binding the contract method 0x1694fc56.
//
// Solidity: function initializeERC20(string name, string symbol) returns()
func (_ERC20Initializable *ERC20InitializableTransactor) InitializeERC20(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _ERC20Initializable.contract.Transact(opts, "initializeERC20", name, symbol)
}

// InitializeERC20 is a paid mutator transaction binding the contract method 0x1694fc56.
//
// Solidity: function initializeERC20(string name, string symbol) returns()
func (_ERC20Initializable *ERC20InitializableSession) InitializeERC20(name string, symbol string) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.InitializeERC20(&_ERC20Initializable.TransactOpts, name, symbol)
}

// InitializeERC20 is a paid mutator transaction binding the contract method 0x1694fc56.
//
// Solidity: function initializeERC20(string name, string symbol) returns()
func (_ERC20Initializable *ERC20InitializableTransactorSession) InitializeERC20(name string, symbol string) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.InitializeERC20(&_ERC20Initializable.TransactOpts, name, symbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Initializable *ERC20InitializableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Initializable *ERC20InitializableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.Transfer(&_ERC20Initializable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Initializable *ERC20InitializableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.Transfer(&_ERC20Initializable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Initializable *ERC20InitializableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Initializable *ERC20InitializableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.TransferFrom(&_ERC20Initializable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Initializable *ERC20InitializableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Initializable.Contract.TransferFrom(&_ERC20Initializable.TransactOpts, sender, recipient, amount)
}

// ERC20InitializableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Initializable contract.
type ERC20InitializableApprovalIterator struct {
	Event *ERC20InitializableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20InitializableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20InitializableApproval)
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
		it.Event = new(ERC20InitializableApproval)
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
func (it *ERC20InitializableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20InitializableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20InitializableApproval represents a Approval event raised by the ERC20Initializable contract.
type ERC20InitializableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Initializable *ERC20InitializableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20InitializableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Initializable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20InitializableApprovalIterator{contract: _ERC20Initializable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Initializable *ERC20InitializableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20InitializableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Initializable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20InitializableApproval)
				if err := _ERC20Initializable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Initializable *ERC20InitializableFilterer) ParseApproval(log types.Log) (*ERC20InitializableApproval, error) {
	event := new(ERC20InitializableApproval)
	if err := _ERC20Initializable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20InitializableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Initializable contract.
type ERC20InitializableTransferIterator struct {
	Event *ERC20InitializableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20InitializableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20InitializableTransfer)
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
		it.Event = new(ERC20InitializableTransfer)
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
func (it *ERC20InitializableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20InitializableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20InitializableTransfer represents a Transfer event raised by the ERC20Initializable contract.
type ERC20InitializableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Initializable *ERC20InitializableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20InitializableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Initializable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20InitializableTransferIterator{contract: _ERC20Initializable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Initializable *ERC20InitializableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20InitializableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Initializable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20InitializableTransfer)
				if err := _ERC20Initializable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Initializable *ERC20InitializableFilterer) ParseTransfer(log types.Log) (*ERC20InitializableTransfer, error) {
	event := new(ERC20InitializableTransfer)
	if err := _ERC20Initializable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InitializableABI is the input ABI used to generate the binding from.
const InitializableABI = "[]"

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// OwnableWithAcceptABI is the input ABI used to generate the binding from.
const OwnableWithAcceptABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableWithAcceptFuncSigs maps the 4-byte function signature to its string representation.
var OwnableWithAcceptFuncSigs = map[string]string{
	"79ba5097": "acceptOwnership()",
	"76c8ef4e": "getNextOwner()",
	"5f53837f": "initializeOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableWithAccept is an auto generated Go binding around an Ethereum contract.
type OwnableWithAccept struct {
	OwnableWithAcceptCaller     // Read-only binding to the contract
	OwnableWithAcceptTransactor // Write-only binding to the contract
	OwnableWithAcceptFilterer   // Log filterer for contract events
}

// OwnableWithAcceptCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableWithAcceptCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableWithAcceptTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableWithAcceptTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableWithAcceptFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableWithAcceptFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableWithAcceptSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableWithAcceptSession struct {
	Contract     *OwnableWithAccept // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableWithAcceptCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableWithAcceptCallerSession struct {
	Contract *OwnableWithAcceptCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OwnableWithAcceptTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableWithAcceptTransactorSession struct {
	Contract     *OwnableWithAcceptTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OwnableWithAcceptRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableWithAcceptRaw struct {
	Contract *OwnableWithAccept // Generic contract binding to access the raw methods on
}

// OwnableWithAcceptCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableWithAcceptCallerRaw struct {
	Contract *OwnableWithAcceptCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableWithAcceptTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableWithAcceptTransactorRaw struct {
	Contract *OwnableWithAcceptTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableWithAccept creates a new instance of OwnableWithAccept, bound to a specific deployed contract.
func NewOwnableWithAccept(address common.Address, backend bind.ContractBackend) (*OwnableWithAccept, error) {
	contract, err := bindOwnableWithAccept(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableWithAccept{OwnableWithAcceptCaller: OwnableWithAcceptCaller{contract: contract}, OwnableWithAcceptTransactor: OwnableWithAcceptTransactor{contract: contract}, OwnableWithAcceptFilterer: OwnableWithAcceptFilterer{contract: contract}}, nil
}

// NewOwnableWithAcceptCaller creates a new read-only instance of OwnableWithAccept, bound to a specific deployed contract.
func NewOwnableWithAcceptCaller(address common.Address, caller bind.ContractCaller) (*OwnableWithAcceptCaller, error) {
	contract, err := bindOwnableWithAccept(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableWithAcceptCaller{contract: contract}, nil
}

// NewOwnableWithAcceptTransactor creates a new write-only instance of OwnableWithAccept, bound to a specific deployed contract.
func NewOwnableWithAcceptTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableWithAcceptTransactor, error) {
	contract, err := bindOwnableWithAccept(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableWithAcceptTransactor{contract: contract}, nil
}

// NewOwnableWithAcceptFilterer creates a new log filterer instance of OwnableWithAccept, bound to a specific deployed contract.
func NewOwnableWithAcceptFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableWithAcceptFilterer, error) {
	contract, err := bindOwnableWithAccept(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableWithAcceptFilterer{contract: contract}, nil
}

// bindOwnableWithAccept binds a generic wrapper to an already deployed contract.
func bindOwnableWithAccept(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableWithAcceptABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableWithAccept *OwnableWithAcceptRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OwnableWithAccept.Contract.OwnableWithAcceptCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableWithAccept *OwnableWithAcceptRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.OwnableWithAcceptTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableWithAccept *OwnableWithAcceptRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.OwnableWithAcceptTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableWithAccept *OwnableWithAcceptCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OwnableWithAccept.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableWithAccept *OwnableWithAcceptTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableWithAccept *OwnableWithAcceptTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.contract.Transact(opts, method, params...)
}

// GetNextOwner is a free data retrieval call binding the contract method 0x76c8ef4e.
//
// Solidity: function getNextOwner() view returns(address)
func (_OwnableWithAccept *OwnableWithAcceptCaller) GetNextOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OwnableWithAccept.contract.Call(opts, out, "getNextOwner")
	return *ret0, err
}

// GetNextOwner is a free data retrieval call binding the contract method 0x76c8ef4e.
//
// Solidity: function getNextOwner() view returns(address)
func (_OwnableWithAccept *OwnableWithAcceptSession) GetNextOwner() (common.Address, error) {
	return _OwnableWithAccept.Contract.GetNextOwner(&_OwnableWithAccept.CallOpts)
}

// GetNextOwner is a free data retrieval call binding the contract method 0x76c8ef4e.
//
// Solidity: function getNextOwner() view returns(address)
func (_OwnableWithAccept *OwnableWithAcceptCallerSession) GetNextOwner() (common.Address, error) {
	return _OwnableWithAccept.Contract.GetNextOwner(&_OwnableWithAccept.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableWithAccept *OwnableWithAcceptCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OwnableWithAccept.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableWithAccept *OwnableWithAcceptSession) Owner() (common.Address, error) {
	return _OwnableWithAccept.Contract.Owner(&_OwnableWithAccept.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableWithAccept *OwnableWithAcceptCallerSession) Owner() (common.Address, error) {
	return _OwnableWithAccept.Contract.Owner(&_OwnableWithAccept.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OwnableWithAccept *OwnableWithAcceptTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableWithAccept.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OwnableWithAccept *OwnableWithAcceptSession) AcceptOwnership() (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.AcceptOwnership(&_OwnableWithAccept.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OwnableWithAccept *OwnableWithAcceptTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.AcceptOwnership(&_OwnableWithAccept.TransactOpts)
}

// InitializeOwner is a paid mutator transaction binding the contract method 0x5f53837f.
//
// Solidity: function initializeOwner() returns()
func (_OwnableWithAccept *OwnableWithAcceptTransactor) InitializeOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableWithAccept.contract.Transact(opts, "initializeOwner")
}

// InitializeOwner is a paid mutator transaction binding the contract method 0x5f53837f.
//
// Solidity: function initializeOwner() returns()
func (_OwnableWithAccept *OwnableWithAcceptSession) InitializeOwner() (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.InitializeOwner(&_OwnableWithAccept.TransactOpts)
}

// InitializeOwner is a paid mutator transaction binding the contract method 0x5f53837f.
//
// Solidity: function initializeOwner() returns()
func (_OwnableWithAccept *OwnableWithAcceptTransactorSession) InitializeOwner() (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.InitializeOwner(&_OwnableWithAccept.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableWithAccept *OwnableWithAcceptTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableWithAccept.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableWithAccept *OwnableWithAcceptSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.RenounceOwnership(&_OwnableWithAccept.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableWithAccept *OwnableWithAcceptTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.RenounceOwnership(&_OwnableWithAccept.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableWithAccept *OwnableWithAcceptTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnableWithAccept.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableWithAccept *OwnableWithAcceptSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.TransferOwnership(&_OwnableWithAccept.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableWithAccept *OwnableWithAcceptTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableWithAccept.Contract.TransferOwnership(&_OwnableWithAccept.TransactOpts, newOwner)
}

// OwnableWithAcceptOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnableWithAccept contract.
type OwnableWithAcceptOwnershipTransferredIterator struct {
	Event *OwnableWithAcceptOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableWithAcceptOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableWithAcceptOwnershipTransferred)
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
		it.Event = new(OwnableWithAcceptOwnershipTransferred)
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
func (it *OwnableWithAcceptOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableWithAcceptOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableWithAcceptOwnershipTransferred represents a OwnershipTransferred event raised by the OwnableWithAccept contract.
type OwnableWithAcceptOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableWithAccept *OwnableWithAcceptFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableWithAcceptOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableWithAccept.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableWithAcceptOwnershipTransferredIterator{contract: _OwnableWithAccept.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableWithAccept *OwnableWithAcceptFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableWithAcceptOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableWithAccept.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableWithAcceptOwnershipTransferred)
				if err := _OwnableWithAccept.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableWithAccept *OwnableWithAcceptFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableWithAcceptOwnershipTransferred, error) {
	event := new(OwnableWithAcceptOwnershipTransferred)
	if err := _OwnableWithAccept.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaletteTokenABI is the input ABI used to generate the binding from.
const PaletteTokenABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CancelVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalVote\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limitToPass\",\"type\":\"uint256\"}],\"name\":\"ProposalFail\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"}],\"name\":\"ProposalMake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ProposalPass\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"cancelVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ccID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"durationLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"excuteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"he\",\"type\":\"address\"}],\"name\":\"getHisDepositedID\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"he\",\"type\":\"address\"}],\"name\":\"getHisProposals\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initializeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFreezed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isGoodToDeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isGoodToVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"makeProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"myTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"enumPaletteToken.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiver\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"setCrossChainID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setDurationLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLockProxy\",\"type\":\"address\"}],\"name\":\"setLockProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newReceiver\",\"type\":\"bytes\"}],\"name\":\"setPaletteReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeTable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unFreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voteBox\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voteCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PaletteTokenFuncSigs maps the 4-byte function signature to its string representation.
var PaletteTokenFuncSigs = map[string]string{
	"79ba5097": "acceptOwnership()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"42966c68": "burn(uint256)",
	"79cc6790": "burnFrom(address,uint256)",
	"e711b6e6": "cancelVote(uint256,uint256)",
	"12fdb896": "ccID()",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"e2bbb158": "deposit(uint256,uint256)",
	"0fb5a6b4": "duration()",
	"ba465d83": "durationLimit()",
	"c63f5552": "excuteProposal(uint256)",
	"62a5af3b": "freeze()",
	"e9f053ee": "getHisDepositedID(address)",
	"8da04a94": "getHisProposals(address)",
	"76c8ef4e": "getNextOwner()",
	"39509351": "increaseAllowance(address,uint256)",
	"8129fc1c": "initialize()",
	"1694fc56": "initializeERC20(string,string)",
	"5f53837f": "initializeOwner()",
	"b9469e1a": "isFreezed()",
	"819e9b3a": "isGoodToDeposit(uint256)",
	"c6ce910d": "isGoodToVote(uint256)",
	"9d4dc021": "lockProxy()",
	"bbf9894f": "makeProposal(uint256,string,uint256)",
	"dd4c5162": "myTotalStake(uint256)",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"371fa854": "proposalID()",
	"013cf08b": "proposals(uint256)",
	"f7260d3e": "receiver()",
	"715018a6": "renounceOwnership()",
	"42ca919e": "setCrossChainID(uint64)",
	"f6be71d1": "setDuration(uint256)",
	"62f4f3a1": "setDurationLimit(uint256)",
	"6f2b6ee6": "setLockProxy(address)",
	"c6fd4b68": "setPaletteReceiver(bytes)",
	"d87886b0": "stakeCounter(uint256)",
	"99c862b9": "stakeTable(uint256,address)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"7cf12b90": "unFreeze()",
	"b384abef": "vote(uint256,uint256)",
	"ad0325d0": "voteBox(uint256,address)",
	"9e1f8dd0": "voteCounter(uint256)",
	"2e1a7d4d": "withdraw(uint256)",
}

// PaletteTokenBin is the compiled bytecode used for deploying new contracts.
var PaletteTokenBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600d81526020016c2830b632ba3a32902a37b5b2b760991b8152506040518060400160405280600381526020016214131560ea1b815250620000646200008360201b60201c565b6200007082826200018d565b506200007d905062000272565b620005f6565b600054610100900460ff16806200009f57506200009f620003d2565b80620000ae575060005460ff16155b620000eb5760405162461bcd60e51b815260040180806020018281038252602e815260200180620055be602e913960400191505060405180910390fd5b600054610100900460ff1615801562000117576000805460ff1961ff0019909116610100171660011790555b600062000123620003d8565b6000805462010000600160b01b031916620100006001600160a01b038416908102919091178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35080156200018a576000805461ff00191690555b50565b600054610100900460ff1680620001a95750620001a9620003d2565b80620001b8575060005460ff16155b620001f55760405162461bcd60e51b815260040180806020018281038252602e815260200180620055be602e913960400191505060405180910390fd5b600054610100900460ff1615801562000221576000805460ff1961ff0019909116610100171660011790555b8251620002369060059060208601906200055a565b5081516200024c9060069060208501906200055a565b506007805460ff1916601217905580156200026d576000805461ff00191690555b505050565b600054610100900460ff16806200028e57506200028e620003d2565b806200029d575060005460ff16155b620002da5760405162461bcd60e51b815260040180806020018281038252602e815260200180620055be602e913960400191505060405180910390fd5b600054610100900460ff1615801562000306576000805460ff1961ff0019909116610100171660011790555b6200031062000083565b620003626040518060400160405280600d81526020016c2830b632ba3a32902a37b5b2b760991b8152506040518060400160405280600381526020016214131560ea1b8152506200018d60201b60201c565b6200039262000370620003d8565b6200037a620003dc565b60ff16600a0a633b9aca0002620003e560201b60201c565b600b805460ff60401b19166801000000000000000017905562127500600c55600160085562015180600d5580156200018a576000805461ff001916905550565b303b1590565b3390565b60075460ff1690565b6001600160a01b03821662000441576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6200044f600083836200026d565b6200046b81600454620004f860201b620042501790919060201c565b6004556001600160a01b038216600090815260026020908152604090912054620004a091839062004250620004f8821b17901c565b6001600160a01b03831660008181526002602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b60008282018381101562000553576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200059d57805160ff1916838001178555620005cd565b82800160010185558215620005cd579182015b82811115620005cd578251825591602001919060010190620005b0565b50620005db929150620005df565b5090565b5b80821115620005db5760008155600101620005e0565b614fb880620006066000396000f3fe608060405234801561001057600080fd5b50600436106102f15760003560e01c8063819e9b3a1161019d578063bbf9894f116100e9578063dd62ed3e116100a2578063e9f053ee1161007c578063e9f053ee14610bb5578063f2fde38b14610bdb578063f6be71d114610c01578063f7260d3e14610c1e576102f1565b8063dd62ed3e14610b41578063e2bbb15814610b6f578063e711b6e614610b92576102f1565b8063bbf9894f1461097c578063c63f555214610a29578063c6ce910d14610a46578063c6fd4b6814610a63578063d87886b014610b07578063dd4c516214610b24576102f1565b80639e1f8dd011610156578063ad0325d011610130578063ad0325d01461091d578063b384abef14610949578063b9469e1a1461096c578063ba465d8314610974576102f1565b80639e1f8dd0146108a8578063a457c2d7146108c5578063a9059cbb146108f1576102f1565b8063819e9b3a146107d15780638da04a94146107ee5780638da5cb5b1461086457806395d89b411461086c57806399c862b9146108745780639d4dc021146108a0576102f1565b806342966c681161025c57806370a082311161021557806379ba5097116101ef57806379ba50971461078d57806379cc6790146107955780637cf12b90146107c15780638129fc1c146107c9576102f1565b806370a082311461073b578063715018a61461076157806376c8ef4e14610769576102f1565b806342966c68146106a457806342ca919e146106c15780635f53837f146106e857806362a5af3b146106f057806362f4f3a1146106f85780636f2b6ee614610715576102f1565b806318160ddd116102ae57806318160ddd146105f757806323b872dd146105ff5780632e1a7d4d14610635578063313ce56714610652578063371fa854146106705780633950935114610678576102f1565b8063013cf08b146102f657806306fdde03146103d0578063095ea7b31461044d5780630fb5a6b41461048d57806312fdb896146104a75780631694fc56146104cc575b600080fd5b6103136004803603602081101561030c57600080fd5b5035610c26565b60405180896001600160a01b0316815260200188815260200187815260200186815260200185815260200184815260200183600281111561035057fe5b815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561038e578181015183820152602001610376565b50505050905090810190601f1680156103bb5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b6103d8610d05565b6040805160208082528351818301528351919283929083019185019080838360005b838110156104125781810151838201526020016103fa565b50505050905090810190601f16801561043f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6104796004803603604081101561046357600080fd5b506001600160a01b038135169060200135610d9c565b604080519115158252519081900360200190f35b610495610dba565b60408051918252519081900360200190f35b6104af610dc0565b6040805167ffffffffffffffff9092168252519081900360200190f35b6105f5600480360360408110156104e257600080fd5b810190602081018135600160201b8111156104fc57600080fd5b82018360208201111561050e57600080fd5b803590602001918460018302840111600160201b8311171561052f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561058157600080fd5b82018360208201111561059357600080fd5b803590602001918460018302840111600160201b831117156105b457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610dd0945050505050565b005b610495610ea9565b6104796004803603606081101561061557600080fd5b506001600160a01b03813581169160208101359091169060400135610eaf565b6105f56004803603602081101561064b57600080fd5b5035610f36565b61065a6114f5565b6040805160ff9092168252519081900360200190f35b6104956114fe565b6104796004803603604081101561068e57600080fd5b506001600160a01b038135169060200135611504565b6105f5600480360360208110156106ba57600080fd5b5035611552565b6105f5600480360360208110156106d757600080fd5b503567ffffffffffffffff16611566565b6105f5611630565b6105f561172f565b6105f56004803603602081101561070e57600080fd5b50356117eb565b6105f56004803603602081101561072b57600080fd5b50356001600160a01b031661184e565b6104956004803603602081101561075157600080fd5b50356001600160a01b0316611916565b6105f5611935565b6107716119e4565b604080516001600160a01b039092168252519081900360200190f35b6105f5611a4c565b6105f5600480360360408110156107ab57600080fd5b506001600160a01b038135169060200135611dad565b6105f5611e02565b6105f5611fd1565b610479600480360360208110156107e757600080fd5b503561210e565b6108146004803603602081101561080457600080fd5b50356001600160a01b0316612267565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610850578181015183820152602001610838565b505050509050019250505060405180910390f35b6107716122d3565b6103d86122e8565b6104956004803603604081101561088a57600080fd5b50803590602001356001600160a01b0316612349565b610771612366565b610495600480360360208110156108be57600080fd5b5035612375565b610479600480360360408110156108db57600080fd5b506001600160a01b038135169060200135612387565b6104796004803603604081101561090757600080fd5b506001600160a01b0381351690602001356123ef565b6104956004803603604081101561093357600080fd5b50803590602001356001600160a01b0316612403565b6105f56004803603604081101561095f57600080fd5b5080359060200135612420565b610479612870565b610495612880565b6104956004803603606081101561099257600080fd5b81359190810190604081016020820135600160201b8111156109b357600080fd5b8201836020820111156109c557600080fd5b803590602001918460018302840111600160201b831117156109e657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250612886915050565b6105f560048036036020811015610a3f57600080fd5b5035612c3c565b61047960048036036020811015610a5c57600080fd5b5035613489565b6105f560048036036020811015610a7957600080fd5b810190602081018135600160201b811115610a9357600080fd5b820183602082011115610aa557600080fd5b803590602001918460018302840111600160201b83111715610ac657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506135df945050505050565b61049560048036036020811015610b1d57600080fd5b503561369c565b61049560048036036020811015610b3a57600080fd5b50356136ae565b61049560048036036040811015610b5757600080fd5b506001600160a01b0381358116916020013516613707565b6105f560048036036040811015610b8557600080fd5b5080359060200135613732565b6105f560048036036040811015610ba857600080fd5b5080359060200135613c21565b61081460048036036020811015610bcb57600080fd5b50356001600160a01b0316613fe6565b6105f560048036036020811015610bf157600080fd5b50356001600160a01b0316614050565b6105f560048036036020811015610c1757600080fd5b5035614115565b6103d86141c2565b6012602090815260009182526040918290208054600180830154600280850154600386015460048701546005880154600689015460078a0180548d51601f6000199b831615610100029b909b01909116979097049889018c90048c0287018c01909c528786526001600160a01b039098169a9599939892979196909560ff909216949293830182828015610cfb5780601f10610cd057610100808354040283529160200191610cfb565b820191906000526020600020905b815481529060010190602001808311610cde57829003601f168201915b5050505050905088565b60058054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610d915780601f10610d6657610100808354040283529160200191610d91565b820191906000526020600020905b815481529060010190602001808311610d7457829003601f168201915b505050505090505b90565b6000610db0610da96142aa565b84846142ae565b5060015b92915050565b600c5481565b600b5467ffffffffffffffff1681565b600054610100900460ff1680610de95750610de961439a565b80610df7575060005460ff16155b610e325760405162461bcd60e51b815260040180806020018281038252602e815260200180614d9d602e913960400191505060405180910390fd5b600054610100900460ff16158015610e5d576000805460ff1961ff0019909116610100171660011790555b8251610e709060059060208601906149a4565b508151610e849060069060208501906149a4565b506007805460ff191660121790558015610ea4576000805461ff00191690555b505050565b60045490565b6000610ebc8484846143a0565b610f2c84610ec86142aa565b610f2785604051806060016040528060288152602001614dec602891396001600160a01b038a16600090815260036020526040812090610f066142aa565b6001600160a01b0316815260208101919091526040016000205491906144fd565b6142ae565b5060019392505050565b600b54600160401b900460ff1615610f7f5760405162461bcd60e51b8152600401808060200182810382526033815260200180614d6a6033913960400191505060405180910390fd5b8060085411610fc3576040805162461bcd60e51b815260206004820152601b6024820152600080516020614e8d833981519152604482015290519081900360640190fd5b610fcb614a22565b60008281526012602090815260409182902082516101008101845281546001600160a01b03168152600182015492810192909252600280820154938301939093526003810154606083015260048101546080830152600581015460a083015260068101549192909160c084019160ff9091169081111561104757fe5b600281111561105257fe5b815260078201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156110e05780601f106110b5576101008083540402835291602001916110e0565b820191906000526020600020905b8154815290600101906020018083116110c357829003601f168201915b5050505050815250509050600060028111156110f857fe5b8160c00151600281111561110857fe5b14156111455760405162461bcd60e51b8152600401808060200182810382526039815260200180614b5f6039913960400191505060405180910390fd5b428160a0015110611194576040805162461bcd60e51b815260206004820152601460248201527369742773207374696c6c20696e20766f74696e6760601b604482015290519081900360640190fd5b600082815260106020526040812061120d90826111af6142aa565b6001600160a01b03166001600160a01b0316815260200190815260200160002054600e600086815260200190815260200160002060006111ed6142aa565b6001600160a01b0316815260208101919091526040016000205490614250565b90506000811161124e5760405162461bcd60e51b8152600401808060200182810382526023815260200180614cc36023913960400191505060405180910390fd5b6112603061125a6142aa565b836143a0565b6000838152600e60205260408120906112776142aa565b6001600160a01b03166001600160a01b03168152602001908152602001600020600090556010600084815260200190815260200160002060006112b86142aa565b6001600160a01b03166001600160a01b031681526020019081526020016000206000905560006001601560006112ec6142aa565b6001600160a01b03166001600160a01b0316815260200190815260200160002060008681526020019081526020016000205403905060006001601460006113316142aa565b6001600160a01b03166001600160a01b03168152602001908152602001600020805490500390506000601460006113666142aa565b6001600160a01b03166001600160a01b03168152602001908152602001600020828154811061139157fe5b9060005260206000200154905080601460006113ab6142aa565b6001600160a01b03166001600160a01b0316815260200190815260200160002084815481106113d657fe5b9060005260206000200181905550601460006113f06142aa565b6001600160a01b03166001600160a01b0316815260200190815260200160002080548061141957fe5b60019003818190600052602060002001600090559055826015600061143c6142aa565b6001600160a01b03166001600160a01b031681526020019081526020016000206000838152602001908152602001600020819055506015600061147d6142aa565b6001600160a01b0316815260208082019290925260409081016000908120898252909252812055856114ad6142aa565b6001600160a01b03167ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568866040518082815260200191505060405180910390a3505050505050565b60075460ff1690565b60085481565b6000610db06115116142aa565b84610f2785600360006115226142aa565b6001600160a01b03908116825260208083019390935260409182016000908120918c168152925290205490614250565b61156361155d6142aa565b82614594565b50565b61156e6142aa565b6000546201000090046001600160a01b039081169116146115c4576040805162461bcd60e51b81526020600482018190526024820152600080516020614e14833981519152604482015290519081900360640190fd5b600b54600160401b900460ff1661160c5760405162461bcd60e51b8152600401808060200182810382526034815260200180614ce66034913960400191505060405180910390fd5b600b805467ffffffffffffffff191667ffffffffffffffff92909216919091179055565b600054610100900460ff1680611649575061164961439a565b80611657575060005460ff16155b6116925760405162461bcd60e51b815260040180806020018281038252602e815260200180614d9d602e913960400191505060405180910390fd5b600054610100900460ff161580156116bd576000805460ff1961ff0019909116610100171660011790555b60006116c76142aa565b6000805462010000600160b01b031916620100006001600160a01b038416908102919091178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3508015611563576000805461ff001916905550565b6117376142aa565b6000546201000090046001600160a01b0390811691161461178d576040805162461bcd60e51b81526020600482018190526024820152600080516020614e14833981519152604482015290519081900360640190fd5b600b54600160401b900460ff16156117d65760405162461bcd60e51b8152600401808060200182810382526033815260200180614d6a6033913960400191505060405180910390fd5b600b805460ff60401b1916600160401b179055565b6117f36142aa565b6000546201000090046001600160a01b03908116911614611849576040805162461bcd60e51b81526020600482018190526024820152600080516020614e14833981519152604482015290519081900360640190fd5b600d55565b6118566142aa565b6000546201000090046001600160a01b039081169116146118ac576040805162461bcd60e51b81526020600482018190526024820152600080516020614e14833981519152604482015290519081900360640190fd5b600b54600160401b900460ff166118f45760405162461bcd60e51b8152600401808060200182810382526034815260200180614ce66034913960400191505060405180910390fd5b600980546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0381166000908152600260205260409020545b919050565b61193d6142aa565b6000546201000090046001600160a01b03908116911614611993576040805162461bcd60e51b81526020600482018190526024820152600080516020614e14833981519152604482015290519081900360640190fd5b60008054604051620100009091046001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805462010000600160b01b0319169055565b6001546000906001600160a01b03166119fb6142aa565b6001600160a01b03161480611a2f5750611a136122d3565b6001600160a01b0316611a246142aa565b6001600160a01b0316145b15611a4657506001546001600160a01b0316610d99565b50600090565b6000611a566119e4565b90506001600160a01b038116611aa9576040805162461bcd60e51b815260206004820152601360248201527233b2ba103d32b937903732bc3a1037bbb732b960691b604482015290519081900360640190fd5b806001600160a01b0316611abb6142aa565b6001600160a01b031614611b0f576040805162461bcd60e51b81526020600482015260166024820152753cb7ba9030b932903737ba103732bc3a1037bbb732b960511b604482015290519081900360640190fd5b6000611b196122d3565b90506000611b2682611916565b9050611b338284836143a0565b611b3b614690565b6001600160a01b038216600090815260146020908152604091829020805483518184028101840190945280845260609392830182828015611b9b57602002820191906000526020600020905b815481526020019060010190808311611b87575b505083519394505082159150611cd690505760005b81811015611cb4576000838281518110611bc657fe5b6020908102919091018101516001600160a01b03891660009081526015835260408082208383529093529190912054909150611c40576001600160a01b0387166000818152601460209081526040808320805460018101825581855283852001869055938352925460158252838320858452909152919020555b6000818152600e602090815260408083206001600160a01b038a81168086529184528285208054918d16808752848720929092558686526010855283862083875285528386208054928752848720929092558590559084528390556015825280832093835292905290812055600101611bb0565b506001600160a01b0384166000908152601460205260408120611cd691614a7b565b611cdf84612267565b805190925015611da65760005b8251811015611d84578560126000858481518110611d0657fe5b6020908102919091018101518252818101929092526040908101600090812080546001600160a01b0319166001600160a01b039586161790559289168352601390915290208351849083908110611d5957fe5b6020908102919091018101518254600181810185556000948552929093209092019190915501611cec565b506001600160a01b0384166000908152601360205260408120611da691614a7b565b5050505050565b6000611de482604051806060016040528060248152602001614e6960249139611ddd86611dd86142aa565b613707565b91906144fd565b9050611df883611df26142aa565b836142ae565b610ea48383614594565b611e0a6142aa565b6000546201000090046001600160a01b03908116911614611e60576040805162461bcd60e51b81526020600482018190526024820152600080516020614e14833981519152604482015290519081900360640190fd5b600b54600160401b900460ff16611ea85760405162461bcd60e51b8152600401808060200182810382526034815260200180614ce66034913960400191505060405180910390fd5b6009546001600160a01b0316611efd576040805162461bcd60e51b81526020600482015260156024820152741b1bd8dac81c1c9bde1e481a5cc81b9bdd081cd95d605a1b604482015290519081900360640190fd5b600a5460026000196101006001841615020190911604611f64576040805162461bcd60e51b815260206004820152601b60248201527f70616c65747465206f70657261746f72206973206e6f74207365740000000000604482015290519081900360640190fd5b600b5467ffffffffffffffff16611fc2576040805162461bcd60e51b815260206004820152601860248201527f70616c6574746520636861696e206964206973207a65726f0000000000000000604482015290519081900360640190fd5b600b805460ff60401b19169055565b600054610100900460ff1680611fea5750611fea61439a565b80611ff8575060005460ff16155b6120335760405162461bcd60e51b815260040180806020018281038252602e815260200180614d9d602e913960400191505060405180910390fd5b600054610100900460ff1615801561205e576000805460ff1961ff0019909116610100171660011790555b612066611630565b6120b06040518060400160405280600d81526020016c2830b632ba3a32902a37b5b2b760991b8152506040518060400160405280600381526020016214131560ea1b815250610dd0565b6120d46120bb6142aa565b6120c36114f5565b60ff16600a0a633b9aca0002614770565b600b805460ff60401b1916600160401b17905562127500600c55600160085562015180600d558015611563576000805461ff001916905550565b6000816008541161212157506000611930565b612129614a22565b60008381526012602090815260409182902082516101008101845281546001600160a01b03168152600182015492810192909252600280820154938301939093526003810154606083015260048101546080830152600581015460a083015260068101549192909160c084019160ff909116908111156121a557fe5b60028111156121b057fe5b815260078201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561223e5780601f106122135761010080835404028352916020019161223e565b820191906000526020600020905b81548152906001019060200180831161222157829003601f168201915b5050505050815250509050428160800151118015612260575042816060015111155b9392505050565b6001600160a01b0381166000908152601360209081526040918290208054835181840281018401909452808452606093928301828280156122c757602002820191906000526020600020905b8154815260200190600101908083116122b3575b50505050509050919050565b6000546201000090046001600160a01b031690565b60068054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610d915780601f10610d6657610100808354040283529160200191610d91565b600e60209081526000928352604080842090915290825290205481565b6009546001600160a01b031681565b60116020526000908152604090205481565b6000610db06123946142aa565b84610f2785604051806060016040528060258152602001614f5e60259139600360006123be6142aa565b6001600160a01b03908116825260208083019390935260409182016000908120918d168152925290205491906144fd565b6000610db06123fc6142aa565b84846143a0565b601060209081526000928352604080842090915290825290205481565b600b54600160401b900460ff16156124695760405162461bcd60e51b8152600401808060200182810382526033815260200180614d6a6033913960400191505060405180910390fd5b81600854116124ad576040805162461bcd60e51b815260206004820152601b6024820152600080516020614e8d833981519152604482015290519081900360640190fd5b60008111612502576040805162461bcd60e51b815260206004820152601c60248201527f616d6f756e74206d75737420626967676572207468616e207a65726f00000000604482015290519081900360640190fd5b61250a614a22565b60008381526012602090815260409182902082516101008101845281546001600160a01b03168152600182015492810192909252600280820154938301939093526003810154606083015260048101546080830152600581015460a083015260068101549192909160c084019160ff9091169081111561258657fe5b600281111561259157fe5b815260078201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561261f5780601f106125f45761010080835404028352916020019161261f565b820191906000526020600020905b81548152906001019060200180831161260257829003601f168201915b50505050508152505090508060800151421015612683576040805162461bcd60e51b815260206004820152601e60248201527f746869732070726f706f73616c206973206e6f74207374617274207965740000604482015290519081900360640190fd5b8060a0015142106126c55760405162461bcd60e51b815260040180806020018281038252602d815260200180614c3b602d913960400191505060405180910390fd5b6000838152600e6020526040812083916126dd6142aa565b6001600160a01b03166001600160a01b0316815260200190815260200160002054101561273b5760405162461bcd60e51b8152600401808060200182810382526035815260200180614e346035913960400191505060405180910390fd5b6000838152601060205260408120612758918491906111ed6142aa565b60008481526010602052604081209061276f6142aa565b6001600160a01b03168152602080820192909252604090810160009081209390935585835260119091529020546127a69083614250565b600084815260116020908152604080832093909355600e90529081206127f1918491906127d16142aa565b6001600160a01b0316815260208101919091526040016000205490614862565b6000848152600e60205260408120906128086142aa565b6001600160a01b031681526020810191909152604001600020558261282b6142aa565b6001600160a01b03167fafd3f234c1f8e944129b26b206d98e5752ad3336a4059938b4a3e990e9588530846040518082815260200191505060405180910390a3505050565b600b54600160401b900460ff1681565b600d5481565b600b54600090600160401b900460ff16156128d25760405162461bcd60e51b8152600401808060200182810382526033815260200180614d6a6033913960400191505060405180910390fd5b6080835111156129135760405162461bcd60e51b815260040180806020018281038252602b815260200180614b98602b913960400191505060405180910390fd5b60006129286064612922610ea9565b906148a4565b9050600061293c6129376142aa565b611916565b90508181101561297d5760405162461bcd60e51b815260040180806020018281038252605b815260200180614c68605b913960600191505060405180910390fd5b600086116129bc5760405162461bcd60e51b8152600401808060200182810382526021815260200180614ece6021913960400191505060405180910390fd5b428410156129c8574293505b6129d0614a22565b6020810187905260608101859052600c548086016080830152600202850160a082015260e08101869052600060c0820152612a096142aa565b6001600160a01b0390811682526040808301858152600854600090815260126020908152929020845181546001600160a01b0319169416939093178355908301516001808401919091559051600280840191909155606084015160038401556080840151600484015560a0840151600584015560c08401516006840180548695949293919260ff19909116918490811115612aa057fe5b021790555060e08201518051612ac09160078401916020909101906149a4565b5090505060136000612ad06142aa565b6001600160a01b0316815260208082019290925260400160009081206008548154600181018355918352929091200155612b0a30846123ef565b612b455760405162461bcd60e51b8152600401808060200182810382526023815260200180614d476023913960400191505060405180910390fd5b612b4d6142aa565b6001600160a01b03166008547f3f23ae9839a06634b42ced8775f9f68c2c5409eb8fb37d17d12acdbf750875c7898885608001518660a0015187604001518d6040518087815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015612be7578181015183820152602001612bcf565b50505050905090810190601f168015612c145780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a35050600880546001810190915595945050505050565b600b54600160401b900460ff1615612c855760405162461bcd60e51b8152600401808060200182810382526033815260200180614d6a6033913960400191505060405180910390fd5b8060085411612cc9576040805162461bcd60e51b815260206004820152601b6024820152600080516020614e8d833981519152604482015290519081900360640190fd5b6009546001600160a01b0316612d26576040805162461bcd60e51b815260206004820152601a60248201527f6c6f636b2070726f7879206973207a65726f2061646472657373000000000000604482015290519081900360640190fd5b600b5467ffffffffffffffff16612d6e5760405162461bcd60e51b8152600401808060200182810382526027815260200180614af06027913960400191505060405180910390fd5b600a5460026000196101006001841615020190911604612dd5576040805162461bcd60e51b815260206004820181905260248201527f70616c65747465206f70657261746f72206973207a65726f2061646472657373604482015290519081900360640190fd5b612ddd614a22565b60008281526012602090815260409182902082516101008101845281546001600160a01b03168152600182015492810192909252600280820154938301939093526003810154606083015260048101546080830152600581015460a083015260068101549192909160c084019160ff90911690811115612e5957fe5b6002811115612e6457fe5b815260078201805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152938201939291830182828015612ef25780601f10612ec757610100808354040283529160200191612ef2565b820191906000526020600020905b815481529060010190602001808311612ed557829003601f168201915b5050505050815250509050428160a0015110612f4c576040805162461bcd60e51b815260206004820152601460248201527369742773207374696c6c20696e20766f74696e6760601b604482015290519081900360640190fd5b60008160c001516002811115612f5e57fe5b14612fb0576040805162461bcd60e51b815260206004820152601760248201527f70726f706f73616c206973206e6f742072756e6e696e67000000000000000000604482015290519081900360640190fd5b612fc330826000015183604001516143a0565b6000828152600f6020526040812054612fe4906003906129229060026148e6565b6001019050806011600085815260200190815260200160002054101561306a576000838152601260209081526040808320600601805460ff191660021790556011825291829020548251908152908101839052815185927f4767847d15cbcbbe7c37bc3cb0f5fc3e12381ae471a69e22d84e416afafd5ebd928290030190a25050611563565b6000838152601260209081526040909120600601805460ff19166001179055820151613097903090614770565b602080830151604080519182525185927facc09279e1464d695159ec4179b1ba33676acec7b80f45f633ccd1a45fcd2ab9928290030190a26000838152601160209081526040808320839055600f82528220919091556009549083015161310b9130916001600160a01b03909116906142ae565b600b546040805130602482015267ffffffffffffffff90921660448084019190915281518084039091018152606490920181526020820180516001600160e01b03166309efb30160e31b1781526009549151835160009460609490936001600160a01b0391909116928492909182918083835b6020831061319d5780518252601f19909201916020918201910161317e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146131ff576040519150601f19603f3d011682016040523d82523d6000602084013e613204565b606091505b5090935091508261325c576040805162461bcd60e51b815260206004820152601d60248201527f6661696c656420746f2063616c6c206173736574486173684d61702829000000604482015290519081900360640190fd5b60408251116132b2576040805162461bcd60e51b815260206004820152601760248201527f6e6f2061737365742062696e64656420666f7220504c54000000000000000000604482015290519081900360640190fd5b600b546020860151604051306024820181815267ffffffffffffffff9094166044830181905260848301849052608060648401908152600a80546002600019600183161561010002019091160460a48601819052606097949693959194929160c40190859080156133645780601f1061333957610100808354040283529160200191613364565b820191906000526020600020905b81548152906001019060200180831161334757829003601f168201915b505060408051601f198184030181529181526020820180516001600160e01b03166384a6d05560e01b17815260095491518351939a506001600160a01b039092169850899750909550859450925090508083835b602083106133d75780518252601f1990920191602091820191016133b8565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114613439576040519150601f19603f3d011682016040523d82523d6000602084013e61343e565b606091505b505080945050836134805760405162461bcd60e51b815260040180806020018281038252602c815260200180614bc3602c913960400191505060405180910390fd5b50505050505050565b6000816008541161349c57506000611930565b6134a4614a22565b60008381526012602090815260409182902082516101008101845281546001600160a01b03168152600182015492810192909252600280820154938301939093526003810154606083015260048101546080830152600581015460a083015260068101549192909160c084019160ff9091169081111561352057fe5b600281111561352b57fe5b815260078201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156135b95780601f1061358e576101008083540402835291602001916135b9565b820191906000526020600020905b81548152906001019060200180831161359c57829003601f168201915b5050505050815250509050428160a0015111801561226057506080015142101592915050565b6135e76142aa565b6000546201000090046001600160a01b0390811691161461363d576040805162461bcd60e51b81526020600482018190526024820152600080516020614e14833981519152604482015290519081900360640190fd5b600b54600160401b900460ff166136855760405162461bcd60e51b8152600401808060200182810382526034815260200180614ce66034913960400191505060405180910390fd5b805161369890600a9060208401906149a4565b5050565b600f6020526000908152604090205481565b6000818152601060205260408120610db490826136c96142aa565b6001600160a01b03166001600160a01b0316815260200190815260200160002054600e600085815260200190815260200160002060006111ed6142aa565b6001600160a01b03918216600090815260036020908152604080832093909416825291909152205490565b600b54600160401b900460ff161561377b5760405162461bcd60e51b8152600401808060200182810382526033815260200180614d6a6033913960400191505060405180910390fd5b600081116137d0576040805162461bcd60e51b815260206004820152601c60248201527f616d6f756e74206d75737420626967676572207468616e207a65726f00000000604482015290519081900360640190fd5b8160085411613814576040805162461bcd60e51b815260206004820152601b6024820152600080516020614e8d833981519152604482015290519081900360640190fd5b61381c614a22565b60008381526012602090815260409182902082516101008101845281546001600160a01b03168152600182015492810192909252600280820154938301939093526003810154606083015260048101546080830152600581015460a083015260068101549192909160c084019160ff9091169081111561389857fe5b60028111156138a357fe5b815260078201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156139315780601f1061390657610100808354040283529160200191613931565b820191906000526020600020905b81548152906001019060200180831161391457829003601f168201915b50505050508152505090504281606001511115613995576040805162461bcd60e51b815260206004820152601e60248201527f746869732070726f706f73616c206973206e6f74207374617274207965740000604482015290519081900360640190fd5b428160800151116139d75760405162461bcd60e51b8152600401808060200182810382526026815260200180614c156026913960400191505060405180910390fd5b60006139e46129376142aa565b905082811015613a3b576040805162461bcd60e51b815260206004820152601e60248201527f796f757220504c54206973206e6f7420656e6f75676820746f206c6f636b0000604482015290519081900360640190fd5b613a4530846123ef565b613a805760405162461bcd60e51b8152600401808060200182810382526023815260200180614d476023913960400191505060405180910390fd5b6000848152600e60205260408120613a9d918591906111ed6142aa565b6000858152600e6020526040812090613ab46142aa565b6001600160a01b031681526020808201929092526040908101600090812093909355868352600f909152902054613aeb9084614250565b6000858152600f6020526040812091909155601590613b086142aa565b6001600160a01b0316815260208082019290925260409081016000908120878252909252902054613bd25760146000613b3f6142aa565b6001600160a01b0316815260208082019290925260400160009081208054600181018255908252918120909101859055601490613b7a6142aa565b6001600160a01b03166001600160a01b031681526020019081526020016000208054905060156000613baa6142aa565b6001600160a01b03168152602080820192909252604090810160009081208882529092529020555b83613bdb6142aa565b6001600160a01b03167f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15856040518082815260200191505060405180910390a350505050565b600b54600160401b900460ff1615613c6a5760405162461bcd60e51b8152600401808060200182810382526033815260200180614d6a6033913960400191505060405180910390fd5b8160085411613cae576040805162461bcd60e51b815260206004820152601b6024820152600080516020614e8d833981519152604482015290519081900360640190fd5b613cb6614a22565b60008381526012602090815260409182902082516101008101845281546001600160a01b03168152600182015492810192909252600280820154938301939093526003810154606083015260048101546080830152600581015460a083015260068101549192909160c084019160ff90911690811115613d3257fe5b6002811115613d3d57fe5b815260078201805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152938201939291830182828015613dcb5780601f10613da057610100808354040283529160200191613dcb565b820191906000526020600020905b815481529060010190602001808311613dae57829003601f168201915b50505050508152505090508060800151421015613e195760405162461bcd60e51b8152600401808060200182810382526026815260200180614f386026913960400191505060405180910390fd5b8060a001514210613e5b5760405162461bcd60e51b815260040180806020018281038252602d815260200180614c3b602d913960400191505060405180910390fd5b60008381526010602052604081208391613e736142aa565b6001600160a01b03166001600160a01b03168152602001908152602001600020541015613ed15760405162461bcd60e51b815260040180806020018281038252602d815260200180614d1a602d913960400191505060405180910390fd5b6000838152601060205260408120613eee918491906127d16142aa565b600084815260106020526040812090613f056142aa565b6001600160a01b0316815260208082019290925260409081016000908120939093558583526011909152902054613f3c9083614862565b600084815260116020908152604080832093909355600e9052908120613f67918491906111ed6142aa565b6000848152600e6020526040812090613f7e6142aa565b6001600160a01b0316815260208101919091526040016000205582613fa16142aa565b6001600160a01b03167fae933f1b11e9ae0890a54a205e35f0e68b56abb232cbaf3e16482391d33ae578846040518082815260200191505060405180910390a3505050565b6001600160a01b0381166000908152601460209081526040918290208054835181840281018401909452808452606093928301828280156122c757602002820191906000526020600020908154815260200190600101908083116122b35750505050509050919050565b6140586142aa565b6000546201000090046001600160a01b039081169116146140ae576040805162461bcd60e51b81526020600482018190526024820152600080516020614e14833981519152604482015290519081900360640190fd5b6001600160a01b0381166140f35760405162461bcd60e51b8152600401808060200182810382526026815260200180614b176026913960400191505060405180910390fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b61411d6142aa565b6000546201000090046001600160a01b03908116911614614173576040805162461bcd60e51b81526020600482018190526024820152600080516020614e14833981519152604482015290519081900360640190fd5b600d548110156141bd576040805162461bcd60e51b815260206004820152601060248201526f6174206c65617374206f6e652064617960801b604482015290519081900360640190fd5b600c55565b600a805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156142485780601f1061421d57610100808354040283529160200191614248565b820191906000526020600020905b81548152906001019060200180831161422b57829003601f168201915b505050505081565b600082820183811015612260576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b3390565b6001600160a01b0383166142f35760405162461bcd60e51b8152600401808060200182810382526024815260200180614f146024913960400191505060405180910390fd5b6001600160a01b0382166143385760405162461bcd60e51b8152600401808060200182810382526022815260200180614b3d6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260036020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b303b1590565b6001600160a01b0383166143e55760405162461bcd60e51b8152600401808060200182810382526025815260200180614eef6025913960400191505060405180910390fd5b6001600160a01b03821661442a5760405162461bcd60e51b8152600401808060200182810382526023815260200180614aab6023913960400191505060405180910390fd5b614435838383610ea4565b61447281604051806060016040528060268152602001614bef602691396001600160a01b03861660009081526002602052604090205491906144fd565b6001600160a01b0380851660009081526002602052604080822093909355908416815220546144a19082614250565b6001600160a01b0380841660008181526002602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561458c5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015614551578181015183820152602001614539565b50505050905090810190601f16801561457e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6001600160a01b0382166145d95760405162461bcd60e51b8152600401808060200182810382526021815260200180614ead6021913960400191505060405180910390fd5b6145e582600083610ea4565b61462281604051806060016040528060228152602001614ace602291396001600160a01b03851660009081526002602052604090205491906144fd565b6001600160a01b0383166000908152600260205260409020556004546146489082614862565b6004556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b6001546001600160a01b03166146a46142aa565b6001600160a01b0316146146f8576040805162461bcd60e51b81526020600482015260166024820152753cb7ba9030b932903737ba103732bc3a1037bbb732b960511b604482015290519081900360640190fd5b600154600080546040516001600160a01b039384169362010000909204909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600180546000805462010000600160b01b0319166001600160a01b03831662010000021790556001600160a01b0319169055565b6001600160a01b0382166147cb576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6147d760008383610ea4565b6004546147e49082614250565b6004556001600160a01b03821660009081526002602052604090205461480a9082614250565b6001600160a01b03831660008181526002602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600061226083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506144fd565b600061226083836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061493f565b6000826148f557506000610db4565b8282028284828161490257fe5b04146122605760405162461bcd60e51b8152600401808060200182810382526021815260200180614dcb6021913960400191505060405180910390fd5b6000818361498e5760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315614551578181015183820152602001614539565b50600083858161499a57fe5b0495945050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106149e557805160ff1916838001178555614a12565b82800160010185558215614a12579182015b82811115614a125782518255916020019190600101906149f7565b50614a1e929150614a95565b5090565b60405180610100016040528060006001600160a01b03168152602001600081526020016000815260200160008152602001600081526020016000815260200160006002811115614a6e57fe5b8152602001606081525090565b508054600082559060005260206000209081019061156391905b5b80821115614a1e5760008155600101614a9656fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636563726f737320636861696e204944206f662050616c6574746520436861696e206973207a65726f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f2061646472657373796f752063616e206e6f7420756e6c6f636b20796f7572207374616b6520756e74696c2070726f706f73616c20697320657863757465642e206c656e677468206f66206465736372697074696f6e206d757374206265206c657373207468616e203132386661696c656420746f2063616c6c206c6f636b2829206f66206c6f636b2070726f787920636f6e747261637445524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365746869732070726f706f73616c206973206f7574206f66207374616b65206475726174696f6e746869732070726f706f73616c20697320616c7265616479206f7574206f6620766f7465206475726174696f6e796f75206e65656420746f206c6f636b206f6e652070657263656e7420504c54206f6620746f74616c20737570706c7920746f2063726561746520612070726f706f73616c20776869636820796f7520646f6e277420686176652e796f752068617665206e6f207374616b6520666f7220746869732070726f706f73616c796f752063616e206f6e6c792063616c6c20746869732066756e63207768656e20636f6e747261637420697320667265657a6564796f7520766f746564207374616b65206973206e6f7420656e6f75676820666f72207468697320616d6f756e746661696c656420746f206c6f636b20796f757220504c5420746f20636f6e7472616374796f752063616e206e6f742063616c6c20746869732066756e63207768656e20636f6e747261637420697320667265657a6564496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7745524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63654f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572796f75206c6f636b6564207374616b65206973206e6f7420656e6f75676820746f20766f746520696e207468697320616d6f756e7445524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e6365746869732070726f706f73616c206973206e6f7420657869737421000000000045524332303a206275726e2066726f6d20746865207a65726f20616464726573736d696e74416d6f756e74206d7573742062652067726561746572207468616e203045524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373766f7465206f6620746869732070726f706f73616c206973206e6f742073746172742079657445524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212203b0be9ef37b5c6c62ebbaae61fb6b9d0e0105825e37cc5ec53d2e548742c04df64736f6c634300060c0033496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564"

// DeployPaletteToken deploys a new Ethereum contract, binding an instance of PaletteToken to it.
func DeployPaletteToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PaletteToken, error) {
	parsed, err := abi.JSON(strings.NewReader(PaletteTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PaletteTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PaletteToken{PaletteTokenCaller: PaletteTokenCaller{contract: contract}, PaletteTokenTransactor: PaletteTokenTransactor{contract: contract}, PaletteTokenFilterer: PaletteTokenFilterer{contract: contract}}, nil
}

// PaletteToken is an auto generated Go binding around an Ethereum contract.
type PaletteToken struct {
	PaletteTokenCaller     // Read-only binding to the contract
	PaletteTokenTransactor // Write-only binding to the contract
	PaletteTokenFilterer   // Log filterer for contract events
}

// PaletteTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaletteTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaletteTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaletteTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaletteTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaletteTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaletteTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaletteTokenSession struct {
	Contract     *PaletteToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaletteTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaletteTokenCallerSession struct {
	Contract *PaletteTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PaletteTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaletteTokenTransactorSession struct {
	Contract     *PaletteTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PaletteTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaletteTokenRaw struct {
	Contract *PaletteToken // Generic contract binding to access the raw methods on
}

// PaletteTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaletteTokenCallerRaw struct {
	Contract *PaletteTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PaletteTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaletteTokenTransactorRaw struct {
	Contract *PaletteTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaletteToken creates a new instance of PaletteToken, bound to a specific deployed contract.
func NewPaletteToken(address common.Address, backend bind.ContractBackend) (*PaletteToken, error) {
	contract, err := bindPaletteToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaletteToken{PaletteTokenCaller: PaletteTokenCaller{contract: contract}, PaletteTokenTransactor: PaletteTokenTransactor{contract: contract}, PaletteTokenFilterer: PaletteTokenFilterer{contract: contract}}, nil
}

// NewPaletteTokenCaller creates a new read-only instance of PaletteToken, bound to a specific deployed contract.
func NewPaletteTokenCaller(address common.Address, caller bind.ContractCaller) (*PaletteTokenCaller, error) {
	contract, err := bindPaletteToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenCaller{contract: contract}, nil
}

// NewPaletteTokenTransactor creates a new write-only instance of PaletteToken, bound to a specific deployed contract.
func NewPaletteTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PaletteTokenTransactor, error) {
	contract, err := bindPaletteToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenTransactor{contract: contract}, nil
}

// NewPaletteTokenFilterer creates a new log filterer instance of PaletteToken, bound to a specific deployed contract.
func NewPaletteTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PaletteTokenFilterer, error) {
	contract, err := bindPaletteToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenFilterer{contract: contract}, nil
}

// bindPaletteToken binds a generic wrapper to an already deployed contract.
func bindPaletteToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaletteTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaletteToken *PaletteTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaletteToken.Contract.PaletteTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaletteToken *PaletteTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaletteToken.Contract.PaletteTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaletteToken *PaletteTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaletteToken.Contract.PaletteTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaletteToken *PaletteTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaletteToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaletteToken *PaletteTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaletteToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaletteToken *PaletteTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaletteToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PaletteToken *PaletteTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PaletteToken *PaletteTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PaletteToken.Contract.Allowance(&_PaletteToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PaletteToken *PaletteTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PaletteToken.Contract.Allowance(&_PaletteToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PaletteToken *PaletteTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PaletteToken *PaletteTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PaletteToken.Contract.BalanceOf(&_PaletteToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PaletteToken *PaletteTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PaletteToken.Contract.BalanceOf(&_PaletteToken.CallOpts, account)
}

// CcID is a free data retrieval call binding the contract method 0x12fdb896.
//
// Solidity: function ccID() view returns(uint64)
func (_PaletteToken *PaletteTokenCaller) CcID(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "ccID")
	return *ret0, err
}

// CcID is a free data retrieval call binding the contract method 0x12fdb896.
//
// Solidity: function ccID() view returns(uint64)
func (_PaletteToken *PaletteTokenSession) CcID() (uint64, error) {
	return _PaletteToken.Contract.CcID(&_PaletteToken.CallOpts)
}

// CcID is a free data retrieval call binding the contract method 0x12fdb896.
//
// Solidity: function ccID() view returns(uint64)
func (_PaletteToken *PaletteTokenCallerSession) CcID() (uint64, error) {
	return _PaletteToken.Contract.CcID(&_PaletteToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PaletteToken *PaletteTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PaletteToken *PaletteTokenSession) Decimals() (uint8, error) {
	return _PaletteToken.Contract.Decimals(&_PaletteToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PaletteToken *PaletteTokenCallerSession) Decimals() (uint8, error) {
	return _PaletteToken.Contract.Decimals(&_PaletteToken.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_PaletteToken *PaletteTokenCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "duration")
	return *ret0, err
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_PaletteToken *PaletteTokenSession) Duration() (*big.Int, error) {
	return _PaletteToken.Contract.Duration(&_PaletteToken.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_PaletteToken *PaletteTokenCallerSession) Duration() (*big.Int, error) {
	return _PaletteToken.Contract.Duration(&_PaletteToken.CallOpts)
}

// DurationLimit is a free data retrieval call binding the contract method 0xba465d83.
//
// Solidity: function durationLimit() view returns(uint256)
func (_PaletteToken *PaletteTokenCaller) DurationLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "durationLimit")
	return *ret0, err
}

// DurationLimit is a free data retrieval call binding the contract method 0xba465d83.
//
// Solidity: function durationLimit() view returns(uint256)
func (_PaletteToken *PaletteTokenSession) DurationLimit() (*big.Int, error) {
	return _PaletteToken.Contract.DurationLimit(&_PaletteToken.CallOpts)
}

// DurationLimit is a free data retrieval call binding the contract method 0xba465d83.
//
// Solidity: function durationLimit() view returns(uint256)
func (_PaletteToken *PaletteTokenCallerSession) DurationLimit() (*big.Int, error) {
	return _PaletteToken.Contract.DurationLimit(&_PaletteToken.CallOpts)
}

// GetHisDepositedID is a free data retrieval call binding the contract method 0xe9f053ee.
//
// Solidity: function getHisDepositedID(address he) view returns(uint256[])
func (_PaletteToken *PaletteTokenCaller) GetHisDepositedID(opts *bind.CallOpts, he common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "getHisDepositedID", he)
	return *ret0, err
}

// GetHisDepositedID is a free data retrieval call binding the contract method 0xe9f053ee.
//
// Solidity: function getHisDepositedID(address he) view returns(uint256[])
func (_PaletteToken *PaletteTokenSession) GetHisDepositedID(he common.Address) ([]*big.Int, error) {
	return _PaletteToken.Contract.GetHisDepositedID(&_PaletteToken.CallOpts, he)
}

// GetHisDepositedID is a free data retrieval call binding the contract method 0xe9f053ee.
//
// Solidity: function getHisDepositedID(address he) view returns(uint256[])
func (_PaletteToken *PaletteTokenCallerSession) GetHisDepositedID(he common.Address) ([]*big.Int, error) {
	return _PaletteToken.Contract.GetHisDepositedID(&_PaletteToken.CallOpts, he)
}

// GetHisProposals is a free data retrieval call binding the contract method 0x8da04a94.
//
// Solidity: function getHisProposals(address he) view returns(uint256[])
func (_PaletteToken *PaletteTokenCaller) GetHisProposals(opts *bind.CallOpts, he common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "getHisProposals", he)
	return *ret0, err
}

// GetHisProposals is a free data retrieval call binding the contract method 0x8da04a94.
//
// Solidity: function getHisProposals(address he) view returns(uint256[])
func (_PaletteToken *PaletteTokenSession) GetHisProposals(he common.Address) ([]*big.Int, error) {
	return _PaletteToken.Contract.GetHisProposals(&_PaletteToken.CallOpts, he)
}

// GetHisProposals is a free data retrieval call binding the contract method 0x8da04a94.
//
// Solidity: function getHisProposals(address he) view returns(uint256[])
func (_PaletteToken *PaletteTokenCallerSession) GetHisProposals(he common.Address) ([]*big.Int, error) {
	return _PaletteToken.Contract.GetHisProposals(&_PaletteToken.CallOpts, he)
}

// GetNextOwner is a free data retrieval call binding the contract method 0x76c8ef4e.
//
// Solidity: function getNextOwner() view returns(address)
func (_PaletteToken *PaletteTokenCaller) GetNextOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "getNextOwner")
	return *ret0, err
}

// GetNextOwner is a free data retrieval call binding the contract method 0x76c8ef4e.
//
// Solidity: function getNextOwner() view returns(address)
func (_PaletteToken *PaletteTokenSession) GetNextOwner() (common.Address, error) {
	return _PaletteToken.Contract.GetNextOwner(&_PaletteToken.CallOpts)
}

// GetNextOwner is a free data retrieval call binding the contract method 0x76c8ef4e.
//
// Solidity: function getNextOwner() view returns(address)
func (_PaletteToken *PaletteTokenCallerSession) GetNextOwner() (common.Address, error) {
	return _PaletteToken.Contract.GetNextOwner(&_PaletteToken.CallOpts)
}

// IsFreezed is a free data retrieval call binding the contract method 0xb9469e1a.
//
// Solidity: function isFreezed() view returns(bool)
func (_PaletteToken *PaletteTokenCaller) IsFreezed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "isFreezed")
	return *ret0, err
}

// IsFreezed is a free data retrieval call binding the contract method 0xb9469e1a.
//
// Solidity: function isFreezed() view returns(bool)
func (_PaletteToken *PaletteTokenSession) IsFreezed() (bool, error) {
	return _PaletteToken.Contract.IsFreezed(&_PaletteToken.CallOpts)
}

// IsFreezed is a free data retrieval call binding the contract method 0xb9469e1a.
//
// Solidity: function isFreezed() view returns(bool)
func (_PaletteToken *PaletteTokenCallerSession) IsFreezed() (bool, error) {
	return _PaletteToken.Contract.IsFreezed(&_PaletteToken.CallOpts)
}

// IsGoodToDeposit is a free data retrieval call binding the contract method 0x819e9b3a.
//
// Solidity: function isGoodToDeposit(uint256 id) view returns(bool)
func (_PaletteToken *PaletteTokenCaller) IsGoodToDeposit(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "isGoodToDeposit", id)
	return *ret0, err
}

// IsGoodToDeposit is a free data retrieval call binding the contract method 0x819e9b3a.
//
// Solidity: function isGoodToDeposit(uint256 id) view returns(bool)
func (_PaletteToken *PaletteTokenSession) IsGoodToDeposit(id *big.Int) (bool, error) {
	return _PaletteToken.Contract.IsGoodToDeposit(&_PaletteToken.CallOpts, id)
}

// IsGoodToDeposit is a free data retrieval call binding the contract method 0x819e9b3a.
//
// Solidity: function isGoodToDeposit(uint256 id) view returns(bool)
func (_PaletteToken *PaletteTokenCallerSession) IsGoodToDeposit(id *big.Int) (bool, error) {
	return _PaletteToken.Contract.IsGoodToDeposit(&_PaletteToken.CallOpts, id)
}

// IsGoodToVote is a free data retrieval call binding the contract method 0xc6ce910d.
//
// Solidity: function isGoodToVote(uint256 id) view returns(bool)
func (_PaletteToken *PaletteTokenCaller) IsGoodToVote(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "isGoodToVote", id)
	return *ret0, err
}

// IsGoodToVote is a free data retrieval call binding the contract method 0xc6ce910d.
//
// Solidity: function isGoodToVote(uint256 id) view returns(bool)
func (_PaletteToken *PaletteTokenSession) IsGoodToVote(id *big.Int) (bool, error) {
	return _PaletteToken.Contract.IsGoodToVote(&_PaletteToken.CallOpts, id)
}

// IsGoodToVote is a free data retrieval call binding the contract method 0xc6ce910d.
//
// Solidity: function isGoodToVote(uint256 id) view returns(bool)
func (_PaletteToken *PaletteTokenCallerSession) IsGoodToVote(id *big.Int) (bool, error) {
	return _PaletteToken.Contract.IsGoodToVote(&_PaletteToken.CallOpts, id)
}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_PaletteToken *PaletteTokenCaller) LockProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "lockProxy")
	return *ret0, err
}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_PaletteToken *PaletteTokenSession) LockProxy() (common.Address, error) {
	return _PaletteToken.Contract.LockProxy(&_PaletteToken.CallOpts)
}

// LockProxy is a free data retrieval call binding the contract method 0x9d4dc021.
//
// Solidity: function lockProxy() view returns(address)
func (_PaletteToken *PaletteTokenCallerSession) LockProxy() (common.Address, error) {
	return _PaletteToken.Contract.LockProxy(&_PaletteToken.CallOpts)
}

// MyTotalStake is a free data retrieval call binding the contract method 0xdd4c5162.
//
// Solidity: function myTotalStake(uint256 id) view returns(uint256)
func (_PaletteToken *PaletteTokenCaller) MyTotalStake(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "myTotalStake", id)
	return *ret0, err
}

// MyTotalStake is a free data retrieval call binding the contract method 0xdd4c5162.
//
// Solidity: function myTotalStake(uint256 id) view returns(uint256)
func (_PaletteToken *PaletteTokenSession) MyTotalStake(id *big.Int) (*big.Int, error) {
	return _PaletteToken.Contract.MyTotalStake(&_PaletteToken.CallOpts, id)
}

// MyTotalStake is a free data retrieval call binding the contract method 0xdd4c5162.
//
// Solidity: function myTotalStake(uint256 id) view returns(uint256)
func (_PaletteToken *PaletteTokenCallerSession) MyTotalStake(id *big.Int) (*big.Int, error) {
	return _PaletteToken.Contract.MyTotalStake(&_PaletteToken.CallOpts, id)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PaletteToken *PaletteTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PaletteToken *PaletteTokenSession) Name() (string, error) {
	return _PaletteToken.Contract.Name(&_PaletteToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PaletteToken *PaletteTokenCallerSession) Name() (string, error) {
	return _PaletteToken.Contract.Name(&_PaletteToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaletteToken *PaletteTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaletteToken *PaletteTokenSession) Owner() (common.Address, error) {
	return _PaletteToken.Contract.Owner(&_PaletteToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaletteToken *PaletteTokenCallerSession) Owner() (common.Address, error) {
	return _PaletteToken.Contract.Owner(&_PaletteToken.CallOpts)
}

// ProposalID is a free data retrieval call binding the contract method 0x371fa854.
//
// Solidity: function proposalID() view returns(uint256)
func (_PaletteToken *PaletteTokenCaller) ProposalID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "proposalID")
	return *ret0, err
}

// ProposalID is a free data retrieval call binding the contract method 0x371fa854.
//
// Solidity: function proposalID() view returns(uint256)
func (_PaletteToken *PaletteTokenSession) ProposalID() (*big.Int, error) {
	return _PaletteToken.Contract.ProposalID(&_PaletteToken.CallOpts)
}

// ProposalID is a free data retrieval call binding the contract method 0x371fa854.
//
// Solidity: function proposalID() view returns(uint256)
func (_PaletteToken *PaletteTokenCallerSession) ProposalID() (*big.Int, error) {
	return _PaletteToken.Contract.ProposalID(&_PaletteToken.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address proposer, uint256 mintAmount, uint256 lockedAmount, uint256 startTime, uint256 depositEndTime, uint256 votingEndTime, uint8 status, string desc)
func (_PaletteToken *PaletteTokenCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Proposer       common.Address
	MintAmount     *big.Int
	LockedAmount   *big.Int
	StartTime      *big.Int
	DepositEndTime *big.Int
	VotingEndTime  *big.Int
	Status         uint8
	Desc           string
}, error) {
	ret := new(struct {
		Proposer       common.Address
		MintAmount     *big.Int
		LockedAmount   *big.Int
		StartTime      *big.Int
		DepositEndTime *big.Int
		VotingEndTime  *big.Int
		Status         uint8
		Desc           string
	})
	out := ret
	err := _PaletteToken.contract.Call(opts, out, "proposals", arg0)
	return *ret, err
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address proposer, uint256 mintAmount, uint256 lockedAmount, uint256 startTime, uint256 depositEndTime, uint256 votingEndTime, uint8 status, string desc)
func (_PaletteToken *PaletteTokenSession) Proposals(arg0 *big.Int) (struct {
	Proposer       common.Address
	MintAmount     *big.Int
	LockedAmount   *big.Int
	StartTime      *big.Int
	DepositEndTime *big.Int
	VotingEndTime  *big.Int
	Status         uint8
	Desc           string
}, error) {
	return _PaletteToken.Contract.Proposals(&_PaletteToken.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address proposer, uint256 mintAmount, uint256 lockedAmount, uint256 startTime, uint256 depositEndTime, uint256 votingEndTime, uint8 status, string desc)
func (_PaletteToken *PaletteTokenCallerSession) Proposals(arg0 *big.Int) (struct {
	Proposer       common.Address
	MintAmount     *big.Int
	LockedAmount   *big.Int
	StartTime      *big.Int
	DepositEndTime *big.Int
	VotingEndTime  *big.Int
	Status         uint8
	Desc           string
}, error) {
	return _PaletteToken.Contract.Proposals(&_PaletteToken.CallOpts, arg0)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(bytes)
func (_PaletteToken *PaletteTokenCaller) Receiver(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "receiver")
	return *ret0, err
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(bytes)
func (_PaletteToken *PaletteTokenSession) Receiver() ([]byte, error) {
	return _PaletteToken.Contract.Receiver(&_PaletteToken.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(bytes)
func (_PaletteToken *PaletteTokenCallerSession) Receiver() ([]byte, error) {
	return _PaletteToken.Contract.Receiver(&_PaletteToken.CallOpts)
}

// StakeCounter is a free data retrieval call binding the contract method 0xd87886b0.
//
// Solidity: function stakeCounter(uint256 ) view returns(uint256)
func (_PaletteToken *PaletteTokenCaller) StakeCounter(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "stakeCounter", arg0)
	return *ret0, err
}

// StakeCounter is a free data retrieval call binding the contract method 0xd87886b0.
//
// Solidity: function stakeCounter(uint256 ) view returns(uint256)
func (_PaletteToken *PaletteTokenSession) StakeCounter(arg0 *big.Int) (*big.Int, error) {
	return _PaletteToken.Contract.StakeCounter(&_PaletteToken.CallOpts, arg0)
}

// StakeCounter is a free data retrieval call binding the contract method 0xd87886b0.
//
// Solidity: function stakeCounter(uint256 ) view returns(uint256)
func (_PaletteToken *PaletteTokenCallerSession) StakeCounter(arg0 *big.Int) (*big.Int, error) {
	return _PaletteToken.Contract.StakeCounter(&_PaletteToken.CallOpts, arg0)
}

// StakeTable is a free data retrieval call binding the contract method 0x99c862b9.
//
// Solidity: function stakeTable(uint256 , address ) view returns(uint256)
func (_PaletteToken *PaletteTokenCaller) StakeTable(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "stakeTable", arg0, arg1)
	return *ret0, err
}

// StakeTable is a free data retrieval call binding the contract method 0x99c862b9.
//
// Solidity: function stakeTable(uint256 , address ) view returns(uint256)
func (_PaletteToken *PaletteTokenSession) StakeTable(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _PaletteToken.Contract.StakeTable(&_PaletteToken.CallOpts, arg0, arg1)
}

// StakeTable is a free data retrieval call binding the contract method 0x99c862b9.
//
// Solidity: function stakeTable(uint256 , address ) view returns(uint256)
func (_PaletteToken *PaletteTokenCallerSession) StakeTable(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _PaletteToken.Contract.StakeTable(&_PaletteToken.CallOpts, arg0, arg1)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PaletteToken *PaletteTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PaletteToken *PaletteTokenSession) Symbol() (string, error) {
	return _PaletteToken.Contract.Symbol(&_PaletteToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PaletteToken *PaletteTokenCallerSession) Symbol() (string, error) {
	return _PaletteToken.Contract.Symbol(&_PaletteToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PaletteToken *PaletteTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PaletteToken *PaletteTokenSession) TotalSupply() (*big.Int, error) {
	return _PaletteToken.Contract.TotalSupply(&_PaletteToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PaletteToken *PaletteTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PaletteToken.Contract.TotalSupply(&_PaletteToken.CallOpts)
}

// VoteBox is a free data retrieval call binding the contract method 0xad0325d0.
//
// Solidity: function voteBox(uint256 , address ) view returns(uint256)
func (_PaletteToken *PaletteTokenCaller) VoteBox(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "voteBox", arg0, arg1)
	return *ret0, err
}

// VoteBox is a free data retrieval call binding the contract method 0xad0325d0.
//
// Solidity: function voteBox(uint256 , address ) view returns(uint256)
func (_PaletteToken *PaletteTokenSession) VoteBox(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _PaletteToken.Contract.VoteBox(&_PaletteToken.CallOpts, arg0, arg1)
}

// VoteBox is a free data retrieval call binding the contract method 0xad0325d0.
//
// Solidity: function voteBox(uint256 , address ) view returns(uint256)
func (_PaletteToken *PaletteTokenCallerSession) VoteBox(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _PaletteToken.Contract.VoteBox(&_PaletteToken.CallOpts, arg0, arg1)
}

// VoteCounter is a free data retrieval call binding the contract method 0x9e1f8dd0.
//
// Solidity: function voteCounter(uint256 ) view returns(uint256)
func (_PaletteToken *PaletteTokenCaller) VoteCounter(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaletteToken.contract.Call(opts, out, "voteCounter", arg0)
	return *ret0, err
}

// VoteCounter is a free data retrieval call binding the contract method 0x9e1f8dd0.
//
// Solidity: function voteCounter(uint256 ) view returns(uint256)
func (_PaletteToken *PaletteTokenSession) VoteCounter(arg0 *big.Int) (*big.Int, error) {
	return _PaletteToken.Contract.VoteCounter(&_PaletteToken.CallOpts, arg0)
}

// VoteCounter is a free data retrieval call binding the contract method 0x9e1f8dd0.
//
// Solidity: function voteCounter(uint256 ) view returns(uint256)
func (_PaletteToken *PaletteTokenCallerSession) VoteCounter(arg0 *big.Int) (*big.Int, error) {
	return _PaletteToken.Contract.VoteCounter(&_PaletteToken.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PaletteToken *PaletteTokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PaletteToken *PaletteTokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _PaletteToken.Contract.AcceptOwnership(&_PaletteToken.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PaletteToken *PaletteTokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PaletteToken.Contract.AcceptOwnership(&_PaletteToken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PaletteToken *PaletteTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PaletteToken *PaletteTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Approve(&_PaletteToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PaletteToken *PaletteTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Approve(&_PaletteToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PaletteToken *PaletteTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PaletteToken *PaletteTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Burn(&_PaletteToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PaletteToken *PaletteTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Burn(&_PaletteToken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_PaletteToken *PaletteTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_PaletteToken *PaletteTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.BurnFrom(&_PaletteToken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_PaletteToken *PaletteTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.BurnFrom(&_PaletteToken.TransactOpts, account, amount)
}

// CancelVote is a paid mutator transaction binding the contract method 0xe711b6e6.
//
// Solidity: function cancelVote(uint256 id, uint256 amount) returns()
func (_PaletteToken *PaletteTokenTransactor) CancelVote(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "cancelVote", id, amount)
}

// CancelVote is a paid mutator transaction binding the contract method 0xe711b6e6.
//
// Solidity: function cancelVote(uint256 id, uint256 amount) returns()
func (_PaletteToken *PaletteTokenSession) CancelVote(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.CancelVote(&_PaletteToken.TransactOpts, id, amount)
}

// CancelVote is a paid mutator transaction binding the contract method 0xe711b6e6.
//
// Solidity: function cancelVote(uint256 id, uint256 amount) returns()
func (_PaletteToken *PaletteTokenTransactorSession) CancelVote(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.CancelVote(&_PaletteToken.TransactOpts, id, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PaletteToken *PaletteTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PaletteToken *PaletteTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.DecreaseAllowance(&_PaletteToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PaletteToken *PaletteTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.DecreaseAllowance(&_PaletteToken.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 id, uint256 amount) returns()
func (_PaletteToken *PaletteTokenTransactor) Deposit(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "deposit", id, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 id, uint256 amount) returns()
func (_PaletteToken *PaletteTokenSession) Deposit(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Deposit(&_PaletteToken.TransactOpts, id, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 id, uint256 amount) returns()
func (_PaletteToken *PaletteTokenTransactorSession) Deposit(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Deposit(&_PaletteToken.TransactOpts, id, amount)
}

// ExcuteProposal is a paid mutator transaction binding the contract method 0xc63f5552.
//
// Solidity: function excuteProposal(uint256 id) returns()
func (_PaletteToken *PaletteTokenTransactor) ExcuteProposal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "excuteProposal", id)
}

// ExcuteProposal is a paid mutator transaction binding the contract method 0xc63f5552.
//
// Solidity: function excuteProposal(uint256 id) returns()
func (_PaletteToken *PaletteTokenSession) ExcuteProposal(id *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.ExcuteProposal(&_PaletteToken.TransactOpts, id)
}

// ExcuteProposal is a paid mutator transaction binding the contract method 0xc63f5552.
//
// Solidity: function excuteProposal(uint256 id) returns()
func (_PaletteToken *PaletteTokenTransactorSession) ExcuteProposal(id *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.ExcuteProposal(&_PaletteToken.TransactOpts, id)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_PaletteToken *PaletteTokenTransactor) Freeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "freeze")
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_PaletteToken *PaletteTokenSession) Freeze() (*types.Transaction, error) {
	return _PaletteToken.Contract.Freeze(&_PaletteToken.TransactOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() returns()
func (_PaletteToken *PaletteTokenTransactorSession) Freeze() (*types.Transaction, error) {
	return _PaletteToken.Contract.Freeze(&_PaletteToken.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PaletteToken *PaletteTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PaletteToken *PaletteTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.IncreaseAllowance(&_PaletteToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PaletteToken *PaletteTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.IncreaseAllowance(&_PaletteToken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PaletteToken *PaletteTokenTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PaletteToken *PaletteTokenSession) Initialize() (*types.Transaction, error) {
	return _PaletteToken.Contract.Initialize(&_PaletteToken.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PaletteToken *PaletteTokenTransactorSession) Initialize() (*types.Transaction, error) {
	return _PaletteToken.Contract.Initialize(&_PaletteToken.TransactOpts)
}

// InitializeERC20 is a paid mutator transaction binding the contract method 0x1694fc56.
//
// Solidity: function initializeERC20(string name, string symbol) returns()
func (_PaletteToken *PaletteTokenTransactor) InitializeERC20(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "initializeERC20", name, symbol)
}

// InitializeERC20 is a paid mutator transaction binding the contract method 0x1694fc56.
//
// Solidity: function initializeERC20(string name, string symbol) returns()
func (_PaletteToken *PaletteTokenSession) InitializeERC20(name string, symbol string) (*types.Transaction, error) {
	return _PaletteToken.Contract.InitializeERC20(&_PaletteToken.TransactOpts, name, symbol)
}

// InitializeERC20 is a paid mutator transaction binding the contract method 0x1694fc56.
//
// Solidity: function initializeERC20(string name, string symbol) returns()
func (_PaletteToken *PaletteTokenTransactorSession) InitializeERC20(name string, symbol string) (*types.Transaction, error) {
	return _PaletteToken.Contract.InitializeERC20(&_PaletteToken.TransactOpts, name, symbol)
}

// InitializeOwner is a paid mutator transaction binding the contract method 0x5f53837f.
//
// Solidity: function initializeOwner() returns()
func (_PaletteToken *PaletteTokenTransactor) InitializeOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "initializeOwner")
}

// InitializeOwner is a paid mutator transaction binding the contract method 0x5f53837f.
//
// Solidity: function initializeOwner() returns()
func (_PaletteToken *PaletteTokenSession) InitializeOwner() (*types.Transaction, error) {
	return _PaletteToken.Contract.InitializeOwner(&_PaletteToken.TransactOpts)
}

// InitializeOwner is a paid mutator transaction binding the contract method 0x5f53837f.
//
// Solidity: function initializeOwner() returns()
func (_PaletteToken *PaletteTokenTransactorSession) InitializeOwner() (*types.Transaction, error) {
	return _PaletteToken.Contract.InitializeOwner(&_PaletteToken.TransactOpts)
}

// MakeProposal is a paid mutator transaction binding the contract method 0xbbf9894f.
//
// Solidity: function makeProposal(uint256 mintAmount, string desc, uint256 startTime) returns(uint256)
func (_PaletteToken *PaletteTokenTransactor) MakeProposal(opts *bind.TransactOpts, mintAmount *big.Int, desc string, startTime *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "makeProposal", mintAmount, desc, startTime)
}

// MakeProposal is a paid mutator transaction binding the contract method 0xbbf9894f.
//
// Solidity: function makeProposal(uint256 mintAmount, string desc, uint256 startTime) returns(uint256)
func (_PaletteToken *PaletteTokenSession) MakeProposal(mintAmount *big.Int, desc string, startTime *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.MakeProposal(&_PaletteToken.TransactOpts, mintAmount, desc, startTime)
}

// MakeProposal is a paid mutator transaction binding the contract method 0xbbf9894f.
//
// Solidity: function makeProposal(uint256 mintAmount, string desc, uint256 startTime) returns(uint256)
func (_PaletteToken *PaletteTokenTransactorSession) MakeProposal(mintAmount *big.Int, desc string, startTime *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.MakeProposal(&_PaletteToken.TransactOpts, mintAmount, desc, startTime)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaletteToken *PaletteTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaletteToken *PaletteTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaletteToken.Contract.RenounceOwnership(&_PaletteToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaletteToken *PaletteTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaletteToken.Contract.RenounceOwnership(&_PaletteToken.TransactOpts)
}

// SetCrossChainID is a paid mutator transaction binding the contract method 0x42ca919e.
//
// Solidity: function setCrossChainID(uint64 chainID) returns()
func (_PaletteToken *PaletteTokenTransactor) SetCrossChainID(opts *bind.TransactOpts, chainID uint64) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "setCrossChainID", chainID)
}

// SetCrossChainID is a paid mutator transaction binding the contract method 0x42ca919e.
//
// Solidity: function setCrossChainID(uint64 chainID) returns()
func (_PaletteToken *PaletteTokenSession) SetCrossChainID(chainID uint64) (*types.Transaction, error) {
	return _PaletteToken.Contract.SetCrossChainID(&_PaletteToken.TransactOpts, chainID)
}

// SetCrossChainID is a paid mutator transaction binding the contract method 0x42ca919e.
//
// Solidity: function setCrossChainID(uint64 chainID) returns()
func (_PaletteToken *PaletteTokenTransactorSession) SetCrossChainID(chainID uint64) (*types.Transaction, error) {
	return _PaletteToken.Contract.SetCrossChainID(&_PaletteToken.TransactOpts, chainID)
}

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _duration) returns()
func (_PaletteToken *PaletteTokenTransactor) SetDuration(opts *bind.TransactOpts, _duration *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "setDuration", _duration)
}

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _duration) returns()
func (_PaletteToken *PaletteTokenSession) SetDuration(_duration *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.SetDuration(&_PaletteToken.TransactOpts, _duration)
}

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _duration) returns()
func (_PaletteToken *PaletteTokenTransactorSession) SetDuration(_duration *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.SetDuration(&_PaletteToken.TransactOpts, _duration)
}

// SetDurationLimit is a paid mutator transaction binding the contract method 0x62f4f3a1.
//
// Solidity: function setDurationLimit(uint256 limit) returns()
func (_PaletteToken *PaletteTokenTransactor) SetDurationLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "setDurationLimit", limit)
}

// SetDurationLimit is a paid mutator transaction binding the contract method 0x62f4f3a1.
//
// Solidity: function setDurationLimit(uint256 limit) returns()
func (_PaletteToken *PaletteTokenSession) SetDurationLimit(limit *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.SetDurationLimit(&_PaletteToken.TransactOpts, limit)
}

// SetDurationLimit is a paid mutator transaction binding the contract method 0x62f4f3a1.
//
// Solidity: function setDurationLimit(uint256 limit) returns()
func (_PaletteToken *PaletteTokenTransactorSession) SetDurationLimit(limit *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.SetDurationLimit(&_PaletteToken.TransactOpts, limit)
}

// SetLockProxy is a paid mutator transaction binding the contract method 0x6f2b6ee6.
//
// Solidity: function setLockProxy(address newLockProxy) returns()
func (_PaletteToken *PaletteTokenTransactor) SetLockProxy(opts *bind.TransactOpts, newLockProxy common.Address) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "setLockProxy", newLockProxy)
}

// SetLockProxy is a paid mutator transaction binding the contract method 0x6f2b6ee6.
//
// Solidity: function setLockProxy(address newLockProxy) returns()
func (_PaletteToken *PaletteTokenSession) SetLockProxy(newLockProxy common.Address) (*types.Transaction, error) {
	return _PaletteToken.Contract.SetLockProxy(&_PaletteToken.TransactOpts, newLockProxy)
}

// SetLockProxy is a paid mutator transaction binding the contract method 0x6f2b6ee6.
//
// Solidity: function setLockProxy(address newLockProxy) returns()
func (_PaletteToken *PaletteTokenTransactorSession) SetLockProxy(newLockProxy common.Address) (*types.Transaction, error) {
	return _PaletteToken.Contract.SetLockProxy(&_PaletteToken.TransactOpts, newLockProxy)
}

// SetPaletteReceiver is a paid mutator transaction binding the contract method 0xc6fd4b68.
//
// Solidity: function setPaletteReceiver(bytes newReceiver) returns()
func (_PaletteToken *PaletteTokenTransactor) SetPaletteReceiver(opts *bind.TransactOpts, newReceiver []byte) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "setPaletteReceiver", newReceiver)
}

// SetPaletteReceiver is a paid mutator transaction binding the contract method 0xc6fd4b68.
//
// Solidity: function setPaletteReceiver(bytes newReceiver) returns()
func (_PaletteToken *PaletteTokenSession) SetPaletteReceiver(newReceiver []byte) (*types.Transaction, error) {
	return _PaletteToken.Contract.SetPaletteReceiver(&_PaletteToken.TransactOpts, newReceiver)
}

// SetPaletteReceiver is a paid mutator transaction binding the contract method 0xc6fd4b68.
//
// Solidity: function setPaletteReceiver(bytes newReceiver) returns()
func (_PaletteToken *PaletteTokenTransactorSession) SetPaletteReceiver(newReceiver []byte) (*types.Transaction, error) {
	return _PaletteToken.Contract.SetPaletteReceiver(&_PaletteToken.TransactOpts, newReceiver)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PaletteToken *PaletteTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PaletteToken *PaletteTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Transfer(&_PaletteToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PaletteToken *PaletteTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Transfer(&_PaletteToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PaletteToken *PaletteTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PaletteToken *PaletteTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.TransferFrom(&_PaletteToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PaletteToken *PaletteTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.TransferFrom(&_PaletteToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaletteToken *PaletteTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaletteToken *PaletteTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PaletteToken.Contract.TransferOwnership(&_PaletteToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaletteToken *PaletteTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PaletteToken.Contract.TransferOwnership(&_PaletteToken.TransactOpts, newOwner)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_PaletteToken *PaletteTokenTransactor) UnFreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "unFreeze")
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_PaletteToken *PaletteTokenSession) UnFreeze() (*types.Transaction, error) {
	return _PaletteToken.Contract.UnFreeze(&_PaletteToken.TransactOpts)
}

// UnFreeze is a paid mutator transaction binding the contract method 0x7cf12b90.
//
// Solidity: function unFreeze() returns()
func (_PaletteToken *PaletteTokenTransactorSession) UnFreeze() (*types.Transaction, error) {
	return _PaletteToken.Contract.UnFreeze(&_PaletteToken.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 id, uint256 amount) returns()
func (_PaletteToken *PaletteTokenTransactor) Vote(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "vote", id, amount)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 id, uint256 amount) returns()
func (_PaletteToken *PaletteTokenSession) Vote(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Vote(&_PaletteToken.TransactOpts, id, amount)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 id, uint256 amount) returns()
func (_PaletteToken *PaletteTokenTransactorSession) Vote(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Vote(&_PaletteToken.TransactOpts, id, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 id) returns()
func (_PaletteToken *PaletteTokenTransactor) Withdraw(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _PaletteToken.contract.Transact(opts, "withdraw", id)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 id) returns()
func (_PaletteToken *PaletteTokenSession) Withdraw(id *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Withdraw(&_PaletteToken.TransactOpts, id)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 id) returns()
func (_PaletteToken *PaletteTokenTransactorSession) Withdraw(id *big.Int) (*types.Transaction, error) {
	return _PaletteToken.Contract.Withdraw(&_PaletteToken.TransactOpts, id)
}

// PaletteTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PaletteToken contract.
type PaletteTokenApprovalIterator struct {
	Event *PaletteTokenApproval // Event containing the contract specifics and raw log

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
func (it *PaletteTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaletteTokenApproval)
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
		it.Event = new(PaletteTokenApproval)
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
func (it *PaletteTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaletteTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaletteTokenApproval represents a Approval event raised by the PaletteToken contract.
type PaletteTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PaletteToken *PaletteTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PaletteTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PaletteToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenApprovalIterator{contract: _PaletteToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PaletteToken *PaletteTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PaletteTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PaletteToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaletteTokenApproval)
				if err := _PaletteToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PaletteToken *PaletteTokenFilterer) ParseApproval(log types.Log) (*PaletteTokenApproval, error) {
	event := new(PaletteTokenApproval)
	if err := _PaletteToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaletteTokenCancelVoteIterator is returned from FilterCancelVote and is used to iterate over the raw logs and unpacked data for CancelVote events raised by the PaletteToken contract.
type PaletteTokenCancelVoteIterator struct {
	Event *PaletteTokenCancelVote // Event containing the contract specifics and raw log

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
func (it *PaletteTokenCancelVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaletteTokenCancelVote)
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
		it.Event = new(PaletteTokenCancelVote)
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
func (it *PaletteTokenCancelVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaletteTokenCancelVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaletteTokenCancelVote represents a CancelVote event raised by the PaletteToken contract.
type PaletteTokenCancelVote struct {
	Voter      common.Address
	ProposalID *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCancelVote is a free log retrieval operation binding the contract event 0xae933f1b11e9ae0890a54a205e35f0e68b56abb232cbaf3e16482391d33ae578.
//
// Solidity: event CancelVote(address indexed voter, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) FilterCancelVote(opts *bind.FilterOpts, voter []common.Address, proposalID []*big.Int) (*PaletteTokenCancelVoteIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.FilterLogs(opts, "CancelVote", voterRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenCancelVoteIterator{contract: _PaletteToken.contract, event: "CancelVote", logs: logs, sub: sub}, nil
}

// WatchCancelVote is a free log subscription operation binding the contract event 0xae933f1b11e9ae0890a54a205e35f0e68b56abb232cbaf3e16482391d33ae578.
//
// Solidity: event CancelVote(address indexed voter, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) WatchCancelVote(opts *bind.WatchOpts, sink chan<- *PaletteTokenCancelVote, voter []common.Address, proposalID []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.WatchLogs(opts, "CancelVote", voterRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaletteTokenCancelVote)
				if err := _PaletteToken.contract.UnpackLog(event, "CancelVote", log); err != nil {
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

// ParseCancelVote is a log parse operation binding the contract event 0xae933f1b11e9ae0890a54a205e35f0e68b56abb232cbaf3e16482391d33ae578.
//
// Solidity: event CancelVote(address indexed voter, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) ParseCancelVote(log types.Log) (*PaletteTokenCancelVote, error) {
	event := new(PaletteTokenCancelVote)
	if err := _PaletteToken.contract.UnpackLog(event, "CancelVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaletteTokenDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the PaletteToken contract.
type PaletteTokenDepositIterator struct {
	Event *PaletteTokenDeposit // Event containing the contract specifics and raw log

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
func (it *PaletteTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaletteTokenDeposit)
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
		it.Event = new(PaletteTokenDeposit)
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
func (it *PaletteTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaletteTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaletteTokenDeposit represents a Deposit event raised by the PaletteToken contract.
type PaletteTokenDeposit struct {
	From       common.Address
	ProposalID *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed from, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address, proposalID []*big.Int) (*PaletteTokenDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.FilterLogs(opts, "Deposit", fromRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenDepositIterator{contract: _PaletteToken.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed from, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PaletteTokenDeposit, from []common.Address, proposalID []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.WatchLogs(opts, "Deposit", fromRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaletteTokenDeposit)
				if err := _PaletteToken.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed from, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) ParseDeposit(log types.Log) (*PaletteTokenDeposit, error) {
	event := new(PaletteTokenDeposit)
	if err := _PaletteToken.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaletteTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PaletteToken contract.
type PaletteTokenOwnershipTransferredIterator struct {
	Event *PaletteTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaletteTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaletteTokenOwnershipTransferred)
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
		it.Event = new(PaletteTokenOwnershipTransferred)
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
func (it *PaletteTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaletteTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaletteTokenOwnershipTransferred represents a OwnershipTransferred event raised by the PaletteToken contract.
type PaletteTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PaletteToken *PaletteTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaletteTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaletteToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenOwnershipTransferredIterator{contract: _PaletteToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PaletteToken *PaletteTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaletteTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaletteToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaletteTokenOwnershipTransferred)
				if err := _PaletteToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PaletteToken *PaletteTokenFilterer) ParseOwnershipTransferred(log types.Log) (*PaletteTokenOwnershipTransferred, error) {
	event := new(PaletteTokenOwnershipTransferred)
	if err := _PaletteToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaletteTokenProposalFailIterator is returned from FilterProposalFail and is used to iterate over the raw logs and unpacked data for ProposalFail events raised by the PaletteToken contract.
type PaletteTokenProposalFailIterator struct {
	Event *PaletteTokenProposalFail // Event containing the contract specifics and raw log

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
func (it *PaletteTokenProposalFailIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaletteTokenProposalFail)
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
		it.Event = new(PaletteTokenProposalFail)
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
func (it *PaletteTokenProposalFailIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaletteTokenProposalFailIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaletteTokenProposalFail represents a ProposalFail event raised by the PaletteToken contract.
type PaletteTokenProposalFail struct {
	ProposalID  *big.Int
	TotalVote   *big.Int
	LimitToPass *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalFail is a free log retrieval operation binding the contract event 0x4767847d15cbcbbe7c37bc3cb0f5fc3e12381ae471a69e22d84e416afafd5ebd.
//
// Solidity: event ProposalFail(uint256 indexed proposalID, uint256 totalVote, uint256 limitToPass)
func (_PaletteToken *PaletteTokenFilterer) FilterProposalFail(opts *bind.FilterOpts, proposalID []*big.Int) (*PaletteTokenProposalFailIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.FilterLogs(opts, "ProposalFail", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenProposalFailIterator{contract: _PaletteToken.contract, event: "ProposalFail", logs: logs, sub: sub}, nil
}

// WatchProposalFail is a free log subscription operation binding the contract event 0x4767847d15cbcbbe7c37bc3cb0f5fc3e12381ae471a69e22d84e416afafd5ebd.
//
// Solidity: event ProposalFail(uint256 indexed proposalID, uint256 totalVote, uint256 limitToPass)
func (_PaletteToken *PaletteTokenFilterer) WatchProposalFail(opts *bind.WatchOpts, sink chan<- *PaletteTokenProposalFail, proposalID []*big.Int) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.WatchLogs(opts, "ProposalFail", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaletteTokenProposalFail)
				if err := _PaletteToken.contract.UnpackLog(event, "ProposalFail", log); err != nil {
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

// ParseProposalFail is a log parse operation binding the contract event 0x4767847d15cbcbbe7c37bc3cb0f5fc3e12381ae471a69e22d84e416afafd5ebd.
//
// Solidity: event ProposalFail(uint256 indexed proposalID, uint256 totalVote, uint256 limitToPass)
func (_PaletteToken *PaletteTokenFilterer) ParseProposalFail(log types.Log) (*PaletteTokenProposalFail, error) {
	event := new(PaletteTokenProposalFail)
	if err := _PaletteToken.contract.UnpackLog(event, "ProposalFail", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaletteTokenProposalMakeIterator is returned from FilterProposalMake and is used to iterate over the raw logs and unpacked data for ProposalMake events raised by the PaletteToken contract.
type PaletteTokenProposalMakeIterator struct {
	Event *PaletteTokenProposalMake // Event containing the contract specifics and raw log

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
func (it *PaletteTokenProposalMakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaletteTokenProposalMake)
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
		it.Event = new(PaletteTokenProposalMake)
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
func (it *PaletteTokenProposalMakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaletteTokenProposalMakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaletteTokenProposalMake represents a ProposalMake event raised by the PaletteToken contract.
type PaletteTokenProposalMake struct {
	ProposalID     *big.Int
	Sender         common.Address
	MintAmount     *big.Int
	StartTime      *big.Int
	DepositEndTime *big.Int
	VotingEndTime  *big.Int
	LockedAmount   *big.Int
	Desc           string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProposalMake is a free log retrieval operation binding the contract event 0x3f23ae9839a06634b42ced8775f9f68c2c5409eb8fb37d17d12acdbf750875c7.
//
// Solidity: event ProposalMake(uint256 indexed proposalID, address indexed sender, uint256 mintAmount, uint256 startTime, uint256 depositEndTime, uint256 votingEndTime, uint256 lockedAmount, string desc)
func (_PaletteToken *PaletteTokenFilterer) FilterProposalMake(opts *bind.FilterOpts, proposalID []*big.Int, sender []common.Address) (*PaletteTokenProposalMakeIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PaletteToken.contract.FilterLogs(opts, "ProposalMake", proposalIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenProposalMakeIterator{contract: _PaletteToken.contract, event: "ProposalMake", logs: logs, sub: sub}, nil
}

// WatchProposalMake is a free log subscription operation binding the contract event 0x3f23ae9839a06634b42ced8775f9f68c2c5409eb8fb37d17d12acdbf750875c7.
//
// Solidity: event ProposalMake(uint256 indexed proposalID, address indexed sender, uint256 mintAmount, uint256 startTime, uint256 depositEndTime, uint256 votingEndTime, uint256 lockedAmount, string desc)
func (_PaletteToken *PaletteTokenFilterer) WatchProposalMake(opts *bind.WatchOpts, sink chan<- *PaletteTokenProposalMake, proposalID []*big.Int, sender []common.Address) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PaletteToken.contract.WatchLogs(opts, "ProposalMake", proposalIDRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaletteTokenProposalMake)
				if err := _PaletteToken.contract.UnpackLog(event, "ProposalMake", log); err != nil {
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

// ParseProposalMake is a log parse operation binding the contract event 0x3f23ae9839a06634b42ced8775f9f68c2c5409eb8fb37d17d12acdbf750875c7.
//
// Solidity: event ProposalMake(uint256 indexed proposalID, address indexed sender, uint256 mintAmount, uint256 startTime, uint256 depositEndTime, uint256 votingEndTime, uint256 lockedAmount, string desc)
func (_PaletteToken *PaletteTokenFilterer) ParseProposalMake(log types.Log) (*PaletteTokenProposalMake, error) {
	event := new(PaletteTokenProposalMake)
	if err := _PaletteToken.contract.UnpackLog(event, "ProposalMake", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaletteTokenProposalPassIterator is returned from FilterProposalPass and is used to iterate over the raw logs and unpacked data for ProposalPass events raised by the PaletteToken contract.
type PaletteTokenProposalPassIterator struct {
	Event *PaletteTokenProposalPass // Event containing the contract specifics and raw log

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
func (it *PaletteTokenProposalPassIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaletteTokenProposalPass)
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
		it.Event = new(PaletteTokenProposalPass)
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
func (it *PaletteTokenProposalPassIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaletteTokenProposalPassIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaletteTokenProposalPass represents a ProposalPass event raised by the PaletteToken contract.
type PaletteTokenProposalPass struct {
	ProposalID *big.Int
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalPass is a free log retrieval operation binding the contract event 0xacc09279e1464d695159ec4179b1ba33676acec7b80f45f633ccd1a45fcd2ab9.
//
// Solidity: event ProposalPass(uint256 indexed proposalID, uint256 value)
func (_PaletteToken *PaletteTokenFilterer) FilterProposalPass(opts *bind.FilterOpts, proposalID []*big.Int) (*PaletteTokenProposalPassIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.FilterLogs(opts, "ProposalPass", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenProposalPassIterator{contract: _PaletteToken.contract, event: "ProposalPass", logs: logs, sub: sub}, nil
}

// WatchProposalPass is a free log subscription operation binding the contract event 0xacc09279e1464d695159ec4179b1ba33676acec7b80f45f633ccd1a45fcd2ab9.
//
// Solidity: event ProposalPass(uint256 indexed proposalID, uint256 value)
func (_PaletteToken *PaletteTokenFilterer) WatchProposalPass(opts *bind.WatchOpts, sink chan<- *PaletteTokenProposalPass, proposalID []*big.Int) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.WatchLogs(opts, "ProposalPass", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaletteTokenProposalPass)
				if err := _PaletteToken.contract.UnpackLog(event, "ProposalPass", log); err != nil {
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

// ParseProposalPass is a log parse operation binding the contract event 0xacc09279e1464d695159ec4179b1ba33676acec7b80f45f633ccd1a45fcd2ab9.
//
// Solidity: event ProposalPass(uint256 indexed proposalID, uint256 value)
func (_PaletteToken *PaletteTokenFilterer) ParseProposalPass(log types.Log) (*PaletteTokenProposalPass, error) {
	event := new(PaletteTokenProposalPass)
	if err := _PaletteToken.contract.UnpackLog(event, "ProposalPass", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaletteTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PaletteToken contract.
type PaletteTokenTransferIterator struct {
	Event *PaletteTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PaletteTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaletteTokenTransfer)
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
		it.Event = new(PaletteTokenTransfer)
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
func (it *PaletteTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaletteTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaletteTokenTransfer represents a Transfer event raised by the PaletteToken contract.
type PaletteTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PaletteToken *PaletteTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PaletteTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PaletteToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenTransferIterator{contract: _PaletteToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PaletteToken *PaletteTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PaletteTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PaletteToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaletteTokenTransfer)
				if err := _PaletteToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PaletteToken *PaletteTokenFilterer) ParseTransfer(log types.Log) (*PaletteTokenTransfer, error) {
	event := new(PaletteTokenTransfer)
	if err := _PaletteToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaletteTokenVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the PaletteToken contract.
type PaletteTokenVoteIterator struct {
	Event *PaletteTokenVote // Event containing the contract specifics and raw log

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
func (it *PaletteTokenVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaletteTokenVote)
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
		it.Event = new(PaletteTokenVote)
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
func (it *PaletteTokenVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaletteTokenVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaletteTokenVote represents a Vote event raised by the PaletteToken contract.
type PaletteTokenVote struct {
	Voter      common.Address
	ProposalID *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0xafd3f234c1f8e944129b26b206d98e5752ad3336a4059938b4a3e990e9588530.
//
// Solidity: event Vote(address indexed voter, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) FilterVote(opts *bind.FilterOpts, voter []common.Address, proposalID []*big.Int) (*PaletteTokenVoteIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.FilterLogs(opts, "Vote", voterRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenVoteIterator{contract: _PaletteToken.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0xafd3f234c1f8e944129b26b206d98e5752ad3336a4059938b4a3e990e9588530.
//
// Solidity: event Vote(address indexed voter, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *PaletteTokenVote, voter []common.Address, proposalID []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.WatchLogs(opts, "Vote", voterRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaletteTokenVote)
				if err := _PaletteToken.contract.UnpackLog(event, "Vote", log); err != nil {
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

// ParseVote is a log parse operation binding the contract event 0xafd3f234c1f8e944129b26b206d98e5752ad3336a4059938b4a3e990e9588530.
//
// Solidity: event Vote(address indexed voter, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) ParseVote(log types.Log) (*PaletteTokenVote, error) {
	event := new(PaletteTokenVote)
	if err := _PaletteToken.contract.UnpackLog(event, "Vote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PaletteTokenWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PaletteToken contract.
type PaletteTokenWithdrawIterator struct {
	Event *PaletteTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *PaletteTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaletteTokenWithdraw)
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
		it.Event = new(PaletteTokenWithdraw)
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
func (it *PaletteTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaletteTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaletteTokenWithdraw represents a Withdraw event raised by the PaletteToken contract.
type PaletteTokenWithdraw struct {
	From       common.Address
	ProposalID *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed from, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) FilterWithdraw(opts *bind.FilterOpts, from []common.Address, proposalID []*big.Int) (*PaletteTokenWithdrawIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.FilterLogs(opts, "Withdraw", fromRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &PaletteTokenWithdrawIterator{contract: _PaletteToken.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed from, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PaletteTokenWithdraw, from []common.Address, proposalID []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _PaletteToken.contract.WatchLogs(opts, "Withdraw", fromRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaletteTokenWithdraw)
				if err := _PaletteToken.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed from, uint256 indexed proposalID, uint256 amount)
func (_PaletteToken *PaletteTokenFilterer) ParseWithdraw(log types.Log) (*PaletteTokenWithdraw, error) {
	event := new(PaletteTokenWithdraw)
	if err := _PaletteToken.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200fcdcb8fd74ddbe995a7e5f549c44b178d440fb82018c7d98ad1e66c69829e4964736f6c634300060c0033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
