package handlers

import (
	"log"
	"net/http"
	"urlShortner/models"
	"urlShortner/utils"
)

func IPHandler(w http.ResponseWriter, r *http.Request) {
	geo, err := utils.GetGeoInfo()
	if err != nil {
		log.Println("GeoInfo error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = models.StoreIpInfo(geo)
	if err != nil {
		log.Println("StoreIpInfo error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, geo)
}
