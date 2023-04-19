package models

type Book struct {
	Base
	Judul    string `json:"judul" form:"judul" gorm:"type:varchar(150)" validate:"required"`
	Penulis  string `json:"penulis" form:"penulis" gorm:"type:varchar(100)" validate:"required"`
	Penerbit string `json:"penerbit" form:"penerbit" gorm:"type:varchar(100)" validate:"required"`
}

type BookRes struct {
	ID       uint   `json:"id"`
	Judul    string `json:"judul"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
}

type BookReqUpdate struct {
	Judul    string `json:"judul"`
	Penulis  string `json:"penulis"`
	Penerbit string `json:"penerbit"`
}