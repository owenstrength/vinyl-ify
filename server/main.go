package main

import (
	"fmt"
	"log"
	"net/http"

	"Server/api"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/login", api.HandleLogin).Methods("GET")
	r.HandleFunc("/spotify-callback", api.HandleSpotifyCallback)

	// Start the server
	port := ":8000"
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
