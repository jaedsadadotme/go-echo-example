package main

type User struct {
	Id    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
