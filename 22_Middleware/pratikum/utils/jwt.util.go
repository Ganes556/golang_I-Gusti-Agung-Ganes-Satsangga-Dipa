package utils

import (
	"os"
	"time"

	models_jwt "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models/jwt"
	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userId uint, name string) (string, error) {

	claims := models_jwt.JwtClaims{
		UserId: userId,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}