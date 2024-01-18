package services

import (
	"backend/modules/users/models"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// GetUserFromContext returns the user token from the given gin context.
//
// The function takes a *gin.Context as a parameter and retrieves the "user" value from it.
// The retrieved value is then type-casted to models.Token and returned.
// The function returns a models.Token.
func GetUserFromContext(c *gin.Context) models.Token {
	user := c.MustGet("user").(models.Token)

	return user
}
