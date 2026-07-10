package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// a simple mongo model for book
type Book struct {
	ID bson.ObjectID `bson:"_id,omitempty"`

	Title string `bson:"title"`

	Author string `bson:"author"`

	Description string `bson:"description"`

	Price int `bson:"price"`

	Quantity int `bson:"quantity"`

	CreatedAt time.Time `bson:"created_at"`

	UpdatedAt time.Time `bson:"updated_at"`
}

/*
When creating a new book, generate the ID like this:
book := Book{
	ID:          primitive.NewObjectID(),
	Title:       "Atomic Habits",
	Author:      "James Clear",
	Description: "A book about building habits",
	Price:       499,
	Quantity:    20,
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

*/
