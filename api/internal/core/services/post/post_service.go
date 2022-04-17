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
	return s.repo.Get(postID)
}

// Gets all posts from the repository
func (s *service) GetAll(offset int) ([]*domain.Post, error) {
	return s.repo.GetAll(offset)
}

// Creates a new post in the repository
func (s *service) Create(post *domain.Post) (*domain.Post, error) {
	return s.repo.Create(post)
}

// Updates a post in the repository
func (s *service) Update(post *domain.Post) (*domain.Post, error) {
	return s.repo.Update(post)
}

// Soft deletes a post in the repository
func (s *service) Delete(post *domain.Post) error {
	return s.repo.Delete(post)
}
