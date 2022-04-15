package post_service

import (
	"github.com/edwardkerckhof/goblog/internal/core/domain"
	"github.com/edwardkerckhof/goblog/internal/core/ports"
)

type service struct {
	repo ports.PostRepository
}

func NewPostService(repo ports.PostRepository) ports.PostService {
	return &service{
		repo: repo,
	}
}

func (s *service) Get(postID uint) (*domain.Post, error) {
	return s.repo.Get(postID)
}

func (s *service) GetAll() ([]*domain.Post, error) {
	return s.repo.GetAll()
}

func (s *service) Create(post *domain.Post) (uint, error) {
	return s.repo.Create(post)
}
