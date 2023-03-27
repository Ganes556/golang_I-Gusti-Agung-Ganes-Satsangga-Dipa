package db

import (
	"github.com/go-playground/validator/v10"
)

type (
	User struct {
		Id       int    `json:"id" form:"id"`
		Name     string `json:"name" form:"name" validate:"required"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8"`
	}
	CustomValidator struct {
    validator *validator.Validate
  }
)


var Users []User