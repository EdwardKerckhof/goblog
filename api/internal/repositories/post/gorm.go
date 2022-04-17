package post_repository

import (
	"gorm.io/gorm"

	"github.com/edwardkerckhof/goblog/internal/core/domain"
	"github.com/edwardkerckhof/goblog/internal/core/ports"
)

type repository struct {
	db *gorm.DB
}

// NewGormRepository creates a new GORM repository for posts
func NewGormRepository(db *gorm.DB) ports.PostRepository {
	return &repository{
		db: db,
	}
}

// Gets a post using GORM
func (r *repository) Get(postID uint) (*domain.Post, error) {
	var post *domain.Post
	err := r.db.First(&post, postID).Error
	return post, err
}

// Gets all posts using GORM
func (r *repository) GetAll(offset int) ([]*domain.Post, error) {
	var posts []*domain.Post
	err := r.db.Find(&posts).Offset(offset).Limit(50).Error
	return posts, err
}

// Creates a post using GORM
func (r *repository) Create(post *domain.Post) (*domain.Post, error) {
	err := r.db.Create(&post).Error
	return post, err
}

// Updates a post using GORM
func (r *repository) Update(post *domain.Post) (*domain.Post, error) {
	err := r.db.Save(&post).Error
	return post, err
}

// Deletes a post using GORM
func (r *repository) Delete(post *domain.Post) error {
	return r.db.Delete(&post).Error
}
