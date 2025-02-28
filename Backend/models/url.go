package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"urlShortner/config"
)

func GenerateShortUrl(OriginalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl)) // converting the original url is smaller byte slice
	data := hasher.Sum(nil)           // storing slice in data
	hash := hex.EncodeToString(data)  // encoding the data into string
	return hash[:8]
}

func CreateUrl(originalUrl string) string {
	shortUrl := GenerateShortUrl(originalUrl)
	_, err := config.DB.Exec("INSERT INTO urls(short_url, original_url) VALUES($1, $2) ON CONFLICT (short_url) DO NOTHING", shortUrl, originalUrl)
	if err != nil {
		log.Fatal(err)
	}
	return shortUrl
}

// function of getting the original url
func GetOriginalUrl(shortUrl string) (string, error) {
	var originalUrl string
	err := config.DB.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", shortUrl).Scan(&originalUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}
	return originalUrl, nil
}

func CreateCustomUrl(originalUrl, customUrl string) error {
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM urls WHERE short_url = $1)", customUrl).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("custom URL already exists")
	}

	_, err = config.DB.Exec("INSERT INTO urls (short_url, original_url) VALUES ($1, $2)", customUrl, originalUrl)
	return err
}
