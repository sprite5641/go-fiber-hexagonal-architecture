package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr string, password string, db int) *RedisClient {
	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
	}
}

func (r *RedisClient) SetKey(key string, value string) error {
	_, err := r.client.Set(ctx, key, value, 0).Result()
	return err
}

func (r *RedisClient) GetKey(key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisClient) DeleteKey(key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisClient) Close() error {
	return r.client.Close()
}

func (r *RedisClient) Ping() error {
	return r.client.Ping(ctx).Err()
}

func (r *RedisClient) SetKeyWithExpire(key string, value string, expiration time.Duration) error {
	_, err := r.client.Set(ctx, key, value, expiration).Result()
	return err
}
