package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *Service
}

func NewUserHandler(svc *Service) *UserHandler {
	return &UserHandler{svc: svc}
}

// REGISTER USER
func (uh *UserHandler) RegisterUserHandler(c *gin.Context) {
	var regsInput RegisterInput
	// small validation
	if err := c.ShouldBindJSON(&regsInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call the service layer
	output, err := uh.svc.RegisterUserService(c.Request.Context(), regsInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)

}

// LOGIN USER
func (uh *UserHandler) LoginUserHandler(c *gin.Context) {
	var loginInput LoginInput
	// small validation
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call the service layer
	output, err := uh.svc.LoginUserService(c.Request.Context(), loginInput, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
