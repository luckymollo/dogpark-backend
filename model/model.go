package model

type User struct {
	UserName string `json:"userName" binding: "required`
	Password string `json:"password" binding: "required"`
}
