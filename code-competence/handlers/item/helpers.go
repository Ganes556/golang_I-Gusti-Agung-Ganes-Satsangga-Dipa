package item

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

	if strings.Contains(err.Error(), "already exists") {
		return http.StatusConflict
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
				if err.Field() == "CategoryID" {
					return errors.New("category_id is required")
				}
				return errors.New(field + " is required")
			}
			
			if err.Tag() == "number" {
				return errors.New(field + " must be a number")
			}
			
		}
	}
	return nil
}

func isReuqestUuid(m string) error {
	validate := validator.New()
	err := validate.Var(m, "uuid4")
	
	if err != nil {
		return errors.New("id must be a uuid")
	}
	return nil
}