package user

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func isRequestValid(m interface{}) error {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			if err.Tag() == "required" {
				return errors.New(field + " is required")
			}
			if err.Tag() == "email" {
				return errors.New(field + " is not valid")
			}
			if err.Tag() == "min" {
				return errors.New(field + " must be at least " + err.Param() + " characters")
			}
		}
	}
	return nil
}