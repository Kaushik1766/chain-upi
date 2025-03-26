package eth

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/Kaushik1766/chain-upi-gin/db"
	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

func InsertEthWallet(PrivateKey string, uid string) error {
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return fmt.Errorf("error casting public key to ECDSA")
	}
	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	parsedUid, err := uuid.Parse(uid)
	if err != nil {
		log.Println(err)
		return err
	}
	wallet := models.Wallet{
		Address:    walletAddress.String(),
		PrivateKey: PrivateKey,
		UserUID:    parsedUid,
		Chain:      "eth",
	}
	err = db.AddWallet(&wallet)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
