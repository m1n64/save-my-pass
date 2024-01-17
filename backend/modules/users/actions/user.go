package actions

import (
	"backend/modules/users/models"
	"backend/modules/users/services"
	services2 "backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserTokenResponse struct {
	Token string `json:"token"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserRegister is a function that handles the registration of a user.
//
// It accepts a *gin.Context parameter.
// It does not return any values.
// @Summary Register user
// @Description Register user by Email, Name and Password
// @Accept  json
// @Produce  json
// @Param   userRegisterRequest  body    UserRegisterRequest  true  "User Registration"
// @Success 200 {object} UserTokenResponse	"{"token": "jakjdslskldaew"}"
// @Failure 400 {object} services2.ErrorResponse "{"error": "error"}"
// @Failure 500 {object} services2.ErrorResponse "{"error": "error"}"
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	var request UserRegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := services.UserService{DB: services2.GetDBConnection()}
	token, err := userService.CreateUser(request.Name, request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserTokenResponse{Token: token.Token})
}

// UserLogin handles the user login functionality.
//
// It expects a JSON object containing the user's email and password as the request body.
// It returns a JSON response with the user's authentication token if the login is successful.
// Otherwise, it returns an error response with the appropriate status code.
// @Summary User login
// @Description Log in a user using email and password
// @Accept  json
// @Produce  json
// @Param   userLoginRequest  body    UserLoginRequest  true  "User Login"
// @Success 200 {object} UserTokenResponse 	"{"token": "jakjdslskldaew"}"
// @Failure 400 {object} services2.ErrorResponse "{"error": "error"}"
// @Failure 500 {object} services2.ErrorResponse "{"error": "error"}"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var request UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := services.UserService{DB: services2.GetDBConnection()}

	token, err := userService.LoginUser(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserTokenResponse{Token: token.Token})
}

// GetUser retrieves the user information from the context and returns it as a JSON response.
//
// Parameters:
//
//	c: a pointer to the gin.Context object.
//
// Return type:
//
//	None.
//
// @Summary Get user info
// @Description Get information about the user
// @Accept  json
// @Produce  json
// @Param   Authorization  header  string  true  "Authorization Bearer Token"
// @Success 200 {object} GetUserResponse "{ 'id': 6, 'name': 'admin', 'email': 'admin@gmail.com' }"
// @Router /user [get]
func GetUser(c *gin.Context) {
	user := c.MustGet("user").(models.Token)
	c.JSON(http.StatusOK, GetUserResponse{
		ID:    user.User.ID,
		Name:  user.User.Name,
		Email: user.User.Email,
	})
}
