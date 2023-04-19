package controllers

import (
	"net/http"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/middlewares"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/services"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/utils"
	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	// your solution here
	var userReq models.UserReqAuth

	c.Bind(&userReq)

	if err := c.Validate(echo.Map{"method":c.Request().Method,"data":&userReq}); err != nil {
		return err
	}
	
	userDB, err := services.GetUserRepo().FindByEmail(userReq.Email)
	
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "user not found")
	}

	if err := utils.ComparePassword(userDB.Password,userReq.Password); err != nil{
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid email or password")
	}

	token, err := middlewares.GetJWTMiddleware().CreateToken(userDB.ID, userDB.Name)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success login",
		"token": token,
	})
}

// get all users
func GetUsers(c echo.Context) error {
	users, err := services.GetUserRepo().FindAll()
	// log.Print(users,err)
	if err != nil {
		return err
	}

  return c.JSON(http.StatusOK, &echo.Map{
		"message": "success get all users",
    "users": users,
  })
}

// get user by id
func GetUser(c echo.Context) error {
  // your solution here
  idstr := c.Param("id")
	user, err := services.GetUserRepo().FindById(idstr)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success get user by id " + idstr,
		"user": user,
	})

}

// create new user
func CreateUser(c echo.Context) error {
  user := models.User{}
	
  c.Bind(&user)

  if err := c.Validate(echo.Map{"method":c.Request().Method,"data":&user}); err != nil {
		return err
  }
	
	err := services.GetUserRepo().Create(user)

	if err != nil {
		return err
	}
	
  return c.JSON(http.StatusCreated, &echo.Map{
    "message": "success create new user",
  })
}

// update user by id
func UpdateUser(c echo.Context) error {
	// your solution here
  var user models.UserReqUpdate
	c.Bind(&user)

	idStr := c.Param("id")
	
	if err := c.Validate(echo.Map{"method":c.Request().Method,"data":&user}); err != nil {
		return err
  }
	
	if err := services.GetUserRepo().Update(idStr, user); err != nil {
		return err
	}	
		
	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success update user by id " + idStr,
	})
}

// delete user by id
func DeleteUser(c echo.Context) error {
	// your solution here
	idStr := c.Param("id")
	err := services.GetUserRepo().Delete(idStr)
	if err != nil {
		return err
	}	
	
	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success delete user by id " + idStr,
	})

}