package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgres(config PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
