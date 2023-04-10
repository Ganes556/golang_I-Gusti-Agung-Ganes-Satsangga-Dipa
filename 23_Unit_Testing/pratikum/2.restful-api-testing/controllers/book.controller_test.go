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

func TestGetBooks(t *testing.T) {
	mockBookRepo := services.MockBookRepo{}
	services.SetBookRepo(&mockBookRepo)

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
				"message": "success to get all books",
				"books": []models.BookRes{
					{
						ID: 1,
						Judul: "Buku test 1",
						Penulis: "Penulis test 1",
						Penerbit: "Penerbit test 1",
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
			req := httptest.NewRequest(echo.GET, "/books", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if !tt.wantErr {
				mockBookRepo.On("FindAll").Return(tt.expectedBody["books"], nil).Once()
			} else {
				mockBookRepo.On("FindAll").Return(tt.expectedCode, echo.NewHTTPError(http.StatusInternalServerError, tt.expectedBody["message"])).Once()
			}

			err := GetBooks(c)
			if tt.wantErr {
				assert.Error(t, err)
				httpErr, ok := err.(*echo.HTTPError)
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCode, httpErr.Code)
				assert.Equal(t, tt.expectedBody["message"], httpErr.Message)
			} else {

				assert.NoError(t, err)
				var ret struct {
					Message string           `json:"message"`
					Books   []models.BookRes `json:"books"`
				}

				err = json.Unmarshal(rec.Body.Bytes(), &ret)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.Equal(t, tt.expectedBody["message"], ret.Message)
				assert.Equal(t, tt.expectedBody["books"], ret.Books)
			}
		})
	}
}


func TestGetBook(t *testing.T) {
	mockBookRepo := services.MockBookRepo{}
	services.SetBookRepo(&mockBookRepo)
	
	tests := []struct {
		name         string
		bookId 			 string
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			bookId: "1",
			expectedBody: echo.Map{
				"message": "success get book by id 1",
				"book": models.BookRes{
					ID:    1,
					Judul: "Buku test 1",
					Penulis: "Penulis test 1",
					Penerbit: "Penerbit test 1",
				},
			},
			expectedCode: http.StatusOK,
			wantErr:      false,
		},
		{
			name: "Failed - Invalid Id",
			bookId: "abcd",
			expectedBody: echo.Map{
				"message": "id must be number",
			},
			expectedCode: http.StatusBadRequest,
			wantErr:      true,
		},
		{
			name: "Failed - Book Not Found",
			bookId: "2",
			expectedBody: echo.Map{
				"message": "book not found",
			},
			expectedCode: http.StatusNotFound,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(echo.GET, "/books", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.bookId)
			
			if !tt.wantErr {
				mockBookRepo.On("FindById", tt.bookId).Return(tt.expectedBody["book"], nil).Once()
			} else {
				mockBookRepo.On("FindById", tt.bookId).Return(models.BookRes{}, echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string))).Once()
			}

			err := GetBook(c)

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
					Book    models.BookRes `json:"book"`
				}

				err = json.Unmarshal(rec.Body.Bytes(), &ret)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.Equal(t, tt.expectedBody["message"], ret.Message)
				assert.Equal(t, tt.expectedBody["book"], ret.Book)
			}
		})
	}
}

func TestCreateBook(t *testing.T) {
	mockBookRepo := services.MockBookRepo{}
	services.SetBookRepo(&mockBookRepo)

	tests := []struct {
		name         string
		payload      models.Book
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		{
			name: "Success",
			payload: models.Book{
				Judul: "Buku test 1",
				Penulis: "Penulis test 1",
				Penerbit: "Penerbit test 1",
			},
			expectedBody: echo.Map{
				"message": "success create new book",
			},
			expectedCode: http.StatusCreated,
			wantErr: false,
		},
		{
			name: "Failed - Required Judul",
			payload: models.Book{
				Penulis: "Penulis test 1",
				Penerbit: "Penerbit test 1",
			},
			expectedBody: echo.Map{
				"message": "judul is required",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - Required Penulis",
			payload: models.Book{
				Judul: "Buku test 1",
				Penerbit: "Penerbit test 1",
			},
			expectedBody: echo.Map{
				"message": "penulis is required",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - Required Penerbit",
			payload: models.Book{
				Judul: "Buku test 1",
				Penulis: "Penulis test 1",
			},
			expectedBody: echo.Map{
				"message": "penerbit is required",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - Book Conflict",
			payload: models.Book{
				Judul: "Buku test 1",
				Penulis: "Penulis test 1",
				Penerbit: "Penerbit test 1",
			},
			expectedBody: echo.Map{
				"message": "book already exist",
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
		
			req := httptest.NewRequest(echo.POST, "/books", bytes.NewReader(payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			
			if !tt.wantErr {
				mockBookRepo.On("Create", tt.payload).Return(nil).Once()
			}else {				
				mockBookRepo.On("Create", tt.payload).Return(echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string))).Once()
			}

			err = CreateBook(c)
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

func TestDeleteBook(t *testing.T) {
	mockBookRepo := services.MockBookRepo{}
	services.SetBookRepo(&mockBookRepo)
	
	tests := []struct {
		name         string
		bookId 			 string
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			bookId: "1",
			expectedBody: echo.Map{
				"message": "success delete book by id 1",
			},
			expectedCode: http.StatusOK,
			wantErr:      false,
		},
		{
			name: "Failed - Invalid Id",
			bookId: "abcd",
			expectedBody: echo.Map{
				"message": "id must be number",
			},
			expectedCode: http.StatusBadRequest,
			wantErr:      true,
		},
		{
			name: "Failed - Book Not Found",
			bookId: "2",
			expectedBody: echo.Map{
				"message": "book not found",
			},
			expectedCode: http.StatusNotFound,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(echo.DELETE, "/Books", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.bookId)
			
			if !tt.wantErr {
				mockBookRepo.On("Delete", tt.bookId).Return(nil).Once()
			} else {
				mockBookRepo.On("Delete", tt.bookId).Return(echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string))).Once()
			}

			err := DeleteBook(c)

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

func TestUpdateBook(t *testing.T) {
	mockBookRepo := services.MockBookRepo{}
	services.SetBookRepo(&mockBookRepo)

	tests := []struct {
		name         string
		bookId	     string
		payload      models.BookReqUpdate
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		{
			name: "Success",
			bookId: "1",
			payload: models.BookReqUpdate{
				Judul: "test2",
				Penulis: "penulis test2",
				Penerbit: "penerbit test2",
			},
			expectedBody: echo.Map{
				"message": "success update book by id 1",
			},
			expectedCode: http.StatusOK,
			wantErr: false,
		},
		
		{
			name: "Failed - Invalid Id",
			bookId: "abcd",
			payload: models.BookReqUpdate{
				Judul: "test2",
				Penulis: "penulis test2",
				Penerbit: "penerbit test2",
			},
			expectedBody: echo.Map{
				"message": "id must be a number",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - Book Not Found",
			bookId: "999",
			payload: models.BookReqUpdate{
				Judul: "test2",
				Penulis: "penulis test2",
				Penerbit: "penerbit test2",
			},
			expectedBody: echo.Map{
				"message": "book not found",
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
		
			req := httptest.NewRequest(echo.PUT, "/books", bytes.NewReader(payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.bookId)
			
			if !tt.wantErr {
				mockBookRepo.On("Update",tt.bookId ,tt.payload).Return(nil).Once()
			}else {
				mockBookRepo.On("Update",tt.bookId , tt.payload).Return(echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string))).Once()
			}

			err = UpdateBook(c)

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