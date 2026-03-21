package repositorie

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type PostgresRepositorie struct {
	db *sql.DB
}

func NewPostgresRepositorie(dsn string) (*PostgresRepositorie, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("the database is not responding: %w", err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS links (
            short_url VARCHAR(10) PRIMARY KEY,
            original_url TEXT NOT NULL UNIQUE
        )
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return &PostgresRepositorie{db: db}, nil
}

func (r *PostgresRepositorie) Save(originalLink, shortLink string) (string, error) {
	_, err := r.db.Exec(
		"INSERT INTO links (short_url, original_url) VALUES ($1, $2)",
		shortLink, originalLink,
	)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			var existingShort string
			err := r.db.QueryRow(
				"SELECT short_url FROM links WHERE original_url = $1",
				originalLink,
			).Scan(&existingShort)
			if err != nil {
				return "", fmt.Errorf("failed to find existing link: %w", err)
			}
			return existingShort, nil
		}
		return "", fmt.Errorf("failed to save link: %w", err)
	}

	return shortLink, nil
}

func (r *PostgresRepositorie) Get(shortLink string) (string, error) {
	var originalURL string
	err := r.db.QueryRow(
		"SELECT original_url FROM links WHERE short_url = $1",
		shortLink,
	).Scan(&originalURL)

	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("failed to get link: %w", err)
	}

	return originalURL, nil
}

func (r *PostgresRepositorie) Check(shortLink string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM links WHERE short_url = $1)",
		shortLink,
	).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("failed to check link: %w", err)
	}

	return exists, nil
}
