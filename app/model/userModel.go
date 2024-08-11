package model


type User struct {
    Username string `json:"username"`
    Password string `json:"-"`
	Role 	 string `json:"role"`
}