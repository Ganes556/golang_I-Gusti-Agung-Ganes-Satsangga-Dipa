package controllers

import (
	"net/http"
	"strings"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/services"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/utils"
	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) error {
	var books = []models.Book{}
	
	if err := services.FindAll(&books); err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, &echo.Map{
		"message" : "success to get all books",
		"books" : books,
	})
}

func GetBook(c echo.Context) error {
	idStr := c.Param("id")
	
	var id int
	err := utils.Id2Int(idStr, &id)

	if err != nil {
		return err
	}

	var book = models.Book{}

	if err := services.FindById(id,&book); err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success get book by id " + idStr,
		"book": book,
	})
}

func CreateBook(c echo.Context) error {
	var book models.Book
	
	c.Bind(&book)

  if err := c.Validate(&book); err != nil {
    return err
  }
	
	if err := services.Create(&book); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success create new book",
		"book": book,
	})
}

func DeleteBook(c echo.Context) error {

	idStr := c.Param("id")
	
	var id int
	err := utils.Id2Int(idStr, &id)

	if err != nil {
		return err
	}
	
	if err := services.DeleteById(id, &models.Book{}); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}	
	
	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success delete book by id " + idStr,
	})

}


func UpdateBook(c echo.Context) error {
	idStr := c.Param("id")
	var id int
	err := utils.Id2Int(idStr, &id)

	if err != nil {
		return err
	}
	
	var book = models.Book{}
	
	c.Bind(&book)
	if err := c.Validate(&book); err != nil {
		if !strings.Contains(err.Error(),"required") {
			return err
		}
  }
	
	book.ID = uint(id)

	if err := services.UpdateById(id, &book); err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success update book by id " + idStr,
		"book": book,
	})
}