package controllers

// import (
// 	"net/http"

// 	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
// 	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/services"
// 	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/utils"
// 	"github.com/labstack/echo/v4"
// )

// func CreateBlog(c echo.Context) error {
	
// 	var blog models.Blog
// 	c.Bind(&blog)
	
// 	if err := c.Validate(&blog); err != nil {
// 		return err
// 	}
	
// 	if err := services.Create(&blog); err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, &echo.Map{
// 		"message": "success create blog",
// 		"blog": blog,
// 	})

// }

// func GetBlogs(c echo.Context) error {
// 	// var blogs []models.Blog
	
// 	blogs, err := services.GetMysqlRepo().FindAll([]models.Blog{})
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, &echo.Map{
// 		"message": "success get blogs",
// 		"blogs": blogs,
// 	})
// }

// func GetBlog(c echo.Context) error {
// 	var blog models.Blog
	
// 	idStr := c.Param("id")
// 	var id int
// 	err := utils.Id2Int(idStr, &id)


// 	if err != nil {
// 		return err
// 	}

// 	if err := services.FindById(id,&blog); err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, &echo.Map{
// 		"message": "success get blog " + idStr,
// 		"blog": blog,
// 	})
// }

// func GetBlogByUserId(c echo.Context) error {
// 	var blogs []struct{
// 		Penulis string `json:"penulis"`
// 		models.Blog
// 	}

// 	idStr := c.Param("id")
// 	var id int
// 	err := utils.Id2Int(idStr, &id)
	
// 	if err != nil {
// 		return err
// 	}

// 	if err := services.FindBlogByUserId(id,&blogs); err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, &echo.Map{
// 		"message": "success get blog by user id " + idStr,
// 		"blogs": blogs,
// 	})
// }

// func UpdateBlog(c echo.Context) error {
// 	var blog models.Blog

// 	idStr := c.Param("id")
// 	var id int
// 	err := utils.Id2Int(idStr, &id)
	
// 	if err != nil {
// 		return err
// 	}

// 	c.Bind(&blog)

// 	if err := services.UpdateById(id,&blog); err != nil {
// 		return err
// 	}

// 	blog.ID = uint(id)

// 	return c.JSON(http.StatusOK, &echo.Map{
// 		"message": "success update blog " + idStr,
// 	})
// }

// func DeleteBlog(c echo.Context) error {
	
// 	idStr := c.Param("id")
// 	var id int
// 	err := utils.Id2Int(idStr, &id)

// 	if err != nil {
// 		return err
// 	}

// 	if err := services.DeleteById(id, &models.Blog{}); err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, &echo.Map{
// 		"message": "success delete blog " + idStr,
// 	})

// }