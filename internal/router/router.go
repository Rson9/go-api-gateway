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

// NewRouter 创建并初始化路由器
func NewRouter(cfg *config.Config) (http.Handler, error) {
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

// ServeHTTP 是路由器的入口点，它实现了 http.Handler 接口
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 遍历排序后的路由，找到第一个匹配的
	for _, route := range r.routes {
		if strings.HasPrefix(req.URL.Path, route.Path) {
			proxy := r.proxies[route.Path]
			// 在转发前，我们可以修改请求头等信息
			// 例如，移除匹配到的前缀，如果后端服务不需要它
			// req.URL.Path = strings.TrimPrefix(req.URL.Path, route.Path)
			proxy.ServeHTTP(w, req)
			return
		}
	}

	// 如果没有找到任何匹配的路由，返回 404
	http.NotFound(w, req)
}
