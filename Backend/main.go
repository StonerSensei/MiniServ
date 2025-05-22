package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"urlShortner/config"
	"urlShortner/handlers"
	"urlShortner/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitDB()
	http.HandleFunc("/", handlers.RootPage)

	http.HandleFunc("/generate_qr/", utils.EnableCORS(handlers.QRCodeHandler))

	http.HandleFunc("/get_ip_info/", utils.EnableCORS(handlers.IPHandler))
	http.HandleFunc("/get_dns/", utils.EnableCORS(handlers.DNSHandler))

	http.HandleFunc("/upload", utils.EnableCORS(handlers.FileUploadHandler))
	http.HandleFunc("/files/", utils.EnableCORS(handlers.ServeFileHandler))

	http.HandleFunc("/paste", utils.EnableCORS(handlers.SavePasteHandler))
	http.HandleFunc("/paste/", utils.EnableCORS(handlers.GetPasteHandler))
	http.HandleFunc("/convert", utils.EnableCORS(handlers.ConvertHandler))
	// FileServer(http.Dir("qrcode")) --> Creates a file server from qr code directory
	// http.Handle("/qrcode/", http.StripPrefix("/qrcode/", fs)) It maps http://localhost:8080/qrcode/ to server files
	// Change reflecting or not
	fs := http.FileServer(http.Dir("qrcode"))
	http.Handle("/qrcode/", http.StripPrefix("/qrcode/", fs))
	go utils.CleanUpUploads()
	go utils.CleanUpQR()
	go utils.CleanUpPaste()
	// Starting a HTTTP Server on Port 8080
	//corsHandler := enableCORS(r)
	fmt.Println("Server Starting on Port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
