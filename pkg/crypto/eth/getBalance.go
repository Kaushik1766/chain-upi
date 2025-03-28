package eth

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func GetBalance(address string) (string, error) {
	baseUrl := os.Getenv("ETH_BASE_URL")
	if baseUrl == "" {
		return "", fmt.Errorf("etherscan base url not found in .env")
	}
	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("ETHERSCAN_API_KEY not found in .env")

	}
	url := baseUrl + fmt.Sprintf("v2/api?chainid=1&module=account&action=balance&address=%s&tag=latest&apikey=%s", address, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ethResp Response
	err = json.NewDecoder(resp.Body).Decode(&ethResp)
	if err != nil {
		return "", err
	}

	if ethResp.Status != "1" {
		return "", fmt.Errorf(ethResp.Message)
	} else {
		wei, _ := new(big.Int).SetString(ethResp.Result, 10)
		return fmt.Sprint(WeiToEth(wei)), nil
	}
}
