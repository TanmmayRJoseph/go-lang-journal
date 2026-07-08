package routes

import (
	"crudapi/internal/handlers"
	"log"
	"net/http"
)

// RegisterRoutes registers all API endpoints.
//
// Think of this function as the receptionist of the application.
// Every incoming request first comes here.
//
// Its only job is:
//
// URL  --->  Handler
//
// It does NOT contain any business logic.
func RegisterRoutes(mux *http.ServeMux) {
	log.Println("Registered POST /studentDto")

	// --------------------------------------------------
	// GET /students
	// --------------------------------------------------
	//
	// If someone sends:
	//
	// GET /students
	//
	// then call:
	//
	// handlers.GetStudent()
	//
	mux.HandleFunc("GET /students", handlers.GetStudent)

	// --------------------------------------------------
	// POST /students
	// --------------------------------------------------
	//
	// If someone sends:
	//
	// POST /students
	//
	// call CreateStudent().
	//
	mux.HandleFunc("POST /students", handlers.CreateStudent)

	// --------------------------------------------------
	// GET /students/{id}
	// --------------------------------------------------
	//
	// Example:
	//
	// GET /students/5
	//
	// This route matches any value after /students/.
	//
	mux.HandleFunc("GET /students/", handlers.GetStudentById)

	// --------------------------------------------------
	// PUT /students/{id}
	// --------------------------------------------------
	//
	// Example:
	//
	// PUT /students/10
	//
	mux.HandleFunc("PUT /students/", handlers.UpdateStudentById)

	// --------------------------------------------------
	// DELETE /students/{id}
	// --------------------------------------------------
	//
	// Example:
	//
	// DELETE /students/25
	//
	mux.HandleFunc("DELETE /students/", handlers.DeleteStudentById)

	// --------------------------------------------------
	// POST /studentDto
	// --------------------------------------------------
	//
	// If someone sends:
	//
	// POST /studentDto
	//
	// call CreateStudentWithDto().
	//
	mux.HandleFunc("POST /studentDto", handlers.CreateStudentWithDto)
}
