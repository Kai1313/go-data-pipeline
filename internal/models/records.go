package models

import "time"

type RawRecord struct {
	ID        int       `json:"id"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

type ProcessedRecord struct {
	ID        	int       `json:"id"`
	DataSummary string    `json:"data_summary"`
	ProcessedAt time.Time `json:"processed_at"`
}