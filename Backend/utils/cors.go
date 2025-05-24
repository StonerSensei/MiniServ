package utils

import (
	"net/http"
	"os"
	"strings"
)

func EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Get allowed origins from environment or use defaults
		allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")
		var allowedOrigins []string

		if allowedOriginsEnv != "" {
			allowedOrigins = strings.Split(allowedOriginsEnv, ",")
		} else {
			// Default allowed origins
			allowedOrigins = []string{
				"http://localhost:3000",              // Local development
				"https://your-vercel-app.vercel.app", // Replace with your actual Vercel URL
			}
		}

		// Check if origin is in allowed list
		isAllowed := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == strings.TrimSpace(allowedOrigin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				isAllowed = true
				break
			}
		}

		// Allow Vercel preview deployments
		if isAllowed || strings.Contains(origin, "vercel.app") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-Requested-With")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
