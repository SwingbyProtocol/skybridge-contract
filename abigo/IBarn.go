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

// IBarnMetaData contains all meta data concerning the IBarn contract.
var IBarnMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBarnABI is the input ABI used to generate the binding from.
// Deprecated: Use IBarnMetaData.ABI instead.
var IBarnABI = IBarnMetaData.ABI

// IBarn is an auto generated Go binding around an Ethereum contract.
type IBarn struct {
	IBarnCaller     // Read-only binding to the contract
	IBarnTransactor // Write-only binding to the contract
	IBarnFilterer   // Log filterer for contract events
}

// IBarnCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBarnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBarnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBarnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBarnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBarnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBarnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBarnSession struct {
	Contract     *IBarn            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBarnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBarnCallerSession struct {
	Contract *IBarnCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IBarnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBarnTransactorSession struct {
	Contract     *IBarnTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBarnRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBarnRaw struct {
	Contract *IBarn // Generic contract binding to access the raw methods on
}

// IBarnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBarnCallerRaw struct {
	Contract *IBarnCaller // Generic read-only contract binding to access the raw methods on
}

// IBarnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBarnTransactorRaw struct {
	Contract *IBarnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBarn creates a new instance of IBarn, bound to a specific deployed contract.
func NewIBarn(address common.Address, backend bind.ContractBackend) (*IBarn, error) {
	contract, err := bindIBarn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBarn{IBarnCaller: IBarnCaller{contract: contract}, IBarnTransactor: IBarnTransactor{contract: contract}, IBarnFilterer: IBarnFilterer{contract: contract}}, nil
}

// NewIBarnCaller creates a new read-only instance of IBarn, bound to a specific deployed contract.
func NewIBarnCaller(address common.Address, caller bind.ContractCaller) (*IBarnCaller, error) {
	contract, err := bindIBarn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBarnCaller{contract: contract}, nil
}

// NewIBarnTransactor creates a new write-only instance of IBarn, bound to a specific deployed contract.
func NewIBarnTransactor(address common.Address, transactor bind.ContractTransactor) (*IBarnTransactor, error) {
	contract, err := bindIBarn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBarnTransactor{contract: contract}, nil
}

// NewIBarnFilterer creates a new log filterer instance of IBarn, bound to a specific deployed contract.
func NewIBarnFilterer(address common.Address, filterer bind.ContractFilterer) (*IBarnFilterer, error) {
	contract, err := bindIBarn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBarnFilterer{contract: contract}, nil
}

// bindIBarn binds a generic wrapper to an already deployed contract.
func bindIBarn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBarnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBarn *IBarnRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBarn.Contract.IBarnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBarn *IBarnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBarn.Contract.IBarnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBarn *IBarnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBarn.Contract.IBarnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBarn *IBarnCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBarn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBarn *IBarnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBarn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBarn *IBarnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBarn.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_IBarn *IBarnCaller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_IBarn *IBarnSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _IBarn.Contract.BalanceOf(&_IBarn.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_IBarn *IBarnCallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _IBarn.Contract.BalanceOf(&_IBarn.CallOpts, user)
}
