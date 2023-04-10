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

func TestGetBlogByUserId(t *testing.T) {
	mockBlogRepo := services.MockBlogRepo{}
	services.SetBlogRepo(&mockBlogRepo)

	tests := []struct {
		name         string
		userId       string
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			userId: "1",
			expectedBody: echo.Map{
				"message": "success get blogs by user id 1",
				"blogs": []models.BlogByUserId{
					{
						Penulis: "Penulis 12",
						Judul: "Judul 1",
						Konten: "Konten 1",
					},
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
			req := httptest.NewRequest(echo.GET, "/blogs/userId", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.userId)

			if !tt.wantErr {
				mockBlogRepo.On("FindBlogByUserId",tt.userId).Return(tt.expectedBody["blogs"], nil).Once()
			} else {
				mockBlogRepo.On("FindBlogByUserId",tt.userId).Return([]models.BlogByUserId{}, echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string))).Once()
			}

			err := GetBlogByUserId(c)
			
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
					Blogs   []models.BlogByUserId `json:"blogs"`
				}

				err = json.Unmarshal(rec.Body.Bytes(), &ret)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.Equal(t, tt.expectedBody["message"], ret.Message)
				assert.Equal(t, tt.expectedBody["blogs"], ret.Blogs)
			}
		})
	}
}


func TestGetBlogs(t *testing.T) {
	mockBlogRepo := services.MockBlogRepo{}
	services.SetBlogRepo(&mockBlogRepo)

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
				"message": "success to get all blogs",
				"blogs": []models.BlogRes{
					{
						ID:       1,
						UserRefer: 1,
						Judul: 	"Judul 1",
						Konten: "Konten 1",
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
			req := httptest.NewRequest(echo.GET, "/blogs", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if !tt.wantErr {
				mockBlogRepo.On("FindAll").Return(tt.expectedBody["blogs"], nil).Once()
			} else {
				mockBlogRepo.On("FindAll").Return([]models.BlogRes{}, echo.NewHTTPError(http.StatusInternalServerError, tt.expectedBody["message"].(string))).Once()
			}

			err := GetBlogs(c)
			
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
					Blogs   []models.BlogRes `json:"blogs"`
				}

				err = json.Unmarshal(rec.Body.Bytes(), &ret)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.Equal(t, tt.expectedBody["message"], ret.Message)
				assert.Equal(t, tt.expectedBody["blogs"], ret.Blogs)
			}
		})
	}
}



func TestGetBlog(t *testing.T) {
	mockBlogRepo := services.MockBlogRepo{}
	services.SetBlogRepo(&mockBlogRepo)
	
	tests := []struct {
		name         string
		blogId 			 string
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			blogId: "1",
			expectedBody: echo.Map{
				"message": "success get blog by id 1",
				"blog": models.BlogRes{
					ID:       1,
					UserRefer: 1,
					Judul: 	"Judul 1",
					Konten: "Konten 1",
				},
			},
			expectedCode: http.StatusOK,
			wantErr:      false,
		},
		{
			name: "Failed - Invalid Id",
			blogId: "abcd",
			expectedBody: echo.Map{
				"message": "id must be number",
			},
			expectedCode: http.StatusBadRequest,
			wantErr:      true,
		},
		{
			name: "Failed - Blog Not Found",
			blogId: "2",
			expectedBody: echo.Map{
				"message": "blog not found",
			},
			expectedCode: http.StatusNotFound,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(echo.GET, "/blogs", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.blogId)
			
			if !tt.wantErr {
				mockBlogRepo.On("FindById", tt.blogId).Return(tt.expectedBody["blog"], nil).Once()
			} else {
				mockBlogRepo.On("FindById", tt.blogId).Return(models.BlogRes{}, echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string))).Once()
			}

			err := GetBlog(c)

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
					Blog    models.BlogRes `json:"blog"`
				}

				err = json.Unmarshal(rec.Body.Bytes(), &ret)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCode, rec.Code)
				assert.Equal(t, tt.expectedBody["message"], ret.Message)
				assert.Equal(t, tt.expectedBody["blog"], ret.Blog)
			}
		})
	}
}

func TestCreateBlog(t *testing.T) {
	mockBlogRepo := services.MockBlogRepo{}
	services.SetBlogRepo(&mockBlogRepo)

	tests := []struct {
		name         string
		payload      models.Blog
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		{
			name: "Success",
			payload: models.Blog{
				UserRefer: 1,
				Judul: 	"Judul 2",
				Konten: "Konten 2",
			},
			expectedBody: echo.Map{
				"message": "success create new blog",
			},
			expectedCode: http.StatusCreated,
			wantErr: false,
		},
		{
			name: "Failed - Required User Refer",
			payload: models.Blog{
				Judul: 	"Judul 1",
				Konten: "Konten 1",
			},
			expectedBody: echo.Map{
				"message": "userrefer is required",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - Required Judul",
			payload: models.Blog{
				UserRefer: 1,
				Konten: "Konten 1",
			},
			expectedBody: echo.Map{
				"message": "judul is required",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - Required Konten",
			payload: models.Blog{
				UserRefer: 1,
				Judul: 	"Judul 2",
			},
			expectedBody: echo.Map{
				"message": "konten is required",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - Blog Conflict",
			payload: models.Blog{
				UserRefer: 1,
				Judul: 	"Judul 2",
				Konten: "Konten 2",
			},
			expectedBody: echo.Map{
				"message": "blog already exist",
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
		
			req := httptest.NewRequest(echo.POST, "/blogs", bytes.NewReader(payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			
			if !tt.wantErr {
				mockBlogRepo.On("Create", tt.payload).Return(nil).Once()
			}else {				
				mockBlogRepo.On("Create", tt.payload).Return(echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string))).Once()
			}

			err = CreateBlog(c)
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

func TestDeleteBlog(t *testing.T) {
	mockBlogRepo := services.MockBlogRepo{}
	services.SetBlogRepo(&mockBlogRepo)
	
	tests := []struct {
		name         string
		blogId 			 string
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			blogId: "1",
			expectedBody: echo.Map{
				"message": "success delete blog by id 1",
			},
			expectedCode: http.StatusOK,
			wantErr:      false,
		},
		{
			name: "Failed - Invalid Id",
			blogId: "abcd",
			expectedBody: echo.Map{
				"message": "id must be number",
			},
			expectedCode: http.StatusBadRequest,
			wantErr:      true,
		},
		{
			name: "Failed - Blog Not Found",
			blogId: "2",
			expectedBody: echo.Map{
				"message": "blog not found",
			},
			expectedCode: http.StatusNotFound,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(echo.DELETE, "/Blogs", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.blogId)
			
			if !tt.wantErr {
				mockBlogRepo.On("Delete", tt.blogId).Return(nil).Once()
			} else {
				mockBlogRepo.On("Delete", tt.blogId).Return(echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string))).Once()
			}

			err := DeleteBlog(c)

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

func TestUpdateBlog(t *testing.T) {
	mockBlogRepo := services.MockBlogRepo{}
	services.SetBlogRepo(&mockBlogRepo)
	
	tests := []struct {
		name         string
		blogId	     string
		payload      models.BlogReqUpdate
		expectedBody echo.Map
		expectedCode int
		wantErr      bool
	}{
		{
			name: "Success",
			blogId: "1",
			payload: models.BlogReqUpdate{
				UserRefer: 1,
				Judul: 	"Judul 3",
				Konten: "Konten 3",
			},
			expectedBody: echo.Map{
				"message": "success update blog by id 1",
			},
			expectedCode: http.StatusOK,
			wantErr: false,
		},
		
		{
			name: "Failed - Invalid Id",
			blogId: "abcd",
			payload: models.BlogReqUpdate{
				UserRefer: 1,
				Judul: 	"Judul 3",
				Konten: "Konten 3",
			},
			expectedBody: echo.Map{
				"message": "id must be a number",
			},
			expectedCode: http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Failed - Blog Not Found",
			blogId: "999",
			payload: models.BlogReqUpdate{
				UserRefer: 1,
				Judul: 	"Judul 3",
				Konten: "Konten 3",
			},
			expectedBody: echo.Map{
				"message": "blog not found",
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
		
			req := httptest.NewRequest(echo.PUT, "/blogs", bytes.NewReader(payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.blogId)
			
			if !tt.wantErr {
				mockBlogRepo.On("Update",tt.blogId ,tt.payload).Return(nil).Once()
			}else {
				mockBlogRepo.On("Update",tt.blogId , tt.payload).Return(echo.NewHTTPError(tt.expectedCode, tt.expectedBody["message"].(string))).Once()
			}

			err = UpdateBlog(c)

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

