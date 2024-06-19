package controllers

import (
	"net/http"
	"strconv"

	"recipes/config"
	"recipes/models"

	"github.com/gin-gonic/gin"
)

// SaveRating godoc
// @Summary Save a new rating
// @Description Save a new rating
// @Tags ratings
// @Accept json
// @Produce json
// @Param rating body models.Rating true "Rating"
// @Success 200 {object} models.Rating
// @Router /ratings [post]
func SaveRating(c *gin.Context) {
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&rating).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rating)
}

// GetRatingsByRecipeId godoc
// @Summary Get ratings by recipe ID
// @Description Get ratings by recipe ID
// @Tags ratings
// @Produce json
// @Param recipeId path uint true "Recipe ID"
// @Success 200 {array} models.Rating
// @Router /ratings/recipe/{recipeId} [get]
func GetRatingsByRecipeId(c *gin.Context) {
	recipeID, err := strconv.ParseUint(c.Param("recipeId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recipe ID"})
		return
	}
	var ratings []models.Rating
	if err := config.DB.Where("recipe_id = ?", recipeID).Find(&ratings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ratings)
}

// DeleteRating godoc
// @Summary Delete a rating by ID
// @Description Delete a rating by ID
// @Tags ratings
// @Param id path uint true "Rating ID"
// @Success 204 {string} string "No Content"
// @Router /ratings/{id} [delete]
func DeleteRating(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid rating ID"})
		return
	}
	if err := config.DB.Delete(&models.Rating{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
