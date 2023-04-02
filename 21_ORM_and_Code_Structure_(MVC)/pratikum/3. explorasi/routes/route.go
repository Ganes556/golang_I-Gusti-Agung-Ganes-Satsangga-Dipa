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
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)

	// books
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	e.POST("/books", controller.CreateBookController)
	e.PUT("/books/:id", controller.UpdateBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)

	// blogs
	e.GET("/blogs", controller.GetBlogsController)
	e.GET("/blogs/:id", controller.GetBlogController)
	e.POST("/blogs", controller.CreateBlogController)
	e.PUT("/blogs/:id", controller.UpdateBlogController)
	e.DELETE("/blogs/:id", controller.DeleteBlogController)
	
	return e
}