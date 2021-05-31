package facade

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/morscino/teady-app-2/handlers"
	"github.com/morscino/teady-app-2/service/userservice"
)

type UserFascade struct {
	ctx         context.Context
	userhandler handlers.UserHandler
}

func NewUserFacade(ctx context.Context, u handlers.UserHandler) *UserFascade {

	return &UserFascade{
		ctx:         ctx,
		userhandler: u,
	}
}

func (f UserFascade) CreateUser(c *gin.Context) {
	var input userservice.UserRegistrationData
	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ensure password is valid
	sevenOrMore, number, upper, special := passwordIsValid(input.Password)

	if !sevenOrMore || !number || !upper || !special {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password, password must contain at least 8 characters, at least one letter, one number and one special character"})
		return
	}
	//Ensure password is equal to confirm password
	if input.Password != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password does not match"})
		return
	}
	//Ensure email is valid
	if !emailIsValid(strings.ToLower(input.Email)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}
	//ensure username is unique
	if !f.usernameIsUnique(strings.ToLower(input.Username)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	//ensure email is unique
	if !f.emailIsUnique(strings.ToLower(input.Email)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with email already exists"})
		return
	}
	//ensure phone number is unique
	if !f.phoneIsUnique(strings.ToLower(input.PhoneNumber)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with Phone Number already exists"})
		return
	}
	user, err := f.userhandler.CreateUser(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func (f UserFascade) phoneIsUnique(phone string) bool {
	exists := true

	PhoneNumber := f.userhandler.UserService.GetUserByPhoneNumber(phone)

	if PhoneNumber != "" {
		exists = false
	}

	return exists
}
func (f UserFascade) usernameIsUnique(username string) bool {
	exists := true

	Username := f.userhandler.UserService.GetUserByUsername(username)
	fmt.Println("username logged " + Username)

	if Username != "" {
		exists = false
	}

	return exists
}

func (f UserFascade) emailIsUnique(email string) bool {
	exists := true

	Email := f.userhandler.UserService.GetUserByEmail(email)

	if Email != "" {
		exists = false
	}

	return exists
}

func emailIsValid(email string) bool {
	if len(email) < 4 || len(email) > 254 {
		return false
	}
	return regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9]+\\.(?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(email)
}

func passwordIsValid(password string) (sevenOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	sevenOrMore = letters >= 8
	return
}
