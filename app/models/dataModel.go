package model

import (
	"math/rand"
)

type Weather struct {
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	WindSpeed   float64 `json:"windSpeed"`
	Temperature float64 `json:"temperature"`
	Rain        float64 `json:"rain"`
}

func randFloats(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func GenerateData() *Weather {
	return &Weather{
		Longitude:   randFloats(-180, 180),
		Latitude:    randFloats(-90, 90),
		WindSpeed:   randFloats(0, 100),
		Temperature: randFloats(-80, 80),
		Rain:        randFloats(0, 100),
	}
}
