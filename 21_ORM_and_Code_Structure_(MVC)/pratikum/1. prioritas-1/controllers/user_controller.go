package controllers

import (
	"net/http"
	"strconv"
	"strings"

	conf "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	
  var users []models.User

  if err := conf.DB.Find(&users).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
	
  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "success get all users",
    "users":   users,
  })
	
}

// get user by id
func GetUserController(c echo.Context) error {
  // your solution here
  idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

  var user models.User
  err = conf.DB.Select("id, name, email, password").First(&user,"id = ?", id).Error  
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"messages": "success get user by id " + idStr,
		"user": user,
	})

}

// create new user
func CreateUserController(c echo.Context) error {
  user := models.User{}
  c.Bind(&user)

  if err := c.Validate(&user); err != nil {
    return err
  }
	
  if conf.DB.First(&models.User{}, "email = ?", user.Email).RowsAffected > 0{
    return echo.NewHTTPError(http.StatusBadRequest, "email already exist")
  }

  if err := conf.DB.Save(&user).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
	
  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "success create new user",
    "user":    user,
  })
}

// delete user by id
func DeleteUserController(c echo.Context) error {
  // your solution here
  idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"messages": "id must be number",
		})
	}
  
	if err = conf.DB.First(&models.User{}, id).Unscoped().Delete(&models.User{}).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}	
  
  var count int64 = -1
  conf.DB.Model(&models.User{}).Count(&count)
  
  if count == 0 {
    // reset auto increment
    conf.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
  }

	return c.JSON(http.StatusOK, &echo.Map{
		"messages": "success delete user by id " + idStr,
	})

}

// update user by id
func UpdateUserController(c echo.Context) error {
  // your solution here
  var user models.User
	c.Bind(&user)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}
	
	if err := c.Validate(&user); err != nil {
		if !strings.Contains(err.Error(),"required") {
			return err
		}
  }

	user.ID = uint(id)

	if err = conf.DB.Updates(&user).First(&user,id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"messages": "success update user by id " + idStr,
		"user": user,
	})
}