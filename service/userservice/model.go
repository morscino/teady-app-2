package userservice

import (
	"github.com/morscino/teady-app-2/models/usermodel"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user usermodel.User) (usermodel.User, error)
	GetUserByPhoneNumber(phone string) string
	GetUserByUsername(username string) string
	GetUserByEmail(email string) string
}

type UserService struct {
	db *gorm.DB
}

type UserRegistrationData struct {
	FirstName       string `json:"firstname" binding:"required"`
	LastName        string `json:"lastname" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	PhoneNumber     string `json:"phone_number"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}
