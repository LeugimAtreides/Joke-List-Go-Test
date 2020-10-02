package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateToken creates a session token
func CreateToken(userid uint64) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(goDotEnvVariable("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Login the user
func performLogin(c *gin.Context) {
	// obtain the POSTed values
	username := c.PostForm("username")
	password := c.PostForm("password")

	var sameSiteCookie http.SameSite

	// check that combination is valid

	if isUserValid(username, password) {
		// if valid then set token in a cookie
		token := CreateToken(username.ID)
		c.SetCookie("token", token, 3600, "", "", sameSiteCookie, false, true)
		c.Set("is_logged_in", true)
		c.JSON(http.StatusAccepted, "Log in successfull!")
	} else {
		// respond with error message if invalid user login
		c.AbortWithStatus(http.StatusUnauthorized, "Please provide valid login details")
	}
}

// logout the user
func logout(c *gin.Context) {
	var sameSiteCookie http.SameSite

	// clear the cookie
	c.SetCookie("token", "", -1, "", "", sameSiteCookie, false, true)
	c.JSON(http.StatusAccepted, "successfully logged out!")

}

func register(c *gin.Context) {
	// obtain values from POST
	username := c.PostForm("username")
	password := c.PostForm("password")

	var sameSiteCookie http.SameSite

	if _, err := registerNewUser(username, password); err == nil {
		// if the user is created, set the token in a cookie and log the user in
		token := CreateToken(username.ID)
		c.SetCookie("token", token, 3600, "", "", sameSiteCookie, false, true)
		c.Set("is_logged_in", true)
		c.JSON(http.StatusAccepted, "Login Successful!")
	} else {
		c.AbortWithStatus(http.StatusBadRequest, "An unexpected error occurred")
	}
}
