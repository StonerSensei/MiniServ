package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"
	"urlShortner/config"
)

func GenerateRandomID() string {
	bytes := make([]byte, 5)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func SavePaste(content string, durationMinutes int) (string, error) {
	pasteID := GenerateRandomID()
	var expiresAt *time.Time
	if durationMinutes > 0 {
		exp := time.Now().UTC().Add(time.Duration(durationMinutes) * time.Minute)
		expiresAt = &exp
	}

	_, err := config.DB.Exec(
		"INSERT INTO pastes (paste_id, content, expires_at) VALUES ($1, $2, $3)",
		pasteID, content, expiresAt,
	)

	if err != nil {
		return "", err
	}

	return pasteID, nil
}

func GetPaste(pasteID string) (string, error) {
	var content string
	var expiresAt sql.NullTime

	err := config.DB.QueryRow(
		"SELECT content, expires_at FROM pastes WHERE paste_id = $1", pasteID,
	).Scan(&content, &expiresAt)

	if err != nil {
		return "", errors.New("paste not found")
	}

	if expiresAt.Valid && expiresAt.Time.Before(time.Now()) {
		return "", errors.New("paste expired")
	}

	return content, nil
}
