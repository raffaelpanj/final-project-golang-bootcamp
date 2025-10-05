package middleware

import (
	"net/http"
	"strings"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Retrieve JWT secret key from environment variables
var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims struct defines the payload structure of the JWT token
type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// AuthMiddleware validates JWT tokens and (optionally) checks user roles
func AuthMiddleware(requiredRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// The header format must be: Bearer <token>
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// Parse the JWT token and extract claims
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Manually check expiration for extra security
		if claims.ExpiresAt.Time.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		// Check if user role matches one of the required roles (if provided)
		if len(requiredRole) > 0 {
			match := false
			for _, role := range requiredRole {
				if claims.Role == role {
					match = true
					break
				}
			}
			if !match {
				c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Access Denied"})
				c.Abort()
				return
			}
		}

		// Store user data in context so it can be accessed by controllers
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		// Continue to the next middleware or handler
		c.Next()
	}
}