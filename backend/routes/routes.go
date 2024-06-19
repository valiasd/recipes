package routes

import (
	"recipes/controllers"
	_ "recipes/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	router.Use(cors.New(config))

	// API group
	api := router.Group("/api")
	{
		// User routes
		userGroup := api.Group("/users")
		{
			userGroup.POST("/register", controllers.RegisterUser)
			userGroup.GET("/:username", controllers.GetUserByUsername)
		}

		// Recipe routes
		recipeGroup := api.Group("/recipes")
		{
			recipeGroup.POST("/", controllers.SaveRecipe)
			recipeGroup.GET("/author/:authorId", controllers.GetRecipesByAuthorId)
			recipeGroup.DELETE("/:id", controllers.DeleteRecipe)
		}

		// Comment routes
		commentGroup := api.Group("/comments")
		{
			commentGroup.POST("/", controllers.SaveComment)
			commentGroup.GET("/recipe/:recipeId", controllers.GetCommentsByRecipeId)
			commentGroup.DELETE("/:id", controllers.DeleteComment)
		}

		// Rating routes
		ratingGroup := api.Group("/ratings")
		{
			ratingGroup.POST("/", controllers.SaveRating)
			ratingGroup.GET("/recipe/:recipeId", controllers.GetRatingsByRecipeId)
			ratingGroup.DELETE("/:id", controllers.DeleteRating)
		}
	}

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
