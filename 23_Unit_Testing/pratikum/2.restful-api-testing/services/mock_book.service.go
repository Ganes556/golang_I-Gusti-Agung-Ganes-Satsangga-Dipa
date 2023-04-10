package services

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/stretchr/testify/mock"
)

type MockBookRepo struct {
	mock.Mock
}

func (m *MockBookRepo) FindAll() ([]models.BookRes, error) {
	args := m.Called()
	return args.Get(0).([]models.BookRes), args.Error(1)
}

func (m *MockBookRepo) FindById(idStr string) (models.BookRes, error) {
	args := m.Called(idStr)
	return args.Get(0).(models.BookRes), args.Error(1)
}

func (m *MockBookRepo) Create(book models.Book) error {
	
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepo) Update(idStr string, book models.BookReqUpdate) error {
	args := m.Called(idStr, book)
	return args.Error(0)
}

func (m *MockBookRepo) Delete(idStr string) error {
	args := m.Called(idStr)
	return args.Error(0)
}