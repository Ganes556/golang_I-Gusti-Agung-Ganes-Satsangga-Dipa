package user

import (
	"net/http"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/dto"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/helpers"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/middlewares"
	"github.com/labstack/echo/v4"
)

func (uc *userHandler) Login(c echo.Context) error {
	req := dto.LoginDTO{}
	c.Bind(&req)

	if err := isRequestValid(&req); err != nil {
		return helpers.ResponseError(http.StatusBadRequest, err)
	}

	user, err := uc.userUC.Login(req)

	if err != nil {
		return helpers.ResponseError(http.StatusBadRequest, err)
	}

	return helpers.ResponseSuccess(c, http.StatusOK, "login success", echo.Map{
		"token": middlewares.JWTToken(user.ID, user.Name),
	})
	
}

func (uc *userHandler) Register(c echo.Context) error {
	req := dto.RegisterDTO{}
	
	c.Bind(&req)

	if err := isRequestValid(&req); err != nil {
		return helpers.ResponseError(http.StatusBadRequest, err)
	}

	err := uc.userUC.Register(req)

	if err != nil {
		return helpers.ResponseError(http.StatusBadRequest, err)
	}

	return helpers.ResponseSuccessWithoutData(c, http.StatusOK, "register success")
}