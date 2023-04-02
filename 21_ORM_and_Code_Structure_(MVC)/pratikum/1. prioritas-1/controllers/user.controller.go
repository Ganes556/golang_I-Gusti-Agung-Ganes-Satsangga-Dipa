package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/services"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/utils"
	"github.com/labstack/echo/v4"
)

// get all users
func GetUsers(c echo.Context) error {
  var users []models.User
	err := services.FindAll(&users)
	if err != nil {
		return err
	}
  return c.JSON(http.StatusOK, &echo.Map{
		"message": "success get all users",
    "users":   users,
  })
	
}

// get user by id
func GetUser(c echo.Context) error {
  // your solution here
  idStr := c.Param("id")
	var id int
	err := utils.Id2Int(idStr, &id)

	if err != nil {
		return err
	}
  var user models.User
	err = services.FindById(id, &user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"messages": "success get user by id " + idStr,
		"user": user,
	})

}

// create new user
func CreateUser(c echo.Context) error {
  user := models.User{}
	
  c.Bind(&user)

  if err := c.Validate(&user); err != nil {
    return err
  }

	err := services.Create(&user)
	if err != nil {
		return err
	}
	
  return c.JSON(http.StatusOK, &echo.Map{
    "message": "success create new user",
    "user":    user,
  })
}

// delete user by id
func DeleteUser(c echo.Context) error {
  // your solution here
  idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"messages": "id must be number",
		})
	}
	
  err = services.DeleteById(id, &models.User{})
	if err != nil {
		return err
	}	
  
	return c.JSON(http.StatusOK, &echo.Map{
		"messages": "success delete user by id " + idStr,
	})

}

// update user by id
func UpdateUser(c echo.Context) error {
  // your solution here
  var user models.User
	c.Bind(&user)

	idStr := c.Param("id")
	var id int
	err := utils.Id2Int(idStr, &id)
	if err != nil {
		return err
	}
	
	if err := c.Validate(&user); err != nil {
		if !strings.Contains(err.Error(),"required") {
			return err
		}
  }
	
	user.ID = uint(id)
	if err := services.UpdateById(id, &user);err != nil {
		return err
	}	

	return c.JSON(http.StatusOK, &echo.Map{
		"messages": "success update user by id " + idStr,
		"user": user,
	})
}