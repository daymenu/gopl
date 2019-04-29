package models

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

