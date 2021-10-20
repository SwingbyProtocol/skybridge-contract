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

// IAugustusSwapperBuyData is an auto generated low-level Go binding around an user-defined struct.
type IAugustusSwapperBuyData struct {
	FromToken     common.Address
	ToToken       common.Address
	FromAmount    *big.Int
	ToAmount      *big.Int
	Beneficiary   common.Address
	Referrer      string
	UseReduxToken bool
	Route         []IAugustusSwapperBuyRoute
}

// IAugustusSwapperBuyRoute is an auto generated low-level Go binding around an user-defined struct.
type IAugustusSwapperBuyRoute struct {
	Exchange       common.Address
	TargetExchange common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	Payload        []byte
	NetworkFee     *big.Int
}

// IAugustusSwapperMegaSwapPath is an auto generated low-level Go binding around an user-defined struct.
type IAugustusSwapperMegaSwapPath struct {
	FromAmountPercent *big.Int
	Path              []IAugustusSwapperPath
}

// IAugustusSwapperMegaSwapSellData is an auto generated low-level Go binding around an user-defined struct.
type IAugustusSwapperMegaSwapSellData struct {
	FromToken      common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	Beneficiary    common.Address
	Referrer       string
	UseReduxToken  bool
	Path           []IAugustusSwapperMegaSwapPath
}

// IAugustusSwapperPath is an auto generated low-level Go binding around an user-defined struct.
type IAugustusSwapperPath struct {
	To              common.Address
	TotalNetworkFee *big.Int
	Routes          []IAugustusSwapperRoute
}

// IAugustusSwapperRoute is an auto generated low-level Go binding around an user-defined struct.
type IAugustusSwapperRoute struct {
	Exchange       common.Address
	TargetExchange common.Address
	Percent        *big.Int
	Payload        []byte
	NetworkFee     *big.Int
}

// IAugustusSwapperSellData is an auto generated low-level Go binding around an user-defined struct.
type IAugustusSwapperSellData struct {
	FromToken      common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	Beneficiary    common.Address
	Referrer       string
	UseReduxToken  bool
	Path           []IAugustusSwapperPath
}

// IAugustusSwapperABI is the input ABI used to generate the binding from.
const IAugustusSwapperABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"referrer\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useReduxToken\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetExchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"internalType\":\"structIAugustusSwapper.BuyRoute[]\",\"name\":\"route\",\"type\":\"tuple[]\"}],\"internalType\":\"structIAugustusSwapper.BuyData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"referrer\",\"type\":\"uint8\"}],\"name\":\"buyOnUniswap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"referrer\",\"type\":\"uint8\"}],\"name\":\"buyOnUniswapFork\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPartnerRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenTransferProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUniswapProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"referrer\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useReduxToken\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromAmountPercent\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalNetworkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetExchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"internalType\":\"structIAugustusSwapper.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"internalType\":\"structIAugustusSwapper.Path[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structIAugustusSwapper.MegaSwapPath[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structIAugustusSwapper.MegaSwapSellData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"megaSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"referrer\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useReduxToken\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalNetworkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetExchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"internalType\":\"structIAugustusSwapper.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"internalType\":\"structIAugustusSwapper.Path[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structIAugustusSwapper.SellData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"multiSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"callees\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"exchangeData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"startIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"referrer\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useReduxToken\",\"type\":\"bool\"}],\"name\":\"simplBuy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"callees\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"exchangeData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"startIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"referrer\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useReduxToken\",\"type\":\"bool\"}],\"name\":\"simpleSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"referrer\",\"type\":\"uint8\"}],\"name\":\"swapOnUniswap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"referrer\",\"type\":\"uint8\"}],\"name\":\"swapOnUniswapFork\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IAugustusSwapper is an auto generated Go binding around an Ethereum contract.
type IAugustusSwapper struct {
	IAugustusSwapperCaller     // Read-only binding to the contract
	IAugustusSwapperTransactor // Write-only binding to the contract
	IAugustusSwapperFilterer   // Log filterer for contract events
}

// IAugustusSwapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAugustusSwapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAugustusSwapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAugustusSwapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAugustusSwapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAugustusSwapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAugustusSwapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAugustusSwapperSession struct {
	Contract     *IAugustusSwapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAugustusSwapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAugustusSwapperCallerSession struct {
	Contract *IAugustusSwapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IAugustusSwapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAugustusSwapperTransactorSession struct {
	Contract     *IAugustusSwapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IAugustusSwapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAugustusSwapperRaw struct {
	Contract *IAugustusSwapper // Generic contract binding to access the raw methods on
}

// IAugustusSwapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAugustusSwapperCallerRaw struct {
	Contract *IAugustusSwapperCaller // Generic read-only contract binding to access the raw methods on
}

// IAugustusSwapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAugustusSwapperTransactorRaw struct {
	Contract *IAugustusSwapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAugustusSwapper creates a new instance of IAugustusSwapper, bound to a specific deployed contract.
func NewIAugustusSwapper(address common.Address, backend bind.ContractBackend) (*IAugustusSwapper, error) {
	contract, err := bindIAugustusSwapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAugustusSwapper{IAugustusSwapperCaller: IAugustusSwapperCaller{contract: contract}, IAugustusSwapperTransactor: IAugustusSwapperTransactor{contract: contract}, IAugustusSwapperFilterer: IAugustusSwapperFilterer{contract: contract}}, nil
}

// NewIAugustusSwapperCaller creates a new read-only instance of IAugustusSwapper, bound to a specific deployed contract.
func NewIAugustusSwapperCaller(address common.Address, caller bind.ContractCaller) (*IAugustusSwapperCaller, error) {
	contract, err := bindIAugustusSwapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAugustusSwapperCaller{contract: contract}, nil
}

// NewIAugustusSwapperTransactor creates a new write-only instance of IAugustusSwapper, bound to a specific deployed contract.
func NewIAugustusSwapperTransactor(address common.Address, transactor bind.ContractTransactor) (*IAugustusSwapperTransactor, error) {
	contract, err := bindIAugustusSwapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAugustusSwapperTransactor{contract: contract}, nil
}

// NewIAugustusSwapperFilterer creates a new log filterer instance of IAugustusSwapper, bound to a specific deployed contract.
func NewIAugustusSwapperFilterer(address common.Address, filterer bind.ContractFilterer) (*IAugustusSwapperFilterer, error) {
	contract, err := bindIAugustusSwapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAugustusSwapperFilterer{contract: contract}, nil
}

// bindIAugustusSwapper binds a generic wrapper to an already deployed contract.
func bindIAugustusSwapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAugustusSwapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAugustusSwapper *IAugustusSwapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IAugustusSwapper.Contract.IAugustusSwapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAugustusSwapper *IAugustusSwapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.IAugustusSwapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAugustusSwapper *IAugustusSwapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.IAugustusSwapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAugustusSwapper *IAugustusSwapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IAugustusSwapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAugustusSwapper *IAugustusSwapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAugustusSwapper *IAugustusSwapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.contract.Transact(opts, method, params...)
}

// GetFeeWallet is a free data retrieval call binding the contract method 0x5459060d.
//
// Solidity: function getFeeWallet() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCaller) GetFeeWallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IAugustusSwapper.contract.Call(opts, out, "getFeeWallet")
	return *ret0, err
}

// GetFeeWallet is a free data retrieval call binding the contract method 0x5459060d.
//
// Solidity: function getFeeWallet() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperSession) GetFeeWallet() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetFeeWallet(&_IAugustusSwapper.CallOpts)
}

// GetFeeWallet is a free data retrieval call binding the contract method 0x5459060d.
//
// Solidity: function getFeeWallet() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCallerSession) GetFeeWallet() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetFeeWallet(&_IAugustusSwapper.CallOpts)
}

// GetPartnerRegistry is a free data retrieval call binding the contract method 0xb69bd7aa.
//
// Solidity: function getPartnerRegistry() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCaller) GetPartnerRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IAugustusSwapper.contract.Call(opts, out, "getPartnerRegistry")
	return *ret0, err
}

// GetPartnerRegistry is a free data retrieval call binding the contract method 0xb69bd7aa.
//
// Solidity: function getPartnerRegistry() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperSession) GetPartnerRegistry() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetPartnerRegistry(&_IAugustusSwapper.CallOpts)
}

// GetPartnerRegistry is a free data retrieval call binding the contract method 0xb69bd7aa.
//
// Solidity: function getPartnerRegistry() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCallerSession) GetPartnerRegistry() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetPartnerRegistry(&_IAugustusSwapper.CallOpts)
}

// GetTokenTransferProxy is a free data retrieval call binding the contract method 0xd2c4b598.
//
// Solidity: function getTokenTransferProxy() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCaller) GetTokenTransferProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IAugustusSwapper.contract.Call(opts, out, "getTokenTransferProxy")
	return *ret0, err
}

// GetTokenTransferProxy is a free data retrieval call binding the contract method 0xd2c4b598.
//
// Solidity: function getTokenTransferProxy() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperSession) GetTokenTransferProxy() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetTokenTransferProxy(&_IAugustusSwapper.CallOpts)
}

// GetTokenTransferProxy is a free data retrieval call binding the contract method 0xd2c4b598.
//
// Solidity: function getTokenTransferProxy() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCallerSession) GetTokenTransferProxy() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetTokenTransferProxy(&_IAugustusSwapper.CallOpts)
}

// GetUniswapProxy is a free data retrieval call binding the contract method 0x08defdee.
//
// Solidity: function getUniswapProxy() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCaller) GetUniswapProxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IAugustusSwapper.contract.Call(opts, out, "getUniswapProxy")
	return *ret0, err
}

// GetUniswapProxy is a free data retrieval call binding the contract method 0x08defdee.
//
// Solidity: function getUniswapProxy() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperSession) GetUniswapProxy() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetUniswapProxy(&_IAugustusSwapper.CallOpts)
}

// GetUniswapProxy is a free data retrieval call binding the contract method 0x08defdee.
//
// Solidity: function getUniswapProxy() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCallerSession) GetUniswapProxy() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetUniswapProxy(&_IAugustusSwapper.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_IAugustusSwapper *IAugustusSwapperCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IAugustusSwapper.contract.Call(opts, out, "getVersion")
	return *ret0, err
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_IAugustusSwapper *IAugustusSwapperSession) GetVersion() (string, error) {
	return _IAugustusSwapper.Contract.GetVersion(&_IAugustusSwapper.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_IAugustusSwapper *IAugustusSwapperCallerSession) GetVersion() (string, error) {
	return _IAugustusSwapper.Contract.GetVersion(&_IAugustusSwapper.CallOpts)
}

// GetWhitelistAddress is a free data retrieval call binding the contract method 0x915eb973.
//
// Solidity: function getWhitelistAddress() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCaller) GetWhitelistAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IAugustusSwapper.contract.Call(opts, out, "getWhitelistAddress")
	return *ret0, err
}

// GetWhitelistAddress is a free data retrieval call binding the contract method 0x915eb973.
//
// Solidity: function getWhitelistAddress() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperSession) GetWhitelistAddress() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetWhitelistAddress(&_IAugustusSwapper.CallOpts)
}

// GetWhitelistAddress is a free data retrieval call binding the contract method 0x915eb973.
//
// Solidity: function getWhitelistAddress() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCallerSession) GetWhitelistAddress() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetWhitelistAddress(&_IAugustusSwapper.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xf95a49eb.
//
// Solidity: function buy((address,address,uint256,uint256,address,string,bool,(address,address,uint256,uint256,bytes,uint256)[]) data) payable returns(uint256)
func (_IAugustusSwapper *IAugustusSwapperTransactor) Buy(opts *bind.TransactOpts, data IAugustusSwapperBuyData) (*types.Transaction, error) {
	return _IAugustusSwapper.contract.Transact(opts, "buy", data)
}

// Buy is a paid mutator transaction binding the contract method 0xf95a49eb.
//
// Solidity: function buy((address,address,uint256,uint256,address,string,bool,(address,address,uint256,uint256,bytes,uint256)[]) data) payable returns(uint256)
func (_IAugustusSwapper *IAugustusSwapperSession) Buy(data IAugustusSwapperBuyData) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.Buy(&_IAugustusSwapper.TransactOpts, data)
}

// Buy is a paid mutator transaction binding the contract method 0xf95a49eb.
//
// Solidity: function buy((address,address,uint256,uint256,address,string,bool,(address,address,uint256,uint256,bytes,uint256)[]) data) payable returns(uint256)
func (_IAugustusSwapper *IAugustusSwapperTransactorSession) Buy(data IAugustusSwapperBuyData) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.Buy(&_IAugustusSwapper.TransactOpts, data)
}

// BuyOnUniswap is a paid mutator transaction binding the contract method 0xf9355f72.
//
// Solidity: function buyOnUniswap(uint256 amountInMax, uint256 amountOut, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperTransactor) BuyOnUniswap(opts *bind.TransactOpts, amountInMax *big.Int, amountOut *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.contract.Transact(opts, "buyOnUniswap", amountInMax, amountOut, path, referrer)
}

// BuyOnUniswap is a paid mutator transaction binding the contract method 0xf9355f72.
//
// Solidity: function buyOnUniswap(uint256 amountInMax, uint256 amountOut, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperSession) BuyOnUniswap(amountInMax *big.Int, amountOut *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.BuyOnUniswap(&_IAugustusSwapper.TransactOpts, amountInMax, amountOut, path, referrer)
}

// BuyOnUniswap is a paid mutator transaction binding the contract method 0xf9355f72.
//
// Solidity: function buyOnUniswap(uint256 amountInMax, uint256 amountOut, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperTransactorSession) BuyOnUniswap(amountInMax *big.Int, amountOut *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.BuyOnUniswap(&_IAugustusSwapper.TransactOpts, amountInMax, amountOut, path, referrer)
}

// BuyOnUniswapFork is a paid mutator transaction binding the contract method 0x33635226.
//
// Solidity: function buyOnUniswapFork(address factory, bytes32 initCode, uint256 amountInMax, uint256 amountOut, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperTransactor) BuyOnUniswapFork(opts *bind.TransactOpts, factory common.Address, initCode [32]byte, amountInMax *big.Int, amountOut *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.contract.Transact(opts, "buyOnUniswapFork", factory, initCode, amountInMax, amountOut, path, referrer)
}

// BuyOnUniswapFork is a paid mutator transaction binding the contract method 0x33635226.
//
// Solidity: function buyOnUniswapFork(address factory, bytes32 initCode, uint256 amountInMax, uint256 amountOut, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperSession) BuyOnUniswapFork(factory common.Address, initCode [32]byte, amountInMax *big.Int, amountOut *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.BuyOnUniswapFork(&_IAugustusSwapper.TransactOpts, factory, initCode, amountInMax, amountOut, path, referrer)
}

// BuyOnUniswapFork is a paid mutator transaction binding the contract method 0x33635226.
//
// Solidity: function buyOnUniswapFork(address factory, bytes32 initCode, uint256 amountInMax, uint256 amountOut, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperTransactorSession) BuyOnUniswapFork(factory common.Address, initCode [32]byte, amountInMax *big.Int, amountOut *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.BuyOnUniswapFork(&_IAugustusSwapper.TransactOpts, factory, initCode, amountInMax, amountOut, path, referrer)
}

// MegaSwap is a paid mutator transaction binding the contract method 0xec1d21dd.
//
// Solidity: function megaSwap((address,uint256,uint256,uint256,address,string,bool,(uint256,(address,uint256,(address,address,uint256,bytes,uint256)[])[])[]) data) payable returns(uint256)
func (_IAugustusSwapper *IAugustusSwapperTransactor) MegaSwap(opts *bind.TransactOpts, data IAugustusSwapperMegaSwapSellData) (*types.Transaction, error) {
	return _IAugustusSwapper.contract.Transact(opts, "megaSwap", data)
}

// MegaSwap is a paid mutator transaction binding the contract method 0xec1d21dd.
//
// Solidity: function megaSwap((address,uint256,uint256,uint256,address,string,bool,(uint256,(address,uint256,(address,address,uint256,bytes,uint256)[])[])[]) data) payable returns(uint256)
func (_IAugustusSwapper *IAugustusSwapperSession) MegaSwap(data IAugustusSwapperMegaSwapSellData) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.MegaSwap(&_IAugustusSwapper.TransactOpts, data)
}

// MegaSwap is a paid mutator transaction binding the contract method 0xec1d21dd.
//
// Solidity: function megaSwap((address,uint256,uint256,uint256,address,string,bool,(uint256,(address,uint256,(address,address,uint256,bytes,uint256)[])[])[]) data) payable returns(uint256)
func (_IAugustusSwapper *IAugustusSwapperTransactorSession) MegaSwap(data IAugustusSwapperMegaSwapSellData) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.MegaSwap(&_IAugustusSwapper.TransactOpts, data)
}

// MultiSwap is a paid mutator transaction binding the contract method 0x8f00eccb.
//
// Solidity: function multiSwap((address,uint256,uint256,uint256,address,string,bool,(address,uint256,(address,address,uint256,bytes,uint256)[])[]) data) payable returns(uint256)
func (_IAugustusSwapper *IAugustusSwapperTransactor) MultiSwap(opts *bind.TransactOpts, data IAugustusSwapperSellData) (*types.Transaction, error) {
	return _IAugustusSwapper.contract.Transact(opts, "multiSwap", data)
}

// MultiSwap is a paid mutator transaction binding the contract method 0x8f00eccb.
//
// Solidity: function multiSwap((address,uint256,uint256,uint256,address,string,bool,(address,uint256,(address,address,uint256,bytes,uint256)[])[]) data) payable returns(uint256)
func (_IAugustusSwapper *IAugustusSwapperSession) MultiSwap(data IAugustusSwapperSellData) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.MultiSwap(&_IAugustusSwapper.TransactOpts, data)
}

// MultiSwap is a paid mutator transaction binding the contract method 0x8f00eccb.
//
// Solidity: function multiSwap((address,uint256,uint256,uint256,address,string,bool,(address,uint256,(address,address,uint256,bytes,uint256)[])[]) data) payable returns(uint256)
func (_IAugustusSwapper *IAugustusSwapperTransactorSession) MultiSwap(data IAugustusSwapperSellData) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.MultiSwap(&_IAugustusSwapper.TransactOpts, data)
}

// SimplBuy is a paid mutator transaction binding the contract method 0xa27e8b6b.
//
// Solidity: function simplBuy(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address[] callees, bytes exchangeData, uint256[] startIndexes, uint256[] values, address beneficiary, string referrer, bool useReduxToken) payable returns()
func (_IAugustusSwapper *IAugustusSwapperTransactor) SimplBuy(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int, toAmount *big.Int, callees []common.Address, exchangeData []byte, startIndexes []*big.Int, values []*big.Int, beneficiary common.Address, referrer string, useReduxToken bool) (*types.Transaction, error) {
	return _IAugustusSwapper.contract.Transact(opts, "simplBuy", fromToken, toToken, fromAmount, toAmount, callees, exchangeData, startIndexes, values, beneficiary, referrer, useReduxToken)
}

// SimplBuy is a paid mutator transaction binding the contract method 0xa27e8b6b.
//
// Solidity: function simplBuy(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address[] callees, bytes exchangeData, uint256[] startIndexes, uint256[] values, address beneficiary, string referrer, bool useReduxToken) payable returns()
func (_IAugustusSwapper *IAugustusSwapperSession) SimplBuy(fromToken common.Address, toToken common.Address, fromAmount *big.Int, toAmount *big.Int, callees []common.Address, exchangeData []byte, startIndexes []*big.Int, values []*big.Int, beneficiary common.Address, referrer string, useReduxToken bool) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.SimplBuy(&_IAugustusSwapper.TransactOpts, fromToken, toToken, fromAmount, toAmount, callees, exchangeData, startIndexes, values, beneficiary, referrer, useReduxToken)
}

// SimplBuy is a paid mutator transaction binding the contract method 0xa27e8b6b.
//
// Solidity: function simplBuy(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, address[] callees, bytes exchangeData, uint256[] startIndexes, uint256[] values, address beneficiary, string referrer, bool useReduxToken) payable returns()
func (_IAugustusSwapper *IAugustusSwapperTransactorSession) SimplBuy(fromToken common.Address, toToken common.Address, fromAmount *big.Int, toAmount *big.Int, callees []common.Address, exchangeData []byte, startIndexes []*big.Int, values []*big.Int, beneficiary common.Address, referrer string, useReduxToken bool) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.SimplBuy(&_IAugustusSwapper.TransactOpts, fromToken, toToken, fromAmount, toAmount, callees, exchangeData, startIndexes, values, beneficiary, referrer, useReduxToken)
}

// SimpleSwap is a paid mutator transaction binding the contract method 0xcfc0afeb.
//
// Solidity: function simpleSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, uint256 expectedAmount, address[] callees, bytes exchangeData, uint256[] startIndexes, uint256[] values, address beneficiary, string referrer, bool useReduxToken) payable returns(uint256 receivedAmount)
func (_IAugustusSwapper *IAugustusSwapperTransactor) SimpleSwap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int, toAmount *big.Int, expectedAmount *big.Int, callees []common.Address, exchangeData []byte, startIndexes []*big.Int, values []*big.Int, beneficiary common.Address, referrer string, useReduxToken bool) (*types.Transaction, error) {
	return _IAugustusSwapper.contract.Transact(opts, "simpleSwap", fromToken, toToken, fromAmount, toAmount, expectedAmount, callees, exchangeData, startIndexes, values, beneficiary, referrer, useReduxToken)
}

// SimpleSwap is a paid mutator transaction binding the contract method 0xcfc0afeb.
//
// Solidity: function simpleSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, uint256 expectedAmount, address[] callees, bytes exchangeData, uint256[] startIndexes, uint256[] values, address beneficiary, string referrer, bool useReduxToken) payable returns(uint256 receivedAmount)
func (_IAugustusSwapper *IAugustusSwapperSession) SimpleSwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, toAmount *big.Int, expectedAmount *big.Int, callees []common.Address, exchangeData []byte, startIndexes []*big.Int, values []*big.Int, beneficiary common.Address, referrer string, useReduxToken bool) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.SimpleSwap(&_IAugustusSwapper.TransactOpts, fromToken, toToken, fromAmount, toAmount, expectedAmount, callees, exchangeData, startIndexes, values, beneficiary, referrer, useReduxToken)
}

// SimpleSwap is a paid mutator transaction binding the contract method 0xcfc0afeb.
//
// Solidity: function simpleSwap(address fromToken, address toToken, uint256 fromAmount, uint256 toAmount, uint256 expectedAmount, address[] callees, bytes exchangeData, uint256[] startIndexes, uint256[] values, address beneficiary, string referrer, bool useReduxToken) payable returns(uint256 receivedAmount)
func (_IAugustusSwapper *IAugustusSwapperTransactorSession) SimpleSwap(fromToken common.Address, toToken common.Address, fromAmount *big.Int, toAmount *big.Int, expectedAmount *big.Int, callees []common.Address, exchangeData []byte, startIndexes []*big.Int, values []*big.Int, beneficiary common.Address, referrer string, useReduxToken bool) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.SimpleSwap(&_IAugustusSwapper.TransactOpts, fromToken, toToken, fromAmount, toAmount, expectedAmount, callees, exchangeData, startIndexes, values, beneficiary, referrer, useReduxToken)
}

// SwapOnUniswap is a paid mutator transaction binding the contract method 0x58b9d179.
//
// Solidity: function swapOnUniswap(uint256 amountIn, uint256 amountOutMin, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperTransactor) SwapOnUniswap(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.contract.Transact(opts, "swapOnUniswap", amountIn, amountOutMin, path, referrer)
}

// SwapOnUniswap is a paid mutator transaction binding the contract method 0x58b9d179.
//
// Solidity: function swapOnUniswap(uint256 amountIn, uint256 amountOutMin, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperSession) SwapOnUniswap(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.SwapOnUniswap(&_IAugustusSwapper.TransactOpts, amountIn, amountOutMin, path, referrer)
}

// SwapOnUniswap is a paid mutator transaction binding the contract method 0x58b9d179.
//
// Solidity: function swapOnUniswap(uint256 amountIn, uint256 amountOutMin, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperTransactorSession) SwapOnUniswap(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.SwapOnUniswap(&_IAugustusSwapper.TransactOpts, amountIn, amountOutMin, path, referrer)
}

// SwapOnUniswapFork is a paid mutator transaction binding the contract method 0x0863b7ac.
//
// Solidity: function swapOnUniswapFork(address factory, bytes32 initCode, uint256 amountIn, uint256 amountOutMin, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperTransactor) SwapOnUniswapFork(opts *bind.TransactOpts, factory common.Address, initCode [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.contract.Transact(opts, "swapOnUniswapFork", factory, initCode, amountIn, amountOutMin, path, referrer)
}

// SwapOnUniswapFork is a paid mutator transaction binding the contract method 0x0863b7ac.
//
// Solidity: function swapOnUniswapFork(address factory, bytes32 initCode, uint256 amountIn, uint256 amountOutMin, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperSession) SwapOnUniswapFork(factory common.Address, initCode [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.SwapOnUniswapFork(&_IAugustusSwapper.TransactOpts, factory, initCode, amountIn, amountOutMin, path, referrer)
}

// SwapOnUniswapFork is a paid mutator transaction binding the contract method 0x0863b7ac.
//
// Solidity: function swapOnUniswapFork(address factory, bytes32 initCode, uint256 amountIn, uint256 amountOutMin, address[] path, uint8 referrer) payable returns()
func (_IAugustusSwapper *IAugustusSwapperTransactorSession) SwapOnUniswapFork(factory common.Address, initCode [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, referrer uint8) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.SwapOnUniswapFork(&_IAugustusSwapper.TransactOpts, factory, initCode, amountIn, amountOutMin, path, referrer)
}
