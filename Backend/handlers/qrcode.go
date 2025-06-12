package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"urlShortner/utils"
)

func QRCodeHandler(w http.ResponseWriter, r *http.Request) {
	var data struct { // struct to hold json data from the frontend
		Content  string `json:"content"`
		FileName string `json:"file_name"`
	}
	err := json.NewDecoder(r.Body).Decode(&data) // decoding the request body if there is error it will return it
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err = utils.GenerateQRCode(data.Content, data.FileName)
	if err != nil {
		http.Error(w, "Unable to Generate the QR Code", http.StatusBadRequest)
		return
	}

	// Get base URL from environment variable or use default
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080" // fallback for local development
	}

	res := struct { // creating a Response Object
		Message string `json:"message"`
		URL     string `json:"url"`
	}{
		Message: "QR Code generated successfully!",
		URL:     baseURL + "/qrcode/" + data.FileName,
	}
	utils.WriteJSON(w, res)
}
