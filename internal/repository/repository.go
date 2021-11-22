package repository

import "log"

//go:generate go run github.com/golang/mock/mockgen -source=repository.go -destination=mocks/mock.go

type ShortURL interface {
	Create(generatedURL, original string) error
	SearchShortURL(shortURL string) (string, error)
	SearchOriginal(original string) (string, error)
}

type Repository struct {
	ShortURL
}

func NewRepository(db Database) *Repository {
	if db.DBType == Postgres {
		return &Repository{
			ShortURL: NewShortURLPostgres(db.Postgres),
		}
	} else if db.DBType == Redis {
		return &Repository{
			ShortURL: NewShortURLRedis(db.Redis),
		}
	} else {
		log.Fatal("Unknown database")
		return &Repository{}
	}
}
