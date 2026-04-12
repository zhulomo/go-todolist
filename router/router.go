package router

import (
	"go-gin-api/handler"
	middleware "go-gin-api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouter() *gin.Engine {

	r := gin.Default()

	protected := r.Group("/")
	//公共接口
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	public := r.Group("/")
	{
		public.POST("/register", handler.Register)
		public.POST("/login", handler.Login)
		public.GET("/ping", handler.Ping)
	}
	//
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/admin/users",
			middleware.RoleMiddleware("admin"),
			handler.AdminGetUsers,
		)
		protected.POST("/tasks/create", handler.CreateTasks)
		protected.GET("/tasks/get", middleware.RequireRole("admin"),
			handler.GetAllTasks)
		protected.GET("/tasks/:id", middleware.OperateMiddleware(),
			handler.GetTaskById)
		protected.PUT("/tasks/:id", middleware.OperateMiddleware(),
			handler.UpdateTaskById)
		protected.DELETE("/tasks/:id", middleware.OperateMiddleware(),
			handler.TaskDelete)

	}
	//r.GET("/ping", handler.Ping)
	// r.POST("/register", handler.Register)
	// r.GET("/users", handler.GetUsers)
	// r.POST("/login", handler.Login)
	return r
}
