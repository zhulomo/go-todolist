package repository

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Content  string
	Status   string `gorm:"default:todo"`
	UserID   uint
	CreateAt time.Time
	UpdateAt time.Time
}

func CreateTask(db *gorm.DB, task *Task) error {
	result := db.Create(task)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTasksByUserID(db *gorm.DB, userID uint) (*Task, error) {
	var task Task
	err := db.Where("UserID = ?", userID).Find(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func GetTaskByID(db *gorm.DB, id uint) (*Task, error) {
	var task Task
	err := db.Where("ID = ?", id).First(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func UpdateTask(db *gorm.DB, task *Task) error {
	err := db.Save(task).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(db *gorm.DB, task *Task) error {
	err := db.Delete(task).Error

	if err != nil {
		return err
	}

	return nil
}

func GetTasks(db *gorm.DB, page int, pageSize int) ([]Task, error, int64) {
	var tasks []Task
	var total int64
	offset := (page - 1) * pageSize
	result := db.Offset(offset).Limit(pageSize).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error, 0
	}
	result.Count(&total)
	return tasks, nil, total
}
