package post_handler

import (
	"time"

	"github.com/edwardkerckhof/goblog/internal/core/domain"
)

// Post resonse
// swagger:response PostResponse
type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreatePostResponse(post *domain.Post) *PostResponse {
	return &PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Body:      post.Body,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func CreatePostListResponse(posts []*domain.Post) []*PostResponse {
	postList := []*PostResponse{}
	for _, p := range posts {
		post := CreatePostResponse(p)
		postList = append(postList, post)
	}
	return postList
}
