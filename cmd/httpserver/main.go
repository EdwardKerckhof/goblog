package main

import (
	"fmt"

	"github.com/edwardkerckhof/blog/internal/core/domain"
)

func main() {
	fmt.Println("Hello!")

	post := domain.Post{
		ID:    1,
		Title: "Test",
		Body:  "Body",
	}

	fmt.Println(post)
}
