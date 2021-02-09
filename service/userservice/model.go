package userservice

import "github.com/morscino/teady-app-2/models/usermodel"

type UserRepository interface {
	Create(user usermodel.User) usermodel.User
}

type UserService struct {
	db interface{}
}

type UserRegistrationData struct {
	FirstName       string `json:"firstname" binding:"required"`
	LastName        string `json:"lastname" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	PhoneNumber     string `json:"phone_number"`
	Password        string `json:"password" `
	ConfirmPassword string `json:"confirm_password"`
}
