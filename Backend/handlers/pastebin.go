package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"urlShortner/models"
	"urlShortner/utils"
)

// SavePasteHandler handles saving new pastes
func SavePasteHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Content          string `json:"content"`
		ExpiresInMinutes int    `json:"expires_in_minutes"` // optional
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	pasteID, err := models.SavePaste(data.Content, data.ExpiresInMinutes)
	if err != nil {
		http.Error(w, "Failed to save paste", http.StatusInternalServerError)
		return
	}

	pasteURL := "http://localhost:8080/paste/" + pasteID
	utils.WriteJSON(w, map[string]string{
		"id":  pasteID,
		"url": pasteURL,
	})
}

// GetPasteHandler handles retrieving pastes by ID
func GetPasteHandler(w http.ResponseWriter, r *http.Request) {
	pasteID := r.URL.Path[len("/paste/"):]
	content, err := models.GetPaste(pasteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, content)
}
