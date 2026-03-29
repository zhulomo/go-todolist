package repository

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Role     string `gorm:"default:user"`
}

func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user User

	err := db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}
func GetRoleByUserID(db *gorm.DB, userID uint) (string, error) {
	var user User

	err := db.Where("ID = ?", userID).First(&user).Error

	if err != nil {
		return "", err
	}
	return user.Role, nil
}
