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
func (r *repository) GetAll() ([]*domain.Post, error) {
	var posts []*domain.Post
	err := r.db.Find(&posts).Error
	return posts, err
}

// Creates a post using GORM
func (r *repository) Create(post *domain.Post) (*domain.Post, error) {
	err := r.db.Create(&post).Error
	return post, err
}

func (r *repository) Delete(post *domain.Post) {
	r.db.Delete(&post)
}
