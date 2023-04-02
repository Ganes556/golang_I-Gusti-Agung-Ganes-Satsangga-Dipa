package controllers

import (
	"net/http"
	"strconv"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/labstack/echo/v4"
)

func CreateBlogController(c echo.Context) error {
	
	var blog models.Blog
	c.Bind(&blog)
	
	if err := c.Validate(&blog); err != nil {
		return err
	}
	
	if err := configs.DB.First(&models.User{}, "id = ?", blog.UserRefer).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "user not found")
	}

	if configs.DB.Where(&blog).First(&blog).RowsAffected > 0{
		return echo.NewHTTPError(http.StatusBadRequest, "blog already exist")
	}
	
	if err := configs.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create blog",
		"blog": blog,
	})

}

func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	if err := configs.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blogs",
		"blogs": blogs,
	})
}

func GetBlogController(c echo.Context) error {
	var blog models.Blog
	
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	if err := configs.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog " + idStr,
		"blog": blog,
	})
}

func UpdateBlogController(c echo.Context) error {
	var blog models.Blog

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	if err := configs.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&blog)

	if err := configs.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog " + idStr,
		"blog": blog,
	})
}

func DeleteBlogController(c echo.Context) error {
	
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	if err := configs.DB.Delete(&models.Blog{}, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var count int64 = -1

	configs.DB.Model(&models.Blog{}).Count(&count)

	if count == 0 {
		configs.DB.Exec("ALTER TABLE blogs AUTO_INCREMENT = 1")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog " + idStr,
	})

}