package utils_aave

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"

	ILendingPool "main/abi/ilendingpool"
	constants "main/constants"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/joho/godotenv"
)

func GenerateBorrowerList() {

	var wg sync.WaitGroup

	wg.Add(4)
	go scripts(11362579, 11396312, &wg)
	go scripts(13500000, 13504000, &wg)
	go scripts(15000000, 15000400, &wg)
	go scripts(16500000, 16500400, &wg)

	wg.Wait()
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

	//fmt.Println("startBlock: ", startBlock, "\nendBlock: ", endBlock)

	contractAddress := common.HexToAddress(constants.AAVE_LENDING_POOL_V2)
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

	contractAbi, err := abi.JSON(strings.NewReader(string(ILendingPool.ILendingPoolABI)))
	if err != nil {
		log.Fatal(err)
	}

	eventSignature := []byte("Borrow(address,address,address,uint256,uint256,uint256,uint16)")
	hash := crypto.Keccak256Hash(eventSignature)
	hashHex := hash.Hex()

	for _, vLog := range logs {

		event := struct {
			Reserve        common.Address
			User           common.Address
			OnBehalfOf     common.Address
			Amount         *big.Int
			BorrowRateMode *big.Int
			BorrowRate     *big.Int
			Referral       uint16
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "Borrow", vLog.Data)

		if err == nil && vLog.Topics[0].Hex() == hashHex {

			fname := "./aave_borrowers.csv"
			// read the file
			f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			w := csv.NewWriter(f)
			a := []string{vLog.TxHash.String(), event.User.String(), strconv.FormatUint(uint64(vLog.BlockNumber), 10)}
			w.Write(a)
			w.Flush()
			/*
				fmt.Println(event)
				fmt.Println("Reserve: ", event.Reserve)
				fmt.Println("User: ", event.User)
				fmt.Println("OnBehalfOf: ", event.OnBehalfOf)
				fmt.Println("Amount: ", event.Amount)
				fmt.Println("BorrowRateMode: ", event.BorrowRateMode)
				fmt.Println("BorrowRate: ", event.BorrowRate)
				fmt.Println("Referral: ", event.Referral)
			*/
		}

	}

	scripts(endBlock, endBlock+4000, wg)

}
