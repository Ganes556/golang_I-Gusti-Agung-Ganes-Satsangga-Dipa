package entities

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	ID uint `json:"id" gorm:"primary_key"`
	Name string    `json:"name" gorm:"type:varchar(150)"`
	Email string   `json:"email" gorm:"type:varchar(150);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(150)"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {	
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hash)
	return nil
}

