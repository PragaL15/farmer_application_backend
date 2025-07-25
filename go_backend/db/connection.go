package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/pashagolub/pgxmock"
)

// Database Interface with QueryRow added
type Database interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row // ✅ Added
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Ping(ctx context.Context) error
	Begin(ctx context.Context) (pgx.Tx, error)
	Close()
}

// PgxDB implements Database interface
type PgxDB struct {
	Pool *pgxpool.Pool
}

func (db *PgxDB) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return db.Pool.Query(ctx, sql, args...)
}

func (db *PgxDB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row { // ✅ Added
	return db.Pool.QueryRow(ctx, sql, args...)
}

func (db *PgxDB) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.Pool.Exec(ctx, sql, args...)
}

func (db *PgxDB) Ping(ctx context.Context) error {
	return db.Pool.Ping(ctx)
}

func (db *PgxDB) Begin(ctx context.Context) (pgx.Tx, error) {
	return db.Pool.Begin(ctx)
}

func (db *PgxDB) Close() {
	db.Pool.Close()
}

// MockDB for testing
type MockDB struct {
	Mock pgxmock.PgxPoolIface
}

func (m *MockDB) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return m.Mock.Query(ctx, sql, args...)
}

func (m *MockDB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row { // ✅ Added
	return m.Mock.QueryRow(ctx, sql, args...)
}

func (m *MockDB) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return m.Mock.Exec(ctx, sql, args...)
}

func (m *MockDB) Ping(ctx context.Context) error {
	return nil
}

func (m *MockDB) Close() {
	m.Mock.Close()
}

// Global Database Pool
var Pool Database

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found.")
	}
}

func ConnectDB() {
	LoadEnv()
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgresql://postgres:Alekya@54321@localhost:5433/brokerdb"
		log.Println("Warning: Using default database connection string.")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
		return
	}

	Pool = &PgxDB{Pool: pool}
	// code has to be modified to handle db related issues
	if err := Pool.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Database connection established.")
}

func CloseDB() {
	if Pool != nil {
		Pool.Close()
		log.Println("Database connection closed.")
	}
}
