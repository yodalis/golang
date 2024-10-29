package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"github.com/yodalis/golang/conclusion-tasks/rate-limiter/config"
	"github.com/yodalis/golang/conclusion-tasks/rate-limiter/internal/limiter"
	"github.com/yodalis/golang/conclusion-tasks/rate-limiter/internal/middleware"
)

func Test_Middleware_RateLimitMiddleware_Success(t *testing.T) {
	db, mock := redismock.NewClientMock()
	defer db.Close()

	cfg := config.Config{IPRequestLimit: 2, BlockTimeSeconds: 10}
	rl := limiter.NewRateLimiter(db, cfg)
	ipKey := "ip:127.0.0.1"
	ipBlockedKey := "ip:127.0.0.1"
	ip := "127.0.0.1"
	apiMessage := "Welcome!"

	mock.ExpectIncr(ipKey).SetVal(2)
	mock.ExpectSet(ipBlockedKey, 0, time.Duration(cfg.BlockTimeSeconds)).SetVal("OK")

	handler := middleware.RateLimitMiddleware(rl, cfg.IPRequestLimit, 0, cfg.BlockTimeSeconds)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(apiMessage))
	}))

	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = ip
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, rec.Body.String(), apiMessage)
}

func Test_Middleware_RateLimitMiddleware_Error(t *testing.T) {
	db, mock := redismock.NewClientMock()
	defer db.Close()

	cfg := config.Config{IPRequestLimit: 2, BlockTimeSeconds: 10}
	rl := limiter.NewRateLimiter(db, cfg)
	ipKey := "ip:127.0.0.1"
	ipBlockedKey := "ip:127.0.0.1"
	ip := "127.0.0.1"
	apiMessage := "you have reached the maximum number of requests"

	mock.ExpectIncr(ipKey).SetVal(3)
	mock.ExpectSet(ipBlockedKey, 1, time.Duration(cfg.BlockTimeSeconds)).SetVal("OK")

	handler := middleware.RateLimitMiddleware(rl, cfg.IPRequestLimit, 0, cfg.BlockTimeSeconds)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	}))

	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = ip
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusTooManyRequests, rec.Code)
	assert.Contains(t, rec.Body.String(), apiMessage)
}
