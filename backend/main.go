package main

import (
	"fmt"
	"log"
	"net/http"
	"spotify-dashboard/backend/handlers"
	"spotify-dashboard/backend/repository"
	"spotify-dashboard/backend/services"
)

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Spotify API - Try: http://localhost:8080/spotify/tracks?q=Levitating")) // Added closing )
	})
	http.HandleFunc("/favorites", favoriteHandler.GetFavorite)
	http.HandleFunc("/favorites/save", favoriteHandler.SaveFavorite)
	http.HandleFunc("/spotify/status", spotifyHandler.HealthCheck)
	http.HandleFunc("/spotify/tracks", spotifyHandler.ShowTracks)
	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
