package model

import "time"


type Node struct {
    // Id          uint        `json:"id"`
    Username    string      `json:"username"`
    Password    string      `json:"-"`
	Role 	    string      `json:"role"`
    Status      bool        `json:"status"`
    LastUpdate  time.Time    `json:"lastUpdate"`
}

