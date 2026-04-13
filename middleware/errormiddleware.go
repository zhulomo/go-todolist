package middleware

import (
	"go-gin-api/dto"
	"go-gin-api/response"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 捕获panic
				response.Error(c, 500, "服务器内部错误")

			}
		}()

		c.Next()

		// 处理 gin 收集的错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := err.(dto.AppError); ok {
				response.Error(c, appErr.Code, appErr.Msg)
				return
			}

			// 默认错误
			response.Error(c, 500, err.Error())
		}
	}
}
