package post_handler_test

import (
	"os"
	"testing"
	"time"

	"github.com/edwardkerckhof/goblog/configs"
	"github.com/edwardkerckhof/goblog/internal/core/domain"
	postHandler "github.com/edwardkerckhof/goblog/internal/handlers/post"
	"github.com/edwardkerckhof/goblog/internal/server"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func newTestServer(t *testing.T, postHandler postHandler.PostHandler) *server.Server {
	config, err := configs.LoadConfig("../../../configs")
	require.NoError(t, err)

	return server.NewServer(config, postHandler)
}

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

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
