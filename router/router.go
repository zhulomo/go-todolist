package router

import (
	"go-gin-api/handler"
	middleeware "go-gin-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {

	r := gin.Default()

	protected := r.Group("/")
	//公共接口
	public := r.Group("/")
	{
		public.POST("/register", handler.Register)
		public.POST("/login", handler.Login)
		public.GET("/ping", handler.Ping)
	}
	//
	protected.Use(middleeware.AuthMiddleware())
	{
		protected.GET("/admin/users",
			middleeware.RoleMiddleware("admin"),
			handler.AdminGetUsers,
		)
		protected.POST("/tasks/create", handler.CreateTasks)
		protected.GET("/tasks/get", handler.GetAllTasks)
		protected.GET("/tasks/:id", middleeware.OperateMiddleware(),
			handler.GetTaskById)
		protected.PUT("/tasks/:id", middleeware.OperateMiddleware(),
			handler.UpdateTaskById)
		protected.DELETE("/tasks/:id", middleeware.OperateMiddleware(),
			handler.TaskDelete)

	}
	//r.GET("/ping", handler.Ping)
	// r.POST("/register", handler.Register)
	// r.GET("/users", handler.GetUsers)
	// r.POST("/login", handler.Login)
	return r
}
