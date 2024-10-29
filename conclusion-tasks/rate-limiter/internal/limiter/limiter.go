package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yodalis/golang/conclusion-tasks/rate-limiter/config"
)

type RateLimiter struct {
	client *redis.Client
	cfg    config.Config
}

func NewRateLimiter(client *redis.Client, cfg config.Config) *RateLimiter {
	return &RateLimiter{client: client, cfg: cfg}
}

func (rl *RateLimiter) Allow(ctx context.Context, ip, token string, ipLimit, tokenLimit, blockTimeSeconds int) bool {
	if token != "" {
		tokenKey := "token:" + token
		return rl.allowRequest(ctx, tokenKey, tokenLimit, blockTimeSeconds)
	}

	ipKey := "ip:" + ip
	return rl.allowRequest(ctx, ipKey, ipLimit, blockTimeSeconds)
}

func (rl *RateLimiter) allowRequest(ctx context.Context, key string, limit int, blockTimeSeconds int) bool {
	blockKey := key + ":blocked"

	if rl.client.Exists(ctx, blockKey).Val() > 0 {
		return false
	}

	count, _ := rl.client.Incr(ctx, key).Result()

	if count > int64(limit) {
		rl.client.Set(ctx, blockKey, 1, time.Duration(blockTimeSeconds)*time.Second)
		rl.client.Del(ctx, key)

		return false
	}

	return true
}
