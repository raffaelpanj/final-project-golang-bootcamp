package controllers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
)

// Alias for Queue model
type Queue = models.Queue
var QueueDatas = [] Queue{}

// Alias for UpdateQueue model
type UpdateQueue = models.UpdateQueue
var UpdateQueueDatas = [] UpdateQueue{}

// CreateQueue handles the creation of a new queue record
func CreateQueue(ctx *gin.Context) {
	var newQueue Queue
	var returnQueue Queue

	// Parse and bind JSON request body
	if err := ctx.ShouldBindJSON(&newQueue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}

	// Insert queue data into database
	returnQueue, err := connection.InsertQueue(newQueue.UserID, newQueue.EventID, newQueue.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create queue",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}

	// Return success response with created queue data
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Queue created successfully",
		"Queue Data": returnQueue,
	})
}

// UpdateQueueById updates the queue status by its ID
func UpdateQueueById(ctx *gin.Context) {
	queueId := ctx.Param("QueueID")
	var updatedQueue UpdateQueue

	// Parse and validate incoming JSON request
	if err := ctx.ShouldBindJSON(&updatedQueue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// Perform update operation in database
	rowsAffected, err := connection.UpdateQueueById(queueId, updatedQueue.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update queue, check your status value",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}

	// If no rows were affected, return 404
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("queue with id %s not found", queueId),
		})
		return
	}

	// Return success response with updated queue data
	updatedQueue.QueueID = queueId
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Queue updated successfully",
		"queue":   updatedQueue,
	})
}

// GetQueueById retrieves queue data by its ID
func GetQueueById(ctx *gin.Context){
	queueId := ctx.Param("QueueID")
	var queueData Queue
	var err error

	// Fetch queue from database
	queueData, err = connection.GetQueueById(queueId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve queue",
		})
		return
	}

	// Return the queue data as JSON response
	ctx.JSON(http.StatusOK, gin.H{
		"queue": queueData,
	})
}

// GetQueuesByEventIdAndStatus retrieves all queues by event ID and status
func GetQueuesByEventIdAndStatus(ctx *gin.Context){
	eventID := ctx.Query("event_id")
    status := ctx.Query("status")

	// Validate query parameters
	if eventID == "" || status == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request"})
        return
    }

	// Retrieve queues from database based on event ID and status
	queues, err := connection.GetQueueByEventIdAndStatus(eventID, status)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
			"error_message": err.Error(),
		})
		return
	}

	// Handle case where no queue data found
	if queues == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Bad Request or Data Not Found",
		})
		return
	}

	// Return list of queues as JSON response
	ctx.JSON(http.StatusOK, gin.H{
		"queues": queues,
	})
}