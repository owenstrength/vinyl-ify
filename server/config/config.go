package config

import (
    "os"
)

// Config struct holds application configuration
type Config struct {
    SpotifyClientID     string
    SpotifyClientSecret string
    SpotifyRedirectURI  string
}

// LoadConfig loads application configuration from environment variables
func LoadConfig() *Config {
    return &Config{
        SpotifyClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
        SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
        SpotifyRedirectURI:  os.Getenv("SPOTIFY_REDIRECT_URI"),
    }
}
