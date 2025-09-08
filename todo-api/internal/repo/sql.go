package repo

import (
	"database/sql"

	"github.com/example/todo-api/internal/model"
)

type SQLTodoRepository struct {
	DB *sql.DB
}

func NewSQLTodoRepository(db *sql.DB) *SQLTodoRepository {
	return &SQLTodoRepository{DB: db}
}

func (r *SQLTodoRepository) Create(todo *model.Todo) (string, error) {
	// Implementation for creating a todo in the SQL database
	return "", nil
}
func (r *SQLTodoRepository) List() ([]*model.Todo, error) {
	// Implementation for listing todos from the SQL database
	return nil, nil
}
func (r *SQLTodoRepository) Get(id int) (*model.Todo, error) {
	// Implementation for getting a todo by ID from the SQL database
	return nil, nil
}
func (r *SQLTodoRepository) Update(todo *model.Todo) error {
	// Implementation for updating a todo in the SQL database
	return nil
}
func (r *SQLTodoRepository) Delete(id int) error {
	// Implementation for deleting a todo by ID from the SQL database
	return nil
}
func (r *SQLTodoRepository) SetDone(id int, done bool) (*model.Todo, error) {
	// Implementation for setting a todo as done/undone in the SQL database
	return nil, nil
}
