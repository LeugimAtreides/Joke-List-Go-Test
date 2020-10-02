package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// EnsureLoggedIn ensures that requests are aborted with error
func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there is an error or the token is empty
		// the user is not logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			// if token, err := c.Cookie("token"); err == nil || token != ""
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// SetUserStatus sets whether the user is logged in or not
func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
