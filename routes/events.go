package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"arsh.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get data"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not pass a value"})
		return
	}

	user := context.GetInt64("userId")
	event.UserId = user

	err = event.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create data", "error": err})
		return
	}

	context.JSON(http.StatusCreated, event)
}

func getEvent(context *gin.Context) {
	evendId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find the data"})
		return
	}

	event, err := models.GetEventById(evendId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find the data"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	ev, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event"})
		return
	}

	if ev.UserId != context.GetInt64("userId") {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not autherized to change event"})
		return
	}
	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated succesfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event"})
		return
	}

	if event.UserId != context.GetInt64("userId") {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not autherized to delete event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Your mentioned event  has been deleted"})
}
