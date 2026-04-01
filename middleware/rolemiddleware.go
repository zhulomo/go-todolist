package middleware

import (
	"fmt"
	"go-gin-api/repository"
	"go-gin-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requireRole string) gin.HandlerFunc {
	return func(c *gin.Context) {

		//role, exists := c.Get("role")
		userIDInterface, exists := c.Get("userID")
		if !exists {
			utils.Error(c, 401, "not user exist")
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"error": "not user exist",
			// })
			c.Abort()
			return
		}
		userID := userIDInterface.(uint)
		fmt.Println("userID:", userID)
		role, err := repository.GetRoleByUserID(repository.DB, userID)

		if err != nil {
			utils.Error(c, 400, "database wrong")
			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"error": "database wrong",
			// })
			c.Abort()
			return
		}
		if role != requireRole {
			utils.Error(c, 401, "permission denied")
			// c.JSON(http.StatusForbidden, gin.H{
			// 	"error": "permission denied",
			// })
			c.Abort()
			return
		}
		c.Next()
	}
}

func OperateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIInterface, exists := c.Get("userID")
		if !exists {
			utils.Error(c, 401, "not user exist")
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"error": "not user exist",
			// })
			c.Abort()
			return
		}
		userID := userIInterface.(uint)
		fmt.Println("userID:", userID)
		taskID, err := strconv.Atoi(c.Param("id"))
		task, err := repository.GetTaskByID(repository.DB, uint(taskID))
		if err != nil {
			utils.Error(c, 404, "taskid doesn't exist")
			// c.JSON(http.StatusNotFound, gin.H{
			// 	"error": "taskid doesn't exist",
			// })
			c.Abort()
			return
		}
		if userID != task.UserID {
			utils.Error(c, 401, "permission denied")
			// c.JSON(http.StatusForbidden, gin.H{
			// 	"error": "you haven't permission",
			// })
			c.Abort()
			return
		}
		c.Next()
	}
}
