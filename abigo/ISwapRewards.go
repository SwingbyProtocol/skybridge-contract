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

// ISwapRewardsMetaData contains all meta data concerning the ISwapRewards contract.
var ISwapRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dest\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_swapped\",\"type\":\"uint256\"}],\"name\":\"pullRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dest\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_receiver\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_swapped\",\"type\":\"uint256[]\"}],\"name\":\"pullRewardsMulti\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISwapRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwapRewardsMetaData.ABI instead.
var ISwapRewardsABI = ISwapRewardsMetaData.ABI

// ISwapRewards is an auto generated Go binding around an Ethereum contract.
type ISwapRewards struct {
	ISwapRewardsCaller     // Read-only binding to the contract
	ISwapRewardsTransactor // Write-only binding to the contract
	ISwapRewardsFilterer   // Log filterer for contract events
}

// ISwapRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapRewardsSession struct {
	Contract     *ISwapRewards     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapRewardsCallerSession struct {
	Contract *ISwapRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ISwapRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapRewardsTransactorSession struct {
	Contract     *ISwapRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ISwapRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapRewardsRaw struct {
	Contract *ISwapRewards // Generic contract binding to access the raw methods on
}

// ISwapRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapRewardsCallerRaw struct {
	Contract *ISwapRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// ISwapRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapRewardsTransactorRaw struct {
	Contract *ISwapRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwapRewards creates a new instance of ISwapRewards, bound to a specific deployed contract.
func NewISwapRewards(address common.Address, backend bind.ContractBackend) (*ISwapRewards, error) {
	contract, err := bindISwapRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwapRewards{ISwapRewardsCaller: ISwapRewardsCaller{contract: contract}, ISwapRewardsTransactor: ISwapRewardsTransactor{contract: contract}, ISwapRewardsFilterer: ISwapRewardsFilterer{contract: contract}}, nil
}

// NewISwapRewardsCaller creates a new read-only instance of ISwapRewards, bound to a specific deployed contract.
func NewISwapRewardsCaller(address common.Address, caller bind.ContractCaller) (*ISwapRewardsCaller, error) {
	contract, err := bindISwapRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapRewardsCaller{contract: contract}, nil
}

// NewISwapRewardsTransactor creates a new write-only instance of ISwapRewards, bound to a specific deployed contract.
func NewISwapRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwapRewardsTransactor, error) {
	contract, err := bindISwapRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapRewardsTransactor{contract: contract}, nil
}

// NewISwapRewardsFilterer creates a new log filterer instance of ISwapRewards, bound to a specific deployed contract.
func NewISwapRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwapRewardsFilterer, error) {
	contract, err := bindISwapRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapRewardsFilterer{contract: contract}, nil
}

// bindISwapRewards binds a generic wrapper to an already deployed contract.
func bindISwapRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISwapRewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapRewards *ISwapRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapRewards.Contract.ISwapRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapRewards *ISwapRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapRewards.Contract.ISwapRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapRewards *ISwapRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapRewards.Contract.ISwapRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapRewards *ISwapRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapRewards *ISwapRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapRewards *ISwapRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapRewards.Contract.contract.Transact(opts, method, params...)
}

// PullRewards is a paid mutator transaction binding the contract method 0xff320aae.
//
// Solidity: function pullRewards(address _dest, address _receiver, uint256 _swapped) returns(bool)
func (_ISwapRewards *ISwapRewardsTransactor) PullRewards(opts *bind.TransactOpts, _dest common.Address, _receiver common.Address, _swapped *big.Int) (*types.Transaction, error) {
	return _ISwapRewards.contract.Transact(opts, "pullRewards", _dest, _receiver, _swapped)
}

// PullRewards is a paid mutator transaction binding the contract method 0xff320aae.
//
// Solidity: function pullRewards(address _dest, address _receiver, uint256 _swapped) returns(bool)
func (_ISwapRewards *ISwapRewardsSession) PullRewards(_dest common.Address, _receiver common.Address, _swapped *big.Int) (*types.Transaction, error) {
	return _ISwapRewards.Contract.PullRewards(&_ISwapRewards.TransactOpts, _dest, _receiver, _swapped)
}

// PullRewards is a paid mutator transaction binding the contract method 0xff320aae.
//
// Solidity: function pullRewards(address _dest, address _receiver, uint256 _swapped) returns(bool)
func (_ISwapRewards *ISwapRewardsTransactorSession) PullRewards(_dest common.Address, _receiver common.Address, _swapped *big.Int) (*types.Transaction, error) {
	return _ISwapRewards.Contract.PullRewards(&_ISwapRewards.TransactOpts, _dest, _receiver, _swapped)
}

// PullRewardsMulti is a paid mutator transaction binding the contract method 0x49e031a5.
//
// Solidity: function pullRewardsMulti(address _dest, address[] _receiver, uint256[] _swapped) returns(bool)
func (_ISwapRewards *ISwapRewardsTransactor) PullRewardsMulti(opts *bind.TransactOpts, _dest common.Address, _receiver []common.Address, _swapped []*big.Int) (*types.Transaction, error) {
	return _ISwapRewards.contract.Transact(opts, "pullRewardsMulti", _dest, _receiver, _swapped)
}

// PullRewardsMulti is a paid mutator transaction binding the contract method 0x49e031a5.
//
// Solidity: function pullRewardsMulti(address _dest, address[] _receiver, uint256[] _swapped) returns(bool)
func (_ISwapRewards *ISwapRewardsSession) PullRewardsMulti(_dest common.Address, _receiver []common.Address, _swapped []*big.Int) (*types.Transaction, error) {
	return _ISwapRewards.Contract.PullRewardsMulti(&_ISwapRewards.TransactOpts, _dest, _receiver, _swapped)
}

// PullRewardsMulti is a paid mutator transaction binding the contract method 0x49e031a5.
//
// Solidity: function pullRewardsMulti(address _dest, address[] _receiver, uint256[] _swapped) returns(bool)
func (_ISwapRewards *ISwapRewardsTransactorSession) PullRewardsMulti(_dest common.Address, _receiver []common.Address, _swapped []*big.Int) (*types.Transaction, error) {
	return _ISwapRewards.Contract.PullRewardsMulti(&_ISwapRewards.TransactOpts, _dest, _receiver, _swapped)
}
