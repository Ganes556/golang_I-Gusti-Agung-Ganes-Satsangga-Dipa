package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type userHandlers struct {}

func (u *userHandlers) Gets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": "success get all users",
		"users":  Users,
	})
}

func (u *userHandlers) Get(w http.ResponseWriter, r *http.Request) {
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

func (u *userHandlers) Update(w http.ResponseWriter, r *http.Request) {
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

	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = id
	Users[index] = user


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{ 
		"messages": "success update user by id",
		"user": user,
	})
}

func (u *userHandlers) Delete(w http.ResponseWriter, r *http.Request) {
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

	Users = append(Users[:index], Users[index+1:]...)
	if len(Users) == 0 {
		Users = nil
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": "success delete user by id",
	})
}

func (u *userHandlers) Create(w http.ResponseWriter, r *http.Request) {
	
	user := User{}
	json.NewDecoder(r.Body).Decode(&user)
	if len(Users) == 0 {
    user.Id = 1
  } else {
    newId := Users[len(Users)-1].Id + 1
    user.Id = newId
  }
	Users = append(Users, user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": "success create user",
		"user": user,
	})
}

var UserHandlers = &userHandlers{}