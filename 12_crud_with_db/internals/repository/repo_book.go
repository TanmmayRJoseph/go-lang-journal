package repository

import (
	create_book_dto "book_store/internals/dtos"
	"book_store/internals/models"
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

/*
Repository Layer

Think of this file as the "Database Employee."

Controller  --->  Service  ---> Repository ---> MongoDB

Its ONLY responsibility is:
- Insert data
- Read data
- Update data
- Delete data

It should NEVER:
- Return HTTP responses
- Know about Gin
- Parse JSON
- Handle authentication
*/

// Repo represents our repository.
//
// It stores the MongoDB collection that this repository will work with.
//
// Since this repository is for books,
// we store a pointer to the "books" collection.
type Repo struct {

	// Pointer to MongoDB collection.
	//
	// Think of Collection as a table in SQL.
	//
	// books collection
	// -------------------------
	// _id
	// title
	// author
	// price
	//
	coll *mongo.Collection
}

// Constructor for Repository.
//
// Instead of creating Repo manually everywhere,
// we create it once using this function.
//
// db represents the connected MongoDB database.
func NewRepo(db *mongo.Database) *Repo {

	// Return a pointer to Repo.
	return &Repo{

		// Tell MongoDB that this repository
		// will work with the "books" collection.
		//
		// If collection doesn't exist,
		// MongoDB creates it automatically
		// when first document is inserted.
		coll: db.Collection("books"),
	}
}

// CreateBook inserts one book into MongoDB.
func (r *Repo) CreateBook(ctx context.Context, book models.Book) (models.Book, error) {

	// ----------------------------------------------------
	// Create a child context with a timeout.
	//
	// If MongoDB takes more than 5 seconds,
	// operation automatically gets cancelled.
	//
	// Why?
	//
	// Imagine MongoDB server hangs forever.
	// Without timeout your API would also hang forever.
	//
	// This protects your application.
	// ----------------------------------------------------
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	// VERY IMPORTANT.
	//
	// defer means:
	// "Run this function just before exiting."
	//
	// It releases resources allocated by context.
	defer cancel()

	// ----------------------------------------------------
	// InsertOne()
	//
	// First argument:
	// Context (contains timeout)
	//
	// Second argument:
	// Book struct that will be inserted.
	//
	// MongoDB converts Go struct
	// into BSON automatically.
	// ----------------------------------------------------
	result, err := r.coll.InsertOne(opCtx, book)
	// If insertion failed...
	if err != nil {

		// Log the error.
		slog.Error(
			"Error inserting book",
			"error",
			err,
		)

		// Return empty Book + error.
		return models.Book{}, err
	}

	// ----------------------------------------------------
	// If insertion succeeds...
	//
	// MongoDB automatically creates an _id
	// if we didn't provide one.
	//
	// InsertOne returns:
	//
	// result.InsertedID
	//
	// whose type is interface{}.
	//
	// We know it is actually an ObjectID,
	// so we perform a type assertion.
	// ----------------------------------------------------
	_ = result

	// Return the book back to Service layer.
	//
	// Now book contains the generated MongoDB ID.
	return book, nil
}

func (r *Repo) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{}
	cursor, err := r.coll.Find(opCtx, filter)
	if err != nil {
		return nil, err
	}
	// cursor must be closed
	defer cursor.Close(opCtx)

	var books []models.Book

	if err := cursor.All(opCtx, &books); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *Repo) GetBookById(ctx context.Context, id bson.ObjectID) (models.Book, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	var book models.Book
	err := r.coll.FindOne(opCtx, filter).Decode(&book)
	return book, err
}

func (r *Repo) UpdateBookById(
	ctx context.Context,
	id bson.ObjectID,
	req create_book_dto.UpdateBookRequest,
) (models.Book, error) {

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": bson.M{
			"title":       req.Title,
			"author":      req.Author,
			"description": req.Description,
			"price":       req.Price,
			"quantity":    req.Quantity,
			"updated_at":  time.Now().UTC(),
		},
	}

	_, err := r.coll.UpdateOne(opCtx, filter, update)
	if err != nil {
		return models.Book{}, err
	}

	var updatedBook models.Book

	err = r.coll.FindOne(opCtx, filter).Decode(&updatedBook)
	if err != nil {
		return models.Book{}, err
	}

	return updatedBook, nil
}

func (r *Repo) DeleteBookById(ctx context.Context, id bson.ObjectID) error {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	_, err := r.coll.DeleteOne(opCtx, filter)
	return err
}
