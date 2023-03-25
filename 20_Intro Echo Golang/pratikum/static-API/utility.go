package main

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