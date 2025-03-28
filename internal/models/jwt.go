package models

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	Email     string
	Name      string
	UID       string
	UpiHandle string
	jwt.RegisteredClaims
}
