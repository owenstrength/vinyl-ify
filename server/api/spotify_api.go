package api

import (
	"Server/service"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	cfg := LoadConfig()
	fmt.Println(cfg)

	// Construct the Spotify authorization URL
	authURL := fmt.Sprintf("https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=%s",
		cfg.SpotifyClientID, cfg.SpotifyRedirectURI, url.QueryEscape(strings.Join(cfg.Scopes[:], " ")))

	// Redirect the user to the Spotify authorization URL
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func HandleSpotifyCallback(w http.ResponseWriter, r *http.Request) {
	// Extract authorization code from query parameters
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing authorization code", http.StatusBadRequest)
		return
	}

	// Exchange authorization code for access token
	accessToken, err := service.ExchangeCodeForAccessToken(code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to exchange authorization code: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Store access token securely in your backend, associate it with the authenticated user
	fmt.Println(accessToken)
	// Redirect or respond to frontend
	http.Redirect(w, r, "/success", http.StatusFound)
}
