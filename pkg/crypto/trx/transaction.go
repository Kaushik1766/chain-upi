package trx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type transaction struct {
	TransactionID string `json:"transactionId"`
	Success       bool   `json:"success"`
}

func SendTrx(senderAddress string, senderPrivateKey string, receiver string, amount float64) (string, error) {
	base := "https://chain-node.azurewebsites.net"
	// fmt.Println(base)
	url := fmt.Sprintf("%s/api/sendTrx", base)
	payloadData := fmt.Sprintf(`{"senderAddress":"%s","receiverAddress":"%s","privateKey":"%s","amount":%f}`, senderAddress, receiver, senderPrivateKey, amount)
	payload := strings.NewReader(payloadData)
	// fmt.Println(url)
	// fmt.Println(payloadData)
	// fmt.Println("payload length:", payload.Len())
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
	var data transaction
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	if data.Success {
		fmt.Println("Transaction ID:", data.TransactionID)
		return "https://shasta.tronscan.org/#/transaction/" + data.TransactionID, nil
	} else {
		return "", fmt.Errorf("transaction failed")
	}
}
