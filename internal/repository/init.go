package repository

import (
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

const (
	Unknown = iota
	Postgres
	Redis
)

type Database struct {
	Postgres *sqlx.DB
	Redis    *redis.Client
	DBType   int
}

func NewDatabase(dbType string) Database {
	if dbType == "postgres" {
		db, err := NewPostgres(PostgresConfig{
			Host:     "localhost",
			Port:     "5432",
			Username: "docker",
			Password: "docker",
			DBName:   "docker",
			SSLMode:  "disable",
		})
		if err != nil {
			log.Fatal(err)
		}
		return Database{
			Postgres: db,
			DBType:   Postgres,
		}
	} else if dbType == "redis" {
		db, err := NewRedis(RedisConfig{
			Host:     "localhost",
			Port:     "6379",
			Password: "",
			DB:       0,
		})
		if err != nil {
			log.Fatal(err)
		}
		return Database{
			Redis:  db,
			DBType: Redis,
		}
	} else {
		log.Fatal("Unknown database")
		return Database{
			DBType: Unknown,
		}
	}
}
