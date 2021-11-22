package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ShortURLPostgres struct {
	db *sqlx.DB
}

func NewShortURLPostgres(db *sqlx.DB) *ShortURLPostgres {
	return &ShortURLPostgres{db: db}
}

func (r *ShortURLPostgres) SearchShortURL(shortURL string) (string, error) {
	var original string
	query := fmt.Sprint("SELECT original FROM urls WHERE short_url = $1")
	row := r.db.QueryRow(query, shortURL)
	err := row.Scan(&original)

	if err == sql.ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return original, nil
}

func (r *ShortURLPostgres) SearchOriginal(original string) (string, error) {
	var shortURL string
	query := fmt.Sprint("SELECT short_url FROM urls WHERE original = $1")
	row := r.db.QueryRow(query, original)
	err := row.Scan(&shortURL)

	if err == sql.ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return shortURL, nil
}

func (r *ShortURLPostgres) Create(generatedURL, original string) error {
	var id int
	query := fmt.Sprint("INSERT INTO urls (short_url, original) VALUES ($1, $2) RETURNING ID")
	row := r.db.QueryRow(query, generatedURL, original)
	err := row.Scan(&id)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return err
}
