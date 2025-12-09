package handlers

import (
	"encoding/json"
	"net/http"
	"spotify-dashboard/backend/models"
	"spotify-dashboard/backend/repository"
	"spotify-dashboard/backend/services"
)

type FavoriteHandler struct {
	spotifyService *services.SpotifyService
	favoriteRepo   *repository.FavoriteRepository
}

func NewFavoriteHandler(spotify *services.SpotifyService, repo *repository.FavoriteRepository) *FavoriteHandler {
	return &FavoriteHandler{
		spotifyService: spotify,
		favoriteRepo:   repo,
	}
}

func (h *FavoriteHandler) SaveFavorite(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		TrackID string `json:"track_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if request.TrackID == "" {
		http.Error(w, "track_id is required", http.StatusBadRequest)
		return
	}

	tracks, err := h.spotifyService.SearchTracks(request.TrackID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(tracks) == 0 {
		http.Error(w, "Track not found", http.StatusNotFound)
		return
	}

	track := tracks[0]
	favorite := models.Favorite{
		ID:         track.ID,
		Title:      track.Title,
		Artists:    track.Artists,
		Album:      track.Album,
		DurationMs: track.DurationMs,
		PreviewURL: track.PreviewURL,
	}

	if err := h.favoriteRepo.SaveTrackAsFavorite(favorite); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "saved",
		"track":  favorite.Title,
	})
}

func (h *FavoriteHandler) GetFavorite(w http.ResponseWriter, r *http.Request) {
	favorites, err := h.favoriteRepo.GetFavorites()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count":  len(favorites),
		"tracks": favorites,
	})
}
