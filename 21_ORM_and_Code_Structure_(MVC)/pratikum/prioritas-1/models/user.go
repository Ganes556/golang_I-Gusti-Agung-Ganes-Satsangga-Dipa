package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id" form:"id"`
	CreatedAt time.Time      `json:"-" form:"-"`
	UpdatedAt time.Time      `json:"-" form:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" form:"-"`
}

type User struct {
	BaseModel
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" gorm:"index" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}