package controllers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"net/http"
	"github.com/gin-gonic/gin"
	// "log"
)
type Queue = models.Queue
var QueueDatas = [] Queue{}

func CreateQueue(ctx *gin.Context) {
	// logger := log.Default()
	var newQueue Queue
	var returnQueue Queue

	if err := ctx.ShouldBindJSON(&newQueue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}
	returnQueue, err := connection.InsertQueue(newQueue.UserID, newQueue.EventID, newQueue.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create queue",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Queue created successfully",
		"Queue Data": returnQueue,
	})
}