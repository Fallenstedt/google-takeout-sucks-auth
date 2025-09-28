package handlers

import (
	"net/http"

	"github.com/Fallenstedt/google-takeout-sucks-auth/internal/google_drive"
)

// Login should redirect the user to Login Screen
func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    authUrl := google_drive.GoogleOAuthEndpoint(r.Context())

    http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)
}

