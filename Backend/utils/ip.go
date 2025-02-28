package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"urlShortner/models"
)

func GetPublicIP() string {
	resp, err := http.Get("https://api64.ipify.org/?format=text")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer resp.Body.Close()

	ip, _ := io.ReadAll(resp.Body)
	return string(ip)
}

func GetGeoInfo() (*models.Geo, error) {
	ip := GetPublicIP()
	url := "http://ip-api.com/json/" + ip

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var geo models.Geo
	err = json.NewDecoder(resp.Body).Decode(&geo)
	if err != nil {
		return nil, err
	}
	return &geo, nil
}
