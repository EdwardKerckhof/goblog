package ports

import "github.com/edwardkerckhof/goblog/internal/core/domain"

type DBConnection struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type PostRepository interface {
	Get(postID uint) (*domain.Post, error)
	GetAll(offset int) ([]*domain.Post, error)
	Create(post *domain.Post) (*domain.Post, error)
	Update(post *domain.Post) (*domain.Post, error)
	Delete(post *domain.Post) error
}
