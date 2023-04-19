package pkg

import (
	"errors"

	"github.com/go-playground/validator"
)

type CustomValidator interface {
	Validate(i interface{}) error
}

type customValidator struct {
	Validator *validator.Validate
}

func NewValidator() CustomValidator {
	return &customValidator{validator.New()}
}

func (cv *customValidator) Validate(i interface{}) error {
	data := i.(map[string]interface{})["data"]
	method := i.(map[string]interface{})["method"]

	if err := cv.Validator.Struct(data); err != nil {		
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" && method == "POST" {
				return errors.New(err.Field() + " is required")
			}
		}
	}
	
	return nil
}

