package main

import (
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"

	"belajar-go-echo/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userService)

	e.Use(middleware.JwtMiddleware())
	e.GET("/users", userController.GetAllUsers)
	e.POST("/users", userController.CreateUser)
	e.POST("/users/login", userController.GetAllUsers)
}