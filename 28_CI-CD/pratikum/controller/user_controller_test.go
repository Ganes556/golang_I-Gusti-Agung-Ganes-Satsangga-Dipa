package controller

import (
	"belajar-go-echo/model"
	usecase "belajar-go-echo/usecase/mock"
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {

	mockUserUsecase := usecase.NewMockUserUsecase()
	userController := NewUserController(mockUserUsecase)
	
	tests := []struct {
		name     string
		expectedCode int
		expectedBody map[string]interface{}
		wantErr  bool
	}{
		{
			name: "success",
			expectedCode: 200,
			expectedBody: map[string]interface{}{
				"data": []model.UserRes{
					{
						ID:    1,
						Email: "testing@gmail.com",
					},
				},
			},
			wantErr: false,
		},
		{
			name:     "Failed - get all users",
			expectedCode: 500,
			expectedBody: map[string]interface{}{
				"message": "internal server error",
			},
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(echo.GET, "/users", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			

			
			if test.wantErr {
				mockUserUsecase.On("GetAll").Return([]model.UserRes{}, errors.New(test.expectedBody["message"].(string))).Once()
				err := userController.GetAllUsers(c)
				assert.NotNil(t, err)

				httpErr, ok := err.(*echo.HTTPError)
				assert.True(t, ok)
				assert.Equal(t, test.expectedCode, httpErr.Code)
				assert.Equal(t, test.expectedBody["message"], httpErr.Message)
				
			} else {
				mockUserUsecase.On("GetAll").Return(test.expectedBody["data"].([]model.UserRes), nil).Once()
				err := userController.GetAllUsers(c)
				assert.Nil(t, err)
				assert.Equal(t, test.expectedCode, rec.Code)

				type response struct {
					Data    []model.UserRes `json:"data"`
				}	
				var resp response
				json.Unmarshal(rec.Body.Bytes(), &resp)
				assert.Equal(t, test.expectedBody["data"], resp.Data)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {

	mockUserUsecase := usecase.NewMockUserUsecase()
	userController := NewUserController(mockUserUsecase)
	
	tests := []struct {
		name     string
		payload model.UserReq
		expectedCode int
		expectedBody map[string]interface{}
		wantErr  bool
	}{
		{
			name: "success",
			payload: model.UserReq{
				Email: "testing@gmail.com",
				Password: "testingpassword",
			},
			expectedCode: 201,
			expectedBody: map[string]interface{}{
				"message": "success create user",
			},
			wantErr: false,
		},
		{
			name:     "Failed - user already exist",
			payload: model.UserReq{
				Email: "testing@gmail.com",
				Password: "testingpassword",
			},
			expectedCode: 409,
			expectedBody: map[string]interface{}{
				"message": "user already exist",
			},
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()
			
			payload, err := json.Marshal(test.payload)
			assert.NoError(t, err)
			req := httptest.NewRequest(echo.POST, "/users", bytes.NewReader(payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			
			
			if test.wantErr {
				mockUserUsecase.On("Create", test.payload).Return(errors.New(test.expectedBody["message"].(string))).Once()
				err := userController.CreateUser(c)
				assert.NotNil(t, err)

				httpErr, ok := err.(*echo.HTTPError)
				assert.True(t, ok)
				assert.Equal(t, test.expectedCode, httpErr.Code)
				assert.Equal(t, test.expectedBody["message"], httpErr.Message)
				
			} else {
				mockUserUsecase.On("Create",test.payload).Return(nil).Once()
				err := userController.CreateUser(c)
				assert.Nil(t, err)
				assert.Equal(t, test.expectedCode, rec.Code)
				
				type response struct {
					Message string `json:"message"`
				}	
				var resp response
				json.Unmarshal(rec.Body.Bytes(), &resp)
				assert.Equal(t, test.expectedBody["message"], resp.Message)
			}
		})
	}
}
func TestLoginUser(t *testing.T) {

	mockUserUsecase := usecase.NewMockUserUsecase()
	userController := NewUserController(mockUserUsecase)
	
	tests := []struct {
		name     string
		payload model.UserReq
		expectedCode int
		expectedBody map[string]interface{}
		wantErr  bool
	}{
		{
			name: "success",
			payload: model.UserReq{
				Email: "testing@gmail.com",
				Password: "testingpassword",
			},
			expectedCode: 200,
			expectedBody: map[string]interface{}{
				"token": "token",
			},
			wantErr: false,
		},
		{
			name:     "Failed - invalid auth",
			payload: model.UserReq{
				Email: "testing@gmail.com",
				Password: "wrongpassword",
			},
			expectedCode: 401,
			expectedBody: map[string]interface{}{
				"message": "invalid email or password",
			},
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()
			
			payload, err := json.Marshal(test.payload)
			assert.NoError(t, err)
			req := httptest.NewRequest(echo.POST, "/users/login", bytes.NewReader(payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			
			if test.wantErr {
				mockUserUsecase.On("Login", test.payload).Return("",errors.New(test.expectedBody["message"].(string))).Once()
				err := userController.LoginUser(c)
				assert.NotNil(t, err)

				httpErr, ok := err.(*echo.HTTPError)
				assert.True(t, ok)
				assert.Equal(t, test.expectedCode, httpErr.Code)
				assert.Equal(t, test.expectedBody["message"], httpErr.Message)
				
			} else {
				mockUserUsecase.On("Login", test.payload).Return("token",nil).Once()
				err := userController.LoginUser(c)
				assert.Nil(t, err)
				assert.Equal(t, test.expectedCode, rec.Code)
				
				type response struct {
					Token string `json:"token"`
				}	
				var resp response
				json.Unmarshal(rec.Body.Bytes(), &resp)
				assert.Equal(t, test.expectedBody["token"], resp.Token)
			}
		})
	}
}
