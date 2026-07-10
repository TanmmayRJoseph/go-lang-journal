package main

import (
	"book_store/internals/config"
	"book_store/internals/db"
	"book_store/internals/routes"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("Welcome in the class of CRUD with DB") // a simple logger message

	// Production:
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default() //setting up the gin router

	// Most important part loading up the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Error loading config", "error", err)
		os.Exit(1) // this means that the program is going to exit with a status code of 1
	}

	//Another important part connecting to a database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		slog.Error("Error connecting to database", slog.Any("error", err))
		os.Exit(1)
	}

	// defer func because we want to disconnect from the database after the main function is finished
	defer func() {
		if err := database.DisconnectDB(db); err != nil {
			slog.Error("Error disconnecting from database", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	// setting up the trusted proxies for the API
	// Never trust all proxies in production.
	// if err := router.SetTrustedProxies(nil); err != nil {
	// 	slog.Error("Failed to set trusted proxies", "error", err)
	// 	os.Exit(1)
	// }

	// These are 2 sample endpoints to test the API
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the API"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Registering the routes
	bookroutes.RegisterBookRoutes(router, db.Database("book_store"))

	// Starting the server
	slog.Info("Server started", "port", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
