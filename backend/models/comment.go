package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model `swaggerignore:"true"`
	RecipeID   uint   `json:"recipe_id"`
	UserID     uint   `json:"user_id"`
	Content    string `json:"content"`
}
