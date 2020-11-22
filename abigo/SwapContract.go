// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// SwapContractABI is the input ABI used to generate the binding from.
const SwapContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"WBTC_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"churnedInCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeRewardsRatio\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"}],\"name\":\"singleTransferERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"_addressesAndAmounts\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"_inputDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"}],\"name\":\"multiTransferERC20TightlyPacked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_contributors\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"}],\"name\":\"multiTransferERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"collectSwapFeesForBTC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_addressesAndAmountOfFloat\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordIncomingFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"issueLPTokensForFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_addressesAndAmountOfLPtoken\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordOutcomingFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"burnLPTokensForFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeNodeRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"_nodeRewardsAddressAndAmounts\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"_churnedInCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_nodeRewardsRatio\",\"type\":\"uint8\"}],\"name\":\"churn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"isTxUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentPriceLP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SwapContract is an auto generated Go binding around an Ethereum contract.
type SwapContract struct {
	SwapContractCaller     // Read-only binding to the contract
	SwapContractTransactor // Write-only binding to the contract
	SwapContractFilterer   // Log filterer for contract events
}

// SwapContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapContractSession struct {
	Contract     *SwapContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapContractCallerSession struct {
	Contract *SwapContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SwapContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapContractTransactorSession struct {
	Contract     *SwapContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SwapContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapContractRaw struct {
	Contract *SwapContract // Generic contract binding to access the raw methods on
}

// SwapContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapContractCallerRaw struct {
	Contract *SwapContractCaller // Generic read-only contract binding to access the raw methods on
}

// SwapContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapContractTransactorRaw struct {
	Contract *SwapContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapContract creates a new instance of SwapContract, bound to a specific deployed contract.
func NewSwapContract(address common.Address, backend bind.ContractBackend) (*SwapContract, error) {
	contract, err := bindSwapContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapContract{SwapContractCaller: SwapContractCaller{contract: contract}, SwapContractTransactor: SwapContractTransactor{contract: contract}, SwapContractFilterer: SwapContractFilterer{contract: contract}}, nil
}

// NewSwapContractCaller creates a new read-only instance of SwapContract, bound to a specific deployed contract.
func NewSwapContractCaller(address common.Address, caller bind.ContractCaller) (*SwapContractCaller, error) {
	contract, err := bindSwapContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapContractCaller{contract: contract}, nil
}

// NewSwapContractTransactor creates a new write-only instance of SwapContract, bound to a specific deployed contract.
func NewSwapContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapContractTransactor, error) {
	contract, err := bindSwapContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapContractTransactor{contract: contract}, nil
}

// NewSwapContractFilterer creates a new log filterer instance of SwapContract, bound to a specific deployed contract.
func NewSwapContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapContractFilterer, error) {
	contract, err := bindSwapContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapContractFilterer{contract: contract}, nil
}

// bindSwapContract binds a generic wrapper to an already deployed contract.
func bindSwapContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapContract *SwapContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwapContract.Contract.SwapContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapContract *SwapContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapContract.Contract.SwapContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapContract *SwapContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapContract.Contract.SwapContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapContract *SwapContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwapContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapContract *SwapContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapContract *SwapContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapContract.Contract.contract.Transact(opts, method, params...)
}

// WBTCADDR is a free data retrieval call binding the contract method 0x4e4317d5.
//
// Solidity: function WBTC_ADDR() view returns(address)
func (_SwapContract *SwapContractCaller) WBTCADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwapContract.contract.Call(opts, out, "WBTC_ADDR")
	return *ret0, err
}

// WBTCADDR is a free data retrieval call binding the contract method 0x4e4317d5.
//
// Solidity: function WBTC_ADDR() view returns(address)
func (_SwapContract *SwapContractSession) WBTCADDR() (common.Address, error) {
	return _SwapContract.Contract.WBTCADDR(&_SwapContract.CallOpts)
}

// WBTCADDR is a free data retrieval call binding the contract method 0x4e4317d5.
//
// Solidity: function WBTC_ADDR() view returns(address)
func (_SwapContract *SwapContractCallerSession) WBTCADDR() (common.Address, error) {
	return _SwapContract.Contract.WBTCADDR(&_SwapContract.CallOpts)
}

// ChurnedInCount is a free data retrieval call binding the contract method 0x0089356f.
//
// Solidity: function churnedInCount() view returns(uint8)
func (_SwapContract *SwapContractCaller) ChurnedInCount(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SwapContract.contract.Call(opts, out, "churnedInCount")
	return *ret0, err
}

// ChurnedInCount is a free data retrieval call binding the contract method 0x0089356f.
//
// Solidity: function churnedInCount() view returns(uint8)
func (_SwapContract *SwapContractSession) ChurnedInCount() (uint8, error) {
	return _SwapContract.Contract.ChurnedInCount(&_SwapContract.CallOpts)
}

// ChurnedInCount is a free data retrieval call binding the contract method 0x0089356f.
//
// Solidity: function churnedInCount() view returns(uint8)
func (_SwapContract *SwapContractCallerSession) ChurnedInCount() (uint8, error) {
	return _SwapContract.Contract.ChurnedInCount(&_SwapContract.CallOpts)
}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256)
func (_SwapContract *SwapContractCaller) GetCurrentPriceLP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapContract.contract.Call(opts, out, "getCurrentPriceLP")
	return *ret0, err
}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256)
func (_SwapContract *SwapContractSession) GetCurrentPriceLP() (*big.Int, error) {
	return _SwapContract.Contract.GetCurrentPriceLP(&_SwapContract.CallOpts)
}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256)
func (_SwapContract *SwapContractCallerSession) GetCurrentPriceLP() (*big.Int, error) {
	return _SwapContract.Contract.GetCurrentPriceLP(&_SwapContract.CallOpts)
}

// IsTxUsed is a free data retrieval call binding the contract method 0xe6ca2084.
//
// Solidity: function isTxUsed(bytes32 _txid) view returns(bool)
func (_SwapContract *SwapContractCaller) IsTxUsed(opts *bind.CallOpts, _txid [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SwapContract.contract.Call(opts, out, "isTxUsed", _txid)
	return *ret0, err
}

// IsTxUsed is a free data retrieval call binding the contract method 0xe6ca2084.
//
// Solidity: function isTxUsed(bytes32 _txid) view returns(bool)
func (_SwapContract *SwapContractSession) IsTxUsed(_txid [32]byte) (bool, error) {
	return _SwapContract.Contract.IsTxUsed(&_SwapContract.CallOpts, _txid)
}

// IsTxUsed is a free data retrieval call binding the contract method 0xe6ca2084.
//
// Solidity: function isTxUsed(bytes32 _txid) view returns(bool)
func (_SwapContract *SwapContractCallerSession) IsTxUsed(_txid [32]byte) (bool, error) {
	return _SwapContract.Contract.IsTxUsed(&_SwapContract.CallOpts, _txid)
}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint8)
func (_SwapContract *SwapContractCaller) NodeRewardsRatio(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SwapContract.contract.Call(opts, out, "nodeRewardsRatio")
	return *ret0, err
}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint8)
func (_SwapContract *SwapContractSession) NodeRewardsRatio() (uint8, error) {
	return _SwapContract.Contract.NodeRewardsRatio(&_SwapContract.CallOpts)
}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint8)
func (_SwapContract *SwapContractCallerSession) NodeRewardsRatio() (uint8, error) {
	return _SwapContract.Contract.NodeRewardsRatio(&_SwapContract.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(bytes32)
func (_SwapContract *SwapContractCaller) Nodes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SwapContract.contract.Call(opts, out, "nodes", arg0)
	return *ret0, err
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(bytes32)
func (_SwapContract *SwapContractSession) Nodes(arg0 *big.Int) ([32]byte, error) {
	return _SwapContract.Contract.Nodes(&_SwapContract.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) view returns(bytes32)
func (_SwapContract *SwapContractCallerSession) Nodes(arg0 *big.Int) ([32]byte, error) {
	return _SwapContract.Contract.Nodes(&_SwapContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapContract *SwapContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SwapContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapContract *SwapContractSession) Owner() (common.Address, error) {
	return _SwapContract.Contract.Owner(&_SwapContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapContract *SwapContractCallerSession) Owner() (common.Address, error) {
	return _SwapContract.Contract.Owner(&_SwapContract.CallOpts)
}

// BurnLPTokensForFloat is a paid mutator transaction binding the contract method 0x1d351b5c.
//
// Solidity: function burnLPTokensForFloat(bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactor) BurnLPTokensForFloat(opts *bind.TransactOpts, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "burnLPTokensForFloat", _txid)
}

// BurnLPTokensForFloat is a paid mutator transaction binding the contract method 0x1d351b5c.
//
// Solidity: function burnLPTokensForFloat(bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractSession) BurnLPTokensForFloat(_txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.BurnLPTokensForFloat(&_SwapContract.TransactOpts, _txid)
}

// BurnLPTokensForFloat is a paid mutator transaction binding the contract method 0x1d351b5c.
//
// Solidity: function burnLPTokensForFloat(bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactorSession) BurnLPTokensForFloat(_txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.BurnLPTokensForFloat(&_SwapContract.TransactOpts, _txid)
}

// Churn is a paid mutator transaction binding the contract method 0x80a8ae30.
//
// Solidity: function churn(address _newOwner, bytes32[] _nodeRewardsAddressAndAmounts, uint8 _churnedInCount, uint8 _nodeRewardsRatio) returns(bool)
func (_SwapContract *SwapContractTransactor) Churn(opts *bind.TransactOpts, _newOwner common.Address, _nodeRewardsAddressAndAmounts [][32]byte, _churnedInCount uint8, _nodeRewardsRatio uint8) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "churn", _newOwner, _nodeRewardsAddressAndAmounts, _churnedInCount, _nodeRewardsRatio)
}

// Churn is a paid mutator transaction binding the contract method 0x80a8ae30.
//
// Solidity: function churn(address _newOwner, bytes32[] _nodeRewardsAddressAndAmounts, uint8 _churnedInCount, uint8 _nodeRewardsRatio) returns(bool)
func (_SwapContract *SwapContractSession) Churn(_newOwner common.Address, _nodeRewardsAddressAndAmounts [][32]byte, _churnedInCount uint8, _nodeRewardsRatio uint8) (*types.Transaction, error) {
	return _SwapContract.Contract.Churn(&_SwapContract.TransactOpts, _newOwner, _nodeRewardsAddressAndAmounts, _churnedInCount, _nodeRewardsRatio)
}

// Churn is a paid mutator transaction binding the contract method 0x80a8ae30.
//
// Solidity: function churn(address _newOwner, bytes32[] _nodeRewardsAddressAndAmounts, uint8 _churnedInCount, uint8 _nodeRewardsRatio) returns(bool)
func (_SwapContract *SwapContractTransactorSession) Churn(_newOwner common.Address, _nodeRewardsAddressAndAmounts [][32]byte, _churnedInCount uint8, _nodeRewardsRatio uint8) (*types.Transaction, error) {
	return _SwapContract.Contract.Churn(&_SwapContract.TransactOpts, _newOwner, _nodeRewardsAddressAndAmounts, _churnedInCount, _nodeRewardsRatio)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x1a4ed638.
//
// Solidity: function collectSwapFeesForBTC(address _feeToken, uint256 _rewardsAmount, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactor) CollectSwapFeesForBTC(opts *bind.TransactOpts, _feeToken common.Address, _rewardsAmount *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "collectSwapFeesForBTC", _feeToken, _rewardsAmount, _txid)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x1a4ed638.
//
// Solidity: function collectSwapFeesForBTC(address _feeToken, uint256 _rewardsAmount, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractSession) CollectSwapFeesForBTC(_feeToken common.Address, _rewardsAmount *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.CollectSwapFeesForBTC(&_SwapContract.TransactOpts, _feeToken, _rewardsAmount, _txid)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x1a4ed638.
//
// Solidity: function collectSwapFeesForBTC(address _feeToken, uint256 _rewardsAmount, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactorSession) CollectSwapFeesForBTC(_feeToken common.Address, _rewardsAmount *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.CollectSwapFeesForBTC(&_SwapContract.TransactOpts, _feeToken, _rewardsAmount, _txid)
}

// DistributeNodeRewards is a paid mutator transaction binding the contract method 0xe9e9bf6a.
//
// Solidity: function distributeNodeRewards() returns(bool)
func (_SwapContract *SwapContractTransactor) DistributeNodeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "distributeNodeRewards")
}

// DistributeNodeRewards is a paid mutator transaction binding the contract method 0xe9e9bf6a.
//
// Solidity: function distributeNodeRewards() returns(bool)
func (_SwapContract *SwapContractSession) DistributeNodeRewards() (*types.Transaction, error) {
	return _SwapContract.Contract.DistributeNodeRewards(&_SwapContract.TransactOpts)
}

// DistributeNodeRewards is a paid mutator transaction binding the contract method 0xe9e9bf6a.
//
// Solidity: function distributeNodeRewards() returns(bool)
func (_SwapContract *SwapContractTransactorSession) DistributeNodeRewards() (*types.Transaction, error) {
	return _SwapContract.Contract.DistributeNodeRewards(&_SwapContract.TransactOpts)
}

// IssueLPTokensForFloat is a paid mutator transaction binding the contract method 0xdb8beb0e.
//
// Solidity: function issueLPTokensForFloat(bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactor) IssueLPTokensForFloat(opts *bind.TransactOpts, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "issueLPTokensForFloat", _txid)
}

// IssueLPTokensForFloat is a paid mutator transaction binding the contract method 0xdb8beb0e.
//
// Solidity: function issueLPTokensForFloat(bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractSession) IssueLPTokensForFloat(_txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.IssueLPTokensForFloat(&_SwapContract.TransactOpts, _txid)
}

// IssueLPTokensForFloat is a paid mutator transaction binding the contract method 0xdb8beb0e.
//
// Solidity: function issueLPTokensForFloat(bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactorSession) IssueLPTokensForFloat(_txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.IssueLPTokensForFloat(&_SwapContract.TransactOpts, _txid)
}

// MultiTransferERC20 is a paid mutator transaction binding the contract method 0xac2aaa3c.
//
// Solidity: function multiTransferERC20(address _token, address[] _contributors, uint256[] _amounts, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractTransactor) MultiTransferERC20(opts *bind.TransactOpts, _token common.Address, _contributors []common.Address, _amounts []*big.Int, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "multiTransferERC20", _token, _contributors, _amounts, _rewardsAmount)
}

// MultiTransferERC20 is a paid mutator transaction binding the contract method 0xac2aaa3c.
//
// Solidity: function multiTransferERC20(address _token, address[] _contributors, uint256[] _amounts, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractSession) MultiTransferERC20(_token common.Address, _contributors []common.Address, _amounts []*big.Int, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.MultiTransferERC20(&_SwapContract.TransactOpts, _token, _contributors, _amounts, _rewardsAmount)
}

// MultiTransferERC20 is a paid mutator transaction binding the contract method 0xac2aaa3c.
//
// Solidity: function multiTransferERC20(address _token, address[] _contributors, uint256[] _amounts, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractTransactorSession) MultiTransferERC20(_token common.Address, _contributors []common.Address, _amounts []*big.Int, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.MultiTransferERC20(&_SwapContract.TransactOpts, _token, _contributors, _amounts, _rewardsAmount)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad833dee.
//
// Solidity: function multiTransferERC20TightlyPacked(address _token, bytes32[] _addressesAndAmounts, uint8 _inputDecimals, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractTransactor) MultiTransferERC20TightlyPacked(opts *bind.TransactOpts, _token common.Address, _addressesAndAmounts [][32]byte, _inputDecimals uint8, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "multiTransferERC20TightlyPacked", _token, _addressesAndAmounts, _inputDecimals, _rewardsAmount)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad833dee.
//
// Solidity: function multiTransferERC20TightlyPacked(address _token, bytes32[] _addressesAndAmounts, uint8 _inputDecimals, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractSession) MultiTransferERC20TightlyPacked(_token common.Address, _addressesAndAmounts [][32]byte, _inputDecimals uint8, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.MultiTransferERC20TightlyPacked(&_SwapContract.TransactOpts, _token, _addressesAndAmounts, _inputDecimals, _rewardsAmount)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad833dee.
//
// Solidity: function multiTransferERC20TightlyPacked(address _token, bytes32[] _addressesAndAmounts, uint8 _inputDecimals, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractTransactorSession) MultiTransferERC20TightlyPacked(_token common.Address, _addressesAndAmounts [][32]byte, _inputDecimals uint8, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.MultiTransferERC20TightlyPacked(&_SwapContract.TransactOpts, _token, _addressesAndAmounts, _inputDecimals, _rewardsAmount)
}

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0xcf10b16b.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactor) RecordIncomingFloat(opts *bind.TransactOpts, _token common.Address, _addressesAndAmountOfFloat [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "recordIncomingFloat", _token, _addressesAndAmountOfFloat, _txid)
}

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0xcf10b16b.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractSession) RecordIncomingFloat(_token common.Address, _addressesAndAmountOfFloat [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordIncomingFloat(&_SwapContract.TransactOpts, _token, _addressesAndAmountOfFloat, _txid)
}

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0xcf10b16b.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactorSession) RecordIncomingFloat(_token common.Address, _addressesAndAmountOfFloat [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordIncomingFloat(&_SwapContract.TransactOpts, _token, _addressesAndAmountOfFloat, _txid)
}

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0xa1e271d3.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactor) RecordOutcomingFloat(opts *bind.TransactOpts, _token common.Address, _addressesAndAmountOfLPtoken [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "recordOutcomingFloat", _token, _addressesAndAmountOfLPtoken, _txid)
}

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0xa1e271d3.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractSession) RecordOutcomingFloat(_token common.Address, _addressesAndAmountOfLPtoken [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordOutcomingFloat(&_SwapContract.TransactOpts, _token, _addressesAndAmountOfLPtoken, _txid)
}

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0xa1e271d3.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactorSession) RecordOutcomingFloat(_token common.Address, _addressesAndAmountOfLPtoken [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordOutcomingFloat(&_SwapContract.TransactOpts, _token, _addressesAndAmountOfLPtoken, _txid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapContract *SwapContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapContract *SwapContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapContract.Contract.RenounceOwnership(&_SwapContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapContract *SwapContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapContract.Contract.RenounceOwnership(&_SwapContract.TransactOpts)
}

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x37ea3269.
//
// Solidity: function singleTransferERC20(address _token, address _to, uint256 _amount, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractTransactor) SingleTransferERC20(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "singleTransferERC20", _token, _to, _amount, _rewardsAmount)
}

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x37ea3269.
//
// Solidity: function singleTransferERC20(address _token, address _to, uint256 _amount, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractSession) SingleTransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.SingleTransferERC20(&_SwapContract.TransactOpts, _token, _to, _amount, _rewardsAmount)
}

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x37ea3269.
//
// Solidity: function singleTransferERC20(address _token, address _to, uint256 _amount, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractTransactorSession) SingleTransferERC20(_token common.Address, _to common.Address, _amount *big.Int, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.SingleTransferERC20(&_SwapContract.TransactOpts, _token, _to, _amount, _rewardsAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapContract *SwapContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapContract *SwapContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapContract.Contract.TransferOwnership(&_SwapContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapContract *SwapContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapContract.Contract.TransferOwnership(&_SwapContract.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SwapContract *SwapContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SwapContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SwapContract *SwapContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SwapContract.Contract.Fallback(&_SwapContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SwapContract *SwapContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SwapContract.Contract.Fallback(&_SwapContract.TransactOpts, calldata)
}

// SwapContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SwapContract contract.
type SwapContractOwnershipTransferredIterator struct {
	Event *SwapContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SwapContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractOwnershipTransferred)
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
		it.Event = new(SwapContractOwnershipTransferred)
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
func (it *SwapContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractOwnershipTransferred represents a OwnershipTransferred event raised by the SwapContract contract.
type SwapContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapContract *SwapContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapContractOwnershipTransferredIterator{contract: _SwapContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapContract *SwapContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractOwnershipTransferred)
				if err := _SwapContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SwapContract *SwapContractFilterer) ParseOwnershipTransferred(log types.Log) (*SwapContractOwnershipTransferred, error) {
	event := new(SwapContractOwnershipTransferred)
	if err := _SwapContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
