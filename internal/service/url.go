package service

import (
	"math/rand"
	"time"

	"github.com/ZeroXKiritsu/urlshortener/internal/repository"
	"github.com/ZeroXKiritsu/urlshortener/structs"
)

const (
	symbols     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	shortURlLen = 10
	host = "localhost:8080"
)

type ShortURLService struct {
	repository repository.ShortURL
}

func NewShortURLService(repository repository.ShortURL) *ShortURLService {
	return &ShortURLService{repository: repository}
}

func (s *ShortURLService) GetOriginal(shortURL string) (string, error) {
	original, err := s.repository.SearchShortURL(shortURL)
	if err != nil {
		return "", err
	}
	return original, nil
}

func (s *ShortURLService) Create(url structs.Requests) (string, error) {
	shortURL, err := s.repository.SearchOriginal(url.URL)
	if err != nil {
		return "", err
	}

	if shortURL != "" {
		return host + "/" + shortURL, nil
	}

	generatedURL, error := s.GenerateShortURL()
	if err != nil {
		return "", error
	}

	return host + "/" + generatedURL, s.repository.Create(generatedURL, url.URL)
}

func (s *ShortURLService) GenerateShortURL() (string, error) {
	shortURL := make([]byte, shortURlLen)

	rand.Seed(time.Now().UnixNano())

	for {
		for i := range shortURL {
			shortURL[i] = symbols[rand.Intn(len(symbols))]
		}

		result, err := s.repository.SearchShortURL(string(shortURL))
		if err != nil {
			return "", err
		}
		if result == "" {
			break
		}
	}

	return string(shortURL), nil
}