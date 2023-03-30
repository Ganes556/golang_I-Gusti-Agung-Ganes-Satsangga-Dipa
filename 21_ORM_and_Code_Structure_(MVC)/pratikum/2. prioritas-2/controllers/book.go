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
	
	
	if err := conf.DB.Find(&books).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
	
	return c.JSON(http.StatusOK, map[string]interface{}{
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
	
	var book = models.Book{}
	
	if err := conf.DB.First(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
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
	
	if err := conf.DB.Where(&book).First(&book).Error; err == nil {
    return echo.NewHTTPError(http.StatusBadRequest, "book already exist")
  }

	if err := conf.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
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
	
	var book = models.Book{}
	
	if err = conf.DB.First(&models.Book{}, id).Unscoped().Delete(&models.Book{}).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}	
	var count int64 = -1
	conf.DB.Model(&models.Book{}).Count(&count)

	if count == 0 { 
		conf.DB.Exec("ALTER TABLE books AUTO_INCREMENT = 1")
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book by id " + idStr,
		"book": book,
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

	if err := c.Validate(&book); err != nil {
		return err
	}
	
	if err := conf.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book by id " + idStr,
		"book": book,
	})
}