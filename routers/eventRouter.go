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
	router.POST("/event", controllers.CreateEvent)
	router.GET("/event/:EventID", controllers.GetEventById)
	router.PUT("/event/:EventID", controllers.UpdateEvent)
	router.POST("/createOrder", controllers.CreateOrder)
	router.POST("/createQueue", controllers.CreateQueue)
	// router.POST("/event", controllers.CreateEvent)
	// router.GET("/event/", controllers.GetAllEvents)
	// router.GET("/event/:EventID", controllers.GetEventsById)
	// router.PUT("/event/:EventID", controllers.UpdateEvent)
	// router.DELETE("/event/:EventID", controllers.DeleteEvent)
	return router
}