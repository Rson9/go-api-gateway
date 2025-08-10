package main

import (
	"log"

	"github.com/rson9/go-api-gateway/internal/proxy"
	"github.com/rson9/go-api-gateway/internal/server"
	"github.com/spf13/viper"
)

func main() {
    // 1. 初始化配置
    viper.SetConfigName("config")    // 配置文件名 (不带扩展名)
    viper.SetConfigType("yaml")      // 配置文件类型
    viper.AddConfigPath("./configs") // 配置文件路径
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }

    // 2. 获取配置
    port := viper.GetString("server.port")
    targetURL := viper.GetString("proxy.target_url")

    if port == "" || targetURL == "" {
        log.Fatal("Server port or proxy target_url not defined in config")
    }

    // 3. 创建反向代理处理器
    proxyHandler, err := proxy.NewProxy(targetURL)
    if err != nil {
        log.Fatalf("Could not create proxy: %s", err)
    }

    // 4. 启动服务器
    server.Start(port, proxyHandler)
}
