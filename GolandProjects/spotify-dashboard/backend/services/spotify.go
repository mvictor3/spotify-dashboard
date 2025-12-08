package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"spotify-dashboard/backend/models"
	"strings"
	"sync"
	"time"
)

const tokenURL = "https://accounts.spotify.com/api/token"

type SpotifyService struct {
	clientID     string
	clientSecret string
	httpClient   *http.Client
	accessToken  string
	tokenExpiry  time.Time
	mu           sync.Mutex
}
type CredentialConfig struct {
	ClientID     string
	ClientSecret string
}

func LoadConfig() (*CredentialConfig, error) {
	config := &CredentialConfig{

		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
	}
	if config.ClientID == "" || config.ClientSecret == "" {
		return nil, errors.New("required CLIENT_ID and CLIENT_SECRET, environment is not set")
	}
	return config, nil
}
func NewSpotifyService(clientID, clientSecret string) *SpotifyService {
	return &SpotifyService{
		clientID:     clientID,
		clientSecret: clientSecret,
		httpClient:   &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *SpotifyService) GetAccessToken() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if time.Now().Before(s.tokenExpiry.Add(-30 * time.Second)) {
		return nil
	}

	authString := s.clientID + ":" + s.clientSecret
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(authString))

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	requestBody := strings.NewReader(data.Encode())

	req, err := http.NewRequest("POST", tokenURL, requestBody)
	if err != nil {
		return err // Return error instead of printing
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err // Return error instead of printing
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("token request failed: %s - %s", resp.Status, string(body))
	}

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// NEW PART: Parse and save the token
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return err
	}

	// SAVE the token and expiry
	s.accessToken = tokenResponse.AccessToken
	s.tokenExpiry = time.Now().Add(time.Duration(tokenResponse.ExpiresIn) * time.Second)

	return nil
}

func (s *SpotifyService) SearchTracks(query string) ([]models.Track, error) {

	searchURL := "https://api.spotify.com/v1/search"
	params := url.Values{}
	params.Add("q", query)
	params.Add("type", "track")
	params.Add("limit", "10")
	fullURL := searchURL + "?" + params.Encode()

	if err := s.GetAccessToken(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+s.accessToken)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("track search request failed: %s - %s", resp.Status, string(body))
	}

	var response models.SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Tracks.Items, nil
}
