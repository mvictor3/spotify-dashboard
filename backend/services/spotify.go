package services

import (
	"fmt"
	"net/http"
	"time"
	"encoding/base64"
	"io"
	"net/url"
	"strings"
)

func GetAccessToken() {
	 	http.HandleFunc("/artist", showArtists)
		fmt.Println("Server running on :8080")
		 http.ListenAndServe(":8080",nil) 
	}

func showArtists(w http.ResponseWriter, r *http.Request){
	 const clientID = "my client id"
		const clientSecret = "client secret id"
		const tokenURL = "https://accounts.spotify.com/api/token"

		authString:= clientID + ":" + clientSecret
		encodedAuth := base64.StdEncoding.EncodeToString([]byte(authString))



		data := url.Values{}
		data.Set("grant_type", "client_credentials")

		requestBody := strings.NewReader(data.Encode())

		client := &http.Client{Timeout: 10 * time.Second}

		req, err:= http.NewRequest("POST", tokenURL, requestBody)
		if err != nil {
			fmt.Println("Error executing request:", err)
			return
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded") 
		req.Header.Set("Authorization", "Basic "+encodedAuth)


		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error executing request:", err)
			return
		}
		defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("Request failed. Status: %s\n", resp.Status)
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Error Body: %s\n", string(body))
	}

		body, err := io.ReadAll(resp.Body)
		if err != nil{
			fmt.Println("Error reading response body:", err)
			return
		}



		fmt.Println("Access Token Request Successful!")
		fmt.Printf("Response Body: %s\n", string(body))
	w.Write(body)
}


