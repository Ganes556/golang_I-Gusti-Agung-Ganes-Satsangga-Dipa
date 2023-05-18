package user

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/dto"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/helpers"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/middlewares"
	"github.com/labstack/echo/v4"
)

func (u *userHandler) Login(c echo.Context) error {
	req := dto.LoginDTO{}
	c.Bind(&req)

	if err := isRequestValid(&req); err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	user, err := u.userUC.Login(req)

	if err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	return helpers.ResponseSuccess(c, getStatus(err), "login success", echo.Map{
		"token": middlewares.JWTToken(user.ID, user.Name),
	})
	
}

func (u *userHandler) Register(c echo.Context) error {
	req := dto.RegisterDTO{}
	
	c.Bind(&req)

	if err := isRequestValid(&req); err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	err := u.userUC.Register(req)

	if err != nil {
		if err.Error() == "email already exist" {
			return helpers.ResponseError(c, getStatus(err), err)
		}
		return helpers.ResponseError(c, getStatus(err), err)
	}

	return helpers.ResponseSuccessWithoutData(c, getStatus(err), "register success")
}