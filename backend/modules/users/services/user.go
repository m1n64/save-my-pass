package services

import (
	"backend/modules/users/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// CreateUser creates a new user with the given name, email, and password.
//
// Parameters:
// - name: the name of the user.
// - email: the email of the user.
// - password: the password of the user.
//
// Returns:
// - models.Token: the token generated for the user.
// - error: any error that occurred during user creation.
func (u *UserService) CreateUser(name, email, password string) (models.Token, error) {
	userModel := models.UserModel{DB: u.DB}
	userToken, err := userModel.CreateUser(name, email, password)
	if err != nil {
		return models.Token{}, err
	}

	return userToken, nil
}

// LoginUser is a function that allows a user to log in.
//
// It takes two parameters: email (string) and password (string).
// It returns a Token (models.Token) and an error (error).
func (u *UserService) LoginUser(email, password string) (models.Token, error) {
	userModel := models.UserModel{DB: u.DB}
	userToken, err := userModel.LoginUser(email, password)

	if err != nil {
		return models.Token{}, err
	}

	return userToken, nil
}

// GetUserByToken retrieves a user by their token.
//
// token: The token of the user.
//
// returns:
//   - The user's token.
//   - An error if the token is invalid or the user does not exist.
func (u *UserService) GetUserByToken(token string) (models.Token, error) {
	userModel := models.UserModel{DB: u.DB}

	return userModel.CheckToken(token)
}
