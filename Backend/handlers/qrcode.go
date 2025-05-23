package handlers

import (
	"encoding/json"
	"net/http"
	"urlShortner/utils"
)

func QRCodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "True")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var data struct { // struct to hold json data fyrom the frontend
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
	res := struct { // creating a Response Object
		Message string `json:"message"`
		URL     string `json:"url"`
	}{
		Message: "QR Code generated successfully!",
		URL:     "http://localhost:8080/qrcode/" + data.FileName,
	}
	utils.WriteJSON(w, res)
}
