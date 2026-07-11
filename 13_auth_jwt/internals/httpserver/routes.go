package httpserver

import (
	"auth_jwt/internals/app"
	"auth_jwt/internals/middleware"
	"auth_jwt/internals/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerRoutes(a *app.App) *gin.Engine {

	// returns a blank instance without any middleware
	r := gin.New()

	// logger middleware logs the request as it goes in and the request as it goes out
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	// ==============================================================================================

	// a simple home route
	r.GET("/", rootHandler)

	// using that helath check that we created above to check if the server is healthy or not.
	r.GET("/health", healthHandler)

	// ===============================================================================================

	// user routes
	userRepo := users.NewUserRepo(a.DB)
	userSvc := users.NewUserService(userRepo, a.Config.JWTSecret)
	userHandler := users.NewUserHandler(userSvc)

	// Unprotected routes
	usergroup := r.Group("/users")
	{
		usergroup.POST("/register", userHandler.RegisterUserHandler)
		usergroup.POST("/login", userHandler.LoginUserHandler)
	}

	// just demonstrating the use of middleware
	protected := r.Group("/protected").Use(middleware.AuthRequireMiddleware(a.Config.JWTSecret))
	{

		// Protected routes only for admin
		protected.GET("/admin", middleware.AdminRequireMiddleware(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "You are an admin"})
		})

		// protected routes only for user
		protected.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "You are a user"})
		})

	}

	return r
}
