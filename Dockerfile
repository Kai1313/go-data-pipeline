# Step 1: Build Stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install git (needed for fetching dependencies)
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# We use CGO_ENABLED=0 to ensure it's a static binary for Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Step 2: Final Runtime Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .
# Copy .env if you want to bundle it (though it's better to pass env vars at runtime)
# COPY --from=builder /app/.env .

# Expose the API port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]