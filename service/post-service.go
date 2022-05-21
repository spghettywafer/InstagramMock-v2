package service

import (
	model "InstagramMock-v2/model"
	repo "InstagramMock-v2/repository"
)

type (
	PostCreateDTO struct {
		UrlPhoto    string `json:"urlPhoto" binding:"required"`
		Description string `json:"description" binding:"required"`
		Username    string `json:"username" binding:"required"`
	}

	PostUpdateDTO struct {
		ID          uint64 `json:"id"`
		Description string `json:"description" binding:"required"`
	}
)

type (
	IPostService interface {
		Create(p PostCreateDTO) (model.Post, error)
		Update(p PostUpdateDTO) (model.Post, error)
		Delete(postID uint64) error
		ShowAll() ([]model.Post, error)
		FindByID(postID uint64) (model.Post, error)
	}

	PostService struct {
		postRepo repo.IPostRepository
	}
)

func NewPostService(r repo.IPostRepository) *PostService {
	return &PostService{r}
}

func (s *PostService) Create(p PostCreateDTO) (model.Post, error) {
	var post model.Post = model.Post{
		UrlPhoto:    p.UrlPhoto,
		Description: p.Description,
		Username:    p.Username,
	}

	return s.postRepo.Create(post)

}

func (s *PostService) Update(p PostUpdateDTO) (model.Post, error) {
	var post model.Post = model.Post{
		ID:          p.ID,
		Description: p.Description,
	}

	return s.postRepo.Update(post)
}

func (s *PostService) Delete(postID uint64) error {
	return s.postRepo.Delete(postID)
}

func (s *PostService) ShowAll() ([]model.Post, error) {
	posts, err := s.postRepo.ShowAll()
	if err != nil {
		return posts, err
	}

	for i := range posts {
		posts[i].Description = posts[i].GetShortDescription()
	}

	return posts, nil
}

func (s *PostService) FindByID(postID uint64) (model.Post, error) {
	return s.postRepo.FindByID(postID)
}
