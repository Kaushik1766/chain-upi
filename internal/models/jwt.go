package models

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	Email string
	Name  string
	UID   string
	jwt.StandardClaims
}
