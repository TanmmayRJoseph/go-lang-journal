package main

import (
	"fmt"
	"log"
	"net/http"

	"crudapi/internal/config"
	"crudapi/internal/routes"
)

func main() {

	// --------------------------------------------------
	// Step 1 : Welcome Message
	// --------------------------------------------------
	//
	// Just to know our application has started.
	//
	fmt.Println("🚀 Starting CRUD API...")

	// --------------------------------------------------
	// Step 2 : Load Configuration
	// --------------------------------------------------
	//
	// This reads the .env file and returns
	// a pointer to the Config struct.
	//
	cfg := config.LoadConfig()

	// --------------------------------------------------
	// Step 3 : Create a Router (ServeMux)
	// --------------------------------------------------
	//
	// ServeMux is Go's built-in router.
	//
	// Think of it as the receptionist
	// that knows which URL should go
	// to which handler.
	//
	mux := http.NewServeMux()

	// --------------------------------------------------
	// Step 4 : Register all API routes
	// --------------------------------------------------
	//
	// This function tells the router:
	//
	// GET /students
	//      ↓
	// GetStudent()
	//
	// POST /students
	//      ↓
	// CreateStudent()
	//
	routes.RegisterRoutes(mux)

	// --------------------------------------------------
	// Step 5 : Print server information
	// --------------------------------------------------
	//
	fmt.Println("--------------------------------")
	fmt.Println("✅ Server Started Successfully")
	fmt.Println("🌍 Running on Port:", cfg.Port)
	fmt.Println("--------------------------------")

	// --------------------------------------------------
	// Step 6 : Start HTTP Server
	// --------------------------------------------------
	//
	// ListenAndServe starts the web server.
	//
	// First argument:
	// ":" + cfg.Port
	//
	// Example:
	//
	// ":8080"
	//
	// Second argument:
	//
	// Our router (ServeMux)
	//
	// From this point onward,
	// the server keeps listening
	// for incoming requests.
	//
	err := http.ListenAndServe(":"+cfg.Port, mux)

	// --------------------------------------------------
	// Step 7 : Check for server errors
	// --------------------------------------------------
	//
	// ListenAndServe blocks forever.
	//
	// It only returns if something
	// goes wrong.
	//
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
