// internal/limiter/token_bucket.go
package limiter

import (
	"sync"
	"time"
)

// TokenBucket 定义了令牌桶限流器所需的所有属性
type TokenBucket struct {
	rate        float64   // 每秒生成的令牌数
	burst       int       // 桶的容量
	tokens      float64   // 当前桶内的令牌数
	lastTokenAt time.Time // 上次放令牌的时间
	mtx         sync.Mutex
}

// NewTokenBucket 创建一个新的令牌桶
func NewTokenBucket(rate float64, burst int) *TokenBucket {
	return &TokenBucket{
		rate:        rate,
		burst:       burst,
		tokens:      float64(burst), // 初始时桶是满的
		lastTokenAt: time.Now(),
	}
}

// Allow 检查是否允许一个请求通过
func (tb *TokenBucket) Allow() bool {
	tb.mtx.Lock()
	defer tb.mtx.Unlock()

	now := time.Now()
	// 计算自上次请求以来，应该生成多少新令牌
	elapsed := now.Sub(tb.lastTokenAt)
	tokensToAdd := elapsed.Seconds() * tb.rate

	tb.tokens += tokensToAdd
	if tb.tokens > float64(tb.burst) {
		tb.tokens = float64(tb.burst) // 令牌数不能超过桶的容量
	}

	// 更新最后一次令牌时间
	tb.lastTokenAt = now

	// 检查是否有足够的令牌
	if tb.tokens >= 1 {
		tb.tokens--
		return true
	}

	return false
}
