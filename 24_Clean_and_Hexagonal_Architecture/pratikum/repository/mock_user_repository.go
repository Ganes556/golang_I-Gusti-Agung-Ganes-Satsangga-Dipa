package repository

import (
	"belajar-go-echo/model"

	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func NewMockUserRepo() *mockUserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) Create(user *model.User) error {
	ret := m.Called(user)
	return ret.Error(0)
}
func (m *mockUserRepository) FindAll() ([]model.User, error) {
    ret := m.Called()
    return ret.Get(0).([]model.User), ret.Error(1)
}