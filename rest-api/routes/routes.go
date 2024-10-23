package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// Routes for events
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// Routes for Users
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/users", getUsers)

}
