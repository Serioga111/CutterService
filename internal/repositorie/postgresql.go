package repositorie

import (
	"database/sql"
	"fmt"

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

	return &PostgresRepositorie{db: db}, nil
}

func (s *PostgresRepositorie) Save(originalLink, shortLink string) error {
	return nil
}

func (s *PostgresRepositorie) Get(shortLink string) (string, error) {
	return "", nil
}

func (s *PostgresRepositorie) Check(shortLink string) (bool, error) {
	return false, nil
}
