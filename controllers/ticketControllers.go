package controllers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

// Alias for Ticket model
type Ticket = models.Ticket
var TicketDatas = [] Ticket{}

// GetTicketByOrderId handles GET request to retrieve tickets by order ID
func GetTicketByOrderId(ctx *gin.Context){
	orderId := ctx.Query("order_id")
	var ticketData []Ticket

	// Retrieve tickets from database using order ID
	ticketData, err = connection.SelectTicketByOrderId(orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve tickets, check your order_id",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}

	// Handle case when no ticket is found
	if ticketData == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Ticket not found",
		})
		return
	}

	// Return success response with ticket data
	ctx.JSON(http.StatusOK, gin.H{
		"tickets": ticketData,
	})
}

// GetTicketById handles GET request to retrieve a single ticket by its ID
func GetTicketById(ctx *gin.Context){
	ticketId := ctx.Param("TicketID")
	var ticketData Ticket

	// Retrieve ticket from database using ticket ID
	ticketData, err = connection.SelectTicketById(ticketId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve ticket, check your ticket_id",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}

	// Return success response with ticket data
	ctx.JSON(http.StatusOK, gin.H{
		"ticket": ticketData,
	})
}
