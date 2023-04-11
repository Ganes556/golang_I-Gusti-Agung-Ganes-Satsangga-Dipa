package model

type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRes struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type UserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}