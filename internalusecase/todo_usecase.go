package usecase

import (
	"errors"
	"time"

	domain "github.com/SenBedotcom/todo-api/internaldomain"
)

var (
	ErrTodoNotFound    = errors.New("todo not found")
	ErrInvalidTodoData = errors.New("invalid todo data")
)

// TodoUseCase handles business logic for todos
type TodoUseCase struct {
	todoRepo domain.TodoRepository
}

// NewTodoUseCase creates a new todo use case
func NewTodoUseCase(todoRepo domain.TodoRepository) *TodoUseCase {
	return &TodoUseCase{
		todoRepo: todoRepo,
	}
}

// CreateTodo creates a new todo item
func (uc *TodoUseCase) CreateTodo(title, description string) (*domain.Todo, error) {
	if title == "" {
		return nil, ErrInvalidTodoData
	}

	todo := &domain.Todo{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := uc.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// GetTodoByID retrieves a todo by its ID
func (uc *TodoUseCase) GetTodoByID(id int) (*domain.Todo, error) {
	if id <= 0 {
		return nil, ErrInvalidTodoData
	}

	todo, err := uc.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if todo == nil {
		return nil, ErrTodoNotFound
	}

	return todo, nil
}

// GetAllTodos retrieves all todos
func (uc *TodoUseCase) GetAllTodos() ([]*domain.Todo, error) {
	return uc.todoRepo.GetAll()
}

// UpdateTodo updates an existing todo
func (uc *TodoUseCase) UpdateTodo(id int, title, description string, completed bool) (*domain.Todo, error) {
	if id <= 0 {
		return nil, ErrInvalidTodoData
	}

	existingTodo, err := uc.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if existingTodo == nil {
		return nil, ErrTodoNotFound
	}

	if title != "" {
		existingTodo.Title = title
	}
	existingTodo.Description = description
	existingTodo.Completed = completed
	existingTodo.UpdatedAt = time.Now()

	err = uc.todoRepo.Update(existingTodo)
	if err != nil {
		return nil, err
	}

	return existingTodo, nil
}

// DeleteTodo deletes a todo by its ID
func (uc *TodoUseCase) DeleteTodo(id int) error {
	if id <= 0 {
		return ErrInvalidTodoData
	}

	existingTodo, err := uc.todoRepo.GetByID(id)
	if err != nil {
		return err
	}

	if existingTodo == nil {
		return ErrTodoNotFound
	}

	return uc.todoRepo.Delete(id)
}

// ToggleTodoComplete toggles the completion status of a todo
func (uc *TodoUseCase) ToggleTodoComplete(id int) (*domain.Todo, error) {
	if id <= 0 {
		return nil, ErrInvalidTodoData
	}

	todo, err := uc.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if todo == nil {
		return nil, ErrTodoNotFound
	}

	todo.Completed = !todo.Completed
	todo.UpdatedAt = time.Now()

	err = uc.todoRepo.Update(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
