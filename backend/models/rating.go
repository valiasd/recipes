package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model `swaggerignore:"true"`
	RecipeID   uint `json:"recipe_id"`
	UserID     uint `json:"user_id"`
	Score      int  `json:"score"`
}
