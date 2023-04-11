package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
)

type UserController interface{
	GetAllUsers(c echo.Context) error
	CreateUser(c echo.Context) error
	LoginUser(c echo.Context) error
}

type userController struct{
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) *userController{
	return &userController{usecase}
}

func (u *userController) GetAllUsers(c echo.Context) error{
	
	users, err := u.usecase.GetAll()
	
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, map[string]interface{}{
		"data": users,
	})
	
}

func (u *userController) CreateUser(c echo.Context) error{

	var user model.UserReq

	c.Bind(&user)

	err := u.usecase.Create(user)

	if err != nil {
		return echo.NewHTTPError(409, err.Error())
	}

	return c.JSON(201, map[string]interface{}{
		"message": "success create user",
	})
	
}

func (u *userController) LoginUser(c echo.Context) error{
	var userAuth model.UserReq

	c.Bind(&userAuth)

	token, err := u.usecase.Login(userAuth)
	
	if err != nil {
		return echo.NewHTTPError(401, err.Error())
	}

	return c.JSON(200, map[string]interface{}{
		"token": token,
	})
	
}