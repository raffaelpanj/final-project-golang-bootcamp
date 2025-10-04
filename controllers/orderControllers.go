package controllers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
)
type Order = models.Order
var OrderDatas = [] Order{}

type OrderRequest = models.OrderRequest
var OrderRequestDatas = [] OrderRequest{}

func CreateOrder(ctx *gin.Context) {
	logger := log.Default()
	var newOrderRequest OrderRequest

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}
	eventCode, err := connection.GetEventCode(newOrderRequest.EventID)
	if err != nil {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": "Failed to create order",
		})
		return
	}
	logger.Println("Event Code:", eventCode) // Log the retrieved event code
	orderId, err := connection.InsertOrder(newOrderRequest.UserID, newOrderRequest.EventID, newOrderRequest.TicketCount, newOrderRequest.PaymentMethod, newOrderRequest.TotalPrice, eventCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order_id": orderId,
	})
}