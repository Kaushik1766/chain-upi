package db

import (
	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/google/uuid"
)

func AddWallet(wallet *models.Wallet) error {
	var storedWallets []models.Wallet
	res := DB.Where(&models.Wallet{Chain: wallet.Chain, UserUID: wallet.UserUID}).Find(&storedWallets)
	if res.Error != nil {
		return res.Error
	}
	if len(storedWallets) == 0 {
		wallet.IsPrimary = true
	}
	res = DB.Create(wallet)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func SetPrimary(walletAddress string, uid uuid.UUID, chain string) error {

	wallet := models.Wallet{
		// Address: walletAddress,
		UserUID: uid,
	}
	res := DB.Model(&wallet).Where(&models.Wallet{IsPrimary: true, Chain: chain}).Update("IsPrimary", false)
	if res.Error != nil {
		return res.Error
	}
	wallet.Address = walletAddress
	res = DB.Model(&wallet).Update("IsPrimary", true)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetWalletsByChain(upiHandle string, chain string) ([]models.Wallet, error) {
	var wallets []models.Wallet
	res := DB.Preload("User").Where(&models.Wallet{User: models.User{UpiHandle: upiHandle}, Chain: chain}).Find(&wallets)
	if res.Error != nil {
		return nil, res.Error
	}
	return wallets, nil
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
