package app

import (
	userHandler "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/handlers/user"
	userRepo "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/repositories/user"
	userUsecase "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/usecases/user"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func startRoute(db *gorm.DB) *echo.Echo {
	e := echo.New()
	
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
		
		// itemGroupRestrict := e.Group("/users", middlewares.JWTMiddleware())
		// {
			
		// }
	}

	return e
}