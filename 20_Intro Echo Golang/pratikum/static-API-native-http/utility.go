package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": "method not found",
	})
}

func SearchUser(id int) (int, User) {
	if id > 0 {
		start := 0
		end := len(users) - 1
		for start <= end {
			mid := (start + end) / 2
			if users[mid].Id == id {
				return mid, users[mid]
			} else if users[mid].Id < id {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1, User{}
}