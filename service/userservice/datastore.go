package userservice

import (
	"github.com/morscino/teady-app-2/models/usermodel"
	"gorm.io/gorm"
)

func NewUserService(database *gorm.DB) UserRepository {
	return &UserService{db: database}
}

func (u UserService) Create(user usermodel.User) (usermodel.User, error) {

	result := u.db.Create(&user)

	if result.Error != nil {

		return user, result.Error

	}

	return user, nil

}

func (u UserService) GetUserByPhoneNumber(phone string) string {
	var userFound usermodel.User
	u.db.Where("phone_number=?", phone).Find(&userFound)

	return userFound.PhoneNumber
}

func (u UserService) GetUserByUsername(username string) string {
	var userFound usermodel.User
	u.db.Where("username=?", username).Find(&userFound)

	return userFound.Username
}

func (u UserService) GetUserByEmail(email string) string {
	var userFound usermodel.User
	u.db.Where("email=?", email).Find(&userFound)

	return userFound.Email
}
