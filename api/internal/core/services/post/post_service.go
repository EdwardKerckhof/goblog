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

// Gets one post from the repository
func (s *service) Get(postID uint) (*domain.Post, error) {
	post, err := s.repo.Get(postID)
	if err != nil {
		return &domain.Post{}, err
	}
	return post, nil
}

// Gets all posts from the repository
func (s *service) GetAll() ([]*domain.Post, error) {
	return s.repo.GetAll()
}

// Creates a new post in the repository
func (s *service) Create(post *domain.Post) (uint, error) {
	return s.repo.Create(post)
}
