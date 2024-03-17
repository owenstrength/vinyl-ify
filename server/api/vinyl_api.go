package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Album struct {
	Title string   `json:"title"`
	Links []string `json:"links"`
}

type Vinyl struct {
	Artist string  `json:"artist"`
	Albums []Album `json:"albums"`
}

func HandleVinylSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	artist := r.URL.Query().Get("artist")

	fmt.Println("Handling vinyl search for", artist)

	// Example data for response
	response := Vinyl{
		Artist: artist,
		Albums: []Album{
			{Title: "Album1", Links: []string{"Link1"}},
			{Title: "Album2", Links: []string{"Link2"}},
		},
	}

	// Convert response to JSON
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to get vinyls", http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
