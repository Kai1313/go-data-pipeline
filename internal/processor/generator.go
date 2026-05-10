package processor

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateBarcode creates: (a)2026-05-08(b)xyz(c)abcdef1234
func GenerateBarcode() string {
	datePart := time.Now().Format("2006-01-02")

	// Generate a short unique random hex string
	b := make([]byte, 8) // 8 bytes = 16 hex characters
	rand.Read(b)
	uniquePart := hex.EncodeToString(b)

	return fmt.Sprintf("(a)%s(b)xyz(c)%s", datePart, uniquePart)
}