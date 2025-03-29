package trx

import (
	"bytes"
	// "crypto/ecdsa"
	// "crypto/rand"
	// "encoding/hex"
	"encoding/json"
	"fmt"

	// "log"
	"math/big"
	"net/http"
	"os"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/crypto"
)

// TronTransaction structure
type TronTransaction struct {
	RawData    TronRawData `json:"raw_data"`
	Signature  []string    `json:"signature"`
	TxID       string      `json:"txID"`
	Visible    bool        `json:"visible"`
	RawDataHex string      `json:"raw_data_hex"`
}

// Raw data inside the transaction
type TronRawData struct {
	Contract      []TronContract `json:"contract"`
	RefBlockBytes string         `json:"ref_block_bytes"`
	RefBlockHash  string         `json:"ref_block_hash"`
	Expiration    int64          `json:"expiration"`
	Timestamp     int64          `json:"timestamp"`
}

// Contract (for TRX transfer)
type TronContract struct {
	Parameter TronParameter `json:"parameter"`
	Type      string        `json:"type"`
}

// Transaction parameters
type TronParameter struct {
	Value   TronValue `json:"value"`
	TypeURL string    `json:"type_url"`
}

// Values (addresses and amount)
type TronValue struct {
	Amount       int64  `json:"amount"`
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address"`
}

type TronResponse struct {
	Result  bool   `json:"result"`
	Txid    string `json:"txid"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func padBytes(input []byte, length int) []byte {
	if len(input) >= length {
		return input
	}
	padded := make([]byte, length-len(input))
	return append(padded, input...)
}

func calculateRecoveryID(r, s *big.Int, privKey *secp256k1.PrivateKey, hash []byte) int {
	curve := secp256k1.S256()
	halfOrder := new(big.Int).Rsh(curve.N, 1) // Half of secp256k1 curve order

	// Enforce lower-S rule: Flip s if it's greater than half of the curve order
	if s.Cmp(halfOrder) > 0 {
		s.Sub(curve.N, s)
	}

	// Manually set `v` based on public key recovery
	v := 27 // Default to Ethereum/Tron standard

	// Check if the public key generated from r and s matches the private key
	pubKey := privKey.PubKey()
	if pubKey.SerializeCompressed()[0] == 0x03 { // If it's an odd key, set v = 28
		v = 28
	}

	return v
}

// createUnsignedTransaction calls TronGrid's createtransaction API
func createUnsignedTransaction(ownerAddr, toAddr string, amount int64) (*TronTransaction, error) {
	baseUrl := os.Getenv("TRX_BASE_URL")
	if baseUrl == "" {
		baseUrl = "https://api.shasta.trongrid.io/" // Default to mainnet
	}

	url := baseUrl + "wallet/createtransaction"
	payload := map[string]interface{}{
		"owner_address": ownerAddr,
		"to_address":    toAddr,
		"amount":        amount,
	}

	jsonPayload, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}
	defer resp.Body.Close()

	var tx *TronTransaction
	err = json.NewDecoder(resp.Body).Decode(&tx)
	if err != nil {
		return nil, fmt.Errorf("failed to decode transaction response: %v", err)
	}

	return tx, nil
}

// broadcastTx sends the signed transaction to TronGrid
func broadcastTx(tx *TronTransaction) error {
	baseUrl := os.Getenv("TRX_BASE_URL")
	if baseUrl == "" {
		baseUrl = "https://api.shasta.trongrid.io/"
	}

	url := baseUrl + "wallet/broadcasttransaction"
	jsonData, _ := json.Marshal(tx)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error broadcasting transaction: %v", err)
	}
	defer resp.Body.Close()

	var res TronResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return fmt.Errorf("failed to decode broadcast response: %v", err)
	}

	if !res.Result {
		return fmt.Errorf("broadcast failed: %v", res.Message)
	}

	fmt.Printf("Broadcast Result: %+v\n", res)
	return nil
}

// SendTrx orchestrates the full transaction process
func SendTrx(sender *models.Wallet, receiver string, amount float64) (string, error) {
	unsignedTx, err := CreateTransaction(sender.Address, receiver, amount)
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(sender.PrivateKey)
	if err != nil {
		return "", err
	}

	signedTx, err := SignTransaction(unsignedTx, privateKey)
	if err != nil {
		return "", err
	}

	_ = BroadcastTransaction(signedTx)
	return "", nil
}
