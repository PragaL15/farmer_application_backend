package db

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// ConnectDB establishes a connection to the PostgreSQL database
func ConnectDB() {
    var err error
    connStr := "user=postgres password=pragalya123 dbname=broker_retailer host=localhost sslmode=disable"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("Failed to ping the database: %v", err)
    }

    log.Println("Database connection established successfully")
}
