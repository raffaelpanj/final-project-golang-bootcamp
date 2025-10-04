package controllers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ticket = models.Ticket
var TicketDatas = [] Ticket{}

func GetTicketByOrderId(ctx *gin.Context){
	orderId := ctx.Query("order_id")
	var ticketData []Ticket
	ticketData, err = connection.SelectTicketByOrderId(orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve tickets",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"tickets": ticketData,
	})
}

func GetTicketById(ctx *gin.Context){
	ticketId := ctx.Param("TicketID")
	var ticketData Ticket
	ticketData, err = connection.SelectTicketById(ticketId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve ticket",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ticket": ticketData,
	})
}