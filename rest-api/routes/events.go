package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	// Call GetAllEvents() to fetch all events from the database
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. Try again later !!"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later !!"})
		return
	}
	context.JSON(http.StatusOK, event)

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
	userId := context.GetInt64("userId")
	event.UserID = userId

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"Message": "event Created", "event": event})
}

func updateEvent(context *gin.Context) {
	fmt.Println("This is update event.")
	// Call GetAllEvents() to fetch all events from the database
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	fmt.Println("This is update event ID:::", eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. Try again later !!"})
		return
	}
	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later !!"})
		return
	}
	var updatedEvent models.Event

	err = context.ShouldBindBodyWithJSON(&updatedEvent)
	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data"})

	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not Update the data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event updated successfully!! "})
}

func deleteEvent(context *gin.Context) {
	fmt.Println("This is delete event.")
	// Call GetAllEvents() to fetch all events from the database
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later !!"})
		return
	}

	fmt.Println("This is delete event ID:::", eventId)

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later !!"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not Update the data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event deleted successfully!! "})
}
