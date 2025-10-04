package routers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/controllers"
	"github.com/gin-gonic/gin"
)
func StartServer() *gin.Engine {
	router := gin.Default()
	connection.ConnectDB()
	router.GET("/users/:UserRole", controllers.GetUserbyRole)

	// Event Routes
	router.POST("/event", controllers.CreateEvent)
	router.GET("/event/:EventID", controllers.GetEventById)
	router.PUT("/event/:EventID", controllers.UpdateEventById)

	// Order Routes
	router.POST("/createOrder", controllers.CreateOrder)
	router.GET("/order", controllers.GetOrderByUserId)

	// Queue Routes
	router.POST("/createQueue", controllers.CreateQueue)
	router.PUT("/queue/:QueueID", controllers.UpdateQueueById)
	router.GET("/queue/:QueueID", controllers.GetQueueById)
	router.GET("/queue", controllers.GetQueuesByEventIdAndStatus)

	// Ticket Routes
	router.GET("/ticket", controllers.GetTicketByOrderId)
	router.GET("/ticket/:TicketID", controllers.GetTicketById)
	return router
}