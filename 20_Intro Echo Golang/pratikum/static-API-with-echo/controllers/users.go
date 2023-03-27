package controllers

import (
	"net/http"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/db"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/utils"
	"github.com/labstack/echo/v4"
)


func GetUserById(id int) (int, db.User, error) {
	var index int = -1
	if id > 0 {
		
		index = utils.BinarySearchUser(id, db.Users, func(i int) bool { return db.Users[i].Id == id })
	}
	if index != -1 {
		return index, db.Users[index], nil
	}

	return index, db.User{}, echo.NewHTTPError(http.StatusNotFound, "user not found")
}

func DeleteUserById(id int) error {
	index, _, err := GetUserById(id)
	if err != nil {
		return err
	}
	db.Users = append(db.Users[:index], db.Users[index+1:]...)
	return nil
}

func UpdateUserById(id int, user db.User) (db.User, error) {
	index, _, err := GetUserById(id)
	if err != nil {
		return db.User{}, err
	}
	user.Id = id
	db.Users[index] = user
	return db.Users[index], nil
}

func AddUser(user db.User) (db.User, error) {
	users := utils.MergeSortUser(db.Users, func(left, right db.User) bool {
		return left.Id < right.Id
	})
	isEmailExist := utils.BinarySearchUser(len(users), users, func(i int) bool {
		return db.Users[i].Email == user.Email
	})
	if isEmailExist != -1 {
		
		return db.User{}, echo.NewHTTPError(http.StatusConflict, "email already exist")
	}

	if len(db.Users) == 0 {
		user.Id = 1
	} else {
		newId := db.Users[len(db.Users)-1].Id + 1
		user.Id = newId
	}

	db.Users = append(db.Users, user)
	return user, nil

}