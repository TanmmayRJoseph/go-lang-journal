package httpserver

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// A simple health check route
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "Server is healthy",
		"status": "ok",
		"time":   time.Now().UTC(),
	})
}
