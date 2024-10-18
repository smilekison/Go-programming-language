package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database connection and creates tables
func InitDB() {
	var err error
	// Connect to SQLite database (api.db)
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Set connection pool settings (optional for SQLite)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Create tables
	createTables()
}

// createTables creates the necessary tables for the application
func createTables() {
	// SQL statement to create events table
	createEventTables := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	);`

	// Execute the SQL statement to create the table
	result, err := DB.Exec(createEventTables)
	if err != nil {
		log.Fatalf("Could not create events table: %v", err)
	}

	// Output the result of the table creation
	fmt.Println("Table creation result: ", result)
}
