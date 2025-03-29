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
	Address                  string               `json:"address"`                               // The account address (base58check format)
	Balance                  int64                `json:"balance"`                               // TRX balance in sun (1 TRX = 1,000,000 sun)
	Votes                    []Vote               `json:"votes,omitempty"`                       // List of votes for witnesses
	Frozen                   []Frozen             `json:"frozen,omitempty"`                      // Frozen balances for staking
	NetWindowSize            int64                `json:"net_window_size,omitempty"`             // Bandwidth window size
	NetWindowOptimized       bool                 `json:"net_window_optimized,omitempty"`        // Whether bandwidth window is optimized
	CreateTime               time.Time            `json:"create_time,omitempty"`                 // Account creation timestamp
	LatestOprationTime       time.Time            `json:"latest_opration_time,omitempty"`        // Last operation timestamp
	Allowance                int64                `json:"allowance,omitempty"`                   // Withdrawable allowance
	LatestWithdrawTime       time.Time            `json:"latest_withdraw_time,omitempty"`        // Last withdrawal timestamp
	Code                     string               `json:"code,omitempty"`                        // Smart contract code (if any)
	IsWitness                bool                 `json:"is_witness,omitempty"`                  // Whether the account is a witness
	IsCommittee              bool                 `json:"is_committee,omitempty"`                // Whether the account is a committee member
	FrozenSupply             []Frozen             `json:"frozen_supply,omitempty"`               // Frozen supply for token issuance
	Asset                    map[string]int64     `json:"asset,omitempty"`                       // Map of asset balances (token name -> amount)
	LatestAssetOperationTime map[string]time.Time `json:"latest_asset_operation_time,omitempty"` // Last operation time per asset
	FreeNetUsage             int64                `json:"free_net_usage,omitempty"`              // Free bandwidth usage
	FreeNetLimit             int64                `json:"free_net_limit,omitempty"`              // Free bandwidth limit
	NetUsage                 int64                `json:"net_usage,omitempty"`                   // Total bandwidth usage
	NetLimit                 int64                `json:"net_limit,omitempty"`                   // Total bandwidth limit
	EnergyUsage              int64                `json:"energy_usage,omitempty"`                // Energy usage
	EnergyLimit              int64                `json:"energy_limit,omitempty"`                // Energy limit
	AccountResource          AccountResource      `json:"account_resource,omitempty"`            // Additional account resources
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
	url := baseUrl + "/wallet/validateaddress"

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
