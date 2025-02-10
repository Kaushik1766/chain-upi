package db

import "github.com/Kaushik1766/chain-upi-gin/internal/models"

func AddWallet(wallet *models.Wallet) error {
	result := DB.Create(wallet)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
