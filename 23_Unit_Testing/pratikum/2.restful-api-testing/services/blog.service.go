package services

import (
	"net/http"
	"strconv"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/labstack/echo/v4"
)

type IBlogService interface {
	FindAll() ([]models.BlogRes, error)
	FindById(idStr string) (models.BlogRes, error)
	Create(blog models.Blog) error
	Update(idStr string, blog models.BlogReqUpdate) error
	Delete(idStr string) error
	FindBlogByUserId(idStr string) ([]models.BlogByUserId, error)
}

type BlogRepo struct{}

var blogRepo IBlogService

func init() {
	blogRepo = &BlogRepo{}
}

func GetBlogRepo() IBlogService {
	return blogRepo
}

func SetBlogRepo(blr IBlogService) {
	blogRepo = blr
}

func (blr *BlogRepo) FindAll() ([]models.BlogRes, error) {
	var blogs []models.BlogRes
	err := configs.DB.Model(&models.Blog{}).Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (blr *BlogRepo) FindById(idStr string) (models.BlogRes, error) {

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return models.BlogRes{}, echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	var blog models.BlogRes
	err = configs.DB.Model(&models.Blog{}).Where("id = ?", id).First(&blog).Error
	if err != nil {
		return models.BlogRes{}, echo.NewHTTPError(http.StatusNotFound, "blog not found")
	}
	return blog, nil
}

func (blr *BlogRepo) Create(blog models.Blog) error {
	userFound := configs.DB.Where("id = ?", blog.UserRefer).First(&models.User{}).RowsAffected > 0

	if !userFound {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	found := configs.DB.Where("judul = ? OR konten = ?", blog.Judul, blog.Konten).First(&models.Blog{}).RowsAffected > 0

	if found {
		return echo.NewHTTPError(http.StatusConflict, "blog already exist")
	}

	err := configs.DB.Create(&blog).Error
	if err != nil {
		return err
	}
	return nil
}

func (blr *BlogRepo) Update(idStr string, blog models.BlogReqUpdate) error {
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	err = configs.DB.Model(&models.Blog{}).Where("id = ?", id).First(&models.BlogReqUpdate{}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "blog not found")
	}

	err = configs.DB.Table("Blogs").Where("id = ?", id).Updates(&blog).Error
	
	if err != nil {
		return err
	}
	return nil
}

func (blr *BlogRepo) Delete(idStr string) error {
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	err = configs.DB.Model(&models.Blog{}).Where("id = ?", id).First(&models.BlogReqUpdate{}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "blog not found")
	}

	err = configs.DB.Delete(&models.Blog{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (blr *BlogRepo) FindBlogByUserId(idStr string) ([]models.BlogByUserId, error) {
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	var blogs []models.BlogByUserId

	err = configs.DB.Table("blogs").Select("users.name AS penulis, blogs.judul, blogs.konten").Joins("LEFT JOIN users ON users.id = blogs.user_refer").Where("users.id = ?", id).Scan(&blogs).Error
	if err != nil {
		return blogs, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return blogs, nil
}