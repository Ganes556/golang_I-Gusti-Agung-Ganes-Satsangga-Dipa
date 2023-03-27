package db

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		var validationErrors validator.ValidationErrors
		if !errors.As(err, &validationErrors) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		for _, e := range validationErrors {
			if e.Field() == "Email" && e.Tag() == "email" {
				return echo.NewHTTPError(http.StatusBadRequest, "email is not valid")
			}
			if e.Field() == "Password" && e.Tag() == "min" {
				return echo.NewHTTPError(http.StatusBadRequest, "password must be at least 8 characters")
			}
			if e.Tag() == "required" {
				return echo.NewHTTPError(http.StatusBadRequest, "name, email and password are required")
			}
		}
  }
	return nil
}

func InitValidate(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
}