package postgres

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func New(storeURL string) (*Store, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", storeURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Store{db: db}, nil
}
