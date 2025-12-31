package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the HTTP routes
func SetupRouter(todoHandler *TodoHandler) *gin.Engine {
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API v1 routes
	api := router.Group("/api/v1")
	{
		// Todo routes
		api.GET("/todos", todoHandler.GetAllTodos)
		api.POST("/todos", todoHandler.CreateTodo)
		api.GET("/todos/:id", todoHandler.GetTodoByID)
		api.PUT("/todos/:id", todoHandler.UpdateTodo)
		api.DELETE("/todos/:id", todoHandler.DeleteTodo)
		api.PATCH("/todos/:id/toggle", todoHandler.ToggleTodoComplete)
	}

	return router
}
