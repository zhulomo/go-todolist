package handler

import (
	"go-gin-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminGetUsers(c *gin.Context) {

	var users []repository.Users
	repository.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
