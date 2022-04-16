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
	GetAll() ([]*domain.Post, error)
	Create(post *domain.Post) (*domain.Post, error)
	Delete(post *domain.Post)
}
