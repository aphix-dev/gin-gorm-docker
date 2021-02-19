package models

type User struct {
	ID       int    `gorm:"primary_id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func NewUser() *User {
	return &User{}
}
