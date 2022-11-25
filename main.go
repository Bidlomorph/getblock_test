package main

import (
	"fmt"
	"getblock/internal"
	"getblock/responses"
	"getblock/tools"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Wait a moment...")
	address, volume := mostChangedAddress(100)

	fmt.Println("Most changed address:", address)
	fmt.Println("Changed volume:", volume)
}

func mostChangedAddress(countOfBlocksBefore int64) (string, *big.Int) {
	var blocks []responses.BlockResponse

	lastBlockNumberHex := internal.TakeLastBlockNumber()
	lastBlockNumber := tools.HexStringToInt(lastBlockNumberHex).Int64()

	for i := lastBlockNumber; i > lastBlockNumber-countOfBlocksBefore; i = i - 1 {
		blockNumber := strconv.FormatInt(i, 16)
		blockData := internal.TakeBlockTransactions(blockNumber)

		blocks = append(blocks, blockData)
	}

	addressesCollection := internal.ProcessTransactions(&blocks)
	address, volume := internal.FindMostChanged(addressesCollection)

	return address, volume
}
