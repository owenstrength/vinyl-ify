package api

import (
	"Server/config"
	"Server/service"
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

	authJSON, err := service.GetAuthToken(accessToken)
	if err != nil {
		http.Error(w, "Failed to get auth token "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("API TOKEN: ", authJSON)

	// Set the token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    authJSON,
		MaxAge:   3600,  // Token expiration time
		HttpOnly: false, // HTTP only cookie for security. set to false for now, bad practice but we can fix this later
	})

	// Redirect or respond to frontend
	http.Redirect(w, r, "http://localhost:3000", http.StatusFound)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	// Get the access token from the cookie
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	cookie, err := r.Cookie("auth_token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accessToken := cookie.Value
	fmt.Println("COOKIE: ", accessToken)

	// Get the user ID
	userData, err := service.GetUserID(accessToken)
	if err != nil {
		http.Error(w, "Failed to get user ID", http.StatusInternalServerError)
		return
	}

	// Respond with the user ID
	w.Write([]byte(userData))
}

func HandleGetArtists(w http.ResponseWriter, r *http.Request) {
	// Get the access token from the cookie
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	cookie, err := r.Cookie("auth_token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accessToken := cookie.Value
	fmt.Println("COOKIE: ", accessToken)

	// Get the user's top artists
	artists, err := service.GetUserTopArtists(accessToken)
	if err != nil {
		http.Error(w, "Failed to get top artists", http.StatusInternalServerError)
		return
	}

	// Respond with the user's top artists
	w.Write([]byte(artists))
}
