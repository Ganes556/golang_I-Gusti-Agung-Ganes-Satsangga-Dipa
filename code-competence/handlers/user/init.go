package user

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/usecases/user"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type userHandler struct {
	userUC user.UserUsecase
}

func NewUserHandler(userUC user.UserUsecase) UserHandler {
	return &userHandler{userUC: userUC}
}
