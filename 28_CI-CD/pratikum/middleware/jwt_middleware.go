package middleware

import (
	"belajar-go-echo/model"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)


type JWT interface {
	JwtMiddleware() echo.MiddlewareFunc
}

type jwtService struct{}

func NewJWTService() JWT {
	return &jwtService{}
}

func (jwtS *jwtService) JwtMiddleware() echo.MiddlewareFunc{
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtClaims)
		},
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/users/login" || (c.Request().Method == "POST" && c.Request().URL.Path == "/users")
		},
		SigningKey: []byte(os.Getenv("JWT_KEY")),
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			log.Print(auth)
			return auth, nil
		},
	})
}