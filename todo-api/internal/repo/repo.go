package repo

import (
	"context"

	"github.com/example/todo-api/internal/model"
)

// CURD
// Create
// Update
// Read
// Delete

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) (string, error)
	List(ctx context.Context) ([]*model.Todo, error)
	Get(ctx context.Context, id int) (*model.Todo, error)
	Update(ctx context.Context, todo *model.Todo) error
	Delete(ctx context.Context, id int) error
	SetDone(ctx context.Context, id int, done bool) (*model.Todo, error)
	RegisterUser(ctx context.Context, user *model.User) (int, error)
	AuthenticateUser(ctx context.Context, username, password string) (int, error)
}
