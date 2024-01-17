package models

import (
	"backend/modules/users/services/tokens"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Pin      int    `gorm:"nullable"`
}

type Token struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"`
	Token  string `gorm:"not null"`
}

type UserModel struct {
	DB *gorm.DB
}

// CreateUser creates a new user with the given name, email, and password.
//
// Parameters:
// - name: The name of the user.
// - email: The email address of the user.
// - password: The password of the user.
//
// Returns:
// - Token: The token generated for the user.
// - error: An error if any occurred during the creation process.
func (u *UserModel) CreateUser(name, email, password string) (Token, error) {
	hashedPassword, err := u.hashPassword(password)
	if err != nil {
		return Token{}, err
	}

	user := User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	result := u.DB.Create(&user)
	if result.Error != nil {
		return Token{}, result.Error
	}

	return u.getToken(user)
}

// LoginUser authenticates a user by their email and password.
//
// Parameters:
// - email: the email of the user.
// - password: the password of the user.
//
// Returns:
// - Token: the authentication token generated for the user.
// - error: an error if there was a problem during authentication.
func (u *UserModel) LoginUser(email, password string) (Token, error) {
	var user User

	result := u.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return Token{}, result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return Token{}, err
	}

	return u.getToken(user)
}

// CheckToken checks the validity of a token.
//
// It takes a token string as a parameter and returns a Token object and an error.
func (u *UserModel) CheckToken(token string) (Token, error) {
	var tokenObject Token
	result := u.DB.Preload("User").Where("token = ?", token).First(&tokenObject)
	if result.Error != nil {
		return Token{}, result.Error
	}

	return tokenObject, nil
}

// getToken generates a token for the given user and saves it to the database.
//
// The user parameter is the user for whom the token is being generated.
// It is of type User.
//
// The function returns the generated token and an error, if any.
// The return type is Token and error.
func (u *UserModel) getToken(user User) (Token, error) {
	tokenString := tokens.CreateToken()

	token := Token{
		UserID: user.ID,
		Token:  tokenString,
	}

	tokenResult := u.DB.Create(&token)
	if tokenResult.Error != nil {
		return Token{}, tokenResult.Error
	}

	return token, nil
}

// hashPassword generates a hashed password from the given string.
//
// password: the password to be hashed.
// returns the hashed password as a string.
// returns an error if the generation fails.
func (u *UserModel) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
