package models
import(

)
type Artist struct{
  ID string `json:"id"`
  Name string `json:"name"`
}

type Track struct {
  
ID string `json:"id"`
  
Artists []Artist `json:"aritst"`
  
Title string `json:"name"`
  
SpotifyURL string `json:"external_urls"`

DurationsMs int `json:"duration_ms"`
}
