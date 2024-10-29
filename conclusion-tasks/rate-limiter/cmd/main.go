package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/yodalis/golang/conclusion-tasks/rate-limiter/config"
	"github.com/yodalis/golang/conclusion-tasks/rate-limiter/internal/limiter"
	"github.com/yodalis/golang/conclusion-tasks/rate-limiter/internal/middleware"
)

func main() {
	cfg := config.LoadConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})

	rl := limiter.NewRateLimiter(rdb, cfg)

	ipLimit := cfg.IPRequestLimit
	tokenLimit := cfg.TokenRequestLimit
	blockTimeSeconds := cfg.BlockTimeSeconds

	http.Handle("/", middleware.RateLimitMiddleware(rl, ipLimit, tokenLimit, blockTimeSeconds)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	})))

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
