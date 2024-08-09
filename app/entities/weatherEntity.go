package entities

import (
	"time"
)

type (
	WeatherEntity struct {
		Id          uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
		Longitude   float64   `json:"longitude"`
		Latitude    float64   `json:"latitude"`
		WindSpeed   float64   `json:"windSpeed"`
		Temperature float64   `json:"temperature"`
		Rain        float64   `json:"rain"`
		CreatedAt   time.Time `json:"createdAt"`
	}
)
