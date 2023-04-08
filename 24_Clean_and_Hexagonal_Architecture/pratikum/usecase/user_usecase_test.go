package usecase_test

import (
	"testing"

	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// create a new instance of the mock repository
	mockRepo := repository.NewMockUserRepo()

	// create a new instance of the usecase, passing the mock repository
	userRepo := usecase.NewUserUsecase(mockRepo)

	// create a new user
	userReq := &model.User{
		Email: "testing@gmail.com",
		Password: "testtingpassword",
	}

	// configure the mock repository to return no error when Create is called
	mockRepo.On("Create", userReq).Return(nil)

	// call the usecase's Create method
	userRepo.Create(userReq)

	// assert that the error is nil
	mockRepo.AssertExpectations(t)
	
	// userRes := &model.User{
	// 	ID: uint(1),
	// 	Email: "testing@gmail.com",
	// 	Password: "testtingpassword",
	// }
}

func TestGetAllUsers(t *testing.T) {
	// create a new instance of the mock repository
	mockRepo := repository.NewMockUserRepo()

	// create a new instance of the usecase, passing the mock repository
	u := usecase.NewUserUsecase(mockRepo)

	// create a slice of user objects
	users := []model.User{
		{
			ID: 1,
			Email: "testing@gmail.com",
			Password: "testtingpassword",
		},
	}

	// configure the mock repository to return the slice of users and no error when FindAll is called
	mockRepo.On("FindAll").Return(users, nil)

	// call the usecase's GetAll method
	result, err := u.GetAll()

	// assert that the error is nil
	assert.Nil(t, err)

	// assert that the result is equal to the slice of users
	assert.Equal(t, users, result)

	// assert that the mock repository's FindAll method was called
	mockRepo.AssertCalled(t, "FindAll")
}

func TestCreateUserError(t *testing.T) {
	// create a new instance of the mock repository
	mockRepo := repository.NewMockUserRepo()

	// create a new instance of the usecase, passing the mock repository
	u := usecase.NewUserUsecase(mockRepo)

	// create a new user
	user := &model.User{
		Email: "testing@gmail.com",
		Password: "testtingpassword",
	}

	// configure the mock repository to return an error when Create is called
	mockRepo.On("Create", user).Return(echo.NewHTTPError(409, "email already exist!"))

	// call the usecase's Create method
	err := u.Create(user)
	// assert that the error is not nil
	assert.NotNil(t, err)
	
	// assert that the mock repository's Create method was called with the user object
	mockRepo.AssertCalled(t, "Create", user)
}

// func TestGetAllUsersError(t *testing.T) {
// 	// create a new instance of the mock repository
// 	mockRepo := new(repository.mockUserRepository)

// 	// create a new instance of the usecase, passing the mock repository
// 	u := usecase.NewUserUsecase(mockRepo)

// 	// configure the mock repository to return an error when FindAll is called
// 	mockRepo.On("Find
