package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("✅ Successfully connected to database")

	createTables()
}

func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS dns_lookups (
			domain TEXT PRIMARY KEY,
			ip_addresses TEXT[],
			name_servers TEXT[],
			mail_servers TEXT[],
			cname TEXT,
			txt_records TEXT[],
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS ip_geolocation (
			ip TEXT PRIMARY KEY,
			country TEXT,
			city TEXT,
			isp TEXT,
			org TEXT,
			latitude DOUBLE PRECISION,
			longitude DOUBLE PRECISION
		);`,

		`CREATE TABLE IF NOT EXISTS pastes (
			paste_id TEXT PRIMARY KEY,
			content TEXT
		);`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatalf("Failed to create table: %v", err)
		}
	}
}
