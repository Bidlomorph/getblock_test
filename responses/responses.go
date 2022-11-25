package responses

type BlockNumberResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  string `json:"result"`
}

type BlockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  struct {
		BaseFeePerGas   string `json:"baseFeePerGas"`
		Difficulty      string `json:"difficulty"`
		ExtraData       string `json:"extraData"`
		GasLimit        string `json:"gasLimit"`
		GasUsed         string `json:"gasUsed"`
		Hash            string `json:"hash"`
		LogsBloom       string `json:"logsBloom"`
		Miner           string `json:"miner"`
		MixHash         string `json:"mixHash"`
		Nonce           string `json:"nonce"`
		Number          string `json:"number"`
		ParentHash      string `json:"parentHash"`
		ReceiptsRoot    string `json:"receiptsRoot"`
		Sha3Uncles      string `json:"sha3Uncles"`
		Size            string `json:"size"`
		StateRoot       string `json:"stateRoot"`
		Timestamp       string `json:"timestamp"`
		TotalDifficulty string `json:"totalDifficulty"`
		Transactions    []struct {
			BlockHash            string `json:"blockHash"`
			BlockNumber          string `json:"blockNumber"`
			From                 string `json:"from"`
			Gas                  string `json:"gas"`
			GasPrice             string `json:"gasPrice"`
			MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
			MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
			Hash                 string `json:"hash"`
			Input                string `json:"input"`
			Nonce                string `json:"nonce"`
			To                   string `json:"to"`
			TransactionIndex     string `json:"transactionIndex"`
			Value                string `json:"value"`
			Type                 string `json:"type"`
			AccessList           []struct {
				Address     string   `json:"address"`
				StorageKeys []string `json:"storageKeys"`
			} `json:"accessList,omitempty"`
			ChainId string `json:"chainId"`
			V       string `json:"v"`
			R       string `json:"r"`
			S       string `json:"s"`
		} `json:"transactions"`
		TransactionsRoot string        `json:"transactionsRoot"`
		Uncles           []interface{} `json:"uncles"`
	} `json:"result"`
}
