// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply_\",\"type\":\"uint256\"},{\"internalType\":\"address[4]\",\"name\":\"addrs\",\"type\":\"address[4]\"},{\"internalType\":\"uint256[4]\",\"name\":\"buyFeeSetting_\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[4]\",\"name\":\"sellFeeSetting_\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceForReward_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"serviceFeeReceiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExcluded\",\"type\":\"bool\"}],\"name\":\"ExcludeFromFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isExcluded\",\"type\":\"bool\"}],\"name\":\"ExcludeMultipleAccountsFromFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"}],\"name\":\"GasForProcessingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newLiquidityWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldLiquidityWallet\",\"type\":\"address\"}],\"name\":\"LiquidityWalletUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claims\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastProcessedIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"automatic\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"processor\",\"type\":\"address\"}],\"name\":\"ProcessedDividendTracker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensSwapped\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SendDividends\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"SetAutomatedMarketMakerPair\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensSwapped\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensIntoLiqudity\",\"type\":\"uint256\"}],\"name\":\"SwapAndLiquify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"}],\"name\":\"UpdateDividendTracker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"}],\"name\":\"UpdateUniswapV2Router\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AmountLiquidityFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AmountMarketingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AmountTokenRewardsFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_marketingWalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"airdropNumbs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"automatedMarketMakerPairs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyDeadFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyLiquidityFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyMarketingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyTokenRewardsFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"dividendTokenBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dividendTracker\",\"outputs\":[{\"internalType\":\"contractBABYTOKENDividendTracker\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"excludeFromDividends\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"excluded\",\"type\":\"bool\"}],\"name\":\"excludeFromFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"excluded\",\"type\":\"bool\"}],\"name\":\"excludeMultipleAccountsFromFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"first\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasForProcessing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountDividendsInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAccountDividendsInfoAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimWait\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastProcessedIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinimumTokenBalanceForDividends\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfDividendTokenHolders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalDividendsDistributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromDividends\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isExcludedFromFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"excluded\",\"type\":\"bool\"}],\"name\":\"multipleBotlistAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"processDividendTracker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellDeadFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellLiquidityFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellMarketingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellTokenRewardsFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setAirdropNumbs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setAutomatedMarketMakerPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadFee\",\"type\":\"uint256\"}],\"name\":\"setBuyTaxes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setDeadWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setKing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"setMarketingWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadFee\",\"type\":\"uint256\"}],\"name\":\"setSelTaxes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setSwapAndLiquifyEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setSwapTokensAtAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapAndLiquifyEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapManual\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapTokensAtAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimWait\",\"type\":\"uint256\"}],\"name\":\"updateClaimWait\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"updateGasForProcessing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"updateMinimumTokenBalanceForDividends\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"updateUniswapV2Router\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawableDividendOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// AmountLiquidityFee is a free data retrieval call binding the contract method 0xf8326795.
//
// Solidity: function AmountLiquidityFee() view returns(uint256)
func (_Token *TokenCaller) AmountLiquidityFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "AmountLiquidityFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountLiquidityFee is a free data retrieval call binding the contract method 0xf8326795.
//
// Solidity: function AmountLiquidityFee() view returns(uint256)
func (_Token *TokenSession) AmountLiquidityFee() (*big.Int, error) {
	return _Token.Contract.AmountLiquidityFee(&_Token.CallOpts)
}

// AmountLiquidityFee is a free data retrieval call binding the contract method 0xf8326795.
//
// Solidity: function AmountLiquidityFee() view returns(uint256)
func (_Token *TokenCallerSession) AmountLiquidityFee() (*big.Int, error) {
	return _Token.Contract.AmountLiquidityFee(&_Token.CallOpts)
}

// AmountMarketingFee is a free data retrieval call binding the contract method 0xcfe0e619.
//
// Solidity: function AmountMarketingFee() view returns(uint256)
func (_Token *TokenCaller) AmountMarketingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "AmountMarketingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountMarketingFee is a free data retrieval call binding the contract method 0xcfe0e619.
//
// Solidity: function AmountMarketingFee() view returns(uint256)
func (_Token *TokenSession) AmountMarketingFee() (*big.Int, error) {
	return _Token.Contract.AmountMarketingFee(&_Token.CallOpts)
}

// AmountMarketingFee is a free data retrieval call binding the contract method 0xcfe0e619.
//
// Solidity: function AmountMarketingFee() view returns(uint256)
func (_Token *TokenCallerSession) AmountMarketingFee() (*big.Int, error) {
	return _Token.Contract.AmountMarketingFee(&_Token.CallOpts)
}

// AmountTokenRewardsFee is a free data retrieval call binding the contract method 0x3926876d.
//
// Solidity: function AmountTokenRewardsFee() view returns(uint256)
func (_Token *TokenCaller) AmountTokenRewardsFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "AmountTokenRewardsFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountTokenRewardsFee is a free data retrieval call binding the contract method 0x3926876d.
//
// Solidity: function AmountTokenRewardsFee() view returns(uint256)
func (_Token *TokenSession) AmountTokenRewardsFee() (*big.Int, error) {
	return _Token.Contract.AmountTokenRewardsFee(&_Token.CallOpts)
}

// AmountTokenRewardsFee is a free data retrieval call binding the contract method 0x3926876d.
//
// Solidity: function AmountTokenRewardsFee() view returns(uint256)
func (_Token *TokenCallerSession) AmountTokenRewardsFee() (*big.Int, error) {
	return _Token.Contract.AmountTokenRewardsFee(&_Token.CallOpts)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0x1cdd3be3.
//
// Solidity: function _isBlacklisted(address ) view returns(bool)
func (_Token *TokenCaller) IsBlacklisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "_isBlacklisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0x1cdd3be3.
//
// Solidity: function _isBlacklisted(address ) view returns(bool)
func (_Token *TokenSession) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _Token.Contract.IsBlacklisted(&_Token.CallOpts, arg0)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0x1cdd3be3.
//
// Solidity: function _isBlacklisted(address ) view returns(bool)
func (_Token *TokenCallerSession) IsBlacklisted(arg0 common.Address) (bool, error) {
	return _Token.Contract.IsBlacklisted(&_Token.CallOpts, arg0)
}

// MarketingWalletAddress is a free data retrieval call binding the contract method 0x4144d9e4.
//
// Solidity: function _marketingWalletAddress() view returns(address)
func (_Token *TokenCaller) MarketingWalletAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "_marketingWalletAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketingWalletAddress is a free data retrieval call binding the contract method 0x4144d9e4.
//
// Solidity: function _marketingWalletAddress() view returns(address)
func (_Token *TokenSession) MarketingWalletAddress() (common.Address, error) {
	return _Token.Contract.MarketingWalletAddress(&_Token.CallOpts)
}

// MarketingWalletAddress is a free data retrieval call binding the contract method 0x4144d9e4.
//
// Solidity: function _marketingWalletAddress() view returns(address)
func (_Token *TokenCallerSession) MarketingWalletAddress() (common.Address, error) {
	return _Token.Contract.MarketingWalletAddress(&_Token.CallOpts)
}

// AirdropNumbs is a free data retrieval call binding the contract method 0xe32759cf.
//
// Solidity: function airdropNumbs() view returns(uint256)
func (_Token *TokenCaller) AirdropNumbs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "airdropNumbs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AirdropNumbs is a free data retrieval call binding the contract method 0xe32759cf.
//
// Solidity: function airdropNumbs() view returns(uint256)
func (_Token *TokenSession) AirdropNumbs() (*big.Int, error) {
	return _Token.Contract.AirdropNumbs(&_Token.CallOpts)
}

// AirdropNumbs is a free data retrieval call binding the contract method 0xe32759cf.
//
// Solidity: function airdropNumbs() view returns(uint256)
func (_Token *TokenCallerSession) AirdropNumbs() (*big.Int, error) {
	return _Token.Contract.AirdropNumbs(&_Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// AutomatedMarketMakerPairs is a free data retrieval call binding the contract method 0xb62496f5.
//
// Solidity: function automatedMarketMakerPairs(address ) view returns(bool)
func (_Token *TokenCaller) AutomatedMarketMakerPairs(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "automatedMarketMakerPairs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AutomatedMarketMakerPairs is a free data retrieval call binding the contract method 0xb62496f5.
//
// Solidity: function automatedMarketMakerPairs(address ) view returns(bool)
func (_Token *TokenSession) AutomatedMarketMakerPairs(arg0 common.Address) (bool, error) {
	return _Token.Contract.AutomatedMarketMakerPairs(&_Token.CallOpts, arg0)
}

// AutomatedMarketMakerPairs is a free data retrieval call binding the contract method 0xb62496f5.
//
// Solidity: function automatedMarketMakerPairs(address ) view returns(bool)
func (_Token *TokenCallerSession) AutomatedMarketMakerPairs(arg0 common.Address) (bool, error) {
	return _Token.Contract.AutomatedMarketMakerPairs(&_Token.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// BuyDeadFee is a free data retrieval call binding the contract method 0x8de743b3.
//
// Solidity: function buyDeadFee() view returns(uint256)
func (_Token *TokenCaller) BuyDeadFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "buyDeadFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyDeadFee is a free data retrieval call binding the contract method 0x8de743b3.
//
// Solidity: function buyDeadFee() view returns(uint256)
func (_Token *TokenSession) BuyDeadFee() (*big.Int, error) {
	return _Token.Contract.BuyDeadFee(&_Token.CallOpts)
}

// BuyDeadFee is a free data retrieval call binding the contract method 0x8de743b3.
//
// Solidity: function buyDeadFee() view returns(uint256)
func (_Token *TokenCallerSession) BuyDeadFee() (*big.Int, error) {
	return _Token.Contract.BuyDeadFee(&_Token.CallOpts)
}

// BuyLiquidityFee is a free data retrieval call binding the contract method 0xf11a24d3.
//
// Solidity: function buyLiquidityFee() view returns(uint256)
func (_Token *TokenCaller) BuyLiquidityFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "buyLiquidityFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyLiquidityFee is a free data retrieval call binding the contract method 0xf11a24d3.
//
// Solidity: function buyLiquidityFee() view returns(uint256)
func (_Token *TokenSession) BuyLiquidityFee() (*big.Int, error) {
	return _Token.Contract.BuyLiquidityFee(&_Token.CallOpts)
}

// BuyLiquidityFee is a free data retrieval call binding the contract method 0xf11a24d3.
//
// Solidity: function buyLiquidityFee() view returns(uint256)
func (_Token *TokenCallerSession) BuyLiquidityFee() (*big.Int, error) {
	return _Token.Contract.BuyLiquidityFee(&_Token.CallOpts)
}

// BuyMarketingFee is a free data retrieval call binding the contract method 0x7bce5a04.
//
// Solidity: function buyMarketingFee() view returns(uint256)
func (_Token *TokenCaller) BuyMarketingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "buyMarketingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyMarketingFee is a free data retrieval call binding the contract method 0x7bce5a04.
//
// Solidity: function buyMarketingFee() view returns(uint256)
func (_Token *TokenSession) BuyMarketingFee() (*big.Int, error) {
	return _Token.Contract.BuyMarketingFee(&_Token.CallOpts)
}

// BuyMarketingFee is a free data retrieval call binding the contract method 0x7bce5a04.
//
// Solidity: function buyMarketingFee() view returns(uint256)
func (_Token *TokenCallerSession) BuyMarketingFee() (*big.Int, error) {
	return _Token.Contract.BuyMarketingFee(&_Token.CallOpts)
}

// BuyTokenRewardsFee is a free data retrieval call binding the contract method 0x02df64d6.
//
// Solidity: function buyTokenRewardsFee() view returns(uint256)
func (_Token *TokenCaller) BuyTokenRewardsFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "buyTokenRewardsFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyTokenRewardsFee is a free data retrieval call binding the contract method 0x02df64d6.
//
// Solidity: function buyTokenRewardsFee() view returns(uint256)
func (_Token *TokenSession) BuyTokenRewardsFee() (*big.Int, error) {
	return _Token.Contract.BuyTokenRewardsFee(&_Token.CallOpts)
}

// BuyTokenRewardsFee is a free data retrieval call binding the contract method 0x02df64d6.
//
// Solidity: function buyTokenRewardsFee() view returns(uint256)
func (_Token *TokenCallerSession) BuyTokenRewardsFee() (*big.Int, error) {
	return _Token.Contract.BuyTokenRewardsFee(&_Token.CallOpts)
}

// DeadWallet is a free data retrieval call binding the contract method 0x85141a77.
//
// Solidity: function deadWallet() view returns(address)
func (_Token *TokenCaller) DeadWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "deadWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeadWallet is a free data retrieval call binding the contract method 0x85141a77.
//
// Solidity: function deadWallet() view returns(address)
func (_Token *TokenSession) DeadWallet() (common.Address, error) {
	return _Token.Contract.DeadWallet(&_Token.CallOpts)
}

// DeadWallet is a free data retrieval call binding the contract method 0x85141a77.
//
// Solidity: function deadWallet() view returns(address)
func (_Token *TokenCallerSession) DeadWallet() (common.Address, error) {
	return _Token.Contract.DeadWallet(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// DividendTokenBalanceOf is a free data retrieval call binding the contract method 0x6843cd84.
//
// Solidity: function dividendTokenBalanceOf(address account) view returns(uint256)
func (_Token *TokenCaller) DividendTokenBalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "dividendTokenBalanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DividendTokenBalanceOf is a free data retrieval call binding the contract method 0x6843cd84.
//
// Solidity: function dividendTokenBalanceOf(address account) view returns(uint256)
func (_Token *TokenSession) DividendTokenBalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.DividendTokenBalanceOf(&_Token.CallOpts, account)
}

// DividendTokenBalanceOf is a free data retrieval call binding the contract method 0x6843cd84.
//
// Solidity: function dividendTokenBalanceOf(address account) view returns(uint256)
func (_Token *TokenCallerSession) DividendTokenBalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.DividendTokenBalanceOf(&_Token.CallOpts, account)
}

// DividendTracker is a free data retrieval call binding the contract method 0x2c1f5216.
//
// Solidity: function dividendTracker() view returns(address)
func (_Token *TokenCaller) DividendTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "dividendTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DividendTracker is a free data retrieval call binding the contract method 0x2c1f5216.
//
// Solidity: function dividendTracker() view returns(address)
func (_Token *TokenSession) DividendTracker() (common.Address, error) {
	return _Token.Contract.DividendTracker(&_Token.CallOpts)
}

// DividendTracker is a free data retrieval call binding the contract method 0x2c1f5216.
//
// Solidity: function dividendTracker() view returns(address)
func (_Token *TokenCallerSession) DividendTracker() (common.Address, error) {
	return _Token.Contract.DividendTracker(&_Token.CallOpts)
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_Token *TokenCaller) First(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "first")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_Token *TokenSession) First() (*big.Int, error) {
	return _Token.Contract.First(&_Token.CallOpts)
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_Token *TokenCallerSession) First() (*big.Int, error) {
	return _Token.Contract.First(&_Token.CallOpts)
}

// GasForProcessing is a free data retrieval call binding the contract method 0x9c1b8af5.
//
// Solidity: function gasForProcessing() view returns(uint256)
func (_Token *TokenCaller) GasForProcessing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "gasForProcessing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasForProcessing is a free data retrieval call binding the contract method 0x9c1b8af5.
//
// Solidity: function gasForProcessing() view returns(uint256)
func (_Token *TokenSession) GasForProcessing() (*big.Int, error) {
	return _Token.Contract.GasForProcessing(&_Token.CallOpts)
}

// GasForProcessing is a free data retrieval call binding the contract method 0x9c1b8af5.
//
// Solidity: function gasForProcessing() view returns(uint256)
func (_Token *TokenCallerSession) GasForProcessing() (*big.Int, error) {
	return _Token.Contract.GasForProcessing(&_Token.CallOpts)
}

// GetAccountDividendsInfo is a free data retrieval call binding the contract method 0xad56c13c.
//
// Solidity: function getAccountDividendsInfo(address account) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Token *TokenCaller) GetAccountDividendsInfo(opts *bind.CallOpts, account common.Address) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getAccountDividendsInfo", account)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetAccountDividendsInfo is a free data retrieval call binding the contract method 0xad56c13c.
//
// Solidity: function getAccountDividendsInfo(address account) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Token *TokenSession) GetAccountDividendsInfo(account common.Address) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.GetAccountDividendsInfo(&_Token.CallOpts, account)
}

// GetAccountDividendsInfo is a free data retrieval call binding the contract method 0xad56c13c.
//
// Solidity: function getAccountDividendsInfo(address account) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Token *TokenCallerSession) GetAccountDividendsInfo(account common.Address) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.GetAccountDividendsInfo(&_Token.CallOpts, account)
}

// GetAccountDividendsInfoAtIndex is a free data retrieval call binding the contract method 0xf27fd254.
//
// Solidity: function getAccountDividendsInfoAtIndex(uint256 index) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Token *TokenCaller) GetAccountDividendsInfoAtIndex(opts *bind.CallOpts, index *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getAccountDividendsInfoAtIndex", index)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetAccountDividendsInfoAtIndex is a free data retrieval call binding the contract method 0xf27fd254.
//
// Solidity: function getAccountDividendsInfoAtIndex(uint256 index) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Token *TokenSession) GetAccountDividendsInfoAtIndex(index *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.GetAccountDividendsInfoAtIndex(&_Token.CallOpts, index)
}

// GetAccountDividendsInfoAtIndex is a free data retrieval call binding the contract method 0xf27fd254.
//
// Solidity: function getAccountDividendsInfoAtIndex(uint256 index) view returns(address, int256, int256, uint256, uint256, uint256, uint256, uint256)
func (_Token *TokenCallerSession) GetAccountDividendsInfoAtIndex(index *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Token.Contract.GetAccountDividendsInfoAtIndex(&_Token.CallOpts, index)
}

// GetClaimWait is a free data retrieval call binding the contract method 0xa26579ad.
//
// Solidity: function getClaimWait() view returns(uint256)
func (_Token *TokenCaller) GetClaimWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getClaimWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimWait is a free data retrieval call binding the contract method 0xa26579ad.
//
// Solidity: function getClaimWait() view returns(uint256)
func (_Token *TokenSession) GetClaimWait() (*big.Int, error) {
	return _Token.Contract.GetClaimWait(&_Token.CallOpts)
}

// GetClaimWait is a free data retrieval call binding the contract method 0xa26579ad.
//
// Solidity: function getClaimWait() view returns(uint256)
func (_Token *TokenCallerSession) GetClaimWait() (*big.Int, error) {
	return _Token.Contract.GetClaimWait(&_Token.CallOpts)
}

// GetLastProcessedIndex is a free data retrieval call binding the contract method 0xe7841ec0.
//
// Solidity: function getLastProcessedIndex() view returns(uint256)
func (_Token *TokenCaller) GetLastProcessedIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getLastProcessedIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastProcessedIndex is a free data retrieval call binding the contract method 0xe7841ec0.
//
// Solidity: function getLastProcessedIndex() view returns(uint256)
func (_Token *TokenSession) GetLastProcessedIndex() (*big.Int, error) {
	return _Token.Contract.GetLastProcessedIndex(&_Token.CallOpts)
}

// GetLastProcessedIndex is a free data retrieval call binding the contract method 0xe7841ec0.
//
// Solidity: function getLastProcessedIndex() view returns(uint256)
func (_Token *TokenCallerSession) GetLastProcessedIndex() (*big.Int, error) {
	return _Token.Contract.GetLastProcessedIndex(&_Token.CallOpts)
}

// GetMinimumTokenBalanceForDividends is a free data retrieval call binding the contract method 0xbdd4f29f.
//
// Solidity: function getMinimumTokenBalanceForDividends() view returns(uint256)
func (_Token *TokenCaller) GetMinimumTokenBalanceForDividends(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getMinimumTokenBalanceForDividends")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumTokenBalanceForDividends is a free data retrieval call binding the contract method 0xbdd4f29f.
//
// Solidity: function getMinimumTokenBalanceForDividends() view returns(uint256)
func (_Token *TokenSession) GetMinimumTokenBalanceForDividends() (*big.Int, error) {
	return _Token.Contract.GetMinimumTokenBalanceForDividends(&_Token.CallOpts)
}

// GetMinimumTokenBalanceForDividends is a free data retrieval call binding the contract method 0xbdd4f29f.
//
// Solidity: function getMinimumTokenBalanceForDividends() view returns(uint256)
func (_Token *TokenCallerSession) GetMinimumTokenBalanceForDividends() (*big.Int, error) {
	return _Token.Contract.GetMinimumTokenBalanceForDividends(&_Token.CallOpts)
}

// GetNumberOfDividendTokenHolders is a free data retrieval call binding the contract method 0x64b0f653.
//
// Solidity: function getNumberOfDividendTokenHolders() view returns(uint256)
func (_Token *TokenCaller) GetNumberOfDividendTokenHolders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getNumberOfDividendTokenHolders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfDividendTokenHolders is a free data retrieval call binding the contract method 0x64b0f653.
//
// Solidity: function getNumberOfDividendTokenHolders() view returns(uint256)
func (_Token *TokenSession) GetNumberOfDividendTokenHolders() (*big.Int, error) {
	return _Token.Contract.GetNumberOfDividendTokenHolders(&_Token.CallOpts)
}

// GetNumberOfDividendTokenHolders is a free data retrieval call binding the contract method 0x64b0f653.
//
// Solidity: function getNumberOfDividendTokenHolders() view returns(uint256)
func (_Token *TokenCallerSession) GetNumberOfDividendTokenHolders() (*big.Int, error) {
	return _Token.Contract.GetNumberOfDividendTokenHolders(&_Token.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_Token *TokenCaller) GetTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_Token *TokenSession) GetTime() (*big.Int, error) {
	return _Token.Contract.GetTime(&_Token.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_Token *TokenCallerSession) GetTime() (*big.Int, error) {
	return _Token.Contract.GetTime(&_Token.CallOpts)
}

// GetTotalDividendsDistributed is a free data retrieval call binding the contract method 0x30bb4cff.
//
// Solidity: function getTotalDividendsDistributed() view returns(uint256)
func (_Token *TokenCaller) GetTotalDividendsDistributed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getTotalDividendsDistributed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDividendsDistributed is a free data retrieval call binding the contract method 0x30bb4cff.
//
// Solidity: function getTotalDividendsDistributed() view returns(uint256)
func (_Token *TokenSession) GetTotalDividendsDistributed() (*big.Int, error) {
	return _Token.Contract.GetTotalDividendsDistributed(&_Token.CallOpts)
}

// GetTotalDividendsDistributed is a free data retrieval call binding the contract method 0x30bb4cff.
//
// Solidity: function getTotalDividendsDistributed() view returns(uint256)
func (_Token *TokenCallerSession) GetTotalDividendsDistributed() (*big.Int, error) {
	return _Token.Contract.GetTotalDividendsDistributed(&_Token.CallOpts)
}

// IsExcludedFromDividends is a free data retrieval call binding the contract method 0xc705c569.
//
// Solidity: function isExcludedFromDividends(address account) view returns(bool)
func (_Token *TokenCaller) IsExcludedFromDividends(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isExcludedFromDividends", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromDividends is a free data retrieval call binding the contract method 0xc705c569.
//
// Solidity: function isExcludedFromDividends(address account) view returns(bool)
func (_Token *TokenSession) IsExcludedFromDividends(account common.Address) (bool, error) {
	return _Token.Contract.IsExcludedFromDividends(&_Token.CallOpts, account)
}

// IsExcludedFromDividends is a free data retrieval call binding the contract method 0xc705c569.
//
// Solidity: function isExcludedFromDividends(address account) view returns(bool)
func (_Token *TokenCallerSession) IsExcludedFromDividends(account common.Address) (bool, error) {
	return _Token.Contract.IsExcludedFromDividends(&_Token.CallOpts, account)
}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Token *TokenCaller) IsExcludedFromFees(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isExcludedFromFees", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Token *TokenSession) IsExcludedFromFees(account common.Address) (bool, error) {
	return _Token.Contract.IsExcludedFromFees(&_Token.CallOpts, account)
}

// IsExcludedFromFees is a free data retrieval call binding the contract method 0x4fbee193.
//
// Solidity: function isExcludedFromFees(address account) view returns(bool)
func (_Token *TokenCallerSession) IsExcludedFromFees(account common.Address) (bool, error) {
	return _Token.Contract.IsExcludedFromFees(&_Token.CallOpts, account)
}

// Kill is a free data retrieval call binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() view returns(uint256)
func (_Token *TokenCaller) Kill(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "kill")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Kill is a free data retrieval call binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() view returns(uint256)
func (_Token *TokenSession) Kill() (*big.Int, error) {
	return _Token.Contract.Kill(&_Token.CallOpts)
}

// Kill is a free data retrieval call binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() view returns(uint256)
func (_Token *TokenCallerSession) Kill() (*big.Int, error) {
	return _Token.Contract.Kill(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Token *TokenCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Token *TokenSession) RewardToken() (common.Address, error) {
	return _Token.Contract.RewardToken(&_Token.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Token *TokenCallerSession) RewardToken() (common.Address, error) {
	return _Token.Contract.RewardToken(&_Token.CallOpts)
}

// SellDeadFee is a free data retrieval call binding the contract method 0x24942a04.
//
// Solidity: function sellDeadFee() view returns(uint256)
func (_Token *TokenCaller) SellDeadFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "sellDeadFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellDeadFee is a free data retrieval call binding the contract method 0x24942a04.
//
// Solidity: function sellDeadFee() view returns(uint256)
func (_Token *TokenSession) SellDeadFee() (*big.Int, error) {
	return _Token.Contract.SellDeadFee(&_Token.CallOpts)
}

// SellDeadFee is a free data retrieval call binding the contract method 0x24942a04.
//
// Solidity: function sellDeadFee() view returns(uint256)
func (_Token *TokenCallerSession) SellDeadFee() (*big.Int, error) {
	return _Token.Contract.SellDeadFee(&_Token.CallOpts)
}

// SellLiquidityFee is a free data retrieval call binding the contract method 0xf6374342.
//
// Solidity: function sellLiquidityFee() view returns(uint256)
func (_Token *TokenCaller) SellLiquidityFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "sellLiquidityFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellLiquidityFee is a free data retrieval call binding the contract method 0xf6374342.
//
// Solidity: function sellLiquidityFee() view returns(uint256)
func (_Token *TokenSession) SellLiquidityFee() (*big.Int, error) {
	return _Token.Contract.SellLiquidityFee(&_Token.CallOpts)
}

// SellLiquidityFee is a free data retrieval call binding the contract method 0xf6374342.
//
// Solidity: function sellLiquidityFee() view returns(uint256)
func (_Token *TokenCallerSession) SellLiquidityFee() (*big.Int, error) {
	return _Token.Contract.SellLiquidityFee(&_Token.CallOpts)
}

// SellMarketingFee is a free data retrieval call binding the contract method 0x92136913.
//
// Solidity: function sellMarketingFee() view returns(uint256)
func (_Token *TokenCaller) SellMarketingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "sellMarketingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellMarketingFee is a free data retrieval call binding the contract method 0x92136913.
//
// Solidity: function sellMarketingFee() view returns(uint256)
func (_Token *TokenSession) SellMarketingFee() (*big.Int, error) {
	return _Token.Contract.SellMarketingFee(&_Token.CallOpts)
}

// SellMarketingFee is a free data retrieval call binding the contract method 0x92136913.
//
// Solidity: function sellMarketingFee() view returns(uint256)
func (_Token *TokenCallerSession) SellMarketingFee() (*big.Int, error) {
	return _Token.Contract.SellMarketingFee(&_Token.CallOpts)
}

// SellTokenRewardsFee is a free data retrieval call binding the contract method 0x08b2a12c.
//
// Solidity: function sellTokenRewardsFee() view returns(uint256)
func (_Token *TokenCaller) SellTokenRewardsFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "sellTokenRewardsFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellTokenRewardsFee is a free data retrieval call binding the contract method 0x08b2a12c.
//
// Solidity: function sellTokenRewardsFee() view returns(uint256)
func (_Token *TokenSession) SellTokenRewardsFee() (*big.Int, error) {
	return _Token.Contract.SellTokenRewardsFee(&_Token.CallOpts)
}

// SellTokenRewardsFee is a free data retrieval call binding the contract method 0x08b2a12c.
//
// Solidity: function sellTokenRewardsFee() view returns(uint256)
func (_Token *TokenCallerSession) SellTokenRewardsFee() (*big.Int, error) {
	return _Token.Contract.SellTokenRewardsFee(&_Token.CallOpts)
}

// SwapAndLiquifyEnabled is a free data retrieval call binding the contract method 0x4a74bb02.
//
// Solidity: function swapAndLiquifyEnabled() view returns(bool)
func (_Token *TokenCaller) SwapAndLiquifyEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "swapAndLiquifyEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapAndLiquifyEnabled is a free data retrieval call binding the contract method 0x4a74bb02.
//
// Solidity: function swapAndLiquifyEnabled() view returns(bool)
func (_Token *TokenSession) SwapAndLiquifyEnabled() (bool, error) {
	return _Token.Contract.SwapAndLiquifyEnabled(&_Token.CallOpts)
}

// SwapAndLiquifyEnabled is a free data retrieval call binding the contract method 0x4a74bb02.
//
// Solidity: function swapAndLiquifyEnabled() view returns(bool)
func (_Token *TokenCallerSession) SwapAndLiquifyEnabled() (bool, error) {
	return _Token.Contract.SwapAndLiquifyEnabled(&_Token.CallOpts)
}

// SwapTokensAtAmount is a free data retrieval call binding the contract method 0xe2f45605.
//
// Solidity: function swapTokensAtAmount() view returns(uint256)
func (_Token *TokenCaller) SwapTokensAtAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "swapTokensAtAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapTokensAtAmount is a free data retrieval call binding the contract method 0xe2f45605.
//
// Solidity: function swapTokensAtAmount() view returns(uint256)
func (_Token *TokenSession) SwapTokensAtAmount() (*big.Int, error) {
	return _Token.Contract.SwapTokensAtAmount(&_Token.CallOpts)
}

// SwapTokensAtAmount is a free data retrieval call binding the contract method 0xe2f45605.
//
// Solidity: function swapTokensAtAmount() view returns(uint256)
func (_Token *TokenCallerSession) SwapTokensAtAmount() (*big.Int, error) {
	return _Token.Contract.SwapTokensAtAmount(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// UniswapPair is a free data retrieval call binding the contract method 0xc816841b.
//
// Solidity: function uniswapPair() view returns(address)
func (_Token *TokenCaller) UniswapPair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "uniswapPair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapPair is a free data retrieval call binding the contract method 0xc816841b.
//
// Solidity: function uniswapPair() view returns(address)
func (_Token *TokenSession) UniswapPair() (common.Address, error) {
	return _Token.Contract.UniswapPair(&_Token.CallOpts)
}

// UniswapPair is a free data retrieval call binding the contract method 0xc816841b.
//
// Solidity: function uniswapPair() view returns(address)
func (_Token *TokenCallerSession) UniswapPair() (common.Address, error) {
	return _Token.Contract.UniswapPair(&_Token.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Token *TokenCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Token *TokenSession) UniswapV2Router() (common.Address, error) {
	return _Token.Contract.UniswapV2Router(&_Token.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Token *TokenCallerSession) UniswapV2Router() (common.Address, error) {
	return _Token.Contract.UniswapV2Router(&_Token.CallOpts)
}

// WithdrawableDividendOf is a free data retrieval call binding the contract method 0xa8b9d240.
//
// Solidity: function withdrawableDividendOf(address account) view returns(uint256)
func (_Token *TokenCaller) WithdrawableDividendOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "withdrawableDividendOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableDividendOf is a free data retrieval call binding the contract method 0xa8b9d240.
//
// Solidity: function withdrawableDividendOf(address account) view returns(uint256)
func (_Token *TokenSession) WithdrawableDividendOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.WithdrawableDividendOf(&_Token.CallOpts, account)
}

// WithdrawableDividendOf is a free data retrieval call binding the contract method 0xa8b9d240.
//
// Solidity: function withdrawableDividendOf(address account) view returns(uint256)
func (_Token *TokenCallerSession) WithdrawableDividendOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.WithdrawableDividendOf(&_Token.CallOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Token *TokenTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Token *TokenSession) Claim() (*types.Transaction, error) {
	return _Token.Contract.Claim(&_Token.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Token *TokenTransactorSession) Claim() (*types.Transaction, error) {
	return _Token.Contract.Claim(&_Token.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, spender, subtractedValue)
}

// ExcludeFromDividends is a paid mutator transaction binding the contract method 0x31e79db0.
//
// Solidity: function excludeFromDividends(address account) returns()
func (_Token *TokenTransactor) ExcludeFromDividends(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "excludeFromDividends", account)
}

// ExcludeFromDividends is a paid mutator transaction binding the contract method 0x31e79db0.
//
// Solidity: function excludeFromDividends(address account) returns()
func (_Token *TokenSession) ExcludeFromDividends(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.ExcludeFromDividends(&_Token.TransactOpts, account)
}

// ExcludeFromDividends is a paid mutator transaction binding the contract method 0x31e79db0.
//
// Solidity: function excludeFromDividends(address account) returns()
func (_Token *TokenTransactorSession) ExcludeFromDividends(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.ExcludeFromDividends(&_Token.TransactOpts, account)
}

// ExcludeFromFees is a paid mutator transaction binding the contract method 0xc0246668.
//
// Solidity: function excludeFromFees(address account, bool excluded) returns()
func (_Token *TokenTransactor) ExcludeFromFees(opts *bind.TransactOpts, account common.Address, excluded bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "excludeFromFees", account, excluded)
}

// ExcludeFromFees is a paid mutator transaction binding the contract method 0xc0246668.
//
// Solidity: function excludeFromFees(address account, bool excluded) returns()
func (_Token *TokenSession) ExcludeFromFees(account common.Address, excluded bool) (*types.Transaction, error) {
	return _Token.Contract.ExcludeFromFees(&_Token.TransactOpts, account, excluded)
}

// ExcludeFromFees is a paid mutator transaction binding the contract method 0xc0246668.
//
// Solidity: function excludeFromFees(address account, bool excluded) returns()
func (_Token *TokenTransactorSession) ExcludeFromFees(account common.Address, excluded bool) (*types.Transaction, error) {
	return _Token.Contract.ExcludeFromFees(&_Token.TransactOpts, account, excluded)
}

// ExcludeMultipleAccountsFromFees is a paid mutator transaction binding the contract method 0xc492f046.
//
// Solidity: function excludeMultipleAccountsFromFees(address[] accounts, bool excluded) returns()
func (_Token *TokenTransactor) ExcludeMultipleAccountsFromFees(opts *bind.TransactOpts, accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "excludeMultipleAccountsFromFees", accounts, excluded)
}

// ExcludeMultipleAccountsFromFees is a paid mutator transaction binding the contract method 0xc492f046.
//
// Solidity: function excludeMultipleAccountsFromFees(address[] accounts, bool excluded) returns()
func (_Token *TokenSession) ExcludeMultipleAccountsFromFees(accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Token.Contract.ExcludeMultipleAccountsFromFees(&_Token.TransactOpts, accounts, excluded)
}

// ExcludeMultipleAccountsFromFees is a paid mutator transaction binding the contract method 0xc492f046.
//
// Solidity: function excludeMultipleAccountsFromFees(address[] accounts, bool excluded) returns()
func (_Token *TokenTransactorSession) ExcludeMultipleAccountsFromFees(accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Token.Contract.ExcludeMultipleAccountsFromFees(&_Token.TransactOpts, accounts, excluded)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, spender, addedValue)
}

// MultipleBotlistAddress is a paid mutator transaction binding the contract method 0x9a55fff0.
//
// Solidity: function multipleBotlistAddress(address[] accounts, bool excluded) returns()
func (_Token *TokenTransactor) MultipleBotlistAddress(opts *bind.TransactOpts, accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "multipleBotlistAddress", accounts, excluded)
}

// MultipleBotlistAddress is a paid mutator transaction binding the contract method 0x9a55fff0.
//
// Solidity: function multipleBotlistAddress(address[] accounts, bool excluded) returns()
func (_Token *TokenSession) MultipleBotlistAddress(accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Token.Contract.MultipleBotlistAddress(&_Token.TransactOpts, accounts, excluded)
}

// MultipleBotlistAddress is a paid mutator transaction binding the contract method 0x9a55fff0.
//
// Solidity: function multipleBotlistAddress(address[] accounts, bool excluded) returns()
func (_Token *TokenTransactorSession) MultipleBotlistAddress(accounts []common.Address, excluded bool) (*types.Transaction, error) {
	return _Token.Contract.MultipleBotlistAddress(&_Token.TransactOpts, accounts, excluded)
}

// ProcessDividendTracker is a paid mutator transaction binding the contract method 0x700bb191.
//
// Solidity: function processDividendTracker(uint256 gas) returns()
func (_Token *TokenTransactor) ProcessDividendTracker(opts *bind.TransactOpts, gas *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "processDividendTracker", gas)
}

// ProcessDividendTracker is a paid mutator transaction binding the contract method 0x700bb191.
//
// Solidity: function processDividendTracker(uint256 gas) returns()
func (_Token *TokenSession) ProcessDividendTracker(gas *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ProcessDividendTracker(&_Token.TransactOpts, gas)
}

// ProcessDividendTracker is a paid mutator transaction binding the contract method 0x700bb191.
//
// Solidity: function processDividendTracker(uint256 gas) returns()
func (_Token *TokenTransactorSession) ProcessDividendTracker(gas *big.Int) (*types.Transaction, error) {
	return _Token.Contract.ProcessDividendTracker(&_Token.TransactOpts, gas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Token *TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Token.Contract.RenounceOwnership(&_Token.TransactOpts)
}

// SetAirdropNumbs is a paid mutator transaction binding the contract method 0xe5c2b0a6.
//
// Solidity: function setAirdropNumbs(uint256 newValue) returns()
func (_Token *TokenTransactor) SetAirdropNumbs(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setAirdropNumbs", newValue)
}

// SetAirdropNumbs is a paid mutator transaction binding the contract method 0xe5c2b0a6.
//
// Solidity: function setAirdropNumbs(uint256 newValue) returns()
func (_Token *TokenSession) SetAirdropNumbs(newValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetAirdropNumbs(&_Token.TransactOpts, newValue)
}

// SetAirdropNumbs is a paid mutator transaction binding the contract method 0xe5c2b0a6.
//
// Solidity: function setAirdropNumbs(uint256 newValue) returns()
func (_Token *TokenTransactorSession) SetAirdropNumbs(newValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetAirdropNumbs(&_Token.TransactOpts, newValue)
}

// SetAutomatedMarketMakerPair is a paid mutator transaction binding the contract method 0x9a7a23d6.
//
// Solidity: function setAutomatedMarketMakerPair(address pair, bool value) returns()
func (_Token *TokenTransactor) SetAutomatedMarketMakerPair(opts *bind.TransactOpts, pair common.Address, value bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setAutomatedMarketMakerPair", pair, value)
}

// SetAutomatedMarketMakerPair is a paid mutator transaction binding the contract method 0x9a7a23d6.
//
// Solidity: function setAutomatedMarketMakerPair(address pair, bool value) returns()
func (_Token *TokenSession) SetAutomatedMarketMakerPair(pair common.Address, value bool) (*types.Transaction, error) {
	return _Token.Contract.SetAutomatedMarketMakerPair(&_Token.TransactOpts, pair, value)
}

// SetAutomatedMarketMakerPair is a paid mutator transaction binding the contract method 0x9a7a23d6.
//
// Solidity: function setAutomatedMarketMakerPair(address pair, bool value) returns()
func (_Token *TokenTransactorSession) SetAutomatedMarketMakerPair(pair common.Address, value bool) (*types.Transaction, error) {
	return _Token.Contract.SetAutomatedMarketMakerPair(&_Token.TransactOpts, pair, value)
}

// SetBuyTaxes is a paid mutator transaction binding the contract method 0xe4bf1bed.
//
// Solidity: function setBuyTaxes(uint256 liquidity, uint256 rewardsFee, uint256 marketingFee, uint256 deadFee) returns()
func (_Token *TokenTransactor) SetBuyTaxes(opts *bind.TransactOpts, liquidity *big.Int, rewardsFee *big.Int, marketingFee *big.Int, deadFee *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setBuyTaxes", liquidity, rewardsFee, marketingFee, deadFee)
}

// SetBuyTaxes is a paid mutator transaction binding the contract method 0xe4bf1bed.
//
// Solidity: function setBuyTaxes(uint256 liquidity, uint256 rewardsFee, uint256 marketingFee, uint256 deadFee) returns()
func (_Token *TokenSession) SetBuyTaxes(liquidity *big.Int, rewardsFee *big.Int, marketingFee *big.Int, deadFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetBuyTaxes(&_Token.TransactOpts, liquidity, rewardsFee, marketingFee, deadFee)
}

// SetBuyTaxes is a paid mutator transaction binding the contract method 0xe4bf1bed.
//
// Solidity: function setBuyTaxes(uint256 liquidity, uint256 rewardsFee, uint256 marketingFee, uint256 deadFee) returns()
func (_Token *TokenTransactorSession) SetBuyTaxes(liquidity *big.Int, rewardsFee *big.Int, marketingFee *big.Int, deadFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetBuyTaxes(&_Token.TransactOpts, liquidity, rewardsFee, marketingFee, deadFee)
}

// SetDeadWallet is a paid mutator transaction binding the contract method 0x61a60d57.
//
// Solidity: function setDeadWallet(address addr) returns()
func (_Token *TokenTransactor) SetDeadWallet(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setDeadWallet", addr)
}

// SetDeadWallet is a paid mutator transaction binding the contract method 0x61a60d57.
//
// Solidity: function setDeadWallet(address addr) returns()
func (_Token *TokenSession) SetDeadWallet(addr common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetDeadWallet(&_Token.TransactOpts, addr)
}

// SetDeadWallet is a paid mutator transaction binding the contract method 0x61a60d57.
//
// Solidity: function setDeadWallet(address addr) returns()
func (_Token *TokenTransactorSession) SetDeadWallet(addr common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetDeadWallet(&_Token.TransactOpts, addr)
}

// SetKing is a paid mutator transaction binding the contract method 0xc7f063d3.
//
// Solidity: function setKing(uint256 newValue) returns()
func (_Token *TokenTransactor) SetKing(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setKing", newValue)
}

// SetKing is a paid mutator transaction binding the contract method 0xc7f063d3.
//
// Solidity: function setKing(uint256 newValue) returns()
func (_Token *TokenSession) SetKing(newValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetKing(&_Token.TransactOpts, newValue)
}

// SetKing is a paid mutator transaction binding the contract method 0xc7f063d3.
//
// Solidity: function setKing(uint256 newValue) returns()
func (_Token *TokenTransactorSession) SetKing(newValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetKing(&_Token.TransactOpts, newValue)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address wallet) returns()
func (_Token *TokenTransactor) SetMarketingWallet(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setMarketingWallet", wallet)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address wallet) returns()
func (_Token *TokenSession) SetMarketingWallet(wallet common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetMarketingWallet(&_Token.TransactOpts, wallet)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address wallet) returns()
func (_Token *TokenTransactorSession) SetMarketingWallet(wallet common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetMarketingWallet(&_Token.TransactOpts, wallet)
}

// SetSelTaxes is a paid mutator transaction binding the contract method 0xeb671e91.
//
// Solidity: function setSelTaxes(uint256 liquidity, uint256 rewardsFee, uint256 marketingFee, uint256 deadFee) returns()
func (_Token *TokenTransactor) SetSelTaxes(opts *bind.TransactOpts, liquidity *big.Int, rewardsFee *big.Int, marketingFee *big.Int, deadFee *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setSelTaxes", liquidity, rewardsFee, marketingFee, deadFee)
}

// SetSelTaxes is a paid mutator transaction binding the contract method 0xeb671e91.
//
// Solidity: function setSelTaxes(uint256 liquidity, uint256 rewardsFee, uint256 marketingFee, uint256 deadFee) returns()
func (_Token *TokenSession) SetSelTaxes(liquidity *big.Int, rewardsFee *big.Int, marketingFee *big.Int, deadFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetSelTaxes(&_Token.TransactOpts, liquidity, rewardsFee, marketingFee, deadFee)
}

// SetSelTaxes is a paid mutator transaction binding the contract method 0xeb671e91.
//
// Solidity: function setSelTaxes(uint256 liquidity, uint256 rewardsFee, uint256 marketingFee, uint256 deadFee) returns()
func (_Token *TokenTransactorSession) SetSelTaxes(liquidity *big.Int, rewardsFee *big.Int, marketingFee *big.Int, deadFee *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetSelTaxes(&_Token.TransactOpts, liquidity, rewardsFee, marketingFee, deadFee)
}

// SetSwapAndLiquifyEnabled is a paid mutator transaction binding the contract method 0xc49b9a80.
//
// Solidity: function setSwapAndLiquifyEnabled(bool _enabled) returns()
func (_Token *TokenTransactor) SetSwapAndLiquifyEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setSwapAndLiquifyEnabled", _enabled)
}

// SetSwapAndLiquifyEnabled is a paid mutator transaction binding the contract method 0xc49b9a80.
//
// Solidity: function setSwapAndLiquifyEnabled(bool _enabled) returns()
func (_Token *TokenSession) SetSwapAndLiquifyEnabled(_enabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetSwapAndLiquifyEnabled(&_Token.TransactOpts, _enabled)
}

// SetSwapAndLiquifyEnabled is a paid mutator transaction binding the contract method 0xc49b9a80.
//
// Solidity: function setSwapAndLiquifyEnabled(bool _enabled) returns()
func (_Token *TokenTransactorSession) SetSwapAndLiquifyEnabled(_enabled bool) (*types.Transaction, error) {
	return _Token.Contract.SetSwapAndLiquifyEnabled(&_Token.TransactOpts, _enabled)
}

// SetSwapTokensAtAmount is a paid mutator transaction binding the contract method 0xafa4f3b2.
//
// Solidity: function setSwapTokensAtAmount(uint256 amount) returns()
func (_Token *TokenTransactor) SetSwapTokensAtAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setSwapTokensAtAmount", amount)
}

// SetSwapTokensAtAmount is a paid mutator transaction binding the contract method 0xafa4f3b2.
//
// Solidity: function setSwapTokensAtAmount(uint256 amount) returns()
func (_Token *TokenSession) SetSwapTokensAtAmount(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetSwapTokensAtAmount(&_Token.TransactOpts, amount)
}

// SetSwapTokensAtAmount is a paid mutator transaction binding the contract method 0xafa4f3b2.
//
// Solidity: function setSwapTokensAtAmount(uint256 amount) returns()
func (_Token *TokenTransactorSession) SetSwapTokensAtAmount(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetSwapTokensAtAmount(&_Token.TransactOpts, amount)
}

// SwapManual is a paid mutator transaction binding the contract method 0xd2a8b440.
//
// Solidity: function swapManual() returns()
func (_Token *TokenTransactor) SwapManual(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "swapManual")
}

// SwapManual is a paid mutator transaction binding the contract method 0xd2a8b440.
//
// Solidity: function swapManual() returns()
func (_Token *TokenSession) SwapManual() (*types.Transaction, error) {
	return _Token.Contract.SwapManual(&_Token.TransactOpts)
}

// SwapManual is a paid mutator transaction binding the contract method 0xd2a8b440.
//
// Solidity: function swapManual() returns()
func (_Token *TokenTransactorSession) SwapManual() (*types.Transaction, error) {
	return _Token.Contract.SwapManual(&_Token.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Token *TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Token.Contract.TransferOwnership(&_Token.TransactOpts, newOwner)
}

// UpdateClaimWait is a paid mutator transaction binding the contract method 0xe98030c7.
//
// Solidity: function updateClaimWait(uint256 claimWait) returns()
func (_Token *TokenTransactor) UpdateClaimWait(opts *bind.TransactOpts, claimWait *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateClaimWait", claimWait)
}

// UpdateClaimWait is a paid mutator transaction binding the contract method 0xe98030c7.
//
// Solidity: function updateClaimWait(uint256 claimWait) returns()
func (_Token *TokenSession) UpdateClaimWait(claimWait *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UpdateClaimWait(&_Token.TransactOpts, claimWait)
}

// UpdateClaimWait is a paid mutator transaction binding the contract method 0xe98030c7.
//
// Solidity: function updateClaimWait(uint256 claimWait) returns()
func (_Token *TokenTransactorSession) UpdateClaimWait(claimWait *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UpdateClaimWait(&_Token.TransactOpts, claimWait)
}

// UpdateGasForProcessing is a paid mutator transaction binding the contract method 0x871c128d.
//
// Solidity: function updateGasForProcessing(uint256 newValue) returns()
func (_Token *TokenTransactor) UpdateGasForProcessing(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateGasForProcessing", newValue)
}

// UpdateGasForProcessing is a paid mutator transaction binding the contract method 0x871c128d.
//
// Solidity: function updateGasForProcessing(uint256 newValue) returns()
func (_Token *TokenSession) UpdateGasForProcessing(newValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UpdateGasForProcessing(&_Token.TransactOpts, newValue)
}

// UpdateGasForProcessing is a paid mutator transaction binding the contract method 0x871c128d.
//
// Solidity: function updateGasForProcessing(uint256 newValue) returns()
func (_Token *TokenTransactorSession) UpdateGasForProcessing(newValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UpdateGasForProcessing(&_Token.TransactOpts, newValue)
}

// UpdateMinimumTokenBalanceForDividends is a paid mutator transaction binding the contract method 0x0dcb2e89.
//
// Solidity: function updateMinimumTokenBalanceForDividends(uint256 val) returns()
func (_Token *TokenTransactor) UpdateMinimumTokenBalanceForDividends(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateMinimumTokenBalanceForDividends", val)
}

// UpdateMinimumTokenBalanceForDividends is a paid mutator transaction binding the contract method 0x0dcb2e89.
//
// Solidity: function updateMinimumTokenBalanceForDividends(uint256 val) returns()
func (_Token *TokenSession) UpdateMinimumTokenBalanceForDividends(val *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UpdateMinimumTokenBalanceForDividends(&_Token.TransactOpts, val)
}

// UpdateMinimumTokenBalanceForDividends is a paid mutator transaction binding the contract method 0x0dcb2e89.
//
// Solidity: function updateMinimumTokenBalanceForDividends(uint256 val) returns()
func (_Token *TokenTransactorSession) UpdateMinimumTokenBalanceForDividends(val *big.Int) (*types.Transaction, error) {
	return _Token.Contract.UpdateMinimumTokenBalanceForDividends(&_Token.TransactOpts, val)
}

// UpdateUniswapV2Router is a paid mutator transaction binding the contract method 0x65b8dbc0.
//
// Solidity: function updateUniswapV2Router(address newAddress) returns()
func (_Token *TokenTransactor) UpdateUniswapV2Router(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateUniswapV2Router", newAddress)
}

// UpdateUniswapV2Router is a paid mutator transaction binding the contract method 0x65b8dbc0.
//
// Solidity: function updateUniswapV2Router(address newAddress) returns()
func (_Token *TokenSession) UpdateUniswapV2Router(newAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.UpdateUniswapV2Router(&_Token.TransactOpts, newAddress)
}

// UpdateUniswapV2Router is a paid mutator transaction binding the contract method 0x65b8dbc0.
//
// Solidity: function updateUniswapV2Router(address newAddress) returns()
func (_Token *TokenTransactorSession) UpdateUniswapV2Router(newAddress common.Address) (*types.Transaction, error) {
	return _Token.Contract.UpdateUniswapV2Router(&_Token.TransactOpts, newAddress)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenSession) Receive() (*types.Transaction, error) {
	return _Token.Contract.Receive(&_Token.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Token *TokenTransactorSession) Receive() (*types.Transaction, error) {
	return _Token.Contract.Receive(&_Token.TransactOpts)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExcludeFromFeesIterator is returned from FilterExcludeFromFees and is used to iterate over the raw logs and unpacked data for ExcludeFromFees events raised by the Token contract.
type TokenExcludeFromFeesIterator struct {
	Event *TokenExcludeFromFees // Event containing the contract specifics and raw log

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
func (it *TokenExcludeFromFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExcludeFromFees)
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
		it.Event = new(TokenExcludeFromFees)
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
func (it *TokenExcludeFromFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExcludeFromFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExcludeFromFees represents a ExcludeFromFees event raised by the Token contract.
type TokenExcludeFromFees struct {
	Account    common.Address
	IsExcluded bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExcludeFromFees is a free log retrieval operation binding the contract event 0x9d8f7706ea1113d1a167b526eca956215946dd36cc7df39eb16180222d8b5df7.
//
// Solidity: event ExcludeFromFees(address indexed account, bool isExcluded)
func (_Token *TokenFilterer) FilterExcludeFromFees(opts *bind.FilterOpts, account []common.Address) (*TokenExcludeFromFeesIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ExcludeFromFees", accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenExcludeFromFeesIterator{contract: _Token.contract, event: "ExcludeFromFees", logs: logs, sub: sub}, nil
}

// WatchExcludeFromFees is a free log subscription operation binding the contract event 0x9d8f7706ea1113d1a167b526eca956215946dd36cc7df39eb16180222d8b5df7.
//
// Solidity: event ExcludeFromFees(address indexed account, bool isExcluded)
func (_Token *TokenFilterer) WatchExcludeFromFees(opts *bind.WatchOpts, sink chan<- *TokenExcludeFromFees, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ExcludeFromFees", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExcludeFromFees)
				if err := _Token.contract.UnpackLog(event, "ExcludeFromFees", log); err != nil {
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

// ParseExcludeFromFees is a log parse operation binding the contract event 0x9d8f7706ea1113d1a167b526eca956215946dd36cc7df39eb16180222d8b5df7.
//
// Solidity: event ExcludeFromFees(address indexed account, bool isExcluded)
func (_Token *TokenFilterer) ParseExcludeFromFees(log types.Log) (*TokenExcludeFromFees, error) {
	event := new(TokenExcludeFromFees)
	if err := _Token.contract.UnpackLog(event, "ExcludeFromFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenExcludeMultipleAccountsFromFeesIterator is returned from FilterExcludeMultipleAccountsFromFees and is used to iterate over the raw logs and unpacked data for ExcludeMultipleAccountsFromFees events raised by the Token contract.
type TokenExcludeMultipleAccountsFromFeesIterator struct {
	Event *TokenExcludeMultipleAccountsFromFees // Event containing the contract specifics and raw log

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
func (it *TokenExcludeMultipleAccountsFromFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenExcludeMultipleAccountsFromFees)
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
		it.Event = new(TokenExcludeMultipleAccountsFromFees)
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
func (it *TokenExcludeMultipleAccountsFromFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenExcludeMultipleAccountsFromFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenExcludeMultipleAccountsFromFees represents a ExcludeMultipleAccountsFromFees event raised by the Token contract.
type TokenExcludeMultipleAccountsFromFees struct {
	Accounts   []common.Address
	IsExcluded bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExcludeMultipleAccountsFromFees is a free log retrieval operation binding the contract event 0x7fdaf542373fa84f4ee8d662c642f44e4c2276a217d7d29e548b6eb29a233b35.
//
// Solidity: event ExcludeMultipleAccountsFromFees(address[] accounts, bool isExcluded)
func (_Token *TokenFilterer) FilterExcludeMultipleAccountsFromFees(opts *bind.FilterOpts) (*TokenExcludeMultipleAccountsFromFeesIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "ExcludeMultipleAccountsFromFees")
	if err != nil {
		return nil, err
	}
	return &TokenExcludeMultipleAccountsFromFeesIterator{contract: _Token.contract, event: "ExcludeMultipleAccountsFromFees", logs: logs, sub: sub}, nil
}

// WatchExcludeMultipleAccountsFromFees is a free log subscription operation binding the contract event 0x7fdaf542373fa84f4ee8d662c642f44e4c2276a217d7d29e548b6eb29a233b35.
//
// Solidity: event ExcludeMultipleAccountsFromFees(address[] accounts, bool isExcluded)
func (_Token *TokenFilterer) WatchExcludeMultipleAccountsFromFees(opts *bind.WatchOpts, sink chan<- *TokenExcludeMultipleAccountsFromFees) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "ExcludeMultipleAccountsFromFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenExcludeMultipleAccountsFromFees)
				if err := _Token.contract.UnpackLog(event, "ExcludeMultipleAccountsFromFees", log); err != nil {
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

// ParseExcludeMultipleAccountsFromFees is a log parse operation binding the contract event 0x7fdaf542373fa84f4ee8d662c642f44e4c2276a217d7d29e548b6eb29a233b35.
//
// Solidity: event ExcludeMultipleAccountsFromFees(address[] accounts, bool isExcluded)
func (_Token *TokenFilterer) ParseExcludeMultipleAccountsFromFees(log types.Log) (*TokenExcludeMultipleAccountsFromFees, error) {
	event := new(TokenExcludeMultipleAccountsFromFees)
	if err := _Token.contract.UnpackLog(event, "ExcludeMultipleAccountsFromFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenGasForProcessingUpdatedIterator is returned from FilterGasForProcessingUpdated and is used to iterate over the raw logs and unpacked data for GasForProcessingUpdated events raised by the Token contract.
type TokenGasForProcessingUpdatedIterator struct {
	Event *TokenGasForProcessingUpdated // Event containing the contract specifics and raw log

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
func (it *TokenGasForProcessingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGasForProcessingUpdated)
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
		it.Event = new(TokenGasForProcessingUpdated)
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
func (it *TokenGasForProcessingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenGasForProcessingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenGasForProcessingUpdated represents a GasForProcessingUpdated event raised by the Token contract.
type TokenGasForProcessingUpdated struct {
	NewValue *big.Int
	OldValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasForProcessingUpdated is a free log retrieval operation binding the contract event 0x40d7e40e79af4e8e5a9b3c57030d8ea93f13d669c06d448c4d631d4ae7d23db7.
//
// Solidity: event GasForProcessingUpdated(uint256 indexed newValue, uint256 indexed oldValue)
func (_Token *TokenFilterer) FilterGasForProcessingUpdated(opts *bind.FilterOpts, newValue []*big.Int, oldValue []*big.Int) (*TokenGasForProcessingUpdatedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "GasForProcessingUpdated", newValueRule, oldValueRule)
	if err != nil {
		return nil, err
	}
	return &TokenGasForProcessingUpdatedIterator{contract: _Token.contract, event: "GasForProcessingUpdated", logs: logs, sub: sub}, nil
}

// WatchGasForProcessingUpdated is a free log subscription operation binding the contract event 0x40d7e40e79af4e8e5a9b3c57030d8ea93f13d669c06d448c4d631d4ae7d23db7.
//
// Solidity: event GasForProcessingUpdated(uint256 indexed newValue, uint256 indexed oldValue)
func (_Token *TokenFilterer) WatchGasForProcessingUpdated(opts *bind.WatchOpts, sink chan<- *TokenGasForProcessingUpdated, newValue []*big.Int, oldValue []*big.Int) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "GasForProcessingUpdated", newValueRule, oldValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenGasForProcessingUpdated)
				if err := _Token.contract.UnpackLog(event, "GasForProcessingUpdated", log); err != nil {
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

// ParseGasForProcessingUpdated is a log parse operation binding the contract event 0x40d7e40e79af4e8e5a9b3c57030d8ea93f13d669c06d448c4d631d4ae7d23db7.
//
// Solidity: event GasForProcessingUpdated(uint256 indexed newValue, uint256 indexed oldValue)
func (_Token *TokenFilterer) ParseGasForProcessingUpdated(log types.Log) (*TokenGasForProcessingUpdated, error) {
	event := new(TokenGasForProcessingUpdated)
	if err := _Token.contract.UnpackLog(event, "GasForProcessingUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenLiquidityWalletUpdatedIterator is returned from FilterLiquidityWalletUpdated and is used to iterate over the raw logs and unpacked data for LiquidityWalletUpdated events raised by the Token contract.
type TokenLiquidityWalletUpdatedIterator struct {
	Event *TokenLiquidityWalletUpdated // Event containing the contract specifics and raw log

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
func (it *TokenLiquidityWalletUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenLiquidityWalletUpdated)
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
		it.Event = new(TokenLiquidityWalletUpdated)
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
func (it *TokenLiquidityWalletUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenLiquidityWalletUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenLiquidityWalletUpdated represents a LiquidityWalletUpdated event raised by the Token contract.
type TokenLiquidityWalletUpdated struct {
	NewLiquidityWallet common.Address
	OldLiquidityWallet common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLiquidityWalletUpdated is a free log retrieval operation binding the contract event 0x6080503d1da552ae8eb4b7b8a20245d9fabed014180510e7d1a05ea08fdb0f3e.
//
// Solidity: event LiquidityWalletUpdated(address indexed newLiquidityWallet, address indexed oldLiquidityWallet)
func (_Token *TokenFilterer) FilterLiquidityWalletUpdated(opts *bind.FilterOpts, newLiquidityWallet []common.Address, oldLiquidityWallet []common.Address) (*TokenLiquidityWalletUpdatedIterator, error) {

	var newLiquidityWalletRule []interface{}
	for _, newLiquidityWalletItem := range newLiquidityWallet {
		newLiquidityWalletRule = append(newLiquidityWalletRule, newLiquidityWalletItem)
	}
	var oldLiquidityWalletRule []interface{}
	for _, oldLiquidityWalletItem := range oldLiquidityWallet {
		oldLiquidityWalletRule = append(oldLiquidityWalletRule, oldLiquidityWalletItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "LiquidityWalletUpdated", newLiquidityWalletRule, oldLiquidityWalletRule)
	if err != nil {
		return nil, err
	}
	return &TokenLiquidityWalletUpdatedIterator{contract: _Token.contract, event: "LiquidityWalletUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityWalletUpdated is a free log subscription operation binding the contract event 0x6080503d1da552ae8eb4b7b8a20245d9fabed014180510e7d1a05ea08fdb0f3e.
//
// Solidity: event LiquidityWalletUpdated(address indexed newLiquidityWallet, address indexed oldLiquidityWallet)
func (_Token *TokenFilterer) WatchLiquidityWalletUpdated(opts *bind.WatchOpts, sink chan<- *TokenLiquidityWalletUpdated, newLiquidityWallet []common.Address, oldLiquidityWallet []common.Address) (event.Subscription, error) {

	var newLiquidityWalletRule []interface{}
	for _, newLiquidityWalletItem := range newLiquidityWallet {
		newLiquidityWalletRule = append(newLiquidityWalletRule, newLiquidityWalletItem)
	}
	var oldLiquidityWalletRule []interface{}
	for _, oldLiquidityWalletItem := range oldLiquidityWallet {
		oldLiquidityWalletRule = append(oldLiquidityWalletRule, oldLiquidityWalletItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "LiquidityWalletUpdated", newLiquidityWalletRule, oldLiquidityWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenLiquidityWalletUpdated)
				if err := _Token.contract.UnpackLog(event, "LiquidityWalletUpdated", log); err != nil {
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

// ParseLiquidityWalletUpdated is a log parse operation binding the contract event 0x6080503d1da552ae8eb4b7b8a20245d9fabed014180510e7d1a05ea08fdb0f3e.
//
// Solidity: event LiquidityWalletUpdated(address indexed newLiquidityWallet, address indexed oldLiquidityWallet)
func (_Token *TokenFilterer) ParseLiquidityWalletUpdated(log types.Log) (*TokenLiquidityWalletUpdated, error) {
	event := new(TokenLiquidityWalletUpdated)
	if err := _Token.contract.UnpackLog(event, "LiquidityWalletUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Token contract.
type TokenOwnershipTransferredIterator struct {
	Event *TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnershipTransferred)
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
		it.Event = new(TokenOwnershipTransferred)
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
func (it *TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnershipTransferred represents a OwnershipTransferred event raised by the Token contract.
type TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenOwnershipTransferredIterator{contract: _Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Token *TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnershipTransferred)
				if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Token *TokenFilterer) ParseOwnershipTransferred(log types.Log) (*TokenOwnershipTransferred, error) {
	event := new(TokenOwnershipTransferred)
	if err := _Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenProcessedDividendTrackerIterator is returned from FilterProcessedDividendTracker and is used to iterate over the raw logs and unpacked data for ProcessedDividendTracker events raised by the Token contract.
type TokenProcessedDividendTrackerIterator struct {
	Event *TokenProcessedDividendTracker // Event containing the contract specifics and raw log

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
func (it *TokenProcessedDividendTrackerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenProcessedDividendTracker)
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
		it.Event = new(TokenProcessedDividendTracker)
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
func (it *TokenProcessedDividendTrackerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenProcessedDividendTrackerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenProcessedDividendTracker represents a ProcessedDividendTracker event raised by the Token contract.
type TokenProcessedDividendTracker struct {
	Iterations         *big.Int
	Claims             *big.Int
	LastProcessedIndex *big.Int
	Automatic          bool
	Gas                *big.Int
	Processor          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProcessedDividendTracker is a free log retrieval operation binding the contract event 0xc864333d6121033635ab41b29ae52f10a22cf4438c3e4f1c4c68518feb2f8a98.
//
// Solidity: event ProcessedDividendTracker(uint256 iterations, uint256 claims, uint256 lastProcessedIndex, bool indexed automatic, uint256 gas, address indexed processor)
func (_Token *TokenFilterer) FilterProcessedDividendTracker(opts *bind.FilterOpts, automatic []bool, processor []common.Address) (*TokenProcessedDividendTrackerIterator, error) {

	var automaticRule []interface{}
	for _, automaticItem := range automatic {
		automaticRule = append(automaticRule, automaticItem)
	}

	var processorRule []interface{}
	for _, processorItem := range processor {
		processorRule = append(processorRule, processorItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "ProcessedDividendTracker", automaticRule, processorRule)
	if err != nil {
		return nil, err
	}
	return &TokenProcessedDividendTrackerIterator{contract: _Token.contract, event: "ProcessedDividendTracker", logs: logs, sub: sub}, nil
}

// WatchProcessedDividendTracker is a free log subscription operation binding the contract event 0xc864333d6121033635ab41b29ae52f10a22cf4438c3e4f1c4c68518feb2f8a98.
//
// Solidity: event ProcessedDividendTracker(uint256 iterations, uint256 claims, uint256 lastProcessedIndex, bool indexed automatic, uint256 gas, address indexed processor)
func (_Token *TokenFilterer) WatchProcessedDividendTracker(opts *bind.WatchOpts, sink chan<- *TokenProcessedDividendTracker, automatic []bool, processor []common.Address) (event.Subscription, error) {

	var automaticRule []interface{}
	for _, automaticItem := range automatic {
		automaticRule = append(automaticRule, automaticItem)
	}

	var processorRule []interface{}
	for _, processorItem := range processor {
		processorRule = append(processorRule, processorItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "ProcessedDividendTracker", automaticRule, processorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenProcessedDividendTracker)
				if err := _Token.contract.UnpackLog(event, "ProcessedDividendTracker", log); err != nil {
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

// ParseProcessedDividendTracker is a log parse operation binding the contract event 0xc864333d6121033635ab41b29ae52f10a22cf4438c3e4f1c4c68518feb2f8a98.
//
// Solidity: event ProcessedDividendTracker(uint256 iterations, uint256 claims, uint256 lastProcessedIndex, bool indexed automatic, uint256 gas, address indexed processor)
func (_Token *TokenFilterer) ParseProcessedDividendTracker(log types.Log) (*TokenProcessedDividendTracker, error) {
	event := new(TokenProcessedDividendTracker)
	if err := _Token.contract.UnpackLog(event, "ProcessedDividendTracker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSendDividendsIterator is returned from FilterSendDividends and is used to iterate over the raw logs and unpacked data for SendDividends events raised by the Token contract.
type TokenSendDividendsIterator struct {
	Event *TokenSendDividends // Event containing the contract specifics and raw log

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
func (it *TokenSendDividendsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSendDividends)
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
		it.Event = new(TokenSendDividends)
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
func (it *TokenSendDividendsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSendDividendsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSendDividends represents a SendDividends event raised by the Token contract.
type TokenSendDividends struct {
	TokensSwapped *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSendDividends is a free log retrieval operation binding the contract event 0x80195cc573b02cc48460cbca6e6e4cc85ddb91959d946e1c3025ea3d87942dc3.
//
// Solidity: event SendDividends(uint256 tokensSwapped, uint256 amount)
func (_Token *TokenFilterer) FilterSendDividends(opts *bind.FilterOpts) (*TokenSendDividendsIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "SendDividends")
	if err != nil {
		return nil, err
	}
	return &TokenSendDividendsIterator{contract: _Token.contract, event: "SendDividends", logs: logs, sub: sub}, nil
}

// WatchSendDividends is a free log subscription operation binding the contract event 0x80195cc573b02cc48460cbca6e6e4cc85ddb91959d946e1c3025ea3d87942dc3.
//
// Solidity: event SendDividends(uint256 tokensSwapped, uint256 amount)
func (_Token *TokenFilterer) WatchSendDividends(opts *bind.WatchOpts, sink chan<- *TokenSendDividends) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "SendDividends")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSendDividends)
				if err := _Token.contract.UnpackLog(event, "SendDividends", log); err != nil {
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

// ParseSendDividends is a log parse operation binding the contract event 0x80195cc573b02cc48460cbca6e6e4cc85ddb91959d946e1c3025ea3d87942dc3.
//
// Solidity: event SendDividends(uint256 tokensSwapped, uint256 amount)
func (_Token *TokenFilterer) ParseSendDividends(log types.Log) (*TokenSendDividends, error) {
	event := new(TokenSendDividends)
	if err := _Token.contract.UnpackLog(event, "SendDividends", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSetAutomatedMarketMakerPairIterator is returned from FilterSetAutomatedMarketMakerPair and is used to iterate over the raw logs and unpacked data for SetAutomatedMarketMakerPair events raised by the Token contract.
type TokenSetAutomatedMarketMakerPairIterator struct {
	Event *TokenSetAutomatedMarketMakerPair // Event containing the contract specifics and raw log

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
func (it *TokenSetAutomatedMarketMakerPairIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSetAutomatedMarketMakerPair)
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
		it.Event = new(TokenSetAutomatedMarketMakerPair)
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
func (it *TokenSetAutomatedMarketMakerPairIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSetAutomatedMarketMakerPairIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSetAutomatedMarketMakerPair represents a SetAutomatedMarketMakerPair event raised by the Token contract.
type TokenSetAutomatedMarketMakerPair struct {
	Pair  common.Address
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAutomatedMarketMakerPair is a free log retrieval operation binding the contract event 0xffa9187bf1f18bf477bd0ea1bcbb64e93b6a98132473929edfce215cd9b16fab.
//
// Solidity: event SetAutomatedMarketMakerPair(address indexed pair, bool indexed value)
func (_Token *TokenFilterer) FilterSetAutomatedMarketMakerPair(opts *bind.FilterOpts, pair []common.Address, value []bool) (*TokenSetAutomatedMarketMakerPairIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "SetAutomatedMarketMakerPair", pairRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &TokenSetAutomatedMarketMakerPairIterator{contract: _Token.contract, event: "SetAutomatedMarketMakerPair", logs: logs, sub: sub}, nil
}

// WatchSetAutomatedMarketMakerPair is a free log subscription operation binding the contract event 0xffa9187bf1f18bf477bd0ea1bcbb64e93b6a98132473929edfce215cd9b16fab.
//
// Solidity: event SetAutomatedMarketMakerPair(address indexed pair, bool indexed value)
func (_Token *TokenFilterer) WatchSetAutomatedMarketMakerPair(opts *bind.WatchOpts, sink chan<- *TokenSetAutomatedMarketMakerPair, pair []common.Address, value []bool) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "SetAutomatedMarketMakerPair", pairRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSetAutomatedMarketMakerPair)
				if err := _Token.contract.UnpackLog(event, "SetAutomatedMarketMakerPair", log); err != nil {
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

// ParseSetAutomatedMarketMakerPair is a log parse operation binding the contract event 0xffa9187bf1f18bf477bd0ea1bcbb64e93b6a98132473929edfce215cd9b16fab.
//
// Solidity: event SetAutomatedMarketMakerPair(address indexed pair, bool indexed value)
func (_Token *TokenFilterer) ParseSetAutomatedMarketMakerPair(log types.Log) (*TokenSetAutomatedMarketMakerPair, error) {
	event := new(TokenSetAutomatedMarketMakerPair)
	if err := _Token.contract.UnpackLog(event, "SetAutomatedMarketMakerPair", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenSwapAndLiquifyIterator is returned from FilterSwapAndLiquify and is used to iterate over the raw logs and unpacked data for SwapAndLiquify events raised by the Token contract.
type TokenSwapAndLiquifyIterator struct {
	Event *TokenSwapAndLiquify // Event containing the contract specifics and raw log

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
func (it *TokenSwapAndLiquifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenSwapAndLiquify)
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
		it.Event = new(TokenSwapAndLiquify)
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
func (it *TokenSwapAndLiquifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenSwapAndLiquifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenSwapAndLiquify represents a SwapAndLiquify event raised by the Token contract.
type TokenSwapAndLiquify struct {
	TokensSwapped      *big.Int
	EthReceived        *big.Int
	TokensIntoLiqudity *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwapAndLiquify is a free log retrieval operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Token *TokenFilterer) FilterSwapAndLiquify(opts *bind.FilterOpts) (*TokenSwapAndLiquifyIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "SwapAndLiquify")
	if err != nil {
		return nil, err
	}
	return &TokenSwapAndLiquifyIterator{contract: _Token.contract, event: "SwapAndLiquify", logs: logs, sub: sub}, nil
}

// WatchSwapAndLiquify is a free log subscription operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Token *TokenFilterer) WatchSwapAndLiquify(opts *bind.WatchOpts, sink chan<- *TokenSwapAndLiquify) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "SwapAndLiquify")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenSwapAndLiquify)
				if err := _Token.contract.UnpackLog(event, "SwapAndLiquify", log); err != nil {
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

// ParseSwapAndLiquify is a log parse operation binding the contract event 0x17bbfb9a6069321b6ded73bd96327c9e6b7212a5cd51ff219cd61370acafb561.
//
// Solidity: event SwapAndLiquify(uint256 tokensSwapped, uint256 ethReceived, uint256 tokensIntoLiqudity)
func (_Token *TokenFilterer) ParseSwapAndLiquify(log types.Log) (*TokenSwapAndLiquify, error) {
	event := new(TokenSwapAndLiquify)
	if err := _Token.contract.UnpackLog(event, "SwapAndLiquify", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenUpdateDividendTrackerIterator is returned from FilterUpdateDividendTracker and is used to iterate over the raw logs and unpacked data for UpdateDividendTracker events raised by the Token contract.
type TokenUpdateDividendTrackerIterator struct {
	Event *TokenUpdateDividendTracker // Event containing the contract specifics and raw log

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
func (it *TokenUpdateDividendTrackerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUpdateDividendTracker)
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
		it.Event = new(TokenUpdateDividendTracker)
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
func (it *TokenUpdateDividendTrackerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUpdateDividendTrackerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUpdateDividendTracker represents a UpdateDividendTracker event raised by the Token contract.
type TokenUpdateDividendTracker struct {
	NewAddress common.Address
	OldAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateDividendTracker is a free log retrieval operation binding the contract event 0x90c7d74461c613da5efa97d90740869367d74ab3aa5837aa4ae9a975f954b7a8.
//
// Solidity: event UpdateDividendTracker(address indexed newAddress, address indexed oldAddress)
func (_Token *TokenFilterer) FilterUpdateDividendTracker(opts *bind.FilterOpts, newAddress []common.Address, oldAddress []common.Address) (*TokenUpdateDividendTrackerIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "UpdateDividendTracker", newAddressRule, oldAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenUpdateDividendTrackerIterator{contract: _Token.contract, event: "UpdateDividendTracker", logs: logs, sub: sub}, nil
}

// WatchUpdateDividendTracker is a free log subscription operation binding the contract event 0x90c7d74461c613da5efa97d90740869367d74ab3aa5837aa4ae9a975f954b7a8.
//
// Solidity: event UpdateDividendTracker(address indexed newAddress, address indexed oldAddress)
func (_Token *TokenFilterer) WatchUpdateDividendTracker(opts *bind.WatchOpts, sink chan<- *TokenUpdateDividendTracker, newAddress []common.Address, oldAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "UpdateDividendTracker", newAddressRule, oldAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUpdateDividendTracker)
				if err := _Token.contract.UnpackLog(event, "UpdateDividendTracker", log); err != nil {
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

// ParseUpdateDividendTracker is a log parse operation binding the contract event 0x90c7d74461c613da5efa97d90740869367d74ab3aa5837aa4ae9a975f954b7a8.
//
// Solidity: event UpdateDividendTracker(address indexed newAddress, address indexed oldAddress)
func (_Token *TokenFilterer) ParseUpdateDividendTracker(log types.Log) (*TokenUpdateDividendTracker, error) {
	event := new(TokenUpdateDividendTracker)
	if err := _Token.contract.UnpackLog(event, "UpdateDividendTracker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenUpdateUniswapV2RouterIterator is returned from FilterUpdateUniswapV2Router and is used to iterate over the raw logs and unpacked data for UpdateUniswapV2Router events raised by the Token contract.
type TokenUpdateUniswapV2RouterIterator struct {
	Event *TokenUpdateUniswapV2Router // Event containing the contract specifics and raw log

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
func (it *TokenUpdateUniswapV2RouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUpdateUniswapV2Router)
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
		it.Event = new(TokenUpdateUniswapV2Router)
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
func (it *TokenUpdateUniswapV2RouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUpdateUniswapV2RouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUpdateUniswapV2Router represents a UpdateUniswapV2Router event raised by the Token contract.
type TokenUpdateUniswapV2Router struct {
	NewAddress common.Address
	OldAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateUniswapV2Router is a free log retrieval operation binding the contract event 0x8fc842bbd331dfa973645f4ed48b11683d501ebf1352708d77a5da2ab49a576e.
//
// Solidity: event UpdateUniswapV2Router(address indexed newAddress, address indexed oldAddress)
func (_Token *TokenFilterer) FilterUpdateUniswapV2Router(opts *bind.FilterOpts, newAddress []common.Address, oldAddress []common.Address) (*TokenUpdateUniswapV2RouterIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "UpdateUniswapV2Router", newAddressRule, oldAddressRule)
	if err != nil {
		return nil, err
	}
	return &TokenUpdateUniswapV2RouterIterator{contract: _Token.contract, event: "UpdateUniswapV2Router", logs: logs, sub: sub}, nil
}

// WatchUpdateUniswapV2Router is a free log subscription operation binding the contract event 0x8fc842bbd331dfa973645f4ed48b11683d501ebf1352708d77a5da2ab49a576e.
//
// Solidity: event UpdateUniswapV2Router(address indexed newAddress, address indexed oldAddress)
func (_Token *TokenFilterer) WatchUpdateUniswapV2Router(opts *bind.WatchOpts, sink chan<- *TokenUpdateUniswapV2Router, newAddress []common.Address, oldAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "UpdateUniswapV2Router", newAddressRule, oldAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUpdateUniswapV2Router)
				if err := _Token.contract.UnpackLog(event, "UpdateUniswapV2Router", log); err != nil {
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

// ParseUpdateUniswapV2Router is a log parse operation binding the contract event 0x8fc842bbd331dfa973645f4ed48b11683d501ebf1352708d77a5da2ab49a576e.
//
// Solidity: event UpdateUniswapV2Router(address indexed newAddress, address indexed oldAddress)
func (_Token *TokenFilterer) ParseUpdateUniswapV2Router(log types.Log) (*TokenUpdateUniswapV2Router, error) {
	event := new(TokenUpdateUniswapV2Router)
	if err := _Token.contract.UnpackLog(event, "UpdateUniswapV2Router", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
