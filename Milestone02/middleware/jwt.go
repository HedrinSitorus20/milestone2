package middleware

import (
	"Milestone02/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func SuperadminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.ExtractToken(c)
		claims, b := utils.ExtractClaims(token)

		if !b {
			c.JSON(400, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		roleID_64 := claims["role_id"]
		roleID := int(roleID_64.(float64))

		if roleID == 2 {
			//bukan superadmin
			fmt.Println(claims["role_id"])
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You're not a Superadmin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
