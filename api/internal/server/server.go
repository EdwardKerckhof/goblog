package server

import (
	"strconv"

	"github.com/edwardkerckhof/goblog/configs"
	postHandler "github.com/edwardkerckhof/goblog/internal/handlers/post"
	rest "github.com/edwardkerckhof/goblog/internal/http"
)

type Server struct {
	config      *configs.Config
	postHandler postHandler.PostHandler
}

// NewServer creates a new HTTP server
func NewServer(config *configs.Config, postHandler postHandler.PostHandler) *Server {
	return &Server{
		config:      config,
		postHandler: postHandler,
	}
}

// Start creates a new REST router and serves the application
func (s *Server) Start() {
	router := rest.NewMuXRouter(s.config)

	router.GET("/posts/{id:[0-9]+}", s.postHandler.Get)

	router.SERVE(":" + strconv.Itoa(s.config.ApiPort))
}
