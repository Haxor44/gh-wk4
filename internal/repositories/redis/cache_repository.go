package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	client *redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{client: client}
}

func (r *CacheRepository) Set(key string, value interface{}, expiration time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		return nil
	}
	return r.client.Set(context.Background(), key, p, expiration).Err()

}

func (r *CacheRepository) Get(key string, dest interface{}) error {
	p, err := r.client.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(p, dest)
}

func (r *CacheRepository) SetSlice(key string, value interface{}, expiration time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(context.Background(), key, jsonData, expiration).Err()
}

func (r *CacheRepository) GetSlice(key string, dest interface{}) error {
	jsonData, err := r.client.Get(context.Background(), key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, dest)
}
