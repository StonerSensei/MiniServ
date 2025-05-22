package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"urlShortner/config"
)

type SupabaseConfig struct {
	URL        string
	AnonKey    string
	ServiceKey string
	BucketName string
}

var supabaseConfig *SupabaseConfig

func InitSupabase() {
	supabaseConfig = &SupabaseConfig{
		URL:        os.Getenv("SUPABASE_URL"),
		AnonKey:    os.Getenv("SUPABASE_ANON_KEY"),
		ServiceKey: os.Getenv("SUPABASE_SERVICE_KEY"),
		BucketName: os.Getenv("SUPABASE_BUCKET_NAME"),
	}

	if supabaseConfig.BucketName == "" {
		supabaseConfig.BucketName = "uploads" // default bucket name
	}
}

func SaveUploadedFileToSupabase(file io.Reader, originalFileName string) (string, error) {
	if supabaseConfig == nil {
		InitSupabase()
	}

	// Generate random filename
	ext := filepath.Ext(originalFileName)
	randomName := fmt.Sprintf("%d%s", rand.Intn(1000000), ext)

	// Read file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file content: %v", err)
	}

	// Upload to Supabase Storage
	uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s",
		supabaseConfig.URL, supabaseConfig.BucketName, randomName)

	req, err := http.NewRequest("POST", uploadURL, bytes.NewReader(fileContent))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+supabaseConfig.ServiceKey)
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Store file metadata in database for cleanup tracking
	err = storeFileMetadataInDB(randomName, originalFileName, int64(len(fileContent)))
	if err != nil {
		// If we can't store metadata, delete the uploaded file to prevent orphans
		DeleteFileFromSupabase(randomName)
		return "", fmt.Errorf("failed to store file metadata: %v", err)
	}

	// Return public URL
	publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s",
		supabaseConfig.URL, supabaseConfig.BucketName, randomName)

	return publicURL, nil
}

// Store file metadata in database for cleanup tracking
func storeFileMetadataInDB(fileName, originalName string, fileSize int64) error {
	query := `
        INSERT INTO uploaded_files (filename, original_name, file_size, uploaded_at, expires_at) 
        VALUES ($1, $2, $3, NOW(), NOW() + INTERVAL '5 minutes')`

	_, err := config.DB.Exec(query, fileName, originalName, fileSize)
	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}

	fmt.Printf("✅ File metadata stored: %s (expires in 5 minutes)\n", fileName)
	return nil
}

// Get expired files from database
func GetExpiredFiles() ([]string, error) {
	query := `
        SELECT filename 
        FROM uploaded_files 
        WHERE expires_at < NOW() AND is_deleted = FALSE`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("database query error: %v", err)
	}
	defer rows.Close()

	var expiredFiles []string
	for rows.Next() {
		var filename string
		if err := rows.Scan(&filename); err != nil {
			continue
		}
		expiredFiles = append(expiredFiles, filename)
	}

	return expiredFiles, nil
}

// Mark file as deleted in database
func MarkFileAsDeleted(fileName string) error {
	query := `UPDATE uploaded_files SET is_deleted = TRUE WHERE filename = $1`

	_, err := config.DB.Exec(query, fileName)
	if err != nil {
		return fmt.Errorf("database update error: %v", err)
	}

	return nil
}

func GetSupabaseFileURL(fileName string) string {
	if supabaseConfig == nil {
		InitSupabase()
	}

	return fmt.Sprintf("%s/storage/v1/object/public/%s/%s",
		supabaseConfig.URL, supabaseConfig.BucketName, fileName)
}

func DeleteFileFromSupabase(fileName string) error {
	if supabaseConfig == nil {
		InitSupabase()
	}

	deleteURL := fmt.Sprintf("%s/storage/v1/object/%s/%s",
		supabaseConfig.URL, supabaseConfig.BucketName, fileName)

	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+supabaseConfig.ServiceKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("delete failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// List files in bucket (useful for debugging)
func ListSupabaseFiles() ([]string, error) {
	if supabaseConfig == nil {
		InitSupabase()
	}

	listURL := fmt.Sprintf("%s/storage/v1/object/list/%s",
		supabaseConfig.URL, supabaseConfig.BucketName)

	req, err := http.NewRequest("POST", listURL, bytes.NewBuffer([]byte(`{"limit": 1000}`)))
	if err != nil {
		return nil, fmt.Errorf("failed to create list request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+supabaseConfig.ServiceKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list request failed with status %d", resp.StatusCode)
	}

	var files []struct {
		Name      string    `json:"name"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name)
	}

	return fileNames, nil
}
