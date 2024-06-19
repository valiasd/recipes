package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorID    uint   `json:"author_id"`
}
