package controllers

import (
	"fmt"
	"net/http"

	"recipes/config"
	"recipes/models"

	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Router /users/register [post]
func RegisterUser(c *gin.Context) {
	fmt.Println("REGISTER!")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetUserByUsername godoc
// @Summary Get user by username
// @Description Get user by username
// @Tags users
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.User
// @Router /users/{username} [get]
func GetUserByUsername(c *gin.Context) {
	fmt.Println("GET USER")
	username := c.Param("username")
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
