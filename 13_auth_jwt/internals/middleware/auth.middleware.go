package middleware

import (
	"auth_jwt/internals/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// store -> auth data info in gin context

const (
	ctxUserIdKey = "auth.userId" //userID _id
	ctxRoleKey   = "auth.role"   //role (admin, user)
)

// Authorization : `Bearer <token>`
func AuthRequireMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// get token from header
		authHeader := strings.TrimSpace(c.Request.Header.Get("Authorization"))
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing or invalid"})
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		// validate token
		claims, err := utils.VerifyToken(jwtSecret, tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// store user id and role in gin context
		c.Set(ctxUserIdKey, claims.Subject)
		c.Set(ctxRoleKey, claims.Role)

		c.Next()
	}
}

func GetUserID(c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxUserIdKey)
	if !ok {
		return "", false
	}

	userId, ok := res.(string)
	return userId, ok
}


func GetUserRole(c *gin.Context) (string, bool) {
	res, ok := c.Get(ctxRoleKey)
	if !ok {
		return "", false
	}

	role, ok := res.(string)
	return role, ok
}
