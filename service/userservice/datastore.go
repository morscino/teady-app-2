package userservice

import (
	"github.com/morscino/teady-app-2/models/usermodel"
)

func NewUserService(database interface{}) UserRepository {
	return &UserService{db: database}
}

func (u UserService) Create(user usermodel.User) usermodel.User {
	//u.users = append(u.users, user)

	return user

}
