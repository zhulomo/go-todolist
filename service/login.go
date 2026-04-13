package service

import (
	"go-gin-api/dto"
	"go-gin-api/repository"
	"go-gin-api/utils"
)

func Login(username, password string) (string, error) {
	if username == "" || password == "" {
		//utils.Error(c, 400, "username and password required")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "username and password required",
		//})
		return "", dto.AppError{Code: 400, Msg: "username and password required"}
	}

	user, err := repository.GetUserByUsername(repository.DB, username)
	if err != nil {
		//utils.Error(c, 400, "username does not exist")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "username does not exist",
		// })
		return "", dto.AppError{Code: 404, Msg: "user does not exist", Err: err}
	}
	if err := utils.CheckPassword(password, user.Password); err != nil {
		//utils.Error(c, 400, "password does not correct")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "password does not correct",
		// })
		return "", dto.AppError{Code: 401, Msg: "password incorrect", Err: err}
	}
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		//utils.Error(c, 500, "failed to generate token")
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"error": "failed to generate token",
		// })
		return "", dto.AppError{Code: 400, Msg: "failed to generate token", Err: err}
	}
	return token, nil

}
