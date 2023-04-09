package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)


type jwtClaims struct {
	UserId uint
	Name   string
	jwt.RegisteredClaims
}

func CreateToken(userId uint, name string) (string, error) {

	claims := jwtClaims{
		UserId: userId,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func JwtMiddleware() echo.MiddlewareFunc{
	return echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path ==  "/users/login" 
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &jwtClaims{}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		},
	})
}
