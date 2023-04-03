package middlewares

import (
	"os"

	models_jwt "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models/jwt"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware() echo.MiddlewareFunc{
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &models_jwt.JwtClaims{}
		},
	})
}