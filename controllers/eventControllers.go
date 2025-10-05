package controllers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

type Event = models.Event
var EventDatas = [] Event{}

type UpdateEvent = models.UpdateEvent
var UpdateEventDatas = [] UpdateEvent{}

var err error

func CreateEvent(ctx *gin.Context){
	var newEvent Event

	if err := ctx.ShouldBindJSON(&newEvent); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request",
        })
        return
    }

	newEvent.EventID, err = connection.InsertEvent(newEvent.EventCode, newEvent.Name, newEvent.Location, newEvent.Date, newEvent.Quota, newEvent.Description)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
			"error_message": err.Error(),
		})
		return
	}
	EventDatas = append(EventDatas, newEvent)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   newEvent,
	})
}
func GetEventById(ctx *gin.Context){
	eventId := ctx.Param("EventID")
	var eventData Event

	eventData, err = connection.SelectEventById(eventId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("Event with id %s not found", eventId),
		})
		return
	}
		ctx.JSON(http.StatusOK, gin.H{
		"Event": eventData,
	})
}

func UpdateEventById(ctx *gin.Context){
	eventId := ctx.Param("EventID")
	var updatedEvent UpdateEvent
	if err := ctx.ShouldBindJSON(&updatedEvent); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request",
        })
        return
    }
	rowsAffected, err := connection.UpdateEventById(eventId, updatedEvent.Name, updatedEvent.Location, updatedEvent.Date, updatedEvent.Quota, updatedEvent.Description)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
			"error_message": err.Error(),
		})
		return
	}
	if rowsAffected == 99 {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
			"error_message": "Quota cannot be decreased",
		})
		return
	}
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("Event with id %s not found", eventId),
		})
		return
	}
	updatedEvent.EventID = eventId
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"Event": updatedEvent,
	})

}