package user

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/dto"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/entities"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/repositories/user"
)



type UserUsecase interface {
	Login(input dto.LoginDTO) (entities.User, error)
	Register(input dto.RegisterDTO) error
}

type userUsecase struct {
	userRepo user.UserRepository
}

func NewUserUsecase(userRepo user.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}