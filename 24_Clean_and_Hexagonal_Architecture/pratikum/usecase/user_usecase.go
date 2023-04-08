package usecase

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type UserUsecase interface {
	Create(user *model.User) error
	GetAll() ([]model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) Create(user *model.User) error {
	return u.userRepository.Create(user)
}

func (u *userUsecase) GetAll() ([]model.User, error){
	return u.userRepository.FindAll()
}