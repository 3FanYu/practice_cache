package cachekit

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
)

const (
	DefaultTTL = 10 * time.Minute
)

func NewCache(addr string) Cache {
	return &cache{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

type cache struct {
	client       *redis.Client
	singleflight singleflight.Group
}

func (c *cache) Get(context context.Context, prefix, key string, container interface{}) error {
	val, err := c.client.Get(context, prefix+key).Result()
	if err != nil {
		return err
	}
	// Unmarshal the value into the container
	err = json.Unmarshal([]byte(val), container)
	if err != nil {
		return err
	}
	return nil
}

func (c *cache) Set(context context.Context, prefix string, key string, value interface{}, duration time.Duration) error {
	// Marshal the value
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	// Set the value in the cache
	return c.client.Set(context, prefix+key, val, duration).Err()
}

func (c *cache) GetByFunc(ctx context.Context, prefix, key string, container interface{}, getter OneTimeGetterFunc) error {
	cacheKey := prefix + ":" + key

	// Use singleflight to handle concurrent requests
	result, err, _ := c.singleflight.Do(cacheKey, func() (interface{}, error) {
		// Check if the value is in the cache
		val, err := c.client.Get(ctx, cacheKey).Result()
		if err == nil {
			return []byte(val), nil
		}

		// Cache miss, use the getter function to retrieve data
		data, err := getter()
		if err != nil {
			return nil, err
		}

		// Marshal the data to JSON
		json, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal value: %v", err)
		}

		// Set the value in Redis cache with a TTL of 10 minutes
		err = c.client.Set(ctx, cacheKey, json, DefaultTTL).Err()
		if err != nil {
			return nil, fmt.Errorf("failed to set value in cache: %v", err)
		}

		return json, nil
	})

	if err != nil {
		return fmt.Errorf("failed to get value: %v", err)
	}

	// Unmarshal the retrieved value into the container
	err = json.Unmarshal(result.([]byte), &container)
	if err != nil {
		return fmt.Errorf("failed to unmarshal value: %v", err)
	}

	return nil

}
