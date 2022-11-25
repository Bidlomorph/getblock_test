package internal

import (
	"encoding/json"
	"fmt"
	"getblock/responses"
	"getblock/tools"
	"log"
	"math/big"
)

func FindMostChanged(addresses map[string]*big.Int) (string, *big.Int) {
	var maxKey string
	var max *big.Int

	for key, value := range addresses {
		absoluteValue := big.NewInt(0).Abs(value)
		if max == nil {
			max = value
			maxKey = key
		}

		absoluteMax := big.NewInt(0).Abs(max)
		if absoluteMax.Cmp(absoluteValue) < 0 {
			max = value
			maxKey = key
		}
	}

	return maxKey, max
}

func ProcessTransactions(blockData *[]responses.BlockResponse) map[string]*big.Int {
	var addressesCollection = map[string]*big.Int{}

	for _, block := range *blockData {
		for _, tr := range block.Result.Transactions {
			if _, ok := addressesCollection[tr.From]; ok {
				addressesCollection[tr.From] = big.NewInt(0).Sub(addressesCollection[tr.From], tools.HexStringToInt(tr.Value))
			}
			addressesCollection[tr.From] = big.NewInt(tools.HexStringToInt(tr.Value).Int64())

			if _, ok := addressesCollection[tr.To]; ok {
				addressesCollection[tr.To] = big.NewInt(0).Add(addressesCollection[tr.To], tools.HexStringToInt(tr.Value))
			}
			addressesCollection[tr.To] = big.NewInt(tools.HexStringToInt(tr.Value).Int64())
		}
	}

	return addressesCollection
}

func TakeBlockTransactions(blockNumber string) responses.BlockResponse {
	var params []interface{}
	params = append(params, fmt.Sprintf("0x%s", blockNumber), true)

	requestBody := &tools.RequestBody{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  params,
		Id:      "getblock.io",
	}

	body := tools.MakeRequestToGetBlock(requestBody)

	data := responses.BlockResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("Unmarshal failed", err)
	}

	return data
}

func TakeLastBlockNumber() string {
	requestBody := &tools.RequestBody{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Id:      "getblock.io",
	}

	body := tools.MakeRequestToGetBlock(requestBody)

	data := responses.BlockNumberResponse{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("Unmarshal failed", err)
	}

	return fmt.Sprintf("%v", data.Result)
}
