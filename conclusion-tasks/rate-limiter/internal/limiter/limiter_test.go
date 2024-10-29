package limiter_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"github.com/yodalis/golang/conclusion-tasks/rate-limiter/config"
	"github.com/yodalis/golang/conclusion-tasks/rate-limiter/internal/limiter"
)

func TestAllow_IPLimit(t *testing.T) {
	db, mock := redismock.NewClientMock()
	defer db.Close()

	cfg := config.Config{IPRequestLimit: 2, BlockTimeSeconds: 30}
	rl := limiter.NewRateLimiter(db, cfg)
	ctx := context.Background()
	blockKey := "ip:test_ip:blocked"
	testIpKey := "ip:test_ip"
	ip := "test_ip"

	mock.ExpectExists(blockKey).SetVal(0)
	mock.ExpectIncr(testIpKey).SetVal(1)
	assert.True(t, rl.Allow(ctx, ip, "", cfg.IPRequestLimit, 0, cfg.BlockTimeSeconds))

	mock.ExpectExists(blockKey).SetVal(0)
	mock.ExpectIncr(testIpKey).SetVal(2)
	assert.True(t, rl.Allow(ctx, ip, "", cfg.IPRequestLimit, 0, cfg.BlockTimeSeconds))

	mock.ExpectExists(blockKey).SetVal(0)
	mock.ExpectIncr(testIpKey).SetVal(3)
	mock.ExpectSet(blockKey, 1, time.Duration(cfg.BlockTimeSeconds)*time.Second).SetVal("OK")
	mock.ExpectDel(testIpKey).SetVal(1)
	assert.False(t, rl.Allow(ctx, ip, "", cfg.IPRequestLimit, 0, cfg.BlockTimeSeconds))

	mock.ExpectExists(blockKey).SetVal(1)
	assert.False(t, rl.Allow(ctx, ip, "", cfg.IPRequestLimit, 0, cfg.BlockTimeSeconds))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %v", err)
	}
}

func TestAllow_TokenLimit(t *testing.T) {
	db, mock := redismock.NewClientMock()
	defer db.Close()

	cfg := config.Config{TokenRequestLimit: 3, BlockTimeSeconds: 10}
	rl := limiter.NewRateLimiter(db, cfg)
	token := "test_token"
	tokenKey := "token:test_token"
	blockedKey := "token:test_token:blocked"

	mock.ExpectIncr(tokenKey).SetVal(1)
	mock.ExpectIncr(tokenKey).SetVal(2)
	mock.ExpectIncr(tokenKey).SetVal(3)
	mock.ExpectIncr(tokenKey).SetVal(4)
	mock.ExpectSet(blockedKey, 1, time.Duration(cfg.BlockTimeSeconds)*time.Second).SetVal("OK")

	ctx := context.Background()
	assert.True(t, rl.Allow(ctx, "", token, 0, cfg.TokenRequestLimit, cfg.BlockTimeSeconds))
	assert.True(t, rl.Allow(ctx, "", token, 0, cfg.TokenRequestLimit, cfg.BlockTimeSeconds))
	assert.True(t, rl.Allow(ctx, "", token, 0, cfg.TokenRequestLimit, cfg.BlockTimeSeconds))
	assert.False(t, rl.Allow(ctx, "", token, 0, cfg.TokenRequestLimit, cfg.BlockTimeSeconds))

	mock.ExpectExists(blockedKey).SetVal(1)
	assert.False(t, rl.Allow(ctx, "", token, 0, cfg.TokenRequestLimit, cfg.BlockTimeSeconds))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %v", err)
	}
}
