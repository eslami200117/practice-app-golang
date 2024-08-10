package entities

import (
	"time"
)

type (
	NodeEntity struct {
		Id 			uint32 		`gorm:"primaryKey;autoIncrement" json:"id"`
		Username 	string 		`json:"username"`
		Password 	string 		`json:"password"`
		Role 		string 		`json:"role"`
		CreateAt 	time.Time 	`json:"createAt"`
	}
)