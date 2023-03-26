package users

import (
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.GET("/users", GetUsersHandler)
	e.GET("/users/:id", GetUserHandler)
  e.POST("/users", CreateUserHandler)
	e.PUT("/users/:id", UpdateUserHandler)
	e.DELETE("/users/:id", DeleteUserHandler)
}