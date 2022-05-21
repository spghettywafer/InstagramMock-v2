package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"

	"InstagramMock-v2/model"
	mocks "InstagramMock-v2/repository/mocks"
)

var (
	timeNow     time.Time             = time.Now()
	mockRepo    mocks.IPostRepository = mocks.IPostRepository{mock.Mock{}}
	postService                       = NewPostService(&mockRepo)
	posts                             = []model.Post{
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

func TestCreate(t *testing.T) {

	var postDTO PostCreateDTO = PostCreateDTO{
		UrlPhoto:    posts[0].UrlPhoto,
		Description: posts[0].Description,
		Username:    posts[0].Username,
	}

	var paramPost model.Post = model.Post{
		UrlPhoto:    postDTO.UrlPhoto,
		Description: postDTO.Description,
		Username:    postDTO.Username,
	}

	mockRepo.Mock.On("Create", paramPost).Return(posts[0], nil)
	result, err := postService.Create(postDTO)
	assert.Nil(t, err)
	assert.NotEmpty(t, result)

	assert.NotNil(t, result.ID)
	assert.Equal(t, posts[0].UrlPhoto, result.UrlPhoto)
	assert.Equal(t, posts[0].Description, result.Description)
	assert.NotNil(t, result.Likes)
	assert.NotNil(t, result.CreatedAt)
	assert.Equal(t, posts[0].Username, result.Username)
}

func TestUpdate(t *testing.T) {
	var postDTO PostUpdateDTO = PostUpdateDTO{
		Description: "123",
	}

	var post model.Post = model.Post{
		Description: "123",
	}

	mockRepo.Mock.On("Update", post).Return(post, nil)
	result, err := postService.Update(postDTO)
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, post.Description, result.Description)
}

func TestDelete(t *testing.T) {
	mockRepo.Mock.On("Delete", uint64(1)).Return(nil)
	err := postService.Delete(uint64(1))
	assert.Nil(t, err)
}

func TestShowAll(t *testing.T) {
	mockRepo.Mock.On("ShowAll").Return(posts, nil)

	result, err := postService.ShowAll()

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	for i, post := range result {
		assert.Equal(t, posts[i], post)
	}
}

func TestFindById(t *testing.T) {
	mockRepo.Mock.On("FindByID", uint64(1)).Return(posts[0], nil)

	result, err := postService.FindByID(uint64(1))
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, posts[0], result)
}
