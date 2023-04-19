package services

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) FindAll() ([]models.UserRes, error){
	args := m.Called()
	return args.Get(0).([]models.UserRes), args.Error(1)
}

func (m *MockUserRepo) FindById(idStr string) (models.UserRes, error){
	args := m.Called(idStr)
	return args.Get(0).(models.UserRes), args.Error(1)
}

func (m *MockUserRepo) FindByEmail(email string) (models.UserResDB, error){
	args := m.Called(email)
	return args.Get(0).(models.UserResDB), args.Error(1)
}

func (m *MockUserRepo) Create(user models.User) error{
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepo) Update(idStr string, user models.UserReqUpdate) error{
	args := m.Called(idStr,user)
	return args.Error(0)
}

func (m *MockUserRepo) Delete(idStr string) error{
	args := m.Called(idStr)
	return args.Error(0)
}