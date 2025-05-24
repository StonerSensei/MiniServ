package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"urlShortner/config"
	"urlShortner/handlers"
	"urlShortner/utils"
)

func main() {
	// Load .env file if it exists (for local development)
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		fmt.Println("No .env file found, using environment variables")
	}

	// Initialize database
	config.InitDB()

	// Routes with CORS enabled
	http.HandleFunc("/", handlers.RootPage)
	http.HandleFunc("/generate_qr/", utils.EnableCORS(handlers.QRCodeHandler))
	http.HandleFunc("/get_ip_info/", utils.EnableCORS(handlers.IPHandler))
	http.HandleFunc("/get_dns/", utils.EnableCORS(handlers.DNSHandler))
	http.HandleFunc("/upload", utils.EnableCORS(handlers.FileUploadHandler))
	http.HandleFunc("/files/", utils.EnableCORS(handlers.ServeFileHandler))
	http.HandleFunc("/paste", utils.EnableCORS(handlers.SavePasteHandler))
	http.HandleFunc("/paste/", utils.EnableCORS(handlers.GetPasteHandler))
	http.HandleFunc("/convert", utils.EnableCORS(handlers.ConvertHandler))

	// Static file server for QR codes
	fs := http.FileServer(http.Dir("qrcode"))
	http.Handle("/qrcode/", utils.EnableCORS(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/qrcode/", fs).ServeHTTP(w, r)
	}))

	// Start cleanup goroutines
	go utils.CleanUpUploads()
	go utils.CleanUpQR()
	go utils.CleanUpPaste()

	// Get port from environment variable (Render sets this automatically)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local development
	}

	fmt.Printf("Server starting on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
