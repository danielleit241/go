package api

import (
	"net/http"

	"github.com/danielleit241/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

type DeleteAPIKeyParam struct {
	ID int `uri:"id" binding:"gt=0"`
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ctrl *AuthController) GenerateAPIKey(c *gin.Context) {
	keyID, apiKey, err := utils.CreateAPIKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate API key"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "API key generated successfully",
		"id":      keyID,
		"api_key": apiKey,
	})
}

func (ctrl *AuthController) DeleteAPIKey(c *gin.Context) {
	var param DeleteAPIKeyParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid API key id"})
		return
	}

	if !utils.DeleteAPIKeyByID(param.ID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "API key not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "API key deleted successfully",
		"id":      param.ID,
	})
}
