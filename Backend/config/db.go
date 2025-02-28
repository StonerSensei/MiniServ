package config

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "host=localhost port=5432 user=myuser password=mypassword dbname=mydatabase sslmode=disable"
	DB, err = sql.Open("postgres", connStr) // ✅ Assign to the global `db`
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("✅ Successfully connected to database")
}
