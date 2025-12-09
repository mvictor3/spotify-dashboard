package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"spotify-dashboard/backend/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type FavoriteRepository struct {
	db *sql.DB
}
type DBConfig struct {
	PostgresPW string
	Host       string
	Port       string
	User       string
	DBname     string
}

func LoadPostgre() (*DBConfig, error) {
	postgresConfig := &DBConfig{

		PostgresPW: os.Getenv("POSTGRES_PASSWORD"),
		Host:       os.Getenv("POSTGRES_HOST"),
		Port:       os.Getenv("POSTGRES_PORT"),
		User:       os.Getenv("POSTGRES_USER"),
		DBname:     os.Getenv("POSTGRES_DB"),
	}
	if postgresConfig.PostgresPW == "" || postgresConfig.Host == "" || postgresConfig.Port == "" || postgresConfig.User == "" || postgresConfig.DBname == "" {
		return nil, errors.New("required POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_USER, POSTGRES_DB AND POSTGRES_PORT")
	}
	return postgresConfig, nil
}
func ConnectDB(config *DBConfig) (*sql.DB, error) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.PostgresPW,
		config.DBname,
	)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to PostgreSQL!")
	return db, nil
}

func NewFavoriteRepository(db *sql.DB) *FavoriteRepository {
	return &FavoriteRepository{db: db}
}
func CreateFavoriteTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS favorites (
		    id SERIAL PRIMARY KEY,
		    spotify_track_id TEXT NOT NULL UNIQUE,
			title TEXT,
			album TEXT,
			artists TEXT,
			preview_url TEXT,
			duration_ms INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
		`
	_, err := db.Exec(query)
	return err
}
func (r *FavoriteRepository) SaveTrackAsFavorite(favorite models.Favorite) error {

	jsonData, err := json.Marshal(favorite.Artists)
	if err != nil {
		return err
	}
	artistsString := string(jsonData)
	_, err = r.db.Exec(`
		INSERT INTO favorites(spotify_track_id, title, album, artists, preview_url, duration_ms)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (spotify_track_id) DO NOTHING`,

		favorite.ID,
		favorite.Title,
		favorite.Album.Name,
		artistsString,
		favorite.PreviewURL,
		favorite.DurationMs,
	)

	return err
}

func (r *FavoriteRepository) GetFavorites() ([]models.Favorite, error) {

	query := `
		select spotify_track_id, title, album, artists, preview_url, duration_ms, created_at
		FROM favorites	
		ORDER BY created_at DESC
		LIMIT 10`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []models.Favorite
	for rows.Next() {
		var f models.Favorite
		var artistsJSON string
		var albumName string

		err := rows.Scan(
			&f.ID,
			&f.Title,
			&albumName,
			&artistsJSON,
			&f.PreviewURL,
			&f.DurationMs,
			&f.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(artistsJSON), &f.Artists); err != nil {
			return nil, err
		}
		f.Album = models.Album{Name: albumName} // ← Missing: set the album
		favorites = append(favorites, f)

	}
	return favorites, rows.Err()
}
