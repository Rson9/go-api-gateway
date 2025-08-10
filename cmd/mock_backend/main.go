package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Mock backend received request: %s %s", r.Method, r.URL.Path)
		fmt.Fprintf(w, "Hello from mock backend!")
	})

	log.Println("Mock backend server starting on :9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("Could not start mock backend server: %s\n", err)
	}
}
