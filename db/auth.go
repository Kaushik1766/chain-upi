package db

import (
	"github.com/Kaushik1766/chain-upi-gin/internal/models"
)

func GetUser(email string) (*models.User, error) {
	var user models.User
	result := DB.Where(models.User{Email: email}).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	result := DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
