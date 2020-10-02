package main

import "github.com/gin-gonic/contrib/static"

func initializeRoutes() {

	// setUserStatus middleware for every route to set flag for authentication
	router.Use(setUserStatus())

	// Handle the index router
	router.Use(static.Serve("/", static.LocalFile("../../bad-jokes/build", true)))

	// Group user related routes
	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", ensureNotLoggedIn())

		userRoutes.POST("/login", ensureNotLoggedIn())

		userRoutes.GET("/logout", ensureLoggedIn())

		userRoutes.GET("/register", ensureNotLoggedIn())

		userRoutes.GET("/register", ensureNotLoggedIn())
	}
}
