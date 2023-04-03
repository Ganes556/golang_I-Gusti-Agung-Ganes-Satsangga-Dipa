package models

type Blog struct {
	Base
	UserRefer uint   `json:"user_refer" form:"user_refer" validate:"required"`
	Judul     string `json:"judul" form:"judul" gorm:"type:varchar(150)" validate:"required"`
	Konten    string `json:"konten" form:"konten" validate:"required"`
}