package eth

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrivateKeyToWallet(PrivateKey string) (*models.Wallet, error) {
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	wallet := models.Wallet{
		Address:    walletAddress.String(),
		PrivateKey: PrivateKey,
		Chain:      "eth",
	}
	// err = db.AddWallet(&wallet)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	return &wallet, nil
}
