package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"urlShortner/models"
)

func GetPublicIP() (string, error) {
	resp, err := http.Get("https://api64.ipify.org/?format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, _ := io.ReadAll(resp.Body)
	return string(ip), nil
}

func GetGeoInfo() (*models.Geo, error) {
	ip, err := GetPublicIP()
	if err != nil {
		return nil, err
	}
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
