package main

import (
	"log"

	"github.com/edwardkerckhof/goblog/configs"
	postService "github.com/edwardkerckhof/goblog/internal/core/services/post"
	postHandler "github.com/edwardkerckhof/goblog/internal/handlers/post"
	postRepository "github.com/edwardkerckhof/goblog/internal/repositories/post"
	"github.com/edwardkerckhof/goblog/internal/server"
	"github.com/edwardkerckhof/goblog/internal/storage"
)

func main() {
	config, err := configs.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("error loading app.env file: %s", err.Error())
	}

	db := storage.NewPostgresConnection(config)

	// TODO: add google/wire
	postRepo := postRepository.NewGormRepository(db)
	postService := postService.NewPostService(postRepo)
	postHandler := postHandler.NewHTTPHandler(postService)

	httpServer := server.NewServer(
		postHandler,
	)
	httpServer.Start()
}
