package models

type User struct {
	Base
	Name     string `json:"name,omitempty" form:"name" validate:"required"`
	Email    string `json:"email,omitempty" form:"email" gorm:"index" validate:"required,email"`
	Password string `json:"password,omitempty" form:"password" validate:"required,min=8"`
}