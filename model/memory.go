package model

import (
	"time"
)

// jsonとmappingさせる構造体
type Memory struct {
	Id          int64     `json:"id"`
	Memory      string    `json:"memory"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
	Episodes    []Episode `json:"episodes"`
	Image       string    `json:"image"`
	AuthorId    int64    `json:"author_id" db:"author_id"`
	Angle       float64   `json:"angle"`
	Seen        bool      `json:"seen"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Episode struct {
	Id        int64  `json:"id" binding:"required"`
	Episode   string  `json:"episode" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
}
