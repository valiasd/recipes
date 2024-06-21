package controllers

import (
	"net/http"
	"strconv"

	"recipes/config"
	"recipes/models"

	"github.com/gin-gonic/gin"
)

// SaveRecipe godoc
// @Summary Save a new recipe
// @Description Save a new recipe
// @Tags recipes
// @Accept json
// @Produce json
// @Param recipe body models.Recipe true "Recipe"
// @Success 200 {object} models.Recipe
// @Router /recipes [post]
func SaveRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&recipe).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

// GetRecipesByAuthorId godoc
// @Summary Get recipes by author ID
// @Description Get recipes by author ID
// @Tags recipes
// @Produce json
// @Param authorId path uint true "Author ID"
// @Success 200 {array} models.Recipe
// @Router /recipes/author/{authorId} [get]
func GetRecipesByAuthorId(c *gin.Context) {
	authorID, err := strconv.ParseUint(c.Param("authorId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}
	var recipes []models.Recipe
	if err := config.DB.Where("author_id = ?", authorID).Find(&recipes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

// GetRecipeById godoc
// @Summary Get single recipe by ID
// @Description Get single recipe by ID
// @Tags recipes
// @Produce json
// @Param recipeId path uint true "Recipe ID"
// @Success 200 {object} models.Recipe
// @Router /recipes/{recipeId} [get]
func GetRecipeById(c *gin.Context) {
	recipeID, err := strconv.ParseUint(c.Param("recipeId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}
	var recipe models.Recipe
	if err := config.DB.Where("ID = ?", recipeID).First(&recipe).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

// DeleteRecipe godoc
// @Summary Delete a recipe by ID
// @Description Delete a recipe by ID
// @Tags recipes
// @Param id path uint true "Recipe ID"
// @Success 204 {string} string "No Content"
// @Router /recipes/{id} [delete]
func DeleteRecipe(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recipe ID"})
		return
	}
	if err := config.DB.Delete(&models.Recipe{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
