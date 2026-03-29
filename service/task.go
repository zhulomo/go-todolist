package service

import (
	"errors"
	"go-gin-api/dto"
	"go-gin-api/repository"
	"time"
)

func UpdateTask(id uint, req dto.UpdateTaskRequest) error {
	task, err := repository.GetTaskByID(repository.DB, id)
	if err != nil {
		return errors.New("task not found")
	}
	task.Title = req.Title
	task.Content = req.Content
	task.Status = req.Status
	task.UpdateAt = time.Now()
	if error := repository.UpdateTask(repository.DB, task); error != nil {
		return errors.New("update failed")
	}
	return nil

}
