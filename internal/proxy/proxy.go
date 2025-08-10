package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// NewProxy 创建一个新的反向代理处理器
func NewProxy(targetURL string) (http.Handler, error) {
	target, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}

	// httputil.NewSingleHostReverseProxy 会创建一个处理所有请求的代理
	// 它会重写 host header，并将请求转发到 target
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 我们可以在这里修改请求或响应，但现在保持简单
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host // 这是关键，确保 Host header 被正确设置
	}

	log.Printf("Forwarding requests to: %s\n", targetURL)
	return proxy, nil
}
