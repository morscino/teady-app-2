package usermodel

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:ID`
	FirstName   string    `json : FirstName`
	LastName    string    `json : LastName`
	Username    string    `json : Username`
	Email       string    `json:Email`
	PhoneNumber string    `json:PhoneNumber`
}
