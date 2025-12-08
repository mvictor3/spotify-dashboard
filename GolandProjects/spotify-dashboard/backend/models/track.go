package models

import "time"

type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ExternalURLs struct {
	Spotify string `json:"spotify"`
}
type Album struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Track struct {
	ID           string       `json:"id"`
	Artists      []Artist     `json:"artists"`
	Title        string       `json:"name"`
	ExternalURLs ExternalURLs `json:"external_urls"`
	DurationMs   int          `json:"duration_ms"`
	Album        Album        `json:"album"`
	Popularity   int          `json:"popularity"`
	PreviewURL   *string      `json:"preview_url"`
}

type SearchResponse struct {
	Tracks struct {
		Items []Track `json:"items"`
		Total int     `json:"total"`
	} `json:"tracks"`
}
type Favorite struct {
	ID         string    `json:"id"`
	Artists    []Artist  `json:"artists"`
	Title      string    `json:"name"`
	DurationMs int       `json:"duration_ms"`
	Album      Album     `json:"album"`
	PreviewURL *string   `json:"preview_url"`
	CreatedAt  time.Time `json:"created_at"`
}
