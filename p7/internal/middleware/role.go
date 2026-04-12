package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.Value("role").(string)

		if role != requiredRole {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized role"})
			return
		}

		c.Next()
	}
}
