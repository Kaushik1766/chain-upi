package trx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Account represents the response structure for the Tron getaccount API
type Account struct {
	Address string `json:"address"` // The account address (base58check format)
	Balance int64  `json:"balance"` // TRX balance in sun (1 TRX = 1,000,000 sun)
}

// Vote represents a vote entry for a witness
type Vote struct {
	VoteAddress string `json:"vote_address"` // Address of the witness voted for
	VoteCount   int64  `json:"vote_count"`   // Number of votes
}

// Frozen represents frozen balance details
type Frozen struct {
	FrozenBalance int64     `json:"frozen_balance"` // Amount frozen (in sun)
	ExpireTime    time.Time `json:"expire_time"`    // Expiration timestamp of the freeze
}

// AccountResource represents additional resource details
type AccountResource struct {
	EnergyUsage  int64 `json:"energy_usage,omitempty"`   // Energy usage in this resource
	EnergyLimit  int64 `json:"energy_limit,omitempty"`   // Energy limit in this resource
	FreeNetUsage int64 `json:"free_net_usage,omitempty"` // Free bandwidth usage in this resource
	FreeNetLimit int64 `json:"free_net_limit,omitempty"` // Free bandwidth limit in this resource
}

func GetBalance(wallet string) (float64, error) {
	baseUrl := os.Getenv("TRX_BASE_URL")
	url := baseUrl + "/wallet/getaccount"

	payload := strings.NewReader("{\"address\":\"" + wallet + "\",\"visible\":true}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	// fmt.Println(string(body))
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var data Account
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return float64(data.Balance) / 1e6, nil
}
