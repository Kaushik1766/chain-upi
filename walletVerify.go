package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the API URL with query parameters
	var address="0x28D38f76B5631BD3CDf3C6E5F827CBCbE10cA2Bc"
	apiURL := "https://api.etherscan.io/v2/api?chainid=1&module=account&action=balance&address="+address+"&tag=latest&apikey=WCQ5TE8DRDU3VWF38WVPEN2MG948CD2KZV"
	// Make HTTP GET request
	resp, err := http.Get(apiURL)

	fmt.Println("Response")
	fmt.Println(resp)
	fmt.Println("Error")
	fmt.Println(err)
}
