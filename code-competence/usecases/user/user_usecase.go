package user

import (
	"errors"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/dto"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/entities"
	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) Login(input dto.LoginDTO) (entities.User, error) {

	
	user, err := u.userRepo.GetByEmail(input.Email)

	if err != nil {
		return user, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return user, errors.New("invalid email or password")
	}

	return user, nil
}

func (u *userUsecase) Register(input dto.RegisterDTO) error {
	
	_, err := u.userRepo.GetByEmail(input.Email)

	if err == nil {
		return errors.New("email already exist")
	}

	user := entities.User{
		Email: input.Email,
		Password: input.Password,
		Name: input.Name,
	}

	err = u.userRepo.Create(user)

	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}