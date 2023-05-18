package app

import (
	itemHandler "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/handlers/item"
	userHandler "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/handlers/user"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/middlewares"
	categoryRepo "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/repositories/category"
	itemRepo "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/repositories/item"
	userRepo "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/repositories/user"
	itemUsecase "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/usecases/item"
	userUsecase "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/usecases/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func startRoute(db *gorm.DB) *echo.Echo {
	e := echo.New()
	
	e.Pre(middleware.RemoveTrailingSlash())

	// users
	{
		userRepo := userRepo.NewUserRepository(db)
		userUsecase := userUsecase.NewUserUsecase(userRepo)
		userHandler := userHandler.NewUserHandler(userUsecase)
		
		userGroup := e.Group("/users")
		{
			userGroup.POST("/login", userHandler.Login)
			userGroup.POST("/register", userHandler.Register)
		}
	}

	// items
	{
		itemRepo := itemRepo.NewItemRepository(db)
		categoryRepo := categoryRepo.NewCategoryRepository(db)
		itemUsecase := itemUsecase.NewItemUsecase(itemRepo, categoryRepo)
		itemHandler := itemHandler.NewItemHandler(itemUsecase)

		itemGroupRestrict := e.Group("/items", middlewares.JWTMiddleware())
		{
			itemGroupRestrict.GET("", itemHandler.GetAll)
			itemGroupRestrict.POST("", itemHandler.Create)
			itemGroupRestrict.DELETE("/:id", itemHandler.Delete)
			itemGroupRestrict.GET("/:id", itemHandler.Get)
			itemGroupRestrict.PUT("/:id", itemHandler.Update)
			itemGroupRestrict.GET("/category/:category_id", itemHandler.GetByCategory)
		}
	}

	return e
}