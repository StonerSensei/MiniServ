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
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	pasteID, err := models.SavePaste(data.Content)
	if err != nil {
		http.Error(w, "Failed to save paste", http.StatusInternalServerError)
		return
	}
	pasteUrl := "http://localhost:8080/paste/" + pasteID
	response := map[string]string{
		"id":  pasteID,
		"url": pasteUrl,
	}
	utils.WriteJSON(w, response)
}

// GetPasteHandler handles retrieving pastes by ID
func GetPasteHandler(w http.ResponseWriter, r *http.Request) {
	pasteID := r.URL.Path[len("/paste/"):]
	content, err := models.GetPaste(pasteID)
	if err != nil {
		http.Error(w, "Paste not found", http.StatusNotFound)
		return
	}
	// ✅ Return the paste content as plain text
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, content)
}
