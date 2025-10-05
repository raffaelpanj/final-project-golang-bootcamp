package controllers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}
	eventCode, err := connection.SelectEventCodeById(newOrderRequest.EventID)
	if err != nil {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": "Failed to create order",
		})
		return
	}
	logger.Println("Event Code:", eventCode) // Log the retrieved event code
	orderId, err := connection.InsertOrder(newOrderRequest.UserID, newOrderRequest.EventID, newOrderRequest.TicketCount, newOrderRequest.PaymentMethod, newOrderRequest.TotalPrice, eventCode)
	if err != nil {
		if err.Error() == "not enough quota" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Not enough quota",
			})
			return
		}
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

func GetOrderByUserId(ctx *gin.Context){
	userId := ctx.Query("user_id")
	var orderData []Order
	var err error
	if orderData, err = connection.SelectOrdersByUserId(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve orders",
		})
		return
	}
	ctx.JSON(http.StatusOK, orderData)
}