package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to my API",
	})
}
