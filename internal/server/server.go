package server

import (
	"fmt"
	"log"
	"net/http"
)

// Start 启动网关服务器
func Start(port string, handler http.Handler) {
	addr := fmt.Sprintf(":%s", port)
	log.Printf("API Gateway starting on %s", addr)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
