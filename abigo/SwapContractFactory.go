// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigo

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

// SwapContractFactoryABI is the input ABI used to generate the binding from.
const SwapContractFactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"Deployed\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wbtc\",\"type\":\"address\"}],\"name\":\"deployNewWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SwapContractFactory is an auto generated Go binding around an Ethereum contract.
type SwapContractFactory struct {
	SwapContractFactoryCaller     // Read-only binding to the contract
	SwapContractFactoryTransactor // Write-only binding to the contract
	SwapContractFactoryFilterer   // Log filterer for contract events
}

// SwapContractFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapContractFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapContractFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapContractFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapContractFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapContractFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapContractFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapContractFactorySession struct {
	Contract     *SwapContractFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SwapContractFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapContractFactoryCallerSession struct {
	Contract *SwapContractFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SwapContractFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapContractFactoryTransactorSession struct {
	Contract     *SwapContractFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SwapContractFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapContractFactoryRaw struct {
	Contract *SwapContractFactory // Generic contract binding to access the raw methods on
}

// SwapContractFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapContractFactoryCallerRaw struct {
	Contract *SwapContractFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SwapContractFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapContractFactoryTransactorRaw struct {
	Contract *SwapContractFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapContractFactory creates a new instance of SwapContractFactory, bound to a specific deployed contract.
func NewSwapContractFactory(address common.Address, backend bind.ContractBackend) (*SwapContractFactory, error) {
	contract, err := bindSwapContractFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapContractFactory{SwapContractFactoryCaller: SwapContractFactoryCaller{contract: contract}, SwapContractFactoryTransactor: SwapContractFactoryTransactor{contract: contract}, SwapContractFactoryFilterer: SwapContractFactoryFilterer{contract: contract}}, nil
}

// NewSwapContractFactoryCaller creates a new read-only instance of SwapContractFactory, bound to a specific deployed contract.
func NewSwapContractFactoryCaller(address common.Address, caller bind.ContractCaller) (*SwapContractFactoryCaller, error) {
	contract, err := bindSwapContractFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapContractFactoryCaller{contract: contract}, nil
}

// NewSwapContractFactoryTransactor creates a new write-only instance of SwapContractFactory, bound to a specific deployed contract.
func NewSwapContractFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapContractFactoryTransactor, error) {
	contract, err := bindSwapContractFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapContractFactoryTransactor{contract: contract}, nil
}

// NewSwapContractFactoryFilterer creates a new log filterer instance of SwapContractFactory, bound to a specific deployed contract.
func NewSwapContractFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapContractFactoryFilterer, error) {
	contract, err := bindSwapContractFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapContractFactoryFilterer{contract: contract}, nil
}

// bindSwapContractFactory binds a generic wrapper to an already deployed contract.
func bindSwapContractFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapContractFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapContractFactory *SwapContractFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapContractFactory.Contract.SwapContractFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapContractFactory *SwapContractFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapContractFactory.Contract.SwapContractFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapContractFactory *SwapContractFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapContractFactory.Contract.SwapContractFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapContractFactory *SwapContractFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapContractFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapContractFactory *SwapContractFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapContractFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapContractFactory *SwapContractFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapContractFactory.Contract.contract.Transact(opts, method, params...)
}

// DeployNewWallet is a paid mutator transaction binding the contract method 0x4aced83c.
//
// Solidity: function deployNewWallet(address _owner, address _lpToken, address _wbtc) returns(address)
func (_SwapContractFactory *SwapContractFactoryTransactor) DeployNewWallet(opts *bind.TransactOpts, _owner common.Address, _lpToken common.Address, _wbtc common.Address) (*types.Transaction, error) {
	return _SwapContractFactory.contract.Transact(opts, "deployNewWallet", _owner, _lpToken, _wbtc)
}

// DeployNewWallet is a paid mutator transaction binding the contract method 0x4aced83c.
//
// Solidity: function deployNewWallet(address _owner, address _lpToken, address _wbtc) returns(address)
func (_SwapContractFactory *SwapContractFactorySession) DeployNewWallet(_owner common.Address, _lpToken common.Address, _wbtc common.Address) (*types.Transaction, error) {
	return _SwapContractFactory.Contract.DeployNewWallet(&_SwapContractFactory.TransactOpts, _owner, _lpToken, _wbtc)
}

// DeployNewWallet is a paid mutator transaction binding the contract method 0x4aced83c.
//
// Solidity: function deployNewWallet(address _owner, address _lpToken, address _wbtc) returns(address)
func (_SwapContractFactory *SwapContractFactoryTransactorSession) DeployNewWallet(_owner common.Address, _lpToken common.Address, _wbtc common.Address) (*types.Transaction, error) {
	return _SwapContractFactory.Contract.DeployNewWallet(&_SwapContractFactory.TransactOpts, _owner, _lpToken, _wbtc)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SwapContractFactory *SwapContractFactoryTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SwapContractFactory.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SwapContractFactory *SwapContractFactorySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SwapContractFactory.Contract.Fallback(&_SwapContractFactory.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SwapContractFactory *SwapContractFactoryTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SwapContractFactory.Contract.Fallback(&_SwapContractFactory.TransactOpts, calldata)
}

// SwapContractFactoryDeployedIterator is returned from FilterDeployed and is used to iterate over the raw logs and unpacked data for Deployed events raised by the SwapContractFactory contract.
type SwapContractFactoryDeployedIterator struct {
	Event *SwapContractFactoryDeployed // Event containing the contract specifics and raw log

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
func (it *SwapContractFactoryDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractFactoryDeployed)
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
		it.Event = new(SwapContractFactoryDeployed)
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
func (it *SwapContractFactoryDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractFactoryDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractFactoryDeployed represents a Deployed event raised by the SwapContractFactory contract.
type SwapContractFactoryDeployed struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeployed is a free log retrieval operation binding the contract event 0xf40fcec21964ffb566044d083b4073f29f7f7929110ea19e1b3ebe375d89055e.
//
// Solidity: event Deployed(address wallet)
func (_SwapContractFactory *SwapContractFactoryFilterer) FilterDeployed(opts *bind.FilterOpts) (*SwapContractFactoryDeployedIterator, error) {

	logs, sub, err := _SwapContractFactory.contract.FilterLogs(opts, "Deployed")
	if err != nil {
		return nil, err
	}
	return &SwapContractFactoryDeployedIterator{contract: _SwapContractFactory.contract, event: "Deployed", logs: logs, sub: sub}, nil
}

// WatchDeployed is a free log subscription operation binding the contract event 0xf40fcec21964ffb566044d083b4073f29f7f7929110ea19e1b3ebe375d89055e.
//
// Solidity: event Deployed(address wallet)
func (_SwapContractFactory *SwapContractFactoryFilterer) WatchDeployed(opts *bind.WatchOpts, sink chan<- *SwapContractFactoryDeployed) (event.Subscription, error) {

	logs, sub, err := _SwapContractFactory.contract.WatchLogs(opts, "Deployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractFactoryDeployed)
				if err := _SwapContractFactory.contract.UnpackLog(event, "Deployed", log); err != nil {
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

// ParseDeployed is a log parse operation binding the contract event 0xf40fcec21964ffb566044d083b4073f29f7f7929110ea19e1b3ebe375d89055e.
//
// Solidity: event Deployed(address wallet)
func (_SwapContractFactory *SwapContractFactoryFilterer) ParseDeployed(log types.Log) (*SwapContractFactoryDeployed, error) {
	event := new(SwapContractFactoryDeployed)
	if err := _SwapContractFactory.contract.UnpackLog(event, "Deployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
