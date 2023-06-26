package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var ctx = context.Background()

type RedisCache struct {
	cache  *InMemoryCache[string, string]
	client *redis.Client
}

func NewRedisCache[K comparable](timeout int64) *RedisCache {
	return &RedisCache{
		NewInMemoryCache[string, string](timeout, true, 0),
		redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		}),
	}
}

func (c *RedisCache) Get(key string) (string, bool) {
	value, ok := c.cache.Get(key)
	if !ok {
		v, _ := c.client.Get(ctx, key).Result()
		c.cache.Set(key, v)
		return v, true
	}
	return value, ok
}

func (c *RedisCache) Set(key string, value string) {
	c.cache.Set(key, value)
	c.client.Set(ctx, key, value, 0)
}
