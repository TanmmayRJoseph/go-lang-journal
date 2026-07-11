package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRequireMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := GetUserRole(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "this route is for admin only"})
			return
		}

		c.Next()
	}
}
