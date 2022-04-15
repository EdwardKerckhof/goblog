package domain

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
