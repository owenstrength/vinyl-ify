package api

import (
	"Server/config"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	cfg := config.LoadConfig()
	fmt.Println(cfg)

	// Construct the Spotify authorization URL
	authURL := fmt.Sprintf("https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=%s",
		cfg.SpotifyClientID, cfg.SpotifyRedirectURI, url.QueryEscape(strings.Join(cfg.Scopes[:], " ")))

	// Redirect the user to the Spotify authorization URL
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func HandleSpotifyCallback(w http.ResponseWriter, r *http.Request) {
	// Extract authorization code from query parameters
	accessToken := r.URL.Query().Get("code")
	fmt.Println(accessToken)
	if accessToken == "" {
		http.Error(w, "Missing authorization code", http.StatusBadRequest)
		return
	}

	// Set the token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    accessToken,
		MaxAge:   3600,  // Token expiration time
		HttpOnly: false, // HTTP only cookie for security. set to false for now, bad practice but we can fix this later
	})

	// Redirect or respond to frontend
	http.Redirect(w, r, "http://localhost:3000", http.StatusFound)
}
