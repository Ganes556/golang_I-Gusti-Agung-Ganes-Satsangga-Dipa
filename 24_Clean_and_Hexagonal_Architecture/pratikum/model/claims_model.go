package model

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
