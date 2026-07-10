package handlers

import (
	create_book_dto "book_store/internals/dtos"
	"book_store/internals/models"
	"book_store/internals/repository"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/v2/bson"
)

/*
Handler Layer

Think of this as the Receptionist.

Client
   │
HTTP Request
   │
   ▼
Handler
   │
Repository
   │
MongoDB

Responsibilities:
✔ Receive HTTP requests
✔ Validate incoming JSON
✔ Create response
✔ Call Repository

Should NOT:
✘ Write MongoDB queries
✘ Contain complex business logic
*/

// BookHandler stores everything needed
// to handle Book-related requests.
type BookHandler struct {

	// Repository used to interact with MongoDB.
	repo *repository.Repo
}

// Constructor.
//
// Instead of creating BookHandler manually,
// we create it through this function.
func NewHandler(repo *repository.Repo) *BookHandler {

	return &BookHandler{

		// Save repository inside handler.
		repo: repo,
	}
}

// CreateBook handles:
//
// POST /books
func (h *BookHandler) CreateBook(c *gin.Context) {

	// req will temporarily store the JSON
	// received from the client.
	var req create_book_dto.CreateBookRequest

	// --------------------------------------------------
	// Read JSON from request body.
	//
	// Incoming JSON:
	//
	// {
	//    "title":"Atomic Habits",
	//    "author":"James Clear"
	// }
	//
	// Gin automatically fills req.
	// --------------------------------------------------
	if err := c.ShouldBindJSON(&req); err != nil {

		// JSON is invalid.
		//
		// Return HTTP 400 Bad Request.
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Current UTC time.
	//
	// Used for CreatedAt and UpdatedAt.
	now := time.Now().UTC()

	// --------------------------------------------------
	// Convert DTO into Database Model.
	//
	// DTO
	// ↓
	// Book Model
	//
	// Why?
	//
	// DTO represents client input.
	//
	// Model represents database structure.
	// --------------------------------------------------
	book := models.Book{

		// Generate MongoDB ObjectID.
		ID: bson.ObjectID{}, //{},

		// Copy values from request.
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,

		// Add timestamps.
		CreatedAt: now,
		UpdatedAt: now,
	}

	// --------------------------------------------------
	// Ask Repository to insert into MongoDB.
	//
	// Handler DOES NOT know how MongoDB works.
	//
	// Repository handles that.
	// --------------------------------------------------
	created, err := h.repo.CreateBook(c, book)

	if err != nil {

		// Database error.
		//
		// Return HTTP 500.
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	fmt.Println("Generated ID:", book.ID.Hex())
	fmt.Printf("Book: %+v\n", book)

	// --------------------------------------------------
	// Success.
	//
	// Return HTTP 201 Created.
	// --------------------------------------------------
	c.JSON(201, gin.H{

		// Send newly created book.
		"data": created,
	})
}

func (h *BookHandler) GetAllBook(c *gin.Context) {

	// --------------------------------------------------
	// Ask Repository to get all books from MongoDB.
	//
	// Handler DOES NOT know how MongoDB works.
	//
	// Repository handles that.
	// --------------------------------------------------
	books, err := h.repo.GetAllBooks(c)

	if err != nil {

		// Database error.
		//
		// Return HTTP 500.
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	// --------------------------------------------------
	// Success.
	//
	// Return HTTP 200 OK.
	// --------------------------------------------------
	c.JSON(200, gin.H{
		"data": books,
	})
}

func (h *BookHandler) GetBookById(c *gin.Context) {
	books, err := h.repo.GetAllBooks(c)
	if err != nil {

		// Database error.
		//
		// Return HTTP 500.
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"data": books,
	})
}

func (h *BookHandler) UpdateBookById(c *gin.Context) {

	idStr := c.Param("id")

	objectID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid book id",
		})
		return
	}

	// Read request body
	var req create_book_dto.UpdateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedBook, err := h.repo.UpdateBookById(
		c.Request.Context(),
		objectID,
		req,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": updatedBook,
	})
}
func (h *BookHandler) DeleteBookById(c *gin.Context) {
	idstr := c.Param("id")

	onbjId, err := bson.ObjectIDFromHex(idstr)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.repo.DeleteBookById(c.Request.Context(), onbjId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Book deleted successfully",
	})
}
