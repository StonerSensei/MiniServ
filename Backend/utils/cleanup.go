package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
	"urlShortner/config"
)

const fileExpiration = 5 * time.Minute

// Clean up expired files from Supabase storage
func CleanUpUploads() {
	for {
		time.Sleep(30 * time.Second) // Check every 30 seconds for more responsive cleanup

		log.Println("Running Supabase file cleanup job...")

		// Get expired files from database
		expiredFiles, err := GetExpiredFiles()
		if err != nil {
			log.Printf("Error getting expired files: %v", err)
			continue
		}

		if len(expiredFiles) == 0 {
			continue // No expired files
		}

		log.Printf("Found %d expired files to delete", len(expiredFiles))

		deletedCount := 0
		for _, fileName := range expiredFiles {
			// Delete from Supabase Storage
			err := DeleteFileFromSupabase(fileName)
			if err != nil {
				log.Printf("Error deleting file %s from Supabase: %v", fileName, err)
				continue
			}

			// Mark as deleted in database
			err = MarkFileAsDeleted(fileName)
			if err != nil {
				log.Printf("Error marking file %s as deleted: %v", fileName, err)
			}

			deletedCount++
			fmt.Printf("Deleted expired file: %s\n", fileName)
		}

		if deletedCount > 0 {
			log.Printf("Supabase cleanup completed: %d files deleted", deletedCount)
		}
	}
}

// Keep the QR code cleanup as is since they're still stored locally
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
				fmt.Println("Removed QR file: ", file.Name())
			}
		}
	}
}

func CleanUpPaste() {
	for {
		time.Sleep(1 * time.Minute)
		log.Println("Running paste cleanup job...")

		res, err := config.DB.Exec(`DELETE FROM pastes WHERE expires_at IS NOT NULL AND expires_at < NOW()`)
		if err != nil {
			log.Println("Error cleaning up expired pastes:", err)
			continue
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			log.Println("Error getting rows affected:", err)
			continue
		}

		if rowsAffected > 0 {
			log.Printf("Expired pastes cleaned up: %d", rowsAffected)
		}
	}
}
