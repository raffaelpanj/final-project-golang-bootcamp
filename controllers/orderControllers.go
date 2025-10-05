package controllers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Alias for Order model
type Order = models.Order
var OrderDatas = [] Order{}

// Alias for OrderRequest model
type OrderRequest = models.OrderRequest
var OrderRequestDatas = [] OrderRequest{}

// CreateOrder handles creating a new order and ticket generation
func CreateOrder(ctx *gin.Context) {
	logger := log.Default()
	var newOrderRequest OrderRequest

	// Parse and bind JSON request body
	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}

	// Retrieve event code from event ID
	eventCode, err := connection.SelectEventCodeById(newOrderRequest.EventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order",
		})
		return
	}

	// Log the event code for debugging
	logger.Println("Event Code:", eventCode)

	// Insert order and related tickets into the database
	orderId, err := connection.InsertOrder(newOrderRequest.UserID, newOrderRequest.EventID, newOrderRequest.TicketCount, newOrderRequest.PaymentMethod, newOrderRequest.TotalPrice, eventCode)
	if err != nil {
		// Handle custom error if event quota is insufficient
		if err.Error() == "not enough quota" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Not enough quota",
			})
			return
		}
		// General database error
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order",
		})
		return
	}

	// Return success response with order ID
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order_id": orderId,
	})
}

// GetOrderByUserId retrieves all orders for a specific user
func GetOrderByUserId(ctx *gin.Context){
	userId := ctx.Query("user_id")
	var orderData []Order
	var err error

	// Fetch orders from the database by user ID
	if orderData, err = connection.SelectOrdersByUserId(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve orders",
		})
		return
	}

	// Return list of orders as JSON response
	ctx.JSON(http.StatusOK, orderData)
}