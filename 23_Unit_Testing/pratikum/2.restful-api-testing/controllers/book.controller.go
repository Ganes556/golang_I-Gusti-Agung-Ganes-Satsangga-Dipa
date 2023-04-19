package controllers

import (
	"net/http"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/services"
	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) error {
	books, err := services.GetBookRepo().FindAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success to get all books",
		"books": books,
	})
}

func GetBook(c echo.Context) error {
	idStr := c.Param("id")

	book, err := services.GetBookRepo().FindById(idStr)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success get book by id " + idStr,
		"book":    book,
	})
}

func CreateBook(c echo.Context) error {
	var book models.Book

	c.Bind(&book)

  if err := c.Validate(echo.Map{"method":c.Request().Method,"data":&book}); err != nil {
    return err
  }
	
	err := services.GetBookRepo().Create(book)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &echo.Map{
		"message": "success create new book",
	})
}

func DeleteBook(c echo.Context) error {

	idStr := c.Param("id")

	err := services.GetBookRepo().Delete(idStr)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success delete book by id " + idStr,
	})

}

func UpdateBook(c echo.Context) error {
	idStr := c.Param("id")

	var book = models.BookReqUpdate{}

	c.Bind(&book)

	if err := c.Validate(echo.Map{"method":c.Request().Method,"data":&book}); err != nil {
		return err
  }

	err := services.GetBookRepo().Update(idStr, book)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success update book by id " + idStr,
	})
}