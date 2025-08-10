# Go API Gateway

一个使用 Golang 实现的高性能、可扩展的 API 网关。

## ✨ 功能列表 (Features)
 
- [x] 反向代理
- [x] 动态路由 (基于路径前缀)
- [x] 请求限流 (基于令牌桶算法)
- [x] Prometheus 指标监控
 
## 🚀 快速开始 (Quick Start)
 
1.  **启动模拟后端服务 (需要 3 个终端)**:
    ```bash
    go run ./cmd/mock_backend/main.go 9091 service_a
    go run ./cmd/mock_backend/main.go 9092 service_b
    go run ./cmd/mock_backend/main.go 9093 service_a_users
    ```
2. **启动 API 网关**:
    ```bash
    go run ./cmd/gateway/main.go
    ```

3. **发送请求**:
    ```bash
    # 请求将被路由到 service_a
    curl http://127.0.0.1:8080/service/a/anything

    # 请求将被路由到 service_a_users (更具体的匹配优先)
    curl http://127.0.0.1:8080/service/a/users/1
    ```