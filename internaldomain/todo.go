package domain

import "time"

// Todo represents a todo item entity
type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TodoRepository defines the interface for todo data operations
type TodoRepository interface {
	Create(todo *Todo) error
	GetByID(id int) (*Todo, error)
	GetAll() ([]*Todo, error)
	Update(todo *Todo) error
	Delete(id int) error
}
