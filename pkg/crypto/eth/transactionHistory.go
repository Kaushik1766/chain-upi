package eth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
)

type EtherscanResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Result  []EthTransaction `json:"result"`
}

type EthTransaction struct {
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

func GetTransactions(walletAddress string) ([]models.Transaction, error) {

	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		log.Println("ETHERSCAN_API_KEY not found in .env")
	}

	baseUrl := os.Getenv("ETH_BASE_URL")
	url := baseUrl + fmt.Sprintf("api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&sort=desc&apikey=%s", walletAddress, apiKey)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	var etherscanResp EtherscanResponse
	if err := json.NewDecoder(resp.Body).Decode(&etherscanResp); err != nil {
		log.Printf("Error decoding JSON: %v", err)
	}

	if etherscanResp.Status == "0" {
		log.Println(etherscanResp.Message)
		return nil, fmt.Errorf("error fetching transactions")
	}
	//fmt.Println(etherscanResp.Message)
	var transactions []models.Transaction
	for _, val := range etherscanResp.Result {
		epochInt, err := strconv.ParseInt(val.TimeStamp, 10, 64)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		timestamp := time.Unix(epochInt, 0).UTC()
		transactions = append(transactions, models.Transaction{
			Amount:          val.Value,
			ReceiverAddress: val.To,
			TransactionHash: val.Hash,
			TimeStamp:       timestamp,
		})
	}
	return transactions, nil

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
