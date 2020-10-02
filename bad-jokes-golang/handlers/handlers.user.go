package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LeugimAtreides/Joke-List-Go-Test/bad-jokes-golang/models"
	"github.com/joho/godotenv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GoDotEnvVariable retrieves an env variable
func GoDotEnvVariable(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// CreateToken creates a session token
func CreateToken(username string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(GoDotEnvVariable("ACCESS_SECRET")))
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

	// check that combination is valid
	if models.IsUserValid(username, password) {
		// if valid then set token in a cookie
		if token, err := CreateToken(username); err == nil {
			c.SetCookie("token", token, 3600, "", "", false, true)
			c.Set("is_logged_in", true)
			c.JSON(http.StatusAccepted, "Log in successfull!")
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	} else {
		// respond with error message if invalid user login
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// logout the user
func logout(c *gin.Context) {
	// clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)
	c.JSON(http.StatusAccepted, "successfully logged out!")

}

func register(c *gin.Context) {
	// obtain values from POST
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := models.RegisterNewUser(username, password); err == nil {
		// if the user is created, set the token in a cookie and log the user in
		if token, err := CreateToken(username); err == nil {
			c.SetCookie("token", token, 3600, "", "", false, true)
			c.Set("is_logged_in", true)
			c.JSON(http.StatusAccepted, "Login Successful!")
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
