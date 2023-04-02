package models


type User struct {
	Base
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" gorm:"index" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}