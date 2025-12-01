package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Events not found"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Failed to register"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Success registering at event"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Failed to cancel"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success cancelling event"})
}
