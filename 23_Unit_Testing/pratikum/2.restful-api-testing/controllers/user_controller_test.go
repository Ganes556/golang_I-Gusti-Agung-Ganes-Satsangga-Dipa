package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
	lib "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/libraries"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init(){	
	configs.InitDB()
}



func InsertDataUserForGetUser(user models.User) error{
	if err := configs.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetUsers(t *testing.T) {
	var user = models.User{
		Base: models.Base{
			ID: 1,
		},
		Name: "Testing Name",
		Email: "testtinguser@gmail.com",
		Password: "testingpassword",
	}

	var testCase = []struct {
		name string
		path string
		expectedCode int
	}{
		{
			name: "[OK] Get Users",
			expectedCode: http.StatusOK,
		},
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// add users
	
	InsertDataUserForGetUser(user)	

	for _, tc := range testCase {
		c.SetPath("/users")
		if assert.NoError(t, GetUsers(c)) {
			assert.Equal(t, tc.expectedCode, rec.Code)
			body := rec.Body.String()

			var resUsers struct{
				Message string `json:"message"`
				Users []models.User `json:"users"`
			}

			err := json.Unmarshal([]byte(body), &resUsers)
			if err != nil {
				assert.Error(t, err, "Error Json Unmarshal User")
			}
			assert.Equal(t, "success get all users", resUsers.Message)
			assert.Condition(t, func() bool {
				return len(resUsers.Users) > 0
			})
		}
	}
}

func TestGetUser(t *testing.T) {
	e := echo.New()

	testCases := []struct {
		name string
		userID string
		expectedStatus int
		expectedBodyUser echo.Map
		expectedBodyMessage string
	}{
		{
			name: "[OK] Get User",
			userID: "1",
			expectedStatus: http.StatusOK,
			expectedBodyUser: echo.Map{
				"id": float64(1),
				"name": "Testing Name",
				"email": "testtinguser@gmail.com",
				"password": "testingpassword",
			},
			expectedBodyMessage: "success get user by id 1",
		},
		{
			name: "[Not Found] Get User",
			userID: "999",
			expectedStatus: http.StatusNotFound,
			expectedBodyMessage: "user not found",
		},
		{
			name: "[Bad Request] Get User - id not number",
			userID: "abcd",
			expectedStatus: http.StatusBadRequest,
			expectedBodyMessage: "id must be number",
		},
	}

	// Loop over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new request with the test user ID
			req := httptest.NewRequest(echo.GET, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.userID)
			// Call the GetUser endpoint
			err := GetUser(c)

			body := rec.Body.String()
			
			var resUser struct{
				Message string `json:"message"`
				User echo.Map `json:"user"`
			}
			
			if err := json.Unmarshal([]byte(body), &resUser); err != nil {
				assert.Error(t, err, "Error Json Unmarshal User")
			}
			
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
          assert.Equal(t,tc.expectedStatus, he.Code)
					assert.Equal(t, tc.expectedBodyMessage, he.Message)
        }

			}else {
				assert.Equal(t, tc.expectedBodyMessage, resUser.Message)
				assert.Equal(t, tc.expectedBodyUser, resUser.User)
				assert.Equal(t, tc.expectedStatus, rec.Code)
			}
			
		})
	}
}

func TestCreateUser(t *testing.T) {
	e := echo.New()
	e.Validator = lib.NewValidator()
	configs.DB.Unscoped().Delete(&models.User{}, "email = ?", "testingcase1234@gmail.com")
	testCases := []struct {
		name string
		user echo.Map
		expectedStatus int
		expectedBodyMessage string
	}{
		{
			name: "[OK] Create User",
			user: echo.Map{
				"name": "Testing Name",
				"email": "testingcase1234@gmail.com",
				"password": "testingpassword123",
			},
			expectedStatus: http.StatusCreated,
			expectedBodyMessage: "success create new user",
		},
		{
			name: "[Bad Request] Create User - invalid email",
			user: echo.Map{
				"name": "Testing Name",
				"email": "testingcase",
				"password": "testingpasswo",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBodyMessage: "email is not valid",
		},
		{
			name: "[Bad Request] Create User - invalid password",
			user: echo.Map{
				"name": "Testing Name",
				"email": "testingcase123@gmail.com",
				"password": "test",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBodyMessage: "password must be at least 8 characters",
		},
		{
			name: "[Bad Request] Create User - email already exist",
			user: echo.Map{
				"name": "Testing Name",
				"email": "testingcase1234@gmail.com",
				"password": "testingpassword123",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBodyMessage: "email already exist",
		},
	}

	// Loop over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new request with the test user ID
			bodyReq, _ := json.Marshal(tc.user)
			// fmt.Println(string(bodyReq))
			req := httptest.NewRequest(echo.POST, "/users", bytes.NewReader(bodyReq))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := CreateUser(c)

			bodyRes := rec.Body.String()
			
			var res struct{
				Message string `json:"message"`
				User echo.Map `json:"user"`
			}

			if err := json.Unmarshal([]byte(bodyRes), &res); err != nil {
				assert.Error(t, err, "Error Json Unmarshal User")
			}
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t,tc.expectedStatus, he.Code)
					assert.Equal(t, tc.expectedBodyMessage, he.Message)
				}

			}else {
				assert.Equal(t, tc.expectedStatus, rec.Code)
				assert.Equal(t, tc.expectedBodyMessage, res.Message)
				delete(res.User, "id")
				assert.Equal(t, tc.user, res.User)
			}
			
		})
	}

}

func TestUpdateUser(t *testing.T) {
	e := echo.New()
	e.Validator = lib.NewValidator()

	testCases := []struct {
		name                string
		userID              string
		updateData          echo.Map
		expectedStatus      int
		expectedBodyMessage string
	}{
		{
			name:   "[OK] Update User",
			userID: "1",
			updateData: echo.Map{
				"name": "Testing Name",
			},
			expectedStatus:      http.StatusOK,
			expectedBodyMessage: "success update user by id 1",
		},
		{
			name:   "[Bad Request] Update User - invalid id",
			userID: "abcd",
			updateData: echo.Map{
				"email": "testing999@gmail.com",
			},
			expectedStatus:      http.StatusBadRequest,
			expectedBodyMessage: "id must be number",
		},
		{
			name:   "[Bad Request] Update User - invalid email",
			userID: "1",
			updateData: echo.Map{
				"email": "testing999gmail.com",
			},
			expectedStatus:      http.StatusBadRequest,
			expectedBodyMessage: "email is not valid",
		},
		{
			name:   "[Bad Request] Update User - invalid password",
			userID: "1",
			updateData: echo.Map{
				"password": "test",
			},
			expectedStatus:      http.StatusBadRequest,
			expectedBodyMessage: "password must be at least 8 characters",
		},
		{
			name:                "[Not Found] Update User",
			userID:              "999",
			updateData: echo.Map{
				"name": "Testing Name",
			},
			expectedStatus:      http.StatusNotFound,
			expectedBodyMessage: "user not found",
		},
	}

	// Loop over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new request with the test user ID
			bodyReq, _ := json.Marshal(tc.updateData)
			req := httptest.NewRequest(echo.PUT, "/", bytes.NewReader(bodyReq))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.userID)
			
			
			// Call the UpdateUser endpoint
			err := UpdateUser(c)

			body := rec.Body.String()

			var res struct {
				Message string `json:"message"`
			}

			if err := json.Unmarshal([]byte(body), &res); err != nil {
				assert.Error(t, err, "Error Json Unmarshal User")
			}
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, tc.expectedStatus, he.Code)
					assert.Equal(t, tc.expectedBodyMessage, he.Message)
				}

			}else {
				assert.Equal(t, tc.expectedStatus, rec.Code)
				assert.Equal(t, tc.expectedBodyMessage, res.Message)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	e := echo.New()
	e.Validator = lib.NewValidator()

	testCases := []struct {
		name                string
		userID              string
		expectedStatus      int
		expectedBodyMessage string
	}{
		{
			name:                "[OK] Delete User",
			userID:              "1",
			expectedStatus:      http.StatusOK,
			expectedBodyMessage: "success delete user by id 1",
		},
		{
			name:                "[Bad Request] Delete User - invalid id",
			userID:              "abcd",
			expectedStatus:      http.StatusBadRequest,
			expectedBodyMessage: "id must be number",
		},
		{
			name:                "[Not Found] Delete User",
			userID:              "999",
			expectedStatus:      http.StatusNotFound,
			expectedBodyMessage: "user not found",
		},
	}

	// Loop over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new request with the test user ID
			req := httptest.NewRequest(echo.DELETE, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(tc.userID)

			// Call the DeleteUser endpoint
			err := DeleteUser(c)

			body := rec.Body.String()

			var res struct {
				Message string `json:"message"`
			}

			if err := json.Unmarshal([]byte(body), &res); err != nil {
				assert.Error(t, err, "Error Json Unmarshal User")
			}
			if err != nil {
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, tc.expectedStatus, he.Code)
					assert.Equal(t, tc.expectedBodyMessage, he.Message)
				}

			}else {
				assert.Equal(t, tc.expectedStatus, rec.Code)
				assert.Equal(t, tc.expectedBodyMessage, res.Message)
			}
		})
	}
}