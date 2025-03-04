package models

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"urlShortner/config"
)

// GenerateRandomID creates a unique paste ID
func GenerateRandomID() string {
	bytes := make([]byte, 5) // 10-character hex string (5 bytes = 10 hex chars)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// SavePaste stores a new paste in the database
func SavePaste(content string) (string, error) {
	pasteID := GenerateRandomID()

	_, err := config.DB.Exec("INSERT INTO pastes (paste_id, content) VALUES ($1, $2)", pasteID, content)
	if err != nil {
		return "", err
	}

	return pasteID, nil
}

// GetPaste retrieves a paste by its ID
func GetPaste(pasteID string) (string, error) {
	var content string
	err := config.DB.QueryRow("SELECT content FROM pastes WHERE paste_id = $1", pasteID).Scan(&content)
	if err != nil {
		return "", errors.New("paste not found")
	}
	return content, nil
}
