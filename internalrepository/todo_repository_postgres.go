package repository

import (
	"database/sql"
	"fmt"

	domain "github.com/SenBedotcom/todo-api/internaldomain"
)

// PostgresTodoRepository implements TodoRepository interface using PostgreSQL
type PostgresTodoRepository struct {
	db *sql.DB
}

// NewPostgresTodoRepository creates a new PostgreSQL todo repository
func NewPostgresTodoRepository(db *sql.DB) *PostgresTodoRepository {
	return &PostgresTodoRepository{
		db: db,
	}
}

// Create creates a new todo in the database
func (r *PostgresTodoRepository) Create(todo *domain.Todo) error {
	query := `
		INSERT INTO todos (title, description, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		todo.Title,
		todo.Description,
		todo.Completed,
		todo.CreatedAt,
		todo.UpdatedAt,
	).Scan(&todo.ID)

	if err != nil {
		return fmt.Errorf("failed to create todo: %w", err)
	}

	return nil
}

// GetByID retrieves a todo by its ID
func (r *PostgresTodoRepository) GetByID(id int) (*domain.Todo, error) {
	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM todos
		WHERE id = $1
	`

	todo := &domain.Todo{}
	err := r.db.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	return todo, nil
}

// GetAll retrieves all todos from the database
func (r *PostgresTodoRepository) GetAll() ([]*domain.Todo, error) {
	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM todos
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get todos: %w", err)
	}
	defer rows.Close()

	var todos []*domain.Todo
	for rows.Next() {
		todo := &domain.Todo{}
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating todos: %w", err)
	}

	return todos, nil
}

// Update updates an existing todo
func (r *PostgresTodoRepository) Update(todo *domain.Todo) error {
	query := `
		UPDATE todos
		SET title = $1, description = $2, completed = $3, updated_at = $4
		WHERE id = $5
	`

	result, err := r.db.Exec(
		query,
		todo.Title,
		todo.Description,
		todo.Completed,
		todo.UpdatedAt,
		todo.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no todo found with id %d", todo.ID)
	}

	return nil
}

// Delete deletes a todo by its ID
func (r *PostgresTodoRepository) Delete(id int) error {
	query := `DELETE FROM todos WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no todo found with id %d", id)
	}

	return nil
}
