package main

import (
	"fmt"
	"log"
	"net/http"
	"spotify-dashboard/backend/handlers"
	"spotify-dashboard/backend/repository"
	"spotify-dashboard/backend/services"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
func main() {
	fmt.Println("🎬 Starting main.go...")
	config, err := services.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	spotifyService := services.NewSpotifyService(config.ClientID, config.ClientSecret)

	dbConfig, err := repository.LoadPostgre()
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.ConnectDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	favRepo := repository.NewFavoriteRepository(db)
	err = repository.CreateFavoriteTable(db)
	if err != nil {
		log.Fatal(err)
	}
	favoriteHandler := handlers.NewFavoriteHandler(spotifyService, favRepo)
	spotifyHandler := handlers.NewSpotifyHandler(spotifyService)

	http.HandleFunc("/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Spotify API - Try: http://localhost:8080/spotify/tracks?q=Levitating")) // Added closing )
	}))
	http.HandleFunc("/favorites/delete", enableCORS(favoriteHandler.DeleteFavoriteTrack))
	http.HandleFunc("/favorites", enableCORS(favoriteHandler.GetFavorite))
	http.HandleFunc("/favorites/save", enableCORS(favoriteHandler.SaveFavorite))
	http.HandleFunc("/spotify/status", enableCORS(spotifyHandler.HealthCheck))
	http.HandleFunc("/spotify/tracks", enableCORS(spotifyHandler.ShowTracks))
	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
