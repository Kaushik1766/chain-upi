package eth


import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)
type ResponseData struct {
	Valid      bool   `json:"valid"`
	ScamReport string `json:"scamReport"`
	Message    string `json:"message"`
}


func ValidateAddress(address string) bool {
	
	url := os.Getenv("WALLET_VERIFY_API")
	apiKey := os.Getenv("WALLET_VERIFY_API_KEY")
	data := map[string]string{
		"address": address,
		"network": "eth",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Set headers
	req.Header.Set("X-Api-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()
	var responseData ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		fmt.Println("Error decoding response:", err)
	}
	result :=true
	if responseData.Valid {
		result = true
	} else {
		result = false
	}
	return result
}
