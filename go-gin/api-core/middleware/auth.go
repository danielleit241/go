package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/danielleit241/utils"
	"github.com/gin-gonic/gin"
)

func ApiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/v1/auth/") {
			c.Next()
			return
		}

		requestAPIKey := c.GetHeader("X-API-Key")
		if requestAPIKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API key is required"})
			c.Abort()
			return
		}

		if !utils.HasAnyAPIKey() {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "API key is not initialized"})
			c.Abort()
			return
		}

		if !utils.IsValidAPIKey(requestAPIKey) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func AdminSecret() gin.HandlerFunc {
	expectedAdminSecret := os.Getenv("ADMIN_SECRET")

	return func(c *gin.Context) {
		if expectedAdminSecret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Server misconfigured: ADMIN_SECRET is not set"})
			c.Abort()
			return
		}

		adminSecret := c.GetHeader("X-Admin-Secret")
		if !utils.IsSecretMatch(adminSecret, expectedAdminSecret) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
