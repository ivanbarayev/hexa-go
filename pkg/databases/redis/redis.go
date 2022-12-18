package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"main/config"
)

var (
	db  *redis.Client
	Ctx context.Context
	err error
)

// NewRedisClient Returns new redis client
func NewRedisClient(cfg *config.Config) *redis.Client {
	println("Driver Redis Initialized")

	redisHost := fmt.Sprintf("%s:%d", cfg.Redis.HOST, cfg.Redis.PORT)

	db = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: cfg.Redis.PASS,
	})

	return db
}
