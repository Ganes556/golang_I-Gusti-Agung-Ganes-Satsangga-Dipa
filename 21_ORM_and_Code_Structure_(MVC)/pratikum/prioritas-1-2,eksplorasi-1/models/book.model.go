package models

type Book struct {
	Base
	Judul    string `json:"judul" form:"judul" gorm:"type:varchar(150)" validate:"required"`
	Penulis  string `json:"penulis" form:"penulis" gorm:"type:varchar(100)" validate:"required"`
	Penerbit string `json:"penerbit" form:"penerbit" gorm:"type:varchar(100)" validate:"required"`
}