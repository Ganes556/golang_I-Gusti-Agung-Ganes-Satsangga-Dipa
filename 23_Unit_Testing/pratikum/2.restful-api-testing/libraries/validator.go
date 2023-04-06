package libraries

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	data := i.(echo.Map)["data"]
	method := i.(echo.Map)["method"]
	if err := cv.validator.Struct(data); err != nil {
		var validationErrors validator.ValidationErrors
		if !errors.As(err, &validationErrors) {
			return echo.NewHTTPError(http.StatusBadGateway, err.Error())
		}

		for _, e := range validationErrors {
			field := strings.ToLower(e.Field())
			if e.Tag() == "required" && method == "POST" {
				return echo.NewHTTPError(http.StatusBadRequest, field + " is required")
			}
			if e.Tag() == "email" {
				return echo.NewHTTPError(http.StatusBadRequest, field + " is not valid")
			}
			if e.Tag() == "min" {
				return echo.NewHTTPError(http.StatusBadRequest, field + " must be at least 8 characters")
			}

		}
	}
	return nil
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}