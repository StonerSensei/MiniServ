package models

import (
	"urlShortner/config"
)

type Geo struct {
	Query     string  `json:"query"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
	ISP       string  `json:"isp"`
	Org       string  `json:"org"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func StoreIpInfo(geo *Geo) error {
	_, err := config.DB.Exec(
		`INSERT INTO ip_geolocation (ip, country, city, isp, org, latitude, longitude) 
		 VALUES ($1, $2, $3, $4, $5, $6, $7) 
		 ON CONFLICT(ip) DO NOTHING`,
		geo.Query, geo.Country, geo.City, geo.ISP, geo.Org, geo.Latitude, geo.Longitude,
	)
	return err
}
