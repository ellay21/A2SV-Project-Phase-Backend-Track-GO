package main
import (
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	taskService := data.NewTaskService()
	taskService.Initialize()

	taskController := controllers.NewTaskController(taskService)

	r := router.SetupRouter(taskController)
	
	// The server will run on http://localhost:8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}