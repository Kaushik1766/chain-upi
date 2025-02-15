package trx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type validateResponse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

func ValidateAddress(address string) bool {
	baseUrl := os.Getenv("TRX_BASE_URL")
	url := baseUrl + "/wallet/validateaddress"

	payload := strings.NewReader("{\"address\":\"" + address + "\",\"visible\":true}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	var data validateResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return data.Result
}
