package utils

import (
	"fmt"
	"os"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJwt(jwtToken string) (*models.JwtClaims, error) {
	_secret := []byte(os.Getenv("SECRET"))
	token, err := jwt.ParseWithClaims(jwtToken, &models.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("INVALID SIGNING METHOD")
		}
		return _secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*models.JwtClaims); ok && token.Valid {
		return claims, nil
	}
	// fmt.Println(token)
	return nil, fmt.Errorf("INVALID TOKEN")
}
