package entities

import (
	"time"

)

type (
	Nodes struct {
		Username 	string 		`json:"username" gorm:"primaryKey"`
		Password 	string 		`json:"password"`
		Role 		string 		`json:"role"`
		Status 		bool 		`json:"status"`
		CreateAt 	time.Time 	`json:"createAt"`
		LastUpdata	time.Time 	`json:"lastUpdata"`
	}
)