package models

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/utils"
	"gorm.io/gorm"
)

type User struct {
	Base
	Name     string `json:"name" form:"name" gorm:"type:varchar(150)" validate:"required"`
	Email    string `json:"email" form:"email" gorm:"type:varchar(255);index" validate:"required,email"`
	Password string `json:"password" form:"password" gorm:"type:varchar(64)" validate:"required,min=8"`
	Blogs    []Blog `json:"-" form:"-" gorm:"foreignkey:UserRefer"`
}

type UserAuth struct {
	Email string `json:"email" form:"email"  validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	return utils.HashPassword(u.Password, &u.Password)
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	return utils.HashPassword(u.Password, &u.Password)
}