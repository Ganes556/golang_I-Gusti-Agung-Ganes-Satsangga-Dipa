package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/libs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/services"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	mockUserRepo := services.MockUserRepo{}
	services.SetUserRepo(&mockUserRepo)

	tests := []struct {
		name         string
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			expectedBody: echo.Map{
				"message": "success get all users",
				"users": []models.UserRes{
					{
						ID:    1,
						Name:  "test",
						Email: "tes@gmail.com",
					},
				},
			},
			expectedCode: http.StatusOK,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(echo.GET, "/users", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if !tt.wantErr {
				mockUserRepo.On("FindAll").Return(tt.expectedBody["users"], nil)
			}else {
				mockUserRepo.On("FindAll").Return(tt.expectedCode, echo.NewHTTPError(http.StatusInternalServerError, tt.expectedBody["message"]))

			}

			err := GetUsers(c)
			if tt.wantErr {
				assert.Error(t, err)
				httpErr, ok := err.(*echo.HTTPError)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCode, httpErr.Code)
				assert.Equal(t, tt.expectedBody["message"], httpErr.Message)
			}else {

				assert.NoError(t, err)
				var ret struct {
					Message string           `json:"message"`
					Users   []models.UserRes `json:"users"`
				}
	
				err = json.Unmarshal(rec.Body.Bytes(), &ret)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.Equal(t, tt.expectedBody["message"], ret.Message)
				assert.Equal(t, tt.expectedBody["users"], ret.Users)
			}

		})
	}
}

func TestGetUser(t *testing.T) {
	mockUserRepo := services.MockUserRepo{}
	services.SetUserRepo(&mockUserRepo)
	
	tests := []struct {
		name         string
		userId 			 string
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			userId: "1",
			expectedBody: echo.Map{
				"message": "success get user by id 1",
				"user": models.UserRes{
					ID:    1,
					Name:  "test",
					Email: "tes@gmail.com",
				},
			},
			expectedCode: http.StatusOK,
			wantErr:      false,
		},
		{
			name: "Failed - Invalid Id",
			userId: "abcd",
			expectedBody: echo.Map{
				"message": "id must be number",
			},
			expectedCode: http.StatusBadRequest,
			wantErr:      true,
		},
		{
			name: "Failed - User Not Found",
			userId: "2",
			expectedBody: echo.Map{
				"message": "user not found",
			},
			expectedCode: http.StatusNotFound,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(echo.GET, "/users", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.userId)
			
			if !tt.wantErr {
				mockUserRepo.On("FindById", tt.userId).Return(tt.expectedBody["user"], nil)
			} else {
				mockUserRepo.On("FindById", tt.userId).Return(models.UserRes{}, echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string)))
			}

			err := GetUser(c)

			if tt.wantErr {
				assert.Error(t, err)
				httpErr, ok := err.(*echo.HTTPError)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCode, httpErr.Code)
				assert.Equal(t, tt.expectedBody["message"], httpErr.Message)
			} else {
				assert.NoError(t, err)

				var ret struct {
					Message string         `json:"message"`
					User    models.UserRes `json:"user"`
				}

				err = json.Unmarshal(rec.Body.Bytes(), &ret)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.Equal(t, tt.expectedBody["message"], ret.Message)
				assert.Equal(t, tt.expectedBody["user"], ret.User)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	mockUserRepo := services.MockUserRepo{}
	services.SetUserRepo(&mockUserRepo)

	tests := []struct {
		name         string
		payload      models.User
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		{
			name: "Success",
			payload: models.User{
				Name: "test2",
				Email: "test2@gmail.com",
				Password: "test2123456",
			},
			expectedBody: echo.Map{
				"message": "success create new user",
			},
			expectedCode: http.StatusCreated,
			wantErr: false,
		},
		{
			name: "Failed - Invalid Email",
			payload: models.User{
				Name: "test2",
				Email: "test2gmail.com",
				Password: "test2123456",
			},
			expectedBody: echo.Map{
				"message": "email is not valid",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - Invalid Password",
			payload: models.User{
				Name: "test2",
				Email: "test2@gmail.com",
				Password: "test",
			},
			expectedBody: echo.Map{
				"message": "password must be at least 8 characters",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - Email Conflict",
			payload: models.User{
				Name: "test",
				Email: "test@gmail.com",
				Password: "test2123456",
			},
			expectedBody: echo.Map{
				"message": "email already exist",
			},
			expectedCode: http.StatusConflict,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			e.Validator = libs.NewValidator()
			payload, err := json.Marshal(tt.payload)
			
			assert.NoError(t, err)
		
			req := httptest.NewRequest(echo.POST, "/users", bytes.NewReader(payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			
			if !tt.wantErr {
				mockUserRepo.On("Create", tt.payload).Return(nil)
			}else {
				mockUserRepo.On("Create", tt.payload).Return(echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string)))
			}

			err = CreateUser(c)

			if tt.wantErr {
				assert.Error(t, err)
				httpErr, ok := err.(*echo.HTTPError)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCode, httpErr.Code)
				assert.Equal(t, tt.expectedBody["message"], httpErr.Message)
			} else {
				assert.NoError(t, err)

				var ret struct {
					Message string `json:"message"`
				}
				
				err = json.Unmarshal(rec.Body.Bytes(), &ret)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.Equal(t, tt.expectedBody["message"], ret.Message)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	mockUserRepo := services.MockUserRepo{}
	services.SetUserRepo(&mockUserRepo)

	tests := []struct {
		name         string
		userId	     string
		payload      models.UserReqUpdate
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		{
			name: "Success",
			userId: "1",
			payload: models.UserReqUpdate{
				Name: "test3",
				Email: "test3@gmail.com",
			},
			expectedBody: echo.Map{
				"message": "success update user by id 1",
			},
			expectedCode: http.StatusOK,
			wantErr: false,
		},
		{
			name: "Failed - Invalid Email",
			userId: "1",
			payload: models.UserReqUpdate{
				Name: "test2",
				Email: "test2gmail.com",
			},
			expectedBody: echo.Map{
				"message": "email is not valid",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - Invalid Id",
			userId: "abcd",
			payload: models.UserReqUpdate{
				Name: "test2",
				Email: "test2@gmail.com",
			},
			expectedBody: echo.Map{
				"message": "id must be a number",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - User Not Found",
			userId: "999",
			payload: models.UserReqUpdate{
				Name: "test2",
				Email: "test2@gmail.com",
			},
			expectedBody: echo.Map{
				"message": "user not found",
			},
			expectedCode: http.StatusNotFound,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			e.Validator = libs.NewValidator()
			payload, err := json.Marshal(tt.payload)
			
			assert.NoError(t, err)
		
			req := httptest.NewRequest(echo.PUT, "/users", bytes.NewReader(payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.userId)
			
			if !tt.wantErr {
				mockUserRepo.On("Update",tt.userId ,tt.payload).Return(nil)
			}else {
				mockUserRepo.On("Update",tt.userId , tt.payload).Return(echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string)))
			}

			err = UpdateUser(c)

			if tt.wantErr {
				assert.Error(t, err)
				httpErr, ok := err.(*echo.HTTPError)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCode, httpErr.Code)
				assert.Equal(t, tt.expectedBody["message"], httpErr.Message)
			} else {
				assert.NoError(t, err)

				var ret struct {
					Message string `json:"message"`
				}
				
				err = json.Unmarshal(rec.Body.Bytes(), &ret)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.Equal(t, tt.expectedBody["message"], ret.Message)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	mockUserRepo := services.MockUserRepo{}
	services.SetUserRepo(&mockUserRepo)
	
	tests := []struct {
		name         string
		userId 			 string
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			userId: "1",
			expectedBody: echo.Map{
				"message": "success delete user by id 1",
			},
			expectedCode: http.StatusOK,
			wantErr:      false,
		},
		{
			name: "Failed - Invalid Id",
			userId: "abcd",
			expectedBody: echo.Map{
				"message": "id must be number",
			},
			expectedCode: http.StatusBadRequest,
			wantErr:      true,
		},
		{
			name: "Failed - User Not Found",
			userId: "2",
			expectedBody: echo.Map{
				"message": "user not found",
			},
			expectedCode: http.StatusNotFound,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(echo.DELETE, "/users", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.userId)
			
			if !tt.wantErr {
				mockUserRepo.On("Delete", tt.userId).Return(nil)
			} else {
				mockUserRepo.On("Delete", tt.userId).Return(echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string)))
			}

			err := DeleteUser(c)

			if tt.wantErr {
				assert.Error(t, err)
				httpErr, ok := err.(*echo.HTTPError)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCode, httpErr.Code)
				assert.Equal(t, tt.expectedBody["message"], httpErr.Message)
			} else {
				assert.NoError(t, err)

				var ret struct {
					Message string         `json:"message"`
				}

				err = json.Unmarshal(rec.Body.Bytes(), &ret)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.Equal(t, tt.expectedBody["message"], ret.Message)
			}
		})
	}
}