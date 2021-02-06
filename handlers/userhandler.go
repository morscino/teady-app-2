package handlers

import (
	"github.com/google/uuid"

	"github.com/morscino/teady-2-app/models/usermodel"
	"github.com/morscino/teady-2-app/service/userservice"
)

type UserHandler struct {
	UserService userservice.UserRepository
}

func NewUserHandler(u userservice.UserRepository) UserHandler {
	return UserHandler{UserService: u}
}

func (u UserHandler) CreateUser(input userservice.UserRegistrationData) usermodel.User {

	user := usermodel.User{
		ID:          uuid.New(),
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Username:    input.Username,
	}
	NewUser := u.UserService.Create(user)

	return NewUser
}
