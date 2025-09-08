package repo

import "github.com/example/todo-api/internal/model"

// CURD
// Create
// Update
// Read
// Delete

type TodoRepository interface {
	Create(todo *model.Todo) (string, error)
	List() ([]*model.Todo, error)
	Get(id int) (*model.Todo, error)
	Update(todo *model.Todo) error
	Delete(id int) error
	SetDone(id int, done bool) (*model.Todo, error)
}
