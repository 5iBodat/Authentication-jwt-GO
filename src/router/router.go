package router

import (
	Auth "authentication-jwt-go/src/controller/auth"
	"authentication-jwt-go/src/controller/home"
	"authentication-jwt-go/src/middleware"

	"github.com/gin-gonic/gin"
)

// Router sets up the routes and returns the router engine
func Router() *gin.Engine {
	route := gin.Default()

	// Define your routes
	route.POST("/login", Auth.LoginHandler)

	//Protected routes
	authorized := route.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/home", home.Home)
	}
	// Return the router
	return route
}
