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

// SwapContractV1spPendingTx is an auto generated low-level Go binding around an user-defined struct.
type SwapContractV1spPendingTx struct {
	SwapID     [32]byte
	DestAddr   string
	RefundAddr common.Address
	AmountWBTC *big.Int
	Timestamp  *big.Int
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

// SwapContractV1MetaData contains all meta data concerning the SwapContractV1 contract.
var SwapContractV1MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfFloat\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPriceLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"}],\"name\":\"BurnLPTokensForFloat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardLPTsForNodes\",\"type\":\"uint256\"}],\"name\":\"DistributeNodeRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfFloat\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPriceLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"}],\"name\":\"IssueLPTokensForFloat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feesToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLPTokensForNode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPriceLP\",\"type\":\"uint256\"}],\"name\":\"RewardsCollection\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"SwapID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"DestAddr\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"RefundAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"AmountWBTC\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Timestamp\",\"type\":\"uint256\"}],\"name\":\"SwapTokensToBTC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Timestamp\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BTCT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeNodeCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_nodes\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_isRemoved\",\"type\":\"bool[]\"},{\"internalType\":\"uint8\",\"name\":\"_churnedInCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tssThreshold\",\"type\":\"uint8\"}],\"name\":\"churn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"churnedInCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_incomingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_spenders\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_swapAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"_isUpdatelimitBTCForSPFlow2\",\"type\":\"bool\"}],\"name\":\"collectSwapFeesForBTC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentPriceLP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nowPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"}],\"name\":\"getFloatReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_btct\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sbBTCPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_params\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swapRewards\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_existingBTCFloat\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ip\",\"outputs\":[{\"internalType\":\"contractIParams\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isNodeStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"isTxUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitBTCForSPFlow2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIBurnableToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_addressesAndAmounts\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalSwapped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_usedTxIds\",\"type\":\"bytes32[]\"}],\"name\":\"multiRecordSkyPoolsTX\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"_addressesAndAmounts\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalSwapped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_redeemedFloatTxIds\",\"type\":\"bytes32[]\"}],\"name\":\"multiTransferERC20TightlyPacked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oldestActiveIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paraswapAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_addressesAndAmountOfFloat\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordIncomingFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_addressesAndAmountOfLPtoken\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_minerFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordOutcomingFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalSwapped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_usedTxIds\",\"type\":\"bytes32[]\"}],\"name\":\"recordSkyPoolsTX\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minerFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordUTXOSweepMinerFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeemERC20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeemEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sbBTCPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSwapped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_redeemedFloatTxIds\",\"type\":\"bytes32[]\"}],\"name\":\"singleTransferERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spCleanUpOldTXs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"spDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"callees\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"exchangeData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"startIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.SimpleData\",\"name\":\"_data\",\"type\":\"tuple\"}],\"name\":\"spFlow1SimpleSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_fork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_initCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"}],\"name\":\"spFlow1Uniswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_destinationAddressForBTC\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"callees\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"exchangeData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"startIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.SimpleData\",\"name\":\"_data\",\"type\":\"tuple\"}],\"name\":\"spFlow2SimpleSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_destinationAddressForBTC\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_fork\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_initCode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"}],\"name\":\"spFlow2Uniswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spGetPendingSwaps\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"SwapID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"DestAddr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"RefundAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"AmountWBTC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structSwapContractV1.spPendingTx[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spPendingTXs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"SwapID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"DestAddr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"RefundAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"AmountWBTC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sw\",\"outputs\":[{\"internalType\":\"contractISwapRewards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SwapContractV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapContractV1MetaData.ABI instead.
var SwapContractV1ABI = SwapContractV1MetaData.ABI

// SwapContractV1 is an auto generated Go binding around an Ethereum contract.
type SwapContractV1 struct {
	SwapContractV1Caller     // Read-only binding to the contract
	SwapContractV1Transactor // Write-only binding to the contract
	SwapContractV1Filterer   // Log filterer for contract events
}

// SwapContractV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type SwapContractV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapContractV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapContractV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapContractV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapContractV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapContractV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapContractV1Session struct {
	Contract     *SwapContractV1   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapContractV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapContractV1CallerSession struct {
	Contract *SwapContractV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SwapContractV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapContractV1TransactorSession struct {
	Contract     *SwapContractV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SwapContractV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type SwapContractV1Raw struct {
	Contract *SwapContractV1 // Generic contract binding to access the raw methods on
}

// SwapContractV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapContractV1CallerRaw struct {
	Contract *SwapContractV1Caller // Generic read-only contract binding to access the raw methods on
}

// SwapContractV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapContractV1TransactorRaw struct {
	Contract *SwapContractV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapContractV1 creates a new instance of SwapContractV1, bound to a specific deployed contract.
func NewSwapContractV1(address common.Address, backend bind.ContractBackend) (*SwapContractV1, error) {
	contract, err := bindSwapContractV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapContractV1{SwapContractV1Caller: SwapContractV1Caller{contract: contract}, SwapContractV1Transactor: SwapContractV1Transactor{contract: contract}, SwapContractV1Filterer: SwapContractV1Filterer{contract: contract}}, nil
}

// NewSwapContractV1Caller creates a new read-only instance of SwapContractV1, bound to a specific deployed contract.
func NewSwapContractV1Caller(address common.Address, caller bind.ContractCaller) (*SwapContractV1Caller, error) {
	contract, err := bindSwapContractV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapContractV1Caller{contract: contract}, nil
}

// NewSwapContractV1Transactor creates a new write-only instance of SwapContractV1, bound to a specific deployed contract.
func NewSwapContractV1Transactor(address common.Address, transactor bind.ContractTransactor) (*SwapContractV1Transactor, error) {
	contract, err := bindSwapContractV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapContractV1Transactor{contract: contract}, nil
}

// NewSwapContractV1Filterer creates a new log filterer instance of SwapContractV1, bound to a specific deployed contract.
func NewSwapContractV1Filterer(address common.Address, filterer bind.ContractFilterer) (*SwapContractV1Filterer, error) {
	contract, err := bindSwapContractV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapContractV1Filterer{contract: contract}, nil
}

// bindSwapContractV1 binds a generic wrapper to an already deployed contract.
func bindSwapContractV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapContractV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapContractV1 *SwapContractV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapContractV1.Contract.SwapContractV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapContractV1 *SwapContractV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SwapContractV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapContractV1 *SwapContractV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SwapContractV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapContractV1 *SwapContractV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapContractV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapContractV1 *SwapContractV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapContractV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapContractV1 *SwapContractV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapContractV1.Contract.contract.Transact(opts, method, params...)
}

// BTCTADDR is a free data retrieval call binding the contract method 0x0f909486.
//
// Solidity: function BTCT_ADDR() view returns(address)
func (_SwapContractV1 *SwapContractV1Caller) BTCTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "BTCT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BTCTADDR is a free data retrieval call binding the contract method 0x0f909486.
//
// Solidity: function BTCT_ADDR() view returns(address)
func (_SwapContractV1 *SwapContractV1Session) BTCTADDR() (common.Address, error) {
	return _SwapContractV1.Contract.BTCTADDR(&_SwapContractV1.CallOpts)
}

// BTCTADDR is a free data retrieval call binding the contract method 0x0f909486.
//
// Solidity: function BTCT_ADDR() view returns(address)
func (_SwapContractV1 *SwapContractV1CallerSession) BTCTADDR() (common.Address, error) {
	return _SwapContractV1.Contract.BTCTADDR(&_SwapContractV1.CallOpts)
}

// ActiveNodeCount is a free data retrieval call binding the contract method 0x75340815.
//
// Solidity: function activeNodeCount() view returns(uint8)
func (_SwapContractV1 *SwapContractV1Caller) ActiveNodeCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "activeNodeCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ActiveNodeCount is a free data retrieval call binding the contract method 0x75340815.
//
// Solidity: function activeNodeCount() view returns(uint8)
func (_SwapContractV1 *SwapContractV1Session) ActiveNodeCount() (uint8, error) {
	return _SwapContractV1.Contract.ActiveNodeCount(&_SwapContractV1.CallOpts)
}

// ActiveNodeCount is a free data retrieval call binding the contract method 0x75340815.
//
// Solidity: function activeNodeCount() view returns(uint8)
func (_SwapContractV1 *SwapContractV1CallerSession) ActiveNodeCount() (uint8, error) {
	return _SwapContractV1.Contract.ActiveNodeCount(&_SwapContractV1.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _token, address _user) view returns(uint256)
func (_SwapContractV1 *SwapContractV1Caller) BalanceOf(opts *bind.CallOpts, _token common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "balanceOf", _token, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _token, address _user) view returns(uint256)
func (_SwapContractV1 *SwapContractV1Session) BalanceOf(_token common.Address, _user common.Address) (*big.Int, error) {
	return _SwapContractV1.Contract.BalanceOf(&_SwapContractV1.CallOpts, _token, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _token, address _user) view returns(uint256)
func (_SwapContractV1 *SwapContractV1CallerSession) BalanceOf(_token common.Address, _user common.Address) (*big.Int, error) {
	return _SwapContractV1.Contract.BalanceOf(&_SwapContractV1.CallOpts, _token, _user)
}

// ChurnedInCount is a free data retrieval call binding the contract method 0x0089356f.
//
// Solidity: function churnedInCount() view returns(uint8)
func (_SwapContractV1 *SwapContractV1Caller) ChurnedInCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "churnedInCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ChurnedInCount is a free data retrieval call binding the contract method 0x0089356f.
//
// Solidity: function churnedInCount() view returns(uint8)
func (_SwapContractV1 *SwapContractV1Session) ChurnedInCount() (uint8, error) {
	return _SwapContractV1.Contract.ChurnedInCount(&_SwapContractV1.CallOpts)
}

// ChurnedInCount is a free data retrieval call binding the contract method 0x0089356f.
//
// Solidity: function churnedInCount() view returns(uint8)
func (_SwapContractV1 *SwapContractV1CallerSession) ChurnedInCount() (uint8, error) {
	return _SwapContractV1.Contract.ChurnedInCount(&_SwapContractV1.CallOpts)
}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_SwapContractV1 *SwapContractV1Caller) GetActiveNodes(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "getActiveNodes")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_SwapContractV1 *SwapContractV1Session) GetActiveNodes() ([]common.Address, error) {
	return _SwapContractV1.Contract.GetActiveNodes(&_SwapContractV1.CallOpts)
}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(address[])
func (_SwapContractV1 *SwapContractV1CallerSession) GetActiveNodes() ([]common.Address, error) {
	return _SwapContractV1.Contract.GetActiveNodes(&_SwapContractV1.CallOpts)
}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256 nowPrice)
func (_SwapContractV1 *SwapContractV1Caller) GetCurrentPriceLP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "getCurrentPriceLP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256 nowPrice)
func (_SwapContractV1 *SwapContractV1Session) GetCurrentPriceLP() (*big.Int, error) {
	return _SwapContractV1.Contract.GetCurrentPriceLP(&_SwapContractV1.CallOpts)
}

// GetCurrentPriceLP is a free data retrieval call binding the contract method 0x45137e27.
//
// Solidity: function getCurrentPriceLP() view returns(uint256 nowPrice)
func (_SwapContractV1 *SwapContractV1CallerSession) GetCurrentPriceLP() (*big.Int, error) {
	return _SwapContractV1.Contract.GetCurrentPriceLP(&_SwapContractV1.CallOpts)
}

// GetFloatReserve is a free data retrieval call binding the contract method 0xec482729.
//
// Solidity: function getFloatReserve(address _tokenA, address _tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_SwapContractV1 *SwapContractV1Caller) GetFloatReserve(opts *bind.CallOpts, _tokenA common.Address, _tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "getFloatReserve", _tokenA, _tokenB)

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
func (_SwapContractV1 *SwapContractV1Session) GetFloatReserve(_tokenA common.Address, _tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _SwapContractV1.Contract.GetFloatReserve(&_SwapContractV1.CallOpts, _tokenA, _tokenB)
}

// GetFloatReserve is a free data retrieval call binding the contract method 0xec482729.
//
// Solidity: function getFloatReserve(address _tokenA, address _tokenB) view returns(uint256 reserveA, uint256 reserveB)
func (_SwapContractV1 *SwapContractV1CallerSession) GetFloatReserve(_tokenA common.Address, _tokenB common.Address) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _SwapContractV1.Contract.GetFloatReserve(&_SwapContractV1.CallOpts, _tokenA, _tokenB)
}

// Ip is a free data retrieval call binding the contract method 0xd023420d.
//
// Solidity: function ip() view returns(address)
func (_SwapContractV1 *SwapContractV1Caller) Ip(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "ip")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ip is a free data retrieval call binding the contract method 0xd023420d.
//
// Solidity: function ip() view returns(address)
func (_SwapContractV1 *SwapContractV1Session) Ip() (common.Address, error) {
	return _SwapContractV1.Contract.Ip(&_SwapContractV1.CallOpts)
}

// Ip is a free data retrieval call binding the contract method 0xd023420d.
//
// Solidity: function ip() view returns(address)
func (_SwapContractV1 *SwapContractV1CallerSession) Ip() (common.Address, error) {
	return _SwapContractV1.Contract.Ip(&_SwapContractV1.CallOpts)
}

// IsNodeStake is a free data retrieval call binding the contract method 0xa742329d.
//
// Solidity: function isNodeStake(address _user) view returns(bool)
func (_SwapContractV1 *SwapContractV1Caller) IsNodeStake(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "isNodeStake", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNodeStake is a free data retrieval call binding the contract method 0xa742329d.
//
// Solidity: function isNodeStake(address _user) view returns(bool)
func (_SwapContractV1 *SwapContractV1Session) IsNodeStake(_user common.Address) (bool, error) {
	return _SwapContractV1.Contract.IsNodeStake(&_SwapContractV1.CallOpts, _user)
}

// IsNodeStake is a free data retrieval call binding the contract method 0xa742329d.
//
// Solidity: function isNodeStake(address _user) view returns(bool)
func (_SwapContractV1 *SwapContractV1CallerSession) IsNodeStake(_user common.Address) (bool, error) {
	return _SwapContractV1.Contract.IsNodeStake(&_SwapContractV1.CallOpts, _user)
}

// IsTxUsed is a free data retrieval call binding the contract method 0xe6ca2084.
//
// Solidity: function isTxUsed(bytes32 _txid) view returns(bool)
func (_SwapContractV1 *SwapContractV1Caller) IsTxUsed(opts *bind.CallOpts, _txid [32]byte) (bool, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "isTxUsed", _txid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTxUsed is a free data retrieval call binding the contract method 0xe6ca2084.
//
// Solidity: function isTxUsed(bytes32 _txid) view returns(bool)
func (_SwapContractV1 *SwapContractV1Session) IsTxUsed(_txid [32]byte) (bool, error) {
	return _SwapContractV1.Contract.IsTxUsed(&_SwapContractV1.CallOpts, _txid)
}

// IsTxUsed is a free data retrieval call binding the contract method 0xe6ca2084.
//
// Solidity: function isTxUsed(bytes32 _txid) view returns(bool)
func (_SwapContractV1 *SwapContractV1CallerSession) IsTxUsed(_txid [32]byte) (bool, error) {
	return _SwapContractV1.Contract.IsTxUsed(&_SwapContractV1.CallOpts, _txid)
}

// LimitBTCForSPFlow2 is a free data retrieval call binding the contract method 0xa827278b.
//
// Solidity: function limitBTCForSPFlow2() view returns(uint256)
func (_SwapContractV1 *SwapContractV1Caller) LimitBTCForSPFlow2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "limitBTCForSPFlow2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitBTCForSPFlow2 is a free data retrieval call binding the contract method 0xa827278b.
//
// Solidity: function limitBTCForSPFlow2() view returns(uint256)
func (_SwapContractV1 *SwapContractV1Session) LimitBTCForSPFlow2() (*big.Int, error) {
	return _SwapContractV1.Contract.LimitBTCForSPFlow2(&_SwapContractV1.CallOpts)
}

// LimitBTCForSPFlow2 is a free data retrieval call binding the contract method 0xa827278b.
//
// Solidity: function limitBTCForSPFlow2() view returns(uint256)
func (_SwapContractV1 *SwapContractV1CallerSession) LimitBTCForSPFlow2() (*big.Int, error) {
	return _SwapContractV1.Contract.LimitBTCForSPFlow2(&_SwapContractV1.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_SwapContractV1 *SwapContractV1Caller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "lpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_SwapContractV1 *SwapContractV1Session) LpToken() (common.Address, error) {
	return _SwapContractV1.Contract.LpToken(&_SwapContractV1.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_SwapContractV1 *SwapContractV1CallerSession) LpToken() (common.Address, error) {
	return _SwapContractV1.Contract.LpToken(&_SwapContractV1.CallOpts)
}

// OldestActiveIndex is a free data retrieval call binding the contract method 0x1bb06930.
//
// Solidity: function oldestActiveIndex() view returns(uint256)
func (_SwapContractV1 *SwapContractV1Caller) OldestActiveIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "oldestActiveIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OldestActiveIndex is a free data retrieval call binding the contract method 0x1bb06930.
//
// Solidity: function oldestActiveIndex() view returns(uint256)
func (_SwapContractV1 *SwapContractV1Session) OldestActiveIndex() (*big.Int, error) {
	return _SwapContractV1.Contract.OldestActiveIndex(&_SwapContractV1.CallOpts)
}

// OldestActiveIndex is a free data retrieval call binding the contract method 0x1bb06930.
//
// Solidity: function oldestActiveIndex() view returns(uint256)
func (_SwapContractV1 *SwapContractV1CallerSession) OldestActiveIndex() (*big.Int, error) {
	return _SwapContractV1.Contract.OldestActiveIndex(&_SwapContractV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapContractV1 *SwapContractV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapContractV1 *SwapContractV1Session) Owner() (common.Address, error) {
	return _SwapContractV1.Contract.Owner(&_SwapContractV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapContractV1 *SwapContractV1CallerSession) Owner() (common.Address, error) {
	return _SwapContractV1.Contract.Owner(&_SwapContractV1.CallOpts)
}

// ParaswapAddress is a free data retrieval call binding the contract method 0xf10d7c35.
//
// Solidity: function paraswapAddress() view returns(address)
func (_SwapContractV1 *SwapContractV1Caller) ParaswapAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "paraswapAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParaswapAddress is a free data retrieval call binding the contract method 0xf10d7c35.
//
// Solidity: function paraswapAddress() view returns(address)
func (_SwapContractV1 *SwapContractV1Session) ParaswapAddress() (common.Address, error) {
	return _SwapContractV1.Contract.ParaswapAddress(&_SwapContractV1.CallOpts)
}

// ParaswapAddress is a free data retrieval call binding the contract method 0xf10d7c35.
//
// Solidity: function paraswapAddress() view returns(address)
func (_SwapContractV1 *SwapContractV1CallerSession) ParaswapAddress() (common.Address, error) {
	return _SwapContractV1.Contract.ParaswapAddress(&_SwapContractV1.CallOpts)
}

// SbBTCPool is a free data retrieval call binding the contract method 0x0085aea1.
//
// Solidity: function sbBTCPool() view returns(address)
func (_SwapContractV1 *SwapContractV1Caller) SbBTCPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "sbBTCPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SbBTCPool is a free data retrieval call binding the contract method 0x0085aea1.
//
// Solidity: function sbBTCPool() view returns(address)
func (_SwapContractV1 *SwapContractV1Session) SbBTCPool() (common.Address, error) {
	return _SwapContractV1.Contract.SbBTCPool(&_SwapContractV1.CallOpts)
}

// SbBTCPool is a free data retrieval call binding the contract method 0x0085aea1.
//
// Solidity: function sbBTCPool() view returns(address)
func (_SwapContractV1 *SwapContractV1CallerSession) SbBTCPool() (common.Address, error) {
	return _SwapContractV1.Contract.SbBTCPool(&_SwapContractV1.CallOpts)
}

// SpGetPendingSwaps is a free data retrieval call binding the contract method 0x20199e99.
//
// Solidity: function spGetPendingSwaps() view returns((bytes32,string,address,uint256,uint256)[] data)
func (_SwapContractV1 *SwapContractV1Caller) SpGetPendingSwaps(opts *bind.CallOpts) ([]SwapContractV1spPendingTx, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "spGetPendingSwaps")

	if err != nil {
		return *new([]SwapContractV1spPendingTx), err
	}

	out0 := *abi.ConvertType(out[0], new([]SwapContractV1spPendingTx)).(*[]SwapContractV1spPendingTx)

	return out0, err

}

// SpGetPendingSwaps is a free data retrieval call binding the contract method 0x20199e99.
//
// Solidity: function spGetPendingSwaps() view returns((bytes32,string,address,uint256,uint256)[] data)
func (_SwapContractV1 *SwapContractV1Session) SpGetPendingSwaps() ([]SwapContractV1spPendingTx, error) {
	return _SwapContractV1.Contract.SpGetPendingSwaps(&_SwapContractV1.CallOpts)
}

// SpGetPendingSwaps is a free data retrieval call binding the contract method 0x20199e99.
//
// Solidity: function spGetPendingSwaps() view returns((bytes32,string,address,uint256,uint256)[] data)
func (_SwapContractV1 *SwapContractV1CallerSession) SpGetPendingSwaps() ([]SwapContractV1spPendingTx, error) {
	return _SwapContractV1.Contract.SpGetPendingSwaps(&_SwapContractV1.CallOpts)
}

// SpPendingTXs is a free data retrieval call binding the contract method 0x1ab7e937.
//
// Solidity: function spPendingTXs(uint256 ) view returns(bytes32 SwapID, string DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1Caller) SpPendingTXs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SwapID     [32]byte
	DestAddr   string
	RefundAddr common.Address
	AmountWBTC *big.Int
	Timestamp  *big.Int
}, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "spPendingTXs", arg0)

	outstruct := new(struct {
		SwapID     [32]byte
		DestAddr   string
		RefundAddr common.Address
		AmountWBTC *big.Int
		Timestamp  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SwapID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DestAddr = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.RefundAddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.AmountWBTC = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SpPendingTXs is a free data retrieval call binding the contract method 0x1ab7e937.
//
// Solidity: function spPendingTXs(uint256 ) view returns(bytes32 SwapID, string DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1Session) SpPendingTXs(arg0 *big.Int) (struct {
	SwapID     [32]byte
	DestAddr   string
	RefundAddr common.Address
	AmountWBTC *big.Int
	Timestamp  *big.Int
}, error) {
	return _SwapContractV1.Contract.SpPendingTXs(&_SwapContractV1.CallOpts, arg0)
}

// SpPendingTXs is a free data retrieval call binding the contract method 0x1ab7e937.
//
// Solidity: function spPendingTXs(uint256 ) view returns(bytes32 SwapID, string DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1CallerSession) SpPendingTXs(arg0 *big.Int) (struct {
	SwapID     [32]byte
	DestAddr   string
	RefundAddr common.Address
	AmountWBTC *big.Int
	Timestamp  *big.Int
}, error) {
	return _SwapContractV1.Contract.SpPendingTXs(&_SwapContractV1.CallOpts, arg0)
}

// Sw is a free data retrieval call binding the contract method 0x00e5cee4.
//
// Solidity: function sw() view returns(address)
func (_SwapContractV1 *SwapContractV1Caller) Sw(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "sw")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sw is a free data retrieval call binding the contract method 0x00e5cee4.
//
// Solidity: function sw() view returns(address)
func (_SwapContractV1 *SwapContractV1Session) Sw() (common.Address, error) {
	return _SwapContractV1.Contract.Sw(&_SwapContractV1.CallOpts)
}

// Sw is a free data retrieval call binding the contract method 0x00e5cee4.
//
// Solidity: function sw() view returns(address)
func (_SwapContractV1 *SwapContractV1CallerSession) Sw() (common.Address, error) {
	return _SwapContractV1.Contract.Sw(&_SwapContractV1.CallOpts)
}

// SwapCount is a free data retrieval call binding the contract method 0x2eff0d9e.
//
// Solidity: function swapCount() view returns(uint256)
func (_SwapContractV1 *SwapContractV1Caller) SwapCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "swapCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapCount is a free data retrieval call binding the contract method 0x2eff0d9e.
//
// Solidity: function swapCount() view returns(uint256)
func (_SwapContractV1 *SwapContractV1Session) SwapCount() (*big.Int, error) {
	return _SwapContractV1.Contract.SwapCount(&_SwapContractV1.CallOpts)
}

// SwapCount is a free data retrieval call binding the contract method 0x2eff0d9e.
//
// Solidity: function swapCount() view returns(uint256)
func (_SwapContractV1 *SwapContractV1CallerSession) SwapCount() (*big.Int, error) {
	return _SwapContractV1.Contract.SwapCount(&_SwapContractV1.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens(address , address ) view returns(uint256)
func (_SwapContractV1 *SwapContractV1Caller) Tokens(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "tokens", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens(address , address ) view returns(uint256)
func (_SwapContractV1 *SwapContractV1Session) Tokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SwapContractV1.Contract.Tokens(&_SwapContractV1.CallOpts, arg0, arg1)
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens(address , address ) view returns(uint256)
func (_SwapContractV1 *SwapContractV1CallerSession) Tokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SwapContractV1.Contract.Tokens(&_SwapContractV1.CallOpts, arg0, arg1)
}

// TssThreshold is a free data retrieval call binding the contract method 0x12d1441e.
//
// Solidity: function tssThreshold() view returns(uint8)
func (_SwapContractV1 *SwapContractV1Caller) TssThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "tssThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TssThreshold is a free data retrieval call binding the contract method 0x12d1441e.
//
// Solidity: function tssThreshold() view returns(uint8)
func (_SwapContractV1 *SwapContractV1Session) TssThreshold() (uint8, error) {
	return _SwapContractV1.Contract.TssThreshold(&_SwapContractV1.CallOpts)
}

// TssThreshold is a free data retrieval call binding the contract method 0x12d1441e.
//
// Solidity: function tssThreshold() view returns(uint8)
func (_SwapContractV1 *SwapContractV1CallerSession) TssThreshold() (uint8, error) {
	return _SwapContractV1.Contract.TssThreshold(&_SwapContractV1.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_SwapContractV1 *SwapContractV1Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "wETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_SwapContractV1 *SwapContractV1Session) WETH() (common.Address, error) {
	return _SwapContractV1.Contract.WETH(&_SwapContractV1.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_SwapContractV1 *SwapContractV1CallerSession) WETH() (common.Address, error) {
	return _SwapContractV1.Contract.WETH(&_SwapContractV1.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_SwapContractV1 *SwapContractV1Caller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SwapContractV1.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_SwapContractV1 *SwapContractV1Session) Whitelist(arg0 common.Address) (bool, error) {
	return _SwapContractV1.Contract.Whitelist(&_SwapContractV1.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_SwapContractV1 *SwapContractV1CallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _SwapContractV1.Contract.Whitelist(&_SwapContractV1.CallOpts, arg0)
}

// Churn is a paid mutator transaction binding the contract method 0x6845a025.
//
// Solidity: function churn(address _newOwner, address[] _nodes, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold) returns(bool)
func (_SwapContractV1 *SwapContractV1Transactor) Churn(opts *bind.TransactOpts, _newOwner common.Address, _nodes []common.Address, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "churn", _newOwner, _nodes, _isRemoved, _churnedInCount, _tssThreshold)
}

// Churn is a paid mutator transaction binding the contract method 0x6845a025.
//
// Solidity: function churn(address _newOwner, address[] _nodes, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold) returns(bool)
func (_SwapContractV1 *SwapContractV1Session) Churn(_newOwner common.Address, _nodes []common.Address, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8) (*types.Transaction, error) {
	return _SwapContractV1.Contract.Churn(&_SwapContractV1.TransactOpts, _newOwner, _nodes, _isRemoved, _churnedInCount, _tssThreshold)
}

// Churn is a paid mutator transaction binding the contract method 0x6845a025.
//
// Solidity: function churn(address _newOwner, address[] _nodes, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold) returns(bool)
func (_SwapContractV1 *SwapContractV1TransactorSession) Churn(_newOwner common.Address, _nodes []common.Address, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8) (*types.Transaction, error) {
	return _SwapContractV1.Contract.Churn(&_SwapContractV1.TransactOpts, _newOwner, _nodes, _isRemoved, _churnedInCount, _tssThreshold)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0xfcf965ca.
//
// Solidity: function collectSwapFeesForBTC(uint256 _incomingAmount, uint256 _rewardsAmount, address[] _spenders, uint256[] _swapAmounts, bool _isUpdatelimitBTCForSPFlow2) returns(bool)
func (_SwapContractV1 *SwapContractV1Transactor) CollectSwapFeesForBTC(opts *bind.TransactOpts, _incomingAmount *big.Int, _rewardsAmount *big.Int, _spenders []common.Address, _swapAmounts []*big.Int, _isUpdatelimitBTCForSPFlow2 bool) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "collectSwapFeesForBTC", _incomingAmount, _rewardsAmount, _spenders, _swapAmounts, _isUpdatelimitBTCForSPFlow2)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0xfcf965ca.
//
// Solidity: function collectSwapFeesForBTC(uint256 _incomingAmount, uint256 _rewardsAmount, address[] _spenders, uint256[] _swapAmounts, bool _isUpdatelimitBTCForSPFlow2) returns(bool)
func (_SwapContractV1 *SwapContractV1Session) CollectSwapFeesForBTC(_incomingAmount *big.Int, _rewardsAmount *big.Int, _spenders []common.Address, _swapAmounts []*big.Int, _isUpdatelimitBTCForSPFlow2 bool) (*types.Transaction, error) {
	return _SwapContractV1.Contract.CollectSwapFeesForBTC(&_SwapContractV1.TransactOpts, _incomingAmount, _rewardsAmount, _spenders, _swapAmounts, _isUpdatelimitBTCForSPFlow2)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0xfcf965ca.
//
// Solidity: function collectSwapFeesForBTC(uint256 _incomingAmount, uint256 _rewardsAmount, address[] _spenders, uint256[] _swapAmounts, bool _isUpdatelimitBTCForSPFlow2) returns(bool)
func (_SwapContractV1 *SwapContractV1TransactorSession) CollectSwapFeesForBTC(_incomingAmount *big.Int, _rewardsAmount *big.Int, _spenders []common.Address, _swapAmounts []*big.Int, _isUpdatelimitBTCForSPFlow2 bool) (*types.Transaction, error) {
	return _SwapContractV1.Contract.CollectSwapFeesForBTC(&_SwapContractV1.TransactOpts, _incomingAmount, _rewardsAmount, _spenders, _swapAmounts, _isUpdatelimitBTCForSPFlow2)
}

// Initialize is a paid mutator transaction binding the contract method 0x1460e390.
//
// Solidity: function initialize(address _lpToken, address _btct, address _wETH, address _sbBTCPool, address _params, address _swapRewards, uint256 _existingBTCFloat) returns()
func (_SwapContractV1 *SwapContractV1Transactor) Initialize(opts *bind.TransactOpts, _lpToken common.Address, _btct common.Address, _wETH common.Address, _sbBTCPool common.Address, _params common.Address, _swapRewards common.Address, _existingBTCFloat *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "initialize", _lpToken, _btct, _wETH, _sbBTCPool, _params, _swapRewards, _existingBTCFloat)
}

// Initialize is a paid mutator transaction binding the contract method 0x1460e390.
//
// Solidity: function initialize(address _lpToken, address _btct, address _wETH, address _sbBTCPool, address _params, address _swapRewards, uint256 _existingBTCFloat) returns()
func (_SwapContractV1 *SwapContractV1Session) Initialize(_lpToken common.Address, _btct common.Address, _wETH common.Address, _sbBTCPool common.Address, _params common.Address, _swapRewards common.Address, _existingBTCFloat *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.Contract.Initialize(&_SwapContractV1.TransactOpts, _lpToken, _btct, _wETH, _sbBTCPool, _params, _swapRewards, _existingBTCFloat)
}

// Initialize is a paid mutator transaction binding the contract method 0x1460e390.
//
// Solidity: function initialize(address _lpToken, address _btct, address _wETH, address _sbBTCPool, address _params, address _swapRewards, uint256 _existingBTCFloat) returns()
func (_SwapContractV1 *SwapContractV1TransactorSession) Initialize(_lpToken common.Address, _btct common.Address, _wETH common.Address, _sbBTCPool common.Address, _params common.Address, _swapRewards common.Address, _existingBTCFloat *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.Contract.Initialize(&_SwapContractV1.TransactOpts, _lpToken, _btct, _wETH, _sbBTCPool, _params, _swapRewards, _existingBTCFloat)
}

// MultiRecordSkyPoolsTX is a paid mutator transaction binding the contract method 0xbf62dcc8.
//
// Solidity: function multiRecordSkyPoolsTX(bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _usedTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1Transactor) MultiRecordSkyPoolsTX(opts *bind.TransactOpts, _addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _usedTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "multiRecordSkyPoolsTX", _addressesAndAmounts, _totalSwapped, _rewardsAmount, _usedTxIds)
}

// MultiRecordSkyPoolsTX is a paid mutator transaction binding the contract method 0xbf62dcc8.
//
// Solidity: function multiRecordSkyPoolsTX(bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _usedTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1Session) MultiRecordSkyPoolsTX(_addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _usedTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.MultiRecordSkyPoolsTX(&_SwapContractV1.TransactOpts, _addressesAndAmounts, _totalSwapped, _rewardsAmount, _usedTxIds)
}

// MultiRecordSkyPoolsTX is a paid mutator transaction binding the contract method 0xbf62dcc8.
//
// Solidity: function multiRecordSkyPoolsTX(bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _usedTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1TransactorSession) MultiRecordSkyPoolsTX(_addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _usedTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.MultiRecordSkyPoolsTX(&_SwapContractV1.TransactOpts, _addressesAndAmounts, _totalSwapped, _rewardsAmount, _usedTxIds)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad289e76.
//
// Solidity: function multiTransferERC20TightlyPacked(address _destToken, bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1Transactor) MultiTransferERC20TightlyPacked(opts *bind.TransactOpts, _destToken common.Address, _addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "multiTransferERC20TightlyPacked", _destToken, _addressesAndAmounts, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad289e76.
//
// Solidity: function multiTransferERC20TightlyPacked(address _destToken, bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1Session) MultiTransferERC20TightlyPacked(_destToken common.Address, _addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.MultiTransferERC20TightlyPacked(&_SwapContractV1.TransactOpts, _destToken, _addressesAndAmounts, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// MultiTransferERC20TightlyPacked is a paid mutator transaction binding the contract method 0xad289e76.
//
// Solidity: function multiTransferERC20TightlyPacked(address _destToken, bytes32[] _addressesAndAmounts, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1TransactorSession) MultiTransferERC20TightlyPacked(_destToken common.Address, _addressesAndAmounts [][32]byte, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.MultiTransferERC20TightlyPacked(&_SwapContractV1.TransactOpts, _destToken, _addressesAndAmounts, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0xcf10b16b.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bytes32 _txid) returns(bool)
func (_SwapContractV1 *SwapContractV1Transactor) RecordIncomingFloat(opts *bind.TransactOpts, _token common.Address, _addressesAndAmountOfFloat [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "recordIncomingFloat", _token, _addressesAndAmountOfFloat, _txid)
}

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0xcf10b16b.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bytes32 _txid) returns(bool)
func (_SwapContractV1 *SwapContractV1Session) RecordIncomingFloat(_token common.Address, _addressesAndAmountOfFloat [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RecordIncomingFloat(&_SwapContractV1.TransactOpts, _token, _addressesAndAmountOfFloat, _txid)
}

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0xcf10b16b.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bytes32 _txid) returns(bool)
func (_SwapContractV1 *SwapContractV1TransactorSession) RecordIncomingFloat(_token common.Address, _addressesAndAmountOfFloat [32]byte, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RecordIncomingFloat(&_SwapContractV1.TransactOpts, _token, _addressesAndAmountOfFloat, _txid)
}

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0x2586c562.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContractV1 *SwapContractV1Transactor) RecordOutcomingFloat(opts *bind.TransactOpts, _token common.Address, _addressesAndAmountOfLPtoken [32]byte, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "recordOutcomingFloat", _token, _addressesAndAmountOfLPtoken, _minerFee, _txid)
}

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0x2586c562.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContractV1 *SwapContractV1Session) RecordOutcomingFloat(_token common.Address, _addressesAndAmountOfLPtoken [32]byte, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RecordOutcomingFloat(&_SwapContractV1.TransactOpts, _token, _addressesAndAmountOfLPtoken, _minerFee, _txid)
}

// RecordOutcomingFloat is a paid mutator transaction binding the contract method 0x2586c562.
//
// Solidity: function recordOutcomingFloat(address _token, bytes32 _addressesAndAmountOfLPtoken, uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContractV1 *SwapContractV1TransactorSession) RecordOutcomingFloat(_token common.Address, _addressesAndAmountOfLPtoken [32]byte, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RecordOutcomingFloat(&_SwapContractV1.TransactOpts, _token, _addressesAndAmountOfLPtoken, _minerFee, _txid)
}

// RecordSkyPoolsTX is a paid mutator transaction binding the contract method 0xb345a358.
//
// Solidity: function recordSkyPoolsTX(address _to, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _usedTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1Transactor) RecordSkyPoolsTX(opts *bind.TransactOpts, _to common.Address, _totalSwapped *big.Int, _rewardsAmount *big.Int, _usedTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "recordSkyPoolsTX", _to, _totalSwapped, _rewardsAmount, _usedTxIds)
}

// RecordSkyPoolsTX is a paid mutator transaction binding the contract method 0xb345a358.
//
// Solidity: function recordSkyPoolsTX(address _to, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _usedTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1Session) RecordSkyPoolsTX(_to common.Address, _totalSwapped *big.Int, _rewardsAmount *big.Int, _usedTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RecordSkyPoolsTX(&_SwapContractV1.TransactOpts, _to, _totalSwapped, _rewardsAmount, _usedTxIds)
}

// RecordSkyPoolsTX is a paid mutator transaction binding the contract method 0xb345a358.
//
// Solidity: function recordSkyPoolsTX(address _to, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _usedTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1TransactorSession) RecordSkyPoolsTX(_to common.Address, _totalSwapped *big.Int, _rewardsAmount *big.Int, _usedTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RecordSkyPoolsTX(&_SwapContractV1.TransactOpts, _to, _totalSwapped, _rewardsAmount, _usedTxIds)
}

// RecordUTXOSweepMinerFee is a paid mutator transaction binding the contract method 0xc810a539.
//
// Solidity: function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContractV1 *SwapContractV1Transactor) RecordUTXOSweepMinerFee(opts *bind.TransactOpts, _minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "recordUTXOSweepMinerFee", _minerFee, _txid)
}

// RecordUTXOSweepMinerFee is a paid mutator transaction binding the contract method 0xc810a539.
//
// Solidity: function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContractV1 *SwapContractV1Session) RecordUTXOSweepMinerFee(_minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RecordUTXOSweepMinerFee(&_SwapContractV1.TransactOpts, _minerFee, _txid)
}

// RecordUTXOSweepMinerFee is a paid mutator transaction binding the contract method 0xc810a539.
//
// Solidity: function recordUTXOSweepMinerFee(uint256 _minerFee, bytes32 _txid) returns(bool)
func (_SwapContractV1 *SwapContractV1TransactorSession) RecordUTXOSweepMinerFee(_minerFee *big.Int, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RecordUTXOSweepMinerFee(&_SwapContractV1.TransactOpts, _minerFee, _txid)
}

// RedeemERC20Token is a paid mutator transaction binding the contract method 0x4db75485.
//
// Solidity: function redeemERC20Token(address _token, uint256 _amount) returns()
func (_SwapContractV1 *SwapContractV1Transactor) RedeemERC20Token(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "redeemERC20Token", _token, _amount)
}

// RedeemERC20Token is a paid mutator transaction binding the contract method 0x4db75485.
//
// Solidity: function redeemERC20Token(address _token, uint256 _amount) returns()
func (_SwapContractV1 *SwapContractV1Session) RedeemERC20Token(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RedeemERC20Token(&_SwapContractV1.TransactOpts, _token, _amount)
}

// RedeemERC20Token is a paid mutator transaction binding the contract method 0x4db75485.
//
// Solidity: function redeemERC20Token(address _token, uint256 _amount) returns()
func (_SwapContractV1 *SwapContractV1TransactorSession) RedeemERC20Token(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RedeemERC20Token(&_SwapContractV1.TransactOpts, _token, _amount)
}

// RedeemEther is a paid mutator transaction binding the contract method 0x8535490f.
//
// Solidity: function redeemEther(uint256 _amount) returns()
func (_SwapContractV1 *SwapContractV1Transactor) RedeemEther(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "redeemEther", _amount)
}

// RedeemEther is a paid mutator transaction binding the contract method 0x8535490f.
//
// Solidity: function redeemEther(uint256 _amount) returns()
func (_SwapContractV1 *SwapContractV1Session) RedeemEther(_amount *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RedeemEther(&_SwapContractV1.TransactOpts, _amount)
}

// RedeemEther is a paid mutator transaction binding the contract method 0x8535490f.
//
// Solidity: function redeemEther(uint256 _amount) returns()
func (_SwapContractV1 *SwapContractV1TransactorSession) RedeemEther(_amount *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.Contract.RedeemEther(&_SwapContractV1.TransactOpts, _amount)
}

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x0d63aca7.
//
// Solidity: function singleTransferERC20(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1Transactor) SingleTransferERC20(opts *bind.TransactOpts, _destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "singleTransferERC20", _destToken, _to, _amount, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x0d63aca7.
//
// Solidity: function singleTransferERC20(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1Session) SingleTransferERC20(_destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SingleTransferERC20(&_SwapContractV1.TransactOpts, _destToken, _to, _amount, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// SingleTransferERC20 is a paid mutator transaction binding the contract method 0x0d63aca7.
//
// Solidity: function singleTransferERC20(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, uint256 _rewardsAmount, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContractV1 *SwapContractV1TransactorSession) SingleTransferERC20(_destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _rewardsAmount *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SingleTransferERC20(&_SwapContractV1.TransactOpts, _destToken, _to, _amount, _totalSwapped, _rewardsAmount, _redeemedFloatTxIds)
}

// SpCleanUpOldTXs is a paid mutator transaction binding the contract method 0xb9ae0adc.
//
// Solidity: function spCleanUpOldTXs() returns()
func (_SwapContractV1 *SwapContractV1Transactor) SpCleanUpOldTXs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "spCleanUpOldTXs")
}

// SpCleanUpOldTXs is a paid mutator transaction binding the contract method 0xb9ae0adc.
//
// Solidity: function spCleanUpOldTXs() returns()
func (_SwapContractV1 *SwapContractV1Session) SpCleanUpOldTXs() (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpCleanUpOldTXs(&_SwapContractV1.TransactOpts)
}

// SpCleanUpOldTXs is a paid mutator transaction binding the contract method 0xb9ae0adc.
//
// Solidity: function spCleanUpOldTXs() returns()
func (_SwapContractV1 *SwapContractV1TransactorSession) SpCleanUpOldTXs() (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpCleanUpOldTXs(&_SwapContractV1.TransactOpts)
}

// SpDeposit is a paid mutator transaction binding the contract method 0xca0707e7.
//
// Solidity: function spDeposit(address _token, uint256 _amount) payable returns()
func (_SwapContractV1 *SwapContractV1Transactor) SpDeposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "spDeposit", _token, _amount)
}

// SpDeposit is a paid mutator transaction binding the contract method 0xca0707e7.
//
// Solidity: function spDeposit(address _token, uint256 _amount) payable returns()
func (_SwapContractV1 *SwapContractV1Session) SpDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpDeposit(&_SwapContractV1.TransactOpts, _token, _amount)
}

// SpDeposit is a paid mutator transaction binding the contract method 0xca0707e7.
//
// Solidity: function spDeposit(address _token, uint256 _amount) payable returns()
func (_SwapContractV1 *SwapContractV1TransactorSession) SpDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpDeposit(&_SwapContractV1.TransactOpts, _token, _amount)
}

// SpFlow1SimpleSwap is a paid mutator transaction binding the contract method 0x511834c3.
//
// Solidity: function spFlow1SimpleSwap((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) returns()
func (_SwapContractV1 *SwapContractV1Transactor) SpFlow1SimpleSwap(opts *bind.TransactOpts, _data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "spFlow1SimpleSwap", _data)
}

// SpFlow1SimpleSwap is a paid mutator transaction binding the contract method 0x511834c3.
//
// Solidity: function spFlow1SimpleSwap((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) returns()
func (_SwapContractV1 *SwapContractV1Session) SpFlow1SimpleSwap(_data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpFlow1SimpleSwap(&_SwapContractV1.TransactOpts, _data)
}

// SpFlow1SimpleSwap is a paid mutator transaction binding the contract method 0x511834c3.
//
// Solidity: function spFlow1SimpleSwap((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) returns()
func (_SwapContractV1 *SwapContractV1TransactorSession) SpFlow1SimpleSwap(_data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpFlow1SimpleSwap(&_SwapContractV1.TransactOpts, _data)
}

// SpFlow1Uniswap is a paid mutator transaction binding the contract method 0x130b7a63.
//
// Solidity: function spFlow1Uniswap(bool _fork, address _factory, bytes32 _initCode, uint256 _amountIn, uint256 _amountOutMin, address[] _path) returns(uint256 receivedAmount)
func (_SwapContractV1 *SwapContractV1Transactor) SpFlow1Uniswap(opts *bind.TransactOpts, _fork bool, _factory common.Address, _initCode [32]byte, _amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "spFlow1Uniswap", _fork, _factory, _initCode, _amountIn, _amountOutMin, _path)
}

// SpFlow1Uniswap is a paid mutator transaction binding the contract method 0x130b7a63.
//
// Solidity: function spFlow1Uniswap(bool _fork, address _factory, bytes32 _initCode, uint256 _amountIn, uint256 _amountOutMin, address[] _path) returns(uint256 receivedAmount)
func (_SwapContractV1 *SwapContractV1Session) SpFlow1Uniswap(_fork bool, _factory common.Address, _initCode [32]byte, _amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpFlow1Uniswap(&_SwapContractV1.TransactOpts, _fork, _factory, _initCode, _amountIn, _amountOutMin, _path)
}

// SpFlow1Uniswap is a paid mutator transaction binding the contract method 0x130b7a63.
//
// Solidity: function spFlow1Uniswap(bool _fork, address _factory, bytes32 _initCode, uint256 _amountIn, uint256 _amountOutMin, address[] _path) returns(uint256 receivedAmount)
func (_SwapContractV1 *SwapContractV1TransactorSession) SpFlow1Uniswap(_fork bool, _factory common.Address, _initCode [32]byte, _amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpFlow1Uniswap(&_SwapContractV1.TransactOpts, _fork, _factory, _initCode, _amountIn, _amountOutMin, _path)
}

// SpFlow2SimpleSwap is a paid mutator transaction binding the contract method 0xb4014fe8.
//
// Solidity: function spFlow2SimpleSwap(string _destinationAddressForBTC, (address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) returns(uint256 receivedAmount)
func (_SwapContractV1 *SwapContractV1Transactor) SpFlow2SimpleSwap(opts *bind.TransactOpts, _destinationAddressForBTC string, _data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "spFlow2SimpleSwap", _destinationAddressForBTC, _data)
}

// SpFlow2SimpleSwap is a paid mutator transaction binding the contract method 0xb4014fe8.
//
// Solidity: function spFlow2SimpleSwap(string _destinationAddressForBTC, (address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) returns(uint256 receivedAmount)
func (_SwapContractV1 *SwapContractV1Session) SpFlow2SimpleSwap(_destinationAddressForBTC string, _data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpFlow2SimpleSwap(&_SwapContractV1.TransactOpts, _destinationAddressForBTC, _data)
}

// SpFlow2SimpleSwap is a paid mutator transaction binding the contract method 0xb4014fe8.
//
// Solidity: function spFlow2SimpleSwap(string _destinationAddressForBTC, (address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) returns(uint256 receivedAmount)
func (_SwapContractV1 *SwapContractV1TransactorSession) SpFlow2SimpleSwap(_destinationAddressForBTC string, _data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpFlow2SimpleSwap(&_SwapContractV1.TransactOpts, _destinationAddressForBTC, _data)
}

// SpFlow2Uniswap is a paid mutator transaction binding the contract method 0xd38c7bb1.
//
// Solidity: function spFlow2Uniswap(string _destinationAddressForBTC, bool _fork, address _factory, bytes32 _initCode, uint256 _amountIn, uint256 _amountOutMin, address[] _path) returns(uint256 receivedAmount)
func (_SwapContractV1 *SwapContractV1Transactor) SpFlow2Uniswap(opts *bind.TransactOpts, _destinationAddressForBTC string, _fork bool, _factory common.Address, _initCode [32]byte, _amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "spFlow2Uniswap", _destinationAddressForBTC, _fork, _factory, _initCode, _amountIn, _amountOutMin, _path)
}

// SpFlow2Uniswap is a paid mutator transaction binding the contract method 0xd38c7bb1.
//
// Solidity: function spFlow2Uniswap(string _destinationAddressForBTC, bool _fork, address _factory, bytes32 _initCode, uint256 _amountIn, uint256 _amountOutMin, address[] _path) returns(uint256 receivedAmount)
func (_SwapContractV1 *SwapContractV1Session) SpFlow2Uniswap(_destinationAddressForBTC string, _fork bool, _factory common.Address, _initCode [32]byte, _amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpFlow2Uniswap(&_SwapContractV1.TransactOpts, _destinationAddressForBTC, _fork, _factory, _initCode, _amountIn, _amountOutMin, _path)
}

// SpFlow2Uniswap is a paid mutator transaction binding the contract method 0xd38c7bb1.
//
// Solidity: function spFlow2Uniswap(string _destinationAddressForBTC, bool _fork, address _factory, bytes32 _initCode, uint256 _amountIn, uint256 _amountOutMin, address[] _path) returns(uint256 receivedAmount)
func (_SwapContractV1 *SwapContractV1TransactorSession) SpFlow2Uniswap(_destinationAddressForBTC string, _fork bool, _factory common.Address, _initCode [32]byte, _amountIn *big.Int, _amountOutMin *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _SwapContractV1.Contract.SpFlow2Uniswap(&_SwapContractV1.TransactOpts, _destinationAddressForBTC, _fork, _factory, _initCode, _amountIn, _amountOutMin, _path)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_SwapContractV1 *SwapContractV1Transactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SwapContractV1.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_SwapContractV1 *SwapContractV1Session) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SwapContractV1.Contract.TransferOwnership(&_SwapContractV1.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_SwapContractV1 *SwapContractV1TransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SwapContractV1.Contract.TransferOwnership(&_SwapContractV1.TransactOpts, _newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapContractV1 *SwapContractV1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapContractV1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapContractV1 *SwapContractV1Session) Receive() (*types.Transaction, error) {
	return _SwapContractV1.Contract.Receive(&_SwapContractV1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapContractV1 *SwapContractV1TransactorSession) Receive() (*types.Transaction, error) {
	return _SwapContractV1.Contract.Receive(&_SwapContractV1.TransactOpts)
}

// SwapContractV1BurnLPTokensForFloatIterator is returned from FilterBurnLPTokensForFloat and is used to iterate over the raw logs and unpacked data for BurnLPTokensForFloat events raised by the SwapContractV1 contract.
type SwapContractV1BurnLPTokensForFloatIterator struct {
	Event *SwapContractV1BurnLPTokensForFloat // Event containing the contract specifics and raw log

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
func (it *SwapContractV1BurnLPTokensForFloatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractV1BurnLPTokensForFloat)
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
		it.Event = new(SwapContractV1BurnLPTokensForFloat)
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
func (it *SwapContractV1BurnLPTokensForFloatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractV1BurnLPTokensForFloatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractV1BurnLPTokensForFloat represents a BurnLPTokensForFloat event raised by the SwapContractV1 contract.
type SwapContractV1BurnLPTokensForFloat struct {
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
func (_SwapContractV1 *SwapContractV1Filterer) FilterBurnLPTokensForFloat(opts *bind.FilterOpts) (*SwapContractV1BurnLPTokensForFloatIterator, error) {

	logs, sub, err := _SwapContractV1.contract.FilterLogs(opts, "BurnLPTokensForFloat")
	if err != nil {
		return nil, err
	}
	return &SwapContractV1BurnLPTokensForFloatIterator{contract: _SwapContractV1.contract, event: "BurnLPTokensForFloat", logs: logs, sub: sub}, nil
}

// WatchBurnLPTokensForFloat is a free log subscription operation binding the contract event 0xa9da34e5e4a956b744307ae7611795634e8365951073bf049d76b64b2ae7a058.
//
// Solidity: event BurnLPTokensForFloat(address token, uint256 amountOfLP, uint256 amountOfFloat, uint256 currentPriceLP, uint256 withdrawal, bytes32 txid)
func (_SwapContractV1 *SwapContractV1Filterer) WatchBurnLPTokensForFloat(opts *bind.WatchOpts, sink chan<- *SwapContractV1BurnLPTokensForFloat) (event.Subscription, error) {

	logs, sub, err := _SwapContractV1.contract.WatchLogs(opts, "BurnLPTokensForFloat")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractV1BurnLPTokensForFloat)
				if err := _SwapContractV1.contract.UnpackLog(event, "BurnLPTokensForFloat", log); err != nil {
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
func (_SwapContractV1 *SwapContractV1Filterer) ParseBurnLPTokensForFloat(log types.Log) (*SwapContractV1BurnLPTokensForFloat, error) {
	event := new(SwapContractV1BurnLPTokensForFloat)
	if err := _SwapContractV1.contract.UnpackLog(event, "BurnLPTokensForFloat", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractV1DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SwapContractV1 contract.
type SwapContractV1DepositIterator struct {
	Event *SwapContractV1Deposit // Event containing the contract specifics and raw log

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
func (it *SwapContractV1DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractV1Deposit)
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
		it.Event = new(SwapContractV1Deposit)
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
func (it *SwapContractV1DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractV1DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractV1Deposit represents a Deposit event raised by the SwapContractV1 contract.
type SwapContractV1Deposit struct {
	Token     common.Address
	User      common.Address
	Amount    *big.Int
	Balance   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address token, address user, uint256 amount, uint256 balance, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1Filterer) FilterDeposit(opts *bind.FilterOpts) (*SwapContractV1DepositIterator, error) {

	logs, sub, err := _SwapContractV1.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &SwapContractV1DepositIterator{contract: _SwapContractV1.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address token, address user, uint256 amount, uint256 balance, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SwapContractV1Deposit) (event.Subscription, error) {

	logs, sub, err := _SwapContractV1.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractV1Deposit)
				if err := _SwapContractV1.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address token, address user, uint256 amount, uint256 balance, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1Filterer) ParseDeposit(log types.Log) (*SwapContractV1Deposit, error) {
	event := new(SwapContractV1Deposit)
	if err := _SwapContractV1.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractV1DistributeNodeRewardsIterator is returned from FilterDistributeNodeRewards and is used to iterate over the raw logs and unpacked data for DistributeNodeRewards events raised by the SwapContractV1 contract.
type SwapContractV1DistributeNodeRewardsIterator struct {
	Event *SwapContractV1DistributeNodeRewards // Event containing the contract specifics and raw log

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
func (it *SwapContractV1DistributeNodeRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractV1DistributeNodeRewards)
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
		it.Event = new(SwapContractV1DistributeNodeRewards)
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
func (it *SwapContractV1DistributeNodeRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractV1DistributeNodeRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractV1DistributeNodeRewards represents a DistributeNodeRewards event raised by the SwapContractV1 contract.
type SwapContractV1DistributeNodeRewards struct {
	RewardLPTsForNodes *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDistributeNodeRewards is a free log retrieval operation binding the contract event 0x86dd6ae252e62228b47b3c1a12762cbf634952b2d29319b04f4f72fb153bcbdd.
//
// Solidity: event DistributeNodeRewards(uint256 rewardLPTsForNodes)
func (_SwapContractV1 *SwapContractV1Filterer) FilterDistributeNodeRewards(opts *bind.FilterOpts) (*SwapContractV1DistributeNodeRewardsIterator, error) {

	logs, sub, err := _SwapContractV1.contract.FilterLogs(opts, "DistributeNodeRewards")
	if err != nil {
		return nil, err
	}
	return &SwapContractV1DistributeNodeRewardsIterator{contract: _SwapContractV1.contract, event: "DistributeNodeRewards", logs: logs, sub: sub}, nil
}

// WatchDistributeNodeRewards is a free log subscription operation binding the contract event 0x86dd6ae252e62228b47b3c1a12762cbf634952b2d29319b04f4f72fb153bcbdd.
//
// Solidity: event DistributeNodeRewards(uint256 rewardLPTsForNodes)
func (_SwapContractV1 *SwapContractV1Filterer) WatchDistributeNodeRewards(opts *bind.WatchOpts, sink chan<- *SwapContractV1DistributeNodeRewards) (event.Subscription, error) {

	logs, sub, err := _SwapContractV1.contract.WatchLogs(opts, "DistributeNodeRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractV1DistributeNodeRewards)
				if err := _SwapContractV1.contract.UnpackLog(event, "DistributeNodeRewards", log); err != nil {
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

// ParseDistributeNodeRewards is a log parse operation binding the contract event 0x86dd6ae252e62228b47b3c1a12762cbf634952b2d29319b04f4f72fb153bcbdd.
//
// Solidity: event DistributeNodeRewards(uint256 rewardLPTsForNodes)
func (_SwapContractV1 *SwapContractV1Filterer) ParseDistributeNodeRewards(log types.Log) (*SwapContractV1DistributeNodeRewards, error) {
	event := new(SwapContractV1DistributeNodeRewards)
	if err := _SwapContractV1.contract.UnpackLog(event, "DistributeNodeRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractV1IssueLPTokensForFloatIterator is returned from FilterIssueLPTokensForFloat and is used to iterate over the raw logs and unpacked data for IssueLPTokensForFloat events raised by the SwapContractV1 contract.
type SwapContractV1IssueLPTokensForFloatIterator struct {
	Event *SwapContractV1IssueLPTokensForFloat // Event containing the contract specifics and raw log

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
func (it *SwapContractV1IssueLPTokensForFloatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractV1IssueLPTokensForFloat)
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
		it.Event = new(SwapContractV1IssueLPTokensForFloat)
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
func (it *SwapContractV1IssueLPTokensForFloatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractV1IssueLPTokensForFloatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractV1IssueLPTokensForFloat represents a IssueLPTokensForFloat event raised by the SwapContractV1 contract.
type SwapContractV1IssueLPTokensForFloat struct {
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
func (_SwapContractV1 *SwapContractV1Filterer) FilterIssueLPTokensForFloat(opts *bind.FilterOpts) (*SwapContractV1IssueLPTokensForFloatIterator, error) {

	logs, sub, err := _SwapContractV1.contract.FilterLogs(opts, "IssueLPTokensForFloat")
	if err != nil {
		return nil, err
	}
	return &SwapContractV1IssueLPTokensForFloatIterator{contract: _SwapContractV1.contract, event: "IssueLPTokensForFloat", logs: logs, sub: sub}, nil
}

// WatchIssueLPTokensForFloat is a free log subscription operation binding the contract event 0xdfaf4e23f7030adafec91a682e05bf9fb30c721f9cec3d8bcfae0459c5362db1.
//
// Solidity: event IssueLPTokensForFloat(address to, uint256 amountOfFloat, uint256 amountOfLP, uint256 currentPriceLP, uint256 depositFees, bytes32 txid)
func (_SwapContractV1 *SwapContractV1Filterer) WatchIssueLPTokensForFloat(opts *bind.WatchOpts, sink chan<- *SwapContractV1IssueLPTokensForFloat) (event.Subscription, error) {

	logs, sub, err := _SwapContractV1.contract.WatchLogs(opts, "IssueLPTokensForFloat")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractV1IssueLPTokensForFloat)
				if err := _SwapContractV1.contract.UnpackLog(event, "IssueLPTokensForFloat", log); err != nil {
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
func (_SwapContractV1 *SwapContractV1Filterer) ParseIssueLPTokensForFloat(log types.Log) (*SwapContractV1IssueLPTokensForFloat, error) {
	event := new(SwapContractV1IssueLPTokensForFloat)
	if err := _SwapContractV1.contract.UnpackLog(event, "IssueLPTokensForFloat", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractV1RewardsCollectionIterator is returned from FilterRewardsCollection and is used to iterate over the raw logs and unpacked data for RewardsCollection events raised by the SwapContractV1 contract.
type SwapContractV1RewardsCollectionIterator struct {
	Event *SwapContractV1RewardsCollection // Event containing the contract specifics and raw log

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
func (it *SwapContractV1RewardsCollectionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractV1RewardsCollection)
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
		it.Event = new(SwapContractV1RewardsCollection)
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
func (it *SwapContractV1RewardsCollectionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractV1RewardsCollectionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractV1RewardsCollection represents a RewardsCollection event raised by the SwapContractV1 contract.
type SwapContractV1RewardsCollection struct {
	FeesToken             common.Address
	Rewards               *big.Int
	AmountLPTokensForNode *big.Int
	CurrentPriceLP        *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRewardsCollection is a free log retrieval operation binding the contract event 0xe3debe835f6848edc082b32a1d729781d3bfcd7e14422d80bcaa6794d3816b2f.
//
// Solidity: event RewardsCollection(address feesToken, uint256 rewards, uint256 amountLPTokensForNode, uint256 currentPriceLP)
func (_SwapContractV1 *SwapContractV1Filterer) FilterRewardsCollection(opts *bind.FilterOpts) (*SwapContractV1RewardsCollectionIterator, error) {

	logs, sub, err := _SwapContractV1.contract.FilterLogs(opts, "RewardsCollection")
	if err != nil {
		return nil, err
	}
	return &SwapContractV1RewardsCollectionIterator{contract: _SwapContractV1.contract, event: "RewardsCollection", logs: logs, sub: sub}, nil
}

// WatchRewardsCollection is a free log subscription operation binding the contract event 0xe3debe835f6848edc082b32a1d729781d3bfcd7e14422d80bcaa6794d3816b2f.
//
// Solidity: event RewardsCollection(address feesToken, uint256 rewards, uint256 amountLPTokensForNode, uint256 currentPriceLP)
func (_SwapContractV1 *SwapContractV1Filterer) WatchRewardsCollection(opts *bind.WatchOpts, sink chan<- *SwapContractV1RewardsCollection) (event.Subscription, error) {

	logs, sub, err := _SwapContractV1.contract.WatchLogs(opts, "RewardsCollection")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractV1RewardsCollection)
				if err := _SwapContractV1.contract.UnpackLog(event, "RewardsCollection", log); err != nil {
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
// Solidity: event RewardsCollection(address feesToken, uint256 rewards, uint256 amountLPTokensForNode, uint256 currentPriceLP)
func (_SwapContractV1 *SwapContractV1Filterer) ParseRewardsCollection(log types.Log) (*SwapContractV1RewardsCollection, error) {
	event := new(SwapContractV1RewardsCollection)
	if err := _SwapContractV1.contract.UnpackLog(event, "RewardsCollection", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractV1SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the SwapContractV1 contract.
type SwapContractV1SwapIterator struct {
	Event *SwapContractV1Swap // Event containing the contract specifics and raw log

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
func (it *SwapContractV1SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractV1Swap)
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
		it.Event = new(SwapContractV1Swap)
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
func (it *SwapContractV1SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractV1SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractV1Swap represents a Swap event raised by the SwapContractV1 contract.
type SwapContractV1Swap struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address from, address to, uint256 amount)
func (_SwapContractV1 *SwapContractV1Filterer) FilterSwap(opts *bind.FilterOpts) (*SwapContractV1SwapIterator, error) {

	logs, sub, err := _SwapContractV1.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &SwapContractV1SwapIterator{contract: _SwapContractV1.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address from, address to, uint256 amount)
func (_SwapContractV1 *SwapContractV1Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *SwapContractV1Swap) (event.Subscription, error) {

	logs, sub, err := _SwapContractV1.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractV1Swap)
				if err := _SwapContractV1.contract.UnpackLog(event, "Swap", log); err != nil {
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
func (_SwapContractV1 *SwapContractV1Filterer) ParseSwap(log types.Log) (*SwapContractV1Swap, error) {
	event := new(SwapContractV1Swap)
	if err := _SwapContractV1.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractV1SwapTokensToBTCIterator is returned from FilterSwapTokensToBTC and is used to iterate over the raw logs and unpacked data for SwapTokensToBTC events raised by the SwapContractV1 contract.
type SwapContractV1SwapTokensToBTCIterator struct {
	Event *SwapContractV1SwapTokensToBTC // Event containing the contract specifics and raw log

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
func (it *SwapContractV1SwapTokensToBTCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractV1SwapTokensToBTC)
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
		it.Event = new(SwapContractV1SwapTokensToBTC)
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
func (it *SwapContractV1SwapTokensToBTCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractV1SwapTokensToBTCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractV1SwapTokensToBTC represents a SwapTokensToBTC event raised by the SwapContractV1 contract.
type SwapContractV1SwapTokensToBTC struct {
	SwapID     [32]byte
	DestAddr   string
	RefundAddr common.Address
	AmountWBTC *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwapTokensToBTC is a free log retrieval operation binding the contract event 0xaaa4c278970337923c885845b00985ecf5082317e2f9abf72a27b4d35255d672.
//
// Solidity: event SwapTokensToBTC(bytes32 SwapID, string DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1Filterer) FilterSwapTokensToBTC(opts *bind.FilterOpts) (*SwapContractV1SwapTokensToBTCIterator, error) {

	logs, sub, err := _SwapContractV1.contract.FilterLogs(opts, "SwapTokensToBTC")
	if err != nil {
		return nil, err
	}
	return &SwapContractV1SwapTokensToBTCIterator{contract: _SwapContractV1.contract, event: "SwapTokensToBTC", logs: logs, sub: sub}, nil
}

// WatchSwapTokensToBTC is a free log subscription operation binding the contract event 0xaaa4c278970337923c885845b00985ecf5082317e2f9abf72a27b4d35255d672.
//
// Solidity: event SwapTokensToBTC(bytes32 SwapID, string DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1Filterer) WatchSwapTokensToBTC(opts *bind.WatchOpts, sink chan<- *SwapContractV1SwapTokensToBTC) (event.Subscription, error) {

	logs, sub, err := _SwapContractV1.contract.WatchLogs(opts, "SwapTokensToBTC")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractV1SwapTokensToBTC)
				if err := _SwapContractV1.contract.UnpackLog(event, "SwapTokensToBTC", log); err != nil {
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

// ParseSwapTokensToBTC is a log parse operation binding the contract event 0xaaa4c278970337923c885845b00985ecf5082317e2f9abf72a27b4d35255d672.
//
// Solidity: event SwapTokensToBTC(bytes32 SwapID, string DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1Filterer) ParseSwapTokensToBTC(log types.Log) (*SwapContractV1SwapTokensToBTC, error) {
	event := new(SwapContractV1SwapTokensToBTC)
	if err := _SwapContractV1.contract.UnpackLog(event, "SwapTokensToBTC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractV1WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SwapContractV1 contract.
type SwapContractV1WithdrawIterator struct {
	Event *SwapContractV1Withdraw // Event containing the contract specifics and raw log

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
func (it *SwapContractV1WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractV1Withdraw)
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
		it.Event = new(SwapContractV1Withdraw)
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
func (it *SwapContractV1WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractV1WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractV1Withdraw represents a Withdraw event raised by the SwapContractV1 contract.
type SwapContractV1Withdraw struct {
	Token     common.Address
	User      common.Address
	Amount    *big.Int
	Balance   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address token, address user, uint256 amount, uint256 balance, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1Filterer) FilterWithdraw(opts *bind.FilterOpts) (*SwapContractV1WithdrawIterator, error) {

	logs, sub, err := _SwapContractV1.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &SwapContractV1WithdrawIterator{contract: _SwapContractV1.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address token, address user, uint256 amount, uint256 balance, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SwapContractV1Withdraw) (event.Subscription, error) {

	logs, sub, err := _SwapContractV1.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractV1Withdraw)
				if err := _SwapContractV1.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address token, address user, uint256 amount, uint256 balance, uint256 Timestamp)
func (_SwapContractV1 *SwapContractV1Filterer) ParseWithdraw(log types.Log) (*SwapContractV1Withdraw, error) {
	event := new(SwapContractV1Withdraw)
	if err := _SwapContractV1.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
