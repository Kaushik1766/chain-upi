package utils

import (
	"os"

	"github.com/Kaushik1766/chain-upi-gin/internal/models"
	"github.com/golang-jwt/jwt"
)

func ValidateJwt(jwtToken string) (*models.JwtClaims, error) {
	_secret := os.Getenv("SECRET")
	claims, err := jwt.ParseWithClaims(jwtToken)
}
