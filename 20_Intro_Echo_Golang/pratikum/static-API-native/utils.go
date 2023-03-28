package main


type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var Users []User

func SearchUser(id int) (int, User) {
	if id > 0 {
		start := 0
		end := len(Users) - 1
		for start <= end {
			mid := (start + end) / 2
			if Users[mid].Id == id {
				return mid, Users[mid]
			} else if Users[mid].Id < id {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1, User{}
}