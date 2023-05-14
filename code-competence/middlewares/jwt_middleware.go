package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type jwtDecode struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func JWTToken(id uint, name string) string {
	claims := &jwtClaims{
		ID: id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))

	return token
}

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtClaims)
		},
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			return jwt.ParseWithClaims(auth, new(jwtClaims), func(token *jwt.Token) (interface{}, error) {
				if token.Claims.(*jwtClaims).ExpiresAt.Time.Before(time.Now()) {
					return nil, fmt.Errorf("token expired")
				}
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
		},
		SuccessHandler: func(c echo.Context) {
			data := c.Get("user").(*jwt.Token).Claims.(*jwtClaims)
			c.Set("user", jwtDecode{
				ID:   data.ID,
				Name: data.Name,
			})
		},
		ErrorHandler: func(c echo.Context, err error) error {
			if err.Error() != "token expired" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		},
	})
}