package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

func ListingsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Print an informational message in the terminal/logs.
		// This helps us know that someone is trying to fetch listings.
		slog.Info("Fetching listings from the database...")

		// db.Query() sends the SQL query to the database.
		// It returns:
		// 1. rows -> the results returned by the database
		// 2. err  -> an error, if something went wrong
		//
		// Here we're asking:
		// "Give me the latest 100 listings."
		rows, err := db.Query(`
			SELECT id, title, description, price, city, created_at
			FROM listings
			ORDER BY created_at DESC
			LIMIT 100
		`)

		// Check whether the database query failed.
		if err != nil {

			// Write the error into our application logs.
			// %v means "print the actual error value".
			slog.Error("query %v", err)

			// Tell the client that something went wrong on our server.
			// We don't expose the actual database error to the client
			// because that could leak internal information.
			http.Error(
				w,
				"internal server error",
				http.StatusInternalServerError,
			)

			// Stop executing this handler.
			// We don't want to continue because we don't have any rows.
			return
		}

		// rows represents the results returned by the database.
		//
		// Once we're finished with the rows, we need to close them.
		//
		// defer means:
		// "Run rows.Close() when this function is finished."
		//
		// So no matter how we leave this handler,
		// Go will try to clean up the database resources.
		defer rows.Close()

		// Create an empty slice that will eventually contain
		// all the listings we get from the database.
		//
		// Think of this as an empty box:
		//
		// listings = []
		//
		// As we read each database row, we'll put a listing inside it.
		listings := []listing{}

		// rows.Next() moves to the next row returned by the database.
		//
		// Imagine the database returned:
		//
		// Row 1
		// Row 2
		// Row 3
		//
		// rows.Next() moves through them one by one.
		//
		// It returns true while there is another row to read.
		for rows.Next() {

			// Create an empty listing struct.
			//
			// This will temporarily hold the data
			// from the current database row.
			var l listing

			// rows.Scan() takes the values from the current database row
			// and puts them into our Go variables.
			//
			// The order here MUST match the order in SELECT:
			//
			// SELECT
			// id,
			// title,
			// description,
			// price,
			// city,
			// created_at
			//
			// Therefore Scan receives:
			//
			// id          -> l.ID
			// title       -> l.Title
			// description -> l.Description
			// price       -> l.Price
			// city        -> l.City
			// created_at  -> l.CreatedAt
			if err := rows.Scan(
				&l.ID,
				&l.Title,
				&l.Description,
				&l.Price,
				&l.City,
				&l.CreatedAt,
			); err != nil {

				// Log the error so we can debug it later.
				log.Printf("Rows.scan %v", err)

				// Tell the client that something went wrong.
				http.Error(
					w,
					"internal server error",
					http.StatusInternalServerError,
				)

				// IMPORTANT:
				// We should return here.
				// Otherwise the code will continue and may send
				// a successful response after already sending an error.
				return
			}

			// Add the listing we just created to our listings slice.
			//
			// Before:
			// listings = []
			//
			// After first row:
			// listings = [listing1]
			//
			// After second row:
			// listings = [listing1, listing2]
			listings = append(listings, l)
		}

		// rows.Err() checks whether something went wrong
		// while we were going through the rows.
		//
		// This is different from the error we checked earlier.
		//
		// db.Query() could succeed,
		// but an error could happen later while reading the rows.
		if err := rows.Err(); err != nil {

			// Log the database row-reading error.
			log.Printf("rows.err %v", err)

			// Tell the client that something went wrong.
			http.Error(
				w,
				"internal server error",
				http.StatusInternalServerError,
			)

			// Stop here because we couldn't successfully read
			// all the database results.
			return
		}

		// We're returning JSON, so we set Content-Type to application/json.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Encode() writes that JSON directly into the HTTP response.
		json.NewEncoder(w).Encode(listings)
	}
}


// DELETE
func DeleteListingsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// take the id from the url
		id := r.PathValue("id")
		fmt.Printf("INFO The id got is: %s", id)

		// write the sql cmd to delete
		_, err := db.Exec(`
		DELETE FROM listings WHERE id=$1
		`, id)
		if err != nil {
			log.Printf("Delete: %w", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		// rows, err := result.RowsAffected()

		w.WriteHeader(http.StatusNoContent)

	}
}
