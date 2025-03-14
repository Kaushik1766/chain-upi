package trx

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type EtherscanResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []Transaction `json:"result"`
}

type Transaction struct {
	BlockNumber   string `json:"blockNumber"`
	TimeStamp     string `json:"timeStamp"`
	Hash          string `json:"hash"`
	From          string `json:"from"`
	To            string `json:"to"`
	Value         string `json:"value"`
	Gas           string `json:"gas"`
	GasPrice      string `json:"gasPrice"`
	Confirmations string `json:"confirmations"`
}

// type RequestBody struct {
// 	WalletAddress string `json:"wallet_address"`
// }

func GetTransactions(walletAddress string) EtherscanResponse {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		log.Fatal("ETHERSCAN_API_KEY not found in .env")
	}

	url := fmt.Sprintf(
		os.Getenv("TXN_HISTORY_URL"),
		walletAddress, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	var etherscanResp EtherscanResponse

	if err := json.NewDecoder(resp.Body).Decode(&etherscanResp); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	// for _, tx := range etherscanResp.Result {
	// 	fmt.Printf("Tx Hash: %s | From: %s | To: %s | Value: %s Wei\n", tx.Hash, tx.From, tx.To, tx.Value)
	// }
	return etherscanResp

}

/*
func transactionHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	GetTransactions(reqBody.WalletAddress)
}
*/
