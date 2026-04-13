package handler

import (
	"go-gin-api/dto"
	"go-gin-api/repository"
	"go-gin-api/response"
	"go-gin-api/service"
	"strconv"

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

// @Summary 新建任务
// @Description 用户新建任务
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param task body dto.CreateTaskRequest true "任务信息"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /tasks/{id} [post]
// @Security BearerAuth
func CreateTasks(c *gin.Context) {
	var req dto.CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		//utils.Error(c, 400, "invalid request")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "invalid  request",
		// })
		c.Error(dto.AppError{
			Code: 400,
			Msg:  "invalid request",
			Err:  err,
		})
		c.Abort()
		return
	}
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)

	if err := service.CreateTask(userID, req); err != nil {
		//utils.Error(c, 400, err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "can't create task",
		// })
		c.Error(err)
		c.Abort()
		return
	}
	response.Success(c, nil)
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "task created!",
	// })
}

// @Summary 获取所有任务
// @Description 获取所有用户的所有任务
// @Tags tasks
// @Produce json
// @Param page query int true "页数"
// @Param pageSize query int true "每页数量"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /tasks/get [get]
// @Security BearerAuth
// @Param role header string false "admin required"
func GetAllTasks(c *gin.Context) {

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

	tasks, total, err := service.GetAllTasks(page, pageSize)
	if err != nil {
		//utils.Error(c, 400, err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "出错了",
		// })
		c.Error(err)
		c.Abort()
		return
	}

	//c.JSON(http.StatusOK, tasks)
	//优化返回结构
	response.Success(c, gin.H{
		//c.JSON(http.StatusOK, gin.H{
		//"code":    200,
		//"message": "success",
		// "data": gin.H{
		// 	"list":     tasks,
		// 	"total":    total,
		// 	"page":     page,
		// 	"pageSize": pageSize,
		// },
		"list":     tasks,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func GetTaskById(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	task, err := repository.GetTaskByID(repository.DB, uint(id))
	if err != nil {
		//utils.Error(c, 400, "taskid doesn't exist")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "taskid doesn't exist",
		// })
		c.Error(dto.AppError{
			Code: 404,
			Msg:  "task does not exist",
			Err:  err,
		})
		c.Abort()
		return
	}
	response.Success(c, gin.H{
		"task": task,
	})
	//c.JSON(http.StatusOK, task)

}

// @Summary 更新任务
// @Description 更新当前用户的任务
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "任务ID"
// @Param task body dto.UpdateTaskRequest true "任务信息"
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /tasks/{id} [put]
// @Security BearerAuth
func UpdateTaskById(c *gin.Context) {
	var req dto.UpdateTaskRequest
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	if err := c.ShouldBindJSON(&req); err != nil {
		//utils.Error(c, 400, "invalid request")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "invalid  request",
		// })
		c.Error(dto.AppError{
			Code: 400,
			Msg:  "invalid request",
			Err:  err,
		})
		c.Abort()
		return
	}
	err := service.UpdateTask(uint(id), req)
	if err != nil {
		//utils.Error(c, 400, err.Error())
		// c.JSON(400, gin.H{
		// 	"error": err.Error(),
		// })
		c.Error(err)
		c.Abort()
		return
	}

	response.Success(c, nil)
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "update successed",
	// })

}

func TaskDelete(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	tasks, err := repository.GetTaskByID(repository.DB, uint(id))
	if err != nil {
		//utils.Error(c, 400, "task not found")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "task not found",
		// })
		c.Error(dto.AppError{
			Code: 404,
			Msg:  "task not found",
			Err:  err,
		})
		c.Abort()
		return
	}
	if err := repository.DeleteTask(repository.DB, tasks); err != nil {
		//utils.Error(c, 400, "can't delete")
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": "Can't delete",
		// })
		c.Error(dto.AppError{
			Code: 500,
			Msg:  "delete falied",
			Err:  err,
		})
		return
	}

	response.Success(c, nil)
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "delete  successed",
	// })
}
