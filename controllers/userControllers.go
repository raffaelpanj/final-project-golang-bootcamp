package controllers

import (
	// "fmt"
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Alias for User model
type User = models.User

// Temporary in-memory data (not used directly with database)
var UserDatas = []User{}

// GetUserbyRole handles GET request to retrieve all users by their role
func GetUserbyRole(ctx *gin.Context){
    role := ctx.Param("UserRole")

	// Validate that role parameter is provided
    if role == "" {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error_status":  "Bad Request",
            "error_message": "role is required",
        })
        return
    }

	// Fetch all users from database filtered by role
	users, err := connection.SelectAllUsersByRole(role)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
			"error_message": err.Error(),
		})
		return
	}

	// Handle case when no users are found
	if users == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": "No users found with the specified role",
		})
		return
	}

	// Mask passwords before sending the response
	for i := range users {
		users[i].Password = "****"
	}

	// Return success response with user data
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}