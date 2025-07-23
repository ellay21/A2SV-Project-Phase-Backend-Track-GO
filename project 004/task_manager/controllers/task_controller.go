package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// i used IndentedJson instead of Json to make the response more readable sinc ethe project isn't for production jus so you know bozos
type TaskController struct {
	TaskService *data.TaskService
}
func NewTaskController(taskService *data.TaskService) *TaskController {
	return &TaskController{
		TaskService: taskService,
	}
}
func (c * TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBind(&task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.TaskService.CTask(task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Task created successfully", "task": task})
}
func (c *TaskController) GetTask(ctx *gin.Context) {
    // Extract ID from path parameter
    idStr := ctx.Param("id")
    
    // Convert string to int
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid task ID format"})
        return
    }
    
    task, err := c.TaskService.GTask(id)
    if err != nil {
        ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.IndentedJSON(http.StatusOK, task)
}

func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks := c.TaskService.GAllTasks()
	ctx.IndentedJSON(http.StatusOK, tasks)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
    
    // Convert string to int
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid task ID format"})
        return
    }
	var task models.Task
	if err := ctx.ShouldBind(&task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = id 
	if err := c.TaskService.UTask(task); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
    idStr := ctx.Param("id")
    
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid task ID format"})
        return
    }
    
    if err := c.TaskService.DTask(id); err != nil {
        ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}