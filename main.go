package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ApeX-Protocol/multicall-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"

	constants "main/constants"
)

func main() {
	AaveHealthCheck(context.Background(), "0xb3D8CcA14f249Cf68E5272C1E7623BFecc705591")
}

func AaveHealthCheck(ctx context.Context, address string) (*big.Int, error) {

	const MulticallABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
	const LendingPoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alchemyEndpoint := os.Getenv("ALCHEMY_ENDPOINT")

	mc := multicall.NewMultiCall(alchemyEndpoint,
		"0xcA11bde05977b3631167028862bE2a173976CA11")

	csvFile, err := os.Open("borrowersGenerator/borrowers.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Loaded Borrowers CSV file")
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
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress(line[1])}})
	}

	result, err := mc.CallTargets(ctx, passData)
	if err != nil {
		panic(err)
	}

	counter := 1

	for i, account := range result {
		colorReset := "\033[0m"
		fmt.Println(string(colorReset), counter, " - ", i, ": ", account[5])
		toInteger := account[5].(*big.Int) //strconv.Atoi(account[5])
		if err != nil {
			return nil, err
		}
		comparison := toInteger.Cmp(ToWei(1.5, 18))
		if comparison == -1 {
			colorRed := "\033[31m"
			fmt.Println(string(colorRed), "	Attention to ", i)

		}

		counter++
	}

	return nil, nil
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}
