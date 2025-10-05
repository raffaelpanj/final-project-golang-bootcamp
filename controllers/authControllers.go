package controllers

import (
	"database/sql"
	"final-project-golang-bootcamp/connection"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"net/mail"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Struct for claims JWT
type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func RegisterCustomer(c *gin.Context){
	var newUser User
	// Parse input JSON
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}
	// Validate email format
	_, err := mail.ParseAddress(newUser.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email format",
		})
		log.Printf("[ERROR] Failed to parse email: %v", err)
		return
	}

	userId, err := connection.InsertUser(newUser.Name, newUser.Email, newUser.Password, "customer")
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to register user",
		})
		log.Printf("[ERROR] Failed to register user: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user_id": userId,
	})
}
func RegisterAdmin(c *gin.Context){
	var newUser User
	// Parse input JSON
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		log.Printf("[ERROR] Failed Error: %v", err)
		return
	}
	// Validate email format
	_, err := mail.ParseAddress(newUser.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email format",
		})
		log.Printf("[ERROR] Failed to parse email: %v", err)
		return
	}

	userId, err := connection.InsertUser(newUser.Name, newUser.Email, newUser.Password, "admin")
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to register user",
		})
		log.Printf("[ERROR] Failed to register user: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user_id": userId,
	})
}

func LoginUser(c *gin.Context){
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	// Parse input JSON
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		log.Printf("[ERROR] Failed to bind JSON: %v", err)
		return
	}
	user, err := connection.SelectUser(loginData.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login user"})
		log.Printf("[ERROR] Failed to login user: %v", err)
		return
	}
	if user.Password != loginData.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}
	
// Generate JWT token
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
		"token":   tokenString,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
