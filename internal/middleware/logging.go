package middleware

import (
	"net/http"
	"time"

	"github.com/Fallenstedt/google-takeout-sucks-auth/internal/logging"
)

// LogRequest returns a middleware that logs request method, path and duration.
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logging.InfoLog.Printf("started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		logging.InfoLog.Printf("completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}
