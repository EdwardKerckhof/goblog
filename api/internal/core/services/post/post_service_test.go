package post_service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/edwardkerckhof/goblog/internal/core/domain"
	"github.com/edwardkerckhof/goblog/internal/core/ports"
	repositoriesMock "github.com/edwardkerckhof/goblog/mocks/repositories"
)

func randomPost() *domain.Post {
	return &domain.Post{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		Title: "Test title",
		Body:  "Test Body",
	}
}

func Test_service_Get(t *testing.T) {
	post := randomPost()

	testCases := []struct {
		name          string
		buildStubs    func(r *repositoriesMock.MockPostRepository)
		checkResponse func(t *testing.T, s ports.PostService)
	}{
		{
			name: "OKReturnsDefault",
			buildStubs: func(r *repositoriesMock.MockPostRepository) {
				r.EXPECT().
					Get(gomock.Eq(post.ID)).
					Times(1).
					Return(post, nil)
			},
			checkResponse: func(t *testing.T, s ports.PostService) {
				res, err := s.Get(post.ID)
				require.NoError(t, err)
				requireEqual(t, res, post)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			r := repositoriesMock.NewMockPostRepository(c)
			tc.buildStubs(r)

			service := NewPostService(r)
			tc.checkResponse(t, service)
		})
	}
}

func Test_service_GetAll(t *testing.T) {
	const n = 5
	posts := make([]*domain.Post, n)
	for i := 0; i < n; i++ {
		posts = append(posts, randomPost())
	}

	testCases := []struct {
		name          string
		buildStubs    func(r *repositoriesMock.MockPostRepository)
		checkResponse func(t *testing.T, s ports.PostService)
	}{
		{
			name: "OKReturnsDefault",
			buildStubs: func(r *repositoriesMock.MockPostRepository) {
				r.EXPECT().
					GetAll().
					Times(1).
					Return(posts, nil)
			},
			checkResponse: func(t *testing.T, s ports.PostService) {
				res, err := s.GetAll()
				require.NoError(t, err)
				requireEquals(t, res, posts)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			r := repositoriesMock.NewMockPostRepository(c)
			tc.buildStubs(r)

			service := NewPostService(r)
			tc.checkResponse(t, service)
		})
	}
}

func requireEquals(t *testing.T, results []*domain.Post, posts []*domain.Post) {
	for i, res := range results {
		if res != nil && posts[i] != nil {
			requireEqual(t, res, posts[i])
		}
	}
}

func requireEqual(t *testing.T, res *domain.Post, post *domain.Post) {
	require.Equal(t, res.ID, post.ID)
	require.Equal(t, res.Title, post.Title)
	require.Equal(t, res.Body, post.Body)
	require.Equal(t, res.DeletedAt, post.DeletedAt)
	require.WithinDuration(t, res.CreatedAt, post.CreatedAt, time.Second)
	require.WithinDuration(t, res.UpdatedAt, post.UpdatedAt, time.Second)
}
