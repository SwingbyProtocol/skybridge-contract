// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigo

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

// ITokenTransferProxyMetaData contains all meta data concerning the ITokenTransferProxy contract.
var ITokenTransferProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokensToFree\",\"type\":\"uint256\"}],\"name\":\"freeReduxTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITokenTransferProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenTransferProxyMetaData.ABI instead.
var ITokenTransferProxyABI = ITokenTransferProxyMetaData.ABI

// ITokenTransferProxy is an auto generated Go binding around an Ethereum contract.
type ITokenTransferProxy struct {
	ITokenTransferProxyCaller     // Read-only binding to the contract
	ITokenTransferProxyTransactor // Write-only binding to the contract
	ITokenTransferProxyFilterer   // Log filterer for contract events
}

// ITokenTransferProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenTransferProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenTransferProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenTransferProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenTransferProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenTransferProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenTransferProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenTransferProxySession struct {
	Contract     *ITokenTransferProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ITokenTransferProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenTransferProxyCallerSession struct {
	Contract *ITokenTransferProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ITokenTransferProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenTransferProxyTransactorSession struct {
	Contract     *ITokenTransferProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ITokenTransferProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenTransferProxyRaw struct {
	Contract *ITokenTransferProxy // Generic contract binding to access the raw methods on
}

// ITokenTransferProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenTransferProxyCallerRaw struct {
	Contract *ITokenTransferProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenTransferProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenTransferProxyTransactorRaw struct {
	Contract *ITokenTransferProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenTransferProxy creates a new instance of ITokenTransferProxy, bound to a specific deployed contract.
func NewITokenTransferProxy(address common.Address, backend bind.ContractBackend) (*ITokenTransferProxy, error) {
	contract, err := bindITokenTransferProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenTransferProxy{ITokenTransferProxyCaller: ITokenTransferProxyCaller{contract: contract}, ITokenTransferProxyTransactor: ITokenTransferProxyTransactor{contract: contract}, ITokenTransferProxyFilterer: ITokenTransferProxyFilterer{contract: contract}}, nil
}

// NewITokenTransferProxyCaller creates a new read-only instance of ITokenTransferProxy, bound to a specific deployed contract.
func NewITokenTransferProxyCaller(address common.Address, caller bind.ContractCaller) (*ITokenTransferProxyCaller, error) {
	contract, err := bindITokenTransferProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenTransferProxyCaller{contract: contract}, nil
}

// NewITokenTransferProxyTransactor creates a new write-only instance of ITokenTransferProxy, bound to a specific deployed contract.
func NewITokenTransferProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenTransferProxyTransactor, error) {
	contract, err := bindITokenTransferProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenTransferProxyTransactor{contract: contract}, nil
}

// NewITokenTransferProxyFilterer creates a new log filterer instance of ITokenTransferProxy, bound to a specific deployed contract.
func NewITokenTransferProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenTransferProxyFilterer, error) {
	contract, err := bindITokenTransferProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenTransferProxyFilterer{contract: contract}, nil
}

// bindITokenTransferProxy binds a generic wrapper to an already deployed contract.
func bindITokenTransferProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenTransferProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenTransferProxy *ITokenTransferProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenTransferProxy.Contract.ITokenTransferProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenTransferProxy *ITokenTransferProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenTransferProxy.Contract.ITokenTransferProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenTransferProxy *ITokenTransferProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenTransferProxy.Contract.ITokenTransferProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenTransferProxy *ITokenTransferProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenTransferProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenTransferProxy *ITokenTransferProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenTransferProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenTransferProxy *ITokenTransferProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenTransferProxy.Contract.contract.Transact(opts, method, params...)
}

// FreeReduxTokens is a paid mutator transaction binding the contract method 0xdd1fe62c.
//
// Solidity: function freeReduxTokens(address user, uint256 tokensToFree) returns()
func (_ITokenTransferProxy *ITokenTransferProxyTransactor) FreeReduxTokens(opts *bind.TransactOpts, user common.Address, tokensToFree *big.Int) (*types.Transaction, error) {
	return _ITokenTransferProxy.contract.Transact(opts, "freeReduxTokens", user, tokensToFree)
}

// FreeReduxTokens is a paid mutator transaction binding the contract method 0xdd1fe62c.
//
// Solidity: function freeReduxTokens(address user, uint256 tokensToFree) returns()
func (_ITokenTransferProxy *ITokenTransferProxySession) FreeReduxTokens(user common.Address, tokensToFree *big.Int) (*types.Transaction, error) {
	return _ITokenTransferProxy.Contract.FreeReduxTokens(&_ITokenTransferProxy.TransactOpts, user, tokensToFree)
}

// FreeReduxTokens is a paid mutator transaction binding the contract method 0xdd1fe62c.
//
// Solidity: function freeReduxTokens(address user, uint256 tokensToFree) returns()
func (_ITokenTransferProxy *ITokenTransferProxyTransactorSession) FreeReduxTokens(user common.Address, tokensToFree *big.Int) (*types.Transaction, error) {
	return _ITokenTransferProxy.Contract.FreeReduxTokens(&_ITokenTransferProxy.TransactOpts, user, tokensToFree)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(address token, address from, address to, uint256 amount) returns()
func (_ITokenTransferProxy *ITokenTransferProxyTransactor) TransferFrom(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenTransferProxy.contract.Transact(opts, "transferFrom", token, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(address token, address from, address to, uint256 amount) returns()
func (_ITokenTransferProxy *ITokenTransferProxySession) TransferFrom(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenTransferProxy.Contract.TransferFrom(&_ITokenTransferProxy.TransactOpts, token, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(address token, address from, address to, uint256 amount) returns()
func (_ITokenTransferProxy *ITokenTransferProxyTransactorSession) TransferFrom(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenTransferProxy.Contract.TransferFrom(&_ITokenTransferProxy.TransactOpts, token, from, to, amount)
}
