package repository

import (
	"InstagramMock-v2/model"

	"gorm.io/gorm"
)

type (
	IPostRepository interface {
		Create(post model.Post) (model.Post, error)
		Update(post model.Post) (model.Post, error)
		Delete(postID uint64) error
		ShowAll() ([]model.Post, error)
		FindByID(postID uint64) (model.Post, error)
	}

	PostRepository struct {
		db *gorm.DB
	}
)

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) Create(post model.Post) (model.Post, error) {
	if err := r.db.Create(&post).Error; err != nil {
		return post, err
	}

	return post, nil
}

func (r *PostRepository) Update(post model.Post) (model.Post, error) {
	if err := r.db.Model(&post).Update("description", post.Description).Error; err != nil {
		return post, err
	}

	return post, nil
}

func (r *PostRepository) Delete(postID uint64) error {
	if err := r.db.Delete(&model.Post{}, postID).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) ShowAll() ([]model.Post, error) {
	var posts []model.Post
	if err := r.db.Order("created_at desc").Find(&posts).Error; err != nil {
		return posts, err
	}
	return posts, nil
}

func (r *PostRepository) FindByID(postID uint64) (model.Post, error) {
	var post model.Post

	if err := r.db.First(&post, postID).Error; err != nil {
		return post, err
	}

	return post, nil
}
