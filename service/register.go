package service

import (
	"go-gin-api/dto"
	"go-gin-api/repository"
	"go-gin-api/utils"
)

func Register(username, password string) (repository.User, error) {
	//校验
	if username == "" || password == "" {
		//utils.Error(c, 400, "username and password required")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "username and password required",
		// })
		return repository.User{}, dto.AppError{Code: 400, Msg: "username and password required"}
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		//utils.Error(c, 500, "failed to hasd password")
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"error": "failed to hash password",
		// })
		return repository.User{}, dto.AppError{Code: 500, Msg: "failed to hash password", Err: err}
	}

	user := repository.User{
		Username: username,
		Password: hashedPassword,
		Role:     "user",
	}

	//检查是否存在
	if err := repository.DB.Create(&user).Error; err != nil {
		//utils.Error(c, 400, "user already exists")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "user already exists",
		// })
		return repository.User{}, dto.AppError{Code: 409, Msg: "user already exist", Err: err}

	}
	return user, nil
}
