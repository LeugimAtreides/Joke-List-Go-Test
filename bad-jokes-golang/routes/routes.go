package routes

import (
	"github.com/LeugimAtreides/Joke-List-Go-Test/bad-jokes-golang/middleware"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func initializeRoutes() {

	router := gin.Default()

	// setUserStatus middleware for every route to set flag for authentication
	router.Use(middleware.SetUserStatus())

	// Handle the index router
	router.Use(static.Serve("/", static.LocalFile("../../bad-jokes/build", true)))

}
