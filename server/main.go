package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/login", handleLogin).Methods("GET")

	// Start the server
	port := ":8000"
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	cfg := LoadConfig()
	fmt.Println(cfg)

	// Construct the Spotify authorization URL
	authURL := fmt.Sprintf("https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=%s",
		cfg.SpotifyClientID, cfg.SpotifyRedirectURI, url.QueryEscape(strings.Join(cfg.Scopes[:], " ")))

	// Redirect the user to the Spotify authorization URL
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}
