package repository

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error

	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	//如果表不存在就建表
	err = DB.AutoMigrate(&User{}, &Task{})
	if err != nil {
		return err
	}
	return nil
}
