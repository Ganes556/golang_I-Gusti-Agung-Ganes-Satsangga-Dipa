package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type MockJWTMiddleware struct {
	mock.Mock
}

func (m *MockJWTMiddleware) CreateToken(userId uint, name string) (string, error) {
	args := m.Called(userId, name)
	return args.String(0), args.Error(1)
}

func (m *MockJWTMiddleware) Middleware() echo.MiddlewareFunc {
	args := m.Called()
	return args.Get(0).(echo.MiddlewareFunc)
}