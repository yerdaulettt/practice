package utils

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func GenerateJWT(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
			// c.Next()
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) { return jwtSecret, nil })

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			// c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			// c.Next()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)

		c.Set("userID", claims["user_id"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
