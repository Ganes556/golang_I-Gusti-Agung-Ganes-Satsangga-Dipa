package models

type Blog struct {
	Base
	UserRefer uint   `json:"user_refer" form:"user_refer" validate:"required"`
	Judul     string `json:"judul" form:"judul" gorm:"type:varchar(150)" validate:"required"`
	Konten    string `json:"konten" form:"konten" validate:"required"`
}

type BlogRes struct {
	ID        uint   `json:"id"`
	UserRefer uint   `json:"user_refer"`
	Judul     string `json:"judul"`
	Konten    string `json:"konten"`
}

type BlogReqUpdate struct {
	UserRefer uint   `json:"user_refer"`
	Judul     string `json:"judul"`
	Konten    string `json:"konten"`
}

type BlogByUserId struct {
	Penulis string `json:"penulis"`
	Judul   string `json:"judul"`
	Konten  string `json:"konten"`
}