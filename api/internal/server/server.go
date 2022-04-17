package server

import (
	"strconv"

	"github.com/edwardkerckhof/goblog/configs"
	"github.com/edwardkerckhof/goblog/internal/core/ports"
	postHandler "github.com/edwardkerckhof/goblog/internal/handlers/post"
	rest "github.com/edwardkerckhof/goblog/internal/http"
)

type Server struct {
	config      *configs.Config
	postHandler postHandler.PostHandler
	Router      ports.HTTPRouter
}

// NewServer creates a new HTTP server
func NewServer(config *configs.Config, postHandler postHandler.PostHandler) *Server {
	server := &Server{
		config:      config,
		postHandler: postHandler,
	}

	server.setupRouter()
	return server
}

func (s *Server) setupRouter() {
	router := rest.NewMuXRouter(s.config)

	router.GET("/posts/{id:[0-9]+}", s.postHandler.Get)
	router.GET("/posts", s.postHandler.GetAll)

	s.Router = router
}

// Start creates a new REST router and serves the application
func (s *Server) Start() {
	s.Router.SERVE(":" + strconv.Itoa(s.config.ApiPort))
}
