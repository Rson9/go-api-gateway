package router

import (
	"net/http"

	"github.com/rson9/go-api-gateway/internal/config"
)

// Matcher 接口定义了路由匹配的能力
type Matcher interface {
	Match(*http.Request) *config.Route
}
