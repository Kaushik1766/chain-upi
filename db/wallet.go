package db

import (
	"fmt"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
)

func AddWallet(wallet *models.Wallet) error {
	result := DB.Create(wallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetWalletByUpiHandle(upiHandle string, chain string) (*models.Wallet, error) {
	var wallet models.Wallet
	result := DB.Where(&models.Wallet{User: models.User{UpiHandle: upiHandle}, IsPrimary: }).First(&wallet)
	fmt.Println(wallet)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wallet, nil
}
