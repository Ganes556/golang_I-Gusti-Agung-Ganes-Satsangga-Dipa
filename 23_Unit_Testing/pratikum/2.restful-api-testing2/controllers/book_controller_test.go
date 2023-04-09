package controllers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
// 	lib "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/libraries"
// 	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

// func init() {
// 	configs.InitDB()
// }

// func InsertDataBookForGetbook(book models.Book) error {
// 	if err := configs.DB.Create(&book).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func TestGetbooks(t *testing.T) {
// 	var book = models.Book{
// 		Base: models.Base{
// 			ID: 1,
// 		},
// 		Judul:    "Testing Judul",
// 		Penulis:  "Testing Penulis",
// 		Penerbit: "Testing Penerbit",
// 	}

// 	var testCase = []struct {
// 		name         string
// 		path         string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "[OK] Get books",
// 			path:         "/books",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	e := echo.New()
// 	req := httptest.NewRequest(echo.GET, "/", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// add books

// 	InsertDataBookForGetbook(book)

// 	for _, tc := range testCase {
// 		c.SetPath(tc.path)
// 		if assert.NoError(t, GetBooks(c)) {
// 			assert.Equal(t, tc.expectedCode, rec.Code)
// 			body := rec.Body.String()

// 			var resbooks struct {
// 				Message string        `json:"message"`
// 				books   []models.book `json:"books"`
// 			}

// 			err := json.Unmarshal([]byte(body), &resbooks)
// 			if err != nil {
// 				assert.Error(t, err, "Error Json Unmarshal book")
// 			}
// 			assert.Equal(t, "success get all books", resbooks.Message)
// 			assert.Condition(t, func() bool {
// 				return len(resbooks.books) > 0
// 			})
// 		}
// 	}
// }

// func TestGetbook(t *testing.T) {
// 	e := echo.New()

// 	testCases := []struct {
// 		name                string
// 		bookID              string
// 		expectedStatus      int
// 		expectedBodybook    echo.Map
// 		expectedBodyMessage string
// 	}{
// 		{
// 			name:           "[OK] Get book",
// 			bookID:         "1",
// 			expectedStatus: http.StatusOK,
// 			expectedBodybook: echo.Map{
// 				"id":       float64(1),
// 				"name":     "Testing Name",
// 				"email":    "testtingbook@gmail.com",
// 				"password": "testingpassword",
// 			},
// 			expectedBodyMessage: "success get book by id 1",
// 		},
// 		{
// 			name:                "[Not Found] Get book",
// 			bookID:              "999",
// 			expectedStatus:      http.StatusNotFound,
// 			expectedBodyMessage: "book not found",
// 		},
// 		{
// 			name:                "[Bad Request] Get book - id not number",
// 			bookID:              "abcd",
// 			expectedStatus:      http.StatusBadRequest,
// 			expectedBodyMessage: "id must be number",
// 		},
// 	}

// 	// Loop over the test cases
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Create a new request with the test book ID
// 			req := httptest.NewRequest(echo.GET, "/", nil)
// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/books/:id")
// 			c.SetParamNames("id")
// 			c.SetParamValues(tc.bookID)
// 			// Call the Getbook endpoint
// 			err := Getbook(c)

// 			body := rec.Body.String()

// 			var resbook struct {
// 				Message string   `json:"message"`
// 				book    echo.Map `json:"book"`
// 			}

// 			if err := json.Unmarshal([]byte(body), &resbook); err != nil {
// 				assert.Error(t, err, "Error Json Unmarshal book")
// 			}

// 			if err != nil {
// 				he, ok := err.(*echo.HTTPError)
// 				if ok {
// 					assert.Equal(t, tc.expectedStatus, he.Code)
// 					assert.Equal(t, tc.expectedBodyMessage, he.Message)
// 				}

// 			} else {
// 				assert.Equal(t, tc.expectedBodyMessage, resbook.Message)
// 				assert.Equal(t, tc.expectedBodybook, resbook.book)
// 				assert.Equal(t, tc.expectedStatus, rec.Code)
// 			}

// 		})
// 	}
// }

// func TestCreatebook(t *testing.T) {
// 	e := echo.New()
// 	e.Validator = lib.NewValidator()
// 	configs.DB.Unscoped().Delete(&models.book{}, "email = ?", "testingcase1234@gmail.com")
// 	testCases := []struct {
// 		name                string
// 		book                echo.Map
// 		expectedStatus      int
// 		expectedBodyMessage string
// 	}{
// 		{
// 			name: "[OK] Create book",
// 			book: echo.Map{
// 				"name":     "Testing Name",
// 				"email":    "testingcase1234@gmail.com",
// 				"password": "testingpassword123",
// 			},
// 			expectedStatus:      http.StatusCreated,
// 			expectedBodyMessage: "success create new book",
// 		},
// 		{
// 			name: "[Bad Request] Create book - invalid email",
// 			book: echo.Map{
// 				"name":     "Testing Name",
// 				"email":    "testingcase",
// 				"password": "testingpasswo",
// 			},
// 			expectedStatus:      http.StatusBadRequest,
// 			expectedBodyMessage: "email is not valid",
// 		},
// 		{
// 			name: "[Bad Request] Create book - invalid password",
// 			book: echo.Map{
// 				"name":     "Testing Name",
// 				"email":    "testingcase123@gmail.com",
// 				"password": "test",
// 			},
// 			expectedStatus:      http.StatusBadRequest,
// 			expectedBodyMessage: "password must be at least 8 characters",
// 		},
// 		{
// 			name: "[Bad Request] Create book - email already exist",
// 			book: echo.Map{
// 				"name":     "Testing Name",
// 				"email":    "testingcase1234@gmail.com",
// 				"password": "testingpassword123",
// 			},
// 			expectedStatus:      http.StatusBadRequest,
// 			expectedBodyMessage: "email already exist",
// 		},
// 	}

// 	// Loop over the test cases
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Create a new request with the test book ID
// 			bodyReq, _ := json.Marshal(tc.book)
// 			// fmt.Println(string(bodyReq))
// 			req := httptest.NewRequest(echo.POST, "/books", bytes.NewReader(bodyReq))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)

// 			err := Createbook(c)

// 			bodyRes := rec.Body.String()

// 			var res struct {
// 				Message string   `json:"message"`
// 				book    echo.Map `json:"book"`
// 			}

// 			if err := json.Unmarshal([]byte(bodyRes), &res); err != nil {
// 				assert.Error(t, err, "Error Json Unmarshal book")
// 			}
// 			if err != nil {
// 				he, ok := err.(*echo.HTTPError)
// 				if ok {
// 					assert.Equal(t, tc.expectedStatus, he.Code)
// 					assert.Equal(t, tc.expectedBodyMessage, he.Message)
// 				}

// 			} else {
// 				assert.Equal(t, tc.expectedStatus, rec.Code)
// 				assert.Equal(t, tc.expectedBodyMessage, res.Message)
// 				delete(res.book, "id")
// 				assert.Equal(t, tc.book, res.book)
// 			}

// 		})
// 	}

// }

// func TestUpdatebook(t *testing.T) {
// 	e := echo.New()
// 	e.Validator = lib.NewValidator()

// 	testCases := []struct {
// 		name                string
// 		bookID              string
// 		updateData          echo.Map
// 		expectedStatus      int
// 		expectedBodyMessage string
// 	}{
// 		{
// 			name:   "[OK] Update book",
// 			bookID: "1",
// 			updateData: echo.Map{
// 				"name": "Testing Name",
// 			},
// 			expectedStatus:      http.StatusOK,
// 			expectedBodyMessage: "success update book by id 1",
// 		},
// 		{
// 			name:   "[Bad Request] Update book - invalid id",
// 			bookID: "abcd",
// 			updateData: echo.Map{
// 				"email": "testing999@gmail.com",
// 			},
// 			expectedStatus:      http.StatusBadRequest,
// 			expectedBodyMessage: "id must be number",
// 		},
// 		{
// 			name:   "[Bad Request] Update book - invalid email",
// 			bookID: "1",
// 			updateData: echo.Map{
// 				"email": "testing999gmail.com",
// 			},
// 			expectedStatus:      http.StatusBadRequest,
// 			expectedBodyMessage: "email is not valid",
// 		},
// 		{
// 			name:   "[Bad Request] Update book - invalid password",
// 			bookID: "1",
// 			updateData: echo.Map{
// 				"password": "test",
// 			},
// 			expectedStatus:      http.StatusBadRequest,
// 			expectedBodyMessage: "password must be at least 8 characters",
// 		},
// 		{
// 			name:   "[Not Found] Update book",
// 			bookID: "999",
// 			updateData: echo.Map{
// 				"name": "Testing Name",
// 			},
// 			expectedStatus:      http.StatusNotFound,
// 			expectedBodyMessage: "book not found",
// 		},
// 	}

// 	// Loop over the test cases
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Create a new request with the test book ID
// 			bodyReq, _ := json.Marshal(tc.updateData)
// 			req := httptest.NewRequest(echo.PUT, "/", bytes.NewReader(bodyReq))
// 			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/books/:id")
// 			c.SetParamNames("id")
// 			c.SetParamValues(tc.bookID)

// 			// Call the Updatebook endpoint
// 			err := Updatebook(c)

// 			body := rec.Body.String()

// 			var res struct {
// 				Message string `json:"message"`
// 			}

// 			if err := json.Unmarshal([]byte(body), &res); err != nil {
// 				assert.Error(t, err, "Error Json Unmarshal book")
// 			}
// 			if err != nil {
// 				he, ok := err.(*echo.HTTPError)
// 				if ok {
// 					assert.Equal(t, tc.expectedStatus, he.Code)
// 					assert.Equal(t, tc.expectedBodyMessage, he.Message)
// 				}

// 			} else {
// 				assert.Equal(t, tc.expectedStatus, rec.Code)
// 				assert.Equal(t, tc.expectedBodyMessage, res.Message)
// 			}
// 		})
// 	}
// }

// func TestDeletebook(t *testing.T) {
// 	e := echo.New()
// 	e.Validator = lib.NewValidator()

// 	testCases := []struct {
// 		name                string
// 		bookID              string
// 		expectedStatus      int
// 		expectedBodyMessage string
// 	}{
// 		{
// 			name:                "[OK] Delete book",
// 			bookID:              "1",
// 			expectedStatus:      http.StatusOK,
// 			expectedBodyMessage: "success delete book by id 1",
// 		},
// 		{
// 			name:                "[Bad Request] Delete book - invalid id",
// 			bookID:              "abcd",
// 			expectedStatus:      http.StatusBadRequest,
// 			expectedBodyMessage: "id must be number",
// 		},
// 		{
// 			name:                "[Not Found] Delete book",
// 			bookID:              "999",
// 			expectedStatus:      http.StatusNotFound,
// 			expectedBodyMessage: "book not found",
// 		},
// 	}

// 	// Loop over the test cases
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			// Create a new request with the test book ID
// 			req := httptest.NewRequest(echo.DELETE, "/", nil)
// 			rec := httptest.NewRecorder()
// 			c := e.NewContext(req, rec)
// 			c.SetPath("/books/:id")
// 			c.SetParamNames("id")
// 			c.SetParamValues(tc.bookID)

// 			// Call the Deletebook endpoint
// 			err := Deletebook(c)

// 			body := rec.Body.String()

// 			var res struct {
// 				Message string `json:"message"`
// 			}

// 			if err := json.Unmarshal([]byte(body), &res); err != nil {
// 				assert.Error(t, err, "Error Json Unmarshal book")
// 			}
// 			if err != nil {
// 				he, ok := err.(*echo.HTTPError)
// 				if ok {
// 					assert.Equal(t, tc.expectedStatus, he.Code)
// 					assert.Equal(t, tc.expectedBodyMessage, he.Message)
// 				}

// 			} else {
// 				assert.Equal(t, tc.expectedStatus, rec.Code)
// 				assert.Equal(t, tc.expectedBodyMessage, res.Message)
// 			}
// 		})
// 	}
// }
