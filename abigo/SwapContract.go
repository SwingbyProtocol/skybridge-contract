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

// SwapContractMetaData contains all meta data concerning the SwapContract contract.
var SwapContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_btct\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sbBTCPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swapRewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buybackAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initBTCFloat\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initWBTCFloat\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfFloat\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPriceLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"}],\"name\":\"BurnLPTokensForFloat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardLPTsForNodes\",\"type\":\"uint256\"}],\"name\":\"DistributeFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfFloat\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPriceLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"}],\"name\":\"IssueLPTokensForFloat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feesToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsLPTTotal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPriceLP\",\"type\":\"uint256\"}],\"name\":\"RewardsCollection\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BTCT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeNodeCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buybackAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buybackRewardsRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_nodes\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_isRemoved\",\"type\":\"bool[]\"},{\"internalType\":\"uint8\",\"name\":\"_churnedInCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tssThreshold\",\"type\":\"uint8\"}],\"name\":\"churn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"churnedInCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_incomingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_spenders\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_swapAmounts\",\"type\":\"uint256[]\"}],\"name\":\"collectSwapFeesForBTC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentPriceLP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nowPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"}],\"name\":\"getFloatReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isNodeStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"isTxUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIBurnableToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"_addressesAndAmounts\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalSwapped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_redeemedFloatTxIds\",\"type\":\"bytes32[]\"}],\"name\":\"multiTransferERC20TightlyPacked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeRewardsRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_addressesAndAmountOfFloat\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordIncomingFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_addressesAndAmountOfLPtoken\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_minerFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordOutcomingFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minerFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordUTXOSweepMinerFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sbBTCPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSwapped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_redeemedFloatTxIds\",\"type\":\"bytes32[]\"}],\"name\":\"singleTransferERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sw\",\"outputs\":[{\"internalType\":\"contractISwapRewards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sbBTCPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buybackAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalFeeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeRewardsRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_buybackRewardsRatio\",\"type\":\"uint256\"}],\"name\":\"updateParams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalFeeBPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SwapContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapContractMetaData.ABI instead.
var SwapContractABI = SwapContractMetaData.ABI

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
func (_SwapContract *SwapContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_SwapContract *SwapContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// BTCTADDR is a free data retrieval call binding the contract method 0x0f909486.
//
// Solidity: function BTCT_ADDR() view returns(address)
func (_SwapContract *SwapContractCaller) BTCTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "BTCT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BTCTADDR is a free data retrieval call binding the contract method 0x0f909486.
//
// Solidity: function BTCT_ADDR() view returns(address)
func (_SwapContract *SwapContractSession) BTCTADDR() (common.Address, error) {
	return _SwapContract.Contract.BTCTADDR(&_SwapContract.CallOpts)
}

// BTCTADDR is a free data retrieval call binding the contract method 0x0f909486.
//
// Solidity: function BTCT_ADDR() view returns(address)
func (_SwapContract *SwapContractCallerSession) BTCTADDR() (common.Address, error) {
	return _SwapContract.Contract.BTCTADDR(&_SwapContract.CallOpts)
}

// ActiveNodeCount is a free data retrieval call binding the contract method 0x75340815.
//
// Solidity: function activeNodeCount() view returns(uint8)
func (_SwapContract *SwapContractCaller) ActiveNodeCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "activeNodeCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ActiveNodeCount is a free data retrieval call binding the contract method 0x75340815.
//
// Solidity: function activeNodeCount() view returns(uint8)
func (_SwapContract *SwapContractSession) ActiveNodeCount() (uint8, error) {
	return _SwapContract.Contract.ActiveNodeCount(&_SwapContract.CallOpts)
}

// ActiveNodeCount is a free data retrieval call binding the contract method 0x75340815.
//
// Solidity: function activeNodeCount() view returns(uint8)
func (_SwapContract *SwapContractCallerSession) ActiveNodeCount() (uint8, error) {
	return _SwapContract.Contract.ActiveNodeCount(&_SwapContract.CallOpts)
}

// BuybackAddress is a free data retrieval call binding the contract method 0xcc2fbd66.
//
// Solidity: function buybackAddress() view returns(address)
func (_SwapContract *SwapContractCaller) BuybackAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "buybackAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BuybackAddress is a free data retrieval call binding the contract method 0xcc2fbd66.
//
// Solidity: function buybackAddress() view returns(address)
func (_SwapContract *SwapContractSession) BuybackAddress() (common.Address, error) {
	return _SwapContract.Contract.BuybackAddress(&_SwapContract.CallOpts)
}

// BuybackAddress is a free data retrieval call binding the contract method 0xcc2fbd66.
//
// Solidity: function buybackAddress() view returns(address)
func (_SwapContract *SwapContractCallerSession) BuybackAddress() (common.Address, error) {
	return _SwapContract.Contract.BuybackAddress(&_SwapContract.CallOpts)
}

// BuybackRewardsRatio is a free data retrieval call binding the contract method 0x0b513f1d.
//
// Solidity: function buybackRewardsRatio() view returns(uint256)
func (_SwapContract *SwapContractCaller) BuybackRewardsRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "buybackRewardsRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuybackRewardsRatio is a free data retrieval call binding the contract method 0x0b513f1d.
//
// Solidity: function buybackRewardsRatio() view returns(uint256)
func (_SwapContract *SwapContractSession) BuybackRewardsRatio() (*big.Int, error) {
	return _SwapContract.Contract.BuybackRewardsRatio(&_SwapContract.CallOpts)
}

// BuybackRewardsRatio is a free data retrieval call binding the contract method 0x0b513f1d.
//
// Solidity: function buybackRewardsRatio() view returns(uint256)
func (_SwapContract *SwapContractCallerSession) BuybackRewardsRatio() (*big.Int, error) {
	return _SwapContract.Contract.BuybackRewardsRatio(&_SwapContract.CallOpts)
}

// ChurnedInCount is a free data retrieval call binding the contract method 0x0089356f.
//
// Solidity: function churnedInCount() view returns(uint8)
func (_SwapContract *SwapContractCaller) ChurnedInCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "churnedInCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

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

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_SwapContract *SwapContractCaller) GetActiveNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "getActiveNodes")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_SwapContract *SwapContractSession) GetActiveNodes() ([]common.Address, error) {
	return _SwapContract.Contract.GetActiveNodes(&_SwapContract.CallOpts)
}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_SwapContract *SwapContractCallerSession) GetActiveNodes() ([]common.Address, error) {
	return _SwapContract.Contract.GetActiveNodes(&_SwapContract.CallOpts)
}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256 nowPrice)
func (_SwapContract *SwapContractCaller) GetCurrentPriceLP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "getCurrentPriceLP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256 nowPrice)
func (_SwapContract *SwapContractSession) GetCurrentPriceLP() (*big.Int, error) {
	return _SwapContract.Contract.GetCurrentPriceLP(&_SwapContract.CallOpts)
}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256 nowPrice)
func (_SwapContract *SwapContractCallerSession) GetCurrentPriceLP() (*big.Int, error) {
	return _SwapContract.Contract.GetCurrentPriceLP(&_SwapContract.CallOpts)
}

// GetFloatReserve is a free data retrieval call binding the contract method 0xec482729.
//
// Solidity: function getFloatReserve(address _tokenA, address _tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_SwapContract *SwapContractCaller) GetFloatReserve(opts *bind.CallOpts, _tokenA common.Address, _tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "getFloatReserve", _tokenA, _tokenB)

	outstruct := new(struct {
		ReserveA *big.Int
		ReserveB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFloatReserve is a free data retrieval call binding the contract method 0xec482729.
//
// Solidity: function getFloatReserve(address _tokenA, address _tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_SwapContract *SwapContractSession) GetFloatReserve(_tokenA common.Address, _tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _SwapContract.Contract.GetFloatReserve(&_SwapContract.CallOpts, _tokenA, _tokenB)
}

// GetFloatReserve is a free data retrieval call binding the contract method 0xec482729.
//
// Solidity: function getFloatReserve(address _tokenA, address _tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_SwapContract *SwapContractCallerSession) GetFloatReserve(_tokenA common.Address, _tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _SwapContract.Contract.GetFloatReserve(&_SwapContract.CallOpts, _tokenA, _tokenB)
}

// IsNodeStake is a free data retrieval call binding the contract method 0xa742329d.
//
// Solidity: function isNodeStake(address _user) view returns(bool)
func (_SwapContract *SwapContractCaller) IsNodeStake(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "isNodeStake", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNodeStake is a free data retrieval call binding the contract method 0xa742329d.
//
// Solidity: function isNodeStake(address _user) view returns(bool)
func (_SwapContract *SwapContractSession) IsNodeStake(_user common.Address) (bool, error) {
	return _SwapContract.Contract.IsNodeStake(&_SwapContract.CallOpts, _user)
}

// IsNodeStake is a free data retrieval call binding the contract method 0xa742329d.
//
// Solidity: function isNodeStake(address _user) view returns(bool)
func (_SwapContract *SwapContractCallerSession) IsNodeStake(_user common.Address) (bool, error) {
	return _SwapContract.Contract.IsNodeStake(&_SwapContract.CallOpts, _user)
}

// IsTxUsed is a free data retrieval call binding the contract method 0xe6ca2084.
//
// Solidity: function isTxUsed(bytes32 _txid) view returns(bool)
func (_SwapContract *SwapContractCaller) IsTxUsed(opts *bind.CallOpts, _txid [32]byte) (bool, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "isTxUsed", _txid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_SwapContract *SwapContractCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "lpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_SwapContract *SwapContractSession) LpToken() (common.Address, error) {
	return _SwapContract.Contract.LpToken(&_SwapContract.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_SwapContract *SwapContractCallerSession) LpToken() (common.Address, error) {
	return _SwapContract.Contract.LpToken(&_SwapContract.CallOpts)
}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint256)
func (_SwapContract *SwapContractCaller) NodeRewardsRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "nodeRewardsRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint256)
func (_SwapContract *SwapContractSession) NodeRewardsRatio() (*big.Int, error) {
	return _SwapContract.Contract.NodeRewardsRatio(&_SwapContract.CallOpts)
}

// NodeRewardsRatio is a free data retrieval call binding the contract method 0x0b68134d.
//
// Solidity: function nodeRewardsRatio() view returns(uint256)
func (_SwapContract *SwapContractCallerSession) NodeRewardsRatio() (*big.Int, error) {
	return _SwapContract.Contract.NodeRewardsRatio(&_SwapContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapContract *SwapContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// SbBTCPool is a free data retrieval call binding the contract method 0x0085aea1.
//
// Solidity: function sbBTCPool() view returns(address)
func (_SwapContract *SwapContractCaller) SbBTCPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "sbBTCPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SbBTCPool is a free data retrieval call binding the contract method 0x0085aea1.
//
// Solidity: function sbBTCPool() view returns(address)
func (_SwapContract *SwapContractSession) SbBTCPool() (common.Address, error) {
	return _SwapContract.Contract.SbBTCPool(&_SwapContract.CallOpts)
}

// SbBTCPool is a free data retrieval call binding the contract method 0x0085aea1.
//
// Solidity: function sbBTCPool() view returns(address)
func (_SwapContract *SwapContractCallerSession) SbBTCPool() (common.Address, error) {
	return _SwapContract.Contract.SbBTCPool(&_SwapContract.CallOpts)
}

// Sw is a free data retrieval call binding the contract method 0x00e5cee4.
//
// Solidity: function sw() view returns(address)
func (_SwapContract *SwapContractCaller) Sw(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "sw")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sw is a free data retrieval call binding the contract method 0x00e5cee4.
//
// Solidity: function sw() view returns(address)
func (_SwapContract *SwapContractSession) Sw() (common.Address, error) {
	return _SwapContract.Contract.Sw(&_SwapContract.CallOpts)
}

// Sw is a free data retrieval call binding the contract method 0x00e5cee4.
//
// Solidity: function sw() view returns(address)
func (_SwapContract *SwapContractCallerSession) Sw() (common.Address, error) {
	return _SwapContract.Contract.Sw(&_SwapContract.CallOpts)
}

// TssThreshold is a free data retrieval call binding the contract method 0x12d1441e.
//
// Solidity: function tssThreshold() view returns(uint8)
func (_SwapContract *SwapContractCaller) TssThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "tssThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TssThreshold is a free data retrieval call binding the contract method 0x12d1441e.
//
// Solidity: function tssThreshold() view returns(uint8)
func (_SwapContract *SwapContractSession) TssThreshold() (uint8, error) {
	return _SwapContract.Contract.TssThreshold(&_SwapContract.CallOpts)
}

// TssThreshold is a free data retrieval call binding the contract method 0x12d1441e.
//
// Solidity: function tssThreshold() view returns(uint8)
func (_SwapContract *SwapContractCallerSession) TssThreshold() (uint8, error) {
	return _SwapContract.Contract.TssThreshold(&_SwapContract.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_SwapContract *SwapContractCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_SwapContract *SwapContractSession) Whitelist(arg0 common.Address) (bool, error) {
	return _SwapContract.Contract.Whitelist(&_SwapContract.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_SwapContract *SwapContractCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _SwapContract.Contract.Whitelist(&_SwapContract.CallOpts, arg0)
}

// WithdrawalFeeBPS is a free data retrieval call binding the contract method 0xb6268e5d.
//
// Solidity: function withdrawalFeeBPS() view returns(uint256)
func (_SwapContract *SwapContractCaller) WithdrawalFeeBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "withdrawalFeeBPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalFeeBPS is a free data retrieval call binding the contract method 0xb6268e5d.
//
// Solidity: function withdrawalFeeBPS() view returns(uint256)
func (_SwapContract *SwapContractSession) WithdrawalFeeBPS() (*big.Int, error) {
	return _SwapContract.Contract.WithdrawalFeeBPS(&_SwapContract.CallOpts)
}

// WithdrawalFeeBPS is a free data retrieval call binding the contract method 0xb6268e5d.
//
// Solidity: function withdrawalFeeBPS() view returns(uint256)
func (_SwapContract *SwapContractCallerSession) WithdrawalFeeBPS() (*big.Int, error) {
	return _SwapContract.Contract.WithdrawalFeeBPS(&_SwapContract.CallOpts)
}

// Churn is a paid mutator transaction binding the contract method 0x6845a025.
//
// Solidity: function churn(address _newOwner, address[] _nodes, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold) returns(bool)
func (_SwapContract *SwapContractTransactor) Churn(opts *bind.TransactOpts, _newOwner common.Address, _nodes []common.Address, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "churn", _newOwner, _nodes, _isRemoved, _churnedInCount, _tssThreshold)
}

// Churn is a paid mutator transaction binding the contract method 0x6845a025.
//
// Solidity: function churn(address _newOwner, address[] _nodes, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold) returns(bool)
func (_SwapContract *SwapContractSession) Churn(_newOwner common.Address, _nodes []common.Address, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8) (*types.Transaction, error) {
	return _SwapContract.Contract.Churn(&_SwapContract.TransactOpts, _newOwner, _nodes, _isRemoved, _churnedInCount, _tssThreshold)
}

// Churn is a paid mutator transaction binding the contract method 0x6845a025.
//
// Solidity: function churn(address _newOwner, address[] _nodes, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold) returns(bool)
func (_SwapContract *SwapContractTransactorSession) Churn(_newOwner common.Address, _nodes []common.Address, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8) (*types.Transaction, error) {
	return _SwapContract.Contract.Churn(&_SwapContract.TransactOpts, _newOwner, _nodes, _isRemoved, _churnedInCount, _tssThreshold)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x2adf9f87.
//
// Solidity: function collectSwapFeesForBTC(uint256 _incomingAmount, uint256 _minerFee, uint256 _rewardsAmount, address[] _spenders, uint256[] _swapAmounts) returns(bool)
func (_SwapContract *SwapContractTransactor) CollectSwapFeesForBTC(opts *bind.TransactOpts, _incomingAmount *big.Int, _minerFee *big.Int, _rewardsAmount *big.Int, _spenders []common.Address, _swapAmounts []*big.Int) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "collectSwapFeesForBTC", _incomingAmount, _minerFee, _rewardsAmount, _spenders, _swapAmounts)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x2adf9f87.
//
// Solidity: function collectSwapFeesForBTC(uint256 _incomingAmount, uint256 _minerFee, uint256 _rewardsAmount, address[] _spenders, uint256[] _swapAmounts) returns(bool)
func (_SwapContract *SwapContractSession) CollectSwapFeesForBTC(_incomingAmount *big.Int, _minerFee *big.Int, _rewardsAmount *big.Int, _spenders []common.Address, _swapAmounts []*big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.CollectSwapFeesForBTC(&_SwapContract.TransactOpts, _incomingAmount, _minerFee, _rewardsAmount, _spenders, _swapAmounts)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x2adf9f87.
//
// Solidity: function collectSwapFeesForBTC(uint256 _incomingAmount, uint256 _minerFee, uint256 _rewardsAmount, address[] _spenders, uint256[] _swapAmounts) returns(bool)
func (_SwapContract *SwapContractTransactorSession) CollectSwapFeesForBTC(_incomingAmount *big.Int, _minerFee *big.Int, _rewardsAmount *big.Int, _spenders []common.Address, _swapAmounts []*big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.CollectSwapFeesForBTC(&_SwapContract.TransactOpts, _incomingAmount, _minerFee, _rewardsAmount, _spenders, _swapAmounts)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad289e76.
//
// Solidity: function multiTransferERC20TightlyPacked(address _destToken, bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContract *SwapContractTransactor) MultiTransferERC20TightlyPacked(opts *bind.TransactOpts, _destToken common.Address, _addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "multiTransferERC20TightlyPacked", _destToken, _addressesAndAmounts, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad289e76.
//
// Solidity: function multiTransferERC20TightlyPacked(address _destToken, bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContract *SwapContractSession) MultiTransferERC20TightlyPacked(_destToken common.Address, _addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.MultiTransferERC20TightlyPacked(&_SwapContract.TransactOpts, _destToken, _addressesAndAmounts, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad289e76.
//
// Solidity: function multiTransferERC20TightlyPacked(address _destToken, bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContract *SwapContractTransactorSession) MultiTransferERC20TightlyPacked(_destToken common.Address, _addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.MultiTransferERC20TightlyPacked(&_SwapContract.TransactOpts, _destToken, _addressesAndAmounts, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
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

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0x2586c562.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactor) RecordOutcomingFloat(opts *bind.TransactOpts, _token common.Address, _addressesAndAmountOfLPtoken [32]byte, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "recordOutcomingFloat", _token, _addressesAndAmountOfLPtoken, _minerFee, _txid)
}

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0x2586c562.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractSession) RecordOutcomingFloat(_token common.Address, _addressesAndAmountOfLPtoken [32]byte, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordOutcomingFloat(&_SwapContract.TransactOpts, _token, _addressesAndAmountOfLPtoken, _minerFee, _txid)
}

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0x2586c562.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactorSession) RecordOutcomingFloat(_token common.Address, _addressesAndAmountOfLPtoken [32]byte, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordOutcomingFloat(&_SwapContract.TransactOpts, _token, _addressesAndAmountOfLPtoken, _minerFee, _txid)
}

// RecordUTXOSweepMinerFee is a paid mutator transaction binding the contract method 0xc810a539.
//
// Solidity: function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactor) RecordUTXOSweepMinerFee(opts *bind.TransactOpts, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "recordUTXOSweepMinerFee", _minerFee, _txid)
}

// RecordUTXOSweepMinerFee is a paid mutator transaction binding the contract method 0xc810a539.
//
// Solidity: function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractSession) RecordUTXOSweepMinerFee(_minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordUTXOSweepMinerFee(&_SwapContract.TransactOpts, _minerFee, _txid)
}

// RecordUTXOSweepMinerFee is a paid mutator transaction binding the contract method 0xc810a539.
//
// Solidity: function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactorSession) RecordUTXOSweepMinerFee(_minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordUTXOSweepMinerFee(&_SwapContract.TransactOpts, _minerFee, _txid)
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

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x0d63aca7.
//
// Solidity: function singleTransferERC20(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContract *SwapContractTransactor) SingleTransferERC20(opts *bind.TransactOpts, _destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "singleTransferERC20", _destToken, _to, _amount, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x0d63aca7.
//
// Solidity: function singleTransferERC20(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContract *SwapContractSession) SingleTransferERC20(_destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.SingleTransferERC20(&_SwapContract.TransactOpts, _destToken, _to, _amount, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x0d63aca7.
//
// Solidity: function singleTransferERC20(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContract *SwapContractTransactorSession) SingleTransferERC20(_destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.SingleTransferERC20(&_SwapContract.TransactOpts, _destToken, _to, _amount, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
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

// UpdateParams is a paid mutator transaction binding the contract method 0xe880afd4.
//
// Solidity: function updateParams(address _sbBTCPool, address _buybackAddress, uint256 _withdrawalFeeBPS, uint256 _nodeRewardsRatio, uint256 _buybackRewardsRatio) returns(bool)
func (_SwapContract *SwapContractTransactor) UpdateParams(opts *bind.TransactOpts, _sbBTCPool common.Address, _buybackAddress common.Address, _withdrawalFeeBPS *big.Int, _nodeRewardsRatio *big.Int, _buybackRewardsRatio *big.Int) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "updateParams", _sbBTCPool, _buybackAddress, _withdrawalFeeBPS, _nodeRewardsRatio, _buybackRewardsRatio)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xe880afd4.
//
// Solidity: function updateParams(address _sbBTCPool, address _buybackAddress, uint256 _withdrawalFeeBPS, uint256 _nodeRewardsRatio, uint256 _buybackRewardsRatio) returns(bool)
func (_SwapContract *SwapContractSession) UpdateParams(_sbBTCPool common.Address, _buybackAddress common.Address, _withdrawalFeeBPS *big.Int, _nodeRewardsRatio *big.Int, _buybackRewardsRatio *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.UpdateParams(&_SwapContract.TransactOpts, _sbBTCPool, _buybackAddress, _withdrawalFeeBPS, _nodeRewardsRatio, _buybackRewardsRatio)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xe880afd4.
//
// Solidity: function updateParams(address _sbBTCPool, address _buybackAddress, uint256 _withdrawalFeeBPS, uint256 _nodeRewardsRatio, uint256 _buybackRewardsRatio) returns(bool)
func (_SwapContract *SwapContractTransactorSession) UpdateParams(_sbBTCPool common.Address, _buybackAddress common.Address, _withdrawalFeeBPS *big.Int, _nodeRewardsRatio *big.Int, _buybackRewardsRatio *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.UpdateParams(&_SwapContract.TransactOpts, _sbBTCPool, _buybackAddress, _withdrawalFeeBPS, _nodeRewardsRatio, _buybackRewardsRatio)
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

// SwapContractBurnLPTokensForFloatIterator is returned from FilterBurnLPTokensForFloat and is used to iterate over the raw logs and unpacked data for BurnLPTokensForFloat events raised by the SwapContract contract.
type SwapContractBurnLPTokensForFloatIterator struct {
	Event *SwapContractBurnLPTokensForFloat // Event containing the contract specifics and raw log

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
func (it *SwapContractBurnLPTokensForFloatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractBurnLPTokensForFloat)
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
		it.Event = new(SwapContractBurnLPTokensForFloat)
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
func (it *SwapContractBurnLPTokensForFloatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractBurnLPTokensForFloatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractBurnLPTokensForFloat represents a BurnLPTokensForFloat event raised by the SwapContract contract.
type SwapContractBurnLPTokensForFloat struct {
	Token          common.Address
	AmountOfLP     *big.Int
	AmountOfFloat  *big.Int
	CurrentPriceLP *big.Int
	Withdrawal     *big.Int
	Txid           [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBurnLPTokensForFloat is a free log retrieval operation binding the contract event 0xa9da34e5e4a956b744307ae7611795634e8365951073bf049d76b64b2ae7a058.
//
// Solidity: event BurnLPTokensForFloat(address token, uint256 amountOfLP, uint256 amountOfFloat, uint256 currentPriceLP, uint256 withdrawal, bytes32 txid)
func (_SwapContract *SwapContractFilterer) FilterBurnLPTokensForFloat(opts *bind.FilterOpts) (*SwapContractBurnLPTokensForFloatIterator, error) {

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "BurnLPTokensForFloat")
	if err != nil {
		return nil, err
	}
	return &SwapContractBurnLPTokensForFloatIterator{contract: _SwapContract.contract, event: "BurnLPTokensForFloat", logs: logs, sub: sub}, nil
}

// WatchBurnLPTokensForFloat is a free log subscription operation binding the contract event 0xa9da34e5e4a956b744307ae7611795634e8365951073bf049d76b64b2ae7a058.
//
// Solidity: event BurnLPTokensForFloat(address token, uint256 amountOfLP, uint256 amountOfFloat, uint256 currentPriceLP, uint256 withdrawal, bytes32 txid)
func (_SwapContract *SwapContractFilterer) WatchBurnLPTokensForFloat(opts *bind.WatchOpts, sink chan<- *SwapContractBurnLPTokensForFloat) (event.Subscription, error) {

	logs, sub, err := _SwapContract.contract.WatchLogs(opts, "BurnLPTokensForFloat")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractBurnLPTokensForFloat)
				if err := _SwapContract.contract.UnpackLog(event, "BurnLPTokensForFloat", log); err != nil {
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

// ParseBurnLPTokensForFloat is a log parse operation binding the contract event 0xa9da34e5e4a956b744307ae7611795634e8365951073bf049d76b64b2ae7a058.
//
// Solidity: event BurnLPTokensForFloat(address token, uint256 amountOfLP, uint256 amountOfFloat, uint256 currentPriceLP, uint256 withdrawal, bytes32 txid)
func (_SwapContract *SwapContractFilterer) ParseBurnLPTokensForFloat(log types.Log) (*SwapContractBurnLPTokensForFloat, error) {
	event := new(SwapContractBurnLPTokensForFloat)
	if err := _SwapContract.contract.UnpackLog(event, "BurnLPTokensForFloat", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractDistributeFeesIterator is returned from FilterDistributeFees and is used to iterate over the raw logs and unpacked data for DistributeFees events raised by the SwapContract contract.
type SwapContractDistributeFeesIterator struct {
	Event *SwapContractDistributeFees // Event containing the contract specifics and raw log

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
func (it *SwapContractDistributeFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractDistributeFees)
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
		it.Event = new(SwapContractDistributeFees)
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
func (it *SwapContractDistributeFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractDistributeFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractDistributeFees represents a DistributeFees event raised by the SwapContract contract.
type SwapContractDistributeFees struct {
	RewardLPTsForNodes *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDistributeFees is a free log retrieval operation binding the contract event 0x98247f7bfe9bc0dd2082599235b06718b8222bf674552bd10b2c055051159bd2.
//
// Solidity: event DistributeFees(uint256 rewardLPTsForNodes)
func (_SwapContract *SwapContractFilterer) FilterDistributeFees(opts *bind.FilterOpts) (*SwapContractDistributeFeesIterator, error) {

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "DistributeFees")
	if err != nil {
		return nil, err
	}
	return &SwapContractDistributeFeesIterator{contract: _SwapContract.contract, event: "DistributeFees", logs: logs, sub: sub}, nil
}

// WatchDistributeFees is a free log subscription operation binding the contract event 0x98247f7bfe9bc0dd2082599235b06718b8222bf674552bd10b2c055051159bd2.
//
// Solidity: event DistributeFees(uint256 rewardLPTsForNodes)
func (_SwapContract *SwapContractFilterer) WatchDistributeFees(opts *bind.WatchOpts, sink chan<- *SwapContractDistributeFees) (event.Subscription, error) {

	logs, sub, err := _SwapContract.contract.WatchLogs(opts, "DistributeFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractDistributeFees)
				if err := _SwapContract.contract.UnpackLog(event, "DistributeFees", log); err != nil {
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

// ParseDistributeFees is a log parse operation binding the contract event 0x98247f7bfe9bc0dd2082599235b06718b8222bf674552bd10b2c055051159bd2.
//
// Solidity: event DistributeFees(uint256 rewardLPTsForNodes)
func (_SwapContract *SwapContractFilterer) ParseDistributeFees(log types.Log) (*SwapContractDistributeFees, error) {
	event := new(SwapContractDistributeFees)
	if err := _SwapContract.contract.UnpackLog(event, "DistributeFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractIssueLPTokensForFloatIterator is returned from FilterIssueLPTokensForFloat and is used to iterate over the raw logs and unpacked data for IssueLPTokensForFloat events raised by the SwapContract contract.
type SwapContractIssueLPTokensForFloatIterator struct {
	Event *SwapContractIssueLPTokensForFloat // Event containing the contract specifics and raw log

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
func (it *SwapContractIssueLPTokensForFloatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractIssueLPTokensForFloat)
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
		it.Event = new(SwapContractIssueLPTokensForFloat)
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
func (it *SwapContractIssueLPTokensForFloatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractIssueLPTokensForFloatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractIssueLPTokensForFloat represents a IssueLPTokensForFloat event raised by the SwapContract contract.
type SwapContractIssueLPTokensForFloat struct {
	To             common.Address
	AmountOfFloat  *big.Int
	AmountOfLP     *big.Int
	CurrentPriceLP *big.Int
	DepositFees    *big.Int
	Txid           [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIssueLPTokensForFloat is a free log retrieval operation binding the contract event 0xdfaf4e23f7030adafec91a682e05bf9fb30c721f9cec3d8bcfae0459c5362db1.
//
// Solidity: event IssueLPTokensForFloat(address to, uint256 amountOfFloat, uint256 amountOfLP, uint256 currentPriceLP, uint256 depositFees, bytes32 txid)
func (_SwapContract *SwapContractFilterer) FilterIssueLPTokensForFloat(opts *bind.FilterOpts) (*SwapContractIssueLPTokensForFloatIterator, error) {

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "IssueLPTokensForFloat")
	if err != nil {
		return nil, err
	}
	return &SwapContractIssueLPTokensForFloatIterator{contract: _SwapContract.contract, event: "IssueLPTokensForFloat", logs: logs, sub: sub}, nil
}

// WatchIssueLPTokensForFloat is a free log subscription operation binding the contract event 0xdfaf4e23f7030adafec91a682e05bf9fb30c721f9cec3d8bcfae0459c5362db1.
//
// Solidity: event IssueLPTokensForFloat(address to, uint256 amountOfFloat, uint256 amountOfLP, uint256 currentPriceLP, uint256 depositFees, bytes32 txid)
func (_SwapContract *SwapContractFilterer) WatchIssueLPTokensForFloat(opts *bind.WatchOpts, sink chan<- *SwapContractIssueLPTokensForFloat) (event.Subscription, error) {

	logs, sub, err := _SwapContract.contract.WatchLogs(opts, "IssueLPTokensForFloat")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractIssueLPTokensForFloat)
				if err := _SwapContract.contract.UnpackLog(event, "IssueLPTokensForFloat", log); err != nil {
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

// ParseIssueLPTokensForFloat is a log parse operation binding the contract event 0xdfaf4e23f7030adafec91a682e05bf9fb30c721f9cec3d8bcfae0459c5362db1.
//
// Solidity: event IssueLPTokensForFloat(address to, uint256 amountOfFloat, uint256 amountOfLP, uint256 currentPriceLP, uint256 depositFees, bytes32 txid)
func (_SwapContract *SwapContractFilterer) ParseIssueLPTokensForFloat(log types.Log) (*SwapContractIssueLPTokensForFloat, error) {
	event := new(SwapContractIssueLPTokensForFloat)
	if err := _SwapContract.contract.UnpackLog(event, "IssueLPTokensForFloat", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	event.Raw = log
	return event, nil
}

// SwapContractRewardsCollectionIterator is returned from FilterRewardsCollection and is used to iterate over the raw logs and unpacked data for RewardsCollection events raised by the SwapContract contract.
type SwapContractRewardsCollectionIterator struct {
	Event *SwapContractRewardsCollection // Event containing the contract specifics and raw log

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
func (it *SwapContractRewardsCollectionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractRewardsCollection)
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
		it.Event = new(SwapContractRewardsCollection)
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
func (it *SwapContractRewardsCollectionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractRewardsCollectionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractRewardsCollection represents a RewardsCollection event raised by the SwapContract contract.
type SwapContractRewardsCollection struct {
	FeesToken       common.Address
	Rewards         *big.Int
	RewardsLPTTotal *big.Int
	CurrentPriceLP  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardsCollection is a free log retrieval operation binding the contract event 0xe3debe835f6848edc082b32a1d729781d3bfcd7e14422d80bcaa6794d3816b2f.
//
// Solidity: event RewardsCollection(address feesToken, uint256 rewards, uint256 rewardsLPTTotal, uint256 currentPriceLP)
func (_SwapContract *SwapContractFilterer) FilterRewardsCollection(opts *bind.FilterOpts) (*SwapContractRewardsCollectionIterator, error) {

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "RewardsCollection")
	if err != nil {
		return nil, err
	}
	return &SwapContractRewardsCollectionIterator{contract: _SwapContract.contract, event: "RewardsCollection", logs: logs, sub: sub}, nil
}

// WatchRewardsCollection is a free log subscription operation binding the contract event 0xe3debe835f6848edc082b32a1d729781d3bfcd7e14422d80bcaa6794d3816b2f.
//
// Solidity: event RewardsCollection(address feesToken, uint256 rewards, uint256 rewardsLPTTotal, uint256 currentPriceLP)
func (_SwapContract *SwapContractFilterer) WatchRewardsCollection(opts *bind.WatchOpts, sink chan<- *SwapContractRewardsCollection) (event.Subscription, error) {

	logs, sub, err := _SwapContract.contract.WatchLogs(opts, "RewardsCollection")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractRewardsCollection)
				if err := _SwapContract.contract.UnpackLog(event, "RewardsCollection", log); err != nil {
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

// ParseRewardsCollection is a log parse operation binding the contract event 0xe3debe835f6848edc082b32a1d729781d3bfcd7e14422d80bcaa6794d3816b2f.
//
// Solidity: event RewardsCollection(address feesToken, uint256 rewards, uint256 rewardsLPTTotal, uint256 currentPriceLP)
func (_SwapContract *SwapContractFilterer) ParseRewardsCollection(log types.Log) (*SwapContractRewardsCollection, error) {
	event := new(SwapContractRewardsCollection)
	if err := _SwapContract.contract.UnpackLog(event, "RewardsCollection", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the SwapContract contract.
type SwapContractSwapIterator struct {
	Event *SwapContractSwap // Event containing the contract specifics and raw log

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
func (it *SwapContractSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractSwap)
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
		it.Event = new(SwapContractSwap)
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
func (it *SwapContractSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractSwap represents a Swap event raised by the SwapContract contract.
type SwapContractSwap struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address from, address to, uint256 amount)
func (_SwapContract *SwapContractFilterer) FilterSwap(opts *bind.FilterOpts) (*SwapContractSwapIterator, error) {

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &SwapContractSwapIterator{contract: _SwapContract.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address from, address to, uint256 amount)
func (_SwapContract *SwapContractFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *SwapContractSwap) (event.Subscription, error) {

	logs, sub, err := _SwapContract.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractSwap)
				if err := _SwapContract.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address from, address to, uint256 amount)
func (_SwapContract *SwapContractFilterer) ParseSwap(log types.Log) (*SwapContractSwap, error) {
	event := new(SwapContractSwap)
	if err := _SwapContract.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
