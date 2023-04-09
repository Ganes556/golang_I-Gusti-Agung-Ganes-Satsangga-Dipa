package services

import (
	"net/http"
	"strconv"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/labstack/echo/v4"
)

type IUserService interface {
	FindAll() ([]models.UserRes, error)
	FindById(idStr string) (models.UserRes, error)
	FindByEmail(email string) (models.User, error)
	Create(user models.User) error
	Update(idStr string, user models.UserReqUpdate) error
	Delete(idStr string) error
}

type UserRepo struct {
	Func IUserService
}

var userRepo IUserService

func init() {
	ur := &UserRepo{}
	userRepo = ur
}

func GetUserRepo() IUserService {
	return userRepo
}

func SetUserRepo(ur IUserService) {
	userRepo = ur
}

func (ur *UserRepo) FindAll() ([]models.UserRes, error){
	var users []models.UserRes
	err := configs.DB.Model(&models.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepo) FindById(idStr string) (models.UserRes, error){

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return models.UserRes{}, echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	var user models.UserRes
	err = configs.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.UserRes{}, echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return user, nil
}

func (ur *UserRepo) FindByEmail(email string) (models.User, error){
	var user models.User
	err := configs.DB.Select("name","email","password").Model(&models.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return user, nil
}

func (ur *UserRepo) Create(user models.User) error{
	
	rowAffect := configs.DB.Model(&models.User{}).Where("email = ?", user.Email).First(&models.User{}).RowsAffected
	if rowAffect > 0 {
		return echo.NewHTTPError(http.StatusConflict, "email already exist")
	}

	err := configs.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) Update(idStr string, user models.UserReqUpdate) error{
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")
	}

	err = configs.DB.Model(&models.User{}).Where("id = ?", id).First(&models.UserReqUpdate{}).Error
	
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	// ignore column `password`
	err = configs.DB.Model(&models.User{}).Where("id = ?", id).Omit("password").Updates(&user).Error

	if err != nil {
		return err
	}
	return nil
}


func (ur *UserRepo) Delete(idStr string) error{
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be number")	
	}

	err = configs.DB.Model(&models.User{}).Where("id = ?", id).First(&models.UserReq{}).Error

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	err = configs.DB.Delete(&models.User{}, id).Error
	if err != nil {
		return err	
	}

	return nil
}




