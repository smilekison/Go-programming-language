package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Save user data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successully"})
}

func getUsers(context *gin.Context) {
	// Call GetAllEvents() to fetch all events from the database
	users, err := models.GetAllUsers()

	// Log the result for debugging
	fmt.Println("These are the events fetched: ", users)

	// If there is an error, return a 500 response with a message
	if err != nil {
		fmt.Println("Error fetching users: ", err) // Log the error for debugging
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users."})
		return
	}

	// Return the events as a JSON response with status code 200
	context.JSON(http.StatusOK, users)
}
