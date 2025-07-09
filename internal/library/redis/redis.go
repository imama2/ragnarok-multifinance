package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

func New(cfg Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}

func Ping(ctx context.Context, client *redis.Client) error {
	return client.Ping(ctx).Err()
}
