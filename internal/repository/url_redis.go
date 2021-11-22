package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type ShortURLRedis struct {
	db *redis.Client
}

func NewShortURLRedis(db *redis.Client) *ShortURLRedis {
	return &ShortURLRedis{db: db}
}

func (r *ShortURLRedis) SearchShortURL(shortURL string) (string, error) {
	original, err := r.db.Get(context.Background(), shortURL).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return original, nil
}

func (r *ShortURLRedis) SearchOriginal(original string) (string, error) {
	shortURL, err := r.db.Get(context.Background(), original).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return shortURL, nil
}

func (r *ShortURLRedis) Create(generatedURL, original string) error {
	if err := r.db.Set(context.Background(), generatedURL, original, 0).Err(); err != nil {
		return fmt.Errorf("%v", err)
	}
	if err := r.db.Set(context.Background(), original, generatedURL, 0).Err(); err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
