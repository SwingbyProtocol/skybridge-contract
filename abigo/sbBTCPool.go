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

// SbBTCPoolMetaData contains all meta data concerning the SbBTCPool contract.
var SbBTCPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceBefore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"barn\",\"outputs\":[{\"internalType\":\"contractIBarn\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"resetUnstakedNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_barn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swap\",\"type\":\"address\"}],\"name\":\"setBarnAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapContract\",\"outputs\":[{\"internalType\":\"contractISwapContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNodeStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"updateAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SbBTCPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use SbBTCPoolMetaData.ABI instead.
var SbBTCPoolABI = SbBTCPoolMetaData.ABI

// SbBTCPool is an auto generated Go binding around an Ethereum contract.
type SbBTCPool struct {
	SbBTCPoolCaller     // Read-only binding to the contract
	SbBTCPoolTransactor // Write-only binding to the contract
	SbBTCPoolFilterer   // Log filterer for contract events
}

// SbBTCPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SbBTCPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SbBTCPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SbBTCPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SbBTCPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SbBTCPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SbBTCPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SbBTCPoolSession struct {
	Contract     *SbBTCPool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SbBTCPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SbBTCPoolCallerSession struct {
	Contract *SbBTCPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SbBTCPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SbBTCPoolTransactorSession struct {
	Contract     *SbBTCPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SbBTCPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SbBTCPoolRaw struct {
	Contract *SbBTCPool // Generic contract binding to access the raw methods on
}

// SbBTCPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SbBTCPoolCallerRaw struct {
	Contract *SbBTCPoolCaller // Generic read-only contract binding to access the raw methods on
}

// SbBTCPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SbBTCPoolTransactorRaw struct {
	Contract *SbBTCPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSbBTCPool creates a new instance of SbBTCPool, bound to a specific deployed contract.
func NewSbBTCPool(address common.Address, backend bind.ContractBackend) (*SbBTCPool, error) {
	contract, err := bindSbBTCPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SbBTCPool{SbBTCPoolCaller: SbBTCPoolCaller{contract: contract}, SbBTCPoolTransactor: SbBTCPoolTransactor{contract: contract}, SbBTCPoolFilterer: SbBTCPoolFilterer{contract: contract}}, nil
}

// NewSbBTCPoolCaller creates a new read-only instance of SbBTCPool, bound to a specific deployed contract.
func NewSbBTCPoolCaller(address common.Address, caller bind.ContractCaller) (*SbBTCPoolCaller, error) {
	contract, err := bindSbBTCPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SbBTCPoolCaller{contract: contract}, nil
}

// NewSbBTCPoolTransactor creates a new write-only instance of SbBTCPool, bound to a specific deployed contract.
func NewSbBTCPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SbBTCPoolTransactor, error) {
	contract, err := bindSbBTCPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SbBTCPoolTransactor{contract: contract}, nil
}

// NewSbBTCPoolFilterer creates a new log filterer instance of SbBTCPool, bound to a specific deployed contract.
func NewSbBTCPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SbBTCPoolFilterer, error) {
	contract, err := bindSbBTCPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SbBTCPoolFilterer{contract: contract}, nil
}

// bindSbBTCPool binds a generic wrapper to an already deployed contract.
func bindSbBTCPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SbBTCPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SbBTCPool *SbBTCPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SbBTCPool.Contract.SbBTCPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SbBTCPool *SbBTCPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SbBTCPool.Contract.SbBTCPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SbBTCPool *SbBTCPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SbBTCPool.Contract.SbBTCPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SbBTCPool *SbBTCPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SbBTCPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SbBTCPool *SbBTCPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SbBTCPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SbBTCPool *SbBTCPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SbBTCPool.Contract.contract.Transact(opts, method, params...)
}

// BalanceBefore is a free data retrieval call binding the contract method 0x94b5798a.
//
// Solidity: function balanceBefore() view returns(uint256)
func (_SbBTCPool *SbBTCPoolCaller) BalanceBefore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SbBTCPool.contract.Call(opts, &out, "balanceBefore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceBefore is a free data retrieval call binding the contract method 0x94b5798a.
//
// Solidity: function balanceBefore() view returns(uint256)
func (_SbBTCPool *SbBTCPoolSession) BalanceBefore() (*big.Int, error) {
	return _SbBTCPool.Contract.BalanceBefore(&_SbBTCPool.CallOpts)
}

// BalanceBefore is a free data retrieval call binding the contract method 0x94b5798a.
//
// Solidity: function balanceBefore() view returns(uint256)
func (_SbBTCPool *SbBTCPoolCallerSession) BalanceBefore() (*big.Int, error) {
	return _SbBTCPool.Contract.BalanceBefore(&_SbBTCPool.CallOpts)
}

// Barn is a free data retrieval call binding the contract method 0x194f480e.
//
// Solidity: function barn() view returns(address)
func (_SbBTCPool *SbBTCPoolCaller) Barn(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SbBTCPool.contract.Call(opts, &out, "barn")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Barn is a free data retrieval call binding the contract method 0x194f480e.
//
// Solidity: function barn() view returns(address)
func (_SbBTCPool *SbBTCPoolSession) Barn() (common.Address, error) {
	return _SbBTCPool.Contract.Barn(&_SbBTCPool.CallOpts)
}

// Barn is a free data retrieval call binding the contract method 0x194f480e.
//
// Solidity: function barn() view returns(address)
func (_SbBTCPool *SbBTCPoolCallerSession) Barn() (common.Address, error) {
	return _SbBTCPool.Contract.Barn(&_SbBTCPool.CallOpts)
}

// CurrentMultiplier is a free data retrieval call binding the contract method 0x6fbaaa1e.
//
// Solidity: function currentMultiplier() view returns(uint256)
func (_SbBTCPool *SbBTCPoolCaller) CurrentMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SbBTCPool.contract.Call(opts, &out, "currentMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentMultiplier is a free data retrieval call binding the contract method 0x6fbaaa1e.
//
// Solidity: function currentMultiplier() view returns(uint256)
func (_SbBTCPool *SbBTCPoolSession) CurrentMultiplier() (*big.Int, error) {
	return _SbBTCPool.Contract.CurrentMultiplier(&_SbBTCPool.CallOpts)
}

// CurrentMultiplier is a free data retrieval call binding the contract method 0x6fbaaa1e.
//
// Solidity: function currentMultiplier() view returns(uint256)
func (_SbBTCPool *SbBTCPoolCallerSession) CurrentMultiplier() (*big.Int, error) {
	return _SbBTCPool.Contract.CurrentMultiplier(&_SbBTCPool.CallOpts)
}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_SbBTCPool *SbBTCPoolCaller) Owed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SbBTCPool.contract.Call(opts, &out, "owed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_SbBTCPool *SbBTCPoolSession) Owed(arg0 common.Address) (*big.Int, error) {
	return _SbBTCPool.Contract.Owed(&_SbBTCPool.CallOpts, arg0)
}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_SbBTCPool *SbBTCPoolCallerSession) Owed(arg0 common.Address) (*big.Int, error) {
	return _SbBTCPool.Contract.Owed(&_SbBTCPool.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SbBTCPool *SbBTCPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SbBTCPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SbBTCPool *SbBTCPoolSession) Owner() (common.Address, error) {
	return _SbBTCPool.Contract.Owner(&_SbBTCPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SbBTCPool *SbBTCPoolCallerSession) Owner() (common.Address, error) {
	return _SbBTCPool.Contract.Owner(&_SbBTCPool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SbBTCPool *SbBTCPoolCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SbBTCPool.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SbBTCPool *SbBTCPoolSession) RewardToken() (common.Address, error) {
	return _SbBTCPool.Contract.RewardToken(&_SbBTCPool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SbBTCPool *SbBTCPoolCallerSession) RewardToken() (common.Address, error) {
	return _SbBTCPool.Contract.RewardToken(&_SbBTCPool.CallOpts)
}

// SwapContract is a free data retrieval call binding the contract method 0x8ea83031.
//
// Solidity: function swapContract() view returns(address)
func (_SbBTCPool *SbBTCPoolCaller) SwapContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SbBTCPool.contract.Call(opts, &out, "swapContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapContract is a free data retrieval call binding the contract method 0x8ea83031.
//
// Solidity: function swapContract() view returns(address)
func (_SbBTCPool *SbBTCPoolSession) SwapContract() (common.Address, error) {
	return _SbBTCPool.Contract.SwapContract(&_SbBTCPool.CallOpts)
}

// SwapContract is a free data retrieval call binding the contract method 0x8ea83031.
//
// Solidity: function swapContract() view returns(address)
func (_SbBTCPool *SbBTCPoolCallerSession) SwapContract() (common.Address, error) {
	return _SbBTCPool.Contract.SwapContract(&_SbBTCPool.CallOpts)
}

// TotalNodeStaked is a free data retrieval call binding the contract method 0x8ca0a7e5.
//
// Solidity: function totalNodeStaked() view returns(uint256)
func (_SbBTCPool *SbBTCPoolCaller) TotalNodeStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SbBTCPool.contract.Call(opts, &out, "totalNodeStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNodeStaked is a free data retrieval call binding the contract method 0x8ca0a7e5.
//
// Solidity: function totalNodeStaked() view returns(uint256)
func (_SbBTCPool *SbBTCPoolSession) TotalNodeStaked() (*big.Int, error) {
	return _SbBTCPool.Contract.TotalNodeStaked(&_SbBTCPool.CallOpts)
}

// TotalNodeStaked is a free data retrieval call binding the contract method 0x8ca0a7e5.
//
// Solidity: function totalNodeStaked() view returns(uint256)
func (_SbBTCPool *SbBTCPoolCallerSession) TotalNodeStaked() (*big.Int, error) {
	return _SbBTCPool.Contract.TotalNodeStaked(&_SbBTCPool.CallOpts)
}

// UserMultiplier is a free data retrieval call binding the contract method 0xb1a03b6b.
//
// Solidity: function userMultiplier(address ) view returns(uint256)
func (_SbBTCPool *SbBTCPoolCaller) UserMultiplier(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SbBTCPool.contract.Call(opts, &out, "userMultiplier", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMultiplier is a free data retrieval call binding the contract method 0xb1a03b6b.
//
// Solidity: function userMultiplier(address ) view returns(uint256)
func (_SbBTCPool *SbBTCPoolSession) UserMultiplier(arg0 common.Address) (*big.Int, error) {
	return _SbBTCPool.Contract.UserMultiplier(&_SbBTCPool.CallOpts, arg0)
}

// UserMultiplier is a free data retrieval call binding the contract method 0xb1a03b6b.
//
// Solidity: function userMultiplier(address ) view returns(uint256)
func (_SbBTCPool *SbBTCPoolCallerSession) UserMultiplier(arg0 common.Address) (*big.Int, error) {
	return _SbBTCPool.Contract.UserMultiplier(&_SbBTCPool.CallOpts, arg0)
}

// AckFunds is a paid mutator transaction binding the contract method 0xacfd9325.
//
// Solidity: function ackFunds() returns()
func (_SbBTCPool *SbBTCPoolTransactor) AckFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SbBTCPool.contract.Transact(opts, "ackFunds")
}

// AckFunds is a paid mutator transaction binding the contract method 0xacfd9325.
//
// Solidity: function ackFunds() returns()
func (_SbBTCPool *SbBTCPoolSession) AckFunds() (*types.Transaction, error) {
	return _SbBTCPool.Contract.AckFunds(&_SbBTCPool.TransactOpts)
}

// AckFunds is a paid mutator transaction binding the contract method 0xacfd9325.
//
// Solidity: function ackFunds() returns()
func (_SbBTCPool *SbBTCPoolTransactorSession) AckFunds() (*types.Transaction, error) {
	return _SbBTCPool.Contract.AckFunds(&_SbBTCPool.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_SbBTCPool *SbBTCPoolTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SbBTCPool.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_SbBTCPool *SbBTCPoolSession) Claim() (*types.Transaction, error) {
	return _SbBTCPool.Contract.Claim(&_SbBTCPool.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_SbBTCPool *SbBTCPoolTransactorSession) Claim() (*types.Transaction, error) {
	return _SbBTCPool.Contract.Claim(&_SbBTCPool.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_SbBTCPool *SbBTCPoolTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SbBTCPool.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_SbBTCPool *SbBTCPoolSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _SbBTCPool.Contract.EmergencyWithdraw(&_SbBTCPool.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_SbBTCPool *SbBTCPoolTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _SbBTCPool.Contract.EmergencyWithdraw(&_SbBTCPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SbBTCPool *SbBTCPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SbBTCPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SbBTCPool *SbBTCPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _SbBTCPool.Contract.RenounceOwnership(&_SbBTCPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SbBTCPool *SbBTCPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SbBTCPool.Contract.RenounceOwnership(&_SbBTCPool.TransactOpts)
}

// ResetUnstakedNode is a paid mutator transaction binding the contract method 0x58411c09.
//
// Solidity: function resetUnstakedNode(address _node) returns()
func (_SbBTCPool *SbBTCPoolTransactor) ResetUnstakedNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _SbBTCPool.contract.Transact(opts, "resetUnstakedNode", _node)
}

// ResetUnstakedNode is a paid mutator transaction binding the contract method 0x58411c09.
//
// Solidity: function resetUnstakedNode(address _node) returns()
func (_SbBTCPool *SbBTCPoolSession) ResetUnstakedNode(_node common.Address) (*types.Transaction, error) {
	return _SbBTCPool.Contract.ResetUnstakedNode(&_SbBTCPool.TransactOpts, _node)
}

// ResetUnstakedNode is a paid mutator transaction binding the contract method 0x58411c09.
//
// Solidity: function resetUnstakedNode(address _node) returns()
func (_SbBTCPool *SbBTCPoolTransactorSession) ResetUnstakedNode(_node common.Address) (*types.Transaction, error) {
	return _SbBTCPool.Contract.ResetUnstakedNode(&_SbBTCPool.TransactOpts, _node)
}

// SetBarnAndSwap is a paid mutator transaction binding the contract method 0x1ed64040.
//
// Solidity: function setBarnAndSwap(address _barn, address _swap) returns()
func (_SbBTCPool *SbBTCPoolTransactor) SetBarnAndSwap(opts *bind.TransactOpts, _barn common.Address, _swap common.Address) (*types.Transaction, error) {
	return _SbBTCPool.contract.Transact(opts, "setBarnAndSwap", _barn, _swap)
}

// SetBarnAndSwap is a paid mutator transaction binding the contract method 0x1ed64040.
//
// Solidity: function setBarnAndSwap(address _barn, address _swap) returns()
func (_SbBTCPool *SbBTCPoolSession) SetBarnAndSwap(_barn common.Address, _swap common.Address) (*types.Transaction, error) {
	return _SbBTCPool.Contract.SetBarnAndSwap(&_SbBTCPool.TransactOpts, _barn, _swap)
}

// SetBarnAndSwap is a paid mutator transaction binding the contract method 0x1ed64040.
//
// Solidity: function setBarnAndSwap(address _barn, address _swap) returns()
func (_SbBTCPool *SbBTCPoolTransactorSession) SetBarnAndSwap(_barn common.Address, _swap common.Address) (*types.Transaction, error) {
	return _SbBTCPool.Contract.SetBarnAndSwap(&_SbBTCPool.TransactOpts, _barn, _swap)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SbBTCPool *SbBTCPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SbBTCPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SbBTCPool *SbBTCPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SbBTCPool.Contract.TransferOwnership(&_SbBTCPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SbBTCPool *SbBTCPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SbBTCPool.Contract.TransferOwnership(&_SbBTCPool.TransactOpts, newOwner)
}

// UpdateAll is a paid mutator transaction binding the contract method 0xa616f345.
//
// Solidity: function updateAll(uint256 _timestamp) returns()
func (_SbBTCPool *SbBTCPoolTransactor) UpdateAll(opts *bind.TransactOpts, _timestamp *big.Int) (*types.Transaction, error) {
	return _SbBTCPool.contract.Transact(opts, "updateAll", _timestamp)
}

// UpdateAll is a paid mutator transaction binding the contract method 0xa616f345.
//
// Solidity: function updateAll(uint256 _timestamp) returns()
func (_SbBTCPool *SbBTCPoolSession) UpdateAll(_timestamp *big.Int) (*types.Transaction, error) {
	return _SbBTCPool.Contract.UpdateAll(&_SbBTCPool.TransactOpts, _timestamp)
}

// UpdateAll is a paid mutator transaction binding the contract method 0xa616f345.
//
// Solidity: function updateAll(uint256 _timestamp) returns()
func (_SbBTCPool *SbBTCPoolTransactorSession) UpdateAll(_timestamp *big.Int) (*types.Transaction, error) {
	return _SbBTCPool.Contract.UpdateAll(&_SbBTCPool.TransactOpts, _timestamp)
}

// SbBTCPoolClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the SbBTCPool contract.
type SbBTCPoolClaimIterator struct {
	Event *SbBTCPoolClaim // Event containing the contract specifics and raw log

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
func (it *SbBTCPoolClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SbBTCPoolClaim)
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
		it.Event = new(SbBTCPoolClaim)
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
func (it *SbBTCPoolClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SbBTCPoolClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SbBTCPoolClaim represents a Claim event raised by the SbBTCPool contract.
type SbBTCPoolClaim struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed user, uint256 amount)
func (_SbBTCPool *SbBTCPoolFilterer) FilterClaim(opts *bind.FilterOpts, user []common.Address) (*SbBTCPoolClaimIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SbBTCPool.contract.FilterLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return &SbBTCPoolClaimIterator{contract: _SbBTCPool.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed user, uint256 amount)
func (_SbBTCPool *SbBTCPoolFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *SbBTCPoolClaim, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SbBTCPool.contract.WatchLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SbBTCPoolClaim)
				if err := _SbBTCPool.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed user, uint256 amount)
func (_SbBTCPool *SbBTCPoolFilterer) ParseClaim(log types.Log) (*SbBTCPoolClaim, error) {
	event := new(SbBTCPoolClaim)
	if err := _SbBTCPool.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SbBTCPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SbBTCPool contract.
type SbBTCPoolOwnershipTransferredIterator struct {
	Event *SbBTCPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SbBTCPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SbBTCPoolOwnershipTransferred)
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
		it.Event = new(SbBTCPoolOwnershipTransferred)
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
func (it *SbBTCPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SbBTCPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SbBTCPoolOwnershipTransferred represents a OwnershipTransferred event raised by the SbBTCPool contract.
type SbBTCPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SbBTCPool *SbBTCPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SbBTCPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SbBTCPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SbBTCPoolOwnershipTransferredIterator{contract: _SbBTCPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SbBTCPool *SbBTCPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SbBTCPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SbBTCPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SbBTCPoolOwnershipTransferred)
				if err := _SbBTCPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SbBTCPool *SbBTCPoolFilterer) ParseOwnershipTransferred(log types.Log) (*SbBTCPoolOwnershipTransferred, error) {
	event := new(SbBTCPoolOwnershipTransferred)
	if err := _SbBTCPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
