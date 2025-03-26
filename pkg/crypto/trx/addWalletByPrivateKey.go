package trx

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58/base58"
)

func privateKeyToAdress(privateKey string) (string, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	var arr []byte

	x := key.X
	arr = append(arr, x.Bytes()...)

	y := key.Y
	arr = append(arr, y.Bytes()...)

	kec256 := crypto.Keccak256(arr)
	last40 := kec256[12:]
	append41, err := hex.DecodeString("41")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	append41 = append(append41, last40...)
	hash1 := sha256.Sum256(append41)
	hash2 := sha256.Sum256(hash1[:])
	checksum := hash2[:4]
	append41 = append(append41, checksum...)
	// fmt.Println(hex.EncodeToString(append41))
	address := base58.Encode(append41)
	// fmt.Println(address)
	return address, nil
}

func PrivateKeyToWallet(privateKey string) (*models.Wallet, error) {
	var wallet models.Wallet
	address, err := privateKeyToAdress(privateKey)
	if err != nil {
		return nil, err
	}
	wallet.Address = address
	wallet.Chain = "trx"
	wallet.PrivateKey = privateKey
	ok := ValidateAddress(wallet.Address)
	if !ok {
		return nil, fmt.Errorf("INVALID WALLET ADDRESS")
	}
	return &wallet, nil
}
