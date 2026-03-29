package handler

import (
	"go-gin-api/repository"
	"go-gin-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	//绑定JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}
	//校验
	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username and password required",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := repository.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     "user",
	}

	//检查是否存在
	if err := repository.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user already exists",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user registered",
	})
}


