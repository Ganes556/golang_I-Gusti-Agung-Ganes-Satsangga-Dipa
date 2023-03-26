package users

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users": Users,
	})
}

func GetUserHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "id must be number",
		})
	}
	_, user, err := GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id " + idStr,
		"user": user,
	})
}

func DeleteUserHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "id must be number",
		})
	}
	err = DeleteUserById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": err.Error(),
		})
	}	

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user by id " + idStr,
	})
}

func UpdateUserHandler(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "id must be number",
		})
	}

	user, err = UpdateUserById(id, user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user by id " + idStr,
		"user": user,
	})
}

func CreateUserHandler(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	if user.Email == "" || user.Name == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid request body",
		})
	}
	user, err := AddUser(user)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]interface{}{
			"messages": err.Error(),
		})
	}
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success create user",
    "user": user,
  })
}