package server

import (
	"strconv"

	"github.com/edwardkerckhof/goblog/configs"
	"github.com/edwardkerckhof/goblog/internal/core/ports"
	rest "github.com/edwardkerckhof/goblog/internal/http"
)

type Server struct {
	config      *configs.Config
	postHandler ports.PostHandler
}

func NewServer(postHandler ports.PostHandler) *Server {
	return &Server{
		postHandler: postHandler,
	}
}

func (s *Server) Start() {
	router := rest.NewMuXRouter(s.config)

	router.GET("/posts/{id:[0-9]+}", s.postHandler.Get)

	router.SERVE(":" + strconv.Itoa(s.config.ApiPort))
}
