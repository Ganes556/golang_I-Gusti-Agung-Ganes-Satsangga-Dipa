package services

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/stretchr/testify/mock"
)

type MockBlogRepo struct {
	mock.Mock
}

func (m *MockBlogRepo) FindBlogByUserId(idStr string) ([]models.BlogByUserId, error){
	args := m.Called(idStr)
	return args.Get(0).([]models.BlogByUserId), args.Error(1)
}

func (m *MockBlogRepo) FindAll() ([]models.BlogRes, error) {
	args := m.Called()
	return args.Get(0).([]models.BlogRes), args.Error(1)
}

func (m *MockBlogRepo) FindById(idStr string) (models.BlogRes, error) {
	args := m.Called(idStr)
	return args.Get(0).(models.BlogRes), args.Error(1)
}

func (m *MockBlogRepo) Create(Blog models.Blog) error {

	args := m.Called(Blog)
	return args.Error(0)
}

func (m *MockBlogRepo) Update(idStr string, Blog models.BlogReqUpdate) error {
	args := m.Called(idStr, Blog)
	return args.Error(0)
}

func (m *MockBlogRepo) Delete(idStr string) error {
	args := m.Called(idStr)
	return args.Error(0)
}

