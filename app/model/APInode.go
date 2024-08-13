package model

import "time"

type APInode struct {
	Username 	string
	Status 		bool
	LastUpdate	time.Time
}