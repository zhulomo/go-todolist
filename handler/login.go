package handler

import (
	"go-gin-api/dto"
	"go-gin-api/response"
	"go-gin-api/service"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户
// @Accept json
// Produce json
// @Param data body LoginRequest true "登录信息"
// Success 200 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {
	var login LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		//utils.Error(c, 400, "invalid request")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "invalid request",
		// })
		c.Error(dto.AppError{
			Code: 400,
			Msg:  "invalid request",
		})
		c.Abort()
		return
	}

	token, error := service.Login(login.Username, login.Password)
	if error != nil {
		c.Error(error)
		c.Abort()
		return
	}

	response.Success(c, token)
	// c.JSON(http.StatusOK, gin.H{
	// 	"token": token,
	// })

}
