package router

import (
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
	"task_manager/middleware"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
	router := gin.Default()
	authMiddleware := middleware.NewAuthMiddleware()
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(authMiddleware.Authenticate(), authMiddleware.Authorize("admin"))
	{
		adminRoutes.GET("/users", userController.GetAllUsers)
		adminRoutes.GET("/users/:id", userController.GetUser)
		adminRoutes.DELETE("/users/:id", userController.DeleteUser)
	}

	userRoutes := router.Group("/auth")
	{
		userRoutes.POST("/users",userController.CreateUser)
		userRoutes.POST("/login",userController.Login)
	}
	taskRoutes := router.Group("/tasks")
	taskRoutes.Use(authMiddleware.Authenticate(), authMiddleware.Authorize("admin", "user"))
	{
		taskRoutes.GET("/", taskController.GetAllTasks)
		taskRoutes.GET("/:id", taskController.GetTask)
		taskRoutes.POST("/", taskController.CreateTask)
		taskRoutes.PUT("/:id", taskController.UpdateTask)
		taskRoutes.DELETE("/:id", taskController.DeleteTask)
	}
	return router
}