package cache

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCacheService struct {
	r *redis.Client
}

func NewRedisCacheService(r *redis.Client) *RedisCacheService {
	return &RedisCacheService{r: r}
}

func (*RedisCacheService) Get(key string) error {
	return nil
}

func (*RedisCacheService) Set(key string, value any, ttl time.Duration) error {
	return nil
}

func (*RedisCacheService) Remove(key string) error {
	return nil
}
