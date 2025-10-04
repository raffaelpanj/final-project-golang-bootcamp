package routers

import (
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/controllers"
	"final-project-golang-bootcamp/middleware"
	"github.com/gin-gonic/gin"
)
func StartServer() *gin.Engine {
	router := gin.Default()
	connection.ConnectDB()

	adminOnly := middleware.AuthMiddleware("admin")
	adminAndCustomer := middleware.AuthMiddleware("admin", "customer")

	// User Routes
	router.GET("/users/:UserRole", adminOnly, controllers.GetUserbyRole)
	router.POST("/users/login", controllers.LoginUser)
	router.POST("/users/registerCustomer", controllers.RegisterCustomer)
	router.POST("/users/registerAdmin", controllers.RegisterAdmin)

	// Event Routes
	router.POST("/event", adminOnly, controllers.CreateEvent)
	router.GET("/event/:EventID", adminAndCustomer, controllers.GetEventById)
	router.PUT("/event/:EventID", adminOnly, controllers.UpdateEventById)

	// Order Routes
	router.POST("/createOrder", adminOnly, controllers.CreateOrder)
	router.GET("/order", adminAndCustomer, controllers.GetOrderByUserId)

	// Queue Routes
	router.POST("/createQueue", adminAndCustomer, controllers.CreateQueue)
	router.PUT("/queue/:QueueID", adminOnly, controllers.UpdateQueueById)
	router.GET("/queue/:QueueID", adminAndCustomer, controllers.GetQueueById)
	router.GET("/queue", adminAndCustomer, controllers.GetQueuesByEventIdAndStatus)

	// Ticket Routes
	router.GET("/ticket", adminAndCustomer, controllers.GetTicketByOrderId)
	router.GET("/ticket/:TicketID", adminAndCustomer, controllers.GetTicketById)

	return router
}