package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sinulingga23/sales-backend/constant"
)

var (
	C *Cache
)

func init() {
	client := redis.NewClient(&redis.Options{
		Network:  "redis-sales-backend:6379",
		Password: "",
		DB:       0,
	})
	C = NewCache(client)
}

type Cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{client: client}
}

func (c *Cache) SetValue(ctx context.Context, key string, value string, duration ...time.Duration) error {

	d := time.Hour * time.Duration(constant.DEFAULT_TIME_TO_LIVE_CACHE)
	if len(duration) > 0 {
		d = duration[0]
	}

	return c.client.Set(ctx, key, value, d).Err()
}

func (c *Cache) GetValue(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *Cache) SetGenericValue(ctx context.Context, key string, value interface{}, duration ...time.Duration) error {
	d := time.Hour * time.Duration(constant.DEFAULT_TIME_TO_LIVE_CACHE)
	if len(duration) > 0 {
		d = duration[0]
	}

	val, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, string(val), d).Err()
}

func (c *Cache) GetGenericValue(ctx context.Context, key string) (interface{}, error) {
	result, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var value interface{}
	err = json.Unmarshal([]byte(result), &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
