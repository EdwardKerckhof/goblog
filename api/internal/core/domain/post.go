package domain

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string `json:"title" gorm:"unique;not null;index"`
	Body  string `json:"body" gorm:"not null"`
}

func (p *Post) TableName() string {
	return "posts"
}
