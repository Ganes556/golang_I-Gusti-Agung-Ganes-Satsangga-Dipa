package helpers

import (
	"errors"
	"strconv"
)

func CheckId(str, field string) (int, error) {
	
	num, err := strconv.Atoi(str)
	if err != nil || num < 1 {
		return 0, errors.New(field + " must be a number and greater than 0")
	}
	
	return num, nil
}