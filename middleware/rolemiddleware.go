package middleware

import (
	"fmt"
	"go-gin-api/repository"
	"go-gin-api/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取role并于requiredrole对比
func RoleMiddleware(requireRole string) gin.HandlerFunc {
	return func(c *gin.Context) {

		//role, exists := c.Get("role")
		userIDInterface, exists := c.Get("userID")
		if !exists {
			response.Error(c, 400, "not user exist")
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"error": "not user exist",
			// })
			// c.Error(dto.AppError{
			// 	Code: 404,
			// 	Msg:  "user not found",
			// })
			c.Abort()
			return
		}
		userID := userIDInterface.(uint)
		fmt.Println("userID:", userID)
		role, err := repository.GetRoleByUserID(repository.DB, userID)

		if err != nil {
			response.Error(c, 500, "服务器内部错误")
			// c.JSON(http.StatusBadRequest, gin.H{
			// 	"error": "database wrong",
			// })
			// c.Error(dto.AppError{
			// 	Code: 500,
			// 	Msg:  "服务器内部错误",
			// 	Err:  err,
			// })
			c.Abort()
			return
		}
		if role != requireRole {
			response.Error(c, 401, "permission denied")
			// c.JSON(http.StatusForbidden, gin.H{
			// 	"error": "permission denied",
			// })
			c.Abort()
			return
		}
		c.Next()
	}
}

// 验证该用户是否是该task的创建者
func OperateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIInterface, exists := c.Get("userID")
		if !exists {
			response.Error(c, 400, "user not found")
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"error": "not user exist",
			// })
			c.Abort()
			return
		}
		userID := userIInterface.(uint)
		fmt.Println("userID:", userID)
		taskID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response.Error(c, 500, "服务器错误")
			c.Abort()
			return
		}
		task, error := repository.GetTaskByID(repository.DB, uint(taskID))
		if error != nil {
			response.Error(c, 404, "taskid doesn't exist")
			// c.JSON(http.StatusNotFound, gin.H{
			// 	"error": "taskid doesn't exist",
			// })
			c.Abort()
			return
		}
		if userID != task.UserID {
			response.Error(c, 401, "permission denied")
			// c.JSON(http.StatusForbidden, gin.H{
			// 	"error": "you haven't permission",
			// })
			c.Abort()
			return
		}
		c.Next()
	}
}

// func RequireRole(role string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		roleInterface, ok := c.Get("role")
// 		if !ok {
// 			response.Error(c, 400, "no role")
// 			c.Abort()
// 			return
// 		}
// 		r, ok := roleInterface.(string)
// 		if !ok || r != role {
// 			response.Error(c, 401, "permission denied")
// 			c.Abort()
// 			return
// 		}
// 		c.Next()

// 	}
// }
