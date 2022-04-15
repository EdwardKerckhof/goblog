package ports

import "github.com/edwardkerckhof/goblog/internal/core/domain"

type PostService interface {
	Get(postID uint) (*domain.Post, error)
	GetAll() ([]*domain.Post, error)
	Create(post *domain.Post) (uint, error)
}
