package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `json:"username" gorm:"unique"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
