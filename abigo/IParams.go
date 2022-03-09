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

// IParamsMetaData contains all meta data concerning the IParams contract.
var IParamsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"depositFeesBPS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expirationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loopCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumSwapAmountForWBTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeRewardsRatio\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paraswapAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_depositFeesBPS\",\"type\":\"uint8\"}],\"name\":\"setDepositFeesBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expirationTime\",\"type\":\"uint256\"}],\"name\":\"setExpirationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_loopCount\",\"type\":\"uint8\"}],\"name\":\"setLoopCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumSwapAmountForWBTC\",\"type\":\"uint256\"}],\"name\":\"setMinimumSwapAmountForWBTC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_nodeRewardsRatio\",\"type\":\"uint8\"}],\"name\":\"setNodeRewardsRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_paraswapAddress\",\"type\":\"address\"}],\"name\":\"setParaswapAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_withdrawalFeeBPS\",\"type\":\"uint8\"}],\"name\":\"setWithdrawalFeeBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalFeeBPS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IParamsABI is the input ABI used to generate the binding from.
// Deprecated: Use IParamsMetaData.ABI instead.
var IParamsABI = IParamsMetaData.ABI

// IParams is an auto generated Go binding around an Ethereum contract.
type IParams struct {
	IParamsCaller     // Read-only binding to the contract
	IParamsTransactor // Write-only binding to the contract
	IParamsFilterer   // Log filterer for contract events
}

// IParamsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IParamsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParamsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IParamsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParamsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IParamsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParamsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IParamsSession struct {
	Contract     *IParams          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IParamsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IParamsCallerSession struct {
	Contract *IParamsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IParamsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IParamsTransactorSession struct {
	Contract     *IParamsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IParamsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IParamsRaw struct {
	Contract *IParams // Generic contract binding to access the raw methods on
}

// IParamsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IParamsCallerRaw struct {
	Contract *IParamsCaller // Generic read-only contract binding to access the raw methods on
}

// IParamsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IParamsTransactorRaw struct {
	Contract *IParamsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIParams creates a new instance of IParams, bound to a specific deployed contract.
func NewIParams(address common.Address, backend bind.ContractBackend) (*IParams, error) {
	contract, err := bindIParams(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IParams{IParamsCaller: IParamsCaller{contract: contract}, IParamsTransactor: IParamsTransactor{contract: contract}, IParamsFilterer: IParamsFilterer{contract: contract}}, nil
}

// NewIParamsCaller creates a new read-only instance of IParams, bound to a specific deployed contract.
func NewIParamsCaller(address common.Address, caller bind.ContractCaller) (*IParamsCaller, error) {
	contract, err := bindIParams(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IParamsCaller{contract: contract}, nil
}

// NewIParamsTransactor creates a new write-only instance of IParams, bound to a specific deployed contract.
func NewIParamsTransactor(address common.Address, transactor bind.ContractTransactor) (*IParamsTransactor, error) {
	contract, err := bindIParams(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IParamsTransactor{contract: contract}, nil
}

// NewIParamsFilterer creates a new log filterer instance of IParams, bound to a specific deployed contract.
func NewIParamsFilterer(address common.Address, filterer bind.ContractFilterer) (*IParamsFilterer, error) {
	contract, err := bindIParams(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IParamsFilterer{contract: contract}, nil
}

// bindIParams binds a generic wrapper to an already deployed contract.
func bindIParams(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IParamsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IParams *IParamsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IParams.Contract.IParamsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IParams *IParamsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IParams.Contract.IParamsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IParams *IParamsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IParams.Contract.IParamsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IParams *IParamsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IParams.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IParams *IParamsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IParams.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IParams *IParamsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IParams.Contract.contract.Transact(opts, method, params...)
}

// DepositFeesBPS is a free data retrieval call binding the contract method 0x42419255.
//
// Solidity: function depositFeesBPS() view returns(uint8)
func (_IParams *IParamsCaller) DepositFeesBPS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IParams.contract.Call(opts, &out, "depositFeesBPS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DepositFeesBPS is a free data retrieval call binding the contract method 0x42419255.
//
// Solidity: function depositFeesBPS() view returns(uint8)
func (_IParams *IParamsSession) DepositFeesBPS() (uint8, error) {
	return _IParams.Contract.DepositFeesBPS(&_IParams.CallOpts)
}

// DepositFeesBPS is a free data retrieval call binding the contract method 0x42419255.
//
// Solidity: function depositFeesBPS() view returns(uint8)
func (_IParams *IParamsCallerSession) DepositFeesBPS() (uint8, error) {
	return _IParams.Contract.DepositFeesBPS(&_IParams.CallOpts)
}

// ExpirationTime is a free data retrieval call binding the contract method 0xda284dcc.
//
// Solidity: function expirationTime() view returns(uint256)
func (_IParams *IParamsCaller) ExpirationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IParams.contract.Call(opts, &out, "expirationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpirationTime is a free data retrieval call binding the contract method 0xda284dcc.
//
// Solidity: function expirationTime() view returns(uint256)
func (_IParams *IParamsSession) ExpirationTime() (*big.Int, error) {
	return _IParams.Contract.ExpirationTime(&_IParams.CallOpts)
}

// ExpirationTime is a free data retrieval call binding the contract method 0xda284dcc.
//
// Solidity: function expirationTime() view returns(uint256)
func (_IParams *IParamsCallerSession) ExpirationTime() (*big.Int, error) {
	return _IParams.Contract.ExpirationTime(&_IParams.CallOpts)
}

// LoopCount is a free data retrieval call binding the contract method 0xe91675b8.
//
// Solidity: function loopCount() view returns(uint8)
func (_IParams *IParamsCaller) LoopCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IParams.contract.Call(opts, &out, "loopCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LoopCount is a free data retrieval call binding the contract method 0xe91675b8.
//
// Solidity: function loopCount() view returns(uint8)
func (_IParams *IParamsSession) LoopCount() (uint8, error) {
	return _IParams.Contract.LoopCount(&_IParams.CallOpts)
}

// LoopCount is a free data retrieval call binding the contract method 0xe91675b8.
//
// Solidity: function loopCount() view returns(uint8)
func (_IParams *IParamsCallerSession) LoopCount() (uint8, error) {
	return _IParams.Contract.LoopCount(&_IParams.CallOpts)
}

// MinimumSwapAmountForWBTC is a free data retrieval call binding the contract method 0x1411c5b7.
//
// Solidity: function minimumSwapAmountForWBTC() view returns(uint256)
func (_IParams *IParamsCaller) MinimumSwapAmountForWBTC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IParams.contract.Call(opts, &out, "minimumSwapAmountForWBTC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumSwapAmountForWBTC is a free data retrieval call binding the contract method 0x1411c5b7.
//
// Solidity: function minimumSwapAmountForWBTC() view returns(uint256)
func (_IParams *IParamsSession) MinimumSwapAmountForWBTC() (*big.Int, error) {
	return _IParams.Contract.MinimumSwapAmountForWBTC(&_IParams.CallOpts)
}

// MinimumSwapAmountForWBTC is a free data retrieval call binding the contract method 0x1411c5b7.
//
// Solidity: function minimumSwapAmountForWBTC() view returns(uint256)
func (_IParams *IParamsCallerSession) MinimumSwapAmountForWBTC() (*big.Int, error) {
	return _IParams.Contract.MinimumSwapAmountForWBTC(&_IParams.CallOpts)
}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint8)
func (_IParams *IParamsCaller) NodeRewardsRatio(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IParams.contract.Call(opts, &out, "nodeRewardsRatio")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint8)
func (_IParams *IParamsSession) NodeRewardsRatio() (uint8, error) {
	return _IParams.Contract.NodeRewardsRatio(&_IParams.CallOpts)
}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint8)
func (_IParams *IParamsCallerSession) NodeRewardsRatio() (uint8, error) {
	return _IParams.Contract.NodeRewardsRatio(&_IParams.CallOpts)
}

// ParaswapAddress is a free data retrieval call binding the contract method 0xf10d7c35.
//
// Solidity: function paraswapAddress() view returns(address)
func (_IParams *IParamsCaller) ParaswapAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IParams.contract.Call(opts, &out, "paraswapAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParaswapAddress is a free data retrieval call binding the contract method 0xf10d7c35.
//
// Solidity: function paraswapAddress() view returns(address)
func (_IParams *IParamsSession) ParaswapAddress() (common.Address, error) {
	return _IParams.Contract.ParaswapAddress(&_IParams.CallOpts)
}

// ParaswapAddress is a free data retrieval call binding the contract method 0xf10d7c35.
//
// Solidity: function paraswapAddress() view returns(address)
func (_IParams *IParamsCallerSession) ParaswapAddress() (common.Address, error) {
	return _IParams.Contract.ParaswapAddress(&_IParams.CallOpts)
}

// WithdrawalFeeBPS is a free data retrieval call binding the contract method 0xb6268e5d.
//
// Solidity: function withdrawalFeeBPS() view returns(uint8)
func (_IParams *IParamsCaller) WithdrawalFeeBPS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IParams.contract.Call(opts, &out, "withdrawalFeeBPS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// WithdrawalFeeBPS is a free data retrieval call binding the contract method 0xb6268e5d.
//
// Solidity: function withdrawalFeeBPS() view returns(uint8)
func (_IParams *IParamsSession) WithdrawalFeeBPS() (uint8, error) {
	return _IParams.Contract.WithdrawalFeeBPS(&_IParams.CallOpts)
}

// WithdrawalFeeBPS is a free data retrieval call binding the contract method 0xb6268e5d.
//
// Solidity: function withdrawalFeeBPS() view returns(uint8)
func (_IParams *IParamsCallerSession) WithdrawalFeeBPS() (uint8, error) {
	return _IParams.Contract.WithdrawalFeeBPS(&_IParams.CallOpts)
}

// SetDepositFeesBPS is a paid mutator transaction binding the contract method 0x6ce81f80.
//
// Solidity: function setDepositFeesBPS(uint8 _depositFeesBPS) returns()
func (_IParams *IParamsTransactor) SetDepositFeesBPS(opts *bind.TransactOpts, _depositFeesBPS uint8) (*types.Transaction, error) {
	return _IParams.contract.Transact(opts, "setDepositFeesBPS", _depositFeesBPS)
}

// SetDepositFeesBPS is a paid mutator transaction binding the contract method 0x6ce81f80.
//
// Solidity: function setDepositFeesBPS(uint8 _depositFeesBPS) returns()
func (_IParams *IParamsSession) SetDepositFeesBPS(_depositFeesBPS uint8) (*types.Transaction, error) {
	return _IParams.Contract.SetDepositFeesBPS(&_IParams.TransactOpts, _depositFeesBPS)
}

// SetDepositFeesBPS is a paid mutator transaction binding the contract method 0x6ce81f80.
//
// Solidity: function setDepositFeesBPS(uint8 _depositFeesBPS) returns()
func (_IParams *IParamsTransactorSession) SetDepositFeesBPS(_depositFeesBPS uint8) (*types.Transaction, error) {
	return _IParams.Contract.SetDepositFeesBPS(&_IParams.TransactOpts, _depositFeesBPS)
}

// SetExpirationTime is a paid mutator transaction binding the contract method 0xc0cc365d.
//
// Solidity: function setExpirationTime(uint256 _expirationTime) returns()
func (_IParams *IParamsTransactor) SetExpirationTime(opts *bind.TransactOpts, _expirationTime *big.Int) (*types.Transaction, error) {
	return _IParams.contract.Transact(opts, "setExpirationTime", _expirationTime)
}

// SetExpirationTime is a paid mutator transaction binding the contract method 0xc0cc365d.
//
// Solidity: function setExpirationTime(uint256 _expirationTime) returns()
func (_IParams *IParamsSession) SetExpirationTime(_expirationTime *big.Int) (*types.Transaction, error) {
	return _IParams.Contract.SetExpirationTime(&_IParams.TransactOpts, _expirationTime)
}

// SetExpirationTime is a paid mutator transaction binding the contract method 0xc0cc365d.
//
// Solidity: function setExpirationTime(uint256 _expirationTime) returns()
func (_IParams *IParamsTransactorSession) SetExpirationTime(_expirationTime *big.Int) (*types.Transaction, error) {
	return _IParams.Contract.SetExpirationTime(&_IParams.TransactOpts, _expirationTime)
}

// SetLoopCount is a paid mutator transaction binding the contract method 0x56151c06.
//
// Solidity: function setLoopCount(uint8 _loopCount) returns()
func (_IParams *IParamsTransactor) SetLoopCount(opts *bind.TransactOpts, _loopCount uint8) (*types.Transaction, error) {
	return _IParams.contract.Transact(opts, "setLoopCount", _loopCount)
}

// SetLoopCount is a paid mutator transaction binding the contract method 0x56151c06.
//
// Solidity: function setLoopCount(uint8 _loopCount) returns()
func (_IParams *IParamsSession) SetLoopCount(_loopCount uint8) (*types.Transaction, error) {
	return _IParams.Contract.SetLoopCount(&_IParams.TransactOpts, _loopCount)
}

// SetLoopCount is a paid mutator transaction binding the contract method 0x56151c06.
//
// Solidity: function setLoopCount(uint8 _loopCount) returns()
func (_IParams *IParamsTransactorSession) SetLoopCount(_loopCount uint8) (*types.Transaction, error) {
	return _IParams.Contract.SetLoopCount(&_IParams.TransactOpts, _loopCount)
}

// SetMinimumSwapAmountForWBTC is a paid mutator transaction binding the contract method 0x5f79ec58.
//
// Solidity: function setMinimumSwapAmountForWBTC(uint256 _minimumSwapAmountForWBTC) returns()
func (_IParams *IParamsTransactor) SetMinimumSwapAmountForWBTC(opts *bind.TransactOpts, _minimumSwapAmountForWBTC *big.Int) (*types.Transaction, error) {
	return _IParams.contract.Transact(opts, "setMinimumSwapAmountForWBTC", _minimumSwapAmountForWBTC)
}

// SetMinimumSwapAmountForWBTC is a paid mutator transaction binding the contract method 0x5f79ec58.
//
// Solidity: function setMinimumSwapAmountForWBTC(uint256 _minimumSwapAmountForWBTC) returns()
func (_IParams *IParamsSession) SetMinimumSwapAmountForWBTC(_minimumSwapAmountForWBTC *big.Int) (*types.Transaction, error) {
	return _IParams.Contract.SetMinimumSwapAmountForWBTC(&_IParams.TransactOpts, _minimumSwapAmountForWBTC)
}

// SetMinimumSwapAmountForWBTC is a paid mutator transaction binding the contract method 0x5f79ec58.
//
// Solidity: function setMinimumSwapAmountForWBTC(uint256 _minimumSwapAmountForWBTC) returns()
func (_IParams *IParamsTransactorSession) SetMinimumSwapAmountForWBTC(_minimumSwapAmountForWBTC *big.Int) (*types.Transaction, error) {
	return _IParams.Contract.SetMinimumSwapAmountForWBTC(&_IParams.TransactOpts, _minimumSwapAmountForWBTC)
}

// SetNodeRewardsRatio is a paid mutator transaction binding the contract method 0xa0011cd4.
//
// Solidity: function setNodeRewardsRatio(uint8 _nodeRewardsRatio) returns()
func (_IParams *IParamsTransactor) SetNodeRewardsRatio(opts *bind.TransactOpts, _nodeRewardsRatio uint8) (*types.Transaction, error) {
	return _IParams.contract.Transact(opts, "setNodeRewardsRatio", _nodeRewardsRatio)
}

// SetNodeRewardsRatio is a paid mutator transaction binding the contract method 0xa0011cd4.
//
// Solidity: function setNodeRewardsRatio(uint8 _nodeRewardsRatio) returns()
func (_IParams *IParamsSession) SetNodeRewardsRatio(_nodeRewardsRatio uint8) (*types.Transaction, error) {
	return _IParams.Contract.SetNodeRewardsRatio(&_IParams.TransactOpts, _nodeRewardsRatio)
}

// SetNodeRewardsRatio is a paid mutator transaction binding the contract method 0xa0011cd4.
//
// Solidity: function setNodeRewardsRatio(uint8 _nodeRewardsRatio) returns()
func (_IParams *IParamsTransactorSession) SetNodeRewardsRatio(_nodeRewardsRatio uint8) (*types.Transaction, error) {
	return _IParams.Contract.SetNodeRewardsRatio(&_IParams.TransactOpts, _nodeRewardsRatio)
}

// SetParaswapAddress is a paid mutator transaction binding the contract method 0xfb278493.
//
// Solidity: function setParaswapAddress(address _paraswapAddress) returns()
func (_IParams *IParamsTransactor) SetParaswapAddress(opts *bind.TransactOpts, _paraswapAddress common.Address) (*types.Transaction, error) {
	return _IParams.contract.Transact(opts, "setParaswapAddress", _paraswapAddress)
}

// SetParaswapAddress is a paid mutator transaction binding the contract method 0xfb278493.
//
// Solidity: function setParaswapAddress(address _paraswapAddress) returns()
func (_IParams *IParamsSession) SetParaswapAddress(_paraswapAddress common.Address) (*types.Transaction, error) {
	return _IParams.Contract.SetParaswapAddress(&_IParams.TransactOpts, _paraswapAddress)
}

// SetParaswapAddress is a paid mutator transaction binding the contract method 0xfb278493.
//
// Solidity: function setParaswapAddress(address _paraswapAddress) returns()
func (_IParams *IParamsTransactorSession) SetParaswapAddress(_paraswapAddress common.Address) (*types.Transaction, error) {
	return _IParams.Contract.SetParaswapAddress(&_IParams.TransactOpts, _paraswapAddress)
}

// SetWithdrawalFeeBPS is a paid mutator transaction binding the contract method 0xdd62a515.
//
// Solidity: function setWithdrawalFeeBPS(uint8 _withdrawalFeeBPS) returns()
func (_IParams *IParamsTransactor) SetWithdrawalFeeBPS(opts *bind.TransactOpts, _withdrawalFeeBPS uint8) (*types.Transaction, error) {
	return _IParams.contract.Transact(opts, "setWithdrawalFeeBPS", _withdrawalFeeBPS)
}

// SetWithdrawalFeeBPS is a paid mutator transaction binding the contract method 0xdd62a515.
//
// Solidity: function setWithdrawalFeeBPS(uint8 _withdrawalFeeBPS) returns()
func (_IParams *IParamsSession) SetWithdrawalFeeBPS(_withdrawalFeeBPS uint8) (*types.Transaction, error) {
	return _IParams.Contract.SetWithdrawalFeeBPS(&_IParams.TransactOpts, _withdrawalFeeBPS)
}

// SetWithdrawalFeeBPS is a paid mutator transaction binding the contract method 0xdd62a515.
//
// Solidity: function setWithdrawalFeeBPS(uint8 _withdrawalFeeBPS) returns()
func (_IParams *IParamsTransactorSession) SetWithdrawalFeeBPS(_withdrawalFeeBPS uint8) (*types.Transaction, error) {
	return _IParams.Contract.SetWithdrawalFeeBPS(&_IParams.TransactOpts, _withdrawalFeeBPS)
}
