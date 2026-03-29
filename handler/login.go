package handler

import (
	"go-gin-api/repository"
	"go-gin-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var login LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	if login.Username == "" || login.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username and password required",
		})
		return
	}
	user, err := repository.GetUserByUsername(repository.DB, login.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username does not exist",
		})
		return
	}
	if err := utils.CheckPassword(login.Password, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password does not correct",
		})
		return
	}
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
