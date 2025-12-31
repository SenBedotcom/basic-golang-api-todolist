package http

import (
	"net/http"
	"strconv"

	usecase "github.com/SenBedotcom/todo-api/internalusecase"
	"github.com/gin-gonic/gin"
)

// TodoHandler handles HTTP requests for todos
type TodoHandler struct {
	todoUseCase *usecase.TodoUseCase
}

// NewTodoHandler creates a new todo handler
func NewTodoHandler(todoUseCase *usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}

// CreateTodoRequest represents the request body for creating a todo
type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// UpdateTodoRequest represents the request body for updating a todo
type UpdateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// CreateTodo handles POST /todos
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	todo, err := h.todoUseCase.CreateTodo(req.Title, req.Description)
	if err != nil {
		if err == usecase.ErrInvalidTodoData {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Todo created successfully",
		Data:    todo,
	})
}

// GetTodoByID handles GET /todos/:id
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	todo, err := h.todoUseCase.GetTodoByID(id)
	if err != nil {
		if err == usecase.ErrTodoNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todo"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Todo retrieved successfully",
		Data:    todo,
	})
}

// GetAllTodos handles GET /todos
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.todoUseCase.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todos"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Todos retrieved successfully",
		Data:    todos,
	})
}

// UpdateTodo handles PUT /todos/:id
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	var req UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	todo, err := h.todoUseCase.UpdateTodo(id, req.Title, req.Description, req.Completed)
	if err != nil {
		if err == usecase.ErrTodoNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err == usecase.ErrInvalidTodoData {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Todo updated successfully",
		Data:    todo,
	})
}

// DeleteTodo handles DELETE /todos/:id
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	err = h.todoUseCase.DeleteTodo(id)
	if err != nil {
		if err == usecase.ErrTodoNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Todo deleted successfully",
	})
}

// ToggleTodoComplete handles PATCH /todos/:id/toggle
func (h *TodoHandler) ToggleTodoComplete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	todo, err := h.todoUseCase.ToggleTodoComplete(id)
	if err != nil {
		if err == usecase.ErrTodoNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to toggle todo"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Todo toggled successfully",
		Data:    todo,
	})
}
