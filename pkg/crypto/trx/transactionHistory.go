package trx

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
)

type TronscanResponse struct {
	Total      int64            `json:"total"`
	RangeTotal int64            `json:"rangeTotal"`
	Message    string           `json:"message"`
	Result     []TrxTransaction `json:"data"`
}

type TrxTransaction struct {
	Block        int64  `json:"block"`
	Hash         string `json:"hash"`
	OwnerAddress string `json:"ownerAddress"`
	To           string `json:"toAddress"`
	ContractType string `json:"contractType"`
	TimeStamp    int64  `json:"timestamp"`
	Confirmed    string `json:"confirmed"`
	ContractData struct {
		Value string `json:"amount"`
	} `json:"contractData"`
}

// type RequestBody struct {
// 	WalletAddress string `json:"wallet_address"`
// }

func GetTransactions(walletAddress string) ([]models.Transaction, error) {

	baseUrl := os.Getenv("TRXSCAN_BASE_URL")
	url := baseUrl + fmt.Sprintf("api/transaction?address=%s", walletAddress)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	var tronscanResp TronscanResponse
	if err := json.NewDecoder(resp.Body).Decode(&tronscanResp); err != nil {
		log.Printf("Error decoding JSON: %v", err)
	}

	if tronscanResp.Message != "" {
		log.Println(tronscanResp.Message)
		return nil, fmt.Errorf("error fetching transactions")
	}
	//fmt.Println(etherscanResp.Message)
	var transactions []models.Transaction
	for _, val := range tronscanResp.Result {
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		timestamp := time.Unix(val.TimeStamp/1000, 0).UTC()
		transactions = append(transactions, models.Transaction{
			Amount:          val.ContractData.Value,
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
