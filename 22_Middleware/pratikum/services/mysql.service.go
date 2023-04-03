package services

import (
	"net/http"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)



func FindAll(data interface{}) error {	
	if err := configs.DB.Find(data).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return nil
}

func FindById(id int, data interface{}) error {

	err := configs.DB.First(&data,id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func FindByEmail(email string, user *models.User) error {
	err := configs.DB.First(user, "email = ?", email).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func FindBlogByUserId(id int, data interface{}) error {
	err := configs.DB.Table("blogs").Select("users.name AS penulis, blogs.*").Joins("LEFT JOIN users ON users.id = blogs.user_refer").Where("users.id = ?", id).Scan(data).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Create(data interface{}) error {
	if user, ok := data.(*models.User); ok {
		if configs.DB.First(user, "email = ?", user.Email).RowsAffected > 0{
			return echo.NewHTTPError(http.StatusBadRequest, "email already exist")
		}
	}else {
		if blog, ok := data.(*models.Blog); ok {

			err := configs.DB.First(&models.User{},"id = ?",blog.UserRefer).Error
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "user not found")
			}

		}
		if configs.DB.Where(data).First(data).RowsAffected > 0{
			return echo.NewHTTPError(http.StatusBadRequest, "data already exist")
		}
	}
	
  if err := configs.DB.Create(data).Error; err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
	return nil
}

func DeleteById(id int, data interface{}) error {
	if err := configs.DB.First(data, id).Unscoped().Delete(data).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var count int64 = -1
  configs.DB.Model(data).Count(&count)

  if count == 0 {
    // reset auto increment
		stmt := &gorm.Statement{DB: configs.DB}
		stmt.Parse(data)
		configs.DB.Exec("ALTER TABLE " + stmt.Schema.Table + " AUTO_INCREMENT = 1")

  }

	return nil
}

func UpdateById(id int, data interface{}) error {
	if blog, ok := data.(*models.Blog); ok {
		err := configs.DB.First(&models.User{},"id = ?",blog.UserRefer).Error
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "user not found")
		}
	} else {
		if user, ok := data.(*models.User); ok {
			if configs.DB.First(user, "email = ?", user.Email).RowsAffected > 0{
				return echo.NewHTTPError(http.StatusBadRequest, "email already exist")
			}
		}
		var count int64
		err := configs.DB.Model(data).Where("id = ?", id).Select("id").Count(&count).Error
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	if err := configs.DB.Where("id = ?", id).Updates(data).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
