// internal/config/config.go
package config

// Config 是整个配置文件的结构
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	RateLimiter  RateLimiterConfig`mapstructure:"rate_limiter"` // 新增
	Routes []*Route     `mapstructure:"routes"`
}

// ServerConfig 是服务器相关的配置
type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// RateLimiterConfig 是限流器相关的配置
type RateLimiterConfig struct {
    Enabled bool    `mapstructure:"enabled"`
    Rate    float64 `mapstructure:"rate"`
    Burst   int     `mapstructure:"burst"`
}

// Route 是一条路由规则
type Route struct {
	Name   string `mapstructure:"name"`
	Path   string `mapstructure:"path"`
	Target string `mapstructure:"target"`
}

