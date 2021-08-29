package model

type Memory struct {
	Memory      string    `json:"memory"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
	Seen_author []string  `json:"seen_author"`
	Episodes    []Episode `json:"episodes"`
	Image       string    `json:"image"`
	Author      string    `json:"author"`
}

type Episode struct {
	Id        string  `json:"id"`
	Episode   string  `json:"episode"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
