package auth

import (
	"fmt"
	"net/http"
	"strings"

	authEntity "github.com/antongoncharik/crypto-knight-api/internal/entity/auth"
	"github.com/antongoncharik/crypto-knight-api/pkg/utilities"
	"github.com/gin-gonic/gin"
)

func UseAuth(keys authEntity.RSAKeys) gin.HandlerFunc {
	return func(c *gin.Context) {
		botHeader := c.GetHeader("Bot")
		fmt.Println("botHeader", botHeader)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		err := utilities.ValidateToken(parts[1], keys.PublicKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
