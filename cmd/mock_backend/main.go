package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
    if len(os.Args) < 3 {
         log.Fatal("Usage: go run ./cmd/mock_backend/main.go <port> <service_name>")
    }
    port := os.Args[1]
    serviceName := os.Args[2]

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request received on %s for path: %s", serviceName, r.URL.Path)
        fmt.Fprintf(w, "Hello from %s! You requested: %s", serviceName, r.URL.Path)
    })

    addr := ":" + port
    log.Printf("%s listening on %s", serviceName, addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("Could not start %s: %s\n", serviceName, err)
    }
}
