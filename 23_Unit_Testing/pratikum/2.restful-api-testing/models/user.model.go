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

type UserRes struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserReqUpdate struct {
	Name string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required,email"`
}

type UserReqAuth struct {
	Email string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type UserResDB struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}