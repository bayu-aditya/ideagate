package infrastructure

import (
	"github.com/bayu-aditya/ideagate/backend/internal/proxy/config"
	"github.com/redis/go-redis/v9"
)

type Infrastructure struct {
	Redis redis.UniversalClient
}

func NewInfrastructure(cfg *config.Config) Infrastructure {
	// Initialize redis
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    cfg.Redis.Addrs,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	return Infrastructure{
		Redis: redisClient,
	}
}
