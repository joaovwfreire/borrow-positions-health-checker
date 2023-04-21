// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CometMainInterface

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

// CometMainInterfaceAssetInfo is an auto generated low-level Go binding around an user-defined struct.
type CometMainInterfaceAssetInfo struct {
	Offset                    uint8
	Asset                     common.Address
	PriceFeed                 common.Address
	Scale                     uint64
	BorrowCollateralFactor    uint64
	LiquidateCollateralFactor uint64
	LiquidationFactor         uint64
	SupplyCap                 *big.Int
}

// CometMainInterfaceMetaData contains all meta data concerning the CometMainInterface contract.
var CometMainInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Absurd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadDecimals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadDiscount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadMinimum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowCFTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientReserves\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidateCFTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSelfTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCollateralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotForSale\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLiquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SupplyCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimestampTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooMuchSlippage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferInFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferOutFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"absorber\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAbsorbed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdValue\",\"type\":\"uint256\"}],\"name\":\"AbsorbCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"absorber\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"basePaidOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdValue\",\"type\":\"uint256\"}],\"name\":\"AbsorbDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"BuyCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supplyPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"withdrawPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"absorbPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"buyPaused\",\"type\":\"bool\"}],\"name\":\"PauseAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SupplyCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawReserves\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"absorber\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"absorb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accrueAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveThis\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseBorrowMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseMinForRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenPriceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTrackingBorrowSpeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTrackingSupplySpeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowKink\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowPerSecondInterestRateBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowPerSecondInterestRateSlopeHigh\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowPerSecondInterestRateSlopeLow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"buyCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extensionDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"offset\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"scale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"structCometMainInterface.AssetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetInfoByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"offset\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"scale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"structCometMainInterface.AssetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"utilization\",\"type\":\"uint256\"}],\"name\":\"getBorrowRate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getCollateralReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"utilization\",\"type\":\"uint256\"}],\"name\":\"getSupplyRate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUtilization\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAbsorbPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBorrowCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBuyPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSupplyPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTransferPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isWithdrawPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numAssets\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"supplyPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"withdrawPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"absorbPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"buyPaused\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"quoteCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storeFrontPriceFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"supply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"supplyFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyKink\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyPerSecondInterestRateBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyPerSecondInterestRateSlopeHigh\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyPerSecondInterestRateSlopeLow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"supplyTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trackingIndexScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferAssetFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CometMainInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use CometMainInterfaceMetaData.ABI instead.
var CometMainInterfaceABI = CometMainInterfaceMetaData.ABI

// CometMainInterface is an auto generated Go binding around an Ethereum contract.
type CometMainInterface struct {
	CometMainInterfaceCaller     // Read-only binding to the contract
	CometMainInterfaceTransactor // Write-only binding to the contract
	CometMainInterfaceFilterer   // Log filterer for contract events
}

// CometMainInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type CometMainInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CometMainInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CometMainInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CometMainInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CometMainInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CometMainInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CometMainInterfaceSession struct {
	Contract     *CometMainInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CometMainInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CometMainInterfaceCallerSession struct {
	Contract *CometMainInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CometMainInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CometMainInterfaceTransactorSession struct {
	Contract     *CometMainInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CometMainInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type CometMainInterfaceRaw struct {
	Contract *CometMainInterface // Generic contract binding to access the raw methods on
}

// CometMainInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CometMainInterfaceCallerRaw struct {
	Contract *CometMainInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// CometMainInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CometMainInterfaceTransactorRaw struct {
	Contract *CometMainInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCometMainInterface creates a new instance of CometMainInterface, bound to a specific deployed contract.
func NewCometMainInterface(address common.Address, backend bind.ContractBackend) (*CometMainInterface, error) {
	contract, err := bindCometMainInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CometMainInterface{CometMainInterfaceCaller: CometMainInterfaceCaller{contract: contract}, CometMainInterfaceTransactor: CometMainInterfaceTransactor{contract: contract}, CometMainInterfaceFilterer: CometMainInterfaceFilterer{contract: contract}}, nil
}

// NewCometMainInterfaceCaller creates a new read-only instance of CometMainInterface, bound to a specific deployed contract.
func NewCometMainInterfaceCaller(address common.Address, caller bind.ContractCaller) (*CometMainInterfaceCaller, error) {
	contract, err := bindCometMainInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceCaller{contract: contract}, nil
}

// NewCometMainInterfaceTransactor creates a new write-only instance of CometMainInterface, bound to a specific deployed contract.
func NewCometMainInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*CometMainInterfaceTransactor, error) {
	contract, err := bindCometMainInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceTransactor{contract: contract}, nil
}

// NewCometMainInterfaceFilterer creates a new log filterer instance of CometMainInterface, bound to a specific deployed contract.
func NewCometMainInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*CometMainInterfaceFilterer, error) {
	contract, err := bindCometMainInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceFilterer{contract: contract}, nil
}

// bindCometMainInterface binds a generic wrapper to an already deployed contract.
func bindCometMainInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CometMainInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CometMainInterface *CometMainInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CometMainInterface.Contract.CometMainInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CometMainInterface *CometMainInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CometMainInterface.Contract.CometMainInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CometMainInterface *CometMainInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CometMainInterface.Contract.CometMainInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CometMainInterface *CometMainInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CometMainInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CometMainInterface *CometMainInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CometMainInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CometMainInterface *CometMainInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CometMainInterface.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CometMainInterface.Contract.BalanceOf(&_CometMainInterface.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CometMainInterface.Contract.BalanceOf(&_CometMainInterface.CallOpts, owner)
}

// BaseBorrowMin is a free data retrieval call binding the contract method 0x300e6beb.
//
// Solidity: function baseBorrowMin() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) BaseBorrowMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "baseBorrowMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseBorrowMin is a free data retrieval call binding the contract method 0x300e6beb.
//
// Solidity: function baseBorrowMin() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) BaseBorrowMin() (*big.Int, error) {
	return _CometMainInterface.Contract.BaseBorrowMin(&_CometMainInterface.CallOpts)
}

// BaseBorrowMin is a free data retrieval call binding the contract method 0x300e6beb.
//
// Solidity: function baseBorrowMin() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) BaseBorrowMin() (*big.Int, error) {
	return _CometMainInterface.Contract.BaseBorrowMin(&_CometMainInterface.CallOpts)
}

// BaseMinForRewards is a free data retrieval call binding the contract method 0x9364e18a.
//
// Solidity: function baseMinForRewards() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) BaseMinForRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "baseMinForRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseMinForRewards is a free data retrieval call binding the contract method 0x9364e18a.
//
// Solidity: function baseMinForRewards() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) BaseMinForRewards() (*big.Int, error) {
	return _CometMainInterface.Contract.BaseMinForRewards(&_CometMainInterface.CallOpts)
}

// BaseMinForRewards is a free data retrieval call binding the contract method 0x9364e18a.
//
// Solidity: function baseMinForRewards() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) BaseMinForRewards() (*big.Int, error) {
	return _CometMainInterface.Contract.BaseMinForRewards(&_CometMainInterface.CallOpts)
}

// BaseScale is a free data retrieval call binding the contract method 0x44c1e5eb.
//
// Solidity: function baseScale() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) BaseScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "baseScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseScale is a free data retrieval call binding the contract method 0x44c1e5eb.
//
// Solidity: function baseScale() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) BaseScale() (*big.Int, error) {
	return _CometMainInterface.Contract.BaseScale(&_CometMainInterface.CallOpts)
}

// BaseScale is a free data retrieval call binding the contract method 0x44c1e5eb.
//
// Solidity: function baseScale() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) BaseScale() (*big.Int, error) {
	return _CometMainInterface.Contract.BaseScale(&_CometMainInterface.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_CometMainInterface *CometMainInterfaceCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "baseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_CometMainInterface *CometMainInterfaceSession) BaseToken() (common.Address, error) {
	return _CometMainInterface.Contract.BaseToken(&_CometMainInterface.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_CometMainInterface *CometMainInterfaceCallerSession) BaseToken() (common.Address, error) {
	return _CometMainInterface.Contract.BaseToken(&_CometMainInterface.CallOpts)
}

// BaseTokenPriceFeed is a free data retrieval call binding the contract method 0xe7dad6bd.
//
// Solidity: function baseTokenPriceFeed() view returns(address)
func (_CometMainInterface *CometMainInterfaceCaller) BaseTokenPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "baseTokenPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseTokenPriceFeed is a free data retrieval call binding the contract method 0xe7dad6bd.
//
// Solidity: function baseTokenPriceFeed() view returns(address)
func (_CometMainInterface *CometMainInterfaceSession) BaseTokenPriceFeed() (common.Address, error) {
	return _CometMainInterface.Contract.BaseTokenPriceFeed(&_CometMainInterface.CallOpts)
}

// BaseTokenPriceFeed is a free data retrieval call binding the contract method 0xe7dad6bd.
//
// Solidity: function baseTokenPriceFeed() view returns(address)
func (_CometMainInterface *CometMainInterfaceCallerSession) BaseTokenPriceFeed() (common.Address, error) {
	return _CometMainInterface.Contract.BaseTokenPriceFeed(&_CometMainInterface.CallOpts)
}

// BaseTrackingBorrowSpeed is a free data retrieval call binding the contract method 0x9ea99a5a.
//
// Solidity: function baseTrackingBorrowSpeed() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) BaseTrackingBorrowSpeed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "baseTrackingBorrowSpeed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseTrackingBorrowSpeed is a free data retrieval call binding the contract method 0x9ea99a5a.
//
// Solidity: function baseTrackingBorrowSpeed() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) BaseTrackingBorrowSpeed() (*big.Int, error) {
	return _CometMainInterface.Contract.BaseTrackingBorrowSpeed(&_CometMainInterface.CallOpts)
}

// BaseTrackingBorrowSpeed is a free data retrieval call binding the contract method 0x9ea99a5a.
//
// Solidity: function baseTrackingBorrowSpeed() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) BaseTrackingBorrowSpeed() (*big.Int, error) {
	return _CometMainInterface.Contract.BaseTrackingBorrowSpeed(&_CometMainInterface.CallOpts)
}

// BaseTrackingSupplySpeed is a free data retrieval call binding the contract method 0x189bb2f1.
//
// Solidity: function baseTrackingSupplySpeed() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) BaseTrackingSupplySpeed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "baseTrackingSupplySpeed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseTrackingSupplySpeed is a free data retrieval call binding the contract method 0x189bb2f1.
//
// Solidity: function baseTrackingSupplySpeed() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) BaseTrackingSupplySpeed() (*big.Int, error) {
	return _CometMainInterface.Contract.BaseTrackingSupplySpeed(&_CometMainInterface.CallOpts)
}

// BaseTrackingSupplySpeed is a free data retrieval call binding the contract method 0x189bb2f1.
//
// Solidity: function baseTrackingSupplySpeed() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) BaseTrackingSupplySpeed() (*big.Int, error) {
	return _CometMainInterface.Contract.BaseTrackingSupplySpeed(&_CometMainInterface.CallOpts)
}

// BorrowBalanceOf is a free data retrieval call binding the contract method 0x374c49b4.
//
// Solidity: function borrowBalanceOf(address account) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) BorrowBalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "borrowBalanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceOf is a free data retrieval call binding the contract method 0x374c49b4.
//
// Solidity: function borrowBalanceOf(address account) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) BorrowBalanceOf(account common.Address) (*big.Int, error) {
	return _CometMainInterface.Contract.BorrowBalanceOf(&_CometMainInterface.CallOpts, account)
}

// BorrowBalanceOf is a free data retrieval call binding the contract method 0x374c49b4.
//
// Solidity: function borrowBalanceOf(address account) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) BorrowBalanceOf(account common.Address) (*big.Int, error) {
	return _CometMainInterface.Contract.BorrowBalanceOf(&_CometMainInterface.CallOpts, account)
}

// BorrowKink is a free data retrieval call binding the contract method 0x9241a561.
//
// Solidity: function borrowKink() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) BorrowKink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "borrowKink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowKink is a free data retrieval call binding the contract method 0x9241a561.
//
// Solidity: function borrowKink() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) BorrowKink() (*big.Int, error) {
	return _CometMainInterface.Contract.BorrowKink(&_CometMainInterface.CallOpts)
}

// BorrowKink is a free data retrieval call binding the contract method 0x9241a561.
//
// Solidity: function borrowKink() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) BorrowKink() (*big.Int, error) {
	return _CometMainInterface.Contract.BorrowKink(&_CometMainInterface.CallOpts)
}

// BorrowPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x7914acc7.
//
// Solidity: function borrowPerSecondInterestRateBase() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) BorrowPerSecondInterestRateBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "borrowPerSecondInterestRateBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x7914acc7.
//
// Solidity: function borrowPerSecondInterestRateBase() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) BorrowPerSecondInterestRateBase() (*big.Int, error) {
	return _CometMainInterface.Contract.BorrowPerSecondInterestRateBase(&_CometMainInterface.CallOpts)
}

// BorrowPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x7914acc7.
//
// Solidity: function borrowPerSecondInterestRateBase() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) BorrowPerSecondInterestRateBase() (*big.Int, error) {
	return _CometMainInterface.Contract.BorrowPerSecondInterestRateBase(&_CometMainInterface.CallOpts)
}

// BorrowPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x2a48cf12.
//
// Solidity: function borrowPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) BorrowPerSecondInterestRateSlopeHigh(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "borrowPerSecondInterestRateSlopeHigh")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x2a48cf12.
//
// Solidity: function borrowPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) BorrowPerSecondInterestRateSlopeHigh() (*big.Int, error) {
	return _CometMainInterface.Contract.BorrowPerSecondInterestRateSlopeHigh(&_CometMainInterface.CallOpts)
}

// BorrowPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x2a48cf12.
//
// Solidity: function borrowPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) BorrowPerSecondInterestRateSlopeHigh() (*big.Int, error) {
	return _CometMainInterface.Contract.BorrowPerSecondInterestRateSlopeHigh(&_CometMainInterface.CallOpts)
}

// BorrowPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x2d05670b.
//
// Solidity: function borrowPerSecondInterestRateSlopeLow() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) BorrowPerSecondInterestRateSlopeLow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "borrowPerSecondInterestRateSlopeLow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x2d05670b.
//
// Solidity: function borrowPerSecondInterestRateSlopeLow() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) BorrowPerSecondInterestRateSlopeLow() (*big.Int, error) {
	return _CometMainInterface.Contract.BorrowPerSecondInterestRateSlopeLow(&_CometMainInterface.CallOpts)
}

// BorrowPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x2d05670b.
//
// Solidity: function borrowPerSecondInterestRateSlopeLow() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) BorrowPerSecondInterestRateSlopeLow() (*big.Int, error) {
	return _CometMainInterface.Contract.BorrowPerSecondInterestRateSlopeLow(&_CometMainInterface.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CometMainInterface *CometMainInterfaceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CometMainInterface *CometMainInterfaceSession) Decimals() (uint8, error) {
	return _CometMainInterface.Contract.Decimals(&_CometMainInterface.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CometMainInterface *CometMainInterfaceCallerSession) Decimals() (uint8, error) {
	return _CometMainInterface.Contract.Decimals(&_CometMainInterface.CallOpts)
}

// ExtensionDelegate is a free data retrieval call binding the contract method 0x44ff241d.
//
// Solidity: function extensionDelegate() view returns(address)
func (_CometMainInterface *CometMainInterfaceCaller) ExtensionDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "extensionDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtensionDelegate is a free data retrieval call binding the contract method 0x44ff241d.
//
// Solidity: function extensionDelegate() view returns(address)
func (_CometMainInterface *CometMainInterfaceSession) ExtensionDelegate() (common.Address, error) {
	return _CometMainInterface.Contract.ExtensionDelegate(&_CometMainInterface.CallOpts)
}

// ExtensionDelegate is a free data retrieval call binding the contract method 0x44ff241d.
//
// Solidity: function extensionDelegate() view returns(address)
func (_CometMainInterface *CometMainInterfaceCallerSession) ExtensionDelegate() (common.Address, error) {
	return _CometMainInterface.Contract.ExtensionDelegate(&_CometMainInterface.CallOpts)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xc8c7fe6b.
//
// Solidity: function getAssetInfo(uint8 i) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_CometMainInterface *CometMainInterfaceCaller) GetAssetInfo(opts *bind.CallOpts, i uint8) (CometMainInterfaceAssetInfo, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "getAssetInfo", i)

	if err != nil {
		return *new(CometMainInterfaceAssetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CometMainInterfaceAssetInfo)).(*CometMainInterfaceAssetInfo)

	return out0, err

}

// GetAssetInfo is a free data retrieval call binding the contract method 0xc8c7fe6b.
//
// Solidity: function getAssetInfo(uint8 i) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_CometMainInterface *CometMainInterfaceSession) GetAssetInfo(i uint8) (CometMainInterfaceAssetInfo, error) {
	return _CometMainInterface.Contract.GetAssetInfo(&_CometMainInterface.CallOpts, i)
}

// GetAssetInfo is a free data retrieval call binding the contract method 0xc8c7fe6b.
//
// Solidity: function getAssetInfo(uint8 i) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_CometMainInterface *CometMainInterfaceCallerSession) GetAssetInfo(i uint8) (CometMainInterfaceAssetInfo, error) {
	return _CometMainInterface.Contract.GetAssetInfo(&_CometMainInterface.CallOpts, i)
}

// GetAssetInfoByAddress is a free data retrieval call binding the contract method 0x3b3bec2e.
//
// Solidity: function getAssetInfoByAddress(address asset) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_CometMainInterface *CometMainInterfaceCaller) GetAssetInfoByAddress(opts *bind.CallOpts, asset common.Address) (CometMainInterfaceAssetInfo, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "getAssetInfoByAddress", asset)

	if err != nil {
		return *new(CometMainInterfaceAssetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CometMainInterfaceAssetInfo)).(*CometMainInterfaceAssetInfo)

	return out0, err

}

// GetAssetInfoByAddress is a free data retrieval call binding the contract method 0x3b3bec2e.
//
// Solidity: function getAssetInfoByAddress(address asset) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_CometMainInterface *CometMainInterfaceSession) GetAssetInfoByAddress(asset common.Address) (CometMainInterfaceAssetInfo, error) {
	return _CometMainInterface.Contract.GetAssetInfoByAddress(&_CometMainInterface.CallOpts, asset)
}

// GetAssetInfoByAddress is a free data retrieval call binding the contract method 0x3b3bec2e.
//
// Solidity: function getAssetInfoByAddress(address asset) view returns((uint8,address,address,uint64,uint64,uint64,uint64,uint128))
func (_CometMainInterface *CometMainInterfaceCallerSession) GetAssetInfoByAddress(asset common.Address) (CometMainInterfaceAssetInfo, error) {
	return _CometMainInterface.Contract.GetAssetInfoByAddress(&_CometMainInterface.CallOpts, asset)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x9fa83b5a.
//
// Solidity: function getBorrowRate(uint256 utilization) view returns(uint64)
func (_CometMainInterface *CometMainInterfaceCaller) GetBorrowRate(opts *bind.CallOpts, utilization *big.Int) (uint64, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "getBorrowRate", utilization)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBorrowRate is a free data retrieval call binding the contract method 0x9fa83b5a.
//
// Solidity: function getBorrowRate(uint256 utilization) view returns(uint64)
func (_CometMainInterface *CometMainInterfaceSession) GetBorrowRate(utilization *big.Int) (uint64, error) {
	return _CometMainInterface.Contract.GetBorrowRate(&_CometMainInterface.CallOpts, utilization)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x9fa83b5a.
//
// Solidity: function getBorrowRate(uint256 utilization) view returns(uint64)
func (_CometMainInterface *CometMainInterfaceCallerSession) GetBorrowRate(utilization *big.Int) (uint64, error) {
	return _CometMainInterface.Contract.GetBorrowRate(&_CometMainInterface.CallOpts, utilization)
}

// GetCollateralReserves is a free data retrieval call binding the contract method 0x9ff567f8.
//
// Solidity: function getCollateralReserves(address asset) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) GetCollateralReserves(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "getCollateralReserves", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralReserves is a free data retrieval call binding the contract method 0x9ff567f8.
//
// Solidity: function getCollateralReserves(address asset) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) GetCollateralReserves(asset common.Address) (*big.Int, error) {
	return _CometMainInterface.Contract.GetCollateralReserves(&_CometMainInterface.CallOpts, asset)
}

// GetCollateralReserves is a free data retrieval call binding the contract method 0x9ff567f8.
//
// Solidity: function getCollateralReserves(address asset) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) GetCollateralReserves(asset common.Address) (*big.Int, error) {
	return _CometMainInterface.Contract.GetCollateralReserves(&_CometMainInterface.CallOpts, asset)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address priceFeed) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) GetPrice(opts *bind.CallOpts, priceFeed common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "getPrice", priceFeed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address priceFeed) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) GetPrice(priceFeed common.Address) (*big.Int, error) {
	return _CometMainInterface.Contract.GetPrice(&_CometMainInterface.CallOpts, priceFeed)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address priceFeed) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) GetPrice(priceFeed common.Address) (*big.Int, error) {
	return _CometMainInterface.Contract.GetPrice(&_CometMainInterface.CallOpts, priceFeed)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(int256)
func (_CometMainInterface *CometMainInterfaceCaller) GetReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "getReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(int256)
func (_CometMainInterface *CometMainInterfaceSession) GetReserves() (*big.Int, error) {
	return _CometMainInterface.Contract.GetReserves(&_CometMainInterface.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(int256)
func (_CometMainInterface *CometMainInterfaceCallerSession) GetReserves() (*big.Int, error) {
	return _CometMainInterface.Contract.GetReserves(&_CometMainInterface.CallOpts)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xd955759d.
//
// Solidity: function getSupplyRate(uint256 utilization) view returns(uint64)
func (_CometMainInterface *CometMainInterfaceCaller) GetSupplyRate(opts *bind.CallOpts, utilization *big.Int) (uint64, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "getSupplyRate", utilization)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetSupplyRate is a free data retrieval call binding the contract method 0xd955759d.
//
// Solidity: function getSupplyRate(uint256 utilization) view returns(uint64)
func (_CometMainInterface *CometMainInterfaceSession) GetSupplyRate(utilization *big.Int) (uint64, error) {
	return _CometMainInterface.Contract.GetSupplyRate(&_CometMainInterface.CallOpts, utilization)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xd955759d.
//
// Solidity: function getSupplyRate(uint256 utilization) view returns(uint64)
func (_CometMainInterface *CometMainInterfaceCallerSession) GetSupplyRate(utilization *big.Int) (uint64, error) {
	return _CometMainInterface.Contract.GetSupplyRate(&_CometMainInterface.CallOpts, utilization)
}

// GetUtilization is a free data retrieval call binding the contract method 0x7eb71131.
//
// Solidity: function getUtilization() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) GetUtilization(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "getUtilization")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUtilization is a free data retrieval call binding the contract method 0x7eb71131.
//
// Solidity: function getUtilization() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) GetUtilization() (*big.Int, error) {
	return _CometMainInterface.Contract.GetUtilization(&_CometMainInterface.CallOpts)
}

// GetUtilization is a free data retrieval call binding the contract method 0x7eb71131.
//
// Solidity: function getUtilization() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) GetUtilization() (*big.Int, error) {
	return _CometMainInterface.Contract.GetUtilization(&_CometMainInterface.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_CometMainInterface *CometMainInterfaceCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_CometMainInterface *CometMainInterfaceSession) Governor() (common.Address, error) {
	return _CometMainInterface.Contract.Governor(&_CometMainInterface.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_CometMainInterface *CometMainInterfaceCallerSession) Governor() (common.Address, error) {
	return _CometMainInterface.Contract.Governor(&_CometMainInterface.CallOpts)
}

// IsAbsorbPaused is a free data retrieval call binding the contract method 0x8d5d814c.
//
// Solidity: function isAbsorbPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceCaller) IsAbsorbPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "isAbsorbPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAbsorbPaused is a free data retrieval call binding the contract method 0x8d5d814c.
//
// Solidity: function isAbsorbPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceSession) IsAbsorbPaused() (bool, error) {
	return _CometMainInterface.Contract.IsAbsorbPaused(&_CometMainInterface.CallOpts)
}

// IsAbsorbPaused is a free data retrieval call binding the contract method 0x8d5d814c.
//
// Solidity: function isAbsorbPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceCallerSession) IsAbsorbPaused() (bool, error) {
	return _CometMainInterface.Contract.IsAbsorbPaused(&_CometMainInterface.CallOpts)
}

// IsBorrowCollateralized is a free data retrieval call binding the contract method 0x38aa813f.
//
// Solidity: function isBorrowCollateralized(address account) view returns(bool)
func (_CometMainInterface *CometMainInterfaceCaller) IsBorrowCollateralized(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "isBorrowCollateralized", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBorrowCollateralized is a free data retrieval call binding the contract method 0x38aa813f.
//
// Solidity: function isBorrowCollateralized(address account) view returns(bool)
func (_CometMainInterface *CometMainInterfaceSession) IsBorrowCollateralized(account common.Address) (bool, error) {
	return _CometMainInterface.Contract.IsBorrowCollateralized(&_CometMainInterface.CallOpts, account)
}

// IsBorrowCollateralized is a free data retrieval call binding the contract method 0x38aa813f.
//
// Solidity: function isBorrowCollateralized(address account) view returns(bool)
func (_CometMainInterface *CometMainInterfaceCallerSession) IsBorrowCollateralized(account common.Address) (bool, error) {
	return _CometMainInterface.Contract.IsBorrowCollateralized(&_CometMainInterface.CallOpts, account)
}

// IsBuyPaused is a free data retrieval call binding the contract method 0xd8e5f611.
//
// Solidity: function isBuyPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceCaller) IsBuyPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "isBuyPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBuyPaused is a free data retrieval call binding the contract method 0xd8e5f611.
//
// Solidity: function isBuyPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceSession) IsBuyPaused() (bool, error) {
	return _CometMainInterface.Contract.IsBuyPaused(&_CometMainInterface.CallOpts)
}

// IsBuyPaused is a free data retrieval call binding the contract method 0xd8e5f611.
//
// Solidity: function isBuyPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceCallerSession) IsBuyPaused() (bool, error) {
	return _CometMainInterface.Contract.IsBuyPaused(&_CometMainInterface.CallOpts)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address account) view returns(bool)
func (_CometMainInterface *CometMainInterfaceCaller) IsLiquidatable(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "isLiquidatable", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address account) view returns(bool)
func (_CometMainInterface *CometMainInterfaceSession) IsLiquidatable(account common.Address) (bool, error) {
	return _CometMainInterface.Contract.IsLiquidatable(&_CometMainInterface.CallOpts, account)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x042e02cf.
//
// Solidity: function isLiquidatable(address account) view returns(bool)
func (_CometMainInterface *CometMainInterfaceCallerSession) IsLiquidatable(account common.Address) (bool, error) {
	return _CometMainInterface.Contract.IsLiquidatable(&_CometMainInterface.CallOpts, account)
}

// IsSupplyPaused is a free data retrieval call binding the contract method 0x0bc47ad1.
//
// Solidity: function isSupplyPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceCaller) IsSupplyPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "isSupplyPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupplyPaused is a free data retrieval call binding the contract method 0x0bc47ad1.
//
// Solidity: function isSupplyPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceSession) IsSupplyPaused() (bool, error) {
	return _CometMainInterface.Contract.IsSupplyPaused(&_CometMainInterface.CallOpts)
}

// IsSupplyPaused is a free data retrieval call binding the contract method 0x0bc47ad1.
//
// Solidity: function isSupplyPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceCallerSession) IsSupplyPaused() (bool, error) {
	return _CometMainInterface.Contract.IsSupplyPaused(&_CometMainInterface.CallOpts)
}

// IsTransferPaused is a free data retrieval call binding the contract method 0xa1a1ef43.
//
// Solidity: function isTransferPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceCaller) IsTransferPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "isTransferPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransferPaused is a free data retrieval call binding the contract method 0xa1a1ef43.
//
// Solidity: function isTransferPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceSession) IsTransferPaused() (bool, error) {
	return _CometMainInterface.Contract.IsTransferPaused(&_CometMainInterface.CallOpts)
}

// IsTransferPaused is a free data retrieval call binding the contract method 0xa1a1ef43.
//
// Solidity: function isTransferPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceCallerSession) IsTransferPaused() (bool, error) {
	return _CometMainInterface.Contract.IsTransferPaused(&_CometMainInterface.CallOpts)
}

// IsWithdrawPaused is a free data retrieval call binding the contract method 0x67800b5f.
//
// Solidity: function isWithdrawPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceCaller) IsWithdrawPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "isWithdrawPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawPaused is a free data retrieval call binding the contract method 0x67800b5f.
//
// Solidity: function isWithdrawPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceSession) IsWithdrawPaused() (bool, error) {
	return _CometMainInterface.Contract.IsWithdrawPaused(&_CometMainInterface.CallOpts)
}

// IsWithdrawPaused is a free data retrieval call binding the contract method 0x67800b5f.
//
// Solidity: function isWithdrawPaused() view returns(bool)
func (_CometMainInterface *CometMainInterfaceCallerSession) IsWithdrawPaused() (bool, error) {
	return _CometMainInterface.Contract.IsWithdrawPaused(&_CometMainInterface.CallOpts)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint8)
func (_CometMainInterface *CometMainInterfaceCaller) NumAssets(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "numAssets")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint8)
func (_CometMainInterface *CometMainInterfaceSession) NumAssets() (uint8, error) {
	return _CometMainInterface.Contract.NumAssets(&_CometMainInterface.CallOpts)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint8)
func (_CometMainInterface *CometMainInterfaceCallerSession) NumAssets() (uint8, error) {
	return _CometMainInterface.Contract.NumAssets(&_CometMainInterface.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_CometMainInterface *CometMainInterfaceCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_CometMainInterface *CometMainInterfaceSession) PauseGuardian() (common.Address, error) {
	return _CometMainInterface.Contract.PauseGuardian(&_CometMainInterface.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_CometMainInterface *CometMainInterfaceCallerSession) PauseGuardian() (common.Address, error) {
	return _CometMainInterface.Contract.PauseGuardian(&_CometMainInterface.CallOpts)
}

// QuoteCollateral is a free data retrieval call binding the contract method 0x7ac88ed1.
//
// Solidity: function quoteCollateral(address asset, uint256 baseAmount) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) QuoteCollateral(opts *bind.CallOpts, asset common.Address, baseAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "quoteCollateral", asset, baseAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuoteCollateral is a free data retrieval call binding the contract method 0x7ac88ed1.
//
// Solidity: function quoteCollateral(address asset, uint256 baseAmount) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) QuoteCollateral(asset common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _CometMainInterface.Contract.QuoteCollateral(&_CometMainInterface.CallOpts, asset, baseAmount)
}

// QuoteCollateral is a free data retrieval call binding the contract method 0x7ac88ed1.
//
// Solidity: function quoteCollateral(address asset, uint256 baseAmount) view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) QuoteCollateral(asset common.Address, baseAmount *big.Int) (*big.Int, error) {
	return _CometMainInterface.Contract.QuoteCollateral(&_CometMainInterface.CallOpts, asset, baseAmount)
}

// StoreFrontPriceFactor is a free data retrieval call binding the contract method 0x1f5954bd.
//
// Solidity: function storeFrontPriceFactor() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) StoreFrontPriceFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "storeFrontPriceFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StoreFrontPriceFactor is a free data retrieval call binding the contract method 0x1f5954bd.
//
// Solidity: function storeFrontPriceFactor() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) StoreFrontPriceFactor() (*big.Int, error) {
	return _CometMainInterface.Contract.StoreFrontPriceFactor(&_CometMainInterface.CallOpts)
}

// StoreFrontPriceFactor is a free data retrieval call binding the contract method 0x1f5954bd.
//
// Solidity: function storeFrontPriceFactor() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) StoreFrontPriceFactor() (*big.Int, error) {
	return _CometMainInterface.Contract.StoreFrontPriceFactor(&_CometMainInterface.CallOpts)
}

// SupplyKink is a free data retrieval call binding the contract method 0xa5b4ff79.
//
// Solidity: function supplyKink() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) SupplyKink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "supplyKink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyKink is a free data retrieval call binding the contract method 0xa5b4ff79.
//
// Solidity: function supplyKink() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) SupplyKink() (*big.Int, error) {
	return _CometMainInterface.Contract.SupplyKink(&_CometMainInterface.CallOpts)
}

// SupplyKink is a free data retrieval call binding the contract method 0xa5b4ff79.
//
// Solidity: function supplyKink() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) SupplyKink() (*big.Int, error) {
	return _CometMainInterface.Contract.SupplyKink(&_CometMainInterface.CallOpts)
}

// SupplyPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x94920cca.
//
// Solidity: function supplyPerSecondInterestRateBase() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) SupplyPerSecondInterestRateBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "supplyPerSecondInterestRateBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x94920cca.
//
// Solidity: function supplyPerSecondInterestRateBase() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) SupplyPerSecondInterestRateBase() (*big.Int, error) {
	return _CometMainInterface.Contract.SupplyPerSecondInterestRateBase(&_CometMainInterface.CallOpts)
}

// SupplyPerSecondInterestRateBase is a free data retrieval call binding the contract method 0x94920cca.
//
// Solidity: function supplyPerSecondInterestRateBase() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) SupplyPerSecondInterestRateBase() (*big.Int, error) {
	return _CometMainInterface.Contract.SupplyPerSecondInterestRateBase(&_CometMainInterface.CallOpts)
}

// SupplyPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x804de71f.
//
// Solidity: function supplyPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) SupplyPerSecondInterestRateSlopeHigh(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "supplyPerSecondInterestRateSlopeHigh")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x804de71f.
//
// Solidity: function supplyPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) SupplyPerSecondInterestRateSlopeHigh() (*big.Int, error) {
	return _CometMainInterface.Contract.SupplyPerSecondInterestRateSlopeHigh(&_CometMainInterface.CallOpts)
}

// SupplyPerSecondInterestRateSlopeHigh is a free data retrieval call binding the contract method 0x804de71f.
//
// Solidity: function supplyPerSecondInterestRateSlopeHigh() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) SupplyPerSecondInterestRateSlopeHigh() (*big.Int, error) {
	return _CometMainInterface.Contract.SupplyPerSecondInterestRateSlopeHigh(&_CometMainInterface.CallOpts)
}

// SupplyPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x5a94b8d1.
//
// Solidity: function supplyPerSecondInterestRateSlopeLow() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) SupplyPerSecondInterestRateSlopeLow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "supplyPerSecondInterestRateSlopeLow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x5a94b8d1.
//
// Solidity: function supplyPerSecondInterestRateSlopeLow() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) SupplyPerSecondInterestRateSlopeLow() (*big.Int, error) {
	return _CometMainInterface.Contract.SupplyPerSecondInterestRateSlopeLow(&_CometMainInterface.CallOpts)
}

// SupplyPerSecondInterestRateSlopeLow is a free data retrieval call binding the contract method 0x5a94b8d1.
//
// Solidity: function supplyPerSecondInterestRateSlopeLow() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) SupplyPerSecondInterestRateSlopeLow() (*big.Int, error) {
	return _CometMainInterface.Contract.SupplyPerSecondInterestRateSlopeLow(&_CometMainInterface.CallOpts)
}

// TargetReserves is a free data retrieval call binding the contract method 0x32176c49.
//
// Solidity: function targetReserves() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) TargetReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "targetReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetReserves is a free data retrieval call binding the contract method 0x32176c49.
//
// Solidity: function targetReserves() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) TargetReserves() (*big.Int, error) {
	return _CometMainInterface.Contract.TargetReserves(&_CometMainInterface.CallOpts)
}

// TargetReserves is a free data retrieval call binding the contract method 0x32176c49.
//
// Solidity: function targetReserves() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) TargetReserves() (*big.Int, error) {
	return _CometMainInterface.Contract.TargetReserves(&_CometMainInterface.CallOpts)
}

// TotalBorrow is a free data retrieval call binding the contract method 0x8285ef40.
//
// Solidity: function totalBorrow() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) TotalBorrow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "totalBorrow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrow is a free data retrieval call binding the contract method 0x8285ef40.
//
// Solidity: function totalBorrow() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) TotalBorrow() (*big.Int, error) {
	return _CometMainInterface.Contract.TotalBorrow(&_CometMainInterface.CallOpts)
}

// TotalBorrow is a free data retrieval call binding the contract method 0x8285ef40.
//
// Solidity: function totalBorrow() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) TotalBorrow() (*big.Int, error) {
	return _CometMainInterface.Contract.TotalBorrow(&_CometMainInterface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) TotalSupply() (*big.Int, error) {
	return _CometMainInterface.Contract.TotalSupply(&_CometMainInterface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _CometMainInterface.Contract.TotalSupply(&_CometMainInterface.CallOpts)
}

// TrackingIndexScale is a free data retrieval call binding the contract method 0xaba7f15e.
//
// Solidity: function trackingIndexScale() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCaller) TrackingIndexScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CometMainInterface.contract.Call(opts, &out, "trackingIndexScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrackingIndexScale is a free data retrieval call binding the contract method 0xaba7f15e.
//
// Solidity: function trackingIndexScale() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceSession) TrackingIndexScale() (*big.Int, error) {
	return _CometMainInterface.Contract.TrackingIndexScale(&_CometMainInterface.CallOpts)
}

// TrackingIndexScale is a free data retrieval call binding the contract method 0xaba7f15e.
//
// Solidity: function trackingIndexScale() view returns(uint256)
func (_CometMainInterface *CometMainInterfaceCallerSession) TrackingIndexScale() (*big.Int, error) {
	return _CometMainInterface.Contract.TrackingIndexScale(&_CometMainInterface.CallOpts)
}

// Absorb is a paid mutator transaction binding the contract method 0xc3cecfd2.
//
// Solidity: function absorb(address absorber, address[] accounts) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) Absorb(opts *bind.TransactOpts, absorber common.Address, accounts []common.Address) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "absorb", absorber, accounts)
}

// Absorb is a paid mutator transaction binding the contract method 0xc3cecfd2.
//
// Solidity: function absorb(address absorber, address[] accounts) returns()
func (_CometMainInterface *CometMainInterfaceSession) Absorb(absorber common.Address, accounts []common.Address) (*types.Transaction, error) {
	return _CometMainInterface.Contract.Absorb(&_CometMainInterface.TransactOpts, absorber, accounts)
}

// Absorb is a paid mutator transaction binding the contract method 0xc3cecfd2.
//
// Solidity: function absorb(address absorber, address[] accounts) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) Absorb(absorber common.Address, accounts []common.Address) (*types.Transaction, error) {
	return _CometMainInterface.Contract.Absorb(&_CometMainInterface.TransactOpts, absorber, accounts)
}

// AccrueAccount is a paid mutator transaction binding the contract method 0xbfe69c8d.
//
// Solidity: function accrueAccount(address account) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) AccrueAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "accrueAccount", account)
}

// AccrueAccount is a paid mutator transaction binding the contract method 0xbfe69c8d.
//
// Solidity: function accrueAccount(address account) returns()
func (_CometMainInterface *CometMainInterfaceSession) AccrueAccount(account common.Address) (*types.Transaction, error) {
	return _CometMainInterface.Contract.AccrueAccount(&_CometMainInterface.TransactOpts, account)
}

// AccrueAccount is a paid mutator transaction binding the contract method 0xbfe69c8d.
//
// Solidity: function accrueAccount(address account) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) AccrueAccount(account common.Address) (*types.Transaction, error) {
	return _CometMainInterface.Contract.AccrueAccount(&_CometMainInterface.TransactOpts, account)
}

// ApproveThis is a paid mutator transaction binding the contract method 0xad14777c.
//
// Solidity: function approveThis(address manager, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) ApproveThis(opts *bind.TransactOpts, manager common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "approveThis", manager, asset, amount)
}

// ApproveThis is a paid mutator transaction binding the contract method 0xad14777c.
//
// Solidity: function approveThis(address manager, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceSession) ApproveThis(manager common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.ApproveThis(&_CometMainInterface.TransactOpts, manager, asset, amount)
}

// ApproveThis is a paid mutator transaction binding the contract method 0xad14777c.
//
// Solidity: function approveThis(address manager, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) ApproveThis(manager common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.ApproveThis(&_CometMainInterface.TransactOpts, manager, asset, amount)
}

// BuyCollateral is a paid mutator transaction binding the contract method 0xe4e6e779.
//
// Solidity: function buyCollateral(address asset, uint256 minAmount, uint256 baseAmount, address recipient) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) BuyCollateral(opts *bind.TransactOpts, asset common.Address, minAmount *big.Int, baseAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "buyCollateral", asset, minAmount, baseAmount, recipient)
}

// BuyCollateral is a paid mutator transaction binding the contract method 0xe4e6e779.
//
// Solidity: function buyCollateral(address asset, uint256 minAmount, uint256 baseAmount, address recipient) returns()
func (_CometMainInterface *CometMainInterfaceSession) BuyCollateral(asset common.Address, minAmount *big.Int, baseAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CometMainInterface.Contract.BuyCollateral(&_CometMainInterface.TransactOpts, asset, minAmount, baseAmount, recipient)
}

// BuyCollateral is a paid mutator transaction binding the contract method 0xe4e6e779.
//
// Solidity: function buyCollateral(address asset, uint256 minAmount, uint256 baseAmount, address recipient) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) BuyCollateral(asset common.Address, minAmount *big.Int, baseAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _CometMainInterface.Contract.BuyCollateral(&_CometMainInterface.TransactOpts, asset, minAmount, baseAmount, recipient)
}

// InitializeStorage is a paid mutator transaction binding the contract method 0x1c9f7fb9.
//
// Solidity: function initializeStorage() returns()
func (_CometMainInterface *CometMainInterfaceTransactor) InitializeStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "initializeStorage")
}

// InitializeStorage is a paid mutator transaction binding the contract method 0x1c9f7fb9.
//
// Solidity: function initializeStorage() returns()
func (_CometMainInterface *CometMainInterfaceSession) InitializeStorage() (*types.Transaction, error) {
	return _CometMainInterface.Contract.InitializeStorage(&_CometMainInterface.TransactOpts)
}

// InitializeStorage is a paid mutator transaction binding the contract method 0x1c9f7fb9.
//
// Solidity: function initializeStorage() returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) InitializeStorage() (*types.Transaction, error) {
	return _CometMainInterface.Contract.InitializeStorage(&_CometMainInterface.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x44c35d07.
//
// Solidity: function pause(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) Pause(opts *bind.TransactOpts, supplyPaused bool, transferPaused bool, withdrawPaused bool, absorbPaused bool, buyPaused bool) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "pause", supplyPaused, transferPaused, withdrawPaused, absorbPaused, buyPaused)
}

// Pause is a paid mutator transaction binding the contract method 0x44c35d07.
//
// Solidity: function pause(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused) returns()
func (_CometMainInterface *CometMainInterfaceSession) Pause(supplyPaused bool, transferPaused bool, withdrawPaused bool, absorbPaused bool, buyPaused bool) (*types.Transaction, error) {
	return _CometMainInterface.Contract.Pause(&_CometMainInterface.TransactOpts, supplyPaused, transferPaused, withdrawPaused, absorbPaused, buyPaused)
}

// Pause is a paid mutator transaction binding the contract method 0x44c35d07.
//
// Solidity: function pause(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) Pause(supplyPaused bool, transferPaused bool, withdrawPaused bool, absorbPaused bool, buyPaused bool) (*types.Transaction, error) {
	return _CometMainInterface.Contract.Pause(&_CometMainInterface.TransactOpts, supplyPaused, transferPaused, withdrawPaused, absorbPaused, buyPaused)
}

// Supply is a paid mutator transaction binding the contract method 0xf2b9fdb8.
//
// Solidity: function supply(address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) Supply(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "supply", asset, amount)
}

// Supply is a paid mutator transaction binding the contract method 0xf2b9fdb8.
//
// Solidity: function supply(address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceSession) Supply(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.Supply(&_CometMainInterface.TransactOpts, asset, amount)
}

// Supply is a paid mutator transaction binding the contract method 0xf2b9fdb8.
//
// Solidity: function supply(address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) Supply(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.Supply(&_CometMainInterface.TransactOpts, asset, amount)
}

// SupplyFrom is a paid mutator transaction binding the contract method 0x90323177.
//
// Solidity: function supplyFrom(address from, address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) SupplyFrom(opts *bind.TransactOpts, from common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "supplyFrom", from, dst, asset, amount)
}

// SupplyFrom is a paid mutator transaction binding the contract method 0x90323177.
//
// Solidity: function supplyFrom(address from, address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceSession) SupplyFrom(from common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.SupplyFrom(&_CometMainInterface.TransactOpts, from, dst, asset, amount)
}

// SupplyFrom is a paid mutator transaction binding the contract method 0x90323177.
//
// Solidity: function supplyFrom(address from, address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) SupplyFrom(from common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.SupplyFrom(&_CometMainInterface.TransactOpts, from, dst, asset, amount)
}

// SupplyTo is a paid mutator transaction binding the contract method 0x4232cd63.
//
// Solidity: function supplyTo(address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) SupplyTo(opts *bind.TransactOpts, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "supplyTo", dst, asset, amount)
}

// SupplyTo is a paid mutator transaction binding the contract method 0x4232cd63.
//
// Solidity: function supplyTo(address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceSession) SupplyTo(dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.SupplyTo(&_CometMainInterface.TransactOpts, dst, asset, amount)
}

// SupplyTo is a paid mutator transaction binding the contract method 0x4232cd63.
//
// Solidity: function supplyTo(address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) SupplyTo(dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.SupplyTo(&_CometMainInterface.TransactOpts, dst, asset, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CometMainInterface *CometMainInterfaceTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CometMainInterface *CometMainInterfaceSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.Transfer(&_CometMainInterface.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CometMainInterface *CometMainInterfaceTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.Transfer(&_CometMainInterface.TransactOpts, dst, amount)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x439e2e45.
//
// Solidity: function transferAsset(address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) TransferAsset(opts *bind.TransactOpts, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "transferAsset", dst, asset, amount)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x439e2e45.
//
// Solidity: function transferAsset(address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceSession) TransferAsset(dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.TransferAsset(&_CometMainInterface.TransactOpts, dst, asset, amount)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x439e2e45.
//
// Solidity: function transferAsset(address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) TransferAsset(dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.TransferAsset(&_CometMainInterface.TransactOpts, dst, asset, amount)
}

// TransferAssetFrom is a paid mutator transaction binding the contract method 0xc1ee2c18.
//
// Solidity: function transferAssetFrom(address src, address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) TransferAssetFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "transferAssetFrom", src, dst, asset, amount)
}

// TransferAssetFrom is a paid mutator transaction binding the contract method 0xc1ee2c18.
//
// Solidity: function transferAssetFrom(address src, address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceSession) TransferAssetFrom(src common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.TransferAssetFrom(&_CometMainInterface.TransactOpts, src, dst, asset, amount)
}

// TransferAssetFrom is a paid mutator transaction binding the contract method 0xc1ee2c18.
//
// Solidity: function transferAssetFrom(address src, address dst, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) TransferAssetFrom(src common.Address, dst common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.TransferAssetFrom(&_CometMainInterface.TransactOpts, src, dst, asset, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CometMainInterface *CometMainInterfaceTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CometMainInterface *CometMainInterfaceSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.TransferFrom(&_CometMainInterface.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CometMainInterface *CometMainInterfaceTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.TransferFrom(&_CometMainInterface.TransactOpts, src, dst, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "withdraw", asset, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceSession) Withdraw(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.Withdraw(&_CometMainInterface.TransactOpts, asset, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) Withdraw(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.Withdraw(&_CometMainInterface.TransactOpts, asset, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x26441318.
//
// Solidity: function withdrawFrom(address src, address to, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) WithdrawFrom(opts *bind.TransactOpts, src common.Address, to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "withdrawFrom", src, to, asset, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x26441318.
//
// Solidity: function withdrawFrom(address src, address to, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceSession) WithdrawFrom(src common.Address, to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.WithdrawFrom(&_CometMainInterface.TransactOpts, src, to, asset, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x26441318.
//
// Solidity: function withdrawFrom(address src, address to, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) WithdrawFrom(src common.Address, to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.WithdrawFrom(&_CometMainInterface.TransactOpts, src, to, asset, amount)
}

// WithdrawReserves is a paid mutator transaction binding the contract method 0xe478795d.
//
// Solidity: function withdrawReserves(address to, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) WithdrawReserves(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "withdrawReserves", to, amount)
}

// WithdrawReserves is a paid mutator transaction binding the contract method 0xe478795d.
//
// Solidity: function withdrawReserves(address to, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceSession) WithdrawReserves(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.WithdrawReserves(&_CometMainInterface.TransactOpts, to, amount)
}

// WithdrawReserves is a paid mutator transaction binding the contract method 0xe478795d.
//
// Solidity: function withdrawReserves(address to, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) WithdrawReserves(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.WithdrawReserves(&_CometMainInterface.TransactOpts, to, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address to, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactor) WithdrawTo(opts *bind.TransactOpts, to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.contract.Transact(opts, "withdrawTo", to, asset, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address to, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceSession) WithdrawTo(to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.WithdrawTo(&_CometMainInterface.TransactOpts, to, asset, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0xc3b35a7e.
//
// Solidity: function withdrawTo(address to, address asset, uint256 amount) returns()
func (_CometMainInterface *CometMainInterfaceTransactorSession) WithdrawTo(to common.Address, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CometMainInterface.Contract.WithdrawTo(&_CometMainInterface.TransactOpts, to, asset, amount)
}

// CometMainInterfaceAbsorbCollateralIterator is returned from FilterAbsorbCollateral and is used to iterate over the raw logs and unpacked data for AbsorbCollateral events raised by the CometMainInterface contract.
type CometMainInterfaceAbsorbCollateralIterator struct {
	Event *CometMainInterfaceAbsorbCollateral // Event containing the contract specifics and raw log

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
func (it *CometMainInterfaceAbsorbCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometMainInterfaceAbsorbCollateral)
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
		it.Event = new(CometMainInterfaceAbsorbCollateral)
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
func (it *CometMainInterfaceAbsorbCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometMainInterfaceAbsorbCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometMainInterfaceAbsorbCollateral represents a AbsorbCollateral event raised by the CometMainInterface contract.
type CometMainInterfaceAbsorbCollateral struct {
	Absorber           common.Address
	Borrower           common.Address
	Asset              common.Address
	CollateralAbsorbed *big.Int
	UsdValue           *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAbsorbCollateral is a free log retrieval operation binding the contract event 0x9850ab1af75177e4a9201c65a2cf7976d5d28e40ef63494b44366f86b2f9412e.
//
// Solidity: event AbsorbCollateral(address indexed absorber, address indexed borrower, address indexed asset, uint256 collateralAbsorbed, uint256 usdValue)
func (_CometMainInterface *CometMainInterfaceFilterer) FilterAbsorbCollateral(opts *bind.FilterOpts, absorber []common.Address, borrower []common.Address, asset []common.Address) (*CometMainInterfaceAbsorbCollateralIterator, error) {

	var absorberRule []interface{}
	for _, absorberItem := range absorber {
		absorberRule = append(absorberRule, absorberItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CometMainInterface.contract.FilterLogs(opts, "AbsorbCollateral", absorberRule, borrowerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceAbsorbCollateralIterator{contract: _CometMainInterface.contract, event: "AbsorbCollateral", logs: logs, sub: sub}, nil
}

// WatchAbsorbCollateral is a free log subscription operation binding the contract event 0x9850ab1af75177e4a9201c65a2cf7976d5d28e40ef63494b44366f86b2f9412e.
//
// Solidity: event AbsorbCollateral(address indexed absorber, address indexed borrower, address indexed asset, uint256 collateralAbsorbed, uint256 usdValue)
func (_CometMainInterface *CometMainInterfaceFilterer) WatchAbsorbCollateral(opts *bind.WatchOpts, sink chan<- *CometMainInterfaceAbsorbCollateral, absorber []common.Address, borrower []common.Address, asset []common.Address) (event.Subscription, error) {

	var absorberRule []interface{}
	for _, absorberItem := range absorber {
		absorberRule = append(absorberRule, absorberItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CometMainInterface.contract.WatchLogs(opts, "AbsorbCollateral", absorberRule, borrowerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometMainInterfaceAbsorbCollateral)
				if err := _CometMainInterface.contract.UnpackLog(event, "AbsorbCollateral", log); err != nil {
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

// ParseAbsorbCollateral is a log parse operation binding the contract event 0x9850ab1af75177e4a9201c65a2cf7976d5d28e40ef63494b44366f86b2f9412e.
//
// Solidity: event AbsorbCollateral(address indexed absorber, address indexed borrower, address indexed asset, uint256 collateralAbsorbed, uint256 usdValue)
func (_CometMainInterface *CometMainInterfaceFilterer) ParseAbsorbCollateral(log types.Log) (*CometMainInterfaceAbsorbCollateral, error) {
	event := new(CometMainInterfaceAbsorbCollateral)
	if err := _CometMainInterface.contract.UnpackLog(event, "AbsorbCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometMainInterfaceAbsorbDebtIterator is returned from FilterAbsorbDebt and is used to iterate over the raw logs and unpacked data for AbsorbDebt events raised by the CometMainInterface contract.
type CometMainInterfaceAbsorbDebtIterator struct {
	Event *CometMainInterfaceAbsorbDebt // Event containing the contract specifics and raw log

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
func (it *CometMainInterfaceAbsorbDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometMainInterfaceAbsorbDebt)
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
		it.Event = new(CometMainInterfaceAbsorbDebt)
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
func (it *CometMainInterfaceAbsorbDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometMainInterfaceAbsorbDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometMainInterfaceAbsorbDebt represents a AbsorbDebt event raised by the CometMainInterface contract.
type CometMainInterfaceAbsorbDebt struct {
	Absorber    common.Address
	Borrower    common.Address
	BasePaidOut *big.Int
	UsdValue    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAbsorbDebt is a free log retrieval operation binding the contract event 0x1547a878dc89ad3c367b6338b4be6a65a5dd74fb77ae044da1e8747ef1f4f62f.
//
// Solidity: event AbsorbDebt(address indexed absorber, address indexed borrower, uint256 basePaidOut, uint256 usdValue)
func (_CometMainInterface *CometMainInterfaceFilterer) FilterAbsorbDebt(opts *bind.FilterOpts, absorber []common.Address, borrower []common.Address) (*CometMainInterfaceAbsorbDebtIterator, error) {

	var absorberRule []interface{}
	for _, absorberItem := range absorber {
		absorberRule = append(absorberRule, absorberItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CometMainInterface.contract.FilterLogs(opts, "AbsorbDebt", absorberRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceAbsorbDebtIterator{contract: _CometMainInterface.contract, event: "AbsorbDebt", logs: logs, sub: sub}, nil
}

// WatchAbsorbDebt is a free log subscription operation binding the contract event 0x1547a878dc89ad3c367b6338b4be6a65a5dd74fb77ae044da1e8747ef1f4f62f.
//
// Solidity: event AbsorbDebt(address indexed absorber, address indexed borrower, uint256 basePaidOut, uint256 usdValue)
func (_CometMainInterface *CometMainInterfaceFilterer) WatchAbsorbDebt(opts *bind.WatchOpts, sink chan<- *CometMainInterfaceAbsorbDebt, absorber []common.Address, borrower []common.Address) (event.Subscription, error) {

	var absorberRule []interface{}
	for _, absorberItem := range absorber {
		absorberRule = append(absorberRule, absorberItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CometMainInterface.contract.WatchLogs(opts, "AbsorbDebt", absorberRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometMainInterfaceAbsorbDebt)
				if err := _CometMainInterface.contract.UnpackLog(event, "AbsorbDebt", log); err != nil {
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

// ParseAbsorbDebt is a log parse operation binding the contract event 0x1547a878dc89ad3c367b6338b4be6a65a5dd74fb77ae044da1e8747ef1f4f62f.
//
// Solidity: event AbsorbDebt(address indexed absorber, address indexed borrower, uint256 basePaidOut, uint256 usdValue)
func (_CometMainInterface *CometMainInterfaceFilterer) ParseAbsorbDebt(log types.Log) (*CometMainInterfaceAbsorbDebt, error) {
	event := new(CometMainInterfaceAbsorbDebt)
	if err := _CometMainInterface.contract.UnpackLog(event, "AbsorbDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometMainInterfaceBuyCollateralIterator is returned from FilterBuyCollateral and is used to iterate over the raw logs and unpacked data for BuyCollateral events raised by the CometMainInterface contract.
type CometMainInterfaceBuyCollateralIterator struct {
	Event *CometMainInterfaceBuyCollateral // Event containing the contract specifics and raw log

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
func (it *CometMainInterfaceBuyCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometMainInterfaceBuyCollateral)
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
		it.Event = new(CometMainInterfaceBuyCollateral)
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
func (it *CometMainInterfaceBuyCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometMainInterfaceBuyCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometMainInterfaceBuyCollateral represents a BuyCollateral event raised by the CometMainInterface contract.
type CometMainInterfaceBuyCollateral struct {
	Buyer            common.Address
	Asset            common.Address
	BaseAmount       *big.Int
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBuyCollateral is a free log retrieval operation binding the contract event 0xf891b2a411b0e66a5f0a6ff1368670fefa287a13f541eb633a386a1a9cc7046b.
//
// Solidity: event BuyCollateral(address indexed buyer, address indexed asset, uint256 baseAmount, uint256 collateralAmount)
func (_CometMainInterface *CometMainInterfaceFilterer) FilterBuyCollateral(opts *bind.FilterOpts, buyer []common.Address, asset []common.Address) (*CometMainInterfaceBuyCollateralIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CometMainInterface.contract.FilterLogs(opts, "BuyCollateral", buyerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceBuyCollateralIterator{contract: _CometMainInterface.contract, event: "BuyCollateral", logs: logs, sub: sub}, nil
}

// WatchBuyCollateral is a free log subscription operation binding the contract event 0xf891b2a411b0e66a5f0a6ff1368670fefa287a13f541eb633a386a1a9cc7046b.
//
// Solidity: event BuyCollateral(address indexed buyer, address indexed asset, uint256 baseAmount, uint256 collateralAmount)
func (_CometMainInterface *CometMainInterfaceFilterer) WatchBuyCollateral(opts *bind.WatchOpts, sink chan<- *CometMainInterfaceBuyCollateral, buyer []common.Address, asset []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CometMainInterface.contract.WatchLogs(opts, "BuyCollateral", buyerRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometMainInterfaceBuyCollateral)
				if err := _CometMainInterface.contract.UnpackLog(event, "BuyCollateral", log); err != nil {
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

// ParseBuyCollateral is a log parse operation binding the contract event 0xf891b2a411b0e66a5f0a6ff1368670fefa287a13f541eb633a386a1a9cc7046b.
//
// Solidity: event BuyCollateral(address indexed buyer, address indexed asset, uint256 baseAmount, uint256 collateralAmount)
func (_CometMainInterface *CometMainInterfaceFilterer) ParseBuyCollateral(log types.Log) (*CometMainInterfaceBuyCollateral, error) {
	event := new(CometMainInterfaceBuyCollateral)
	if err := _CometMainInterface.contract.UnpackLog(event, "BuyCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometMainInterfacePauseActionIterator is returned from FilterPauseAction and is used to iterate over the raw logs and unpacked data for PauseAction events raised by the CometMainInterface contract.
type CometMainInterfacePauseActionIterator struct {
	Event *CometMainInterfacePauseAction // Event containing the contract specifics and raw log

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
func (it *CometMainInterfacePauseActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometMainInterfacePauseAction)
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
		it.Event = new(CometMainInterfacePauseAction)
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
func (it *CometMainInterfacePauseActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometMainInterfacePauseActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometMainInterfacePauseAction represents a PauseAction event raised by the CometMainInterface contract.
type CometMainInterfacePauseAction struct {
	SupplyPaused   bool
	TransferPaused bool
	WithdrawPaused bool
	AbsorbPaused   bool
	BuyPaused      bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPauseAction is a free log retrieval operation binding the contract event 0x3be39979091ae7ca962aa1c44e645f2df3c221b79f324afa5f44aedc8d2f690d.
//
// Solidity: event PauseAction(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused)
func (_CometMainInterface *CometMainInterfaceFilterer) FilterPauseAction(opts *bind.FilterOpts) (*CometMainInterfacePauseActionIterator, error) {

	logs, sub, err := _CometMainInterface.contract.FilterLogs(opts, "PauseAction")
	if err != nil {
		return nil, err
	}
	return &CometMainInterfacePauseActionIterator{contract: _CometMainInterface.contract, event: "PauseAction", logs: logs, sub: sub}, nil
}

// WatchPauseAction is a free log subscription operation binding the contract event 0x3be39979091ae7ca962aa1c44e645f2df3c221b79f324afa5f44aedc8d2f690d.
//
// Solidity: event PauseAction(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused)
func (_CometMainInterface *CometMainInterfaceFilterer) WatchPauseAction(opts *bind.WatchOpts, sink chan<- *CometMainInterfacePauseAction) (event.Subscription, error) {

	logs, sub, err := _CometMainInterface.contract.WatchLogs(opts, "PauseAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometMainInterfacePauseAction)
				if err := _CometMainInterface.contract.UnpackLog(event, "PauseAction", log); err != nil {
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

// ParsePauseAction is a log parse operation binding the contract event 0x3be39979091ae7ca962aa1c44e645f2df3c221b79f324afa5f44aedc8d2f690d.
//
// Solidity: event PauseAction(bool supplyPaused, bool transferPaused, bool withdrawPaused, bool absorbPaused, bool buyPaused)
func (_CometMainInterface *CometMainInterfaceFilterer) ParsePauseAction(log types.Log) (*CometMainInterfacePauseAction, error) {
	event := new(CometMainInterfacePauseAction)
	if err := _CometMainInterface.contract.UnpackLog(event, "PauseAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometMainInterfaceSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the CometMainInterface contract.
type CometMainInterfaceSupplyIterator struct {
	Event *CometMainInterfaceSupply // Event containing the contract specifics and raw log

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
func (it *CometMainInterfaceSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometMainInterfaceSupply)
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
		it.Event = new(CometMainInterfaceSupply)
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
func (it *CometMainInterfaceSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometMainInterfaceSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometMainInterfaceSupply represents a Supply event raised by the CometMainInterface contract.
type CometMainInterfaceSupply struct {
	From   common.Address
	Dst    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0xd1cf3d156d5f8f0d50f6c122ed609cec09d35c9b9fb3fff6ea0959134dae424e.
//
// Solidity: event Supply(address indexed from, address indexed dst, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) FilterSupply(opts *bind.FilterOpts, from []common.Address, dst []common.Address) (*CometMainInterfaceSupplyIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _CometMainInterface.contract.FilterLogs(opts, "Supply", fromRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceSupplyIterator{contract: _CometMainInterface.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0xd1cf3d156d5f8f0d50f6c122ed609cec09d35c9b9fb3fff6ea0959134dae424e.
//
// Solidity: event Supply(address indexed from, address indexed dst, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *CometMainInterfaceSupply, from []common.Address, dst []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _CometMainInterface.contract.WatchLogs(opts, "Supply", fromRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometMainInterfaceSupply)
				if err := _CometMainInterface.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0xd1cf3d156d5f8f0d50f6c122ed609cec09d35c9b9fb3fff6ea0959134dae424e.
//
// Solidity: event Supply(address indexed from, address indexed dst, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) ParseSupply(log types.Log) (*CometMainInterfaceSupply, error) {
	event := new(CometMainInterfaceSupply)
	if err := _CometMainInterface.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometMainInterfaceSupplyCollateralIterator is returned from FilterSupplyCollateral and is used to iterate over the raw logs and unpacked data for SupplyCollateral events raised by the CometMainInterface contract.
type CometMainInterfaceSupplyCollateralIterator struct {
	Event *CometMainInterfaceSupplyCollateral // Event containing the contract specifics and raw log

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
func (it *CometMainInterfaceSupplyCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometMainInterfaceSupplyCollateral)
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
		it.Event = new(CometMainInterfaceSupplyCollateral)
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
func (it *CometMainInterfaceSupplyCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometMainInterfaceSupplyCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometMainInterfaceSupplyCollateral represents a SupplyCollateral event raised by the CometMainInterface contract.
type CometMainInterfaceSupplyCollateral struct {
	From   common.Address
	Dst    common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSupplyCollateral is a free log retrieval operation binding the contract event 0xfa56f7b24f17183d81894d3ac2ee654e3c26388d17a28dbd9549b8114304e1f4.
//
// Solidity: event SupplyCollateral(address indexed from, address indexed dst, address indexed asset, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) FilterSupplyCollateral(opts *bind.FilterOpts, from []common.Address, dst []common.Address, asset []common.Address) (*CometMainInterfaceSupplyCollateralIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CometMainInterface.contract.FilterLogs(opts, "SupplyCollateral", fromRule, dstRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceSupplyCollateralIterator{contract: _CometMainInterface.contract, event: "SupplyCollateral", logs: logs, sub: sub}, nil
}

// WatchSupplyCollateral is a free log subscription operation binding the contract event 0xfa56f7b24f17183d81894d3ac2ee654e3c26388d17a28dbd9549b8114304e1f4.
//
// Solidity: event SupplyCollateral(address indexed from, address indexed dst, address indexed asset, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) WatchSupplyCollateral(opts *bind.WatchOpts, sink chan<- *CometMainInterfaceSupplyCollateral, from []common.Address, dst []common.Address, asset []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CometMainInterface.contract.WatchLogs(opts, "SupplyCollateral", fromRule, dstRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometMainInterfaceSupplyCollateral)
				if err := _CometMainInterface.contract.UnpackLog(event, "SupplyCollateral", log); err != nil {
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

// ParseSupplyCollateral is a log parse operation binding the contract event 0xfa56f7b24f17183d81894d3ac2ee654e3c26388d17a28dbd9549b8114304e1f4.
//
// Solidity: event SupplyCollateral(address indexed from, address indexed dst, address indexed asset, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) ParseSupplyCollateral(log types.Log) (*CometMainInterfaceSupplyCollateral, error) {
	event := new(CometMainInterfaceSupplyCollateral)
	if err := _CometMainInterface.contract.UnpackLog(event, "SupplyCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometMainInterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CometMainInterface contract.
type CometMainInterfaceTransferIterator struct {
	Event *CometMainInterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *CometMainInterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometMainInterfaceTransfer)
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
		it.Event = new(CometMainInterfaceTransfer)
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
func (it *CometMainInterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometMainInterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometMainInterfaceTransfer represents a Transfer event raised by the CometMainInterface contract.
type CometMainInterfaceTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CometMainInterfaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CometMainInterface.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceTransferIterator{contract: _CometMainInterface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CometMainInterfaceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CometMainInterface.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometMainInterfaceTransfer)
				if err := _CometMainInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) ParseTransfer(log types.Log) (*CometMainInterfaceTransfer, error) {
	event := new(CometMainInterfaceTransfer)
	if err := _CometMainInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometMainInterfaceTransferCollateralIterator is returned from FilterTransferCollateral and is used to iterate over the raw logs and unpacked data for TransferCollateral events raised by the CometMainInterface contract.
type CometMainInterfaceTransferCollateralIterator struct {
	Event *CometMainInterfaceTransferCollateral // Event containing the contract specifics and raw log

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
func (it *CometMainInterfaceTransferCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometMainInterfaceTransferCollateral)
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
		it.Event = new(CometMainInterfaceTransferCollateral)
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
func (it *CometMainInterfaceTransferCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometMainInterfaceTransferCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometMainInterfaceTransferCollateral represents a TransferCollateral event raised by the CometMainInterface contract.
type CometMainInterfaceTransferCollateral struct {
	From   common.Address
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferCollateral is a free log retrieval operation binding the contract event 0x29db89d45e1a802b4d55e202984fce9faf1d30aedf86503ff1ea0ed9ebb64201.
//
// Solidity: event TransferCollateral(address indexed from, address indexed to, address indexed asset, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) FilterTransferCollateral(opts *bind.FilterOpts, from []common.Address, to []common.Address, asset []common.Address) (*CometMainInterfaceTransferCollateralIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CometMainInterface.contract.FilterLogs(opts, "TransferCollateral", fromRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceTransferCollateralIterator{contract: _CometMainInterface.contract, event: "TransferCollateral", logs: logs, sub: sub}, nil
}

// WatchTransferCollateral is a free log subscription operation binding the contract event 0x29db89d45e1a802b4d55e202984fce9faf1d30aedf86503ff1ea0ed9ebb64201.
//
// Solidity: event TransferCollateral(address indexed from, address indexed to, address indexed asset, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) WatchTransferCollateral(opts *bind.WatchOpts, sink chan<- *CometMainInterfaceTransferCollateral, from []common.Address, to []common.Address, asset []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CometMainInterface.contract.WatchLogs(opts, "TransferCollateral", fromRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometMainInterfaceTransferCollateral)
				if err := _CometMainInterface.contract.UnpackLog(event, "TransferCollateral", log); err != nil {
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

// ParseTransferCollateral is a log parse operation binding the contract event 0x29db89d45e1a802b4d55e202984fce9faf1d30aedf86503ff1ea0ed9ebb64201.
//
// Solidity: event TransferCollateral(address indexed from, address indexed to, address indexed asset, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) ParseTransferCollateral(log types.Log) (*CometMainInterfaceTransferCollateral, error) {
	event := new(CometMainInterfaceTransferCollateral)
	if err := _CometMainInterface.contract.UnpackLog(event, "TransferCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometMainInterfaceWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the CometMainInterface contract.
type CometMainInterfaceWithdrawIterator struct {
	Event *CometMainInterfaceWithdraw // Event containing the contract specifics and raw log

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
func (it *CometMainInterfaceWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometMainInterfaceWithdraw)
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
		it.Event = new(CometMainInterfaceWithdraw)
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
func (it *CometMainInterfaceWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometMainInterfaceWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometMainInterfaceWithdraw represents a Withdraw event raised by the CometMainInterface contract.
type CometMainInterfaceWithdraw struct {
	Src    common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed src, address indexed to, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) FilterWithdraw(opts *bind.FilterOpts, src []common.Address, to []common.Address) (*CometMainInterfaceWithdrawIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CometMainInterface.contract.FilterLogs(opts, "Withdraw", srcRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceWithdrawIterator{contract: _CometMainInterface.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed src, address indexed to, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CometMainInterfaceWithdraw, src []common.Address, to []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CometMainInterface.contract.WatchLogs(opts, "Withdraw", srcRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometMainInterfaceWithdraw)
				if err := _CometMainInterface.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed src, address indexed to, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) ParseWithdraw(log types.Log) (*CometMainInterfaceWithdraw, error) {
	event := new(CometMainInterfaceWithdraw)
	if err := _CometMainInterface.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometMainInterfaceWithdrawCollateralIterator is returned from FilterWithdrawCollateral and is used to iterate over the raw logs and unpacked data for WithdrawCollateral events raised by the CometMainInterface contract.
type CometMainInterfaceWithdrawCollateralIterator struct {
	Event *CometMainInterfaceWithdrawCollateral // Event containing the contract specifics and raw log

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
func (it *CometMainInterfaceWithdrawCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometMainInterfaceWithdrawCollateral)
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
		it.Event = new(CometMainInterfaceWithdrawCollateral)
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
func (it *CometMainInterfaceWithdrawCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometMainInterfaceWithdrawCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometMainInterfaceWithdrawCollateral represents a WithdrawCollateral event raised by the CometMainInterface contract.
type CometMainInterfaceWithdrawCollateral struct {
	Src    common.Address
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCollateral is a free log retrieval operation binding the contract event 0xd6d480d5b3068db003533b170d67561494d72e3bf9fa40a266471351ebba9e16.
//
// Solidity: event WithdrawCollateral(address indexed src, address indexed to, address indexed asset, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) FilterWithdrawCollateral(opts *bind.FilterOpts, src []common.Address, to []common.Address, asset []common.Address) (*CometMainInterfaceWithdrawCollateralIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CometMainInterface.contract.FilterLogs(opts, "WithdrawCollateral", srcRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceWithdrawCollateralIterator{contract: _CometMainInterface.contract, event: "WithdrawCollateral", logs: logs, sub: sub}, nil
}

// WatchWithdrawCollateral is a free log subscription operation binding the contract event 0xd6d480d5b3068db003533b170d67561494d72e3bf9fa40a266471351ebba9e16.
//
// Solidity: event WithdrawCollateral(address indexed src, address indexed to, address indexed asset, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) WatchWithdrawCollateral(opts *bind.WatchOpts, sink chan<- *CometMainInterfaceWithdrawCollateral, src []common.Address, to []common.Address, asset []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _CometMainInterface.contract.WatchLogs(opts, "WithdrawCollateral", srcRule, toRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometMainInterfaceWithdrawCollateral)
				if err := _CometMainInterface.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
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

// ParseWithdrawCollateral is a log parse operation binding the contract event 0xd6d480d5b3068db003533b170d67561494d72e3bf9fa40a266471351ebba9e16.
//
// Solidity: event WithdrawCollateral(address indexed src, address indexed to, address indexed asset, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) ParseWithdrawCollateral(log types.Log) (*CometMainInterfaceWithdrawCollateral, error) {
	event := new(CometMainInterfaceWithdrawCollateral)
	if err := _CometMainInterface.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CometMainInterfaceWithdrawReservesIterator is returned from FilterWithdrawReserves and is used to iterate over the raw logs and unpacked data for WithdrawReserves events raised by the CometMainInterface contract.
type CometMainInterfaceWithdrawReservesIterator struct {
	Event *CometMainInterfaceWithdrawReserves // Event containing the contract specifics and raw log

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
func (it *CometMainInterfaceWithdrawReservesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CometMainInterfaceWithdrawReserves)
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
		it.Event = new(CometMainInterfaceWithdrawReserves)
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
func (it *CometMainInterfaceWithdrawReservesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CometMainInterfaceWithdrawReservesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CometMainInterfaceWithdrawReserves represents a WithdrawReserves event raised by the CometMainInterface contract.
type CometMainInterfaceWithdrawReserves struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawReserves is a free log retrieval operation binding the contract event 0xec4431f2ba1a9382f6b0c4352b888cba6f7db91667d9f776abe5ad8ddc5401b6.
//
// Solidity: event WithdrawReserves(address indexed to, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) FilterWithdrawReserves(opts *bind.FilterOpts, to []common.Address) (*CometMainInterfaceWithdrawReservesIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CometMainInterface.contract.FilterLogs(opts, "WithdrawReserves", toRule)
	if err != nil {
		return nil, err
	}
	return &CometMainInterfaceWithdrawReservesIterator{contract: _CometMainInterface.contract, event: "WithdrawReserves", logs: logs, sub: sub}, nil
}

// WatchWithdrawReserves is a free log subscription operation binding the contract event 0xec4431f2ba1a9382f6b0c4352b888cba6f7db91667d9f776abe5ad8ddc5401b6.
//
// Solidity: event WithdrawReserves(address indexed to, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) WatchWithdrawReserves(opts *bind.WatchOpts, sink chan<- *CometMainInterfaceWithdrawReserves, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CometMainInterface.contract.WatchLogs(opts, "WithdrawReserves", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CometMainInterfaceWithdrawReserves)
				if err := _CometMainInterface.contract.UnpackLog(event, "WithdrawReserves", log); err != nil {
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

// ParseWithdrawReserves is a log parse operation binding the contract event 0xec4431f2ba1a9382f6b0c4352b888cba6f7db91667d9f776abe5ad8ddc5401b6.
//
// Solidity: event WithdrawReserves(address indexed to, uint256 amount)
func (_CometMainInterface *CometMainInterfaceFilterer) ParseWithdrawReserves(log types.Log) (*CometMainInterfaceWithdrawReserves, error) {
	event := new(CometMainInterfaceWithdrawReserves)
	if err := _CometMainInterface.contract.UnpackLog(event, "WithdrawReserves", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
