package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client *redis.Client
}

var ctx = context.Background()

func Connect() *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})

	log.Println("Redis server is running")

	return &Cache{client: rdb}
}

func (c *Cache) Get(key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Cache) Set(key string, value string, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *Cache) Close() error {
	return c.client.Close()
}
