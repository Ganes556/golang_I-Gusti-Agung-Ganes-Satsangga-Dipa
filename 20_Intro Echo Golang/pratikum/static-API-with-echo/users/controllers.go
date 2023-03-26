package users

type errMsg struct {
	message string
}

func (err *errMsg) Error() string {
	return err.message
}

func GetUserById(id int) (int, User, error) {
	var index int = -1
	if id > 0 {
		index = BinarySearchUser(id, Users, func(i int) bool { return Users[i].Id == id })
	}
	if index != -1 {
		return index, Users[index], nil
	}
	return index, User{}, &errMsg{"user not found"}
}

func DeleteUserById(id int) error {
	index, _, err := GetUserById(id)
	if err != nil {
		return err
	}
	Users = append(Users[:index], Users[index+1:]...)
	return nil
}

func UpdateUserById(id int, user User) (User, error) {
	index, _, err := GetUserById(id)
	if err != nil {
		return User{}, err
	}
	user.Id = id
	Users[index] = user
	return Users[index], nil
}

func AddUser(user User) (User, error) {
	users := MergeSortUser(Users, func(left, right User) bool {
		return left.Id < right.Id
	})
	isEmailExist := BinarySearchUser(len(users), users, func(i int) bool {
		return Users[i].Email == user.Email
	})
	if isEmailExist != -1 {
		return User{}, &errMsg{"email already exist"}
	}

	if len(Users) == 0 {
		user.Id = 1
	} else {
		newId := Users[len(Users)-1].Id + 1
		user.Id = newId
	}

	Users = append(Users, user)
	return user, nil

}