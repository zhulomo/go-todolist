package handler

import (
	"fmt"
	"go-gin-api/dto"
	"go-gin-api/repository"
	"go-gin-api/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//	type Tasks struct {
//		Title    string    `json:"title"`
//		Content  string    `json:"content"`
//		Status   string    `json:"status"`
//		UserID   uint      `json:"userid"`
//		CreateAt time.Time `json:"createAt"`
//		UpdateAt time.Time `json:"updateAt"`
//	}

func CreateTasks(c *gin.Context) {
	var req dto.CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid  request",
		})
		return
	}
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)
	task := repository.Task{
		Title:    req.Title,
		Content:  req.Content,
		Status:   req.Status,
		UserID:   userID,
		CreateAt: time.Now(),
	}
	if err := repository.CreateTask(repository.DB, &task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "can't create task",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "task created!",
	})
}

func GetAllTasks(c *gin.Context) {
	role, _ := c.Get("role")
	fmt.Println("role", role)
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "permission denied",
		})
		return
	}
	pageStr := c.Query("page")
	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		page = 1
	}
	pageSizestr := c.Query("pagesize")
	pageSize, _ := strconv.Atoi(pageSizestr)
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	tasks, err, total := repository.GetTasks(repository.DB, page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "出错了",
		})
		return
	}

	//c.JSON(http.StatusOK, tasks)
	//优化返回结构
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": gin.H{
			"list":     tasks,
			"total":    total,
			"page":     page,
			"pageSize": pageSize,
		},
	})
}

func GetTaskById(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	task, err := repository.GetTaskByID(repository.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "taskid doesn't exist",
		})
		return
	}
	c.JSON(http.StatusOK, task)

}

func UpdateTaskById(c *gin.Context) {
	var req dto.UpdateTaskRequest
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid  request",
		})
		return
	}
	err := service.UpdateTask(uint(id), req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update successed",
	})

}

func TaskDelete(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	tasks, err := repository.GetTaskByID(repository.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task not found",
		})
		return
	}
	if err := repository.DeleteTask(repository.DB, tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can't delete",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete  successed",
	})
}
