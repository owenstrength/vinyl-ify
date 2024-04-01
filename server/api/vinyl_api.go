package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type Album struct {
	Title string `json:"title"`
	Link  string `json:"link"`
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

	fmt.Println("Got vinyls:", string(out))

	cmd2 := exec.Command("python", "./service/vinyl_10k_search.py", artist)
	out2, err := cmd2.CombinedOutput()
	if err != nil {
		log.Println("Failed to get vinyls:", err)
		http.Error(w, "Failed to get vinyls", http.StatusInternalServerError)
		return
	}

	data := make(map[string]string)

	// Unmarshal the JSON string into the map
	err = json.Unmarshal(out2, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var albums []Album
	albums = append(albums, Album{Title: "Artist Website", Link: string(out)})

	for key, value := range data {
		albums = append(albums, Album{Title: key, Link: value})
	}

	response := Vinyl{
		Artist: artist,
		Albums: albums,
	}

	// Convert response to JSON
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to get vinyls", http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
