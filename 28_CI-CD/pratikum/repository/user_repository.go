package repository

import (
	"belajar-go-echo/model"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user model.UserReq) error
	FindAll() ([]model.UserRes, error)
	FindByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(user model.UserReq) error {
	found := u.db.Table("users").Where("email = ?", user.Email).First(&model.UserReq{}).RowsAffected > 0
	if found {
		return errors.New("user already exist")
	}
	
	if err := u.db.Table("users").Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) FindAll() ([]model.UserRes,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	
	defer cancel()

	var users []model.UserRes

	err := u.db.WithContext(ctx).Table("users").Find(&users).Error

	if err != nil {
		return users, err
	}
	return users, nil
}

func (u *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	found := u.db.Table("users").Where("email = ?", email).First(&user).RowsAffected > 0
	if !found {
		return model.User{} , errors.New("user not found")
	}
	return user, nil
}
