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
		return err
	}

	return c.JSON(200, echo.Map{
		"data": users,
	})
	
}

func (u *userController) CreateUser(c echo.Context) error{

	var user model.User

	c.Bind(&user)

	err := u.usecase.Create(&user)

	if err != nil {
		return err
	}
	return c.JSON(200, echo.Map{
		"data": user,
	})
	
}

// func (u *userController) LoginUser(c echo.Context) error{
	
// }