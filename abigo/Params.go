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

// ParamsMetaData contains all meta data concerning the Params contract.
var ParamsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"depositFeesBPS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expirationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loopCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumSwapAmountForWBTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeRewardsRatio\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paraswapAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_depositFeesBPS\",\"type\":\"uint8\"}],\"name\":\"setDepositFeesBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expirationTime\",\"type\":\"uint256\"}],\"name\":\"setExpirationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_loopCount\",\"type\":\"uint8\"}],\"name\":\"setLoopCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumSwapAmountForWBTC\",\"type\":\"uint256\"}],\"name\":\"setMinimumSwapAmountForWBTC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_nodeRewardsRatio\",\"type\":\"uint8\"}],\"name\":\"setNodeRewardsRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_paraswapAddress\",\"type\":\"address\"}],\"name\":\"setParaswapAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_withdrawalFeeBPS\",\"type\":\"uint8\"}],\"name\":\"setWithdrawalFeeBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalFeeBPS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ParamsABI is the input ABI used to generate the binding from.
// Deprecated: Use ParamsMetaData.ABI instead.
var ParamsABI = ParamsMetaData.ABI

// Params is an auto generated Go binding around an Ethereum contract.
type Params struct {
	ParamsCaller     // Read-only binding to the contract
	ParamsTransactor // Write-only binding to the contract
	ParamsFilterer   // Log filterer for contract events
}

// ParamsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParamsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParamsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParamsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParamsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParamsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParamsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParamsSession struct {
	Contract     *Params           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParamsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParamsCallerSession struct {
	Contract *ParamsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ParamsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParamsTransactorSession struct {
	Contract     *ParamsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParamsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParamsRaw struct {
	Contract *Params // Generic contract binding to access the raw methods on
}

// ParamsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParamsCallerRaw struct {
	Contract *ParamsCaller // Generic read-only contract binding to access the raw methods on
}

// ParamsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParamsTransactorRaw struct {
	Contract *ParamsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParams creates a new instance of Params, bound to a specific deployed contract.
func NewParams(address common.Address, backend bind.ContractBackend) (*Params, error) {
	contract, err := bindParams(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Params{ParamsCaller: ParamsCaller{contract: contract}, ParamsTransactor: ParamsTransactor{contract: contract}, ParamsFilterer: ParamsFilterer{contract: contract}}, nil
}

// NewParamsCaller creates a new read-only instance of Params, bound to a specific deployed contract.
func NewParamsCaller(address common.Address, caller bind.ContractCaller) (*ParamsCaller, error) {
	contract, err := bindParams(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParamsCaller{contract: contract}, nil
}

// NewParamsTransactor creates a new write-only instance of Params, bound to a specific deployed contract.
func NewParamsTransactor(address common.Address, transactor bind.ContractTransactor) (*ParamsTransactor, error) {
	contract, err := bindParams(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParamsTransactor{contract: contract}, nil
}

// NewParamsFilterer creates a new log filterer instance of Params, bound to a specific deployed contract.
func NewParamsFilterer(address common.Address, filterer bind.ContractFilterer) (*ParamsFilterer, error) {
	contract, err := bindParams(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParamsFilterer{contract: contract}, nil
}

// bindParams binds a generic wrapper to an already deployed contract.
func bindParams(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParamsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Params *ParamsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Params.Contract.ParamsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Params *ParamsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Params.Contract.ParamsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Params *ParamsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Params.Contract.ParamsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Params *ParamsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Params.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Params *ParamsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Params.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Params *ParamsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Params.Contract.contract.Transact(opts, method, params...)
}

// DepositFeesBPS is a free data retrieval call binding the contract method 0x42419255.
//
// Solidity: function depositFeesBPS() view returns(uint8)
func (_Params *ParamsCaller) DepositFeesBPS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "depositFeesBPS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DepositFeesBPS is a free data retrieval call binding the contract method 0x42419255.
//
// Solidity: function depositFeesBPS() view returns(uint8)
func (_Params *ParamsSession) DepositFeesBPS() (uint8, error) {
	return _Params.Contract.DepositFeesBPS(&_Params.CallOpts)
}

// DepositFeesBPS is a free data retrieval call binding the contract method 0x42419255.
//
// Solidity: function depositFeesBPS() view returns(uint8)
func (_Params *ParamsCallerSession) DepositFeesBPS() (uint8, error) {
	return _Params.Contract.DepositFeesBPS(&_Params.CallOpts)
}

// ExpirationTime is a free data retrieval call binding the contract method 0xda284dcc.
//
// Solidity: function expirationTime() view returns(uint256)
func (_Params *ParamsCaller) ExpirationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "expirationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpirationTime is a free data retrieval call binding the contract method 0xda284dcc.
//
// Solidity: function expirationTime() view returns(uint256)
func (_Params *ParamsSession) ExpirationTime() (*big.Int, error) {
	return _Params.Contract.ExpirationTime(&_Params.CallOpts)
}

// ExpirationTime is a free data retrieval call binding the contract method 0xda284dcc.
//
// Solidity: function expirationTime() view returns(uint256)
func (_Params *ParamsCallerSession) ExpirationTime() (*big.Int, error) {
	return _Params.Contract.ExpirationTime(&_Params.CallOpts)
}

// LoopCount is a free data retrieval call binding the contract method 0xe91675b8.
//
// Solidity: function loopCount() view returns(uint8)
func (_Params *ParamsCaller) LoopCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "loopCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LoopCount is a free data retrieval call binding the contract method 0xe91675b8.
//
// Solidity: function loopCount() view returns(uint8)
func (_Params *ParamsSession) LoopCount() (uint8, error) {
	return _Params.Contract.LoopCount(&_Params.CallOpts)
}

// LoopCount is a free data retrieval call binding the contract method 0xe91675b8.
//
// Solidity: function loopCount() view returns(uint8)
func (_Params *ParamsCallerSession) LoopCount() (uint8, error) {
	return _Params.Contract.LoopCount(&_Params.CallOpts)
}

// MinimumSwapAmountForWBTC is a free data retrieval call binding the contract method 0x1411c5b7.
//
// Solidity: function minimumSwapAmountForWBTC() view returns(uint256)
func (_Params *ParamsCaller) MinimumSwapAmountForWBTC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "minimumSwapAmountForWBTC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumSwapAmountForWBTC is a free data retrieval call binding the contract method 0x1411c5b7.
//
// Solidity: function minimumSwapAmountForWBTC() view returns(uint256)
func (_Params *ParamsSession) MinimumSwapAmountForWBTC() (*big.Int, error) {
	return _Params.Contract.MinimumSwapAmountForWBTC(&_Params.CallOpts)
}

// MinimumSwapAmountForWBTC is a free data retrieval call binding the contract method 0x1411c5b7.
//
// Solidity: function minimumSwapAmountForWBTC() view returns(uint256)
func (_Params *ParamsCallerSession) MinimumSwapAmountForWBTC() (*big.Int, error) {
	return _Params.Contract.MinimumSwapAmountForWBTC(&_Params.CallOpts)
}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint8)
func (_Params *ParamsCaller) NodeRewardsRatio(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "nodeRewardsRatio")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint8)
func (_Params *ParamsSession) NodeRewardsRatio() (uint8, error) {
	return _Params.Contract.NodeRewardsRatio(&_Params.CallOpts)
}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint8)
func (_Params *ParamsCallerSession) NodeRewardsRatio() (uint8, error) {
	return _Params.Contract.NodeRewardsRatio(&_Params.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Params *ParamsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Params *ParamsSession) Owner() (common.Address, error) {
	return _Params.Contract.Owner(&_Params.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Params *ParamsCallerSession) Owner() (common.Address, error) {
	return _Params.Contract.Owner(&_Params.CallOpts)
}

// ParaswapAddress is a free data retrieval call binding the contract method 0xf10d7c35.
//
// Solidity: function paraswapAddress() view returns(address)
func (_Params *ParamsCaller) ParaswapAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "paraswapAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParaswapAddress is a free data retrieval call binding the contract method 0xf10d7c35.
//
// Solidity: function paraswapAddress() view returns(address)
func (_Params *ParamsSession) ParaswapAddress() (common.Address, error) {
	return _Params.Contract.ParaswapAddress(&_Params.CallOpts)
}

// ParaswapAddress is a free data retrieval call binding the contract method 0xf10d7c35.
//
// Solidity: function paraswapAddress() view returns(address)
func (_Params *ParamsCallerSession) ParaswapAddress() (common.Address, error) {
	return _Params.Contract.ParaswapAddress(&_Params.CallOpts)
}

// WithdrawalFeeBPS is a free data retrieval call binding the contract method 0xb6268e5d.
//
// Solidity: function withdrawalFeeBPS() view returns(uint8)
func (_Params *ParamsCaller) WithdrawalFeeBPS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "withdrawalFeeBPS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// WithdrawalFeeBPS is a free data retrieval call binding the contract method 0xb6268e5d.
//
// Solidity: function withdrawalFeeBPS() view returns(uint8)
func (_Params *ParamsSession) WithdrawalFeeBPS() (uint8, error) {
	return _Params.Contract.WithdrawalFeeBPS(&_Params.CallOpts)
}

// WithdrawalFeeBPS is a free data retrieval call binding the contract method 0xb6268e5d.
//
// Solidity: function withdrawalFeeBPS() view returns(uint8)
func (_Params *ParamsCallerSession) WithdrawalFeeBPS() (uint8, error) {
	return _Params.Contract.WithdrawalFeeBPS(&_Params.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Params *ParamsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Params.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Params *ParamsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Params.Contract.RenounceOwnership(&_Params.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Params *ParamsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Params.Contract.RenounceOwnership(&_Params.TransactOpts)
}

// SetDepositFeesBPS is a paid mutator transaction binding the contract method 0x6ce81f80.
//
// Solidity: function setDepositFeesBPS(uint8 _depositFeesBPS) returns()
func (_Params *ParamsTransactor) SetDepositFeesBPS(opts *bind.TransactOpts, _depositFeesBPS uint8) (*types.Transaction, error) {
	return _Params.contract.Transact(opts, "setDepositFeesBPS", _depositFeesBPS)
}

// SetDepositFeesBPS is a paid mutator transaction binding the contract method 0x6ce81f80.
//
// Solidity: function setDepositFeesBPS(uint8 _depositFeesBPS) returns()
func (_Params *ParamsSession) SetDepositFeesBPS(_depositFeesBPS uint8) (*types.Transaction, error) {
	return _Params.Contract.SetDepositFeesBPS(&_Params.TransactOpts, _depositFeesBPS)
}

// SetDepositFeesBPS is a paid mutator transaction binding the contract method 0x6ce81f80.
//
// Solidity: function setDepositFeesBPS(uint8 _depositFeesBPS) returns()
func (_Params *ParamsTransactorSession) SetDepositFeesBPS(_depositFeesBPS uint8) (*types.Transaction, error) {
	return _Params.Contract.SetDepositFeesBPS(&_Params.TransactOpts, _depositFeesBPS)
}

// SetExpirationTime is a paid mutator transaction binding the contract method 0xc0cc365d.
//
// Solidity: function setExpirationTime(uint256 _expirationTime) returns()
func (_Params *ParamsTransactor) SetExpirationTime(opts *bind.TransactOpts, _expirationTime *big.Int) (*types.Transaction, error) {
	return _Params.contract.Transact(opts, "setExpirationTime", _expirationTime)
}

// SetExpirationTime is a paid mutator transaction binding the contract method 0xc0cc365d.
//
// Solidity: function setExpirationTime(uint256 _expirationTime) returns()
func (_Params *ParamsSession) SetExpirationTime(_expirationTime *big.Int) (*types.Transaction, error) {
	return _Params.Contract.SetExpirationTime(&_Params.TransactOpts, _expirationTime)
}

// SetExpirationTime is a paid mutator transaction binding the contract method 0xc0cc365d.
//
// Solidity: function setExpirationTime(uint256 _expirationTime) returns()
func (_Params *ParamsTransactorSession) SetExpirationTime(_expirationTime *big.Int) (*types.Transaction, error) {
	return _Params.Contract.SetExpirationTime(&_Params.TransactOpts, _expirationTime)
}

// SetLoopCount is a paid mutator transaction binding the contract method 0x56151c06.
//
// Solidity: function setLoopCount(uint8 _loopCount) returns()
func (_Params *ParamsTransactor) SetLoopCount(opts *bind.TransactOpts, _loopCount uint8) (*types.Transaction, error) {
	return _Params.contract.Transact(opts, "setLoopCount", _loopCount)
}

// SetLoopCount is a paid mutator transaction binding the contract method 0x56151c06.
//
// Solidity: function setLoopCount(uint8 _loopCount) returns()
func (_Params *ParamsSession) SetLoopCount(_loopCount uint8) (*types.Transaction, error) {
	return _Params.Contract.SetLoopCount(&_Params.TransactOpts, _loopCount)
}

// SetLoopCount is a paid mutator transaction binding the contract method 0x56151c06.
//
// Solidity: function setLoopCount(uint8 _loopCount) returns()
func (_Params *ParamsTransactorSession) SetLoopCount(_loopCount uint8) (*types.Transaction, error) {
	return _Params.Contract.SetLoopCount(&_Params.TransactOpts, _loopCount)
}

// SetMinimumSwapAmountForWBTC is a paid mutator transaction binding the contract method 0x5f79ec58.
//
// Solidity: function setMinimumSwapAmountForWBTC(uint256 _minimumSwapAmountForWBTC) returns()
func (_Params *ParamsTransactor) SetMinimumSwapAmountForWBTC(opts *bind.TransactOpts, _minimumSwapAmountForWBTC *big.Int) (*types.Transaction, error) {
	return _Params.contract.Transact(opts, "setMinimumSwapAmountForWBTC", _minimumSwapAmountForWBTC)
}

// SetMinimumSwapAmountForWBTC is a paid mutator transaction binding the contract method 0x5f79ec58.
//
// Solidity: function setMinimumSwapAmountForWBTC(uint256 _minimumSwapAmountForWBTC) returns()
func (_Params *ParamsSession) SetMinimumSwapAmountForWBTC(_minimumSwapAmountForWBTC *big.Int) (*types.Transaction, error) {
	return _Params.Contract.SetMinimumSwapAmountForWBTC(&_Params.TransactOpts, _minimumSwapAmountForWBTC)
}

// SetMinimumSwapAmountForWBTC is a paid mutator transaction binding the contract method 0x5f79ec58.
//
// Solidity: function setMinimumSwapAmountForWBTC(uint256 _minimumSwapAmountForWBTC) returns()
func (_Params *ParamsTransactorSession) SetMinimumSwapAmountForWBTC(_minimumSwapAmountForWBTC *big.Int) (*types.Transaction, error) {
	return _Params.Contract.SetMinimumSwapAmountForWBTC(&_Params.TransactOpts, _minimumSwapAmountForWBTC)
}

// SetNodeRewardsRatio is a paid mutator transaction binding the contract method 0xa0011cd4.
//
// Solidity: function setNodeRewardsRatio(uint8 _nodeRewardsRatio) returns()
func (_Params *ParamsTransactor) SetNodeRewardsRatio(opts *bind.TransactOpts, _nodeRewardsRatio uint8) (*types.Transaction, error) {
	return _Params.contract.Transact(opts, "setNodeRewardsRatio", _nodeRewardsRatio)
}

// SetNodeRewardsRatio is a paid mutator transaction binding the contract method 0xa0011cd4.
//
// Solidity: function setNodeRewardsRatio(uint8 _nodeRewardsRatio) returns()
func (_Params *ParamsSession) SetNodeRewardsRatio(_nodeRewardsRatio uint8) (*types.Transaction, error) {
	return _Params.Contract.SetNodeRewardsRatio(&_Params.TransactOpts, _nodeRewardsRatio)
}

// SetNodeRewardsRatio is a paid mutator transaction binding the contract method 0xa0011cd4.
//
// Solidity: function setNodeRewardsRatio(uint8 _nodeRewardsRatio) returns()
func (_Params *ParamsTransactorSession) SetNodeRewardsRatio(_nodeRewardsRatio uint8) (*types.Transaction, error) {
	return _Params.Contract.SetNodeRewardsRatio(&_Params.TransactOpts, _nodeRewardsRatio)
}

// SetParaswapAddress is a paid mutator transaction binding the contract method 0xfb278493.
//
// Solidity: function setParaswapAddress(address _paraswapAddress) returns()
func (_Params *ParamsTransactor) SetParaswapAddress(opts *bind.TransactOpts, _paraswapAddress common.Address) (*types.Transaction, error) {
	return _Params.contract.Transact(opts, "setParaswapAddress", _paraswapAddress)
}

// SetParaswapAddress is a paid mutator transaction binding the contract method 0xfb278493.
//
// Solidity: function setParaswapAddress(address _paraswapAddress) returns()
func (_Params *ParamsSession) SetParaswapAddress(_paraswapAddress common.Address) (*types.Transaction, error) {
	return _Params.Contract.SetParaswapAddress(&_Params.TransactOpts, _paraswapAddress)
}

// SetParaswapAddress is a paid mutator transaction binding the contract method 0xfb278493.
//
// Solidity: function setParaswapAddress(address _paraswapAddress) returns()
func (_Params *ParamsTransactorSession) SetParaswapAddress(_paraswapAddress common.Address) (*types.Transaction, error) {
	return _Params.Contract.SetParaswapAddress(&_Params.TransactOpts, _paraswapAddress)
}

// SetWithdrawalFeeBPS is a paid mutator transaction binding the contract method 0xdd62a515.
//
// Solidity: function setWithdrawalFeeBPS(uint8 _withdrawalFeeBPS) returns()
func (_Params *ParamsTransactor) SetWithdrawalFeeBPS(opts *bind.TransactOpts, _withdrawalFeeBPS uint8) (*types.Transaction, error) {
	return _Params.contract.Transact(opts, "setWithdrawalFeeBPS", _withdrawalFeeBPS)
}

// SetWithdrawalFeeBPS is a paid mutator transaction binding the contract method 0xdd62a515.
//
// Solidity: function setWithdrawalFeeBPS(uint8 _withdrawalFeeBPS) returns()
func (_Params *ParamsSession) SetWithdrawalFeeBPS(_withdrawalFeeBPS uint8) (*types.Transaction, error) {
	return _Params.Contract.SetWithdrawalFeeBPS(&_Params.TransactOpts, _withdrawalFeeBPS)
}

// SetWithdrawalFeeBPS is a paid mutator transaction binding the contract method 0xdd62a515.
//
// Solidity: function setWithdrawalFeeBPS(uint8 _withdrawalFeeBPS) returns()
func (_Params *ParamsTransactorSession) SetWithdrawalFeeBPS(_withdrawalFeeBPS uint8) (*types.Transaction, error) {
	return _Params.Contract.SetWithdrawalFeeBPS(&_Params.TransactOpts, _withdrawalFeeBPS)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Params *ParamsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Params.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Params *ParamsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Params.Contract.TransferOwnership(&_Params.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Params *ParamsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Params.Contract.TransferOwnership(&_Params.TransactOpts, newOwner)
}

// ParamsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Params contract.
type ParamsOwnershipTransferredIterator struct {
	Event *ParamsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ParamsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParamsOwnershipTransferred)
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
		it.Event = new(ParamsOwnershipTransferred)
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
func (it *ParamsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParamsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParamsOwnershipTransferred represents a OwnershipTransferred event raised by the Params contract.
type ParamsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Params *ParamsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ParamsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Params.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ParamsOwnershipTransferredIterator{contract: _Params.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Params *ParamsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ParamsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Params.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParamsOwnershipTransferred)
				if err := _Params.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Params *ParamsFilterer) ParseOwnershipTransferred(log types.Log) (*ParamsOwnershipTransferred, error) {
	event := new(ParamsOwnershipTransferred)
	if err := _Params.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
