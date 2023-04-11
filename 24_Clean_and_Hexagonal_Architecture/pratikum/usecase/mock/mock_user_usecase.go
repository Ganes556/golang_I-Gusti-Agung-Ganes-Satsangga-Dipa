package usecase

import (
	"belajar-go-echo/model"

	"github.com/stretchr/testify/mock"
)

type mockUserUsecase struct {
	mock.Mock
}

func NewMockUserUsecase() *mockUserUsecase {
	return &mockUserUsecase{}
}

func (m *mockUserUsecase) Create(user model.UserReq) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserUsecase) GetAll() ([]model.UserRes, error) {
	args := m.Called()
	return args.Get(0).([]model.UserRes), args.Error(1)
}

func (m *mockUserUsecase) Login(auth model.UserReq) (string, error) {
	args := m.Called(auth)
	return args.Get(0).(string), args.Error(1)
}
