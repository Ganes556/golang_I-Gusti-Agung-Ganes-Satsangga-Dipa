package pkg

import "github.com/stretchr/testify/mock"

type mockPassword struct {
	mock.Mock
}

func NewMockPassword() *mockPassword {
	return &mockPassword{}
}

func (m *mockPassword) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.Get(0).(string), args.Error(1)
}

func (m *mockPassword) CheckPasswordHash(password, hash string) bool {
	args := m.Called(password, hash)
	return args.Bool(0)
}