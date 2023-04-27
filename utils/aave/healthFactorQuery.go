package utils_aave

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

	"main/abi"
	constants "main/constants"
)

func QueryHealth() {
	AaveHealthCheck(context.Background(), "0xb3D8CcA14f249Cf68E5272C1E7623BFecc705591")
}

func AaveHealthCheck(ctx context.Context, address string) (*big.Int, error) {

	const LendingPoolABI = abi.AaveLendingPoolABI
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alchemyEndpoint := os.Getenv("ALCHEMY_ENDPOINT")

	mc := multicall.NewMultiCall(alchemyEndpoint,
		"0xcA11bde05977b3631167028862bE2a173976CA11")

	csvFile, err := os.Open("aave_borrowers.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Loaded AAVE Borrowers CSV file")
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
		//colorReset := "\033[0m"
		//fmt.Println(string(colorReset), counter, " - ", i, ": ", account[5])

		if err != nil {
			return nil, err
		}
		colorRed := "\033[31m"
		colorYellow := "\033[33m"
		comparison := account[5].(*big.Int).Cmp(ToWei(1.5, 18))

		if comparison == -1 {

			lowerComparison := account[5].(*big.Int).Cmp(ToWei(1.0, 18))

			if lowerComparison == -1 {
				fmt.Println(string(colorRed), "\nEvaluate: ", i, "\nCurrent Health Factor: ", account[5])
			} else {
				fmt.Println(string(colorYellow), "\nAttention: ", i, "\nCurrent Health Factor: ", account[5])
			}

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
