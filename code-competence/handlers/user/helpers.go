package user

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func getStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if strings.Contains(err.Error(), "not found") {
		return http.StatusNotFound
	}

	if strings.Contains(err.Error(), "already exist") {
		return http.StatusConflict
	}

	if err.Error() == "invalid email or password" {
		return http.StatusUnauthorized
	}

	if err.Error() == "internal server error" {
		return http.StatusInternalServerError
	}

	return http.StatusBadRequest
}

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