package repository

import (
	"fmt"
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewRedis(config RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: config.Host + ":" + config.Port,
		Password: config.Password,
		DB: config.DB,
	})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return client, nil
}