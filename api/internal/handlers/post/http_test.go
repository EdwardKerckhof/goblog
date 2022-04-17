package post_handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/edwardkerckhof/goblog/internal/core/domain"
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

func Test_handler_Get(t *testing.T) {
	post := randomPost()

	testCases := []struct {
		name          string
		id            uint
		buildStubs    func(r *repositoriesMock.MockPostRepository)
		checkResponse func(t *testing.T, rr *httptest.ResponseRecorder)
	}{
		{
			name: "StatusOKReturnsDefault",
			id:   post.ID,
			buildStubs: func(r *repositoriesMock.MockPostRepository) {
				r.EXPECT().
					Get(gomock.Eq(post.ID)).
					Times(1).
					Return(post, nil)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rr.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
		})
	}
}
