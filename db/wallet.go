package db

import (
	"fmt"

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

func SetPrimary(walletAddress string, uid string, chain string) error {

	// var wallet models.Wallet
	fmt.Println("setprimarydb" + uid)
	parsedUid, _ := uuid.Parse(uid)
	res := DB.Model(&models.Wallet{}).
		Where("wallets.is_primary = ? and wallets.chain = ? and wallets.user_uid = ?", true, chain, parsedUid).
		Update("IsPrimary", false)
	// res := DB.Model(&wallet).Where(&models.Wallet{IsPrimary: true, Chain: chain}).Update("IsPrimary", false)
	if res.Error != nil {
		return res.Error
	}
	res = DB.Model(&models.Wallet{}).
		Where("wallets.address = ? and wallets.user_uid = ? and wallets.chain = ?", walletAddress, parsedUid, chain).
		Update("IsPrimary", true)
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

func GetWalletsByUid(uid string) ([]models.Wallet, error) {
	var wallets []models.Wallet
	parsedUid, _ := uuid.Parse(uid)
	res := DB.Where(models.Wallet{UserUID: parsedUid}).Find(&wallets)
	if res.Error != nil {
		return nil, res.Error
	}
	return wallets, nil

}

// func to get full wallet details by uid, address and chain
func GetUserWallet(uid string, address string, chain string) (*models.Wallet, error) {
	var wallet models.Wallet
	parsedUid, _ := uuid.Parse(uid)

	res := DB.Where(models.Wallet{UserUID: parsedUid, Address: address, Chain: chain}).First(&wallet)
	if res.Error != nil {
		fmt.Println(res.Error.Error())
		return nil, res.Error
	}
	return &wallet, nil
}

func GetPrimaryWalletByUid(uid string, chain string) (*models.Wallet, error) {
	var wallet models.Wallet
	parsedUid, _ := uuid.Parse(uid)

	res := DB.Where(models.Wallet{UserUID: parsedUid, IsPrimary: true, Chain: chain}).First(&wallet)
	if res.Error != nil {
		fmt.Println(res.Error.Error())
		return nil, res.Error
	}
	return &wallet, nil
}
