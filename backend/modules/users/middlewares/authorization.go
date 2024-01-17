package middlewares

import (
	"backend/modules/users/services"
	services2 "backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware is a middleware function that handles authorization for API endpoints.
//
// It checks if the 'Authorization' header is present in the request. If not, it aborts the request and returns a JSON response with an error message.
//
// If the 'Authorization' header is present, it checks if it starts with 'Bearer '. If not, it aborts the request and returns a JSON response with an error message.
//
// If the 'Authorization' header starts with 'Bearer ', it extracts the token from the header and queries the user service to validate the token.
//
// If the token is valid, it sets the 'user' key in the gin.Context with the token and continues to the next middleware or route handler.
//
// Parameters:
//   - c: A pointer to the gin.Context object representing the current HTTP request.
//
// Return:
//   - gin.HandlerFunc: A function that handles the request and response for the API endpoint.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		// Проверяем, начинается ли заголовок с 'Bearer '
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}

		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
		userService := services.UserService{DB: services2.GetDBConnection()}

		token, err := userService.GetUserByToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("user", token)

		c.Next()
	}
}
