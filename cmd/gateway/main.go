package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rson9/go-api-gateway/internal/config"
	"github.com/rson9/go-api-gateway/internal/limiter"
	"github.com/rson9/go-api-gateway/internal/middleware"
	"github.com/rson9/go-api-gateway/internal/router"
	"github.com/rson9/go-api-gateway/internal/server"
	"github.com/spf13/viper"
)

func main() {
	// 1. 初始化配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 2. 将配置 unmarshal 到我们的结构体中
	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	if cfg.Server.Port == "" {
		log.Fatal("Server port not defined in config")
	}
	if len(cfg.Routes) == 0 {
		log.Fatal("No routes defined in config")
	}

	// 3. 创建我们的路由器
	routerHandler, err := router.NewRouter(&cfg)
	if err != nil {
		log.Fatalf("Could not create router: %s", err)
	}

	// 4. 应用中间件
	var finalHandler http.Handler = routerHandler
	// Metrics 中间件应该包裹在最外层，以便捕获所有请求，包括被限流的
	// 它需要一个 routeMatcher 来获取路由信息
	finalHandler = middleware.Metrics(finalHandler, routerHandler)

	if cfg.RateLimiter.Enabled {
		log.Printf("Rate limiter is enabled: rate=%.2f, burst=%d", cfg.RateLimiter.Rate, cfg.RateLimiter.Burst)
		tokenBucket := limiter.NewTokenBucket(cfg.RateLimiter.Rate, cfg.RateLimiter.Burst)
		// 将我们的路由器包裹在限流中间件中
		finalHandler = middleware.RateLimit(tokenBucket)(finalHandler)
	}

	// 5. 启动管理服务器 (用于 /metrics)
	go func() {
		log.Printf("Management server starting on port %s", cfg.Management.Port)
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler()) // 暴露 Prometheus metrics
		addr := fmt.Sprintf(":%s", cfg.Management.Port)
		if err := http.ListenAndServe(addr, mux); err != nil {
			log.Fatalf("Could not start management server: %s", err)
		}
	}()

	// 6. 启动服务器
	server.Start(cfg.Server.Port, finalHandler)
}
