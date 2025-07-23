package router

import (
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
)

func SetupRouter(taskController *controllers.TaskController) *gin.Engine {
	router := gin.Default()

	// Task routes
	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.GET("/", taskController.GetAllTasks)
		taskRoutes.GET("/:id", taskController.GetTask) 

		taskRoutes.POST("/", taskController.CreateTask)
		taskRoutes.PUT("/:id", taskController.UpdateTask)
		taskRoutes.DELETE("/:id", taskController.DeleteTask)
	}

	return router
}