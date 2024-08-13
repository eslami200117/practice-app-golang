package model

import "time"

type APIuser struct {
	Username	string
	LastLogin	time.Time
}