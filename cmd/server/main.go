package main

import (
	"log"
	"net/http"

	"github.com/Fallenstedt/google-takeout-sucks-auth/internal/handlers"
	"github.com/Fallenstedt/google-takeout-sucks-auth/internal/middleware"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/status", handlers.Status)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/oauth2/callback", handlers.Callback)
	mux.HandleFunc("/", handlers.Home)

	// Wrap mux with logging middleware
	handler := middleware.LogRequest(mux)

	addr := ":8080"
	log.Printf("starting server on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
