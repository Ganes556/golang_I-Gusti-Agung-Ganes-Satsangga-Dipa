package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)



type IJWTMiddleware interface {
	CreateToken(userId uint, name string) (string, error)
	Middleware() echo.MiddlewareFunc
}

var jwtMid IJWTMiddleware

func init(){
	jwtMid = &JWTMiddleware{}
}

func SetJWTMiddleware(jwtM IJWTMiddleware) {
	jwtMid = jwtM
}

func GetJWTMiddleware() IJWTMiddleware {
	return jwtMid
}

type jwtClaims struct {
	UserId uint
	Name   string
	jwt.RegisteredClaims
}


type JWTMiddleware struct{
	
}


func (j *JWTMiddleware) CreateToken(userId uint, name string) (string, error) {
	
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

func (j *JWTMiddleware) Middleware() echo.MiddlewareFunc{
	return echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path ==  "/users/login" || (c.Request().Method == "POST" && c.Request().URL.Path == "/users")
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
