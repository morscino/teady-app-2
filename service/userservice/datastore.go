package userservice

import "github.com/morscino/teady-2-app/models/usermodel"

func NewUserService() UserRepository {
	return &UserService{}
}

func (u UserService) Create(user usermodel.User) usermodel.User {
	u.users = append(u.users, user)
	return user
}
