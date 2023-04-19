package services

import (
	"net/http"
	"strconv"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/labstack/echo/v4"
)

type IBookService interface {
	FindAll() ([]models.BookRes, error)
	FindById(idStr string) (models.BookRes, error)
	Create(book models.Book) error
	Update(idStr string, book models.BookReqUpdate) error
	Delete(idStr string) error
}

type BookRepo struct {}

var bookRepo IBookService

func init() {
	bookRepo = &BookRepo{}
}

func GetBookRepo() IBookService {
	return bookRepo
}

func SetBookRepo(br IBookService) {
	bookRepo = br
}


func (br *BookRepo) FindAll() ([]models.BookRes, error) {
	var books []models.BookRes
	err := configs.DB.Model(&models.Book{}).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (br *BookRepo) FindById(idStr string) (models.BookRes, error) {

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return models.BookRes{}, echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	var book models.BookRes
	err = configs.DB.Model(&models.Book{}).Where("id = ?", id).First(&book).Error
	if err != nil {
		return models.BookRes{}, echo.NewHTTPError(http.StatusNotFound, "book not found")
	}
	return book, nil
}


func (br *BookRepo) Create(book models.Book) error {

	found := configs.DB.Where("judul = ?",book.Judul).First(&models.Book{}).RowsAffected > 0
	
	if found {
		return echo.NewHTTPError(http.StatusConflict, "book already exist")
	}

	err := configs.DB.Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (br *BookRepo) Update(idStr string, book models.BookReqUpdate) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	err = configs.DB.Model(&models.Book{}).Where("id = ?", id).First(&models.BookReqUpdate{}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "book not found")
	}

	err = configs.DB.Model(&models.Book{}).Where("id = ?", id).Updates(&book).Error

	if err != nil {
		return err
	}
	return nil
}

func (br *BookRepo) Delete(idStr string) error {
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	err = configs.DB.Model(&models.Book{}).Where("id = ?", id).First(&models.BookReqUpdate{}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "book not found")
	}

	err = configs.DB.Delete(&models.Book{}, id).Error
	if err != nil {
		return err
	}

	return nil
}