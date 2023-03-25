package main

func SearchUserById(id int) (int, User) {
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

func CheckUserByEmail(email string) bool {
	usersSorted := SortUsersByEmail(users)
	start := 0
	end := len(usersSorted) - 1
	for start <= end {
		mid := (start + end) / 2
		if usersSorted[mid].Email == email {
			return true
		} else if usersSorted[mid].Email < email {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

// sort users by email ascending with merge sort
func SortUsersByEmail(users []User) []User {
	if len(users) < 2 {
		return users
	}
	mid := len(users) / 2
	left := SortUsersByEmail(users[:mid])
	right := SortUsersByEmail(users[mid:])

	usersRes := make([]User, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			usersRes = append(usersRes, right...)
			break
		}
		if len(right) == 0 {
			usersRes = append(usersRes, left...)
			break
		}
		if left[0].Email < right[0].Email {
			usersRes = append(usersRes, left[0])
			left = left[1:]
		} else {
			usersRes = append(usersRes, right[0])
			right = right[1:]
		}
	}

	return usersRes
}