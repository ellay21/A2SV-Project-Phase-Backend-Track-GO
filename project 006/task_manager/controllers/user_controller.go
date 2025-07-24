package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"task_manager/middleware"
)

type UserController struct {
	UserService *data.UserService
}

func NewUserController(userService *data.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser, err := c.UserService.CUser(user, ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	auth_middleware := middleware.NewAuthMiddleware()
	token, err := auth_middleware.GenerateJWT(newUser.Username, newUser.Role)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	ctx.IndentedJSON(201, gin.H{"message": "User created successfully", "user": newUser, "token": token})
}
func (c *UserController) Login(ctx *gin.Context) {
    var loginRequest struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Get the actual user from database
    authenticatedUser, err := c.UserService.AuthenticateUser(loginRequest.Username, loginRequest.Password)
    if err != nil {
        ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    auth_middleware := middleware.NewAuthMiddleware()
    token, err := auth_middleware.GenerateJWT(authenticatedUser.Username, authenticatedUser.Role)
    if err != nil {
        ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    ctx.IndentedJSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
        "user": gin.H{
            "username": authenticatedUser.Username,
            "role":    authenticatedUser.Role,
        },
    })
}
func (c *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.UserService.GUser(id, ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.UserService.GAllUsers(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, users)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.UserService.UUser(id, user, ctx); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
    id := ctx.Param("id")
    if err := c.UserService.DUser(id, ctx); err != nil {
        ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
