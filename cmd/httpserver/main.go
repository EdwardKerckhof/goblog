package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/edwardkerckhof/goblog/configs"
	postService "github.com/edwardkerckhof/goblog/internal/core/services/post"
	postHandler "github.com/edwardkerckhof/goblog/internal/handlers/post"
	rest "github.com/edwardkerckhof/goblog/internal/http"
	postRepository "github.com/edwardkerckhof/goblog/internal/repositories/post"
	"github.com/edwardkerckhof/goblog/internal/storage"
)

func main() {
	config, err := configs.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("error loading app.env file: %s", err.Error())
	}
	fmt.Println(config.ApiPort)

	db := storage.NewPostgresConnection(config)

	// TODO: add google/wire
	postRepo := postRepository.NewGormRepository(db)
	postService := postService.NewPostService(postRepo)
	postHandler := postHandler.NewHTTPHandler(postService)

	router := rest.NewMuXRouter(config)
	router.GET("/posts/{id:[0-9]+}", postHandler.Get)

	router.SERVE(":" + strconv.Itoa(config.ApiPort))
}
