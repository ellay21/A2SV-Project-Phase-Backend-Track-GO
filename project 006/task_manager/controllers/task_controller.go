package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// i used IndentedJson instead of Json to make the response more readable since the project isn't for production just so you know bozos
type TaskController struct {
	TaskService *data.TaskService
}

func NewTaskController(taskService *data.TaskService) *TaskController {
	return &TaskController{
		TaskService: taskService,
	}
}
func (c *TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBind(&task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTask, err := c.TaskService.CTask(task, ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(201, gin.H{"message": "Task created successfully", "task": newTask})
}
func (c *TaskController) GetTask(ctx *gin.Context) {
	// Extracting ID from path parameter
	id := ctx.Param("id")
	task, err := c.TaskService.GTask(id, ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, task)
}

func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.TaskService.GUAllTasks(ctx)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err})
	}
	ctx.IndentedJSON(http.StatusOK, tasks)
}
func (c *TaskController) GetAllUserTasks(ctx *gin.Context) {
	tasks := c.TaskService.GAllTasks(ctx)
	ctx.IndentedJSON(http.StatusOK, tasks)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var task models.Task
	if err := ctx.ShouldBind(&task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.TaskService.UTask(task, id, ctx); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.TaskService.DTask(id, ctx); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
