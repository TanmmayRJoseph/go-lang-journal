package handlers

import (
	"crudapi/internal/dtos"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// APIResponse represents the common structure
// that every API endpoint will return.
//
// Instead of returning random JSON for every API,
// we return the same structure everywhere.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ----------------------------------------------------
// GET /students
// ----------------------------------------------------
func GetStudent(w http.ResponseWriter, r *http.Request) {

	// Tell the client that the response will be JSON.
	w.Header().Set("Content-Type", "application/json")

	// Set HTTP Status Code = 200 OK
	w.WriteHeader(http.StatusOK)

	// Create our response object.
	response := APIResponse{
		Success: true,
		Message: "All students fetched successfully",
		Data:    nil, // Later this will contain []Student
	}

	// Convert Go struct -> JSON and send it back.
	json.NewEncoder(w).Encode(response)
}

// ----------------------------------------------------
// POST /students
// ----------------------------------------------------
func CreateStudent(w http.ResponseWriter, r *http.Request) {

	// Tell browser/Postman that we're returning JSON.
	w.Header().Set("Content-Type", "application/json")

	// Resource created successfully.
	w.WriteHeader(http.StatusCreated)

	response := APIResponse{
		Success: true,
		Message: "Student created successfully",
		Data:    nil,
	}

	json.NewEncoder(w).Encode(response)
}

// ----------------------------------------------------
// GET /students/{id}
// ----------------------------------------------------
func GetStudentById(w http.ResponseWriter, r *http.Request) {

	// Example URL:
	// /students/5
	//
	// r.URL.Path gives:
	// "/students/5"

	path := r.URL.Path

	// Remove "/students/" from the beginning.
	//
	// "/students/5"
	//
	// becomes
	//
	// "5"
	id := strings.TrimPrefix(path, "/students/")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := APIResponse{
		Success: true,
		Message: "Student fetched successfully",
		Data: map[string]string{
			"id": id,
		},
	}

	json.NewEncoder(w).Encode(response)
}

// ----------------------------------------------------
// PUT /students/{id}
// ----------------------------------------------------
func UpdateStudentById(w http.ResponseWriter, r *http.Request) {

	// Extract the student ID from URL.
	id := strings.TrimPrefix(r.URL.Path, "/students/")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := APIResponse{
		Success: true,
		Message: "Student updated successfully",
		Data: map[string]string{
			"id": id,
		},
	}

	json.NewEncoder(w).Encode(response)
}

// ----------------------------------------------------
// DELETE /students/{id}
// ----------------------------------------------------
func DeleteStudentById(w http.ResponseWriter, r *http.Request) {

	// Extract the student ID from URL.
	id := strings.TrimPrefix(r.URL.Path, "/students/")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := APIResponse{
		Success: true,
		Message: "Student deleted successfully",
		Data: map[string]string{
			"id": id,
		},
	}

	json.NewEncoder(w).Encode(response)
}

// used Dto here
func CreateStudentWithDto(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside CreateStudentWithDto")
	// Tell the client that we're returning JSON.
	w.Header().Set("Content-Type", "application/json")

	// Create an empty DTO.
	//
	// Think of this as an empty admission form.
	var req dtos.CreateStudentRequest

	// Read the JSON body sent by the client
	// and fill the req struct.
	//
	// Example JSON:
	//
	// {
	//     "name": "Tanmmay",
	//     "age": 21
	// }
	//
	// After decoding:
	//
	// req.Name = "Tanmmay"
	// req.Age = 21
	err := json.NewDecoder(r.Body).Decode(&req)

	// If the JSON is invalid,
	// return HTTP 400 Bad Request.
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Invalid request body",
			Data:    nil,
		})

		return
	}

	// Everything is valid.
	w.WriteHeader(http.StatusCreated)

	// Return the received data back to the client.
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "Student created successfully",
		Data:    req,
	})
}
