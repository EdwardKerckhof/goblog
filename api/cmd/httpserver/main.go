package main

import (
	"log"

	"github.com/edwardkerckhof/goblog/configs"
	"github.com/edwardkerckhof/goblog/internal/server"
	"github.com/edwardkerckhof/goblog/internal/storage"
)

func main() {
	// load environment variables
	config, err := configs.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	// setup app storage
	db := storage.NewPostgresConnection(config)

	// setup handlers
	postHandler := SetupPostHandler(db)

	// create and start server
	httpServer := server.NewServer(
		config,
		postHandler,
	)
	httpServer.Start()
}
