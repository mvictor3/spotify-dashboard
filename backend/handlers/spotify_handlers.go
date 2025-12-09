package handlers

import (
	"encoding/json"
	"net/http"
	"spotify-dashboard/backend/services"
)

type SpotifyHandler struct {
	spotifyService *services.SpotifyService
}

func NewSpotifyHandler(svc *services.SpotifyService) *SpotifyHandler {
	return &SpotifyHandler{spotifyService: svc}
}

func (h *SpotifyHandler) ShowTracks(w http.ResponseWriter, r *http.Request) {
	
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing 'q' parameter. Usage:/spotify/tracks?q=YOUR_SEARCH", http.StatusBadRequest)
		return
	}

	tracks, err := h.spotifyService.SearchTracks(query)
	if err != nil {
		http.Error(w, "Search failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"query":  query,
		"count":  len(tracks),
		"tracks": tracks,
	})

}

func (h *SpotifyHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := h.spotifyService.GetAccessToken(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "ok",
		"spotify": "connected",
	})
}
