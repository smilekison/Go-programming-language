package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
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

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	// fmt.Println("This is user Email details:: ", &user)

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authenticated."})
		return
	}
	fmt.Println("This is email and user id: ", user.Email, user.ID)
	token, err := utils.GenerateToken(user.Email, user.ID)
	fmt.Println("THis is the token ::: ", token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Token error."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
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
