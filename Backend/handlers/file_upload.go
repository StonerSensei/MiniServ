package handlers

import (
	"fmt"
	"net/http"
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

	fileUrl, err := utils.SaveUploadedFileToSupabase(file, handler.Filename)
	if err != nil {
		http.Error(w, "File upload failed", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"url": fileUrl}
	utils.WriteJSON(w, response)

	fmt.Println("✅ File Uploaded to Supabase: ", fileUrl)
}

func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	// Since files are now served directly from Supabase,
	// this handler can redirect to the Supabase URL or proxy the request
	fileName := r.URL.Path[len("/files/"):]

	supabaseUrl := utils.GetSupabaseFileURL(fileName)
	if supabaseUrl == "" {
		http.Error(w, "File Not Found", http.StatusNotFound)
		return
	}

	// Redirect to Supabase URL
	http.Redirect(w, r, supabaseUrl, http.StatusTemporaryRedirect)
}
