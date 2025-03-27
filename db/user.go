package db

import (
	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/google/uuid"
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

func GetUserByUid(uid string) (*models.User, error) {
	var user models.User
	parsedUid, err := uuid.Parse(uid)
	if err != nil {
		return nil, err
	}
	res := DB.First(&user, parsedUid)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
