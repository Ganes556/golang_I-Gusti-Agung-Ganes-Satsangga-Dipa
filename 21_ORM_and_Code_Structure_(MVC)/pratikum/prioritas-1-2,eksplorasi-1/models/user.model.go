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

func (u *User) BeforeCreate(tx *gorm.DB) error {
	err := utils.HashPassword(u.Password, &u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	return u.BeforeCreate(tx)
}