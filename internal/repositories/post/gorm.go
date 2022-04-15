package post_repository

import (
	"gorm.io/gorm"

	"github.com/edwardkerckhof/goblog/internal/core/domain"
	"github.com/edwardkerckhof/goblog/internal/core/ports"
)

type repository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) ports.PostRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(postID uint) (*domain.Post, error) {
	var post domain.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *repository) GetAll() ([]*domain.Post, error) {
	return []*domain.Post{}, nil
}

func (r *repository) Create(post *domain.Post) (uint, error) {
	result := r.db.Create(&post)
	return post.ID, result.Error
}
