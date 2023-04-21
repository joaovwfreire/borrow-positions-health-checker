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

	"github.com/ethereum/go-ethereum"
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
			fname := "./compound_borrowers.csv"
			// read the file
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
