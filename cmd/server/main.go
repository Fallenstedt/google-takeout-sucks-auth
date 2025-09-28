package main

import (
	"log"
	"net/http"

	"github.com/Fallenstedt/google-takeout-sucks-auth/internal/handlers"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/login", handlers.Login)

    addr := ":8080"
    log.Printf("starting server on %s", addr)
    if err := http.ListenAndServe(addr, mux); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}

