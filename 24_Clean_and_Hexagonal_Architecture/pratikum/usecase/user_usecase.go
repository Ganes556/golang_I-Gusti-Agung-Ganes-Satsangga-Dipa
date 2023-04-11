package usecase

import (
	"belajar-go-echo/model"
	"belajar-go-echo/pkg"
	"belajar-go-echo/repository"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserUsecase interface {
	Create(user model.UserReq) error
	GetAll() ([]model.UserRes, error)
	Login(auth model.UserReq) (string, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
	password pkg.Password
}

func NewUserUsecase(userRepo repository.UserRepository, password pkg.Password) *userUsecase {
	return &userUsecase{userRepo, password}
}

func (u *userUsecase) Create(user model.UserReq) error {
	hash, _ := u.password.HashPassword(user.Password)
	user.Password = hash
	err := u.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) GetAll() ([]model.UserRes, error){
	users, err := u.userRepository.FindAll()
	if err != nil {
		return []model.UserRes{}, err
	}
	return users, nil
}

func (u *userUsecase) createToken(id uint, email string) (string, error) {
	claims := &model.JwtClaims{
		ID: id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	return t, nil
}

func (u *userUsecase) Login(auth model.UserReq) (string, error) {
	user, err := u.userRepository.FindByEmail(auth.Email)
	if err != nil {
		return "", err
	}
	
	pass := u.password.CheckPasswordHash(auth.Password, user.Password)
	if !pass {
		return "", errors.New("invalid email or password")
	}

	return u.createToken(user.ID, user.Email)
}