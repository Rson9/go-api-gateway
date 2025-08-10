// internal/middleware/ratelimit.go
package middleware

import (
	"net/http"

	"github.com/rson9/go-api-gateway/internal/limiter"
)

// RateLimit 是一个中间件工厂，它返回一个限流中间件
func RateLimit(limiter *limiter.TokenBucket) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
