package entities

import "time"

type User struct {
	Username 	string 		`gorm:"primaryKey;autoIncrement" json:"id"`
	Password 	string 		`json:"password"`
	Role 		string 		`json:"role"`
	CreateAt 	time.Time 	`json:"createAt"`
	LastLogin	time.Time	`json:"lastLogin"`
}