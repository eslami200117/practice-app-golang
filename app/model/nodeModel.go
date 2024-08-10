package model


type Node struct {
    Id       uint   `json:"id"`
    Username string `json:"username"`
    Password string `json:"-"`
	Role 	 string `json:"role"`
}

