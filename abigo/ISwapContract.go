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

// ISwapContractMetaData contains all meta data concerning the ISwapContract contract.
var ISwapContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BTCT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_nodes\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_isRemoved\",\"type\":\"bool[]\"},{\"internalType\":\"uint8\",\"name\":\"_churnedInCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tssThreshold\",\"type\":\"uint8\"}],\"name\":\"churn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_incomingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_spenders\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_swapAmounts\",\"type\":\"uint256[]\"}],\"name\":\"collectSwapFeesForBTC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentPriceLP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"}],\"name\":\"getFloatReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isNodeStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"isTxUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIBurnableToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"_addressesAndAmounts\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalSwapped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_redeemedFloatTxIds\",\"type\":\"bytes32[]\"}],\"name\":\"multiTransferERC20TightlyPacked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_addressesAndAmountOfFloat\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordIncomingFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_addressesAndAmountOfLPtoken\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_minerFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordOutcomingFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minerFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordUTXOSweepMinerFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSwapped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_redeemedFloatTxIds\",\"type\":\"bytes32[]\"}],\"name\":\"singleTransferERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sbBTCPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buybackAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalFeeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeRewardsRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_buybackRewardsRatio\",\"type\":\"uint256\"}],\"name\":\"updateParams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISwapContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwapContractMetaData.ABI instead.
var ISwapContractABI = ISwapContractMetaData.ABI

// ISwapContract is an auto generated Go binding around an Ethereum contract.
type ISwapContract struct {
	ISwapContractCaller     // Read-only binding to the contract
	ISwapContractTransactor // Write-only binding to the contract
	ISwapContractFilterer   // Log filterer for contract events
}

// ISwapContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapContractSession struct {
	Contract     *ISwapContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapContractCallerSession struct {
	Contract *ISwapContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISwapContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapContractTransactorSession struct {
	Contract     *ISwapContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISwapContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapContractRaw struct {
	Contract *ISwapContract // Generic contract binding to access the raw methods on
}

// ISwapContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapContractCallerRaw struct {
	Contract *ISwapContractCaller // Generic read-only contract binding to access the raw methods on
}

// ISwapContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapContractTransactorRaw struct {
	Contract *ISwapContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwapContract creates a new instance of ISwapContract, bound to a specific deployed contract.
func NewISwapContract(address common.Address, backend bind.ContractBackend) (*ISwapContract, error) {
	contract, err := bindISwapContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwapContract{ISwapContractCaller: ISwapContractCaller{contract: contract}, ISwapContractTransactor: ISwapContractTransactor{contract: contract}, ISwapContractFilterer: ISwapContractFilterer{contract: contract}}, nil
}

// NewISwapContractCaller creates a new read-only instance of ISwapContract, bound to a specific deployed contract.
func NewISwapContractCaller(address common.Address, caller bind.ContractCaller) (*ISwapContractCaller, error) {
	contract, err := bindISwapContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapContractCaller{contract: contract}, nil
}

// NewISwapContractTransactor creates a new write-only instance of ISwapContract, bound to a specific deployed contract.
func NewISwapContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwapContractTransactor, error) {
	contract, err := bindISwapContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapContractTransactor{contract: contract}, nil
}

// NewISwapContractFilterer creates a new log filterer instance of ISwapContract, bound to a specific deployed contract.
func NewISwapContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwapContractFilterer, error) {
	contract, err := bindISwapContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapContractFilterer{contract: contract}, nil
}

// bindISwapContract binds a generic wrapper to an already deployed contract.
func bindISwapContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISwapContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapContract *ISwapContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapContract.Contract.ISwapContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapContract *ISwapContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapContract.Contract.ISwapContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapContract *ISwapContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapContract.Contract.ISwapContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapContract *ISwapContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapContract *ISwapContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapContract *ISwapContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapContract.Contract.contract.Transact(opts, method, params...)
}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_ISwapContract *ISwapContractCaller) GetActiveNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ISwapContract.contract.Call(opts, &out, "getActiveNodes")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_ISwapContract *ISwapContractSession) GetActiveNodes() ([]common.Address, error) {
	return _ISwapContract.Contract.GetActiveNodes(&_ISwapContract.CallOpts)
}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_ISwapContract *ISwapContractCallerSession) GetActiveNodes() ([]common.Address, error) {
	return _ISwapContract.Contract.GetActiveNodes(&_ISwapContract.CallOpts)
}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256)
func (_ISwapContract *ISwapContractCaller) GetCurrentPriceLP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwapContract.contract.Call(opts, &out, "getCurrentPriceLP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256)
func (_ISwapContract *ISwapContractSession) GetCurrentPriceLP() (*big.Int, error) {
	return _ISwapContract.Contract.GetCurrentPriceLP(&_ISwapContract.CallOpts)
}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256)
func (_ISwapContract *ISwapContractCallerSession) GetCurrentPriceLP() (*big.Int, error) {
	return _ISwapContract.Contract.GetCurrentPriceLP(&_ISwapContract.CallOpts)
}

// IsTxUsed is a free data retrieval call binding the contract method 0xe6ca2084.
//
// Solidity: function isTxUsed(bytes32 _txid) view returns(bool)
func (_ISwapContract *ISwapContractCaller) IsTxUsed(opts *bind.CallOpts, _txid [32]byte) (bool, error) {
	var out []interface{}
	err := _ISwapContract.contract.Call(opts, &out, "isTxUsed", _txid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTxUsed is a free data retrieval call binding the contract method 0xe6ca2084.
//
// Solidity: function isTxUsed(bytes32 _txid) view returns(bool)
func (_ISwapContract *ISwapContractSession) IsTxUsed(_txid [32]byte) (bool, error) {
	return _ISwapContract.Contract.IsTxUsed(&_ISwapContract.CallOpts, _txid)
}

// IsTxUsed is a free data retrieval call binding the contract method 0xe6ca2084.
//
// Solidity: function isTxUsed(bytes32 _txid) view returns(bool)
func (_ISwapContract *ISwapContractCallerSession) IsTxUsed(_txid [32]byte) (bool, error) {
	return _ISwapContract.Contract.IsTxUsed(&_ISwapContract.CallOpts, _txid)
}

// BTCTADDR is a paid mutator transaction binding the contract method 0x0f909486.
//
// Solidity: function BTCT_ADDR() returns(address)
func (_ISwapContract *ISwapContractTransactor) BTCTADDR(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "BTCT_ADDR")
}

// BTCTADDR is a paid mutator transaction binding the contract method 0x0f909486.
//
// Solidity: function BTCT_ADDR() returns(address)
func (_ISwapContract *ISwapContractSession) BTCTADDR() (*types.Transaction, error) {
	return _ISwapContract.Contract.BTCTADDR(&_ISwapContract.TransactOpts)
}

// BTCTADDR is a paid mutator transaction binding the contract method 0x0f909486.
//
// Solidity: function BTCT_ADDR() returns(address)
func (_ISwapContract *ISwapContractTransactorSession) BTCTADDR() (*types.Transaction, error) {
	return _ISwapContract.Contract.BTCTADDR(&_ISwapContract.TransactOpts)
}

// Churn is a paid mutator transaction binding the contract method 0x6845a025.
//
// Solidity: function churn(address _newOwner, address[] _nodes, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold) returns(bool)
func (_ISwapContract *ISwapContractTransactor) Churn(opts *bind.TransactOpts, _newOwner common.Address, _nodes []common.Address, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "churn", _newOwner, _nodes, _isRemoved, _churnedInCount, _tssThreshold)
}

// Churn is a paid mutator transaction binding the contract method 0x6845a025.
//
// Solidity: function churn(address _newOwner, address[] _nodes, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold) returns(bool)
func (_ISwapContract *ISwapContractSession) Churn(_newOwner common.Address, _nodes []common.Address, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8) (*types.Transaction, error) {
	return _ISwapContract.Contract.Churn(&_ISwapContract.TransactOpts, _newOwner, _nodes, _isRemoved, _churnedInCount, _tssThreshold)
}

// Churn is a paid mutator transaction binding the contract method 0x6845a025.
//
// Solidity: function churn(address _newOwner, address[] _nodes, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold) returns(bool)
func (_ISwapContract *ISwapContractTransactorSession) Churn(_newOwner common.Address, _nodes []common.Address, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8) (*types.Transaction, error) {
	return _ISwapContract.Contract.Churn(&_ISwapContract.TransactOpts, _newOwner, _nodes, _isRemoved, _churnedInCount, _tssThreshold)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x2adf9f87.
//
// Solidity: function collectSwapFeesForBTC(uint256 _incomingAmount, uint256 _minerFee, uint256 _rewardsAmount, address[] _spenders, uint256[] _swapAmounts) returns(bool)
func (_ISwapContract *ISwapContractTransactor) CollectSwapFeesForBTC(opts *bind.TransactOpts, _incomingAmount *big.Int, _minerFee *big.Int, _rewardsAmount *big.Int, _spenders []common.Address, _swapAmounts []*big.Int) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "collectSwapFeesForBTC", _incomingAmount, _minerFee, _rewardsAmount, _spenders, _swapAmounts)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x2adf9f87.
//
// Solidity: function collectSwapFeesForBTC(uint256 _incomingAmount, uint256 _minerFee, uint256 _rewardsAmount, address[] _spenders, uint256[] _swapAmounts) returns(bool)
func (_ISwapContract *ISwapContractSession) CollectSwapFeesForBTC(_incomingAmount *big.Int, _minerFee *big.Int, _rewardsAmount *big.Int, _spenders []common.Address, _swapAmounts []*big.Int) (*types.Transaction, error) {
	return _ISwapContract.Contract.CollectSwapFeesForBTC(&_ISwapContract.TransactOpts, _incomingAmount, _minerFee, _rewardsAmount, _spenders, _swapAmounts)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x2adf9f87.
//
// Solidity: function collectSwapFeesForBTC(uint256 _incomingAmount, uint256 _minerFee, uint256 _rewardsAmount, address[] _spenders, uint256[] _swapAmounts) returns(bool)
func (_ISwapContract *ISwapContractTransactorSession) CollectSwapFeesForBTC(_incomingAmount *big.Int, _minerFee *big.Int, _rewardsAmount *big.Int, _spenders []common.Address, _swapAmounts []*big.Int) (*types.Transaction, error) {
	return _ISwapContract.Contract.CollectSwapFeesForBTC(&_ISwapContract.TransactOpts, _incomingAmount, _minerFee, _rewardsAmount, _spenders, _swapAmounts)
}

// GetFloatReserve is a paid mutator transaction binding the contract method 0xec482729.
//
// Solidity: function getFloatReserve(address _tokenA, address _tokenB) returns(uint256 reserveA, uint256 reserveB)
func (_ISwapContract *ISwapContractTransactor) GetFloatReserve(opts *bind.TransactOpts, _tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "getFloatReserve", _tokenA, _tokenB)
}

// GetFloatReserve is a paid mutator transaction binding the contract method 0xec482729.
//
// Solidity: function getFloatReserve(address _tokenA, address _tokenB) returns(uint256 reserveA, uint256 reserveB)
func (_ISwapContract *ISwapContractSession) GetFloatReserve(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _ISwapContract.Contract.GetFloatReserve(&_ISwapContract.TransactOpts, _tokenA, _tokenB)
}

// GetFloatReserve is a paid mutator transaction binding the contract method 0xec482729.
//
// Solidity: function getFloatReserve(address _tokenA, address _tokenB) returns(uint256 reserveA, uint256 reserveB)
func (_ISwapContract *ISwapContractTransactorSession) GetFloatReserve(_tokenA common.Address, _tokenB common.Address) (*types.Transaction, error) {
	return _ISwapContract.Contract.GetFloatReserve(&_ISwapContract.TransactOpts, _tokenA, _tokenB)
}

// IsNodeStake is a paid mutator transaction binding the contract method 0xa742329d.
//
// Solidity: function isNodeStake(address _user) returns(bool)
func (_ISwapContract *ISwapContractTransactor) IsNodeStake(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "isNodeStake", _user)
}

// IsNodeStake is a paid mutator transaction binding the contract method 0xa742329d.
//
// Solidity: function isNodeStake(address _user) returns(bool)
func (_ISwapContract *ISwapContractSession) IsNodeStake(_user common.Address) (*types.Transaction, error) {
	return _ISwapContract.Contract.IsNodeStake(&_ISwapContract.TransactOpts, _user)
}

// IsNodeStake is a paid mutator transaction binding the contract method 0xa742329d.
//
// Solidity: function isNodeStake(address _user) returns(bool)
func (_ISwapContract *ISwapContractTransactorSession) IsNodeStake(_user common.Address) (*types.Transaction, error) {
	return _ISwapContract.Contract.IsNodeStake(&_ISwapContract.TransactOpts, _user)
}

// LpToken is a paid mutator transaction binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() returns(address)
func (_ISwapContract *ISwapContractTransactor) LpToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "lpToken")
}

// LpToken is a paid mutator transaction binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() returns(address)
func (_ISwapContract *ISwapContractSession) LpToken() (*types.Transaction, error) {
	return _ISwapContract.Contract.LpToken(&_ISwapContract.TransactOpts)
}

// LpToken is a paid mutator transaction binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() returns(address)
func (_ISwapContract *ISwapContractTransactorSession) LpToken() (*types.Transaction, error) {
	return _ISwapContract.Contract.LpToken(&_ISwapContract.TransactOpts)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad289e76.
//
// Solidity: function multiTransferERC20TightlyPacked(address _destToken, bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_ISwapContract *ISwapContractTransactor) MultiTransferERC20TightlyPacked(opts *bind.TransactOpts, _destToken common.Address, _addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "multiTransferERC20TightlyPacked", _destToken, _addressesAndAmounts, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad289e76.
//
// Solidity: function multiTransferERC20TightlyPacked(address _destToken, bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_ISwapContract *ISwapContractSession) MultiTransferERC20TightlyPacked(_destToken common.Address, _addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _ISwapContract.Contract.MultiTransferERC20TightlyPacked(&_ISwapContract.TransactOpts, _destToken, _addressesAndAmounts, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad289e76.
//
// Solidity: function multiTransferERC20TightlyPacked(address _destToken, bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_ISwapContract *ISwapContractTransactorSession) MultiTransferERC20TightlyPacked(_destToken common.Address, _addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _ISwapContract.Contract.MultiTransferERC20TightlyPacked(&_ISwapContract.TransactOpts, _destToken, _addressesAndAmounts, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0xcf10b16b.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bytes32 _txid) returns(bool)
func (_ISwapContract *ISwapContractTransactor) RecordIncomingFloat(opts *bind.TransactOpts, _token common.Address, _addressesAndAmountOfFloat [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "recordIncomingFloat", _token, _addressesAndAmountOfFloat, _txid)
}

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0xcf10b16b.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bytes32 _txid) returns(bool)
func (_ISwapContract *ISwapContractSession) RecordIncomingFloat(_token common.Address, _addressesAndAmountOfFloat [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _ISwapContract.Contract.RecordIncomingFloat(&_ISwapContract.TransactOpts, _token, _addressesAndAmountOfFloat, _txid)
}

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0xcf10b16b.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bytes32 _txid) returns(bool)
func (_ISwapContract *ISwapContractTransactorSession) RecordIncomingFloat(_token common.Address, _addressesAndAmountOfFloat [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _ISwapContract.Contract.RecordIncomingFloat(&_ISwapContract.TransactOpts, _token, _addressesAndAmountOfFloat, _txid)
}

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0x2586c562.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, uint256 _minerFee, bytes32 _txid) returns(bool)
func (_ISwapContract *ISwapContractTransactor) RecordOutcomingFloat(opts *bind.TransactOpts, _token common.Address, _addressesAndAmountOfLPtoken [32]byte, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "recordOutcomingFloat", _token, _addressesAndAmountOfLPtoken, _minerFee, _txid)
}

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0x2586c562.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, uint256 _minerFee, bytes32 _txid) returns(bool)
func (_ISwapContract *ISwapContractSession) RecordOutcomingFloat(_token common.Address, _addressesAndAmountOfLPtoken [32]byte, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _ISwapContract.Contract.RecordOutcomingFloat(&_ISwapContract.TransactOpts, _token, _addressesAndAmountOfLPtoken, _minerFee, _txid)
}

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0x2586c562.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, uint256 _minerFee, bytes32 _txid) returns(bool)
func (_ISwapContract *ISwapContractTransactorSession) RecordOutcomingFloat(_token common.Address, _addressesAndAmountOfLPtoken [32]byte, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _ISwapContract.Contract.RecordOutcomingFloat(&_ISwapContract.TransactOpts, _token, _addressesAndAmountOfLPtoken, _minerFee, _txid)
}

// RecordUTXOSweepMinerFee is a paid mutator transaction binding the contract method 0xc810a539.
//
// Solidity: function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid) returns(bool)
func (_ISwapContract *ISwapContractTransactor) RecordUTXOSweepMinerFee(opts *bind.TransactOpts, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "recordUTXOSweepMinerFee", _minerFee, _txid)
}

// RecordUTXOSweepMinerFee is a paid mutator transaction binding the contract method 0xc810a539.
//
// Solidity: function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid) returns(bool)
func (_ISwapContract *ISwapContractSession) RecordUTXOSweepMinerFee(_minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _ISwapContract.Contract.RecordUTXOSweepMinerFee(&_ISwapContract.TransactOpts, _minerFee, _txid)
}

// RecordUTXOSweepMinerFee is a paid mutator transaction binding the contract method 0xc810a539.
//
// Solidity: function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid) returns(bool)
func (_ISwapContract *ISwapContractTransactorSession) RecordUTXOSweepMinerFee(_minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _ISwapContract.Contract.RecordUTXOSweepMinerFee(&_ISwapContract.TransactOpts, _minerFee, _txid)
}

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x0d63aca7.
//
// Solidity: function singleTransferERC20(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_ISwapContract *ISwapContractTransactor) SingleTransferERC20(opts *bind.TransactOpts, _destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "singleTransferERC20", _destToken, _to, _amount, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x0d63aca7.
//
// Solidity: function singleTransferERC20(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_ISwapContract *ISwapContractSession) SingleTransferERC20(_destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _ISwapContract.Contract.SingleTransferERC20(&_ISwapContract.TransactOpts, _destToken, _to, _amount, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x0d63aca7.
//
// Solidity: function singleTransferERC20(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_ISwapContract *ISwapContractTransactorSession) SingleTransferERC20(_destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _ISwapContract.Contract.SingleTransferERC20(&_ISwapContract.TransactOpts, _destToken, _to, _amount, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xe880afd4.
//
// Solidity: function updateParams(address _sbBTCPool, address _buybackAddress, uint256 _withdrawalFeeBPS, uint256 _nodeRewardsRatio, uint256 _buybackRewardsRatio) returns(bool)
func (_ISwapContract *ISwapContractTransactor) UpdateParams(opts *bind.TransactOpts, _sbBTCPool common.Address, _buybackAddress common.Address, _withdrawalFeeBPS *big.Int, _nodeRewardsRatio *big.Int, _buybackRewardsRatio *big.Int) (*types.Transaction, error) {
	return _ISwapContract.contract.Transact(opts, "updateParams", _sbBTCPool, _buybackAddress, _withdrawalFeeBPS, _nodeRewardsRatio, _buybackRewardsRatio)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xe880afd4.
//
// Solidity: function updateParams(address _sbBTCPool, address _buybackAddress, uint256 _withdrawalFeeBPS, uint256 _nodeRewardsRatio, uint256 _buybackRewardsRatio) returns(bool)
func (_ISwapContract *ISwapContractSession) UpdateParams(_sbBTCPool common.Address, _buybackAddress common.Address, _withdrawalFeeBPS *big.Int, _nodeRewardsRatio *big.Int, _buybackRewardsRatio *big.Int) (*types.Transaction, error) {
	return _ISwapContract.Contract.UpdateParams(&_ISwapContract.TransactOpts, _sbBTCPool, _buybackAddress, _withdrawalFeeBPS, _nodeRewardsRatio, _buybackRewardsRatio)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xe880afd4.
//
// Solidity: function updateParams(address _sbBTCPool, address _buybackAddress, uint256 _withdrawalFeeBPS, uint256 _nodeRewardsRatio, uint256 _buybackRewardsRatio) returns(bool)
func (_ISwapContract *ISwapContractTransactorSession) UpdateParams(_sbBTCPool common.Address, _buybackAddress common.Address, _withdrawalFeeBPS *big.Int, _nodeRewardsRatio *big.Int, _buybackRewardsRatio *big.Int) (*types.Transaction, error) {
	return _ISwapContract.Contract.UpdateParams(&_ISwapContract.TransactOpts, _sbBTCPool, _buybackAddress, _withdrawalFeeBPS, _nodeRewardsRatio, _buybackRewardsRatio)
}
