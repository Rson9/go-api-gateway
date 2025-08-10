package main

import (
	"log"

	"github.com/rson9/go-api-gateway/internal/config"
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

	// 4. 启动服务器
	server.Start(cfg.Server.Port, routerHandler)
}
