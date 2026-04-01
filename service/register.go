package service

import (
	"errors"
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
		return repository.User{}, errors.New("username and password required")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		//utils.Error(c, 500, "failed to hasd password")
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"error": "failed to hash password",
		// })
		return repository.User{}, errors.New("failed to hash password")
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
		return repository.User{}, errors.New("user already exists")

	}
	return user, nil
}
