package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Fallenstedt/google-takeout-sucks-auth/internal/google_auth"
	"github.com/Fallenstedt/google-takeout-sucks-auth/internal/logging"
)

// Login should redirect the user to Login Screen
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

  auth := google_auth.GoogleAuth{}
	authUrl := auth.GoogleOAuthEndpoint(r.Context())

	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)
}

// Callback should compare state token and use Code
func Callback(w http.ResponseWriter, r *http.Request) {
  	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
  
  state := r.URL.Query().Get("state");
	if state != "state-token" {
    logging.ErrorLog.Printf("discovered mismatch state, %s %s", state, "state-token")
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
