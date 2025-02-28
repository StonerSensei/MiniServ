package handlers

import (
	"encoding/json"
	"net/http"
	"urlShortner/models"
	"urlShortner/utils"
)

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var data struct {
		Url      string `json:"url"`
		CustomID string `json:"custom_id,omitempty"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var shortUrl string
	if data.CustomID != "" {
		err := models.CreateCustomUrl(data.Url, data.CustomID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		shortUrl = data.CustomID
	} else {
		shortUrl = models.CreateUrl(data.Url)
	}

	res := struct {
		NewUrl string `json:"new_url"`
	}{NewUrl: shortUrl}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetOriginalUrlHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[len("/url/"):] // Extract short URL ID

	originalUrl, err := models.GetOriginalUrl(shortUrl)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	response := map[string]string{"original_url": originalUrl}
	utils.WriteJSON(w, response)
}

func RedirectToUrl(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	originalUrl, err := models.GetOriginalUrl(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalUrl, http.StatusFound)
}
