package middleware

import (
	"net/http"
	"strings"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))


// Struct untuk claim JWT
type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Middleware untuk validasi token JWT
func AuthMiddleware(requiredRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Format header harus: Bearer <token>
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// Parse token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Cek expired manual (biar lebih aman)
		if claims.ExpiresAt.Time.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		// Cek role (kalau middleware dikasih role tertentu)
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

		// Simpan data user di context biar bisa diakses controller
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
