package handlers

import (
	"net/http"
	"strconv"

	"github.com/LeugimAtreides/Joke-List-Go-Test/bad-jokes-golang/models"

	"github.com/gin-gonic/gin"
)

// JokeHandler wil retrieve list of jokes
func JokeHandler(c *gin.Context) {
	jokes := models.ReturnAllJokes()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, &jokes)
}

// GetJoke will retrieve a joke based on id
func GetJoke(c *gin.Context) {
	// Check if joke id is valid
	if jokeID, err := strconv.Atoi(c.Param("joke_id")); err == nil {
		// Check if joke exists
		if joke, err := models.GetJokeByID(jokeID); err == nil {
			c.JSON(http.StatusOK, joke)
		} else {
			// if the joke doesn't exist abort with error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// if an invalid jokeID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// GiveALike will attempt to increase the number of likes a joke has
func GiveALike(c *gin.Context) {
	// check joke id is valid
	if jokeid, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// check that joke exists
		if joke, err := models.LikeJoke(jokeid); err == nil {
			c.JSON(http.StatusAccepted, &joke)
			JokeHandler(c)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// AddAJoke will run the http request to add a new joke to the jokes list
func AddAJoke(c *gin.Context) {
	// Obtain the POSTed joke
	joke := c.PostForm("joke")

	if a, err := models.AddNewJoke(joke); err == nil {
		c.JSON(http.StatusAccepted, &a)
		// If the joke is created successfully return list of updated jokes
		JokeHandler(c)
	} else {
		// if there was an error then abort
		c.AbortWithStatus(http.StatusBadRequest)
	}

}
