package userservice

import (
	"github.com/morscino/teady-2-app/models/usermodel"
)

type UserRepository interface {
	Create(user usermodel.User) usermodel.User
}

type UserService struct {
	users []usermodel.User
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
