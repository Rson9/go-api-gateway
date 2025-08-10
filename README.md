# Go API Gateway

一个使用 Golang 实现的高性能、可扩展的 API 网关。

## ✨ 功能列表 (Features)

- [x] 反向代理
- [ ] 动态路由 (开发中...)

## 🚀 快速开始 (Quick Start)

1.  **启动模拟后端服务 (用于测试)**:
    ```bash
    go run ./cmd/mock_backend/main.go
    ```

2. **启动 API 网关**:
    ```bash
    go run ./cmd/gateway/main.go
    ```

3. **发送请求**:
    ```bash
    curl http://127.0.0.1:8080/any/path
    ```