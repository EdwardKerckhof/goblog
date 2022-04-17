package post_handler_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/edwardkerckhof/goblog/internal/core/domain"
	postService "github.com/edwardkerckhof/goblog/internal/core/services/post"
	postHandler "github.com/edwardkerckhof/goblog/internal/handlers/post"
	repositoriesMock "github.com/edwardkerckhof/goblog/mocks/repositories"
)

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
				requireBodyEqualsPost(t, rr.Body, post)
				require.Equal(t, http.StatusOK, rr.Code)
			},
		},
		{
			name: "StatusNotFound",
			id:   post.ID,
			buildStubs: func(r *repositoriesMock.MockPostRepository) {
				r.EXPECT().
					Get(gomock.Eq(post.ID)).
					Times(1).
					Return(&domain.Post{}, gorm.ErrRecordNotFound)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, rr.Code)
			},
		},
		{
			name: "StatusInternalServerError",
			id:   post.ID,
			buildStubs: func(r *repositoriesMock.MockPostRepository) {
				r.EXPECT().
					Get(gomock.Eq(post.ID)).
					Times(1).
					Return(&domain.Post{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, rr.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			r := repositoriesMock.NewMockPostRepository(c)
			tc.buildStubs(r)

			rr := httptest.NewRecorder()

			s := postService.NewPostService(r)
			h := postHandler.NewHTTPHandler(s)
			server := newTestServer(t, h)

			url := fmt.Sprintf("/api/v1/posts/%d", tc.id)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.Router.SERVEHTTP(rr, req)

			tc.checkResponse(t, rr)
		})
	}
}

func requireBodyEqualsPost(t *testing.T, body *bytes.Buffer, post *domain.Post) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var bodyDTO postHandler.PostResponse
	err = json.Unmarshal(data, &bodyDTO)
	require.NoError(t, err)
	require.NotEmpty(t, bodyDTO)

	postResponse := postHandler.CreatePostResponse(post)
	requireBodyResponseMatchPostResponse(t, bodyDTO, *postResponse)
}

func requireBodyResponseMatchPostResponse(t *testing.T, bodyResponse postHandler.PostResponse, postResponse postHandler.PostResponse) {
	require.Equal(t, bodyResponse.ID, postResponse.ID)
	require.Equal(t, bodyResponse.Title, postResponse.Title)
	require.Equal(t, bodyResponse.Body, postResponse.Body)
	require.Equal(t, bodyResponse.DeletedAt, postResponse.DeletedAt)
	require.WithinDuration(t, bodyResponse.CreatedAt, postResponse.CreatedAt, time.Second)
	require.WithinDuration(t, bodyResponse.UpdatedAt, postResponse.UpdatedAt, time.Second)
}
