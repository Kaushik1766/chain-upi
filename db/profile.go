package db

import (
	"fmt"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
)

func UpdatePassword(uid string, password string) error {
	result := DB.Model(&models.User{}).Where("uid = ?", uid).Update("password", password)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return result.Error
	}
	return nil
}
