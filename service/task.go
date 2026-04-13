package service

import (
	"go-gin-api/dto"
	"go-gin-api/repository"
	"time"
)

func UpdateTask(id uint, req dto.UpdateTaskRequest) error {
	task, err := repository.GetTaskByID(repository.DB, id)
	if err != nil {
		return dto.AppError{Code: 400, Msg: "task not found", Err: err}
	}
	task.Title = req.Title
	task.Content = req.Content
	task.Status = req.Status
	task.UpdateAt = time.Now()
	if error := repository.UpdateTask(repository.DB, task); error != nil {
		return dto.AppError{Code: 500, Msg: "update failed", Err: error}
	}
	return nil

}

func CreateTask(id uint, req dto.CreateTaskRequest) error {
	task := repository.Task{
		Title:    req.Title,
		Content:  req.Content,
		Status:   req.Status,
		UserID:   id,
		CreateAt: time.Now(),
	}
	if error := repository.CreateTask(repository.DB, &task); error != nil {
		return dto.AppError{Code: 500, Msg: "create task failed"}
	}
	return nil
}

func GetAllTasks(page, pageSize int) ([]repository.Task, int64, error) {
	tasks, err, total := repository.GetTasks(repository.DB, page, pageSize)
	if err != nil {
		return []repository.Task{}, 0, dto.AppError{Code: 500, Msg: "服务器出错了", Err: err}
	}
	return tasks, total, nil
}
