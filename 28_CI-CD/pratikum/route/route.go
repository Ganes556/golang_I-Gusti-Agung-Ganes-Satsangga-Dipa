package route

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/controller"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/pkg"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/repository"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/usecase"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/middleware"

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