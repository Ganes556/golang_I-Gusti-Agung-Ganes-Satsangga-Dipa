package utils

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/db"
)

func BinarySearchUser(i int, users []db.User, f func(int) bool) int {
	start := 0
	end := len(users) - 1
	for start <= end {
		mid := (start + end) / 2
		if f(mid) {
			return mid
		} else if users[mid].Id < i {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func MergeSortUser(users []db.User, comp func(left, right db.User) bool) []db.User {
	if len(users) < 2 {
		return users
	}
	mid := len(users) / 2
	left := MergeSortUser(users[:mid], comp)
	right := MergeSortUser(users[mid:], comp)

	resArr := make([]db.User, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			resArr = append(resArr, right...)
			break
		}
		if len(right) == 0 {
			resArr = append(resArr, left...)
			break
		}
		if comp(left[0], right[0]) {
			resArr = append(resArr, left[0])
			left = left[1:]
		} else {
			resArr = append(resArr, right[0])
			right = right[1:]
		}
	}

	return resArr
}