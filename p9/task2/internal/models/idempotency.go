package models

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type CachedResponse struct {
	StatusCode int    `redis:"status_code"`
	Body       []byte `redis:"body"`
	Completed  bool   `redis:"completed"`
}

type RedisCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCache(c *redis.Client, t time.Duration) *RedisCache {
	return &RedisCache{client: c, ttl: t}
}

func (r *RedisCache) Get(key string) (*CachedResponse, bool) {
	ctx := context.Background()

	data, err := r.client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, false
	}
	if err != nil {
		return nil, false
	}

	var response CachedResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, false
	}

	return &response, true
}

func (r *RedisCache) StartProcessing(key string) bool {
	ctx := context.Background()

	status := r.client.SetArgs(ctx, key, "processing", redis.SetArgs{Mode: "NX", TTL: 5 * time.Minute})
	if err := status.Err(); err != nil {
		return false
	}

	if status.Val() == "OK" {
		return true
	}

	return false
}

func (r *RedisCache) Finish(key string, status int, body []byte) {
	ctx := context.Background()

	response := &CachedResponse{StatusCode: status, Body: body, Completed: true}
	data, err := json.Marshal(response)
	if err != nil {
		return
	}

	err = r.client.Set(ctx, key, data, r.ttl).Err()
	if err != nil {
		return
	}
}
