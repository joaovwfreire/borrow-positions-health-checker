/*
This script generates a list of borrowers from the protocol queried. It's main purpose is to enable the execution of other code at this repo, such as the health factor script.
In order to make it quick, I've opted to run it in parallel through wait groups. It's important to note that the script will generate a file called `aave_borrowers.csv` at the root folder.
This file will be used by the other scripts.
Bear in mind this will generate thousands of address entries, so make sure to use specific block number filters in order to make the most out of this tool.
This tool does not take into account duplicated addresses, nor does it take into account the amount of debt the address has. It's just a list of addresses that have borrowed from the protocol.
One should be able to manipulate this data in order to get accurate results.
*/

package utils_compound

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"sync"

	constants "main/constants"

	"github.com/ApeX-Protocol/multicall-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// generates a dirty borrowers list
func GenerateWithdrawersList() {
	var wg sync.WaitGroup

	wg.Add(4)
	go scripts(11362579, 11396312, &wg)
	go scripts(13500000, 13504000, &wg)
	go scripts(15000000, 15000400, &wg)
	go scripts(16500000, 16500400, &wg)

	wg.Wait()
}

func GenerateBorrowerList() {

	CleanBorrowerList(context.Background(), "0xb3D8CcA14f249Cf68E5272C1E7623BFecc705591")
	GenerateWithdrawersList()

}

func CleanBorrowerList(ctx context.Context, address string) (*big.Int, error) {
	// multicall addresses at comet
	// userBasic principal must be negative
	// userBasic assetsIn must be > 0, if not - it's liquidatable

	// borrowBalanceOf(address) * oraclePrice / scale
	// this is the amount of debt in USD
	// if debt > 0, then it's liquidatable and should be added to the list
	// if debt < 0, then it's not liquidatable and should be removed from the list

	// how to get borrowHealthFactor?

	// borrowBalanceOf - the array iterator of (balances * oracle prices / scale)

	const MulticallABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
	const LendingPoolABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauseGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseTokenPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"extensionDelegate\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"supplyKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"supplyPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowKink\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeLow\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateSlopeHigh\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowPerYearInterestRateBase\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"storeFrontPriceFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"trackingIndexScale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingSupplySpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingBorrowSpeed\",\"type\":\"uint64\"},{\"internalType\":\"uint104\",\"name\":\"baseMinForRewards\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"baseBorrowMin\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"targetReserves\",\"type\":\"uint104\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"struct CometConfiguration.AssetConfig[]\",\"name\":\"assetConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"struct CometConfiguration.Configuration\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Absurd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadDecimals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadDiscount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadMinimum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowCFTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientReserves\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInt104\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInt256\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUInt104\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUInt128\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUInt64\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidateCFTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NegativeNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSelfTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCollateralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotForSale\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLiquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SupplyCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimestampTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooMuchSlippage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferInFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferOutFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"absorber\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAbsorbed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdValue\",\"type\":\"uint256\"}],\"name\":\"AbsorbCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"absorber\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"basePaidOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdValue\",\"type\":\"uint256\"}],\"name\":\"AbsorbDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"BuyCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supplyPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"withdrawPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"absorbPaused\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"buyPaused\",\"type\":\"bool\"}],\"name\":\"PauseAction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SupplyCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawReserves\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"absorber\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"absorb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accrueAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveThis\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseBorrowMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseMinForRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenPriceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTrackingBorrowSpeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTrackingSupplySpeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowKink\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowPerSecondInterestRateBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowPerSecondInterestRateSlopeHigh\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowPerSecondInterestRateSlopeLow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"buyCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extensionDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"i\",\"type\":\"uint8\"}],\"name\":\"getAssetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"offset\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"scale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"struct CometCore.AssetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetInfoByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"offset\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"scale\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"borrowCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidateCollateralFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidationFactor\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"supplyCap\",\"type\":\"uint128\"}],\"internalType\":\"struct CometCore.AssetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"utilization\",\"type\":\"uint256\"}],\"name\":\"getBorrowRate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getCollateralReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"utilization\",\"type\":\"uint256\"}],\"name\":\"getSupplyRate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUtilization\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAbsorbPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBorrowCollateralized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBuyPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isLiquidatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSupplyPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTransferPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isWithdrawPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidatorPoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"numAbsorbs\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"numAbsorbed\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"approxSpend\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"_reserved\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numAssets\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"supplyPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"withdrawPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"absorbPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"buyPaused\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseAmount\",\"type\":\"uint256\"}],\"name\":\"quoteCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storeFrontPriceFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"supply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"supplyFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyKink\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyPerSecondInterestRateBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyPerSecondInterestRateSlopeHigh\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyPerSecondInterestRateSlopeLow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"supplyTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalsCollateral\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"totalSupplyAsset\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_reserved\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trackingIndexScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferAssetFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBasic\",\"outputs\":[{\"internalType\":\"int104\",\"name\":\"principal\",\"type\":\"int104\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"baseTrackingAccrued\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"assetsIn\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_reserved\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userCollateral\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"balance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_reserved\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alchemyEndpoint := os.Getenv("ALCHEMY_ENDPOINT")

	mc := multicall.NewMultiCall(alchemyEndpoint,
		"0xcA11bde05977b3631167028862bE2a173976CA11")

	csvFile, err := os.Open("compound_withdrawers.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Loaded Compound Borrowers CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err)
	}

	passData := []*multicall.ViewCall{}

	for _, line := range csvLines {
		passData = append(passData, &multicall.ViewCall{
			line[1],
			LendingPoolABI,
			constants.COMPOUND_CUSDCV3,
			"borrowBalanceOf",
			[]interface{}{common.HexToAddress(line[1])}})
	}

	result, err := mc.CallTargets(ctx, passData)
	if err != nil {
		panic(err)
	}

	fname := "./compound_borrowers.csv"
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	counter := 0
	for _, accountBalance := range result {

		borrowBalance := accountBalance[0].(*big.Int)
		comparison := borrowBalance.Cmp(big.NewInt(1))

		if comparison == 1 {
			fmt.Println("Borrower: ", csvLines[counter][1], " - ", accountBalance)

			w := csv.NewWriter(f)
			a := []string{csvLines[counter][1], borrowBalance.String()}
			w.Write(a)
			w.Flush()
		}
		counter++
	}

	return nil, nil

}

func scripts(startBlock, endBlock int64, wg *sync.WaitGroup) {
	defer wg.Done()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alchemyEndpoint := os.Getenv("ALCHEMY_ENDPOINT")

	client, err := ethclient.Dial(alchemyEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(constants.COMPOUND_CUSDCV3)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(startBlock),
		ToBlock:   big.NewInt(endBlock),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		fmt.Println(err)
		scripts(startBlock, endBlock-100, wg)
	}

	eventSignature := []byte("Withdraw(address,address,uint256)")
	hash := crypto.Keccak256Hash(eventSignature)
	hashHex := hash.Hex()

	for _, vLog := range logs {

		if vLog.Topics[0].Hex() == hashHex {

			accountAddress := common.HexToAddress((vLog.Topics[1]).Hex())
			fname := "./compound_withdrawers.csv"
			f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			w := csv.NewWriter(f)
			a := []string{vLog.TxHash.String(), accountAddress.String(), strconv.FormatUint(uint64(vLog.BlockNumber), 10)}
			w.Write(a)
			w.Flush()
		}

	}

	scripts(endBlock, endBlock+4000, wg)

}
