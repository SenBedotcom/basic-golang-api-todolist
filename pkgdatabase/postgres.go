package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(cfg Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// InitSchema initializes the database schema
func InitSchema(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		completed BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to initialize schema: %w", err)
	}

	return nil
}
