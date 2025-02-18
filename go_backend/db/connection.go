package db

import (
    "context"
    "log"
    "os"
    "strconv"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: No .env file found or unable to load.")
    }
}
func ConnectDB() {
    LoadEnv()

    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        dbURL = "postgresql://postgres:pragalya123@localhost:5432/broker_retailer"
        log.Println("Warning: Using default database connection string. Consider using environment variables.")
    }
    maxConnsStr := os.Getenv("DB_MAX_CONNS")
    minConnsStr := os.Getenv("DB_MIN_CONNS")
    maxConnLifetimeStr := os.Getenv("DB_MAX_CONN_LIFETIME")
    maxConnIdleTimeStr := os.Getenv("DB_MAX_CONN_IDLE_TIME")
    healthCheckPeriodStr := os.Getenv("DB_HEALTH_CHECK_PERIOD")

    maxConns, err := strconv.Atoi(maxConnsStr)
    if err != nil {
        log.Fatalf("Invalid value for DB_MAX_CONNS: %v", err)
    }
    minConns, err := strconv.Atoi(minConnsStr)
    if err != nil {
        log.Fatalf("Invalid value for DB_MIN_CONNS: %v", err)
    }
    maxConnLifetime, err := strconv.Atoi(maxConnLifetimeStr)
    if err != nil {
        log.Fatalf("Invalid value for DB_MAX_CONN_LIFETIME: %v", err)
    }
    maxConnIdleTime, err := strconv.Atoi(maxConnIdleTimeStr)
    if err != nil {
        log.Fatalf("Invalid value for DB_MAX_CONN_IDLE_TIME: %v", err)
    }
    healthCheckPeriod, err := strconv.Atoi(healthCheckPeriodStr)
    if err != nil {
        log.Fatalf("Invalid value for DB_HEALTH_CHECK_PERIOD: %v", err)
    }
    config, err := pgxpool.ParseConfig(dbURL)
    if err != nil {
        log.Fatalf("Unable to parse database URL: %v", err)
    }
    config.MaxConns = int32(maxConns)
    config.MinConns = int32(minConns)
    config.MaxConnLifetime = time.Duration(maxConnLifetime) * time.Minute
    config.MaxConnIdleTime = time.Duration(maxConnIdleTime) * time.Minute
    config.HealthCheckPeriod = time.Duration(healthCheckPeriod) * time.Minute
    Pool, err = pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        log.Fatalf("Failed to create connection pool: %v", err)
    }
    if err := Pool.Ping(context.Background()); err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
    log.Println("Database connection pool established successfully!")
}
func CloseDB() {
    if Pool != nil {
        Pool.Close()
        log.Println("Database connection pool closed.")
    }
}
