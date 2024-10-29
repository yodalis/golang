package middleware

import (
	"net/http"

	"github.com/yodalis/golang/conclusion-tasks/rate-limiter/internal/limiter"
)

type rateLimiterResult func(h http.Handler) http.Handler

func RateLimitMiddleware(rl *limiter.RateLimiter, ipLimit, tokenLimit, blockTimeSeconds int) rateLimiterResult {
	return func(h http.Handler) http.Handler {
		return rateLimiterHandler(rl, ipLimit, tokenLimit, blockTimeSeconds, h)
	}
}

func rateLimiterHandler(rl *limiter.RateLimiter, ipLimit, tokenLimit int, blockTimeSeconds int, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		token := r.Header.Get("API_KEY")

		if !rl.Allow(r.Context(), ip, token, ipLimit, tokenLimit, blockTimeSeconds) {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("you have reached the maximum number of requests or actions allowed within a certain time frame"))
			return
		}

		h.ServeHTTP(w, r)
	})
}
