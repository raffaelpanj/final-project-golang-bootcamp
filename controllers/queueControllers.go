package controllers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
)
type Queue = models.Queue
var QueueDatas = [] Queue{}

type UpdateQueue = models.UpdateQueue
var UpdateQueueDatas = [] UpdateQueue{}

func CreateQueue(ctx *gin.Context) {
	var newQueue Queue
	var returnQueue Queue

	if err := ctx.ShouldBindJSON(&newQueue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}
	returnQueue, err := connection.InsertQueue(newQueue.UserID, newQueue.EventID, newQueue.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create queue",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Queue created successfully",
		"Queue Data": returnQueue,
	})
}
func UpdateQueueById(ctx *gin.Context) {
	queueId := ctx.Param("QueueID")
	var updatedQueue UpdateQueue

	if err := ctx.ShouldBindJSON(&updatedQueue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}
	rowsAffected, err := connection.UpdateQueueById(queueId, updatedQueue.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update queue, check your status value",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("queue with id %s not found", queueId),
		})
		return
	}
	updatedQueue.QueueID = queueId
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Queue updated successfully",
		"queue":   updatedQueue,
	})
}

func GetQueueById(ctx *gin.Context){
	queueId := ctx.Param("QueueID")
	var queueData Queue
	var err error
	queueData, err = connection.GetQueueById(queueId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve queue",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"queue": queueData,
	})
}
func GetQueuesByEventIdAndStatus(ctx *gin.Context){
	eventID := ctx.Query("event_id")
    status := ctx.Query("status")
	if eventID == "" || status == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request"})
        return
    }
	queues, err := connection.GetQueueByEventIdAndStatus(eventID, status)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
			"error_message": err.Error(),
		})
		return
	}
	if queues == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Bad Request or Data Not Found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"queues": queues,
	})
}