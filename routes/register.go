package routes

import (
	"net/http"
	"strconv"

	"arsh.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register an event"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.Cancel(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mesage": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration canceled"})
}
