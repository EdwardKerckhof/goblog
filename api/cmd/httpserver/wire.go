//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	postService "github.com/edwardkerckhof/goblog/internal/core/services/post"
	postHandler "github.com/edwardkerckhof/goblog/internal/handlers/post"
	postRepository "github.com/edwardkerckhof/goblog/internal/repositories/post"
)

func SetupPostHandler(db *gorm.DB) postHandler.PostHandler {
	wire.Build(postHandler.NewHTTPHandler, postService.NewPostService, postRepository.NewGormRepository)
	return &postHandler.PostHandlerImpl{}
}
