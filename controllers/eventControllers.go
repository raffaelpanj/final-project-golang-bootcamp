package controllers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Alias for Event model
type Event = models.Event
var EventDatas = [] Event{}

// Alias for UpdateEvent model
type UpdateEvent = models.UpdateEvent
var UpdateEventDatas = [] UpdateEvent{}

var err error

// CreateEvent handles creating a new event
func CreateEvent(ctx *gin.Context){
	var newEvent Event

	// Parse and bind JSON input to Event struct
	if err := ctx.ShouldBindJSON(&newEvent); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Bad request",
        })
		log.Printf("[ERROR] Failed Error: %v", err)
        return
    }

	// Validate date format (must be YYYY-MM-DD)
	layout := "2006-01-02"
	_, err := time.Parse(layout, newEvent.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format, must be YYYY-MM-DD"})
		return
	}

	// Insert new event into the database
	newEvent.EventID, err = connection.InsertEvent(newEvent.EventCode, newEvent.Name, newEvent.Location, newEvent.Date, newEvent.Quota, newEvent.Description)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
			"error_message": "Please use unique event code or valid Quota",
		})
		log.Printf("[ERROR] Failed to insert event: %v", err)
		return
	}

	// Set created_at timestamp as string
	timeNow := time.Now()
	timeNowStr := timeNow.Format("2006-01-02 15:04:05")
	newEvent.CreatedAt = timeNowStr

	// Add event to in-memory slice (optional cache)
	EventDatas = append(EventDatas, newEvent)

	// Return success response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   newEvent,
	})
}

// GetEventById retrieves an event by its ID
func GetEventById(ctx *gin.Context){
	eventId := ctx.Param("EventID")
	var eventData Event

	// Fetch event from the database
	eventData, err = connection.SelectEventById(eventId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}

	// Return event data
	ctx.JSON(http.StatusOK, gin.H{
		"Event": eventData,
	})
}

// UpdateEventById updates event data based on its ID
func UpdateEventById(ctx *gin.Context){
	eventId := ctx.Param("EventID")
	var updatedEvent UpdateEvent

	// Parse and bind JSON input
	if err := ctx.ShouldBindJSON(&updatedEvent); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request",
        })
        return
    }

	// Execute update query
	rowsAffected, err := connection.UpdateEventById(eventId, updatedEvent.Name, updatedEvent.Location, updatedEvent.Date, updatedEvent.Quota, updatedEvent.Description)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
			"error_message": "Bad Request or Data Not Found",
		})
		return
	}

	// Handle specific update constraints (e.g., cannot decrease quota)
	if rowsAffected == 99 {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
			"error_message": "Quota cannot be decreased",
		})
		return
	}

	// Handle case where no rows were updated
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}

	// Return updated event info
	updatedEvent.EventID = eventId
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"Event": updatedEvent,
	})
}
