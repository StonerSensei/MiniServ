package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const fileExpiration = 5 * time.Minute

func CleanUpUploads() {
	for {
		time.Sleep(1 * time.Minute)
		files, err := os.ReadDir("uploads")
		if err != nil {
			log.Println(err)
			continue
		}
		now := time.Now()
		for _, file := range files {
			filePath := filepath.Join("uploads", file.Name())
			info, err := os.Stat(filePath)
			if err == nil && now.Sub(info.ModTime()) > fileExpiration {
				os.Remove(filePath)
				fmt.Println("Removed file: ", file.Name())
			}
		}
	}
}

func CleanUpQR() {
	const qrcode = "qrcode/"
	for {
		time.Sleep(1 * time.Minute)
		files, err := os.ReadDir(qrcode)
		if err != nil {
			log.Println(err)
			continue
		}

		now := time.Now()
		for _, file := range files {
			filePath := filepath.Join(qrcode, file.Name())
			info, err := os.Stat(filePath)
			if err == nil && now.Sub(info.ModTime()) > fileExpiration {
				os.Remove(filePath)
				fmt.Println("Removed file: ", file.Name())
			}
		}
	}
}
