package main

import (
	"log"

	"github.com/edwardkerckhof/goblog/configs"
	"github.com/edwardkerckhof/goblog/internal/server"
	"github.com/edwardkerckhof/goblog/internal/storage"
)

func main() {
	config, err := configs.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("error loading app.env file: %s", err.Error())
	}

	db := storage.NewPostgresConnection(config)

	postHandler := SetupPostHandler(db)

	httpServer := server.NewServer(
		postHandler,
	)
	httpServer.Start()
}
