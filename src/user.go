package main

type User struct {
	Id       int       `json:"id"`
	FullName string    `json:"fullName"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
}

type Users []User