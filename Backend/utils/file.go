package utils

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
)

const UploadDir = "uploads/"

func SaveUploadedFile(file io.Reader, originalFileName string) (string, error) {
	ext := filepath.Ext(originalFileName)
	randomName := fmt.Sprintf("%d%s", rand.Intn(1000000), ext)
	filePath := filepath.Join(UploadDir, randomName)

	if _, err := os.Stat(UploadDir); os.IsNotExist(err) {
		err := os.MkdirAll(UploadDir, 0755) // Create the directory if it doesn't exist
		if err != nil {
			return "", err
		}
	}

	saveFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer saveFile.Close()

	_, err = io.Copy(saveFile, file)
	if err != nil {
		return "", err
	}

	fileUrl := fmt.Sprintf("http://localhost:8080/files/%s", randomName)
	return fileUrl, nil
}
