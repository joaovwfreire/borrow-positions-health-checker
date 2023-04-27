package utils_compound

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ApeX-Protocol/multicall-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"

	"main/abi"
	constants "main/constants"
)

const MulticallABI = abi.MulticallABI
const LendingPoolABI = abi.CompoundLendingPoolABI

var alchemyEndpoint string = "not_initialized"

func QueryHealth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	alchemyEndpoint = os.Getenv("ALCHEMY_ENDPOINT")

	tokenData := GetTokens(context.Background())
	CompoundHealthCheck(context.Background(), tokenData)
}

func CompoundHealthCheck(ctx context.Context, tokens []TokenData) (*big.Int, error) {

	mc := multicall.NewMultiCall(alchemyEndpoint,
		"0xcA11bde05977b3631167028862bE2a173976CA11")
	csvFile, err := os.Open("./compound_borrowers.csv")
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
			line[0],
			LendingPoolABI,
			constants.COMPOUND_CUSDCV3,
			"borrowBalanceOf",
			[]interface{}{common.HexToAddress(line[0])}})

		// this is utterly stupid, but this POC needs to be done
		// the best way to do this is to deploy an-chain contract that does this
		for _, token := range tokens {
			passData = append(passData, &multicall.ViewCall{
				(line[0] + token.address),
				LendingPoolABI,
				constants.COMPOUND_CUSDCV3,
				"userCollateral",
				[]interface{}{common.HexToAddress(line[0]), common.HexToAddress(token.address)}})
		}

	}

	result, err := mc.CallTargets(ctx, passData)
	if err != nil {
		panic(err)
	}

	for _, line := range csvLines {

		healthFactor := big.NewInt(0)

		for _, token := range tokens {
			valueToAdd := big.NewInt(1)
			valueToAdd.Mul(valueToAdd, (result[line[0]+token.address])[0].(*big.Int))
			valueToAdd.Mul(valueToAdd, token.lastPrice)
			valueToAdd.Div(valueToAdd, token.scale)
			healthFactor.Add(healthFactor, valueToAdd)

		}

		healthFactor.Sub(healthFactor, ((result[line[0]])[0].(*big.Int)))
		fmt.Println("Health Factor for ", line[0], " is ", healthFactor)

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

func sliceToString(values []interface{}) string {
	s := make([]string, len(values)) // Pre-allocate the right size
	for index := range values {
		s[index] = fmt.Sprintf("%v", values[index])
	}
	return strings.Join(s, ",")
}

func GetTokens(ctx context.Context) []TokenData {

	mc := multicall.NewMultiCall(alchemyEndpoint,
		"0xcA11bde05977b3631167028862bE2a173976CA11")

	numAssets := []*multicall.ViewCall{}

	numAssets = append(numAssets, &multicall.ViewCall{
		"0",
		LendingPoolABI,
		constants.COMPOUND_CUSDCV3,
		"numAssets",
		[]interface{}{}})

	assetsN, err := mc.CallTargets(ctx, numAssets)
	if err != nil {
		panic(err)
	}

	passData := []*multicall.ViewCall{}

	for i := 0; i < int((assetsN["0"][0].(uint8))); i++ {

		iToString := strconv.Itoa(i)

		passData = append(passData, &multicall.ViewCall{
			iToString,
			LendingPoolABI,
			constants.COMPOUND_CUSDCV3,
			"getAssetInfo",
			[]interface{}{uint8(i)}})

	}

	result, err := mc.CallTargets(ctx, passData)
	if err != nil {
		panic(err)
	}

	tokens := []TokenData{}
	for i := range result {
		addressString := sliceToString(result[i])[3:45]
		priceFeedString := sliceToString(result[i])[46:88]
		scale := strings.Fields(sliceToString(result[i]))

		tokenScale := new(big.Int)
		tokenScale, _ = tokenScale.SetString(scale[3], 10)

		tokens = append(tokens, TokenData{addressString, priceFeedString, tokenScale, big.NewInt(0)})

	}

	const ChainLinkABI = abi.ChainLinkABI

	feedData := []*multicall.ViewCall{}
	for i := 0; i < int((assetsN["0"][0].(uint8))); i++ {

		iToString := strconv.Itoa(i)

		feedData = append(feedData, &multicall.ViewCall{
			Id:        iToString,
			TargetAbi: ChainLinkABI,
			Target:    tokens[i].priceFeed,
			Method:    "latestRoundData",
			Arguments: []interface{}{}})

	}

	feedResults, err := mc.CallTargets(ctx, feedData)
	if err != nil {
		panic(err)
	}

	for i := 0; i < int((assetsN["0"][0].(uint8))); i++ {

		iToString := strconv.Itoa(i)
		tokens[i].lastPrice = feedResults[iToString][1].(*big.Int)

	}

	return tokens
}

type TokenData struct {
	address   string
	priceFeed string
	scale     *big.Int
	lastPrice *big.Int
}
