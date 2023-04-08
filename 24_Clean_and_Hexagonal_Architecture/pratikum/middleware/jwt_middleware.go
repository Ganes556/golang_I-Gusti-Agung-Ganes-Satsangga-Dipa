package middleware

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomeClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(id uint, email string) (string, error) {
	claims := &jwtCustomeClaims{
		ID: id,
		Email: email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func JwtMiddleware() echo.MiddlewareFunc{
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &jwtCustomeClaims{}
		},
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/users/login"
		},
		SigningKey: []byte(os.Getenv("JWT_KEY")),
	})
}