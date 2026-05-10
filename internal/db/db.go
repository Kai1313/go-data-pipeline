package db

import (
	"context"
	"os"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func Connect() (*pgxpool.Pool, error) {
    // 1. Force load env here just in case main.go missed it
    godotenv.Load() 

    connStr := os.Getenv("DATABASE_URL")
    fmt.Printf("--- Connecting to: %s ---\n", connStr) // Add this line!

    config, err := pgxpool.ParseConfig(connStr)
    if err != nil {
        return nil, err
    }
    
    return pgxpool.NewWithConfig(context.Background(), config)
}