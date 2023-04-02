package models

type Book struct {
	Base
	Judul    string `json:"judul" form:"judul" validate:"required"`
	Penulis  string `json:"penulis" form:"penulis" validate:"required"`
	Penerbit string `json:"penerbit" form:"penerbit" validate:"required"`
}