package main

import (
	"recipes/config"
	"recipes/models"
	"recipes/routes"
)

// @title Recipe App API
// @version 1.0
// @description This is the server for a recipe application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{}, &models.Recipe{}, &models.Comment{}, &models.Rating{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
