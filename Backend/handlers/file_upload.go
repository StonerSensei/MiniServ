package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"urlShortner/utils"
)

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // Limit file size to 10MB
	if err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file") // Get the uploaded file
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileUrl, err := utils.SaveUploadedFile(file, handler.Filename)
	if err != nil {
		http.Error(w, "File upload failed", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"url": fileUrl}
	utils.WriteJSON(w, response)

	fmt.Println("✅ File Uploaded: ", fileUrl)
}

func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := filepath.Base(r.URL.Path) // Extract file name from URL
	filePath := filepath.Join(utils.UploadDir, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File Not Found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filePath)
}
