package service

import (
	"github.com/ZeroXKiritsu/urlshortener/internal/repository"
	"github.com/ZeroXKiritsu/urlshortener/structs"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=mocks/mock.go

type ShortURL interface {
	Create(url structs.Requests) (string, error)
	GetOriginal(shortURL string) (string, error)
	GenerateShortURL() (string, error)
}

type Service struct {
	ShortURL
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		ShortURL: NewShortURLService(r.ShortURL),
	}
}
