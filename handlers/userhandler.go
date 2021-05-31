package handlers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/morscino/teady-app-2/models/usermodel"
	"github.com/morscino/teady-app-2/service/userservice"
	"github.com/morscino/teady-app-2/utility/password"
)

type UserHandler struct {
	UserService userservice.UserRepository
}

func NewUserHandler(u userservice.UserRepository) UserHandler {
	return UserHandler{UserService: u}
}

func (u UserHandler) CreateUser(input userservice.UserRegistrationData) (usermodel.User, error) {
	//hash password
	hashedPassword, err := password.NewPasswordHash(input.Password)
	if err != nil {
		fmt.Println(err)
	}
	user := usermodel.User{
		ID:          uuid.New(),
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Username:    input.Username,
		Password:    hashedPassword.Hash,
		Salt:        hashedPassword.Salt,
		CreatedAt:   time.Now(),
	}
	NewUser, err := u.UserService.Create(user)

	if err != nil {
		//log error and notify facade
		return NewUser, err

	}

	return NewUser, nil
}
