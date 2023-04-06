package controllers

import (
	"net/http"
	"strconv"

	conf "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) error {
	var books = []models.Book{}

	conf.DB.Find(&books)

	return c.JSON(http.StatusOK, echo.Map{
		"message" : "success to get all books",
		"books" : books,
	})
}

func GetBook(c echo.Context) error {
	idStr := c.Param("id")
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}
	
	var book models.Book
	
	if err := conf.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "book not found")
	}
	
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get book by id " + idStr,
		"book": book,
	})
}

func CreateBook(c echo.Context) error {
	var book models.Book
	
	c.Bind(&book)

  if err := c.Validate(echo.Map{"method":c.Request().Method,"data":&book}); err != nil {
    return err
  }
	
	if conf.DB.Where(&book).First(&models.Book{}).RowsAffected > 0{
    return echo.NewHTTPError(http.StatusBadRequest, "book already exist")
  }

	conf.DB.Create(&book)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create new book",
		"book": book,
	})
}

func DeleteBook(c echo.Context) error {

	idStr := c.Param("id")
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}
	
	if err := conf.DB.First(&models.Book{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "book not found")
	}

	conf.DB.Unscoped().Delete(&models.Book{}, id)
	
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete book by id " + idStr,
	})

}


func UpdateBook(c echo.Context) error {
	idStr := c.Param("id")
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}
	
	var book = models.Book{}
	
	if err := conf.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	
	c.Bind(&book)

	if err := c.Validate(echo.Map{"method":c.Request().Method,"data":&book}); err != nil {
		return err
	}
	
	conf.DB.Model(&book).Where("id = ?", id).Updates(&book)
	
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update book by id " + idStr,
	})
}