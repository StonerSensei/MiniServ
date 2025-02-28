package utils

import (
	"github.com/skip2/go-qrcode"
	"os"
)

func GenerateQRCode(data, filename string) error {
	dir := "qrcode"
	// os.Stat(dir) checks if qrcode  folder exist if not then it creates it
	// err := os.MkdirAll(dir, 0755)  creates directory if not present with permission of 0755 means read and execute
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755) // creating a directory if not exist
		if err != nil {
			return err
		}
	}
	// Creates Full file path directory/filename
	filePath := dir + "/" + filename
	// Generates QR code and save it to file path
	return qrcode.WriteFile(data, qrcode.Medium, 256, filePath)
}
