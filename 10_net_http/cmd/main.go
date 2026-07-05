package main

import (
	"encoding/json" // package that helps convert Go data (like structs/maps) ↔ JSON text
	"fmt"           // package for printing stuff to the terminal
	"net/http"      // the main package that lets Go act like a web server
	"time"          // package for working with date and time
)

func main() {
	fmt.Println("Net-HTTP") // just prints "Net-HTTP" to your terminal when the server starts

	// ─── REGISTERING ROUTES ───────────────────────────────────────────────────
	// http.HandleFunc(pattern, handlerFunc)
	// Think of it like a receptionist: "when someone visits THIS URL, call THIS function"
	// The first arg is the URL path, the second is the function that handles it

	http.HandleFunc("/", rootHandler) // when someone visits http://localhost:8080/  → call rootHandler

	http.HandleFunc("/hello", helloHandler) // when someone visits /hello  → call helloHandler

	http.HandleFunc("/search", queryHandler) // when someone visits /search  → call queryHandler

	http.HandleFunc("/json", successHandlerJson) // when someone visits /json  → call successHandlerJson

	// ─── STARTING THE SERVER ──────────────────────────────────────────────────
	// http.ListenAndServe(address, handler)
	// This actually STARTS the web server and makes it listen for incoming requests
	// ":8080" means "listen on port 8080 of this machine"
	// nil means "use the default router" (the one we set up with HandleFunc above)
	// This function BLOCKS forever — the code below it only runs if the server crashes
	err := http.ListenAndServe(":8080", nil)

	// These two lines only print if the server fails to start (e.g., port already in use)
	fmt.Println("Server started at port 8080")
	fmt.Println(err)
}

// ============================================================
// =======================ROUTE HANDLERS=======================
// ============================================================

// Every handler function MUST have these exact two parameters:
//
//	w http.ResponseWriter → like a notepad you write your REPLY on and hand back to the user
//	r *http.Request       → the actual REQUEST that came in (contains URL, method, body, headers, etc.)
func helloHandler(w http.ResponseWriter, r *http.Request) {

	// r.Method → tells you what HTTP method was used (GET, POST, PUT, DELETE, etc.)
	// http.MethodGet is just the string "GET" — using the constant is cleaner than typing "GET"
	// If the method is NOT GET, we reject the request
	if r.Method != http.MethodGet {
		// http.Error(w, message, statusCode)
		// A shortcut to send an error response back to the client
		// It writes the message as plain text AND sets the HTTP status code
		// http.StatusMethodNotAllowed is the number 405 — "you used the wrong method"
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return // stop here, don't run any more code in this function
	}

	// w.Write([]byte("..."))
	// Sends text back to the client (the person visiting the URL in their browser)
	// []byte(...) converts a regular string into raw bytes because Write() needs bytes
	// The _ means "I don't care about the return values" (how many bytes written, and any error)
	_, _ = w.Write([]byte("Hello World"))
}

// Home route — handles the root "/" path
func rootHandler(w http.ResponseWriter, r *http.Request) {

	// r.URL.Path → the path part of the URL (e.g., for http://localhost:8080/hello it's "/hello")
	// Because "/" matches EVERYTHING in Go's default router, we manually check
	// if the path is exactly "/" to avoid this handler catching all unknown routes
	if r.URL.Path != "/" {
		// http.StatusNotFound is the number 404 — "page not found"
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		// http.StatusMethodNotAllowed is 405 — "wrong HTTP method used"
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _ = w.Write([]byte("Welcome to my API"))
}

// queryHandler reads a "query parameter" from the URL
// Example URL: http://localhost:8080/search?name=Tanmay
// The part after "?" is called a "query string" — it carries extra info
func queryHandler(w http.ResponseWriter, r *http.Request) {

	// r.URL.Query() → parses the query string of the URL into a map
	// .Get("name")  → looks up the value for the key "name" in that map
	// If someone visits /search?name=Tanmay, this gives us "Tanmay"
	// If there's no "name" key, it returns an empty string ""
	name := r.URL.Query().Get("name")

	// If no name was given in the URL, use "Guest" as a default
	if name == "" {
		name = "Guest"
	}

	// Combine "Hello " + the name and send it back as a response
	_, _ = w.Write([]byte("Hello " + name))
}

// successHandlerJson sends a JSON response (like a REST API would)
func successHandlerJson(w http.ResponseWriter, r *http.Request) {

	// w.Header() → gives you access to the HTTP response HEADERS
	// Headers are extra metadata sent before the actual response body
	// .Set("Content-Type", "application/json") tells the browser/client:
	// "hey, what I'm sending back is JSON, not plain text or HTML"
	// IMPORTANT: You MUST set headers BEFORE calling WriteHeader or Write
	w.Header().Set("Content-Type", "application/json")

	// w.WriteHeader(statusCode)
	// Sends the HTTP status code as part of the response
	// http.StatusOK is 200 — the standard "everything went fine" code
	// Once you call this, you can NO longer change headers
	w.WriteHeader(http.StatusOK)

	// Creating a Go map to hold our response data
	// map[string]any → a dictionary where keys are strings and values can be anything
	res := map[string]any{
		"ok":       true,
		"message":  "JSON is working",
		"datetime": time.Now().UTC(), // time.Now() gives current time, .UTC() converts to UTC timezone
	}

	// json.NewEncoder(w) → creates a JSON encoder that writes directly into the response (w)
	// .Encode(res)       → converts our Go map into JSON text and sends it to the client
	// It's like saying: "turn this Go map into { "ok": true, ... } and send it out"
	// The _ discards any encoding error (in production you'd want to handle it)
	_ = json.NewEncoder(w).Encode(res)
}

// ─── HELPER FUNCTION ─────────────────────────────────────────────────────────
// writeJson is a reusable helper — instead of repeating the same 3 lines
// (set header, write status, encode JSON) in every handler, we wrap them here
// status → the HTTP status code to send (200, 400, 404, etc.)
// data   → any Go value (map, struct, etc.) that will be turned into JSON
func writeJson(w http.ResponseWriter, status int, data any) {
	// Tell the client we're sending JSON
	w.Header().Set("Content-Type", "application/json")

	// Send the HTTP status code (e.g., 200 OK, 400 Bad Request)
	w.WriteHeader(status)

	// Convert 'data' to JSON and write it to the response
	_ = json.NewEncoder(w).Encode(data)
}

// ─── STRUCT ───────────────────────────────────────────────────────────────────
// A struct is like a template / blueprint for a specific type of data
// This User struct represents the shape of JSON we EXPECT to receive in a POST request
// The backtick tags (json:"name") tell the JSON decoder what JSON key maps to which field
// e.g., incoming JSON { "name": "Tanmay" } → goes into the Name field
type User struct {
	Name    string `json:"name"`    // maps JSON key "name" to Go field Name
	Email   string `json:"email"`   // maps JSON key "email" to Go field Email
	Message string `json:"message"` // maps JSON key "message" to Go field Message
}

// decodeJson handles a POST request where the client sends JSON in the request body
// Example: POST /some-route with body → { "name": "Tanmay", "email": "t@t.com", "message": "hi" }
func decodeJson(w http.ResponseWriter, r *http.Request) {

	// Only allow POST requests — reject everything else
	if r.Method != http.MethodPost {
		// http.StatusMethodNotAllowed = 405
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"success": false,
			"message": "Method not allowed",
		})
		return
	}

	// defer r.Body.Close()
	// "defer" means: run this line WHEN the function is about to exit (no matter what)
	// r.Body is like an open pipe carrying the request's data
	// .Close() shuts it down to free up memory — always do this after reading the body!
	defer r.Body.Close()

	// Declare an empty User variable — we'll fill it with data from the request
	var req User

	// json.NewDecoder(r.Body) → creates a decoder that READS from the request body
	// Think of it like: "set up a reader pointed at the incoming data stream"
	dec := json.NewDecoder(r.Body)

	// dec.Decode(&req)
	// Reads the JSON from the request body and fills our 'req' User struct with the values
	// &req means "give me the memory address of req" so Decode can write INTO it directly
	// If the JSON is malformed or doesn't match the struct, this returns an error
	if err := dec.Decode(&req); err != nil {
		// http.StatusBadRequest = 400 — "the data you sent is wrong/broken"
		writeJson(w, http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Bad request",
		})
		return
	}

	// Basic validation — check that none of the required fields are empty strings
	if req.Name == "" || req.Email == "" || req.Message == "" {
		writeJson(w, http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Bad request",
		})
		return
	}

	// Everything looks good! Send a 200 OK response with the data echoed back
	// http.StatusOK = 200 — "everything worked perfectly"
	writeJson(w, http.StatusOK, map[string]any{
		"success":  true,
		"message":  "JSON is working",
		"datetime": time.Now().UTC(), // current time in UTC
		"data":     req,             // echo back the parsed User struct as JSON
	})
}
