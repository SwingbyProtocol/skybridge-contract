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

// UtilsAdapter is an auto generated low-level Go binding around an user-defined struct.
type UtilsAdapter struct {
	Adapter    common.Address
	Percent    *big.Int
	NetworkFee *big.Int
	Route      []UtilsRoute
}

// UtilsMegaSwapPath is an auto generated low-level Go binding around an user-defined struct.
type UtilsMegaSwapPath struct {
	FromAmountPercent *big.Int
	Path              []UtilsPath
}

// UtilsMegaSwapSellData is an auto generated low-level Go binding around an user-defined struct.
type UtilsMegaSwapSellData struct {
	FromToken      common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	Beneficiary    common.Address
	Path           []UtilsMegaSwapPath
	Partner        common.Address
	FeePercent     *big.Int
	Permit         []byte
	Deadline       *big.Int
	Uuid           [16]byte
}

// UtilsPath is an auto generated low-level Go binding around an user-defined struct.
type UtilsPath struct {
	To              common.Address
	TotalNetworkFee *big.Int
	Adapters        []UtilsAdapter
}

// UtilsRoute is an auto generated low-level Go binding around an user-defined struct.
type UtilsRoute struct {
	Index          *big.Int
	TargetExchange common.Address
	Percent        *big.Int
	Payload        []byte
	NetworkFee     *big.Int
}

// UtilsSellData is an auto generated low-level Go binding around an user-defined struct.
type UtilsSellData struct {
	FromToken      common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	Beneficiary    common.Address
	Path           []UtilsPath
	Partner        common.Address
	FeePercent     *big.Int
	Permit         []byte
	Deadline       *big.Int
	Uuid           [16]byte
}

// UtilsSimpleData is an auto generated low-level Go binding around an user-defined struct.
type UtilsSimpleData struct {
	FromToken      common.Address
	ToToken        common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	Callees        []common.Address
	ExchangeData   []byte
	StartIndexes   []*big.Int
	Values         []*big.Int
	Beneficiary    common.Address
	Partner        common.Address
	FeePercent     *big.Int
	Permit         []byte
	Deadline       *big.Int
	Uuid           [16]byte
}

// IParaswapMetaData contains all meta data concerning the IParaswap contract.
var IParaswapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"buyOnUniswap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"buyOnUniswapFork\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"buyOnUniswapV2Fork\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromAmountPercent\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalNetworkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetExchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"internalType\":\"structUtils.Route[]\",\"name\":\"route\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Adapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Path[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.MegaSwapPath[]\",\"name\":\"path\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.MegaSwapSellData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"megaSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalNetworkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetExchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"internalType\":\"structUtils.Route[]\",\"name\":\"route\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Adapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Path[]\",\"name\":\"path\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.SellData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"multiSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromAmountPercent\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalNetworkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetExchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"internalType\":\"structUtils.Route[]\",\"name\":\"route\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Adapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Path[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.MegaSwapPath[]\",\"name\":\"path\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.MegaSwapSellData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"protectedMegaSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalNetworkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"targetExchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"networkFee\",\"type\":\"uint256\"}],\"internalType\":\"structUtils.Route[]\",\"name\":\"route\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Adapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structUtils.Path[]\",\"name\":\"path\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.SellData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"protectedMultiSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"callees\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"exchangeData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"startIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.SimpleData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"protectedSimpleBuy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"callees\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"exchangeData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"startIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.SimpleData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"protectedSimpleSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"callees\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"exchangeData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"startIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.SimpleData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"simpleBuy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"callees\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"exchangeData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"startIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.SimpleData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"simpleSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"swapOnUniswap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"initCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"swapOnUniswapFork\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"swapOnUniswapV2Fork\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"swapOnZeroXv2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"swapOnZeroXv4\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IParaswapABI is the input ABI used to generate the binding from.
// Deprecated: Use IParaswapMetaData.ABI instead.
var IParaswapABI = IParaswapMetaData.ABI

// IParaswap is an auto generated Go binding around an Ethereum contract.
type IParaswap struct {
	IParaswapCaller     // Read-only binding to the contract
	IParaswapTransactor // Write-only binding to the contract
	IParaswapFilterer   // Log filterer for contract events
}

// IParaswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type IParaswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParaswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IParaswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParaswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IParaswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IParaswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IParaswapSession struct {
	Contract     *IParaswap        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IParaswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IParaswapCallerSession struct {
	Contract *IParaswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IParaswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IParaswapTransactorSession struct {
	Contract     *IParaswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IParaswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type IParaswapRaw struct {
	Contract *IParaswap // Generic contract binding to access the raw methods on
}

// IParaswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IParaswapCallerRaw struct {
	Contract *IParaswapCaller // Generic read-only contract binding to access the raw methods on
}

// IParaswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IParaswapTransactorRaw struct {
	Contract *IParaswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIParaswap creates a new instance of IParaswap, bound to a specific deployed contract.
func NewIParaswap(address common.Address, backend bind.ContractBackend) (*IParaswap, error) {
	contract, err := bindIParaswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IParaswap{IParaswapCaller: IParaswapCaller{contract: contract}, IParaswapTransactor: IParaswapTransactor{contract: contract}, IParaswapFilterer: IParaswapFilterer{contract: contract}}, nil
}

// NewIParaswapCaller creates a new read-only instance of IParaswap, bound to a specific deployed contract.
func NewIParaswapCaller(address common.Address, caller bind.ContractCaller) (*IParaswapCaller, error) {
	contract, err := bindIParaswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IParaswapCaller{contract: contract}, nil
}

// NewIParaswapTransactor creates a new write-only instance of IParaswap, bound to a specific deployed contract.
func NewIParaswapTransactor(address common.Address, transactor bind.ContractTransactor) (*IParaswapTransactor, error) {
	contract, err := bindIParaswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IParaswapTransactor{contract: contract}, nil
}

// NewIParaswapFilterer creates a new log filterer instance of IParaswap, bound to a specific deployed contract.
func NewIParaswapFilterer(address common.Address, filterer bind.ContractFilterer) (*IParaswapFilterer, error) {
	contract, err := bindIParaswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IParaswapFilterer{contract: contract}, nil
}

// bindIParaswap binds a generic wrapper to an already deployed contract.
func bindIParaswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IParaswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IParaswap *IParaswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IParaswap.Contract.IParaswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IParaswap *IParaswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IParaswap.Contract.IParaswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IParaswap *IParaswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IParaswap.Contract.IParaswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IParaswap *IParaswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IParaswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IParaswap *IParaswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IParaswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IParaswap *IParaswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IParaswap.Contract.contract.Transact(opts, method, params...)
}

// BuyOnUniswap is a paid mutator transaction binding the contract method 0x935fb84b.
//
// Solidity: function buyOnUniswap(uint256 amountInMax, uint256 amountOut, address[] path) payable returns()
func (_IParaswap *IParaswapTransactor) BuyOnUniswap(opts *bind.TransactOpts, amountInMax *big.Int, amountOut *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "buyOnUniswap", amountInMax, amountOut, path)
}

// BuyOnUniswap is a paid mutator transaction binding the contract method 0x935fb84b.
//
// Solidity: function buyOnUniswap(uint256 amountInMax, uint256 amountOut, address[] path) payable returns()
func (_IParaswap *IParaswapSession) BuyOnUniswap(amountInMax *big.Int, amountOut *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.Contract.BuyOnUniswap(&_IParaswap.TransactOpts, amountInMax, amountOut, path)
}

// BuyOnUniswap is a paid mutator transaction binding the contract method 0x935fb84b.
//
// Solidity: function buyOnUniswap(uint256 amountInMax, uint256 amountOut, address[] path) payable returns()
func (_IParaswap *IParaswapTransactorSession) BuyOnUniswap(amountInMax *big.Int, amountOut *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.Contract.BuyOnUniswap(&_IParaswap.TransactOpts, amountInMax, amountOut, path)
}

// BuyOnUniswapFork is a paid mutator transaction binding the contract method 0xc03786b0.
//
// Solidity: function buyOnUniswapFork(address factory, bytes32 initCode, uint256 amountInMax, uint256 amountOut, address[] path) payable returns()
func (_IParaswap *IParaswapTransactor) BuyOnUniswapFork(opts *bind.TransactOpts, factory common.Address, initCode [32]byte, amountInMax *big.Int, amountOut *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "buyOnUniswapFork", factory, initCode, amountInMax, amountOut, path)
}

// BuyOnUniswapFork is a paid mutator transaction binding the contract method 0xc03786b0.
//
// Solidity: function buyOnUniswapFork(address factory, bytes32 initCode, uint256 amountInMax, uint256 amountOut, address[] path) payable returns()
func (_IParaswap *IParaswapSession) BuyOnUniswapFork(factory common.Address, initCode [32]byte, amountInMax *big.Int, amountOut *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.Contract.BuyOnUniswapFork(&_IParaswap.TransactOpts, factory, initCode, amountInMax, amountOut, path)
}

// BuyOnUniswapFork is a paid mutator transaction binding the contract method 0xc03786b0.
//
// Solidity: function buyOnUniswapFork(address factory, bytes32 initCode, uint256 amountInMax, uint256 amountOut, address[] path) payable returns()
func (_IParaswap *IParaswapTransactorSession) BuyOnUniswapFork(factory common.Address, initCode [32]byte, amountInMax *big.Int, amountOut *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.Contract.BuyOnUniswapFork(&_IParaswap.TransactOpts, factory, initCode, amountInMax, amountOut, path)
}

// BuyOnUniswapV2Fork is a paid mutator transaction binding the contract method 0xb2f1e6db.
//
// Solidity: function buyOnUniswapV2Fork(address tokenIn, uint256 amountInMax, uint256 amountOut, address weth, uint256[] pools) payable returns()
func (_IParaswap *IParaswapTransactor) BuyOnUniswapV2Fork(opts *bind.TransactOpts, tokenIn common.Address, amountInMax *big.Int, amountOut *big.Int, weth common.Address, pools []*big.Int) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "buyOnUniswapV2Fork", tokenIn, amountInMax, amountOut, weth, pools)
}

// BuyOnUniswapV2Fork is a paid mutator transaction binding the contract method 0xb2f1e6db.
//
// Solidity: function buyOnUniswapV2Fork(address tokenIn, uint256 amountInMax, uint256 amountOut, address weth, uint256[] pools) payable returns()
func (_IParaswap *IParaswapSession) BuyOnUniswapV2Fork(tokenIn common.Address, amountInMax *big.Int, amountOut *big.Int, weth common.Address, pools []*big.Int) (*types.Transaction, error) {
	return _IParaswap.Contract.BuyOnUniswapV2Fork(&_IParaswap.TransactOpts, tokenIn, amountInMax, amountOut, weth, pools)
}

// BuyOnUniswapV2Fork is a paid mutator transaction binding the contract method 0xb2f1e6db.
//
// Solidity: function buyOnUniswapV2Fork(address tokenIn, uint256 amountInMax, uint256 amountOut, address weth, uint256[] pools) payable returns()
func (_IParaswap *IParaswapTransactorSession) BuyOnUniswapV2Fork(tokenIn common.Address, amountInMax *big.Int, amountOut *big.Int, weth common.Address, pools []*big.Int) (*types.Transaction, error) {
	return _IParaswap.Contract.BuyOnUniswapV2Fork(&_IParaswap.TransactOpts, tokenIn, amountInMax, amountOut, weth, pools)
}

// MegaSwap is a paid mutator transaction binding the contract method 0x46c67b6d.
//
// Solidity: function megaSwap((address,uint256,uint256,uint256,address,(uint256,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapTransactor) MegaSwap(opts *bind.TransactOpts, data UtilsMegaSwapSellData) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "megaSwap", data)
}

// MegaSwap is a paid mutator transaction binding the contract method 0x46c67b6d.
//
// Solidity: function megaSwap((address,uint256,uint256,uint256,address,(uint256,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapSession) MegaSwap(data UtilsMegaSwapSellData) (*types.Transaction, error) {
	return _IParaswap.Contract.MegaSwap(&_IParaswap.TransactOpts, data)
}

// MegaSwap is a paid mutator transaction binding the contract method 0x46c67b6d.
//
// Solidity: function megaSwap((address,uint256,uint256,uint256,address,(uint256,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapTransactorSession) MegaSwap(data UtilsMegaSwapSellData) (*types.Transaction, error) {
	return _IParaswap.Contract.MegaSwap(&_IParaswap.TransactOpts, data)
}

// MultiSwap is a paid mutator transaction binding the contract method 0xa94e78ef.
//
// Solidity: function multiSwap((address,uint256,uint256,uint256,address,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapTransactor) MultiSwap(opts *bind.TransactOpts, data UtilsSellData) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "multiSwap", data)
}

// MultiSwap is a paid mutator transaction binding the contract method 0xa94e78ef.
//
// Solidity: function multiSwap((address,uint256,uint256,uint256,address,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapSession) MultiSwap(data UtilsSellData) (*types.Transaction, error) {
	return _IParaswap.Contract.MultiSwap(&_IParaswap.TransactOpts, data)
}

// MultiSwap is a paid mutator transaction binding the contract method 0xa94e78ef.
//
// Solidity: function multiSwap((address,uint256,uint256,uint256,address,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapTransactorSession) MultiSwap(data UtilsSellData) (*types.Transaction, error) {
	return _IParaswap.Contract.MultiSwap(&_IParaswap.TransactOpts, data)
}

// ProtectedMegaSwap is a paid mutator transaction binding the contract method 0x37809db4.
//
// Solidity: function protectedMegaSwap((address,uint256,uint256,uint256,address,(uint256,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapTransactor) ProtectedMegaSwap(opts *bind.TransactOpts, data UtilsMegaSwapSellData) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "protectedMegaSwap", data)
}

// ProtectedMegaSwap is a paid mutator transaction binding the contract method 0x37809db4.
//
// Solidity: function protectedMegaSwap((address,uint256,uint256,uint256,address,(uint256,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapSession) ProtectedMegaSwap(data UtilsMegaSwapSellData) (*types.Transaction, error) {
	return _IParaswap.Contract.ProtectedMegaSwap(&_IParaswap.TransactOpts, data)
}

// ProtectedMegaSwap is a paid mutator transaction binding the contract method 0x37809db4.
//
// Solidity: function protectedMegaSwap((address,uint256,uint256,uint256,address,(uint256,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapTransactorSession) ProtectedMegaSwap(data UtilsMegaSwapSellData) (*types.Transaction, error) {
	return _IParaswap.Contract.ProtectedMegaSwap(&_IParaswap.TransactOpts, data)
}

// ProtectedMultiSwap is a paid mutator transaction binding the contract method 0x2478ba3e.
//
// Solidity: function protectedMultiSwap((address,uint256,uint256,uint256,address,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapTransactor) ProtectedMultiSwap(opts *bind.TransactOpts, data UtilsSellData) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "protectedMultiSwap", data)
}

// ProtectedMultiSwap is a paid mutator transaction binding the contract method 0x2478ba3e.
//
// Solidity: function protectedMultiSwap((address,uint256,uint256,uint256,address,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapSession) ProtectedMultiSwap(data UtilsSellData) (*types.Transaction, error) {
	return _IParaswap.Contract.ProtectedMultiSwap(&_IParaswap.TransactOpts, data)
}

// ProtectedMultiSwap is a paid mutator transaction binding the contract method 0x2478ba3e.
//
// Solidity: function protectedMultiSwap((address,uint256,uint256,uint256,address,(address,uint256,(address,uint256,uint256,(uint256,address,uint256,bytes,uint256)[])[])[],address,uint256,bytes,uint256,bytes16) data) payable returns(uint256)
func (_IParaswap *IParaswapTransactorSession) ProtectedMultiSwap(data UtilsSellData) (*types.Transaction, error) {
	return _IParaswap.Contract.ProtectedMultiSwap(&_IParaswap.TransactOpts, data)
}

// ProtectedSimpleBuy is a paid mutator transaction binding the contract method 0xfab13517.
//
// Solidity: function protectedSimpleBuy((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns()
func (_IParaswap *IParaswapTransactor) ProtectedSimpleBuy(opts *bind.TransactOpts, data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "protectedSimpleBuy", data)
}

// ProtectedSimpleBuy is a paid mutator transaction binding the contract method 0xfab13517.
//
// Solidity: function protectedSimpleBuy((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns()
func (_IParaswap *IParaswapSession) ProtectedSimpleBuy(data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.Contract.ProtectedSimpleBuy(&_IParaswap.TransactOpts, data)
}

// ProtectedSimpleBuy is a paid mutator transaction binding the contract method 0xfab13517.
//
// Solidity: function protectedSimpleBuy((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns()
func (_IParaswap *IParaswapTransactorSession) ProtectedSimpleBuy(data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.Contract.ProtectedSimpleBuy(&_IParaswap.TransactOpts, data)
}

// ProtectedSimpleSwap is a paid mutator transaction binding the contract method 0xa8795e3d.
//
// Solidity: function protectedSimpleSwap((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns(uint256 receivedAmount)
func (_IParaswap *IParaswapTransactor) ProtectedSimpleSwap(opts *bind.TransactOpts, data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "protectedSimpleSwap", data)
}

// ProtectedSimpleSwap is a paid mutator transaction binding the contract method 0xa8795e3d.
//
// Solidity: function protectedSimpleSwap((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns(uint256 receivedAmount)
func (_IParaswap *IParaswapSession) ProtectedSimpleSwap(data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.Contract.ProtectedSimpleSwap(&_IParaswap.TransactOpts, data)
}

// ProtectedSimpleSwap is a paid mutator transaction binding the contract method 0xa8795e3d.
//
// Solidity: function protectedSimpleSwap((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns(uint256 receivedAmount)
func (_IParaswap *IParaswapTransactorSession) ProtectedSimpleSwap(data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.Contract.ProtectedSimpleSwap(&_IParaswap.TransactOpts, data)
}

// SimpleBuy is a paid mutator transaction binding the contract method 0x2298207a.
//
// Solidity: function simpleBuy((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns()
func (_IParaswap *IParaswapTransactor) SimpleBuy(opts *bind.TransactOpts, data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "simpleBuy", data)
}

// SimpleBuy is a paid mutator transaction binding the contract method 0x2298207a.
//
// Solidity: function simpleBuy((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns()
func (_IParaswap *IParaswapSession) SimpleBuy(data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.Contract.SimpleBuy(&_IParaswap.TransactOpts, data)
}

// SimpleBuy is a paid mutator transaction binding the contract method 0x2298207a.
//
// Solidity: function simpleBuy((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns()
func (_IParaswap *IParaswapTransactorSession) SimpleBuy(data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.Contract.SimpleBuy(&_IParaswap.TransactOpts, data)
}

// SimpleSwap is a paid mutator transaction binding the contract method 0x54e3f31b.
//
// Solidity: function simpleSwap((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns(uint256 receivedAmount)
func (_IParaswap *IParaswapTransactor) SimpleSwap(opts *bind.TransactOpts, data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "simpleSwap", data)
}

// SimpleSwap is a paid mutator transaction binding the contract method 0x54e3f31b.
//
// Solidity: function simpleSwap((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns(uint256 receivedAmount)
func (_IParaswap *IParaswapSession) SimpleSwap(data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.Contract.SimpleSwap(&_IParaswap.TransactOpts, data)
}

// SimpleSwap is a paid mutator transaction binding the contract method 0x54e3f31b.
//
// Solidity: function simpleSwap((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) data) payable returns(uint256 receivedAmount)
func (_IParaswap *IParaswapTransactorSession) SimpleSwap(data UtilsSimpleData) (*types.Transaction, error) {
	return _IParaswap.Contract.SimpleSwap(&_IParaswap.TransactOpts, data)
}

// SwapOnUniswap is a paid mutator transaction binding the contract method 0x54840d1a.
//
// Solidity: function swapOnUniswap(uint256 amountIn, uint256 amountOutMin, address[] path) payable returns()
func (_IParaswap *IParaswapTransactor) SwapOnUniswap(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "swapOnUniswap", amountIn, amountOutMin, path)
}

// SwapOnUniswap is a paid mutator transaction binding the contract method 0x54840d1a.
//
// Solidity: function swapOnUniswap(uint256 amountIn, uint256 amountOutMin, address[] path) payable returns()
func (_IParaswap *IParaswapSession) SwapOnUniswap(amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.Contract.SwapOnUniswap(&_IParaswap.TransactOpts, amountIn, amountOutMin, path)
}

// SwapOnUniswap is a paid mutator transaction binding the contract method 0x54840d1a.
//
// Solidity: function swapOnUniswap(uint256 amountIn, uint256 amountOutMin, address[] path) payable returns()
func (_IParaswap *IParaswapTransactorSession) SwapOnUniswap(amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.Contract.SwapOnUniswap(&_IParaswap.TransactOpts, amountIn, amountOutMin, path)
}

// SwapOnUniswapFork is a paid mutator transaction binding the contract method 0xf5661034.
//
// Solidity: function swapOnUniswapFork(address factory, bytes32 initCode, uint256 amountIn, uint256 amountOutMin, address[] path) payable returns()
func (_IParaswap *IParaswapTransactor) SwapOnUniswapFork(opts *bind.TransactOpts, factory common.Address, initCode [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "swapOnUniswapFork", factory, initCode, amountIn, amountOutMin, path)
}

// SwapOnUniswapFork is a paid mutator transaction binding the contract method 0xf5661034.
//
// Solidity: function swapOnUniswapFork(address factory, bytes32 initCode, uint256 amountIn, uint256 amountOutMin, address[] path) payable returns()
func (_IParaswap *IParaswapSession) SwapOnUniswapFork(factory common.Address, initCode [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.Contract.SwapOnUniswapFork(&_IParaswap.TransactOpts, factory, initCode, amountIn, amountOutMin, path)
}

// SwapOnUniswapFork is a paid mutator transaction binding the contract method 0xf5661034.
//
// Solidity: function swapOnUniswapFork(address factory, bytes32 initCode, uint256 amountIn, uint256 amountOutMin, address[] path) payable returns()
func (_IParaswap *IParaswapTransactorSession) SwapOnUniswapFork(factory common.Address, initCode [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _IParaswap.Contract.SwapOnUniswapFork(&_IParaswap.TransactOpts, factory, initCode, amountIn, amountOutMin, path)
}

// SwapOnUniswapV2Fork is a paid mutator transaction binding the contract method 0x0b86a4c1.
//
// Solidity: function swapOnUniswapV2Fork(address tokenIn, uint256 amountIn, uint256 amountOutMin, address weth, uint256[] pools) payable returns()
func (_IParaswap *IParaswapTransactor) SwapOnUniswapV2Fork(opts *bind.TransactOpts, tokenIn common.Address, amountIn *big.Int, amountOutMin *big.Int, weth common.Address, pools []*big.Int) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "swapOnUniswapV2Fork", tokenIn, amountIn, amountOutMin, weth, pools)
}

// SwapOnUniswapV2Fork is a paid mutator transaction binding the contract method 0x0b86a4c1.
//
// Solidity: function swapOnUniswapV2Fork(address tokenIn, uint256 amountIn, uint256 amountOutMin, address weth, uint256[] pools) payable returns()
func (_IParaswap *IParaswapSession) SwapOnUniswapV2Fork(tokenIn common.Address, amountIn *big.Int, amountOutMin *big.Int, weth common.Address, pools []*big.Int) (*types.Transaction, error) {
	return _IParaswap.Contract.SwapOnUniswapV2Fork(&_IParaswap.TransactOpts, tokenIn, amountIn, amountOutMin, weth, pools)
}

// SwapOnUniswapV2Fork is a paid mutator transaction binding the contract method 0x0b86a4c1.
//
// Solidity: function swapOnUniswapV2Fork(address tokenIn, uint256 amountIn, uint256 amountOutMin, address weth, uint256[] pools) payable returns()
func (_IParaswap *IParaswapTransactorSession) SwapOnUniswapV2Fork(tokenIn common.Address, amountIn *big.Int, amountOutMin *big.Int, weth common.Address, pools []*big.Int) (*types.Transaction, error) {
	return _IParaswap.Contract.SwapOnUniswapV2Fork(&_IParaswap.TransactOpts, tokenIn, amountIn, amountOutMin, weth, pools)
}

// SwapOnZeroXv2 is a paid mutator transaction binding the contract method 0x81033120.
//
// Solidity: function swapOnZeroXv2(address fromToken, address toToken, uint256 fromAmount, uint256 amountOutMin, address exchange, bytes payload) payable returns()
func (_IParaswap *IParaswapTransactor) SwapOnZeroXv2(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int, amountOutMin *big.Int, exchange common.Address, payload []byte) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "swapOnZeroXv2", fromToken, toToken, fromAmount, amountOutMin, exchange, payload)
}

// SwapOnZeroXv2 is a paid mutator transaction binding the contract method 0x81033120.
//
// Solidity: function swapOnZeroXv2(address fromToken, address toToken, uint256 fromAmount, uint256 amountOutMin, address exchange, bytes payload) payable returns()
func (_IParaswap *IParaswapSession) SwapOnZeroXv2(fromToken common.Address, toToken common.Address, fromAmount *big.Int, amountOutMin *big.Int, exchange common.Address, payload []byte) (*types.Transaction, error) {
	return _IParaswap.Contract.SwapOnZeroXv2(&_IParaswap.TransactOpts, fromToken, toToken, fromAmount, amountOutMin, exchange, payload)
}

// SwapOnZeroXv2 is a paid mutator transaction binding the contract method 0x81033120.
//
// Solidity: function swapOnZeroXv2(address fromToken, address toToken, uint256 fromAmount, uint256 amountOutMin, address exchange, bytes payload) payable returns()
func (_IParaswap *IParaswapTransactorSession) SwapOnZeroXv2(fromToken common.Address, toToken common.Address, fromAmount *big.Int, amountOutMin *big.Int, exchange common.Address, payload []byte) (*types.Transaction, error) {
	return _IParaswap.Contract.SwapOnZeroXv2(&_IParaswap.TransactOpts, fromToken, toToken, fromAmount, amountOutMin, exchange, payload)
}

// SwapOnZeroXv4 is a paid mutator transaction binding the contract method 0x64466805.
//
// Solidity: function swapOnZeroXv4(address fromToken, address toToken, uint256 fromAmount, uint256 amountOutMin, address exchange, bytes payload) payable returns()
func (_IParaswap *IParaswapTransactor) SwapOnZeroXv4(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromAmount *big.Int, amountOutMin *big.Int, exchange common.Address, payload []byte) (*types.Transaction, error) {
	return _IParaswap.contract.Transact(opts, "swapOnZeroXv4", fromToken, toToken, fromAmount, amountOutMin, exchange, payload)
}

// SwapOnZeroXv4 is a paid mutator transaction binding the contract method 0x64466805.
//
// Solidity: function swapOnZeroXv4(address fromToken, address toToken, uint256 fromAmount, uint256 amountOutMin, address exchange, bytes payload) payable returns()
func (_IParaswap *IParaswapSession) SwapOnZeroXv4(fromToken common.Address, toToken common.Address, fromAmount *big.Int, amountOutMin *big.Int, exchange common.Address, payload []byte) (*types.Transaction, error) {
	return _IParaswap.Contract.SwapOnZeroXv4(&_IParaswap.TransactOpts, fromToken, toToken, fromAmount, amountOutMin, exchange, payload)
}

// SwapOnZeroXv4 is a paid mutator transaction binding the contract method 0x64466805.
//
// Solidity: function swapOnZeroXv4(address fromToken, address toToken, uint256 fromAmount, uint256 amountOutMin, address exchange, bytes payload) payable returns()
func (_IParaswap *IParaswapTransactorSession) SwapOnZeroXv4(fromToken common.Address, toToken common.Address, fromAmount *big.Int, amountOutMin *big.Int, exchange common.Address, payload []byte) (*types.Transaction, error) {
	return _IParaswap.Contract.SwapOnZeroXv4(&_IParaswap.TransactOpts, fromToken, toToken, fromAmount, amountOutMin, exchange, payload)
}
