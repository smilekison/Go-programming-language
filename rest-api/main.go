package main

import (
	"fmt"
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	// Call GetAllEvents() to fetch all events from the database
	events, err := models.GetAllEvents()

	// Log the result for debugging
	fmt.Println("These are the events fetched: ", events)

	// If there is an error, return a 500 response with a message
	if err != nil {
		fmt.Println("Error fetching events: ", err) // Log the error for debugging
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events."})
		return
	}

	// Return the events as a JSON response with status code 200
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindBodyWithJSON(&event)
	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data"})

	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"Message": "event Created", "event": event})
}
