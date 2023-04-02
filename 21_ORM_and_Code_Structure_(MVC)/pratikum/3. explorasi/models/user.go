package models

type User struct {
	Base
	Name     string `json:"name" form:"name" gorm:"type:varchar(100)" validate:"required"`
	Email    string `json:"email" form:"email" gorm:"type:varchar(200);uniqueIndex" validate:"required,email"`
	Password string `json:"password" form:"password" gorm:"type:varchar(200)" validate:"required,min=8"`
	Blogs    []Blog `json:"-" form:"-" gorm:"foreignkey:UserRefer"`
}