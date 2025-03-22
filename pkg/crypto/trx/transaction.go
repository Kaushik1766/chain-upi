package trx

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/ethereum/go-ethereum/crypto"
)

// TronTransaction represents the structure of the TRX transaction response
type TronTransaction struct {
	TxID    string `json:"txID"`
	RawData struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					OwnerAddress string `json:"owner_address"`
					ToAddress    string `json:"to_address"`
					Amount       int64  `json:"amount"`
				} `json:"value"`
			} `json:"parameter"`
		} `json:"contract"`
	} `json:"raw_data"`
	Signature []string `json:"signature,omitempty"`
}

// CreateTransaction generates an unsigned TRX transaction
func CreateTransaction(sender string, receiver string, amount float64) ([]byte, error) {
	url := os.Getenv("TRX_BASE_URL") + "/wallet/createtransaction"
	var sunAmount int64 = int64(amount * 1e6) // Convert TRX to SUN
	payload := map[string]interface{}{
		"amount":        sunAmount,
		"owner_address": sender,
		"to_address":    receiver,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to create transaction")
	}

	return io.ReadAll(res.Body)
}

// SignTransaction signs a given TRX transaction
func SignTransaction(transaction []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256(transaction)
	signature, err := crypto.Sign(hash[:], privateKey)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// BroadcastTransaction sends the signed TRX transaction to the blockchain
func BroadcastTransaction(signedTx []byte) error {
	url := os.Getenv("TRX_BASE_URL") + "/wallet/broadcasttransaction"

	res, err := http.Post(url, "application/json", bytes.NewBuffer(signedTx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("failed to broadcast transaction")
	}

	return nil
}

// SendTrx orchestrates the full transaction process
func SendTrx(sender *models.Wallet, receiver string, amount float64) error {
	unsignedTx, err := CreateTransaction(sender.Address, receiver, amount)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(sender.PrivateKey)
	if err != nil {
		return err
	}

	signedTx, err := SignTransaction(unsignedTx, privateKey)
	if err != nil {
		return err
	}

	return BroadcastTransaction(signedTx)
}
