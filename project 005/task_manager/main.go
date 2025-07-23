package main
import (
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	taskService, err := data.NewTaskService()
	if err != nil {
		panic(err)
	}

	taskController := controllers.NewTaskController(taskService)

	r := router.SetupRouter(taskController)
	
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}