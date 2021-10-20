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

// SwapContractspPendingTx is an auto generated low-level Go binding around an user-defined struct.
type SwapContractspPendingTx struct {
	SwapID     [32]byte
	DestAddr   [32]byte
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

// SwapContractMetaData contains all meta data concerning the SwapContract contract.
var SwapContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_btct\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wETH\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_existingBTCFloat\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfFloat\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPriceLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"}],\"name\":\"BurnLPTokensForFloat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardLPTsForNodes\",\"type\":\"uint256\"}],\"name\":\"DistributeNodeRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfFloat\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPriceLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"}],\"name\":\"IssueLPTokensForFloat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feesToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLPTokensForNode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPriceLP\",\"type\":\"uint256\"}],\"name\":\"RewardsCollection\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ExpirationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Timestamp\",\"type\":\"uint256\"}],\"name\":\"SetExpirationTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"SwapID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"DestAddr\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"RefundAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"AmountWBTC\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Timestamp\",\"type\":\"uint256\"}],\"name\":\"SwapTokensToBTC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Timestamp\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"BTCT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"_rewardAddressAndAmounts\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"_isRemoved\",\"type\":\"bool[]\"},{\"internalType\":\"uint8\",\"name\":\"_churnedInCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_tssThreshold\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_nodeRewardsRatio\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_withdrawalFeeBPS\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_expirationTime\",\"type\":\"uint256\"}],\"name\":\"churn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"churnedInCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_incomingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"}],\"name\":\"collectSwapFeesForBTC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositFeesBPS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeNodeRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expirationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feesLPTokensForNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveNodes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentPriceLP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nowPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOfFloat\",\"type\":\"uint256\"}],\"name\":\"getDepositFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositFeeRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenB\",\"type\":\"address\"}],\"name\":\"getFloatReserve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"isTxUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockedLPTokensForNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"_addressesAndAmounts\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalSwapped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_redeemedFloatTxIds\",\"type\":\"bytes32[]\"}],\"name\":\"multiTransferERC20TightlyPacked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeRewardsRatio\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_addressesAndAmountOfFloat\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_zerofee\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordIncomingFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_addressesAndAmountOfLPtoken\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_minerFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordOutcomingFloat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSwapped\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_redeemedFloatTxIds\",\"type\":\"bytes32[]\"}],\"name\":\"recordSkyPoolsTX\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minerFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txid\",\"type\":\"bytes32\"}],\"name\":\"recordUTXOSweepMinerFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeemERC20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeemEther\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSwapped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_redeemedFloatTxIds\",\"type\":\"bytes32[]\"}],\"name\":\"singleTransferERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_loopCount\",\"type\":\"uint256\"}],\"name\":\"spCleanUpOldTXs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"spDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spGetPendingSwaps\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"SwapID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"DestAddr\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"RefundAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"AmountWBTC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structSwapContract.spPendingTx[]\",\"name\":\"data\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"callees\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"exchangeData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"startIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.SimpleData\",\"name\":\"_data\",\"type\":\"tuple\"}],\"name\":\"spParaSwapBTC2Token\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_destinationAddressForBTC\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"callees\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"exchangeData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"startIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.SimpleData\",\"name\":\"_data\",\"type\":\"tuple\"}],\"name\":\"spParaSwapToken2BTC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spPendingTXs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"SwapID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"DestAddr\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"RefundAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"AmountWBTC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tssThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalFeeBPS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _token, address _user) view returns(uint256)
func (_SwapContract *SwapContractCaller) BalanceOf(opts *bind.CallOpts, _token common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "balanceOf", _token, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _token, address _user) view returns(uint256)
func (_SwapContract *SwapContractSession) BalanceOf(_token common.Address, _user common.Address) (*big.Int, error) {
	return _SwapContract.Contract.BalanceOf(&_SwapContract.CallOpts, _token, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address _token, address _user) view returns(uint256)
func (_SwapContract *SwapContractCallerSession) BalanceOf(_token common.Address, _user common.Address) (*big.Int, error) {
	return _SwapContract.Contract.BalanceOf(&_SwapContract.CallOpts, _token, _user)
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

// DepositFeesBPS is a free data retrieval call binding the contract method 0x42419255.
//
// Solidity: function depositFeesBPS() view returns(uint8)
func (_SwapContract *SwapContractCaller) DepositFeesBPS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "depositFeesBPS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DepositFeesBPS is a free data retrieval call binding the contract method 0x42419255.
//
// Solidity: function depositFeesBPS() view returns(uint8)
func (_SwapContract *SwapContractSession) DepositFeesBPS() (uint8, error) {
	return _SwapContract.Contract.DepositFeesBPS(&_SwapContract.CallOpts)
}

// DepositFeesBPS is a free data retrieval call binding the contract method 0x42419255.
//
// Solidity: function depositFeesBPS() view returns(uint8)
func (_SwapContract *SwapContractCallerSession) DepositFeesBPS() (uint8, error) {
	return _SwapContract.Contract.DepositFeesBPS(&_SwapContract.CallOpts)
}

// ExpirationTime is a free data retrieval call binding the contract method 0xda284dcc.
//
// Solidity: function expirationTime() view returns(uint256)
func (_SwapContract *SwapContractCaller) ExpirationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "expirationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpirationTime is a free data retrieval call binding the contract method 0xda284dcc.
//
// Solidity: function expirationTime() view returns(uint256)
func (_SwapContract *SwapContractSession) ExpirationTime() (*big.Int, error) {
	return _SwapContract.Contract.ExpirationTime(&_SwapContract.CallOpts)
}

// ExpirationTime is a free data retrieval call binding the contract method 0xda284dcc.
//
// Solidity: function expirationTime() view returns(uint256)
func (_SwapContract *SwapContractCallerSession) ExpirationTime() (*big.Int, error) {
	return _SwapContract.Contract.ExpirationTime(&_SwapContract.CallOpts)
}

// FeesLPTokensForNode is a free data retrieval call binding the contract method 0xf7824bc7.
//
// Solidity: function feesLPTokensForNode() view returns(uint256)
func (_SwapContract *SwapContractCaller) FeesLPTokensForNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "feesLPTokensForNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeesLPTokensForNode is a free data retrieval call binding the contract method 0xf7824bc7.
//
// Solidity: function feesLPTokensForNode() view returns(uint256)
func (_SwapContract *SwapContractSession) FeesLPTokensForNode() (*big.Int, error) {
	return _SwapContract.Contract.FeesLPTokensForNode(&_SwapContract.CallOpts)
}

// FeesLPTokensForNode is a free data retrieval call binding the contract method 0xf7824bc7.
//
// Solidity: function feesLPTokensForNode() view returns(uint256)
func (_SwapContract *SwapContractCallerSession) FeesLPTokensForNode() (*big.Int, error) {
	return _SwapContract.Contract.FeesLPTokensForNode(&_SwapContract.CallOpts)
}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(bytes32[])
func (_SwapContract *SwapContractCaller) GetActiveNodes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "getActiveNodes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(bytes32[])
func (_SwapContract *SwapContractSession) GetActiveNodes() ([][32]byte, error) {
	return _SwapContract.Contract.GetActiveNodes(&_SwapContract.CallOpts)
}

// GetActiveNodes is a free data retrieval call binding the contract method 0x6b51e919.
//
// Solidity: function getActiveNodes() view returns(bytes32[])
func (_SwapContract *SwapContractCallerSession) GetActiveNodes() ([][32]byte, error) {
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

// GetDepositFeeRate is a free data retrieval call binding the contract method 0x3b6fa127.
//
// Solidity: function getDepositFeeRate(address _token, uint256 _amountOfFloat) view returns(uint256 depositFeeRate)
func (_SwapContract *SwapContractCaller) GetDepositFeeRate(opts *bind.CallOpts, _token common.Address, _amountOfFloat *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "getDepositFeeRate", _token, _amountOfFloat)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositFeeRate is a free data retrieval call binding the contract method 0x3b6fa127.
//
// Solidity: function getDepositFeeRate(address _token, uint256 _amountOfFloat) view returns(uint256 depositFeeRate)
func (_SwapContract *SwapContractSession) GetDepositFeeRate(_token common.Address, _amountOfFloat *big.Int) (*big.Int, error) {
	return _SwapContract.Contract.GetDepositFeeRate(&_SwapContract.CallOpts, _token, _amountOfFloat)
}

// GetDepositFeeRate is a free data retrieval call binding the contract method 0x3b6fa127.
//
// Solidity: function getDepositFeeRate(address _token, uint256 _amountOfFloat) view returns(uint256 depositFeeRate)
func (_SwapContract *SwapContractCallerSession) GetDepositFeeRate(_token common.Address, _amountOfFloat *big.Int) (*big.Int, error) {
	return _SwapContract.Contract.GetDepositFeeRate(&_SwapContract.CallOpts, _token, _amountOfFloat)
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

// InitialExchangeRate is a free data retrieval call binding the contract method 0x6d019f35.
//
// Solidity: function initialExchangeRate() view returns(uint256)
func (_SwapContract *SwapContractCaller) InitialExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "initialExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialExchangeRate is a free data retrieval call binding the contract method 0x6d019f35.
//
// Solidity: function initialExchangeRate() view returns(uint256)
func (_SwapContract *SwapContractSession) InitialExchangeRate() (*big.Int, error) {
	return _SwapContract.Contract.InitialExchangeRate(&_SwapContract.CallOpts)
}

// InitialExchangeRate is a free data retrieval call binding the contract method 0x6d019f35.
//
// Solidity: function initialExchangeRate() view returns(uint256)
func (_SwapContract *SwapContractCallerSession) InitialExchangeRate() (*big.Int, error) {
	return _SwapContract.Contract.InitialExchangeRate(&_SwapContract.CallOpts)
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

// LockedLPTokensForNode is a free data retrieval call binding the contract method 0x2d55282d.
//
// Solidity: function lockedLPTokensForNode() view returns(uint256)
func (_SwapContract *SwapContractCaller) LockedLPTokensForNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "lockedLPTokensForNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedLPTokensForNode is a free data retrieval call binding the contract method 0x2d55282d.
//
// Solidity: function lockedLPTokensForNode() view returns(uint256)
func (_SwapContract *SwapContractSession) LockedLPTokensForNode() (*big.Int, error) {
	return _SwapContract.Contract.LockedLPTokensForNode(&_SwapContract.CallOpts)
}

// LockedLPTokensForNode is a free data retrieval call binding the contract method 0x2d55282d.
//
// Solidity: function lockedLPTokensForNode() view returns(uint256)
func (_SwapContract *SwapContractCallerSession) LockedLPTokensForNode() (*big.Int, error) {
	return _SwapContract.Contract.LockedLPTokensForNode(&_SwapContract.CallOpts)
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
// Solidity: function nodeRewardsRatio() view returns(uint8)
func (_SwapContract *SwapContractCaller) NodeRewardsRatio(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "nodeRewardsRatio")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

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

// SpGetPendingSwaps is a free data retrieval call binding the contract method 0x20199e99.
//
// Solidity: function spGetPendingSwaps() view returns((bytes32,bytes32,address,uint256,uint256)[] data)
func (_SwapContract *SwapContractCaller) SpGetPendingSwaps(opts *bind.CallOpts) ([]SwapContractspPendingTx, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "spGetPendingSwaps")

	if err != nil {
		return *new([]SwapContractspPendingTx), err
	}

	out0 := *abi.ConvertType(out[0], new([]SwapContractspPendingTx)).(*[]SwapContractspPendingTx)

	return out0, err

}

// SpGetPendingSwaps is a free data retrieval call binding the contract method 0x20199e99.
//
// Solidity: function spGetPendingSwaps() view returns((bytes32,bytes32,address,uint256,uint256)[] data)
func (_SwapContract *SwapContractSession) SpGetPendingSwaps() ([]SwapContractspPendingTx, error) {
	return _SwapContract.Contract.SpGetPendingSwaps(&_SwapContract.CallOpts)
}

// SpGetPendingSwaps is a free data retrieval call binding the contract method 0x20199e99.
//
// Solidity: function spGetPendingSwaps() view returns((bytes32,bytes32,address,uint256,uint256)[] data)
func (_SwapContract *SwapContractCallerSession) SpGetPendingSwaps() ([]SwapContractspPendingTx, error) {
	return _SwapContract.Contract.SpGetPendingSwaps(&_SwapContract.CallOpts)
}

// SpPendingTXs is a free data retrieval call binding the contract method 0x1ab7e937.
//
// Solidity: function spPendingTXs(uint256 ) view returns(bytes32 SwapID, bytes32 DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContract *SwapContractCaller) SpPendingTXs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SwapID     [32]byte
	DestAddr   [32]byte
	RefundAddr common.Address
	AmountWBTC *big.Int
	Timestamp  *big.Int
}, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "spPendingTXs", arg0)

	outstruct := new(struct {
		SwapID     [32]byte
		DestAddr   [32]byte
		RefundAddr common.Address
		AmountWBTC *big.Int
		Timestamp  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SwapID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DestAddr = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.RefundAddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.AmountWBTC = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SpPendingTXs is a free data retrieval call binding the contract method 0x1ab7e937.
//
// Solidity: function spPendingTXs(uint256 ) view returns(bytes32 SwapID, bytes32 DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContract *SwapContractSession) SpPendingTXs(arg0 *big.Int) (struct {
	SwapID     [32]byte
	DestAddr   [32]byte
	RefundAddr common.Address
	AmountWBTC *big.Int
	Timestamp  *big.Int
}, error) {
	return _SwapContract.Contract.SpPendingTXs(&_SwapContract.CallOpts, arg0)
}

// SpPendingTXs is a free data retrieval call binding the contract method 0x1ab7e937.
//
// Solidity: function spPendingTXs(uint256 ) view returns(bytes32 SwapID, bytes32 DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContract *SwapContractCallerSession) SpPendingTXs(arg0 *big.Int) (struct {
	SwapID     [32]byte
	DestAddr   [32]byte
	RefundAddr common.Address
	AmountWBTC *big.Int
	Timestamp  *big.Int
}, error) {
	return _SwapContract.Contract.SpPendingTXs(&_SwapContract.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens(address , address ) view returns(uint256)
func (_SwapContract *SwapContractCaller) Tokens(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "tokens", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens(address , address ) view returns(uint256)
func (_SwapContract *SwapContractSession) Tokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SwapContract.Contract.Tokens(&_SwapContract.CallOpts, arg0, arg1)
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens(address , address ) view returns(uint256)
func (_SwapContract *SwapContractCallerSession) Tokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SwapContract.Contract.Tokens(&_SwapContract.CallOpts, arg0, arg1)
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

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_SwapContract *SwapContractCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "wETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_SwapContract *SwapContractSession) WETH() (common.Address, error) {
	return _SwapContract.Contract.WETH(&_SwapContract.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_SwapContract *SwapContractCallerSession) WETH() (common.Address, error) {
	return _SwapContract.Contract.WETH(&_SwapContract.CallOpts)
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
// Solidity: function withdrawalFeeBPS() view returns(uint8)
func (_SwapContract *SwapContractCaller) WithdrawalFeeBPS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SwapContract.contract.Call(opts, &out, "withdrawalFeeBPS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// WithdrawalFeeBPS is a free data retrieval call binding the contract method 0xb6268e5d.
//
// Solidity: function withdrawalFeeBPS() view returns(uint8)
func (_SwapContract *SwapContractSession) WithdrawalFeeBPS() (uint8, error) {
	return _SwapContract.Contract.WithdrawalFeeBPS(&_SwapContract.CallOpts)
}

// WithdrawalFeeBPS is a free data retrieval call binding the contract method 0xb6268e5d.
//
// Solidity: function withdrawalFeeBPS() view returns(uint8)
func (_SwapContract *SwapContractCallerSession) WithdrawalFeeBPS() (uint8, error) {
	return _SwapContract.Contract.WithdrawalFeeBPS(&_SwapContract.CallOpts)
}

// Churn is a paid mutator transaction binding the contract method 0x7c747cf9.
//
// Solidity: function churn(address _newOwner, bytes32[] _rewardAddressAndAmounts, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold, uint8 _nodeRewardsRatio, uint8 _withdrawalFeeBPS, uint256 _expirationTime) returns(bool)
func (_SwapContract *SwapContractTransactor) Churn(opts *bind.TransactOpts, _newOwner common.Address, _rewardAddressAndAmounts [][32]byte, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8, _nodeRewardsRatio uint8, _withdrawalFeeBPS uint8, _expirationTime *big.Int) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "churn", _newOwner, _rewardAddressAndAmounts, _isRemoved, _churnedInCount, _tssThreshold, _nodeRewardsRatio, _withdrawalFeeBPS, _expirationTime)
}

// Churn is a paid mutator transaction binding the contract method 0x7c747cf9.
//
// Solidity: function churn(address _newOwner, bytes32[] _rewardAddressAndAmounts, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold, uint8 _nodeRewardsRatio, uint8 _withdrawalFeeBPS, uint256 _expirationTime) returns(bool)
func (_SwapContract *SwapContractSession) Churn(_newOwner common.Address, _rewardAddressAndAmounts [][32]byte, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8, _nodeRewardsRatio uint8, _withdrawalFeeBPS uint8, _expirationTime *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.Churn(&_SwapContract.TransactOpts, _newOwner, _rewardAddressAndAmounts, _isRemoved, _churnedInCount, _tssThreshold, _nodeRewardsRatio, _withdrawalFeeBPS, _expirationTime)
}

// Churn is a paid mutator transaction binding the contract method 0x7c747cf9.
//
// Solidity: function churn(address _newOwner, bytes32[] _rewardAddressAndAmounts, bool[] _isRemoved, uint8 _churnedInCount, uint8 _tssThreshold, uint8 _nodeRewardsRatio, uint8 _withdrawalFeeBPS, uint256 _expirationTime) returns(bool)
func (_SwapContract *SwapContractTransactorSession) Churn(_newOwner common.Address, _rewardAddressAndAmounts [][32]byte, _isRemoved []bool, _churnedInCount uint8, _tssThreshold uint8, _nodeRewardsRatio uint8, _withdrawalFeeBPS uint8, _expirationTime *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.Churn(&_SwapContract.TransactOpts, _newOwner, _rewardAddressAndAmounts, _isRemoved, _churnedInCount, _tssThreshold, _nodeRewardsRatio, _withdrawalFeeBPS, _expirationTime)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x8006cc3e.
//
// Solidity: function collectSwapFeesForBTC(address _destToken, uint256 _incomingAmount, uint256 _minerFee, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractTransactor) CollectSwapFeesForBTC(opts *bind.TransactOpts, _destToken common.Address, _incomingAmount *big.Int, _minerFee *big.Int, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "collectSwapFeesForBTC", _destToken, _incomingAmount, _minerFee, _rewardsAmount)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x8006cc3e.
//
// Solidity: function collectSwapFeesForBTC(address _destToken, uint256 _incomingAmount, uint256 _minerFee, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractSession) CollectSwapFeesForBTC(_destToken common.Address, _incomingAmount *big.Int, _minerFee *big.Int, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.CollectSwapFeesForBTC(&_SwapContract.TransactOpts, _destToken, _incomingAmount, _minerFee, _rewardsAmount)
}

// CollectSwapFeesForBTC is a paid mutator transaction binding the contract method 0x8006cc3e.
//
// Solidity: function collectSwapFeesForBTC(address _destToken, uint256 _incomingAmount, uint256 _minerFee, uint256 _rewardsAmount) returns(bool)
func (_SwapContract *SwapContractTransactorSession) CollectSwapFeesForBTC(_destToken common.Address, _incomingAmount *big.Int, _minerFee *big.Int, _rewardsAmount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.CollectSwapFeesForBTC(&_SwapContract.TransactOpts, _destToken, _incomingAmount, _minerFee, _rewardsAmount)
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

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0x63f289d9.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bool _zerofee, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactor) RecordIncomingFloat(opts *bind.TransactOpts, _token common.Address, _addressesAndAmountOfFloat [32]byte, _zerofee bool, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "recordIncomingFloat", _token, _addressesAndAmountOfFloat, _zerofee, _txid)
}

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0x63f289d9.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bool _zerofee, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractSession) RecordIncomingFloat(_token common.Address, _addressesAndAmountOfFloat [32]byte, _zerofee bool, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordIncomingFloat(&_SwapContract.TransactOpts, _token, _addressesAndAmountOfFloat, _zerofee, _txid)
}

// RecordIncomingFloat is a paid mutator transaction binding the contract method 0x63f289d9.
//
// Solidity: function recordIncomingFloat(address _token, bytes32 _addressesAndAmountOfFloat, bool _zerofee, bytes32 _txid) returns(bool)
func (_SwapContract *SwapContractTransactorSession) RecordIncomingFloat(_token common.Address, _addressesAndAmountOfFloat [32]byte, _zerofee bool, _txid [32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordIncomingFloat(&_SwapContract.TransactOpts, _token, _addressesAndAmountOfFloat, _zerofee, _txid)
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

// RecordSkyPoolsTX is a paid mutator transaction binding the contract method 0xd6b2779a.
//
// Solidity: function recordSkyPoolsTX(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContract *SwapContractTransactor) RecordSkyPoolsTX(opts *bind.TransactOpts, _destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "recordSkyPoolsTX", _destToken, _to, _amount, _totalSwapped, _redeemedFloatTxIds)
}

// RecordSkyPoolsTX is a paid mutator transaction binding the contract method 0xd6b2779a.
//
// Solidity: function recordSkyPoolsTX(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContract *SwapContractSession) RecordSkyPoolsTX(_destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordSkyPoolsTX(&_SwapContract.TransactOpts, _destToken, _to, _amount, _totalSwapped, _redeemedFloatTxIds)
}

// RecordSkyPoolsTX is a paid mutator transaction binding the contract method 0xd6b2779a.
//
// Solidity: function recordSkyPoolsTX(address _destToken, address _to, uint256 _amount, uint256 _totalSwapped, bytes32[] _redeemedFloatTxIds) returns(bool)
func (_SwapContract *SwapContractTransactorSession) RecordSkyPoolsTX(_destToken common.Address, _to common.Address, _amount *big.Int, _totalSwapped *big.Int, _redeemedFloatTxIds [][32]byte) (*types.Transaction, error) {
	return _SwapContract.Contract.RecordSkyPoolsTX(&_SwapContract.TransactOpts, _destToken, _to, _amount, _totalSwapped, _redeemedFloatTxIds)
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

// RedeemERC20Token is a paid mutator transaction binding the contract method 0x4db75485.
//
// Solidity: function redeemERC20Token(address _token, uint256 _amount) returns()
func (_SwapContract *SwapContractTransactor) RedeemERC20Token(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "redeemERC20Token", _token, _amount)
}

// RedeemERC20Token is a paid mutator transaction binding the contract method 0x4db75485.
//
// Solidity: function redeemERC20Token(address _token, uint256 _amount) returns()
func (_SwapContract *SwapContractSession) RedeemERC20Token(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.RedeemERC20Token(&_SwapContract.TransactOpts, _token, _amount)
}

// RedeemERC20Token is a paid mutator transaction binding the contract method 0x4db75485.
//
// Solidity: function redeemERC20Token(address _token, uint256 _amount) returns()
func (_SwapContract *SwapContractTransactorSession) RedeemERC20Token(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.RedeemERC20Token(&_SwapContract.TransactOpts, _token, _amount)
}

// RedeemEther is a paid mutator transaction binding the contract method 0x8535490f.
//
// Solidity: function redeemEther(uint256 _amount) payable returns()
func (_SwapContract *SwapContractTransactor) RedeemEther(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "redeemEther", _amount)
}

// RedeemEther is a paid mutator transaction binding the contract method 0x8535490f.
//
// Solidity: function redeemEther(uint256 _amount) payable returns()
func (_SwapContract *SwapContractSession) RedeemEther(_amount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.RedeemEther(&_SwapContract.TransactOpts, _amount)
}

// RedeemEther is a paid mutator transaction binding the contract method 0x8535490f.
//
// Solidity: function redeemEther(uint256 _amount) payable returns()
func (_SwapContract *SwapContractTransactorSession) RedeemEther(_amount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.RedeemEther(&_SwapContract.TransactOpts, _amount)
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

// SpCleanUpOldTXs is a paid mutator transaction binding the contract method 0x21959d3c.
//
// Solidity: function spCleanUpOldTXs(uint256 _loopCount) returns()
func (_SwapContract *SwapContractTransactor) SpCleanUpOldTXs(opts *bind.TransactOpts, _loopCount *big.Int) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "spCleanUpOldTXs", _loopCount)
}

// SpCleanUpOldTXs is a paid mutator transaction binding the contract method 0x21959d3c.
//
// Solidity: function spCleanUpOldTXs(uint256 _loopCount) returns()
func (_SwapContract *SwapContractSession) SpCleanUpOldTXs(_loopCount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.SpCleanUpOldTXs(&_SwapContract.TransactOpts, _loopCount)
}

// SpCleanUpOldTXs is a paid mutator transaction binding the contract method 0x21959d3c.
//
// Solidity: function spCleanUpOldTXs(uint256 _loopCount) returns()
func (_SwapContract *SwapContractTransactorSession) SpCleanUpOldTXs(_loopCount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.SpCleanUpOldTXs(&_SwapContract.TransactOpts, _loopCount)
}

// SpDeposit is a paid mutator transaction binding the contract method 0xca0707e7.
//
// Solidity: function spDeposit(address _token, uint256 _amount) payable returns()
func (_SwapContract *SwapContractTransactor) SpDeposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "spDeposit", _token, _amount)
}

// SpDeposit is a paid mutator transaction binding the contract method 0xca0707e7.
//
// Solidity: function spDeposit(address _token, uint256 _amount) payable returns()
func (_SwapContract *SwapContractSession) SpDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.SpDeposit(&_SwapContract.TransactOpts, _token, _amount)
}

// SpDeposit is a paid mutator transaction binding the contract method 0xca0707e7.
//
// Solidity: function spDeposit(address _token, uint256 _amount) payable returns()
func (_SwapContract *SwapContractTransactorSession) SpDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SwapContract.Contract.SpDeposit(&_SwapContract.TransactOpts, _token, _amount)
}

// SpParaSwapBTC2Token is a paid mutator transaction binding the contract method 0xb59b54c3.
//
// Solidity: function spParaSwapBTC2Token((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) payable returns()
func (_SwapContract *SwapContractTransactor) SpParaSwapBTC2Token(opts *bind.TransactOpts, _data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "spParaSwapBTC2Token", _data)
}

// SpParaSwapBTC2Token is a paid mutator transaction binding the contract method 0xb59b54c3.
//
// Solidity: function spParaSwapBTC2Token((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) payable returns()
func (_SwapContract *SwapContractSession) SpParaSwapBTC2Token(_data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContract.Contract.SpParaSwapBTC2Token(&_SwapContract.TransactOpts, _data)
}

// SpParaSwapBTC2Token is a paid mutator transaction binding the contract method 0xb59b54c3.
//
// Solidity: function spParaSwapBTC2Token((address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) payable returns()
func (_SwapContract *SwapContractTransactorSession) SpParaSwapBTC2Token(_data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContract.Contract.SpParaSwapBTC2Token(&_SwapContract.TransactOpts, _data)
}

// SpParaSwapToken2BTC is a paid mutator transaction binding the contract method 0x08f6263b.
//
// Solidity: function spParaSwapToken2BTC(bytes32 _destinationAddressForBTC, (address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) returns()
func (_SwapContract *SwapContractTransactor) SpParaSwapToken2BTC(opts *bind.TransactOpts, _destinationAddressForBTC [32]byte, _data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContract.contract.Transact(opts, "spParaSwapToken2BTC", _destinationAddressForBTC, _data)
}

// SpParaSwapToken2BTC is a paid mutator transaction binding the contract method 0x08f6263b.
//
// Solidity: function spParaSwapToken2BTC(bytes32 _destinationAddressForBTC, (address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) returns()
func (_SwapContract *SwapContractSession) SpParaSwapToken2BTC(_destinationAddressForBTC [32]byte, _data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContract.Contract.SpParaSwapToken2BTC(&_SwapContract.TransactOpts, _destinationAddressForBTC, _data)
}

// SpParaSwapToken2BTC is a paid mutator transaction binding the contract method 0x08f6263b.
//
// Solidity: function spParaSwapToken2BTC(bytes32 _destinationAddressForBTC, (address,address,uint256,uint256,uint256,address[],bytes,uint256[],uint256[],address,address,uint256,bytes,uint256,bytes16) _data) returns()
func (_SwapContract *SwapContractTransactorSession) SpParaSwapToken2BTC(_destinationAddressForBTC [32]byte, _data UtilsSimpleData) (*types.Transaction, error) {
	return _SwapContract.Contract.SpParaSwapToken2BTC(&_SwapContract.TransactOpts, _destinationAddressForBTC, _data)
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

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapContract *SwapContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapContract *SwapContractSession) Receive() (*types.Transaction, error) {
	return _SwapContract.Contract.Receive(&_SwapContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapContract *SwapContractTransactorSession) Receive() (*types.Transaction, error) {
	return _SwapContract.Contract.Receive(&_SwapContract.TransactOpts)
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

// SwapContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SwapContract contract.
type SwapContractDepositIterator struct {
	Event *SwapContractDeposit // Event containing the contract specifics and raw log

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
func (it *SwapContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractDeposit)
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
		it.Event = new(SwapContractDeposit)
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
func (it *SwapContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractDeposit represents a Deposit event raised by the SwapContract contract.
type SwapContractDeposit struct {
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
func (_SwapContract *SwapContractFilterer) FilterDeposit(opts *bind.FilterOpts) (*SwapContractDepositIterator, error) {

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &SwapContractDepositIterator{contract: _SwapContract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address token, address user, uint256 amount, uint256 balance, uint256 Timestamp)
func (_SwapContract *SwapContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SwapContractDeposit) (event.Subscription, error) {

	logs, sub, err := _SwapContract.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractDeposit)
				if err := _SwapContract.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_SwapContract *SwapContractFilterer) ParseDeposit(log types.Log) (*SwapContractDeposit, error) {
	event := new(SwapContractDeposit)
	if err := _SwapContract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractDistributeNodeRewardsIterator is returned from FilterDistributeNodeRewards and is used to iterate over the raw logs and unpacked data for DistributeNodeRewards events raised by the SwapContract contract.
type SwapContractDistributeNodeRewardsIterator struct {
	Event *SwapContractDistributeNodeRewards // Event containing the contract specifics and raw log

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
func (it *SwapContractDistributeNodeRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractDistributeNodeRewards)
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
		it.Event = new(SwapContractDistributeNodeRewards)
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
func (it *SwapContractDistributeNodeRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractDistributeNodeRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractDistributeNodeRewards represents a DistributeNodeRewards event raised by the SwapContract contract.
type SwapContractDistributeNodeRewards struct {
	RewardLPTsForNodes *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDistributeNodeRewards is a free log retrieval operation binding the contract event 0x86dd6ae252e62228b47b3c1a12762cbf634952b2d29319b04f4f72fb153bcbdd.
//
// Solidity: event DistributeNodeRewards(uint256 rewardLPTsForNodes)
func (_SwapContract *SwapContractFilterer) FilterDistributeNodeRewards(opts *bind.FilterOpts) (*SwapContractDistributeNodeRewardsIterator, error) {

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "DistributeNodeRewards")
	if err != nil {
		return nil, err
	}
	return &SwapContractDistributeNodeRewardsIterator{contract: _SwapContract.contract, event: "DistributeNodeRewards", logs: logs, sub: sub}, nil
}

// WatchDistributeNodeRewards is a free log subscription operation binding the contract event 0x86dd6ae252e62228b47b3c1a12762cbf634952b2d29319b04f4f72fb153bcbdd.
//
// Solidity: event DistributeNodeRewards(uint256 rewardLPTsForNodes)
func (_SwapContract *SwapContractFilterer) WatchDistributeNodeRewards(opts *bind.WatchOpts, sink chan<- *SwapContractDistributeNodeRewards) (event.Subscription, error) {

	logs, sub, err := _SwapContract.contract.WatchLogs(opts, "DistributeNodeRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractDistributeNodeRewards)
				if err := _SwapContract.contract.UnpackLog(event, "DistributeNodeRewards", log); err != nil {
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
func (_SwapContract *SwapContractFilterer) ParseDistributeNodeRewards(log types.Log) (*SwapContractDistributeNodeRewards, error) {
	event := new(SwapContractDistributeNodeRewards)
	if err := _SwapContract.contract.UnpackLog(event, "DistributeNodeRewards", log); err != nil {
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
	FeesToken             common.Address
	Rewards               *big.Int
	AmountLPTokensForNode *big.Int
	CurrentPriceLP        *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRewardsCollection is a free log retrieval operation binding the contract event 0xe3debe835f6848edc082b32a1d729781d3bfcd7e14422d80bcaa6794d3816b2f.
//
// Solidity: event RewardsCollection(address feesToken, uint256 rewards, uint256 amountLPTokensForNode, uint256 currentPriceLP)
func (_SwapContract *SwapContractFilterer) FilterRewardsCollection(opts *bind.FilterOpts) (*SwapContractRewardsCollectionIterator, error) {

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "RewardsCollection")
	if err != nil {
		return nil, err
	}
	return &SwapContractRewardsCollectionIterator{contract: _SwapContract.contract, event: "RewardsCollection", logs: logs, sub: sub}, nil
}

// WatchRewardsCollection is a free log subscription operation binding the contract event 0xe3debe835f6848edc082b32a1d729781d3bfcd7e14422d80bcaa6794d3816b2f.
//
// Solidity: event RewardsCollection(address feesToken, uint256 rewards, uint256 amountLPTokensForNode, uint256 currentPriceLP)
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
// Solidity: event RewardsCollection(address feesToken, uint256 rewards, uint256 amountLPTokensForNode, uint256 currentPriceLP)
func (_SwapContract *SwapContractFilterer) ParseRewardsCollection(log types.Log) (*SwapContractRewardsCollection, error) {
	event := new(SwapContractRewardsCollection)
	if err := _SwapContract.contract.UnpackLog(event, "RewardsCollection", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractSetExpirationTimeIterator is returned from FilterSetExpirationTime and is used to iterate over the raw logs and unpacked data for SetExpirationTime events raised by the SwapContract contract.
type SwapContractSetExpirationTimeIterator struct {
	Event *SwapContractSetExpirationTime // Event containing the contract specifics and raw log

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
func (it *SwapContractSetExpirationTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractSetExpirationTime)
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
		it.Event = new(SwapContractSetExpirationTime)
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
func (it *SwapContractSetExpirationTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractSetExpirationTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractSetExpirationTime represents a SetExpirationTime event raised by the SwapContract contract.
type SwapContractSetExpirationTime struct {
	ExpirationTime *big.Int
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetExpirationTime is a free log retrieval operation binding the contract event 0xa1cc12e33e9907b789fd3c10ea0afd9d3e563e2db44d84c7269534f1046d1ca1.
//
// Solidity: event SetExpirationTime(uint256 ExpirationTime, uint256 Timestamp)
func (_SwapContract *SwapContractFilterer) FilterSetExpirationTime(opts *bind.FilterOpts) (*SwapContractSetExpirationTimeIterator, error) {

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "SetExpirationTime")
	if err != nil {
		return nil, err
	}
	return &SwapContractSetExpirationTimeIterator{contract: _SwapContract.contract, event: "SetExpirationTime", logs: logs, sub: sub}, nil
}

// WatchSetExpirationTime is a free log subscription operation binding the contract event 0xa1cc12e33e9907b789fd3c10ea0afd9d3e563e2db44d84c7269534f1046d1ca1.
//
// Solidity: event SetExpirationTime(uint256 ExpirationTime, uint256 Timestamp)
func (_SwapContract *SwapContractFilterer) WatchSetExpirationTime(opts *bind.WatchOpts, sink chan<- *SwapContractSetExpirationTime) (event.Subscription, error) {

	logs, sub, err := _SwapContract.contract.WatchLogs(opts, "SetExpirationTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractSetExpirationTime)
				if err := _SwapContract.contract.UnpackLog(event, "SetExpirationTime", log); err != nil {
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

// ParseSetExpirationTime is a log parse operation binding the contract event 0xa1cc12e33e9907b789fd3c10ea0afd9d3e563e2db44d84c7269534f1046d1ca1.
//
// Solidity: event SetExpirationTime(uint256 ExpirationTime, uint256 Timestamp)
func (_SwapContract *SwapContractFilterer) ParseSetExpirationTime(log types.Log) (*SwapContractSetExpirationTime, error) {
	event := new(SwapContractSetExpirationTime)
	if err := _SwapContract.contract.UnpackLog(event, "SetExpirationTime", log); err != nil {
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

// SwapContractSwapTokensToBTCIterator is returned from FilterSwapTokensToBTC and is used to iterate over the raw logs and unpacked data for SwapTokensToBTC events raised by the SwapContract contract.
type SwapContractSwapTokensToBTCIterator struct {
	Event *SwapContractSwapTokensToBTC // Event containing the contract specifics and raw log

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
func (it *SwapContractSwapTokensToBTCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractSwapTokensToBTC)
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
		it.Event = new(SwapContractSwapTokensToBTC)
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
func (it *SwapContractSwapTokensToBTCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractSwapTokensToBTCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractSwapTokensToBTC represents a SwapTokensToBTC event raised by the SwapContract contract.
type SwapContractSwapTokensToBTC struct {
	SwapID     [32]byte
	DestAddr   [32]byte
	RefundAddr common.Address
	AmountWBTC *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwapTokensToBTC is a free log retrieval operation binding the contract event 0x058800492a0e9d1ce3d60aa19b5a711354aaf0d831a5fa20197617dc4c0a573e.
//
// Solidity: event SwapTokensToBTC(bytes32 SwapID, bytes32 DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContract *SwapContractFilterer) FilterSwapTokensToBTC(opts *bind.FilterOpts) (*SwapContractSwapTokensToBTCIterator, error) {

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "SwapTokensToBTC")
	if err != nil {
		return nil, err
	}
	return &SwapContractSwapTokensToBTCIterator{contract: _SwapContract.contract, event: "SwapTokensToBTC", logs: logs, sub: sub}, nil
}

// WatchSwapTokensToBTC is a free log subscription operation binding the contract event 0x058800492a0e9d1ce3d60aa19b5a711354aaf0d831a5fa20197617dc4c0a573e.
//
// Solidity: event SwapTokensToBTC(bytes32 SwapID, bytes32 DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContract *SwapContractFilterer) WatchSwapTokensToBTC(opts *bind.WatchOpts, sink chan<- *SwapContractSwapTokensToBTC) (event.Subscription, error) {

	logs, sub, err := _SwapContract.contract.WatchLogs(opts, "SwapTokensToBTC")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractSwapTokensToBTC)
				if err := _SwapContract.contract.UnpackLog(event, "SwapTokensToBTC", log); err != nil {
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

// ParseSwapTokensToBTC is a log parse operation binding the contract event 0x058800492a0e9d1ce3d60aa19b5a711354aaf0d831a5fa20197617dc4c0a573e.
//
// Solidity: event SwapTokensToBTC(bytes32 SwapID, bytes32 DestAddr, address RefundAddr, uint256 AmountWBTC, uint256 Timestamp)
func (_SwapContract *SwapContractFilterer) ParseSwapTokensToBTC(log types.Log) (*SwapContractSwapTokensToBTC, error) {
	event := new(SwapContractSwapTokensToBTC)
	if err := _SwapContract.contract.UnpackLog(event, "SwapTokensToBTC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SwapContract contract.
type SwapContractWithdrawIterator struct {
	Event *SwapContractWithdraw // Event containing the contract specifics and raw log

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
func (it *SwapContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapContractWithdraw)
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
		it.Event = new(SwapContractWithdraw)
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
func (it *SwapContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapContractWithdraw represents a Withdraw event raised by the SwapContract contract.
type SwapContractWithdraw struct {
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
func (_SwapContract *SwapContractFilterer) FilterWithdraw(opts *bind.FilterOpts) (*SwapContractWithdrawIterator, error) {

	logs, sub, err := _SwapContract.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &SwapContractWithdrawIterator{contract: _SwapContract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address token, address user, uint256 amount, uint256 balance, uint256 Timestamp)
func (_SwapContract *SwapContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SwapContractWithdraw) (event.Subscription, error) {

	logs, sub, err := _SwapContract.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapContractWithdraw)
				if err := _SwapContract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_SwapContract *SwapContractFilterer) ParseWithdraw(log types.Log) (*SwapContractWithdraw, error) {
	event := new(SwapContractWithdraw)
	if err := _SwapContract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
