package routes

import (
	"github.com/labstack/echo/v4"

	controller "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/controllers"
	lib "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/libraries"
)

func New() *echo.Echo{
	e := echo.New()
	
	// Set up validator
	e.Validator = lib.NewValidator()

	// users
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)

	// books
	e.GET("/books", controller.GetBooks)
	e.GET("/books/:id", controller.GetBook)
	e.POST("/books", controller.CreateBook)
	e.PUT("/books/:id", controller.UpdateBook)
	e.DELETE("/books/:id", controller.DeleteBook)
	
	return e
}