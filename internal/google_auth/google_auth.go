package google_auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

var _config *oauth2.Config

type IGoogleAuth interface {
	GoogleOAuthEndpoint(ctx context.Context, state string) string
	ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error)
	GenerateStateToken() (string, error)
}

type GoogleAuth struct{}

// Ensure GoogleAuth implements IGoogleAuth.
var _ IGoogleAuth = &GoogleAuth{}

func (g *GoogleAuth) GoogleOAuthEndpoint(ctx context.Context, state string) string {
	config := g.getGoogleConfig()

	authURL := config.AuthCodeURL(state, oauth2.AccessTypeOffline)

	return authURL
}

func (g *GoogleAuth) ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error) {
	config := g.getGoogleConfig()
	return config.Exchange(ctx, code)
}

func (g *GoogleAuth) getGoogleConfig() *oauth2.Config {
	if _config != nil {
		return _config
	}

	// TODO use correct credential file based on environment
	b, err := os.ReadFile("dev-credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope, drive.DriveReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	_config = config

	return _config
}

func (g *GoogleAuth) GenerateStateToken() (string, error) {
	// Generate a secure random state (32 bytes -> 64 hex chars)
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	state := hex.EncodeToString(b)

	return state, nil
}
