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

// SwapRewardsMetaData contains all meta data concerning the SwapRewards contract.
var SwapRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swingby\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swap\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Paid\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dest\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_swapped\",\"type\":\"uint256\"}],\"name\":\"pullRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dest\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_receiver\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_swapped\",\"type\":\"uint256[]\"}],\"name\":\"pullRewardsMulti\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebateRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swap\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newRebateRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_thresholdRatio\",\"type\":\"uint256\"}],\"name\":\"setSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapContract\",\"outputs\":[{\"internalType\":\"contractISwapContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thresholdRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SwapRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapRewardsMetaData.ABI instead.
var SwapRewardsABI = SwapRewardsMetaData.ABI

// SwapRewards is an auto generated Go binding around an Ethereum contract.
type SwapRewards struct {
	SwapRewardsCaller     // Read-only binding to the contract
	SwapRewardsTransactor // Write-only binding to the contract
	SwapRewardsFilterer   // Log filterer for contract events
}

// SwapRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapRewardsSession struct {
	Contract     *SwapRewards      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapRewardsCallerSession struct {
	Contract *SwapRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SwapRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapRewardsTransactorSession struct {
	Contract     *SwapRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SwapRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRewardsRaw struct {
	Contract *SwapRewards // Generic contract binding to access the raw methods on
}

// SwapRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapRewardsCallerRaw struct {
	Contract *SwapRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// SwapRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapRewardsTransactorRaw struct {
	Contract *SwapRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapRewards creates a new instance of SwapRewards, bound to a specific deployed contract.
func NewSwapRewards(address common.Address, backend bind.ContractBackend) (*SwapRewards, error) {
	contract, err := bindSwapRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapRewards{SwapRewardsCaller: SwapRewardsCaller{contract: contract}, SwapRewardsTransactor: SwapRewardsTransactor{contract: contract}, SwapRewardsFilterer: SwapRewardsFilterer{contract: contract}}, nil
}

// NewSwapRewardsCaller creates a new read-only instance of SwapRewards, bound to a specific deployed contract.
func NewSwapRewardsCaller(address common.Address, caller bind.ContractCaller) (*SwapRewardsCaller, error) {
	contract, err := bindSwapRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapRewardsCaller{contract: contract}, nil
}

// NewSwapRewardsTransactor creates a new write-only instance of SwapRewards, bound to a specific deployed contract.
func NewSwapRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapRewardsTransactor, error) {
	contract, err := bindSwapRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapRewardsTransactor{contract: contract}, nil
}

// NewSwapRewardsFilterer creates a new log filterer instance of SwapRewards, bound to a specific deployed contract.
func NewSwapRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapRewardsFilterer, error) {
	contract, err := bindSwapRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapRewardsFilterer{contract: contract}, nil
}

// bindSwapRewards binds a generic wrapper to an already deployed contract.
func bindSwapRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapRewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapRewards *SwapRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapRewards.Contract.SwapRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapRewards *SwapRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRewards.Contract.SwapRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapRewards *SwapRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapRewards.Contract.SwapRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapRewards *SwapRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapRewards *SwapRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapRewards *SwapRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapRewards.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapRewards *SwapRewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapRewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapRewards *SwapRewardsSession) Owner() (common.Address, error) {
	return _SwapRewards.Contract.Owner(&_SwapRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapRewards *SwapRewardsCallerSession) Owner() (common.Address, error) {
	return _SwapRewards.Contract.Owner(&_SwapRewards.CallOpts)
}

// RebateRate is a free data retrieval call binding the contract method 0xb681a1e7.
//
// Solidity: function rebateRate() view returns(uint256)
func (_SwapRewards *SwapRewardsCaller) RebateRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapRewards.contract.Call(opts, &out, "rebateRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebateRate is a free data retrieval call binding the contract method 0xb681a1e7.
//
// Solidity: function rebateRate() view returns(uint256)
func (_SwapRewards *SwapRewardsSession) RebateRate() (*big.Int, error) {
	return _SwapRewards.Contract.RebateRate(&_SwapRewards.CallOpts)
}

// RebateRate is a free data retrieval call binding the contract method 0xb681a1e7.
//
// Solidity: function rebateRate() view returns(uint256)
func (_SwapRewards *SwapRewardsCallerSession) RebateRate() (*big.Int, error) {
	return _SwapRewards.Contract.RebateRate(&_SwapRewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SwapRewards *SwapRewardsCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapRewards.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SwapRewards *SwapRewardsSession) RewardToken() (common.Address, error) {
	return _SwapRewards.Contract.RewardToken(&_SwapRewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_SwapRewards *SwapRewardsCallerSession) RewardToken() (common.Address, error) {
	return _SwapRewards.Contract.RewardToken(&_SwapRewards.CallOpts)
}

// SwapContract is a free data retrieval call binding the contract method 0x8ea83031.
//
// Solidity: function swapContract() view returns(address)
func (_SwapRewards *SwapRewardsCaller) SwapContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapRewards.contract.Call(opts, &out, "swapContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapContract is a free data retrieval call binding the contract method 0x8ea83031.
//
// Solidity: function swapContract() view returns(address)
func (_SwapRewards *SwapRewardsSession) SwapContract() (common.Address, error) {
	return _SwapRewards.Contract.SwapContract(&_SwapRewards.CallOpts)
}

// SwapContract is a free data retrieval call binding the contract method 0x8ea83031.
//
// Solidity: function swapContract() view returns(address)
func (_SwapRewards *SwapRewardsCallerSession) SwapContract() (common.Address, error) {
	return _SwapRewards.Contract.SwapContract(&_SwapRewards.CallOpts)
}

// ThresholdRatio is a free data retrieval call binding the contract method 0xc0324da1.
//
// Solidity: function thresholdRatio() view returns(uint256)
func (_SwapRewards *SwapRewardsCaller) ThresholdRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapRewards.contract.Call(opts, &out, "thresholdRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ThresholdRatio is a free data retrieval call binding the contract method 0xc0324da1.
//
// Solidity: function thresholdRatio() view returns(uint256)
func (_SwapRewards *SwapRewardsSession) ThresholdRatio() (*big.Int, error) {
	return _SwapRewards.Contract.ThresholdRatio(&_SwapRewards.CallOpts)
}

// ThresholdRatio is a free data retrieval call binding the contract method 0xc0324da1.
//
// Solidity: function thresholdRatio() view returns(uint256)
func (_SwapRewards *SwapRewardsCallerSession) ThresholdRatio() (*big.Int, error) {
	return _SwapRewards.Contract.ThresholdRatio(&_SwapRewards.CallOpts)
}

// PullRewards is a paid mutator transaction binding the contract method 0xff320aae.
//
// Solidity: function pullRewards(address _dest, address _receiver, uint256 _swapped) returns()
func (_SwapRewards *SwapRewardsTransactor) PullRewards(opts *bind.TransactOpts, _dest common.Address, _receiver common.Address, _swapped *big.Int) (*types.Transaction, error) {
	return _SwapRewards.contract.Transact(opts, "pullRewards", _dest, _receiver, _swapped)
}

// PullRewards is a paid mutator transaction binding the contract method 0xff320aae.
//
// Solidity: function pullRewards(address _dest, address _receiver, uint256 _swapped) returns()
func (_SwapRewards *SwapRewardsSession) PullRewards(_dest common.Address, _receiver common.Address, _swapped *big.Int) (*types.Transaction, error) {
	return _SwapRewards.Contract.PullRewards(&_SwapRewards.TransactOpts, _dest, _receiver, _swapped)
}

// PullRewards is a paid mutator transaction binding the contract method 0xff320aae.
//
// Solidity: function pullRewards(address _dest, address _receiver, uint256 _swapped) returns()
func (_SwapRewards *SwapRewardsTransactorSession) PullRewards(_dest common.Address, _receiver common.Address, _swapped *big.Int) (*types.Transaction, error) {
	return _SwapRewards.Contract.PullRewards(&_SwapRewards.TransactOpts, _dest, _receiver, _swapped)
}

// PullRewardsMulti is a paid mutator transaction binding the contract method 0x49e031a5.
//
// Solidity: function pullRewardsMulti(address _dest, address[] _receiver, uint256[] _swapped) returns()
func (_SwapRewards *SwapRewardsTransactor) PullRewardsMulti(opts *bind.TransactOpts, _dest common.Address, _receiver []common.Address, _swapped []*big.Int) (*types.Transaction, error) {
	return _SwapRewards.contract.Transact(opts, "pullRewardsMulti", _dest, _receiver, _swapped)
}

// PullRewardsMulti is a paid mutator transaction binding the contract method 0x49e031a5.
//
// Solidity: function pullRewardsMulti(address _dest, address[] _receiver, uint256[] _swapped) returns()
func (_SwapRewards *SwapRewardsSession) PullRewardsMulti(_dest common.Address, _receiver []common.Address, _swapped []*big.Int) (*types.Transaction, error) {
	return _SwapRewards.Contract.PullRewardsMulti(&_SwapRewards.TransactOpts, _dest, _receiver, _swapped)
}

// PullRewardsMulti is a paid mutator transaction binding the contract method 0x49e031a5.
//
// Solidity: function pullRewardsMulti(address _dest, address[] _receiver, uint256[] _swapped) returns()
func (_SwapRewards *SwapRewardsTransactorSession) PullRewardsMulti(_dest common.Address, _receiver []common.Address, _swapped []*big.Int) (*types.Transaction, error) {
	return _SwapRewards.Contract.PullRewardsMulti(&_SwapRewards.TransactOpts, _dest, _receiver, _swapped)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapRewards *SwapRewardsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapRewards.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapRewards *SwapRewardsSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapRewards.Contract.RenounceOwnership(&_SwapRewards.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapRewards *SwapRewardsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapRewards.Contract.RenounceOwnership(&_SwapRewards.TransactOpts)
}

// SetSwap is a paid mutator transaction binding the contract method 0xa9d6e083.
//
// Solidity: function setSwap(address _swap, uint256 _newRebateRate, uint256 _thresholdRatio) returns()
func (_SwapRewards *SwapRewardsTransactor) SetSwap(opts *bind.TransactOpts, _swap common.Address, _newRebateRate *big.Int, _thresholdRatio *big.Int) (*types.Transaction, error) {
	return _SwapRewards.contract.Transact(opts, "setSwap", _swap, _newRebateRate, _thresholdRatio)
}

// SetSwap is a paid mutator transaction binding the contract method 0xa9d6e083.
//
// Solidity: function setSwap(address _swap, uint256 _newRebateRate, uint256 _thresholdRatio) returns()
func (_SwapRewards *SwapRewardsSession) SetSwap(_swap common.Address, _newRebateRate *big.Int, _thresholdRatio *big.Int) (*types.Transaction, error) {
	return _SwapRewards.Contract.SetSwap(&_SwapRewards.TransactOpts, _swap, _newRebateRate, _thresholdRatio)
}

// SetSwap is a paid mutator transaction binding the contract method 0xa9d6e083.
//
// Solidity: function setSwap(address _swap, uint256 _newRebateRate, uint256 _thresholdRatio) returns()
func (_SwapRewards *SwapRewardsTransactorSession) SetSwap(_swap common.Address, _newRebateRate *big.Int, _thresholdRatio *big.Int) (*types.Transaction, error) {
	return _SwapRewards.Contract.SetSwap(&_SwapRewards.TransactOpts, _swap, _newRebateRate, _thresholdRatio)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapRewards *SwapRewardsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SwapRewards.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapRewards *SwapRewardsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapRewards.Contract.TransferOwnership(&_SwapRewards.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapRewards *SwapRewardsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapRewards.Contract.TransferOwnership(&_SwapRewards.TransactOpts, newOwner)
}

// SwapRewardsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SwapRewards contract.
type SwapRewardsOwnershipTransferredIterator struct {
	Event *SwapRewardsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SwapRewardsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapRewardsOwnershipTransferred)
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
		it.Event = new(SwapRewardsOwnershipTransferred)
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
func (it *SwapRewardsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapRewardsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapRewardsOwnershipTransferred represents a OwnershipTransferred event raised by the SwapRewards contract.
type SwapRewardsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapRewards *SwapRewardsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapRewardsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapRewards.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapRewardsOwnershipTransferredIterator{contract: _SwapRewards.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapRewards *SwapRewardsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapRewardsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapRewards.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapRewardsOwnershipTransferred)
				if err := _SwapRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SwapRewards *SwapRewardsFilterer) ParseOwnershipTransferred(log types.Log) (*SwapRewardsOwnershipTransferred, error) {
	event := new(SwapRewardsOwnershipTransferred)
	if err := _SwapRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapRewardsPaidIterator is returned from FilterPaid and is used to iterate over the raw logs and unpacked data for Paid events raised by the SwapRewards contract.
type SwapRewardsPaidIterator struct {
	Event *SwapRewardsPaid // Event containing the contract specifics and raw log

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
func (it *SwapRewardsPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapRewardsPaid)
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
		it.Event = new(SwapRewardsPaid)
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
func (it *SwapRewardsPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapRewardsPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapRewardsPaid represents a Paid event raised by the SwapRewards contract.
type SwapRewardsPaid struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaid is a free log retrieval operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address to, uint256 amount)
func (_SwapRewards *SwapRewardsFilterer) FilterPaid(opts *bind.FilterOpts) (*SwapRewardsPaidIterator, error) {

	logs, sub, err := _SwapRewards.contract.FilterLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return &SwapRewardsPaidIterator{contract: _SwapRewards.contract, event: "Paid", logs: logs, sub: sub}, nil
}

// WatchPaid is a free log subscription operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address to, uint256 amount)
func (_SwapRewards *SwapRewardsFilterer) WatchPaid(opts *bind.WatchOpts, sink chan<- *SwapRewardsPaid) (event.Subscription, error) {

	logs, sub, err := _SwapRewards.contract.WatchLogs(opts, "Paid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapRewardsPaid)
				if err := _SwapRewards.contract.UnpackLog(event, "Paid", log); err != nil {
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

// ParsePaid is a log parse operation binding the contract event 0x737c69225d647e5994eab1a6c301bf6d9232beb2759ae1e27a8966b4732bc489.
//
// Solidity: event Paid(address to, uint256 amount)
func (_SwapRewards *SwapRewardsFilterer) ParsePaid(log types.Log) (*SwapRewardsPaid, error) {
	event := new(SwapRewardsPaid)
	if err := _SwapRewards.contract.UnpackLog(event, "Paid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
