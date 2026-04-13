package handler

import (
	"go-gin-api/repository"
	"go-gin-api/response"

	"github.com/gin-gonic/gin"
)

func AdminGetUsers(c *gin.Context) {

	var users []repository.User
	repository.DB.Find(&users)

	var resp []repository.Users
	for _, u := range users {
		resp = append(resp, repository.Users{
			ID:       u.ID,
			Username: u.Username,
			Role:     u.Role,
		})
	}
	//c.JSON(http.StatusOK, users)
	response.Success(c, resp)
}
