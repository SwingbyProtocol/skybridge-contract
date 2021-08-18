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

// IBurnableTokenMetaData contains all meta data concerning the IBurnableToken contract.
var IBurnableTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBurnableTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IBurnableTokenMetaData.ABI instead.
var IBurnableTokenABI = IBurnableTokenMetaData.ABI

// IBurnableToken is an auto generated Go binding around an Ethereum contract.
type IBurnableToken struct {
	IBurnableTokenCaller     // Read-only binding to the contract
	IBurnableTokenTransactor // Write-only binding to the contract
	IBurnableTokenFilterer   // Log filterer for contract events
}

// IBurnableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBurnableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBurnableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBurnableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBurnableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBurnableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBurnableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBurnableTokenSession struct {
	Contract     *IBurnableToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBurnableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBurnableTokenCallerSession struct {
	Contract *IBurnableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IBurnableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBurnableTokenTransactorSession struct {
	Contract     *IBurnableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IBurnableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBurnableTokenRaw struct {
	Contract *IBurnableToken // Generic contract binding to access the raw methods on
}

// IBurnableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBurnableTokenCallerRaw struct {
	Contract *IBurnableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IBurnableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBurnableTokenTransactorRaw struct {
	Contract *IBurnableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBurnableToken creates a new instance of IBurnableToken, bound to a specific deployed contract.
func NewIBurnableToken(address common.Address, backend bind.ContractBackend) (*IBurnableToken, error) {
	contract, err := bindIBurnableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBurnableToken{IBurnableTokenCaller: IBurnableTokenCaller{contract: contract}, IBurnableTokenTransactor: IBurnableTokenTransactor{contract: contract}, IBurnableTokenFilterer: IBurnableTokenFilterer{contract: contract}}, nil
}

// NewIBurnableTokenCaller creates a new read-only instance of IBurnableToken, bound to a specific deployed contract.
func NewIBurnableTokenCaller(address common.Address, caller bind.ContractCaller) (*IBurnableTokenCaller, error) {
	contract, err := bindIBurnableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBurnableTokenCaller{contract: contract}, nil
}

// NewIBurnableTokenTransactor creates a new write-only instance of IBurnableToken, bound to a specific deployed contract.
func NewIBurnableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IBurnableTokenTransactor, error) {
	contract, err := bindIBurnableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBurnableTokenTransactor{contract: contract}, nil
}

// NewIBurnableTokenFilterer creates a new log filterer instance of IBurnableToken, bound to a specific deployed contract.
func NewIBurnableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IBurnableTokenFilterer, error) {
	contract, err := bindIBurnableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBurnableTokenFilterer{contract: contract}, nil
}

// bindIBurnableToken binds a generic wrapper to an already deployed contract.
func bindIBurnableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBurnableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBurnableToken *IBurnableTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBurnableToken.Contract.IBurnableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBurnableToken *IBurnableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBurnableToken.Contract.IBurnableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBurnableToken *IBurnableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBurnableToken.Contract.IBurnableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBurnableToken *IBurnableTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBurnableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBurnableToken *IBurnableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBurnableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBurnableToken *IBurnableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBurnableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address spender) view returns(uint256)
func (_IBurnableToken *IBurnableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBurnableToken.contract.Call(opts, &out, "allowance", _owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address spender) view returns(uint256)
func (_IBurnableToken *IBurnableTokenSession) Allowance(_owner common.Address, spender common.Address) (*big.Int, error) {
	return _IBurnableToken.Contract.Allowance(&_IBurnableToken.CallOpts, _owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address spender) view returns(uint256)
func (_IBurnableToken *IBurnableTokenCallerSession) Allowance(_owner common.Address, spender common.Address) (*big.Int, error) {
	return _IBurnableToken.Contract.Allowance(&_IBurnableToken.CallOpts, _owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IBurnableToken *IBurnableTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBurnableToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IBurnableToken *IBurnableTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IBurnableToken.Contract.BalanceOf(&_IBurnableToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IBurnableToken *IBurnableTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IBurnableToken.Contract.BalanceOf(&_IBurnableToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IBurnableToken *IBurnableTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IBurnableToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IBurnableToken *IBurnableTokenSession) Decimals() (uint8, error) {
	return _IBurnableToken.Contract.Decimals(&_IBurnableToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IBurnableToken *IBurnableTokenCallerSession) Decimals() (uint8, error) {
	return _IBurnableToken.Contract.Decimals(&_IBurnableToken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_IBurnableToken *IBurnableTokenCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBurnableToken.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_IBurnableToken *IBurnableTokenSession) GetOwner() (common.Address, error) {
	return _IBurnableToken.Contract.GetOwner(&_IBurnableToken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_IBurnableToken *IBurnableTokenCallerSession) GetOwner() (common.Address, error) {
	return _IBurnableToken.Contract.GetOwner(&_IBurnableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IBurnableToken *IBurnableTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IBurnableToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IBurnableToken *IBurnableTokenSession) Name() (string, error) {
	return _IBurnableToken.Contract.Name(&_IBurnableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IBurnableToken *IBurnableTokenCallerSession) Name() (string, error) {
	return _IBurnableToken.Contract.Name(&_IBurnableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IBurnableToken *IBurnableTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IBurnableToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IBurnableToken *IBurnableTokenSession) Symbol() (string, error) {
	return _IBurnableToken.Contract.Symbol(&_IBurnableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IBurnableToken *IBurnableTokenCallerSession) Symbol() (string, error) {
	return _IBurnableToken.Contract.Symbol(&_IBurnableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBurnableToken *IBurnableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBurnableToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBurnableToken *IBurnableTokenSession) TotalSupply() (*big.Int, error) {
	return _IBurnableToken.Contract.TotalSupply(&_IBurnableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBurnableToken *IBurnableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IBurnableToken.Contract.TotalSupply(&_IBurnableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.Contract.Approve(&_IBurnableToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.Contract.Approve(&_IBurnableToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.Contract.Burn(&_IBurnableToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.Contract.Burn(&_IBurnableToken.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address target, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenTransactor) Mint(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.contract.Transact(opts, "mint", target, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address target, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenSession) Mint(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.Contract.Mint(&_IBurnableToken.TransactOpts, target, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address target, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenTransactorSession) Mint(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.Contract.Mint(&_IBurnableToken.TransactOpts, target, amount)
}

// Mintable is a paid mutator transaction binding the contract method 0x4bf365df.
//
// Solidity: function mintable() returns(bool)
func (_IBurnableToken *IBurnableTokenTransactor) Mintable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBurnableToken.contract.Transact(opts, "mintable")
}

// Mintable is a paid mutator transaction binding the contract method 0x4bf365df.
//
// Solidity: function mintable() returns(bool)
func (_IBurnableToken *IBurnableTokenSession) Mintable() (*types.Transaction, error) {
	return _IBurnableToken.Contract.Mintable(&_IBurnableToken.TransactOpts)
}

// Mintable is a paid mutator transaction binding the contract method 0x4bf365df.
//
// Solidity: function mintable() returns(bool)
func (_IBurnableToken *IBurnableTokenTransactorSession) Mintable() (*types.Transaction, error) {
	return _IBurnableToken.Contract.Mintable(&_IBurnableToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.Contract.Transfer(&_IBurnableToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.Contract.Transfer(&_IBurnableToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.Contract.TransferFrom(&_IBurnableToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IBurnableToken *IBurnableTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBurnableToken.Contract.TransferFrom(&_IBurnableToken.TransactOpts, sender, recipient, amount)
}

// IBurnableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IBurnableToken contract.
type IBurnableTokenApprovalIterator struct {
	Event *IBurnableTokenApproval // Event containing the contract specifics and raw log

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
func (it *IBurnableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnableTokenApproval)
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
		it.Event = new(IBurnableTokenApproval)
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
func (it *IBurnableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnableTokenApproval represents a Approval event raised by the IBurnableToken contract.
type IBurnableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IBurnableToken *IBurnableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IBurnableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IBurnableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IBurnableTokenApprovalIterator{contract: _IBurnableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IBurnableToken *IBurnableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IBurnableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IBurnableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnableTokenApproval)
				if err := _IBurnableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IBurnableToken *IBurnableTokenFilterer) ParseApproval(log types.Log) (*IBurnableTokenApproval, error) {
	event := new(IBurnableTokenApproval)
	if err := _IBurnableToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBurnableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IBurnableToken contract.
type IBurnableTokenTransferIterator struct {
	Event *IBurnableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *IBurnableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBurnableTokenTransfer)
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
		it.Event = new(IBurnableTokenTransfer)
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
func (it *IBurnableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBurnableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBurnableTokenTransfer represents a Transfer event raised by the IBurnableToken contract.
type IBurnableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IBurnableToken *IBurnableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IBurnableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBurnableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IBurnableTokenTransferIterator{contract: _IBurnableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IBurnableToken *IBurnableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IBurnableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBurnableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBurnableTokenTransfer)
				if err := _IBurnableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IBurnableToken *IBurnableTokenFilterer) ParseTransfer(log types.Log) (*IBurnableTokenTransfer, error) {
	event := new(IBurnableTokenTransfer)
	if err := _IBurnableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
