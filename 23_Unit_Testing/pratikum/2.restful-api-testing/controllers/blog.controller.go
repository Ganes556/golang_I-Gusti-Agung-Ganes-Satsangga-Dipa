package controllers

import (
	"net/http"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/services"
	"github.com/labstack/echo/v4"
)

func CreateBlog(c echo.Context) error {
	
	var blog models.Blog
	c.Bind(&blog)
	
	if err := c.Validate(echo.Map{"method":c.Request().Method,"data":&blog}); err != nil {
		return err
	}
	
	if err := services.GetBlogRepo().Create(blog); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &echo.Map{
		"message": "success create new blog",
	})

}

func GetBlogs(c echo.Context) error {
	
	blogs, err := services.GetBlogRepo().FindAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success to get all blogs",
		"blogs": blogs,
	})
}

func GetBlog(c echo.Context) error {
	
	idStr := c.Param("id")
	
	blog, err := services.GetBlogRepo().FindById(idStr)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success get blog by id " + idStr,
		"blog": blog,
	})
}

func GetBlogByUserId(c echo.Context) error {
	

	idStr := c.Param("id")

	blogs, err := services.GetBlogRepo().FindBlogByUserId(idStr)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success get blogs by user id " + idStr,
		"blogs": blogs,
	})
}

func UpdateBlog(c echo.Context) error {
	var blog models.BlogReqUpdate

	idStr := c.Param("id")

	c.Bind(&blog)

	if err := c.Validate(echo.Map{"method":c.Request().Method,"data":&blog}); err != nil {
		return err
	}

	if err := services.GetBlogRepo().Update(idStr, blog); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success update blog by id " + idStr,
	})
}

func DeleteBlog(c echo.Context) error {
	
	idStr := c.Param("id")

	if err := services.GetBlogRepo().Delete(idStr); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &echo.Map{
		"message": "success delete blog by id " + idStr,
	})

}