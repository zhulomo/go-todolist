package handler

import (
	"go-gin-api/service"
	"go-gin-api/utils"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// @Summary 用户注册
// @Description 用户注册接口
// @Tags 用户
// @Accept json
// Produce json
// @Param data body RegisterRequest true "登录信息"
// Success 200 {object} map[string]string
// @Router /register [post]
func Register(c *gin.Context) {
	var req RegisterRequest

	//绑定JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, 400, "invalid request")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "invalid request",
		// })
		return
	}

	user, err := service.Register(req.Username, req.Password)
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}
	resp := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}

	utils.Success(c, gin.H{
		"user": resp,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "user registered",
	// })
}
