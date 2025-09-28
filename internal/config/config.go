package config

import (
	"os"
	"strings"
)

// Config holds application-wide configuration.
type Config struct {
    // Env should be "dev" or "prod"
    Env string
}

// Cfg is the global configuration instance.
var Cfg = load()

func load() *Config {
    env := os.Getenv("APP_ENV")
    if env == "" {
        env = "dev"
    }
    env = strings.ToLower(env)
    return &Config{Env: env}
}

// CredentialsFile returns the credentials file name based on environment.
func (c *Config) CredentialsFile() string {
    if c.Env == "prod" {
        return "credentials.json"
    }
    return "dev-credentials.json"
}

