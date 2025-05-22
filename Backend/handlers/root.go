package handlers

import (
	"fmt"
	"net/http"
)

func RootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the URL Shortener API! \nAvailable Endpoints:\n"+
		"- POST /shorturl → Create a short URL\n"+
		"- GET /redirect/{id} → Redirect to original URL\n"+
		"- GET /url/{id} → Get original URL without redirect\n"+
		"- POST /generate_qr → Generate a QR Code\n"+
		"- GET /get_ip_info → Get IP geolocation info\n"+
		"- GET /get_dns?domain=example.com → Get DNS records\n"+
		"- POST /upload → Upload a file\n"+
		"- GET /files/{filename} → Retrieve a file")
}
