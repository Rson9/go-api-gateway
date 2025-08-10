// internal/middleware/metrics.go
package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/rson9/go-api-gateway/internal/metrics" // 你的模块路径
	"github.com/rson9/go-api-gateway/internal/router"  // 需要访问路由信息
)

// responseWriterInterceptor 用于捕获 status code
type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Metrics 是一个中间件，用于记录 Prometheus 指标
func Metrics(next http.Handler, routeMatcher router.Matcher) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		metrics.HTTPRequestsInFlight.Inc()
		defer metrics.HTTPRequestsInFlight.Dec()

		// 使用自定义的 responseWriter 来捕获状态码
		interceptor := &responseWriterInterceptor{ResponseWriter: w, statusCode: http.StatusOK}

		// 找到匹配的路由，以便添加 'route' 标签
		matchedRoute := routeMatcher.Match(r)
		routeName := "unknown"
		if matchedRoute != nil {
			routeName = matchedRoute.Name
		}

		defer func() {
			duration := time.Since(start)
			// 记录延迟
			metrics.HTTPRequestDuration.WithLabelValues(routeName, r.Method).Observe(duration.Seconds())
			// 记录总请求数
			metrics.HTTPRequestsTotal.WithLabelValues(routeName, r.Method, strconv.Itoa(interceptor.statusCode)).Inc()
		}()

		next.ServeHTTP(interceptor, r)
	})
}
