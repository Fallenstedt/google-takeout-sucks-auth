package handlers

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"runtime"
	"time"

	"github.com/yuin/goldmark"

	"github.com/Fallenstedt/google-takeout-sucks-auth/internal/google_auth"
	"github.com/Fallenstedt/google-takeout-sucks-auth/internal/logging"
)

var startTime = time.Now()
var homeETag string

func init() {
	sum := sha256.Sum256(homeContent)
	homeETag = fmt.Sprintf(`"%x"`, sum[:])
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Convert Markdown to HTML using goldmark
	md := goldmark.New()
	var buf bytes.Buffer
	if err := md.Convert(homeContent, &buf); err != nil {
		logging.ErrorLog.Printf("failed to convert markdown: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	t, err := template.New("home").Parse(page)
	if err != nil {
		logging.ErrorLog.Printf("failed to parse template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	w.Header().Set("ETag", homeETag)
	if err := t.Execute(w, struct{ Content template.HTML }{Content: template.HTML(buf.String())}); err != nil {
		logging.ErrorLog.Printf("failed to execute template: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func Status(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type healthResp struct {
		Status     string `json:"status"`
		Timestamp  string `json:"timestamp"`
		Uptime     string `json:"uptime"`
		Goroutines int    `json:"goroutines"`
		AllocBytes uint64 `json:"alloc_bytes"`
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	resp := healthResp{
		Status:     "ok",
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Uptime:     time.Since(startTime).String(),
		Goroutines: runtime.NumGoroutine(),
		AllocBytes: m.Alloc,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logging.ErrorLog.Printf("failed to encode status response: %v", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

// Login should redirect the user to Login Screen
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	auth := google_auth.GoogleAuth{}
	state, err := auth.GenerateStateToken()
	if err != nil {
		logging.ErrorLog.Printf("failed to generate state token %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	authUrl := auth.GoogleOAuthEndpoint(r.Context(), state)

	// Store state in cookie for verification in callback
	cookie := &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(10 * time.Minute),
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)
}

// Callback should compare state token and use Code
func Callback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	state := r.URL.Query().Get("state")
	// Read stored state from cookie
	c, err := r.Cookie("oauth_state")
	if err != nil {
		logging.ErrorLog.Printf("state cookie missing: %v", err)
		http.Error(w, "state cookie missing", http.StatusBadRequest)
		return
	}
	if state == "" || c.Value != state {
		logging.ErrorLog.Printf("discovered mismatch state, query=%s cookie=%s", state, c.Value)
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		logging.ErrorLog.Println("code not found")
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	auth := google_auth.GoogleAuth{}
	token, err := auth.ExchangeToken(r.Context(), code)
	if err != nil {
		logging.ErrorLog.Printf("failed to exchange token %v", err)
		http.Error(w, "Failed exchange", http.StatusBadRequest)
		return
	}
	// Return the token as JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		logging.ErrorLog.Printf("failed to encode token to json: %v", err)
		// If encoding fails, there's not much we can do; send a 500
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
