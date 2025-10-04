package controllers

import (
	// "fmt"
	"final-project-golang-bootcamp/connection"
	"final-project-golang-bootcamp/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User = models.User
var UserDatas = []User{}

func GetUserbyRole(ctx *gin.Context){
    role := ctx.Param("UserRole")
    if role == "" {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error_status":  "Bad Request",
            "error_message": "role is required",
        })
        return
    }
	users, err := connection.SelectAllUsersByRole(role)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status": "Internal Server Error",
			"error_message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}