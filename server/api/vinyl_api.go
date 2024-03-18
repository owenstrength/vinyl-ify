package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
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

	// run python ./service/vinyl_search.py
	cmd := exec.Command("python", "./service/vinyl_search.py", artist)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Failed to get vinyls:", err)
		http.Error(w, "Failed to get vinyls", http.StatusInternalServerError)
		return
	}

	out = string(out)

	response := Vinyl{
		Artist: artist,
		Albums: []Album{
			{Title: "ArtistSite", Links: []string{out}},
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
