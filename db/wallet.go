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
	res := DB.Preload("User").
		Joins("JOIN users on users.uid=wallets.user_uid").
		Where("wallets.chain=? and users.upi_handle=? and wallets.is_primary=?", chain, upiHandle, true).
		First(&wallet)
	// fmt.Println(wallet.ToString())
	if res.Error != nil {
		return nil, res.Error
	}
	return &wallet, nil
}

func GetPrimaryWalletsByUpiHandle(upiHandle string) ([]models.Wallet, error) {
	var wallets []models.Wallet
	res := DB.Preload("User").
		Joins("JOIN users on users.uid=wallets.user_uid").
		Where("users.upi_handle=? and wallets.is_primary=?", upiHandle, true).Find(&wallets)
	if res.Error != nil {
		return nil, res.Error
	}
	return wallets, nil
}

func VerifyWallet(walletAddress string, uid string) error {
	var wallet models.Wallet
	res := DB.Where("wallets.address=? and wallets.user_uid=?", walletAddress, uid).First(&wallet)
	if res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}
