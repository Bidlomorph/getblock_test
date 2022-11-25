package tools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
)

type RequestBody struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      string        `json:"id"`
}

func MakeRequestToGetBlock(requestBody *RequestBody) []byte {
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(requestBody)

	req, err := http.NewRequest("POST", "https://eth.getblock.io/mainnet/", payloadBuf)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header = http.Header{
		"x-api-key":    {os.Getenv("API_KEY")},
		"Content-Type": {"application/json"},
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func HexStringToInt(inputValue string) *big.Int {
	n := new(big.Int)
	n.SetString(inputValue, 0)

	return n
}
