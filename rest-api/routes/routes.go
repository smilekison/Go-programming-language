package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Routes for events
	server.POST("/events", createEvent)
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	// Routes for Users
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/users", getUsers)

}
