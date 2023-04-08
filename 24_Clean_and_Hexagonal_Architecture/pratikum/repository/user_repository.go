package repository

import (
	"belajar-go-echo/model"
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindAll() ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(user *model.User) error {
	if u.db.Model(user).Where("email = ?", user.Email).Updates(user).RowsAffected == 0 {
		if err := u.db.Create(user).Error; err != nil {
			return echo.NewHTTPError(500, err.Error())
		}
		return nil
	}
	return echo.NewHTTPError(409, "email already exist!")
}

func (u *userRepository) FindAll() ([]model.User,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
		
	var users []model.User
	err := u.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return users, echo.NewHTTPError(500, err.Error())
	}
	return users, nil
}