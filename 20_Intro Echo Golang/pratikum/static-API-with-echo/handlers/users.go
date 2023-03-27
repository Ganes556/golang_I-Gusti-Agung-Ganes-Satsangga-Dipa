package handlers

import (
	"net/http"
	"strconv"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/controllers"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/db"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"messages": "success get all users",
		"users": db.Users,
	})
}

func GetUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}
	_, user, err :=  controllers.GetUserById(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &echo.Map{
		"messages": "success get user by id " + idStr,
		"user": user,
	})
}

func DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"messages": "id must be number",
		})
	}
	err = controllers.DeleteUserById(id)
	if err != nil {
		return err
	}	

	return c.JSON(http.StatusOK, &echo.Map{
		"messages": "success delete user by id " + idStr,
	})
}

func UpdateUser(c echo.Context) error {
	user := db.User{}
	c.Bind(&user)
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}
	
	if err := c.Validate(&user); err != nil {
    return err
  }

	user, err = controllers.UpdateUserById(id, user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"messages": "success update user by id " + idStr,
		"user": user,
	})
}

func CreateUser(c echo.Context) error {
	user := db.User{}
	c.Bind(&user)
	if err := c.Validate(&user); err != nil {
    return err
  }
	user, err := controllers.AddUser(user)
	if err != nil {
		return err
	}
  return c.JSON(http.StatusOK, &echo.Map{
    "messages": "success create user",
    "user": user,
  })
}