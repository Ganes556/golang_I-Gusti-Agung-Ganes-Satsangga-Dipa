package route

import (
	"belajar-go-echo/controller"
	"belajar-go-echo/pkg"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"

	"belajar-go-echo/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	pkgPassword := pkg.NewPassword()

	userService := usecase.NewUserUsecase(userRepository, pkgPassword)	

	userController := controller.NewUserController(userService)

	jwtMid := middleware.NewJWTService()
	
	e.Use(jwtMid.JwtMiddleware())

	e.GET("/users", userController.GetAllUsers)
	e.POST("/users", userController.CreateUser)
	e.POST("/users/login", userController.LoginUser)
}