package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/owenstrength/vinyl-ify/server/config"
)

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/login", handleLogin).Methods("GET")

	// Start the server
	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Load Spotify configuration from local file
	cfg := config.LoadConfig("config.json")

	// Construct the Spotify authorization URL
	authURL := fmt.Sprintf("https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=%s",
		cfg.SpotifyClientID, cfg.SpotifyRedirectURI, url.QueryEscape(Scopes))

	// Redirect the user to the Spotify authorization URL
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}
