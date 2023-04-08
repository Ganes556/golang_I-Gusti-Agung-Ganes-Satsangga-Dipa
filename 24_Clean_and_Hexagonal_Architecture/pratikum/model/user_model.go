package model

type User struct {
	ID        uint `json:"id" gorm:"primarykey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
