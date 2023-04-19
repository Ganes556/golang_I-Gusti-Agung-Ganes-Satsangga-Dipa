package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	controller "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/controllers"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/libs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/middlewares"
)

func New() *echo.Echo{
	e := echo.New()
	
	// Set up validator
	e.Validator = libs.NewValidator()
	
	// Set up middleware
	e.Use(middlewares.LoggerMiddleware())
	e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middlewares.JwtMiddleware())
	
	// users
	usersGroup := e.Group("/users")
	{
		usersGroup.GET("", controller.GetUsers)
		usersGroup.GET("/:id", controller.GetUser)
		usersGroup.POST("/login", controller.LoginUser)
		usersGroup.POST("", controller.CreateUser)
		usersGroup.PUT("/:id", controller.UpdateUser)
		usersGroup.DELETE("/:id", controller.DeleteUser)
	}

	// books
	booksGroup := e.Group("/books")
	{
		booksGroup.GET("", controller.GetBooks)
		booksGroup.GET("/:id", controller.GetBook)
		booksGroup.POST("", controller.CreateBook)
		booksGroup.PUT("/:id", controller.UpdateBook)
		booksGroup.DELETE("/:id", controller.DeleteBook)
	}

	// blogs
	blogsGroup := e.Group("/blogs")
	{
		blogsGroup.GET("", controller.GetBlogs)
		blogsGroup.GET("/:id", controller.GetBlog)
		blogsGroup.GET("/userId/:id", controller.GetBlogByUserId)
		blogsGroup.POST("", controller.CreateBlog)
		blogsGroup.PUT("/:id", controller.UpdateBlog)
		blogsGroup.DELETE("/:id", controller.DeleteBlog)
	}
	
	return e
}