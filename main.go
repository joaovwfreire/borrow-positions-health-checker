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
	A(context.Background(), "0xb3D8CcA14f249Cf68E5272C1E7623BFecc705591")
}

func A(ctx context.Context, address string) (*big.Int, error) {

	const MulticallABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
	const LendingPoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alchemyEndpoint := os.Getenv("ALCHEMY_ENDPOINT")

	mc := multicall.NewMultiCall(alchemyEndpoint,
		"0xcA11bde05977b3631167028862bE2a173976CA11")

	csvFile, err := os.Open("borrowers.csv")
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
		return nil, err
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

/*
	calls := []*multicall.ViewCall{
		{"key-1", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xb3D8CcA14f249Cf68E5272C1E7623BFecc705591")}},
		{"key-2", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x2797ff1C04A20ABEc5B8bC2a5b76A41D70d097C3")}},
		{"key-3", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x22F505099c3d54c0ee3dC44ad080b6E375B6E4E4")}},
		{"key-4", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x12Df0C95D2c549bbBC96cf8FbA02cA4Bc541aFD9")}},
		{"key-5", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xFaF1031B02A994b80f12Cc1ee4C0dCeBbB946aA0")}},
		{"key-6", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x3689c216f8f6ce7e2CE2a27c81a23096A787F532")}},
		{"key-7", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xEFFC18fC3b7eb8E676dac549E0c693ad50D1Ce31")}},
		{"key-8", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x820E72515f48d78C46807767bccB3037d75d9706")}},
		{"key-9", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x93dABae1444dAafd630d6B4a2d544EcFA4955579")}},
		{"key-10", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xD275E5cb559D6Dc236a5f8002A5f0b4c8e610701")}},
		{"key-11", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x6213d83C4aE95865477fC46d2aDa3F3516cf78cB")}},
		{"key-12", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xB821A5DA474C1Baa1f3e3e6E440171Cf08C7396a")}},
		{"key-13", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x7817c4100E8E51c147c7B7EaC1259abBdF47E7f0")}},
		{"key-14", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xA0109D15EeF50A5c6f3c58B4AB942b1c1e58812E")}},
		{"key-15", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xDdB78a939EB9ebcBeB77c19ec9e5dE94c5395D94")}},
		{"key-16", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x83Bc7e5785F7Cb2A114cd3070Ef820577A2D8888")}},
		{"key-17", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x777777c9898D384F785Ee44Acfe945efDFf5f3E0")}},
		{"key1", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xb3D8CcA14f249Cf68E5272C1E7623BFecc705591")}},
		{"key2", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x2797ff1C04A20ABEc5B8bC2a5b76A41D70d097C3")}},
		{"key3", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x22F505099c3d54c0ee3dC44ad080b6E375B6E4E4")}},
		{"key4", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x12Df0C95D2c549bbBC96cf8FbA02cA4Bc541aFD9")}},
		{"key5", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xFaF1031B02A994b80f12Cc1ee4C0dCeBbB946aA0")}},
		{"key6", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x3689c216f8f6ce7e2CE2a27c81a23096A787F532")}},
		{"key7", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xEFFC18fC3b7eb8E676dac549E0c693ad50D1Ce31")}},
		{"key8", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x820E72515f48d78C46807767bccB3037d75d9706")}},
		{"key9", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x93dABae1444dAafd630d6B4a2d544EcFA4955579")}},
		{"key10", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xD275E5cb559D6Dc236a5f8002A5f0b4c8e610701")}},
		{"key11", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x6213d83C4aE95865477fC46d2aDa3F3516cf78cB")}},
		{"key12", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xB821A5DA474C1Baa1f3e3e6E440171Cf08C7396a")}},
		{"key13", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x7817c4100E8E51c147c7B7EaC1259abBdF47E7f0")}},
		{"key14", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xA0109D15EeF50A5c6f3c58B4AB942b1c1e58812E")}},
		{"key15", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xDdB78a939EB9ebcBeB77c19ec9e5dE94c5395D94")}},
		{"key16", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x83Bc7e5785F7Cb2A114cd3070Ef820577A2D8888")}},
		{"key17", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x777777c9898D384F785Ee44Acfe945efDFf5f3E0")}},
		{"a-1", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xb3D8CcA14f249Cf68E5272C1E7623BFecc705591")}},
		{"a-2", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x2797ff1C04A20ABEc5B8bC2a5b76A41D70d097C3")}},
		{"a-3", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x22F505099c3d54c0ee3dC44ad080b6E375B6E4E4")}},
		{"a-4", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x12Df0C95D2c549bbBC96cf8FbA02cA4Bc541aFD9")}},
		{"a-5", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xFaF1031B02A994b80f12Cc1ee4C0dCeBbB946aA0")}},
		{"a-6", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x3689c216f8f6ce7e2CE2a27c81a23096A787F532")}},
		{"a-7", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xEFFC18fC3b7eb8E676dac549E0c693ad50D1Ce31")}},
		{"a-8", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x820E72515f48d78C46807767bccB3037d75d9706")}},
		{"a-9", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x93dABae1444dAafd630d6B4a2d544EcFA4955579")}},
		{"a-10", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xD275E5cb559D6Dc236a5f8002A5f0b4c8e610701")}},
		{"a-11", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x6213d83C4aE95865477fC46d2aDa3F3516cf78cB")}},
		{"a-12", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xB821A5DA474C1Baa1f3e3e6E440171Cf08C7396a")}},
		{"a-13", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x7817c4100E8E51c147c7B7EaC1259abBdF47E7f0")}},
		{"a-14", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xA0109D15EeF50A5c6f3c58B4AB942b1c1e58812E")}},
		{"a-15", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xDdB78a939EB9ebcBeB77c19ec9e5dE94c5395D94")}},
		{"a-16", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x83Bc7e5785F7Cb2A114cd3070Ef820577A2D8888")}},
		{"a-17", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x777777c9898D384F785Ee44Acfe945efDFf5f3E0")}},
		{"a1", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xb3D8CcA14f249Cf68E5272C1E7623BFecc705591")}},
		{"a2", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x2797ff1C04A20ABEc5B8bC2a5b76A41D70d097C3")}},
		{"a3", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x22F505099c3d54c0ee3dC44ad080b6E375B6E4E4")}},
		{"a4", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x12Df0C95D2c549bbBC96cf8FbA02cA4Bc541aFD9")}},
		{"a5", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xFaF1031B02A994b80f12Cc1ee4C0dCeBbB946aA0")}},
		{"a6", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x3689c216f8f6ce7e2CE2a27c81a23096A787F532")}},
		{"a7", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xEFFC18fC3b7eb8E676dac549E0c693ad50D1Ce31")}},
		{"a8", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x820E72515f48d78C46807767bccB3037d75d9706")}},
		{"a9", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x93dABae1444dAafd630d6B4a2d544EcFA4955579")}},
		{"a10", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xD275E5cb559D6Dc236a5f8002A5f0b4c8e610701")}},
		{"a11", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x6213d83C4aE95865477fC46d2aDa3F3516cf78cB")}},
		{"a12", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xB821A5DA474C1Baa1f3e3e6E440171Cf08C7396a")}},
		{"a13", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x7817c4100E8E51c147c7B7EaC1259abBdF47E7f0")}},
		{"a14", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xA0109D15EeF50A5c6f3c58B4AB942b1c1e58812E")}},
		{"a15", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0xDdB78a939EB9ebcBeB77c19ec9e5dE94c5395D94")}},
		{"a16", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x83Bc7e5785F7Cb2A114cd3070Ef820577A2D8888")}},
		{"a17", LendingPoolABI,
			constants.AAVE_LENDING_POOL_V2,
			"getUserAccountData",
			[]interface{}{common.HexToAddress("0x777777c9898D384F785Ee44Acfe945efDFf5f3E0")}},
	}

	result, err := mc.CallTargets(ctx, calls)
	if err != nil {
		return nil, err
	}

	fmt.Println(result["a17"])

	return nil, nil
}

/*
const MulticallABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"
const LendingPoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

func main() {
	/*
		var wg sync.WaitGroup

		wg.Add(4)
		go scripts(11362579, 11396312, &wg)
		go scripts(13500000, 13504000, &wg)
		go scripts(15000000, 15000400, &wg)
		go scripts(16500000, 16500400, &wg)

		wg.Wait()
*/
/*

	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/T7McdPHytjMPxs0RpmXo2eUpfI5lWxbv")
	if err != nil {
		log.Fatal(err)
	}

	msg := ethereum.CallMsg{}
	/*
	bytecodeBytes, err := hex.DecodeString(strings.ReplaceAll(bytecode, "0x", ""))
	if err != nil {
		panic(err)
	}

*/
/*
	address := common.HexToAddress(constants.AAVE_LENDING_POOL_V2)

	LendingPool, err := abi.JSON(strings.NewReader(string(LendingPoolABI)))
	if err != nil {
		panic(err)
	}

	abiMultiCall, err := abi.JSON(strings.NewReader(string(MulticallABI)))
	if err != nil {
		panic(err)
	}

	method := LendingPool.Methods["getUserAccountData"]
	id := method.ID
	inputs, err := method.Inputs.Pack("0xbc5a1a0da9e05fbf0b2804b063e3bcd6c2aaab7a")
	if err != nil {
		panic(err)
	}

	inputs = append(id[:], inputs[:]...)

	targets := []common.Address{address}
	datas := [][]byte{inputs}

	inputData, err := abiMultiCall.Constructor.Inputs.Pack(targets, datas)
	if err != nil {
		panic(err)
	}

	msg.Data = append(bytecodeBytes[:], inputData[:]...)
	result, err := client.PendingCallContract(context.TODO(), msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(result))

	ilendingpoolGetUserAccountData := ILendingPool.ILendingPoolCaller.GetUserAccountData()

		getUserAccountDataId := ilendingpoolGetUserAccountData.ID
		inputs, err := ilendingpoolGetUserAccountData.Inputs.Pack("0xbc5a1a0da9e05fbf0b2804b063e3bcd6c2aaab7a")

	inputs = append(getUserAccountDataId[:], inputs[:]...)
		data := Call{target: address, callData: inputs}

		multiCallAggregate = multicall2.StoreABI.Methods["aggregate"]
		multiCallInputs, err := multiCallAggregate.Inputs.Pack([]Call{data})

		fmt.Println(response.HealthFactor)
		fmt.Println(response)
*/

/*

	ilendingpoolGetUserAccountData := ilendingpool.GetUserAccountData()

	getUserAccountDataId := ilendingpoolGetUserAccountData.ID
	inputs, err := ilendingpoolGetUserAccountData.Inputs.Pack("0xbc5a1a0da9e05fbf0b2804b063e3bcd6c2aaab7a")

	inputs = append(getUserAccountDataId[:], inputs[:]...)
	data := Call{target: address, callData: inputs}

	multiCallAggregate = multicall2.StoreABI.Methods["aggregate"]
	multiCallInputs, err := multiCallAggregate.Inputs.Pack([]Call{data})

	fmt.Println(response.HealthFactor)
	fmt.Println(response)
*/
