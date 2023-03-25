package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
  Id    int    `json:"id" form:"id"`
  Name  string `json:"name" form:"name"`
  Email string `json:"email" form:"email"`
  Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success get all users",
    "users": users,
  })
}

// get user by id
func GetUserController(c echo.Context) error {
  // your solution here
	
	id, err := strconv.Atoi(c.Param("id"))
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "id must be number",
		})
	}
	
	index, userRes := SearchUserById(id)

	if index == -1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "user not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id",
		"user": userRes,
	})
	
}

// delete user by id
func DeleteUserController(c echo.Context) error {
  // your solution here

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "id must be number",
		})
	}
	index, _ := SearchUserById(id)
	if index == -1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "user not found",
		})
	}
	
	// remove userDel from users
	users = append(users[:index], users[index+1:]...)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user by id " + c.Param("id"),
	})
}
// update user by id
func UpdateUserController(c echo.Context) error {
  // your solution here
	user := User{}
	c.Bind(&user)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "id must be number",
		})
	}

	index, _ := SearchUserById(id)
	if index == -1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "user not found",
		})
	}

	user.Id = id
	users[index] = user

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user by id " + c.Param("id"),
		"user": user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
  // binding data
  user := User{}
  c.Bind(&user)
	if user.Email == "" || user.Name == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid request body",
		})
	}
	if CheckUserByEmail(user.Email) {
		return c.JSON(http.StatusConflict, map[string]interface{}{
			"messages": "user already exist",
		})
	}
  if len(users) == 0 {
    user.Id = 1
  } else {
    newId := users[len(users)-1].Id + 1
    user.Id = newId
  }
  users = append(users, user)
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success create user",
    "user": user,
  })
}

// ---------------------------------------------------
func main() {
  e := echo.New()
  // routing with query parameter
  e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
  e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}