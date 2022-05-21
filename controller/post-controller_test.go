package controller

import (
	// "InstagramMock-v2/service"
	"InstagramMock-v2/model"
	"InstagramMock-v2/service"
	"InstagramMock-v2/service/mocks"
	"bytes"
	"errors"

	// "encoding/json"
	// "go-test/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	// mockService    = mocks.IPostService{mock.Mock{}}
	// postController = NewPostController(&mockService)
	// testRouter     = InitRoute()
	timeNow = time.Now()
	posts   = []model.Post{
		model.Post{
			ID:          1,
			UrlPhoto:    "abc.jpg",
			Description: "abc",
			Likes:       0,
			CreatedAt:   timeNow,
			Username:    "abc",
		},
		model.Post{
			ID:          2,
			UrlPhoto:    "def.jpg",
			Description: "def",
			Likes:       0,
			CreatedAt:   timeNow,
			Username:    "def",
		},
		model.Post{
			ID:          1,
			UrlPhoto:    "ghi.jpg",
			Description: "ghi",
			Likes:       0,
			CreatedAt:   timeNow,
			Username:    "ghi",
		},
	}
)

func InitRoute(postController IPostController) *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.Static("/css", "../templates/css")
	router.Static("/upload", "../upload")
	router.LoadHTMLGlob("../templates/*.html")
	//html

	//read, findbyid, create, update, delete,
	//pageDetail, pageUpdate, pageCreate

	router.GET("/", postController.ShowAll)
	router.GET("/:id", postController.FindByID)
	// router.POST("/create", postController.Create)
	router.GET("/create", postController.PageCreate)
	router.POST("/:id/edit", postController.Update)
	router.POST("/:id", postController.Delete)
	router.GET("/:id/edit", postController.PageUpdate)

	return router
}

func TestShowAllError(t *testing.T) {
	mockService := mocks.IPostService{mock.Mock{}}
	mockService.Mock.On("ShowAll").Return(nil, errors.New("PostService Error"))
	postController := NewPostController(&mockService)
	testRouter := InitRoute(postController)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()
	// gin.CreateTestContext(resp)
	// hf := http.HandlerFunc(postController.ShowAll())
	testRouter.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	mockService.AssertExpectations(t)
}

func TestShowAllSuccess(t *testing.T) {
	mockService := mocks.IPostService{mock.Mock{}}
	mockService.Mock.On("ShowAll").Return(posts, nil)
	postController := NewPostController(&mockService)
	testRouter := InitRoute(postController)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	contentType := resp.Result().Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}

func TestFindByIDSuccess(t *testing.T) {
	mockService := mocks.IPostService{mock.Mock{}}
	mockService.Mock.On("FindByID", uint64(1)).Return(posts[0], nil)
	postController := NewPostController(&mockService)
	testRouter := InitRoute(postController)

	req, err := http.NewRequest("GET", "/1", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	contentType := resp.Result().Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"
	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}

func TestFindByIDErrorParsingID(t *testing.T) {
	mockService := mocks.IPostService{mock.Mock{}}
	postController := NewPostController(&mockService)
	testRouter := InitRoute(postController)

	req, err := http.NewRequest("GET", "/aaaa", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	contentType := resp.Result().Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"
	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}

func TestFindByIDNotFound(t *testing.T) {
	mockService := mocks.IPostService{mock.Mock{}}
	mockService.Mock.On("FindByID", uint64(1)).Return(model.Post{}, errors.New("Not Found"))
	postController := NewPostController(&mockService)
	testRouter := InitRoute(postController)

	req, err := http.NewRequest("GET", "/1", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNotFound, resp.Code)

	contentType := resp.Result().Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"
	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}

func TestPageCreateSuccess(t *testing.T) {
	mockService := mocks.IPostService{mock.Mock{}}
	postController := NewPostController(&mockService)
	testRouter := InitRoute(postController)

	req, err := http.NewRequest("GET", "/create", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	contentType := resp.Result().Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"
	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}

}

func TestUpdateSuccess(t *testing.T) {
	mockService := mocks.IPostService{mock.Mock{}}
	var postDTO service.PostUpdateDTO = service.PostUpdateDTO{
		ID:          uint64(1),
		Description: "123",
	}
	mockService.Mock.On("FindByID", uint64(1)).Return(posts[0], nil)
	mockService.Mock.On("Update", postDTO).Return(model.Post{ID: uint64(1), Description: "123"}, nil)
	postController := NewPostController(&mockService)
	testRouter := InitRoute(postController)

	var payload = []byte(`{
		"description": "123"
	}`)

	req, err := http.NewRequest("POST", "/1/edit", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json;")
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusFound, resp.Code)
}

func TestDeleteSuccess(t *testing.T) {
	mockService := mocks.IPostService{mock.Mock{}}
	mockService.Mock.On("FindByID", uint64(1)).Return(posts[0], nil)
	mockService.Mock.On("Delete", uint64(1)).Return(nil)
	postController := NewPostController(&mockService)
	testRouter := InitRoute(postController)

	req, err := http.NewRequest("POST", "/1", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusFound, resp.Code)
}

func TestPageUpdate(t *testing.T) {
	mockService := mocks.IPostService{mock.Mock{}}
	mockService.Mock.On("FindByID", uint64(1)).Return(posts[0], nil)
	postController := NewPostController(&mockService)
	testRouter := InitRoute(postController)

	req, err := http.NewRequest("GET", "/1/edit", nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	contentType := resp.Result().Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"
	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}

}
