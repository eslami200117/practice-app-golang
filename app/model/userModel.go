package model

import "time"


type User struct {
    Username    string      `json:"username"`
    Password    string      `json:"-"`
	Role 	    string      `json:"role"`
    LastLogin   time.Time   `json:"lastLogin"`
}