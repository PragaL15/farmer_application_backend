package db

import (
    "context"
    "log"
    "github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectDB() {
    var err error
    connStr := "postgresql://postgres:pragalya123@localhost:5432/broker_retailer"

    Pool, err = pgxpool.New(context.Background(), connStr)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }


    if err := Pool.Ping(context.Background()); err != nil {
        log.Fatalf("Failed to ping the database: %v", err)
    }

    log.Println("Database connection established successfully")
}
