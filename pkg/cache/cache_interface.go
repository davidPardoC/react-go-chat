package cache

import "time"

type ICacheService interface {
	Get(key string) error
	Set(key string, value any, ttl time.Duration) error
	Remove(key string) error
}
