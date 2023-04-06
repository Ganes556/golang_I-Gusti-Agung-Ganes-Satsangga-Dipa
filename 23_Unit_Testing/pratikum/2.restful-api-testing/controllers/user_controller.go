package controllers

import (
	"net/http"
	"strconv"

	conf "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/labstack/echo/v4"
)

// get all users
func GetUsers(c echo.Context) error {
	
  var users []models.User

  conf.DB.Find(&users)
	
  return c.JSON(http.StatusOK, echo.Map{
    "message": "success get all users",
    "users":   users,
  })
	
}

// get user by id
func GetUser(c echo.Context) error {
  // your solution here
  idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

  var user models.User
  
	if err := conf.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success get user by id " + idStr,
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
	
  if conf.DB.First(&models.User{}, "email = ?", user.Email).RowsAffected > 0{
    return echo.NewHTTPError(http.StatusBadRequest, "email already exist")
  }

  conf.DB.Create(&user)
	
  return c.JSON(http.StatusCreated, echo.Map{
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
			"message": "id must be number",
		})
	}
  
	if err := conf.DB.First(&models.User{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	conf.DB.Unscoped().Delete(&models.User{}, id)

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success delete user by id " + idStr,
	})

}

// update user by id
func UpdateUser(c echo.Context) error {
  // your solution here
  var user models.User
	c.Bind(&user)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}
	
	if err := c.Validate(echo.Map{"method":c.Request().Method,"data":&user}); err != nil {
		return err
  }
	err = conf.DB.First(&models.User{}, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	conf.DB.Model(&user).Where("id = ?", id).Updates(&user)

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success update user by id " + idStr,
	})
}