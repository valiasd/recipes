package controllers

import (
	"net/http"
	"strconv"

	"recipes/config"
	"recipes/models"

	"github.com/gin-gonic/gin"
)

// SaveComment godoc
// @Summary Save a new comment
// @Description Save a new comment
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Comment"
// @Success 200 {object} models.Comment
// @Router /comments [post]
func SaveComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

// GetCommentsByRecipeId godoc
// @Summary Get comments by recipe ID
// @Description Get comments by recipe ID
// @Tags comments
// @Produce json
// @Param recipeId path uint true "Recipe ID"
// @Success 200 {array} models.Comment
// @Router /comments/recipe/{recipeId} [get]
func GetCommentsByRecipeId(c *gin.Context) {
	recipeID, err := strconv.ParseUint(c.Param("recipeId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recipe ID"})
		return
	}
	var comments []models.Comment
	if err := config.DB.Where("recipe_id = ?", recipeID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// DeleteComment godoc
// @Summary Delete a comment by ID
// @Description Delete a comment by ID
// @Tags comments
// @Param id path uint true "Comment ID"
// @Success 204 {string} string "No Content"
// @Router /comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment ID"})
		return
	}
	if err := config.DB.Delete(&models.Comment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
