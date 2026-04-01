package service

import (
	"errors"
	"go-gin-api/repository"
	"go-gin-api/utils"
)

func Login(username, password string) (string, error) {
	if username == "" || password == "" {
		//utils.Error(c, 400, "username and password required")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "username and password required",
		//})
		return "", errors.New("username and password required ")
	}

	user, err := repository.GetUserByUsername(repository.DB, username)
	if err != nil {
		//utils.Error(c, 400, "username does not exist")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "username does not exist",
		// })
		return "", errors.New("user does not exist")
	}
	if err := utils.CheckPassword(password, user.Password); err != nil {
		//utils.Error(c, 400, "password does not correct")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "password does not correct",
		// })
		return "", errors.New("password does not correct")
	}
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		//utils.Error(c, 500, "failed to generate token")
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"error": "failed to generate token",
		// })
		return "", errors.New("failed to generate token")
	}
	return token, nil

}
