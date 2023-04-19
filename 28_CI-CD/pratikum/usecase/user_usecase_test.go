package usecase

import (
	"errors"
	"testing"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/model"
	pkg "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/pkg/mock"
	repository "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/repository/mock"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	mockRepo := repository.NewMockUserRepo()
	mockPassword := pkg.NewMockPassword()
	
	userRepo := NewUserUsecase(mockRepo, mockPassword)
	tests := []struct{
		name string
		payload model.UserReq
		errMessage string
		wantErr bool
	}{
		{
			name: "success",
			payload: model.UserReq{
				Email:    "testing@gmail.com",
				Password: "testtingpassword",
			},
			wantErr: false,
		},
		{
			name: "failed - email already exist",
			payload: model.UserReq{
				Email: "testing@gmail.com",
				Password: "testtingpassword",
			},
			errMessage: "email already exist",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			mockPassword.On("HashPassword", test.payload.Password).Return(test.payload.Password, nil).Once()
			if test.wantErr {
				mockRepo.On("Create", test.payload).Return(errors.New(test.errMessage)).Once()
				err := userRepo.Create(test.payload)
				assert.NotNil(t, err)
				assert.Equal(t, test.errMessage, err.Error())

			}else {
				mockRepo.On("Create", test.payload).Return(nil).Once()
				err := userRepo.Create(test.payload)
				assert.Nil(t, err)
			}
		})
	}

}

func TestGetAllUsers(t *testing.T) {

	mockRepo := repository.NewMockUserRepo()
	mockPassword := pkg.NewMockPassword()
	userRepo := NewUserUsecase(mockRepo, mockPassword)
	
	tests := []struct{
		name string
		expected []model.UserRes
		wantErr bool
	}{
		{
			name: "success",
			expected: []model.UserRes{
				{
					ID: 1,
					Email: "testing@gmail.com",
				},
			},
			wantErr: false,
		},
		{
			name: "Failed - error",
			expected: []model.UserRes{},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.wantErr {
				mockRepo.On("FindAll").Return(test.expected, errors.New("error")).Once()
				users, err := userRepo.GetAll()
				assert.NotNil(t, err)
				assert.Equal(t, []model.UserRes{}, users)
			}else {
				mockRepo.On("FindAll").Return(test.expected, nil).Once()
				users, err := userRepo.GetAll()
				assert.Nil(t, err)
				assert.Equal(t, test.expected, users)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {

	mockRepo := repository.NewMockUserRepo()
	mockPassword := pkg.NewMockPassword()
	userRepo := NewUserUsecase(mockRepo, mockPassword)
	
	tests := []struct{
		name string
		payload model.UserReq
		hashPassword string
		wantErr bool
		errMessage string
	}{
		{
			name: "success",
			payload: model.UserReq{
				Email:    "testing@gmail.com",
				Password: "testtingpassword",
			},
			hashPassword: "$2y$10$Vn4AG3B4MchRuNPlfuOfMuASVJmAYjWznGVihdvS3f.TAbJLliL3W",
			wantErr: false,
		},
		{
			name: "failed - email not found",
			payload: model.UserReq{
				Email:    "testing2@gmail.com",
				Password: "testtingpassword",
			},
			wantErr: true,
			errMessage: "email not found",
		},
		{
			name: "failed - invalid email or password",
			payload: model.UserReq{
				Email:    "testing@gmail.com",
				Password: "testtingpasswordtest",
			},
			hashPassword: "$2y$10$z8sKOVJFS73YbwdBBME0nO1StO3uOB2K/89QUYAj1aK.z0fNrOFW6",
			wantErr: true,
			errMessage: "invalid email or password",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.wantErr {
				switch test.errMessage {
					case "email not found":
						mockRepo.On("FindByEmail", test.payload.Email).Return(model.User{}, errors.New(test.errMessage)).Once()
					case "invalid email or password":
						mockRepo.On("FindByEmail", test.payload.Email).Return(model.User{
							ID: 1,
							Email: test.payload.Email,
							Password: test.hashPassword,
						}, nil).Once()
						mockPassword.On("CheckPasswordHash", test.payload.Password, test.hashPassword).Return(false).Once()
				}

				token, err := userRepo.Login(test.payload)
				assert.NotNil(t, err)
				assert.Equal(t, "", token)
				assert.Equal(t, test.errMessage, err.Error())

			}else {
				mockRepo.On("FindByEmail", test.payload.Email).Return(model.User{
					ID: 1,
					Email: test.payload.Email,
					Password: test.hashPassword,
				}, nil).Once()
				
				mockPassword.On("CheckPasswordHash", test.payload.Password, test.hashPassword).Return(true).Once()

				token, err := userRepo.Login(test.payload)
				
				assert.Nil(t, err)
				assert.NotEqual(t, "", token)

			}
		})
	}

}
