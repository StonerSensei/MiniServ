package handlers

import (
	"net/http"
	"urlShortner/models"
	"urlShortner/utils"
)

func DNSHandler(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	if domain == "" {
		http.Error(w, "Domain is required", http.StatusBadRequest)
		return
	}

	dnsInfo, err := utils.DNSLookup(domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = models.StoreDnsInfo(dnsInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, dnsInfo)
}
