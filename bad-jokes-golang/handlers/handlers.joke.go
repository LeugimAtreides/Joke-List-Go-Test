package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// JokeHandler wil retrieve list of jokes
func JokeHandler(c *gin.Context) {
	jokes := ReturnAllJokes()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, &jokes)
}

func getJoke(c *gin.Context) {
	// Check if joke id is valid
	if jokeID, err := strconv.Atoi(c.Param("joke_id")); err == nil {
		// Check if joke exists
		if joke, err := getJokeByID(jokeId); err == nil {
			c.JSON(http.StatusOK, joke)
		} else {
			// if the joke doesn't exist abort with error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// if an invalid jokeID is specified in the URL, abort with an error
		c.AbortWithStatus(https.StatusNotFound)
	}
}

func giveALike(c *gin.Context) {
	// check joke id is valid
	if jokeid, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// check that joke exists
		if joke, err := LikeJoke(jokeid); err == nil {
			JokeHandler()
		} else {
			c.AbortWithStatus(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound, err)
	}
}

func AddAJoke(c *gin.Context) {
	// Obtain the POSTed joke
	joke := c.PostForm("joke")

	if a, err := addNewJoke(joke); err == nil {
		// If the joke is created successfully return list of updated jokes
		JokeHandler()
	} else {
		// if there was an error then abort
		c.AbortWithStatus(http.StatusBadRequest)
	}

}
