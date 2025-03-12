package db

import (
	"github.com/Kaushik1766/chain-upi-gin/internal/models"
)

func AddWallet(wallet *models.Wallet) error {
	result := DB.Create(wallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetPrimaryWalletByUpiHandle(upiHandle string, chain string) (*models.Wallet, error) {
	var wallet models.Wallet
	res := DB.Preload("User").Where(&models.Wallet{User: models.User{UpiHandle: upiHandle}, Chain: chain}).First(&wallet)
	// fmt.Println(wallet.ToString())
	if res.Error != nil {
		return nil, res.Error
	}
	return &wallet, nil

}
