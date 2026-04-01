package middleware

import (
	"fmt"
	"go-gin-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		fmt.Println("Authorization Header:", authHeader)
		if authHeader == "" {
			utils.Error(c, 401, "authorization header required")
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"error": "authorization header required",
			// })
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Error(c, 401, "invalid authorization format")
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"error": "invaild authorization format",
			// })
			c.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.Error(c, 401, "invalid expired token")
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"error": "invalid expired token",
			// })
			c.Abort()
			return
		}
		//把用户信息存入上下文
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}
