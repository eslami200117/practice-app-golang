package entities

import (
	"time"

)

type (
	Nodes struct {
		Id 			uint32 		`gorm:"primaryKey;autoIncrement" json:"id"`
		Username 	string 		`json:"username"`
		Password 	string 		`json:"password"`
		Role 		string 		`json:"role"`
		Status 		bool 		`json:"status"`
		CreateAt 	time.Time 	`json:"createAt"`
		UpdateAt	time.Time 	`json:"updateAt"`
	}
)