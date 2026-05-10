# Go Data Pipeline 🚀

A high-performance Go-based service designed to generate and ingest large-scale serialization data into PostgreSQL at speed. 

This project demonstrates how to handle millions of records efficiently using Go's concurrency model and the PostgreSQL **Binary Lead Protocol (COPY)**.

## 🛠 Features

- **High-Speed Ingestion:** Capable of generating and storing 1,000,000 records in ~23 seconds (~43,000 records/sec).
- **Collision Resistance:** Advanced barcode generation logic with built-in retry mechanisms and "Circuit Breaker" batch logic.
- **Memory Efficient:** Uses streaming JSON decoding and buffered batching (5,000 records per batch) to maintain a low RAM footprint.
- **PostgreSQL Optimized:** Leverages `pgx` and the `COPY` protocol to bypass standard SQL overhead.

## 📋 Barcode Logic
Generated barcodes follow the pattern:
`(a)YYYY-MM-DD(b)xyz(c)<unique_hex>`

## 🚀 Getting Started

### 1. Prerequisites
- Go 1.21+
- PostgreSQL (Laragon or native)

### 2. Installation
```bash
git clone [https://github.com/Kai1313/go-data-pipeline.git](https://github.com/Kai1313/go-data-pipeline.git)
cd go-data-pipeline
go mod tidy