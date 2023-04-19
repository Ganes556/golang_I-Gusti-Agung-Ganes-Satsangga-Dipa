package repository

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/model"

	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func NewMockUserRepo() *mockUserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) Create(user model.UserReq) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *mockUserRepository) FindAll() ([]model.UserRes, error) {
	args := m.Called()
	return args.Get(0).([]model.UserRes), args.Error(1)
}

func (m *mockUserRepository) FindByEmail(email string) (model.User, error){
	args := m.Called(email)
	return args.Get(0).(model.User) , args.Error(1)
}
