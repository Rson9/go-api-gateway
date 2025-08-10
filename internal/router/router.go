// internal/router/router.go
package router

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sort"
	"strings"

	"github.com/rson9/go-api-gateway/internal/config"
)

// router 结构体负责处理所有路由逻辑
type router struct {
	routes []*config.Route
	// 预先为每个路由创建一个代理，提高性能
	proxies map[string]*httputil.ReverseProxy
}
 
// 确保 router 实现了 Matcher 接口
var _ Matcher = (*router)(nil)

// 实现Handler接口
var _ http.Handler = (*router)(nil)

// NewRouter 创建并初始化路由器
func NewRouter(cfg *config.Config) (*router, error) {
	r := &router{
		routes:  cfg.Routes,
		proxies: make(map[string]*httputil.ReverseProxy),
	}

	// 为了正确处理路径覆盖（例如 /a/b 应该优先于 /a），
	// 我们需要按路径长度降序排序。
	sort.Slice(r.routes, func(i, j int) bool {
		return len(r.routes[i].Path) > len(r.routes[j].Path)
	})

	for _, route := range r.routes {
		targetURL, err := url.Parse(route.Target)
		if err != nil {
			return nil, err
		}
		r.proxies[route.Path] = httputil.NewSingleHostReverseProxy(targetURL)
	}
	return r, nil
}

// Match 根据请求查找匹配的路由规则，但不执行代理
func (r *router) Match(req *http.Request) *config.Route {
    for _, route := range r.routes {
        if strings.HasPrefix(req.URL.Path, route.Path) {
            return route
        }
    }
    return nil
}

// ServeHTTP 是路由器的入口点，它实现了 http.Handler 接口
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	matchedRoute := r.Match(req)
	if matchedRoute != nil {
		proxy := r.proxies[matchedRoute.Path]
		proxy.ServeHTTP(w, req)
		return
	}
	http.NotFound(w, req)
}
