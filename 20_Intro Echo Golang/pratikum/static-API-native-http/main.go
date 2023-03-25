package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)


func GetUsersController(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	idStr := strings.Split(r.URL.Path, "/users/")[1]
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"messages": "id must be number",
		})
		return
	}
	index, userRes := SearchUser(id)

	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{ 
			"messages": "user not found",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{ 
		"messages": "success get user by id",
		"user": userRes,
	})
}

func DeleteUserController(w http.ResponseWriter, r *http.Request) {
	
	idStr := strings.Split(r.URL.Path, "/users/")[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"messages": "id must be number",
		})
		return
	}
	index, _ := SearchUser(id)
	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"messages": "user not found",
		})
		return
	}

	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": "success delete user by id " + idStr,
	})

	users = append(users[:index], users[index+1:]...)
}

func UpdateUserController(w http.ResponseWriter, r *http.Request) {
	
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	idStr := strings.Split(r.URL.Path, "/users/")[1]
	// check if id is not number
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"messages": "id must be number",
		})
		return
	}
	index, _ := SearchUser(id)
	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"messages": "user not found",
		})
		return
	}
	user.Id = id
	users[index] = user
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": "success update user by id",
		"user":     users[index],
	})
}

func CreateUserController(w http.ResponseWriter, r *http.Request) {
	
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	if len(users) == 0 {
    user.Id = 1
  } else {
    newId := users[len(users)-1].Id + 1
    user.Id = newId
  }
	users = append(users, user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": "success create user",
		"user": user,
	})
}

func CreateUsersController(w http.ResponseWriter, r *http.Request) {
	
	usersReq := []User{}
	err := json.NewDecoder(r.Body).Decode(&usersReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"messages": "invalid request body",
		})
		return
	}
	for i, user := range usersReq {
		if len(users) == 0 {
			user.Id = 1
		} else {
			newId := users[len(users)-1].Id + 1
			user.Id = newId
		}
		users = append(users, user)
		usersReq[i] = user
		
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": "success create multiple user",
		"users":   usersReq,
	})
}

func MainController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	switch r.Method {
		case http.MethodGet:
			if strings.Contains(r.URL.Path,"/users/") {
				GetUserController(w, r)
				return
			}
			GetUsersController(w, r)
		case http.MethodPost:
			if r.URL.Path == "/users/multiple" {
				CreateUsersController(w, r)
				return
			}
			CreateUserController(w, r)
		case http.MethodPut:
			UpdateUserController(w, r)
		case http.MethodDelete:
			DeleteUserController(w, r)
		default:
			NotFound(w, r)
	}
}

func main() {
	// posible link
	http.HandleFunc("/users", MainController)
	http.HandleFunc("/users/", MainController)
	http.HandleFunc("/users/multiple", MainController)

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}