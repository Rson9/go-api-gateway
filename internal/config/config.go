// internal/config/config.go
package config

// Config 是整个配置文件的结构
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Routes []*Route     `mapstructure:"routes"`
}

// ServerConfig 是服务器相关的配置
type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// Route 是一条路由规则
type Route struct {
	Name   string `mapstructure:"name"`
	Path   string `mapstructure:"path"`
	Target string `mapstructure:"target"`
}
