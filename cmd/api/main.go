package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go-data-pipeline/internal/db"
	"go-data-pipeline/internal/processor"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type GenerateRequest struct {
	Amount int `json:"amount"`
}

func main() {
	// Connect to the database
	pool, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	log.Println("welcome")

	r := chi.NewRouter()

	// Define the Generate route
	r.Post("/generate", func(w http.ResponseWriter, r *http.Request) {
		var req GenerateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		start := time.Now()
		log.Printf("start at %v", start)
		log.Println("Generating records...")

		// Settings for batching
		batchSize := 5000
		totalCreated := 0

		for totalCreated < req.Amount {
			currentBatch := batchSize
			if req.Amount-totalCreated < batchSize {
				currentBatch = req.Amount - totalCreated
			}

			// Prepare the batch rows
			rows := [][]interface{}{}
			for i := 0; i < currentBatch; i++ {
				rows = append(rows, []interface{}{processor.GenerateBarcode()})
			}

			// Bulk insert using PostgreSQL's COPY protocol
			_, err := pool.CopyFrom(
				context.Background(),
				pgx.Identifier{"staging_table"},
				[]string{"barcode"},
				pgx.CopyFromRows(rows),
			)

			if err != nil {
				log.Printf("Batch insert failed: %v", err)
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
			totalCreated += currentBatch
		}

		duration := time.Since(start)
		msg := fmt.Sprintf("Generated %d records in %v", req.Amount, duration)
		log.Println(msg)
		w.Write([]byte(msg))
	})

	log.Println("API server running on :8080")
	http.ListenAndServe(":8080", r)
}